package kgo

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
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
	fh, err := os.Open(fpath)
	if err == nil {
		var lastLineSize int
		reader := bufio.NewReader(fh)

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
		numRead, _ := fh.ReadAt(buffer, offset)
		res = buffer[:numRead]
	}
	defer func() {
		_ = fh.Close()
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

// GetFileMode 获取路径的权限模式.
func (kf *LkkFile) GetFileMode(fpath string) (os.FileMode, error) {
	finfo, err := os.Lstat(fpath)
	if err != nil {
		return 0, err
	}
	return finfo.Mode(), nil
}

// AppendFile 插入文件内容.
func (kf *LkkFile) AppendFile(fpath string, data []byte) error {
	if fpath == "" {
		return errors.New("[AppendFile] no path provided")
	}

	var file *os.File
	filePerm, err := kf.GetFileMode(fpath)
	if err != nil {
		// create the file
		file, err = os.Create(fpath)
	} else {
		// open for append
		file, err = os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
	}
	if err != nil {
		// failed to create or open-for-append the file
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.Write(data)

	return err
}

// GetMime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取.
func (kf *LkkFile) GetMime(fpath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(fpath)
		//当unix系统中没有相关的mime.types文件时,将返回空
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(fpath)
		if err != nil {
			return res
		}

		buffer := make([]byte, 512)
		_, err = srcFile.Read(buffer)
		if err != nil {
			return res
		}

		res = http.DetectContentType(buffer)
	}

	return res
}

// FileSize 获取文件大小(bytes字节);注意:文件不存在或无法访问时返回-1 .
func (kf *LkkFile) FileSize(fpath string) int64 {
	f, err := os.Stat(fpath)
	if nil != err {
		return -1
	}
	return f.Size()
}

// DirSize 获取目录大小(bytes字节).
func (kf *LkkFile) DirSize(fpath string) int64 {
	var size int64
	//filepath.Walk压测很慢
	_ = filepath.Walk(fpath, func(_ string, info os.FileInfo, err error) error {
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

// IsExist 路径(文件/目录)是否存在.
func (kf *LkkFile) IsExist(fpath string) bool {
	_, err := os.Stat(fpath)
	return err == nil || os.IsExist(err)
}

// IsLink 是否链接文件(软链接,且存在).
func (kf *LkkFile) IsLink(fpath string) bool {
	f, err := os.Lstat(fpath)
	if err != nil {
		return false
	}

	return f.Mode()&os.ModeSymlink == os.ModeSymlink
}

// IsFile 是否(某类型)文件,且存在.
// ftype为枚举(FILE_TYPE_ANY、FILE_TYPE_LINK、FILE_TYPE_REGULAR、FILE_TYPE_COMMON),默认FILE_TYPE_ANY;
func (kf *LkkFile) IsFile(fpath string, ftype ...LkkFileType) (res bool) {
	var t LkkFileType = FILE_TYPE_ANY
	if len(ftype) > 0 {
		t = ftype[0]
	}

	var f os.FileInfo
	var e error
	var musLink, musRegular bool

	if t == FILE_TYPE_LINK {
		musLink = true
	} else if t == FILE_TYPE_REGULAR {
		musRegular = true
	} else if t == FILE_TYPE_COMMON {
		musLink = true
		musRegular = true
	}

	if (!musLink && !musRegular) || musRegular {
		f, e := os.Stat(fpath)
		if musRegular {
			res = (e == nil) && f.Mode().IsRegular()
		} else {
			res = (e == nil) && !f.IsDir()
		}
	}

	if !res && musLink {
		f, e = os.Lstat(fpath)
		res = (e == nil) && (f.Mode()&os.ModeSymlink == os.ModeSymlink)
	}

	return
}

// IsDir 是否目录(且存在).
func (kf *LkkFile) IsDir(fpath string) bool {
	f, err := os.Lstat(fpath)
	if os.IsNotExist(err) || nil != err {
		return false
	}
	return f.IsDir()
}

// IsBinary 是否二进制文件(且存在).
func (kf *LkkFile) IsBinary(fpath string) bool {
	cont, err := kf.ReadFile(fpath)
	if err != nil {
		return false
	}

	return isBinary(string(cont))
}

// IsImg 是否图片文件(仅检查后缀).
func (kf *LkkFile) IsImg(fpath string) bool {
	ext := kf.GetExt(fpath)
	switch ext {
	case "jpg", "jpeg", "bmp", "gif", "png", "svg", "ico", "webp":
		return true
	default:
		return false
	}
}

// Mkdir 新建目录,允许多级目录.
func (kf *LkkFile) Mkdir(fpath string, mode os.FileMode) error {
	return os.MkdirAll(fpath, mode)
}

// AbsPath 获取绝对路径,path可允许不存在.
func (kf *LkkFile) AbsPath(fpath string) string {
	fullPath := ""
	res, err := filepath.Abs(fpath) // filepath.Abs最终使用到os.Getwd()检查
	if err != nil {
		fullPath = filepath.Clean(filepath.Join(`/`, fpath))
	} else {
		fullPath = res
	}

	return fullPath
}

// RealPath 返回规范化的真实绝对路径名.path必须存在,若路径不存在则返回空字符串.
func (kf *LkkFile) RealPath(fpath string) string {
	fullPath := fpath
	if !filepath.IsAbs(fpath) {
		wd, err := os.Getwd()
		if err != nil {
			return ""
		}
		fullPath = filepath.Clean(wd + `/` + fpath)
	}

	_, err := os.Stat(fullPath)
	if err != nil {
		return ""
	}

	return fullPath
}

// Touch 快速创建指定大小的文件,size为字节.
func (kf *LkkFile) Touch(fpath string, size int64) bool {
	//创建目录
	destDir := filepath.Dir(fpath)
	if destDir != "" && !kf.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0777); err != nil {
			return false
		}
	}

	fd, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return false
	}
	defer func() {
		_ = fd.Close()
	}()

	if size > 1 {
		_, _ = fd.Seek(size-1, 0)
		_, _ = fd.Write([]byte{0})
	}

	return true
}

// Rename 重命名(或移动)文件/目录.
func (kf *LkkFile) Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Unlink 删除文件.
func (kf *LkkFile) Unlink(fpath string) error {
	return os.Remove(fpath)
}

// CopyFile 拷贝源文件到目标文件,cover为枚举(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (kf *LkkFile) CopyFile(source string, dest string, cover LkkFileCover) (int64, error) {
	if source == dest {
		return 0, nil
	}

	sourceStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	} else if !sourceStat.Mode().IsRegular() {
		return 0, fmt.Errorf("[CopyFile]`source %s is not a regular file", source)
	}

	if cover != FILE_COVER_ALLOW {
		if _, err := os.Stat(dest); err == nil {
			if cover == FILE_COVER_IGNORE {
				return 0, nil
			} else if cover == FILE_COVER_DENY {
				return 0, fmt.Errorf("[CopyFile]`dest File %s already exists", dest)
			}
		}
	}

	sourceFile, _ := os.Open(source)
	defer func() {
		_ = sourceFile.Close()
	}()

	//创建目录
	destDir := filepath.Dir(dest)
	if destDir != "" && !kf.IsDir(destDir) {
		if err = os.MkdirAll(destDir, 0766); err != nil {
			return 0, err
		}
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = destFile.Close()
	}()

	var nBytes int64
	sourceSize := sourceStat.Size()
	if sourceSize <= 1048576 { //1M以内小文件使用buffer拷贝
		var total int
		var bufferSize int = 102400
		if sourceSize < 524288 {
			bufferSize = 51200
		}

		buf := make([]byte, bufferSize)
		for {
			n, err := sourceFile.Read(buf)
			if err != nil && err != io.EOF {
				return int64(total), err
			} else if n == 0 {
				break
			}

			if _, err := destFile.Write(buf[:n]); err != nil || !kf.IsExist(dest) {
				return int64(total), err
			}

			total += n
		}
		nBytes = int64(total)
	} else {
		nBytes, err = io.Copy(destFile, sourceFile)
		if err == nil {
			err = os.Chmod(dest, sourceStat.Mode())
		}
	}

	return nBytes, err
}

// FastCopy 快速拷贝源文件到目标文件,不做安全检查.
func (kf *LkkFile) FastCopy(source string, dest string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = sourceFile.Close()
	}()

	//创建目录
	destDir := filepath.Dir(dest)
	if !kf.IsDir(destDir) {
		if err = os.MkdirAll(destDir, 0766); err != nil {
			return 0, err
		}
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = destFile.Close()
	}()

	var bufferSize int = 32768
	var nBytes int
	buf := make([]byte, bufferSize)
	for {
		n, err := sourceFile.Read(buf)
		if err != nil && err != io.EOF {
			return int64(nBytes), err
		} else if n == 0 {
			break
		}

		if _, err := destFile.Write(buf[:n]); err != nil || !kf.IsExist(dest) {
			return int64(nBytes), err
		}

		nBytes += n
	}

	return int64(nBytes), err
}
