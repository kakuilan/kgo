package kgo

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
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
	var err error
	dir := path.Dir(fpath)
	if err = os.MkdirAll(dir, os.ModePerm); err == nil {
		var p os.FileMode = 0655
		if len(perm) > 0 {
			p = perm[0]
		}
		err = os.WriteFile(fpath, data, p)
	}

	return err
}

// GetFileMode 获取路径的权限模式.
func (kf *LkkFile) GetFileMode(fpath string) (os.FileMode, error) {
	finfo, err := os.Lstat(fpath)
	if err != nil {
		return 0, err
	}
	return finfo.Mode(), nil
}

// AppendFile 插入文件内容.若文件不存在,则自动创建.
func (kf *LkkFile) AppendFile(fpath string, data []byte) error {
	var err error
	var file *os.File

	dir := path.Dir(fpath)
	if err = os.MkdirAll(dir, os.ModePerm); err == nil {
		var filePerm os.FileMode
		filePerm, err = kf.GetFileMode(fpath)
		if err != nil {
			file, err = os.Create(fpath)
		} else {
			file, err = os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, filePerm)
		}
		defer func() {
			_ = file.Close()
		}()

		if err == nil {
			_, err = file.Write(data)
		}
	}

	return err
}

// GetMime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取.
func (kf *LkkFile) GetMime(fpath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(fpath)
		//若unix系统中没有相关的mime.types文件时,将返回空
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(fpath)
		if err == nil {
			buffer := make([]byte, 512)
			_, err = srcFile.Read(buffer)
			if err == nil {
				res = http.DetectContentType(buffer)
			}
		}
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
	var res bool
	f, err := os.Lstat(fpath)
	if err == nil {
		res = f.Mode()&os.ModeSymlink == os.ModeSymlink
	}

	return res
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

	return err == nil && isBinary(string(cont))
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

// Mkdir 创建目录,允许多级.
func (kf *LkkFile) Mkdir(fpath string, mode os.FileMode) error {
	return os.MkdirAll(fpath, mode)
}

// AbsPath 获取绝对路径,path可允许不存在.
func (kf *LkkFile) AbsPath(fpath string) string {
	res, err := filepath.Abs(fpath) // filepath.Abs最终使用到os.Getwd()检查
	if err != nil {
		res = filepath.Clean(filepath.Join(string(os.PathSeparator), fpath))
	}

	return res
}

// RealPath 返回规范化的真实绝对路径名.path必须存在,若路径不存在则返回空字符串.
func (kf *LkkFile) RealPath(fpath string) string {
	res := fpath
	if !filepath.IsAbs(fpath) {
		wd, _ := os.Getwd()
		res = filepath.Clean(wd + `/` + fpath)
	}

	_, err := os.Stat(res)
	if err != nil {
		return ""
	}

	return res
}

// Touch 快速创建指定大小的文件,fpath为文件路径,size为字节.
func (kf *LkkFile) Touch(fpath string, size int64) (res bool) {
	//检查文件是否存在
	if !kf.IsExist(fpath) {
		//创建目录
		destDir := filepath.Dir(fpath)
		err := os.MkdirAll(destDir, 0766)
		if err == nil {
			fd, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
			defer func() {
				_ = fd.Close()
			}()
			if err == nil {
				res = true
				if size > 1 {
					_, _ = fd.Seek(size-1, 0)
					_, _ = fd.Write([]byte{0})
				}
			}
		}
	}

	return
}

// Rename 重命名(或移动)文件/目录.
func (kf *LkkFile) Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Unlink 删除文件.
func (kf *LkkFile) Unlink(fpath string) error {
	return os.Remove(fpath)
}

// CopyFile 拷贝source源文件到dest目标文件,cover为是否覆盖,枚举值(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (kf *LkkFile) CopyFile(source string, dest string, cover LkkFileCover) (int64, error) {
	if source == dest {
		return 0, nil
	}

	// os.Stat,获取文件信息对象,符号链接将跳转
	sourceStat, err := os.Stat(source)
	if err != nil {
		return 0, err
	} else if !sourceStat.Mode().IsRegular() { //是否普通文件
		return 0, fmt.Errorf("[CopyFile]`source %s is not a regular file", source)
	}

	//非覆盖模式
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

	//源目录
	srcDirStat, _ := os.Stat(filepath.Dir(source))

	//创建目录
	destDir := filepath.Dir(dest)
	if err = os.MkdirAll(destDir, srcDirStat.Mode()); err != nil {
		return 0, err
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
		var total, num int
		var bufferSize int = 102400
		if sourceSize < 524288 {
			bufferSize = 51200
		}

		buf := make([]byte, bufferSize)
		for {
			num, err = sourceFile.Read(buf)
			if err == nil && num > 0 {
				total += num
				_, err = destFile.Write(buf[:num])
			}

			if num == 0 || err != nil {
				break
			}
		}
		nBytes = int64(total)
	} else {
		nBytes, err = io.Copy(destFile, sourceFile)
	}

	if err == nil || err == io.EOF {
		err = os.Chmod(dest, sourceStat.Mode())
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
	if err = os.MkdirAll(destDir, 0766); err != nil {
		return 0, err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = destFile.Close()
	}()

	var bufferSize int = 32768
	var nBytes, num int
	buf := make([]byte, bufferSize)
	for {
		num, err = sourceFile.Read(buf)
		if err == nil && num > 0 {
			nBytes += num
			_, err = destFile.Write(buf[:num])
		}
		if num == 0 || err != nil {
			break
		}
	}
	if err == io.EOF {
		err = nil
	}

	return int64(nBytes), err
}

// CopyLink 拷贝链接.source为源链接,dest为目标链接,cover为是否覆盖,枚举值(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (kf *LkkFile) CopyLink(source string, dest string, cover LkkFileCover) error {
	if source == dest {
		return nil
	}

	//获取原文件地址
	srcFile, err := os.Readlink(source)
	if err != nil {
		return err
	}

	// os.Lstat,获取文件信息对象,符号链接不跳转
	_, err = os.Lstat(dest)
	if err == nil {
		if cover == FILE_COVER_IGNORE { //忽略,不覆盖
			return nil
		} else if cover == FILE_COVER_DENY { //禁止覆盖
			return fmt.Errorf("[CopyLink]`dest File %s already exists", dest)
		} else { //移除已存在的目标文件
			_ = os.Remove(dest)
		}
	}

	//源目录
	srcDirStat, _ := os.Stat(filepath.Dir(source))

	//创建目录
	destDir := filepath.Dir(dest)
	if err = os.MkdirAll(destDir, srcDirStat.Mode()); err != nil {
		return err
	}

	return os.Symlink(srcFile, dest)
}

