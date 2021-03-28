// +build linux darwin

package kgo

import (
	"golang.org/x/sys/unix"
)

// IsReadable 路径是否可读.
func (kf *LkkFile) IsReadable(fpath string) bool {
	err := unix.Access(fpath, unix.O_RDONLY)
	if err != nil {
		return false
	}
	return true
}

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) bool {
	err := unix.Access(fpath, unix.O_RDWR)
	if err != nil {
		return false
	}
	return true
}
