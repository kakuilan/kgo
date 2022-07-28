//go:build linux || darwin
// +build linux darwin

package kgo

import (
	"os"
	"syscall"
)

// IsProcessExists 进程是否存在.
func (ko *LkkOS) IsProcessExists(pid int) (res bool) {
	if pid > 0 {
		process, err := os.FindProcess(pid)
		if err == nil {
			if err = process.Signal(os.Signal(syscall.Signal(0))); err == nil {
				res = true
			}
		}
	}

	return
}
