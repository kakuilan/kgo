package gohelper

import (
	"io"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
)

// GetExt get the file extension without a dot.
func (kf *LkkFile) GetExt(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// GetSize get the length in bytes of file of the specified path.
func (kf *LkkFile) GetSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return f.Size()
}

// IsExist determines whether the path spcified by the given is exists.
func (kf *LkkFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Writeable determines whether the path spcified by the given path is writeable.
func (kf *LkkFile) IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

// IsReadable determines whether the path spcified by the given path is readable.
func (kf *LkkFile) IsReadable(path string) bool {
	err := syscall.Access(path, syscall.O_RDONLY)
	if err != nil {
		return false
	}
	return true
}

// IsFile returns true if path exists and is a file (or a link to a file) and false otherwise
func (kf *LkkFile) IsFile(path string) bool {
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
	} else if nil != err {
		return false
	}
	return f.IsDir()
}

// IsBinary determines whether the specified content is a binary file content.
func (kf *LkkFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsImg determines whether the specified path is a image.
func (kf *LkkFile) IsImg(path string) bool {
	ext := kf.GetExt(path)
	switch ext {
	case "jpg", "jpeg", "bmp", "gif", "png", "svg", "ico":
		return true
	default:
		return false
	}
}

// AbsPath returns an absolute representation of path. Works like filepath.Abs
func (kf *LkkFile) AbsPath(path string) string {
	fullPath := ""
	res, err := filepath.Abs(path)
	if err != nil {
		fullPath = filepath.Join("/", path)
	} else {
		fullPath = res
	}
	return fullPath
}

// CopyFile copies the source file to the dest file.
func (kf *LkkFile) CopyFile(source string, dest string) (int64, error) {
	sourceFileStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	}else if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", source)
	}

	sourcefile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destfile.Close()

	nBytes, err := io.Copy(destfile, sourcefile)
	if err == nil {
		err = os.Chmod(dest, sourceFileStat.Mode())
	}

	return nBytes, err
}