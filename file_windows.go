// +build windows

package kgo

import (
	"os"
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
