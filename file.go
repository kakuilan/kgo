package kgo

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
)

// GetExt 获取文件的小写扩展名,不包括点"." .
func (kf *LkkFile) GetExt(fpath string) string {
	suffix := filepath.Ext(fpath)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// ReadInArray 把整个文件读入一个数组中,每行作为一个元素.
func (kf *LkkFile) ReadInArray(fpath string) ([]string, error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

// ReadFile 读取文件内容.
func (kf *LkkFile) ReadFile(fpath string) ([]byte, error) {
	data, err := ioutil.ReadFile(fpath)
	return data, err
}

// WriteFile 将内容写入文件.
// fpath为文件路径,data为内容,perm为权限.
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

	return ioutil.WriteFile(fpath, data, p)
}

// AppendFile 插入文件内容.
func (kf *LkkFile) AppendFile(fpath string, data []byte) error {
	if fpath == "" {
		return errors.New("No path provided")
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

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

// GetFileMode 获取路径的权限模式.
func (kf *LkkFile) GetFileMode(fpath string) (os.FileMode, error) {
	finfo, err := os.Lstat(fpath)
	if err != nil {
		return 0, err
	}
	return finfo.Mode(), nil
}

// GetMime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取.
func (kf *LkkFile) GetMime(fpath string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(fpath)
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

// FileSize 获取文件大小(bytes字节),注意:文件不存在或无法访问返回-1 .
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

// IsWritable 路径是否可写.
func (kf *LkkFile) IsWritable(fpath string) bool {
	err := syscall.Access(fpath, syscall.O_RDWR)
	if err != nil {
		return false
	}
	return true
}

// IsReadable 路径是否可读.
func (kf *LkkFile) IsReadable(fpath string) bool {
	err := syscall.Access(fpath, syscall.O_RDONLY)
	if err != nil {
		return false
	}
	return true
}

// IsExecutable 是否可执行文件.
func (kf *LkkFile) IsExecutable(fpath string) bool {
	info, err := os.Stat(fpath)
	return err == nil && info.Mode().IsRegular() && (info.Mode()&0111) != 0
}

// IsFile 是否常规文件(且存在).
func (kf *LkkFile) IsFile(fpath string) bool {
	stat, err := os.Stat(fpath)
	if err != nil {
		return false
	}
	//常规文件,不包括链接
	return stat.Mode().IsRegular()
}

// IsLink 是否链接文件(且存在).
func (kf *LkkFile) IsLink(fpath string) bool {
	f, err := os.Lstat(fpath)
	if err != nil {
		return false
	}

	return f.Mode()&os.ModeSymlink == os.ModeSymlink
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

	return KConv.IsBinary(string(cont))
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

// RealPath 返回规范化的真实绝对路径名,path必须存在.若路径不存在则返回空字符串.
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
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return false
		}
	}

	fd, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
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

// Rename 重命名文件或目录.
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
		return 0, fmt.Errorf("%s is not a regular file", source)
	}

	if cover != FILE_COVER_ALLOW {
		if _, err := os.Stat(dest); err == nil {
			if cover == FILE_COVER_IGNORE {
				return 0, nil
			} else if cover == FILE_COVER_DENY {
				return 0, fmt.Errorf("File %s already exists", dest)
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

// CopyLink 拷贝链接.
func (kf *LkkFile) CopyLink(source string, dest string) error {
	if source == dest {
		return nil
	}

	source, err := os.Readlink(source)
	if err != nil {
		return err
	}

	_, err = os.Lstat(dest)
	if err == nil {
		_ = os.Remove(dest)
	}

	//创建目录
	destDir := filepath.Dir(dest)
	if !kf.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return err
		}
	}

	return os.Symlink(source, dest)
}

// CopyDir 拷贝源目录到目标目录,cover为枚举(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY).
func (kf *LkkFile) CopyDir(source string, dest string, cover LkkFileCover) (int64, error) {
	var total, nBytes int64
	var err error
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return 0, err
	} else if !sourceInfo.IsDir() {
		return 0, fmt.Errorf("%s is not a directory", source)
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceInfo.Mode())
	if err != nil {
		return 0, err
	}

	directory, _ := os.Open(source)
	defer func() {
		_ = directory.Close()
	}()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return 0, err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			nBytes, err = kf.CopyDir(srcFilePath, destFilePath, cover)
		} else {
			destFileInfo, err := os.Stat(destFilePath)
			if err == nil {
				if cover != FILE_COVER_ALLOW || os.SameFile(obj, destFileInfo) {
					continue
				}
			}

			if obj.Mode()&os.ModeSymlink != 0 {
				// a link
				_ = kf.CopyLink(srcFilePath, destFilePath)
			} else {
				nBytes, err = kf.CopyFile(srcFilePath, destFilePath, cover)
			}
		}

		if err == nil {
			total += nBytes
		}
	}

	return total, err
}

