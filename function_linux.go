// +build linux

package kgo

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

// getPidByInode 根据套接字的inode获取PID.须root权限.
func getPidByInode(inode string, procDirs []string) (pid int) {
	if len(procDirs) == 0 {
		procDirs, _ = filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	}

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
func getProcessPathByPid(pid int) string {
	exe := fmt.Sprintf("/proc/%d/exe", pid)
	path, _ := os.Readlink(exe)
	return path
}
