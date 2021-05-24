// +build linux

package kgo

import (
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
)

// MemoryUsage 获取内存使用率(仅支持linux),单位字节.
// 参数 virtual,是否取虚拟内存.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) MemoryUsage(virtual bool) (used, free, total uint64) {
	if virtual {
		// 虚拟机的内存
		contents, err := ioutil.ReadFile("/proc/meminfo")
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
			total = sysi.Totalram * uint64(syscall.Getpagesize()/1024)
			free = sysi.Freeram * uint64(syscall.Getpagesize()/1024)
			used = total - free
		}
	}

	return
}
