// +build windows

package kgo

import (
	"fmt"
	"github.com/StackExchange/wmi"
)

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) (res string) {
	var dst []Win32_Process
	query := fmt.Sprintf("WHERE ProcessId = %d", pid)
	q := wmi.CreateQuery(&dst, query)
	err := wmi.Query(q, &dst)
	if err == nil && len(dst) > 0 {
		res = *dst[0].ExecutablePath
	}

	return
}
