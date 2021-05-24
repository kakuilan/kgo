// +build darwin

package kgo

import (
	"encoding/binary"
	"golang.org/x/sys/unix"
	"os/exec"
	"strconv"
	"strings"
)

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

	ret, out, _ := ko.Exec(vm_stat)
	lines := strings.Split(string(out), "\n")
	pagesize := uint64(unix.Getpagesize())
	var vm_free, vm_inactive uint64
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) < 2 {
			continue
		}
		key := strings.TrimSpace(fields[0])
		value := strings.Trim(fields[1], " .")
		switch key {
		case "Pages free":
			free, e := strconv.ParseUint(value, 10, 64)
			if e != nil {
				err = e
			}
			vm_free = free * pagesize
		case "Pages inactive":
			inactive, e := strconv.ParseUint(value, 10, 64)
			if e != nil {
				err = e
			}
			vm_inactive = inactive * pagesize
		}
	}

	// unix.sysctl() helpfully assumes the result is a null-terminated string and
	// removes the last byte of the result if it's 0 :/
	totalStr += "\x00"
	total = uint64(binary.LittleEndian.Uint64([]byte(totalStr)))
	used = total - (vm_free + vm_inactive)
	return
}
