package kgo

import (
	"os"
	"path/filepath"
	"strings"
)

// GetExt 获取文件的小写扩展名,不包括点"." .
func (kf *LkkFile) GetExt(fpath string) string {
	suffix := filepath.Ext(fpath)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// ReadFile 读取文件内容.
func (kf *LkkFile) ReadFile(fpath string) ([]byte, error) {
	data, err := os.ReadFile(fpath)
	return data, err
}

// ReadInArray 把整个文件读入一个数组中,每行作为一个元素.
func (kf *LkkFile) ReadInArray(fpath string) ([]string, error) {
	data, err := os.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