// CopyDir 拷贝目录.source为源目录,dest为目标目录,cover为是否覆盖,枚举值(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (kf *LkkFile) CopyDir(source string, dest string, cover LkkFileCover) (int64, error) {
	var total, nBytes int64

	if source == "" || source == dest {
		return 0, nil
	}

	sourceInfo, err := os.Stat(source)
	if err != nil {
		return 0, err
	} else if !sourceInfo.IsDir() {
		return 0, fmt.Errorf("[CopyDir]`source %s is not a directory", source)
	}

	// 创建目录
	if err = os.MkdirAll(dest, sourceInfo.Mode()); err != nil {
		return 0, err
	}

	var directory *os.File
	directory, err = os.Open(source)
	defer func() {
		_ = directory.Close()
	}()
	if err != nil {
		return 0, err
	}

	var objects []os.FileInfo
	objects, _ = directory.Readdir(-1)

	var destFileInfo os.FileInfo
	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// 递归创建子目录
			nBytes, err = kf.CopyDir(srcFilePath, destFilePath, cover)
		} else {
			destFileInfo, err = os.Stat(destFilePath)
			if err == nil {
				if cover != FILE_COVER_ALLOW || os.SameFile(obj, destFileInfo) {
					continue
				}
			}

			if obj.Mode()&os.ModeSymlink != 0 {
				// 链接文件
				err = kf.CopyLink(srcFilePath, destFilePath, cover)
			} else {
				nBytes, err = kf.CopyFile(srcFilePath, destFilePath, cover)
			}
		}

		if err == nil {
			total += nBytes
		} else {
			break
		}
	}

	return total, err
}

