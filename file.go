package gohelper

import (
	"io"
	"os"
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
)

// GetExt 获取文件扩展名,不包括"."
func (kf *LkkFile) GetExt(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// GetSize 获取文件大小(bytes字节)
func (kf *LkkFile) GetSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return f.Size()
}

// IsExist 路径(文件/目录)是否存在
func (kf *LkkFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// Writeable 路径是否可写
func (kf *LkkFile) IsWritable(path string) bool {
	err := syscall.Access(path, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

// IsReadable 路径是否可读
func (kf *LkkFile) IsReadable(path string) bool {
	err := syscall.Access(path, syscall.O_RDONLY)
	if err != nil {
		return false
	}
	return true
}

// IsFile 是否正常的文件(或文件链接)
func (kf *LkkFile) IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}

// IsDir 是否目录
func (kf *LkkFile) IsDir(path string) bool {
	f, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	} else if nil != err {
		return false
	}
	return f.IsDir()
}

// IsBinary 是否二进制文件
func (kf *LkkFile) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsImg 是否图片文件
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

// CopyFile 拷贝源文件到目标文件,cover为枚举(FCOVER_ALLOW、FCOVER_IGNORE、FCOVER_DENY)
func (kf *LkkFile) CopyFile(source string, dest string, cover LkkFileCover) (int64, error) {
	if cover != FCOVER_ALLOW {
		if _, err := os.Stat(dest); err ==nil {
			if cover == FCOVER_IGNORE {
				return 0, nil
			}else if cover == FCOVER_DENY {
				return 0, fmt.Errorf("File %s already exists.", dest)
			}
		}
	}

	sourceStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	}else if !sourceStat.Mode().IsRegular() {
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
		err = os.Chmod(dest, sourceStat.Mode())
	}

	return nBytes, err
}


func (kf *LkkFile) FastCopy(source string, dest string, cover LkkFileCover) (int64, error) {

}