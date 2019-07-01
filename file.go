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

// FileSize get the length in bytes of the file.
func (kf *LkkFile) FileSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return f.Size()
}

// DirSize get the length in bytes of the directory.
func (kf *LkkFile) DirSize(path string) int64 {
	var size int64
	//filepath.Walk压测很慢
	_ = filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size
}

// IsExist determines whether the path is exists.
func (kf *LkkFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Writeable determines whether the path is writeable.
func (kf *LkkFile) IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

// IsReadable determines whether the path is readable.
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

// IsDir determines whether the path is a directory.
func (kf *LkkFile) IsDir(path string) bool {
	f, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	} else if nil != err {
		return false
	}
	return f.IsDir()
}

// IsBinary determines whether the content is a binary file content.
// TODO 转到字符串中
func (kf *LkkFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsImg determines whether the path is a image.
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

// CopyFile copies the source file to the dest file,cover is enum(FCOVER_ALLOW、FCOVER_IGNORE、FCOVER_DENY)
func (kf *LkkFile) CopyFile(source string, dest string, cover LkkFileCover) (int64, error) {
	if(source == dest) {
		return 0, nil
	}

	sourceStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	}else if !sourceStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", source)
	}

	if cover != FCOVER_ALLOW {
		if _, err := os.Stat(dest); err == nil {
			if cover == FCOVER_IGNORE {
				return 0, nil
			}else if cover == FCOVER_DENY {
				return 0, fmt.Errorf("File %s already exists", dest)
			}
		}
	}

	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer destFile.Close()

	var nBytes int64
	sourceSize := sourceStat.Size()
	if(sourceSize <= 1048576) { //1M以内小文件使用buffer拷贝
		var total int  =0
		var bufferSize int = 102400
		if sourceSize < 524288 {
			bufferSize = 51200
		}

		buf := make([]byte, bufferSize)
		for {
			n, err := sourceFile.Read(buf)
			if err != nil && err != io.EOF {
				return int64(total), err
			}else if n == 0 {
				break
			}

			if _, err := destFile.Write(buf[:n]); err != nil {
				return int64(total), err
			}

			total += n
		}
		nBytes = int64(total)
	}else{
		nBytes, err = io.Copy(destFile, sourceFile)
		if err == nil {
			err = os.Chmod(dest, sourceStat.Mode())
		}
	}

	return nBytes, err
}

// FastCopy fast copies the source file to the dest file,no safety check.
func (kf *LkkFile) FastCopy(source string, dest string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}

	var bufferSize int = 32768
	var nBytes int = 0
	buf := make([]byte, bufferSize)
	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			return int64(nBytes), err
		}else if n == 0 {
			break
		}

		if _, err := destFile.Write(buf[:n]); err != nil {
			return int64(nBytes), err
		}

		nBytes += n
	}

	return int64(nBytes), err
}

// CopyLink copy link file.
func (kf *LkkFile) CopyLink(source string, dest string) error {
	if(source == dest) {
		return nil
	}

	source, err := os.Readlink(source)
	if err != nil {
		return err
	}
	return os.Symlink(source, dest)
}