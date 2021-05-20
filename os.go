package kgo

import (
	"os"
	"path/filepath"
	"runtime"
)

// IsWindows 当前操作系统是否Windows.
func (ko *LkkOS) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux 当前操作系统是否Linux.
func (ko *LkkOS) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsMac 当前操作系统是否Mac OS/X.
func (ko *LkkOS) IsMac() bool {
	return "darwin" == runtime.GOOS
}

// Pwd 获取当前程序运行所在的路径,注意和Getwd有所不同.
// 若当前执行的是链接文件,则会指向真实二进制程序的所在目录.
func (ko *LkkOS) Pwd() string {
	var dir, ex string
	var err error
	ex, err = os.Executable()
	if err == nil {
		exReal, _ := filepath.EvalSymlinks(ex)
		exReal, _ = filepath.Abs(exReal)
		dir = filepath.Dir(exReal)
	}

	return dir
}
