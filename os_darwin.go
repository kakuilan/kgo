//go:build darwin
// +build darwin

package kgo

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/sys/unix"
	"os/exec"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// cachedBootTime must be accessed via atomic.Load/StoreUint64
var cachedBootTime uint64

//系统IO信息
var cacheIOInfos []byte

// bootTime 获取系统启动时间,秒.
func bootTime() (uint64, error) {
	t := atomic.LoadUint64(&cachedBootTime)
	if t != 0 {
		return t, nil
	}
	tv, err := unix.SysctlTimeval("kern.boottime")
	if err != nil {
		return 0, err
	}

	atomic.StoreUint64(&cachedBootTime, uint64(tv.Sec))

	return uint64(tv.Sec), nil
}

// getIOInfos 获取系统IO信息
func (ko *LkkOS) getIOInfos() []byte {
	if len(cacheIOInfos) == 0 {
		_, cacheIOInfos, _ = ko.System("ioreg -l")
	}
	return cacheIOInfos
}

// MemoryUsage 获取内存使用率,单位字节.
// 参数 virtual(仅支持linux),是否取虚拟内存.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) MemoryUsage(virtual bool) (used, free, total uint64) {
	totalStr, err := unix.Sysctl("hw.memsize")
	if err != nil {
		return
	}

	vm_stat, err := exec.LookPath("vm_stat")
	if err != nil {
		return
	}

	_, out, _ := ko.Exec(vm_stat)
	lines := strings.Split(string(out), "\n")
	pagesize := uint64(unix.Getpagesize())
	var inactive uint64
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.Trim(fields[1], " .")
		switch key {
		case "Pages free":
			f, e := strconv.ParseUint(value, 10, 64)
			if e != nil {
				err = e
			}
			free = f * pagesize
		case "Pages inactive":
			ina, e := strconv.ParseUint(value, 10, 64)
			if e != nil {
				err = e
			}
			inactive = ina * pagesize
		}
	}

	// unix.sysctl() helpfully assumes the result is a null-terminated string and
	// removes the last byte of the result if it's 0 :/
	totalStr += "\x00"
	total = uint64(binary.LittleEndian.Uint64([]byte(totalStr)))
	used = total - (free + inactive)
	return
}

// DiskUsage 获取磁盘(目录)使用情况,单位字节.参数path为路径.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) DiskUsage(path string) (used, free, total uint64) {
	stat := unix.Statfs_t{}
	err := unix.Statfs(path, &stat)
	if err != nil {
		return
	}

	total = uint64(stat.Blocks) * uint64(stat.Bsize)
	free = uint64(stat.Bavail) * uint64(stat.Bsize)
	used = (uint64(stat.Blocks) - uint64(stat.Bfree)) * uint64(stat.Bsize)
	return
}

// Uptime 获取系统运行时间,秒.
func (ko *LkkOS) Uptime() (uint64, error) {
	boot, err := bootTime()
	if err != nil {
		return 0, err
	}

	res := uint64(time.Now().Unix()) - boot
	return res, nil
}

// GetBiosInfo 获取BIOS信息.
// 注意:Mac机器没有BIOS信息,它使用EFI.
func (ko *LkkOS) GetBiosInfo() *BiosInfo {
	res := &BiosInfo{
		Vendor:  "",
		Version: "",
		Date:    "",
	}

	infos := ko.getIOInfos()
	if len(infos) > 0 {
		infoStr := string(infos)
		res.Version = trim(KStr.GetEquationValue(infoStr, "SMBIOS-EPS"), "<", ">", `"`, `'`)
	}

	return res
}

// GetBoardInfo 获取Board信息.
func (ko *LkkOS) GetBoardInfo() *BoardInfo {
	res := &BoardInfo{
		Name:     "",
		Vendor:   "",
		Version:  "",
		Serial:   "",
		AssetTag: "",
	}

	infos := ko.getIOInfos()
	if len(infos) > 0 {
		infoStr := string(infos)
		res.Name = trim(KStr.GetEquationValue(infoStr, "product-name"), "<", ">", `"`, `'`)
		res.Version = trim(KStr.GetEquationValue(infoStr, "board-id"), "<", ">", `"`, `'`)
		res.Serial = trim(KStr.GetEquationValue(infoStr, "IOPlatformUUID"), "<", ">", `"`, `'`)
	}

	return res
}

// GetCpuInfo 获取CPU信息.
func (ko *LkkOS) GetCpuInfo() *CpuInfo {
	var res = &CpuInfo{
		Vendor:  "",
		Model:   "",
		Speed:   "",
		Cache:   0,
		Cpus:    0,
		Cores:   0,
		Threads: 0,
	}

	res.Model, _ = unix.Sysctl("machdep.cpu.brand_string")
	res.Vendor, _ = unix.Sysctl("machdep.cpu.vendor")

	cacheSize, _ := unix.SysctlUint32("machdep.cpu.cache.size")
	cpus, _ := unix.SysctlUint32("hw.physicalcpu")
	cores, _ := unix.SysctlUint32("machdep.cpu.core_count")
	threads, _ := unix.SysctlUint32("machdep.cpu.thread_count")
	res.Cache = uint(cacheSize)
	res.Cpus = uint(cpus)
	res.Cores = uint(cores)
	res.Threads = uint(threads)

	// Use the rated frequency of the CPU. This is a static value and does not
	// account for low power or Turbo Boost modes.
	cpuFrequency, err := unix.SysctlUint64("hw.cpufrequency")
	if err == nil {
		speed := float64(cpuFrequency) / 1000000.0
		res.Speed = KNum.NumberFormat(speed, 2, ".", "")
	}

	return res
}

// GetPidByPort 根据端口号获取监听的进程PID.
// linux可能要求root权限;
// darwin依赖lsof;
// windows依赖netstat.
func (ko *LkkOS) GetPidByPort(port int) (pid int) {
	command := fmt.Sprintf("lsof -i tcp:%d", port)
	_, out, _ := ko.System(command)
	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) >= 9 && isNumeric(fields[1]) {
			p := toInt(fields[1])
			if p > 0 {
				pid = p
				break
			}
		}
	}

	return
}
