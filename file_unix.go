//go:build linux || darwin
// +build linux darwin

package kgo

import (
	"golang.org/x/sys/unix"
	"path/filepath"
	"strings"
)

// IsReadable 路径是否可读.
func (kf *LkkFile) IsReadable(fpath string) bool {
	err := unix.Access(fpath, unix.R_OK)
	return err == nil
}

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) bool {
	err := unix.Access(fpath, unix.W_OK)
	return err == nil
}

// IsExecutable 是否可执行文件.
func (kf *LkkFile) IsExecutable(fpath string) bool {
	err := unix.Access(fpath, unix.X_OK)
	return err == nil
}

// FormatPath 格式化路径.
func (kf *LkkFile) FormatPath(fpath string) string {
	if fpath == "" {
		return ""
	}

	fpath = formatPath(fpath)
	dir := filepath.Dir(fpath)

	if dir == `.` {
		return fpath
	}

	return strings.TrimRight(dir, "/") + "/" + filepath.Base(fpath)
}
