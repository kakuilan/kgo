// +build linux

package kgo

import (
	"io/ioutil"
	"strconv"
	"strings"
	"syscall"
)

// MemoryUsage 获取内存使用率,单位字节.
// 参数 virtual(仅支持linux),是否取虚拟内存.
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
	contents, _ := ioutil.ReadFile("/proc/stat")
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
