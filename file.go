package gohelper

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// GetExt get the file extension without a dot.
func(kf *LkkFile) GetExt(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// GetSize get the length in bytes of file of the specified path.
func(kf *LkkFile) GetSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return f.Size()
}

// IsExist determines whether the path spcified by the given is exists.
func(kf *LkkFile)  IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Writeable determines whether the path spcified by the given path is writeable.
func(kf *LkkFile) IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

// IsReadable determines whether the path spcified by the given path is readable.
func(kf *LkkFile) IsReadable(path string) bool {
	if !kf.IsExist(path) {
		return false
	}
	return syscall.Access(path, syscall.O_RDONLY) == nil
}

// IsFile returns true if path exists and is a file (or a link to a file) and false otherwise
func(kf *LkkFile) IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}

// IsDir determines whether the specified path is a directory.
func (kf *LkkFile) IsDir(path string) bool {
	f, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}else if nil != err {
		return false
	}

	return f.IsDir()
}