// DelDir 删除目录.delete为true时连该目录一起删除;为false时只清空该目录.
func (kf *LkkFile) DelDir(dir string, delete bool) error {
	realPath := kf.AbsPath(dir)
	if !kf.IsDir(realPath) {
		return fmt.Errorf("[DelDir]`dir %s not exists", dir)
	}

	files, err := os.ReadDir(realPath)
	if err != nil {
		return err
	}

	for _, item := range files {
		file := path.Join(realPath, item.Name())
		err = os.RemoveAll(file)
	}

	//删除目录
	if delete {
		err = os.RemoveAll(realPath)
	}

	return err
}

// Img2Base64 读取图片文件,并转换为base64字符串.
func (kf *LkkFile) Img2Base64(fpath string) (string, error) {
	if !kf.IsImg(fpath) {
		return "", fmt.Errorf("[Img2Base64]`fpath %s is not a image", fpath)
	}

	imgBuffer, err := os.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	ext := kf.GetExt(fpath)
	return img2Base64(imgBuffer, ext), nil
}

// FileTree 获取目录的文件树列表.
// ftype为枚举(FILE_TREE_ALL、FILE_TREE_DIR、FILE_TREE_FILE);
// recursive为是否递归;
// filters为一个或多个文件过滤器函数,FileFilter类型.
func (kf *LkkFile) FileTree(fpath string, ftype LkkFileTree, recursive bool, filters ...FileFilter) []string {
	var trees []string

	fpath = strings.TrimRight(fpath, "/\\")
	files, err := filepath.Glob(filepath.Join(fpath, "*"))
	if err != nil || len(files) == 0 {
		return trees
	}

	for _, file := range files {
		file = strings.ReplaceAll(file, "\\", "/")
		if kf.IsDir(file) {
			if ftype != FILE_TREE_FILE {
				trees = append(trees, file)
			}

			if recursive {
				subs := kf.FileTree(file, ftype, recursive, filters...)
				trees = append(trees, subs...)
			}
		} else if ftype != FILE_TREE_DIR {
			//文件过滤
			chk := true
			if len(filters) > 0 {
				for _, filter := range filters {
					chk = filter(file)
					if !chk {
						break
					}
				}
			}
			if !chk {
				continue
			}

			trees = append(trees, file)
		}
	}

	return trees
}

// FormatDir 格式化目录,将"\","//"替换为"/",且以"/"结尾.
func (kf *LkkFile) FormatDir(fpath string) string {
	return formatDir(fpath)
}

// Md5File 获取文件md5值,fpath为文件路径,length指定结果长度32/16.
func (kf *LkkFile) Md5File(fpath string, length uint8) (string, error) {
	var res []byte
	f, err := os.Open(fpath)
	defer func() {
		_ = f.Close()
	}()

	if err == nil {
		res, err = md5Reader(f, length)
	}

	return string(res), err
}

// Md5Reader 计算Reader的 MD5 散列值.
func (kf *LkkFile) Md5Reader(reader io.Reader, length uint8) (string, error) {
	res, err := md5Reader(reader, length)
	return string(res), err
}

// ShaXFile 计算文件的 shaX 散列值,fpath为文件路径,x为1/256/512.
func (kf *LkkFile) ShaXFile(fpath string, x uint16) (string, error) {
	var res []byte
	f, err := os.Open(fpath)
	defer func() {
		_ = f.Close()
	}()

	if err == nil {
		res, err = shaXReader(f, x)
	}

	return string(res), err
}

// ShaXReader 计算Reader的 shaX 散列值,x为1/256/512.
func (kf *LkkFile) ShaXReader(reader io.Reader, x uint16) (string, error) {
	res, err := shaXReader(reader, x)
	return string(res), err
}

