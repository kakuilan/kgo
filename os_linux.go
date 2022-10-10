//go:build linux
// +build linux

package kgo

import (
	"bufio"
	"fmt"
	"golang.org/x/sys/unix"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)

// getPidByInode 根据套接字的inode获取PID.须root权限.
func getPidByInode(inode string, procDirs []string) (pid int) {
	re := regexp.MustCompile(inode)
	for _, item := range procDirs {
		path, _ := os.Readlink(item)
		out := re.FindString(path)
		if len(out) != 0 {
			pid, _ = strconv.Atoi(strings.Split(item, "/")[2])
			break
		}
	}

	return pid
}

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) (res string) {
	exe := fmt.Sprintf("/proc/%d/exe", pid)
	res, _ = os.Readlink(exe)

	return
}

// MemoryUsage 获取内存使用率,单位字节.
// 参数 virtual(仅支持linux),是否取虚拟内存.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) MemoryUsage(virtual bool) (used, free, total uint64) {
	if virtual {
		// 虚拟机的内存
		contents, err := os.ReadFile("/proc/meminfo")
		if err == nil {
			lines := strings.Split(string(contents), "\n")
			for _, line := range lines {
				fields := strings.Fields(line)
				if len(fields) == 3 {
					val, _ := strconv.ParseUint(fields[1], 10, 64) // kB

					if strings.HasPrefix(fields[0], "MemTotal") {
						total = val * 1024
					} else if strings.HasPrefix(fields[0], "MemFree") {
						free = val * 1024
					}
				}
			}

			//计算已用内存
			used = total - free
		}
	} else {
		// 真实物理机内存
		sysi := &syscall.Sysinfo_t{}
		err := syscall.Sysinfo(sysi)
		if err == nil {
			total = sysi.Totalram * uint64(sysi.Unit)
			free = sysi.Freeram * uint64(sysi.Unit)
			used = total - free
		}
	}

	return
}

// CpuUsage 获取CPU使用率(darwin系统必须使用cgo),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (ko *LkkOS) CpuUsage() (user, idle, total uint64) {
	contents, _ := os.ReadFile("/proc/stat")
	if len(contents) > 0 {
		lines := strings.Split(string(contents), "\n")
		for _, line := range lines {
			fields := strings.Fields(line)
			if fields[0] == "cpu" {
				//CPU指标：user，nice, system, idle, iowait, irq, softirq
				// cpu  130216 19944 162525 1491240 3784 24749 17773 0 0 0

				numFields := len(fields)
				for i := 1; i < numFields; i++ {
					val, _ := strconv.ParseUint(fields[i], 10, 64)
					total += val // tally up all the numbers to get total ticks
					if i == 1 {
						user = val
					} else if i == 4 { // idle is the 5th field in the cpu line
						idle = val
					}
				}
				break
			}
		}
	}

	return
}

// DiskUsage 获取磁盘(目录)使用情况,单位字节.参数path为路径.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) DiskUsage(path string) (used, free, total uint64) {
	fs := &syscall.Statfs_t{}
	err := syscall.Statfs(path, fs)
	if err == nil {
		total = fs.Blocks * uint64(fs.Bsize)
		free = fs.Bfree * uint64(fs.Bsize)
		used = total - free
	}

	return
}

// Uptime 获取系统运行时间,秒.
func (ko *LkkOS) Uptime() (res uint64, err error) {
	info := &unix.Sysinfo_t{}
	err = unix.Sysinfo(info)
	if err == nil {
		res = uint64(info.Uptime)
	}

	return
}

// GetBiosInfo 获取BIOS信息.
// 注意:Mac机器没有BIOS信息,它使用EFI.
func (ko *LkkOS) GetBiosInfo() *BiosInfo {
	return &BiosInfo{
		Vendor:  strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/bios_vendor"))),
		Version: strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/bios_version"))),
		Date:    strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/bios_date"))),
	}
}

// GetBoardInfo 获取Board信息.
func (ko *LkkOS) GetBoardInfo() *BoardInfo {
	return &BoardInfo{
		Name:     strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/board_name"))),
		Vendor:   strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/board_vendor"))),
		Version:  strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/board_version"))),
		Serial:   strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/board_serial"))),
		AssetTag: strings.TrimSpace(string(KFile.ReadFirstLine("/sys/class/dmi/id/board_asset_tag"))),
	}
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

	res.Threads = uint(runtime.NumCPU())
	f, err := os.Open("/proc/cpuinfo")
	if err == nil {
		cpu := make(map[string]bool)
		core := make(map[string]bool)
		var cpuID string

		s := bufio.NewScanner(f)
		for s.Scan() {
			if sl := cpuRegTwoColumns.Split(s.Text(), 2); sl != nil {
				switch sl[0] {
				case "physical id":
					cpuID = sl[1]
					cpu[cpuID] = true
				case "core id":
					coreID := fmt.Sprintf("%s/%s", cpuID, sl[1])
					core[coreID] = true
				case "vendor_id":
					if res.Vendor == "" {
						res.Vendor = sl[1]
					}
				case "model name":
					if res.Model == "" {
						// CPU model, as reported by /proc/cpuinfo, can be a bit ugly. Clean up...
						model := cpuRegExtraSpace.ReplaceAllLiteralString(sl[1], " ")
						res.Model = strings.Replace(model, "- ", "-", 1)
					}
				case "cpu MHz":
					if res.Speed == "" {
						res.Speed = sl[1]
					}
				case "cache size":
					if res.Cache == 0 {
						if m := cpuRegCacheSize.FindStringSubmatch(sl[1]); m != nil {
							if cache, err := strconv.ParseUint(m[1], 10, 64); err == nil {
								res.Cache = uint(cache)
							}
						}
					}
				}
			}
		}

		res.Cpus = uint(len(cpu))
		res.Cores = uint(len(core))
	}
	defer func() {
		_ = f.Close()
	}()

	return res
}

// GetPidByPort 根据端口号获取监听的进程PID.
// linux可能要求root权限;
// darwin依赖lsof;
// windows依赖netstat.
func (ko *LkkOS) GetPidByPort(port int) (pid int) {
	files := []string{
		"/proc/net/tcp",
		"/proc/net/udp",
		"/proc/net/tcp6",
		"/proc/net/udp6",
	}

	procDirs, _ := filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	for _, fpath := range files {
		lines, _ := KFile.ReadInArray(fpath)
		for _, line := range lines[1:] {
			fields := strings.Fields(line)
			if len(fields) < 10 {
				continue
			}

			//非 LISTEN 监听状态
			if fields[3] != "0A" {
				continue
			}

			//本地ip和端口
			ipport := strings.Split(fields[1], ":")
			locPort, _ := KConv.Hex2Dec(ipport[1])

			// 非该端口
			if int(locPort) != port {
				continue
			}

			pid = getPidByInode(fields[9], procDirs)
			if pid > 0 {
				return
			}
		}
	}

	return
}
