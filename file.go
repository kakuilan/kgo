package gohelper

import (
	"os"
)

// IsExist determines whether the file spcified by the given path is exists.
func(* LkkFile)  IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	return err == nil || os.IsExist(err)
}