package gohelper

import (
	"os"
	"path/filepath"
	"strings"
)

// IsExist determines whether the file spcified by the given path is exists.
func(* LkkFile)  IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
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