// Pathinfo 获取文件路径的信息.
// option为要返回的信息,枚举值如下:
// -1: all; 1: dirname; 2: basename; 4: extension; 8: filename;
// 若要查看某几项,则为它们之间的和.
func (kf *LkkFile) Pathinfo(fpath string, option int) map[string]string {
	if option == -1 {
		option = 1 | 2 | 4 | 8
	}
	res := make(map[string]string)
	if (option & 1) == 1 {
		res["dirname"] = filepath.Dir(fpath)
	}
	if (option & 2) == 2 {
		res["basename"] = filepath.Base(fpath)
	}
	if ((option & 4) == 4) || ((option & 8) == 8) {
		basename := ""
		if (option & 2) == 2 {
			basename = res["basename"]
		} else {
			basename = filepath.Base(fpath)
		}
		p := strings.LastIndex(basename, ".")
		filename, extension := "", ""
		if p > 0 {
			filename, extension = basename[:p], basename[p+1:]
		} else if p == -1 {
			filename = basename
		} else if p == 0 {
			extension = basename[p+1:]
		}
		if (option & 4) == 4 {
			res["extension"] = extension
		}
		if (option & 8) == 8 {
			res["filename"] = filename
		}
	}
	return res
}

// Basename 返回路径中的文件名部分(包括后缀),空路径时返回".".
func (kf *LkkFile) Basename(fpath string) string {
	return filepath.Base(fpath)
}

// Dirname 返回路径中的目录部分,注意空路径或无目录的返回".".
func (kf *LkkFile) Dirname(fpath string) string {
	return filepath.Dir(fpath)
}

// GetModTime 获取文件的修改时间戳,秒.
func (kf *LkkFile) GetModTime(fpath string) (res int64) {
	fileinfo, err := os.Stat(fpath)
	if err == nil {
		res = fileinfo.ModTime().Unix()
	}
	return
}

// Glob 寻找与模式匹配的文件路径.
func (kf *LkkFile) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// SafeFileName 将文件名转换为安全可用的字符串.
func (kf *LkkFile) SafeFileName(str string) string {
	name := path.Clean(path.Base(str))
	name = strings.Trim(name, " ")
	separators, err := regexp.Compile(`[ &_=+:]`)
	if err == nil {
		name = separators.ReplaceAllString(name, "-")
	}
	legal, err := regexp.Compile(`[^[:alnum:]-.]`)
	if err == nil {
		name = legal.ReplaceAllString(name, "")
	}
	for strings.Contains(name, "--") {
		name = strings.Replace(name, "--", "-", -1)
	}
	return name
}

