// +build windows

package kgo

import (
	"os"
)

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) (res bool) {
	info, err := os.Stat(fpath)
	if err != nil {
		return
	} else if info.Mode().Perm()&(1<<(uint(7))) == 0 {
		return
	}

	res = true
	return
}
