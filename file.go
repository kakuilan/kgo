package kgo

import (
	"bufio"
	"io"
	"os"
	"path"
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

// ReadFirstLine 读取文件首行.
func (kf *LkkFile) ReadFirstLine(fpath string) []byte {
	var res []byte
	fh, err := os.Open(fpath)
	if err == nil {
		scanner := bufio.NewScanner(fh)
		for scanner.Scan() {
			res = scanner.Bytes()
			break
		}
	}
	defer func() {
		_ = fh.Close()
	}()

	return res
}

// ReadLastLine 读取文件末行.
func (kf *LkkFile) ReadLastLine(fpath string) []byte {
	var res []byte
	file, err := os.Open(fpath)
	if err == nil {
		var lastLineSize int
		reader := bufio.NewReader(file)

		for {
			bs, err := reader.ReadBytes('\n')
			lastLineSize = len(bs)
			if err == io.EOF {
				break
			}
		}

		fileInfo, _ := os.Stat(fpath)

		// make a buffer size according to the lastLineSize
		buffer := make([]byte, lastLineSize)
		offset := fileInfo.Size() - int64(lastLineSize)
		numRead, _ := file.ReadAt(buffer, offset)
		res = buffer[:numRead]
	}
	defer func() {
		_ = file.Close()
	}()

	return res
}

// WriteFile 将内容写入文件.
// fpath为文件路径;data为内容;perm为权限,默认为0655.
func (kf *LkkFile) WriteFile(fpath string, data []byte, perm ...os.FileMode) error {
	if dir := path.Dir(fpath); dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	var p os.FileMode = 0655
	if len(perm) > 0 {
		p = perm[0]
	}

	return os.WriteFile(fpath, data, p)
}