// TarGz 打包压缩tar.gz.
// src为源文件或目录,dstTar为打包的路径名,ignorePatterns为要忽略的文件正则.
func (kf *LkkFile) TarGz(src string, dstTar string, ignorePatterns ...string) (bool, error) {
	//过滤器,检查要忽略的文件
	var filter = func(file string) bool {
		res := true
		for _, pattern := range ignorePatterns {
			re, err := regexp.Compile(pattern)
			if err == nil {
				chk := re.MatchString(file)
				if chk {
					res = false
					break
				}
			}
		}
		return res
	}

	src = kf.AbsPath(src)
	dstTar = kf.AbsPath(dstTar)

	//创建目录
	dstDir := filepath.Dir(dstTar)
	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return false, err
	}

	files := kf.FileTree(src, FILE_TREE_ALL, true, filter)
	if len(files) == 0 {
		return false, fmt.Errorf("[TarGz]`src no files to tar.gz")
	}

	// dest file write
	fw, err := os.Create(dstTar)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fw.Close()
	}()
	// gzip write
	gw := gzip.NewWriter(fw)
	defer func() {
		_ = gw.Close()
	}()

	// tar write
	tw := tar.NewWriter(gw)
	defer func() {
		_ = tw.Close()
	}()

	parentDir := filepath.Dir(src)
	for _, file := range files {
		if file == dstTar {
			continue
		}
		fi, err := os.Stat(file)
		if err != nil {
			continue
		}
		newName := strings.Replace(file, parentDir, "", -1)
		newName = strings.ReplaceAll(newName, ":", "") //防止wins下 tmp/D: 创建失败

		// Create tar header
		hdr := new(tar.Header)
		hdr.Format = tar.FormatGNU

		if fi.IsDir() {
			// if last character of header name is '/' it also can be directory
			// but if you don't set Typeflag, error will occur when you untargz
			hdr.Name = newName + "/"
			hdr.Typeflag = tar.TypeDir
			hdr.Size = 0
			//hdr.Mode = 0755 | c_ISDIR
			hdr.Mode = int64(fi.Mode())
			hdr.ModTime = fi.ModTime()

			// Write hander
			err := tw.WriteHeader(hdr)
			if err != nil {
				return false, fmt.Errorf("[TarGz] DirErr: %s file:%s\n", err.Error(), file)
			}
		} else {
			// File reader
			fr, err := os.Open(file)
			if err != nil {
				return false, fmt.Errorf("[TarGz] OpenErr: %s file:%s\n", err.Error(), file)
			}
			defer func() {
				_ = fr.Close()
			}()

			hdr.Name = newName
			hdr.Size = fi.Size()
			hdr.Mode = int64(fi.Mode())
			hdr.ModTime = fi.ModTime()

			// Write hander
			err = tw.WriteHeader(hdr)
			if err != nil {
				return false, fmt.Errorf("[TarGz] FileErr: %s file:%s\n", err.Error(), file)
			}

			// Write file data
			_, err = io.Copy(tw, fr)
			if err != nil {
				return false, fmt.Errorf("[TarGz] CopyErr: %s file:%s\n", err.Error(), file)
			}
			_ = fr.Close()
		}
	}

	return true, nil
}

// UnTarGz 将tar.gz文件解压缩.
// srcTar为压缩包,dstDir为解压目录.
func (kf *LkkFile) UnTarGz(srcTar, dstDir string) (bool, error) {
	fr, err := os.Open(srcTar)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fr.Close()
	}()

	//创建目录
	dstDir = strings.TrimRight(kf.AbsPath(dstDir), "/\\")
	if err = os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return false, err
	}

	// Gzip reader
	gr, err := gzip.NewReader(fr)
	if err != nil {
		return false, err
	}

	// Tar reader
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			// End of tar archive
			break
		} else if err != nil {
			return false, err
		}

		// Create diretory before create file
		newPath := dstDir + "/" + strings.TrimLeft(hdr.Name, "/\\")
		parentDir := path.Dir(newPath)
		if err = os.MkdirAll(parentDir, os.ModePerm); err != nil {
			return false, err
		}

		if hdr.Typeflag != tar.TypeDir {
			// Write data to file
			fw, err := os.Create(newPath)
			if err != nil {
				return false, fmt.Errorf("[UnTarGz] CreateErr: %s file:%s\n", err.Error(), newPath)
			}

			_, err = io.Copy(fw, tr)
			if err != nil {
				return false, fmt.Errorf("[UnTarGz] CopyErr: %s file:%s\n", err.Error(), newPath)
			}

			_ = fw.Close()
		}
	}

	return true, nil
}

// ChmodBatch 批量改变路径权限模式(包括子目录和所属文件).
// filemode为文件权限模式,dirmode为目录权限模式.
func (kf *LkkFile) ChmodBatch(fpath string, filemode, dirmode os.FileMode) (res bool) {
	err := filepath.Walk(fpath, func(fpath string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			err = os.Chmod(fpath, dirmode)
		} else {
			err = os.Chmod(fpath, filemode)
		}

		return err
	})

	if err == nil {
		res = true
	}

	return
}

