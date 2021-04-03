// +build windows

package kgo

import (
	"os"
	"path/filepath"
	"strings"
)

// IsReadable 路径是否可读.
func (kf *LkkFile) IsReadable(fpath string) (res bool) {
	info, err := os.Stat(fpath)

	if err != nil {
		return
	}

	if info.Mode().Perm()&(1<<(uint(8))) == 0 {
		return
	}

	res = true
	return
}

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) (res bool) {
	info, err := os.Stat(fpath)

	if err != nil {
		return
	}

	if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return
	}

	res = true
	return
}

// IsExecutable 是否可执行文件.
func (kf *LkkFile) IsExecutable(fpath string) bool {
	info, err := os.Stat(fpath)
	return err == nil && info.Mode().IsRegular() && (info.Mode()&0111) != 0
}

// FormatPath 格式化路径.
func (kf *LkkFile) FormatPath(fpath string) string {
	if fpath == "" {
		return ""
	}

	fpath = formatPath(fpath)
	dir := formatPath(filepath.Dir(fpath))

	if dir == `.` {
		return fpath
	}

	return strings.TrimRight(dir, "/") + "/" + filepath.Base(fpath)
}
