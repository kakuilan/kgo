// +build linux darwin

package kgo

import (
	"golang.org/x/sys/unix"
)

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) bool {
	err := unix.Access(fpath, unix.O_RDWR)
	if err != nil {
		return false
	}
	return true
}