// CountLines 统计文件行数.buffLength为缓冲长度,kb.
func (kf *LkkFile) CountLines(fpath string, buffLength int) (int, error) {
	fh, err := os.Open(fpath)
	if err != nil {
		return -1, err
	}
	defer func() {
		_ = fh.Close()
	}()

	count := 0
	lineSep := []byte{'\n'}

	if buffLength <= 0 {
		buffLength = 32
	}

	r := bufio.NewReader(fh)
	buf := make([]byte, buffLength*1024)
	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil
		case err != nil:
			return count, err
		}
	}
}

// Zip 将文件或目录进行zip打包.fpaths为源文件或目录的路径.
func (kf *LkkFile) Zip(dst string, fpaths ...string) (bool, error) {
	dst = kf.AbsPath(dst)
	dstDir := kf.Dirname(dst)
	if !kf.IsDir(dstDir) {
		err := os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	fzip, err := os.Create(dst)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fzip.Close()
	}()

	if len(fpaths) == 0 {
		return false, errors.New("[Zip] no input files")
	}

	var allfiles, files []string
	var fpath string
	for _, fpath = range fpaths {
		if kf.IsDir(fpath) {
			files = kf.FileTree(fpath, FILE_TREE_FILE, true)
			if len(files) != 0 {
				allfiles = append(allfiles, files...)
			}
		} else if fpath != "" {
			allfiles = append(allfiles, fpath)
		}
	}

	if len(allfiles) == 0 {
		return false, errors.New("[Zip] no exist files")
	}

	zipw := zip.NewWriter(fzip)
	defer func() {
		_ = zipw.Close()
	}()

	keys := make(map[string]bool)
	for _, fpath = range allfiles {
		if _, ok := keys[fpath]; ok || kf.AbsPath(fpath) == dst {
			continue
		}

		fileToZip, err := os.Open(fpath)
		if err != nil {
			return false, fmt.Errorf("[Zip] failed to open %s: %s", fpath, err)
		}
		defer func() {
			_ = fileToZip.Close()
		}()

		wr, _ := zipw.Create(fpath)
		keys[fpath] = true
		if _, err := io.Copy(wr, fileToZip); err != nil {
			return false, fmt.Errorf("[Zip] failed to write %s to zip: %s", fpath, err)
		}
	}

	return true, nil
}

// UnZip 解压zip文件.srcZip为zip文件路径,dstDir为解压目录.
func (kf *LkkFile) UnZip(srcZip, dstDir string) (bool, error) {
	reader, err := zip.OpenReader(srcZip)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = reader.Close()
	}()

	dstDir = strings.TrimRight(kf.AbsPath(dstDir), "/\\")
	if !kf.IsDir(dstDir) {
		err := os.MkdirAll(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	// 迭代压缩文件中的文件
	for _, f := range reader.File {
		// Create diretory before create file
		//newPath := dstDir + string(os.PathSeparator) + strings.TrimLeft(f.Name, string(os.PathSeparator))
		newPath := dstDir + "/" + strings.TrimLeft(f.Name, "/\\")
		parentDir := path.Dir(newPath)
		if !kf.IsDir(parentDir) {
			err := os.MkdirAll(parentDir, os.ModePerm)
			if err != nil {
				return false, err
			}
		}

		if !f.FileInfo().IsDir() {
			if fcreate, err := os.Create(newPath); err == nil {
				if rc, err := f.Open(); err == nil {
					_, _ = io.Copy(fcreate, rc)
					_ = rc.Close() //不要用defer来关闭，如果文件太多的话，会报too many open files 的错误
					_ = fcreate.Close()
				} else {
					_ = fcreate.Close()
					return false, err
				}
			} else {
				return false, err
			}
		}
	}

	return true, nil
}

// IsZip 是否zip文件.
func (kf *LkkFile) IsZip(fpath string) (bool, error) {
	var res bool
	var err error

	ext := kf.GetExt(fpath)
	if ext == "zip" {
		var f *os.File
		var n int
		f, err = os.Open(fpath)
		defer func() {
			_ = f.Close()
		}()
		if err == nil {
			buf := make([]byte, 4)
			n, err = f.Read(buf)
			res = err == nil && n == 4 && bytes.Equal(buf, []byte("PK\x03\x04"))
		}
	}

	return res, err
}
