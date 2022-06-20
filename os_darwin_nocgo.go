// +build darwin
// +build !cgo

package kgo

import (
	"fmt"
	"os/exec"
	"strings"
)

// CpuUsage 获取CPU使用率(darwin系统必须使用cgo),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (ko *LkkOS) CpuUsage() (user, idle, total uint64) {
	//CPU counters for darwin is unavailable without cgo
	return
}

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) (res string) {
	lsof, err := exec.LookPath("lsof")
	if err != nil {
		return ""
	}
	command := fmt.Sprintf("%s -p %d -Fpfn", lsof, pid)
	_, out, _ := KOS.System(command)
	txtFound := 0
	lines := strings.Split(string(out), "\n")
	for i := 1; i < len(lines); i++ {
		if lines[i] == "ftxt" {
			txtFound++
			if txtFound == 2 {
				return lines[i-1][1:]
			}
		}
	}

	return ""
}
