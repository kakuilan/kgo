// +build linux darwin

package kgo

import (
	"golang.org/x/sys/unix"
)

// IsReadable 路径是否可读.
func (kf *LkkFile) IsReadable(fpath string) bool {
	err := unix.Access(fpath, unix.R_OK)
	if err != nil {
		return false
	}
	return true
}

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) bool {
	err := unix.Access(fpath, unix.W_OK)
	if err != nil {
		return false
	}
	return true
}

// IsExecutable 是否可执行文件.
func (kf *LkkFile) IsExecutable(fpath string) bool {
	err := unix.Access(fpath, unix.X_OK)
	if err != nil {
		return false
	}
	return true
}