// Img2Base64 读取图片文件,并转换为base64字符串.
func (kf *LkkFile) Img2Base64(fpath string) (string, error) {
	if !kf.IsImg(fpath) {
		return "", fmt.Errorf("%s is not a image", fpath)
	}

	imgBuffer, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	ext := kf.GetExt(fpath)
	return KStr.Img2Base64(imgBuffer, ext), nil
}

// DelDir 删除目录;delRoot为true时连该目录一起删除;为false时只清空该目录.
func (kf *LkkFile) DelDir(dir string, delRoot bool) error {
	realPath := kf.AbsPath(dir)
	if !kf.IsDir(realPath) {
		return fmt.Errorf("Dir %s not exists", dir)
	}

	names, err := ioutil.ReadDir(realPath)
	if err != nil {
		return err
	}

	for _, entery := range names {
		file := path.Join([]string{realPath, entery.Name()}...)
		err = os.RemoveAll(file)
	}

	//删除根节点(指定的目录)
	if delRoot {
		err = os.RemoveAll(realPath)
	}

	return err
}

// FileTree 获取目录的文件树列表.
// ftype为枚举(FILE_TREE_ALL、FILE_TREE_DIR、FILE_TREE_FILE);
// recursive为是否递归;
// filters为一个或多个文件过滤器函数,FileFilter类型.
func (kf *LkkFile) FileTree(fpath string, ftype LkkFileTree, recursive bool, filters ...FileFilter) []string {
	var trees []string

	if kf.IsFile(fpath) || kf.IsLink(fpath) {
		if ftype != FILE_TREE_DIR {
			trees = append(trees, fpath)
		}
		return trees
	}

	fpath = strings.TrimRight(fpath, "/")
	files, err := filepath.Glob(filepath.Join(fpath, "*"))
	if err != nil || len(files) == 0 {
		return trees
	}

	for _, file := range files {
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

		if kf.IsDir(file) {
			if ftype != FILE_TREE_FILE {
				trees = append(trees, file)
			}

			if recursive {
				subs := kf.FileTree(file, ftype, recursive)
				trees = append(trees, subs...)
			}
		} else if ftype != FILE_TREE_DIR {
			trees = append(trees, file)
		}
	}

	return trees
}

// FormatDir 格式化路径,将"\","//"替换为"/",且以"/"结尾.
func (kf *LkkFile) FormatDir(fpath string) string {
	if fpath == "" {
		return ""
	}

	// 将"\"替换为"/"
	fpath = strings.ReplaceAll(fpath, "\\", "/")

	str := RegFormatDir.ReplaceAllString(fpath, "/")
	return strings.TrimRight(str, "/") + "/"
}

// Md5 获取文件md5值,length指定结果长度32/16.
func (kf *LkkFile) Md5(fpath string, length uint8) (string, error) {
	var res string
	f, err := os.Open(fpath)
	if err != nil {
		return res, err
	}
	defer func() {
		_ = f.Close()
	}()

	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return res, err
	}

	hashInBytes := hash.Sum(nil)
	if length > 0 && length < 32 {
		dst := make([]byte, hex.EncodedLen(len(hashInBytes)))
		hex.Encode(dst, hashInBytes)
		res = string(dst[:length])
	} else {
		res = hex.EncodeToString(hashInBytes)
	}

	return res, nil
}

// ShaX 计算文件的 shaX 散列值,x为1/256/512.
func (kf *LkkFile) ShaX(fpath string, x uint16) (string, error) {
	data, err := ioutil.ReadFile(fpath)
	if err != nil {
		return "", err
	}

	return string(shaXStr(data, x)), nil
}

// Pathinfo 获取文件路径的信息,options的值为-1: all; 1: dirname; 2: basename; 4: extension; 8: filename.
func (kf *LkkFile) Pathinfo(fpath string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}
	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(fpath)
	}
	if (options & 2) == 2 {
		info["basename"] = filepath.Base(fpath)
	}
	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""
		if (options & 2) == 2 {
			basename, _ = info["basename"]
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
		if (options & 4) == 4 {
			info["extension"] = extension
		}
		if (options & 8) == 8 {
			info["filename"] = filename
		}
	}
	return info
}

// Basename 返回路径中的文件名部分.
func (kf *LkkFile) Basename(fpath string) string {
	return filepath.Base(fpath)
}

// Dirname 返回路径中的目录部分,注意空路径或无目录的返回"." .
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

