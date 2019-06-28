package gohelper

import (
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

// IsExist determines whether the path spcified by the given is exists.
func(* LkkFile)  IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// GetExt get the file extension without a dot.
func(* LkkFile) GetExt(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}

	return suffix
}

// GetSize get the length in bytes of file of the specified path.
func(* LkkFile) GetSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}

	return f.Size()
}

// Writeable determines whether the path spcified by the given path is writeable.
func(* LkkFile) IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}

	return true
}