// TarGz 打包压缩tar.gz;src为源文件或目录,dstTar为打包的路径名,ignorePatterns为要忽略的文件正则.
func (kf *LkkFile) TarGz(src string, dstTar string, ignorePatterns ...string) (bool, error) {
	//过滤器,检查要忽略的文件
	var filter = func(file string) bool {
		res := true
		for _, pattern := range ignorePatterns {
			re, err := regexp.Compile(pattern)
			if err != nil {
				continue
			}
			chk := re.MatchString(file)
			if chk {
				res = false
				break
			}
		}
		return res
	}

	src = kf.AbsPath(src)
	dstTar = kf.AbsPath(dstTar)

	dstDir := kf.Dirname(dstTar)
	if !kf.IsExist(dstDir) {
		_ = kf.Mkdir(dstDir, os.ModePerm)
	}

	files := kf.FileTree(src, FILE_TREE_ALL, true, filter)
	if len(files) == 0 {
		return false, fmt.Errorf("src no files to tar.gz")
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
				return false, fmt.Errorf("DirErr: %s file:%s\n", err.Error(), file)
			}
		} else {
			// File reader
			fr, err := os.Open(file)
			if err != nil {
				return false, fmt.Errorf("OpenErr: %s file:%s\n", err.Error(), file)
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
				return false, fmt.Errorf("FileErr: %s file:%s\n", err.Error(), file)
			}

			// Write file data
			_, err = io.Copy(tw, fr)
			if err != nil {
				return false, fmt.Errorf("CopyErr: %s file:%s\n", err.Error(), file)
			}
			_ = fr.Close()
		}
	}

	return true, nil
}

// UnTarGz 将tar.gz文件解压缩;srcTar为压缩包,dstDir为解压目录.
func (kf *LkkFile) UnTarGz(srcTar, dstDir string) (bool, error) {
	fr, err := os.Open(srcTar)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fr.Close()
	}()

	dstDir = strings.TrimRight(kf.AbsPath(dstDir), "/\\")
	if !kf.IsExist(dstDir) {
		err := kf.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	// Gzip reader
	gr, err := gzip.NewReader(fr)

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
		if !kf.IsExist(parentDir) {
			_ = os.MkdirAll(parentDir, os.ModePerm)
		}

		if hdr.Typeflag != tar.TypeDir {
			// Write data to file
			fw, err := os.Create(newPath)
			if err != nil {
				return false, fmt.Errorf("CreateErr: %s file:%s\n", err.Error(), newPath)
			}

			_, err = io.Copy(fw, tr)
			if err != nil {
				return false, fmt.Errorf("CopyErr: %s file:%s\n", err.Error(), newPath)
			}

			_ = fw.Close()
		}
	}

	return true, nil
}

// SafeFileName 将文件名转换为安全可用的字符串.
func (kf *LkkFile) SafeFileName(str string) string {
	name := strings.ToLower(str)
	name = path.Clean(path.Base(name))
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

// ChmodBatch 批量改变路径权限模式(包括子目录和所属文件).filemode为文件权限模式,dirmode为目录权限模式.
func (kf *LkkFile) ChmodBatch(fpath string, filemode, dirmode os.FileMode) (res bool) {
	var err error
	err = filepath.Walk(fpath, func(fpath string, f os.FileInfo, err error) error {
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
	file, err := os.Open(fpath)
	if err != nil {
		return -1, err
	}
	defer func() {
		_ = file.Close()
	}()

	count := 0
	lineSep := []byte{'\n'}

	if buffLength <= 0 {
		buffLength = 32
	}

	r := bufio.NewReader(file)
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

// Zip 将文件目录进行zip打包.fpaths为文件或目录的路径.
func (kf *LkkFile) Zip(dst string, fpaths ...string) (bool, error) {
	dst = kf.AbsPath(dst)
	dstDir := kf.Dirname(dst)
	if !kf.IsExist(dstDir) {
		_ = kf.Mkdir(dstDir, os.ModePerm)
	}

	fzip, err := os.Create(dst)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = fzip.Close()
	}()

	if len(fpaths) == 0 {
		return false, errors.New("No input files.")
	}

	var allfiles, files []string
	var fpath string
	for _, fpath = range fpaths {
		fpath = KStr.Trim(fpath)
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
		return false, errors.New("No exist files.")
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
			return false, fmt.Errorf("Failed to open %s: %s", fpath, err)
		}
		defer func() {
			_ = fileToZip.Close()
		}()

		wr, _ := zipw.Create(fpath)
		keys[fpath] = true
		if _, err := io.Copy(wr, fileToZip); err != nil {
			return false, fmt.Errorf("Failed to write %s to zip: %s", fpath, err)
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
	if !kf.IsExist(dstDir) {
		err := kf.Mkdir(dstDir, os.ModePerm)
		if err != nil {
			return false, err
		}
	}

	// 迭代压缩文件中的文件
	for _, f := range reader.File {
		// Create diretory before create file
		newPath := dstDir + "/" + strings.TrimLeft(f.Name, "/\\")
		parentDir := path.Dir(newPath)
		if !kf.IsExist(parentDir) {
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
func (kf *LkkFile) IsZip(fpath string) bool {
	ext := kf.GetExt(fpath)
	if ext != "zip" {
		return false
	}

	f, err := os.Open(fpath)
	if err != nil {
		return false
	}
	defer func() {
		_ = f.Close()
	}()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}
