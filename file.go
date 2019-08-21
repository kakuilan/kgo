package kgo

import (
	"archive/tar"
	"compress/gzip"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
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

// GetExt 获取文件扩展名,不包括点"."
func (kf *LkkFile) GetExt(path string) string {
	suffix := filepath.Ext(path)
	if suffix != "" {
		return strings.ToLower(suffix[1:])
	}
	return suffix
}

// GetContents 获取文件内容作为字符串
func (kf *LkkFile) GetContents(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}

// PutContents 将一个字符串写入文件
func (kf *LkkFile) PutContents(fpath string, data []byte) error {
	if dir := path.Dir(fpath); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(fpath, data, 0644)
}

// GetMime 获取文件mime类型;fast为true时根据后缀快速获取;为false时读取文件头获取
func (kf *LkkFile) GetMime(path string, fast bool) string {
	var res string
	if fast {
		suffix := filepath.Ext(path)
		res = mime.TypeByExtension(suffix)
	} else {
		srcFile, err := os.Open(path)
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

// FileSize 获取文件大小(bytes字节),注意:文件不存在或无法访问返回-1
func (kf *LkkFile) FileSize(path string) int64 {
	f, err := os.Stat(path)
	if nil != err {
		return -1
	}
	return f.Size()
}

// DirSize 获取目录大小(bytes字节)
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

// IsExist 路径(文件/目录)是否存在
func (kf *LkkFile) IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// IsWritable 路径是否可写
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

// IsFile 是否常规文件(且存在)
func (kf *LkkFile) IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	//常规文件,不包括链接
	return stat.Mode().IsRegular()
}

// IsLink 是否链接文件(且存在)
func (kf *LkkFile) IsLink(path string) bool {
	f, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return f.Mode()&os.ModeSymlink == os.ModeSymlink
}

// IsDir 是否目录(且存在)
func (kf *LkkFile) IsDir(path string) bool {
	f, err := os.Lstat(path)
	if os.IsNotExist(err) || nil != err {
		return false
	}
	return f.IsDir()
}

// IsBinary 是否二进制文件(且存在)
func (kf *LkkFile) IsBinary(path string) bool {
	cont, err := kf.GetContents(path)
	if err != nil {
		return false
	}

	return KStr.IsBinary(string(cont))
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

// Mkdir 新建目录
func (kf *LkkFile) Mkdir(filename string, mode os.FileMode) error {
	return os.MkdirAll(filename, mode)
}

// AbsPath 获取绝对路径,path可允许不存在
func (kf *LkkFile) AbsPath(path string) string {
	fullPath := ""
	res, err := filepath.Abs(path)
	if err != nil {
		fullPath = filepath.Clean(filepath.Join(`/`, path))
	} else {
		fullPath = res
	}

	return fullPath
}

// Realpath 返回规范化的真实绝对路径名,path必须存在
func (kf *LkkFile) Realpath(path string) string {
	_, err := os.Stat(path)
	if err != nil {
		return ""
	}

	if filepath.IsAbs(path) {
		return path
	}

	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	return filepath.Clean(wd + `/` + path)
}

// Touch 快速创建指定大小的文件,size为字节
func (kf *LkkFile) Touch(path string, size int64) bool {
	//创建目录
	destDir := filepath.Dir(path)
	if destDir != "" && !kf.IsDir(destDir) {
		if err := os.MkdirAll(destDir, 0766); err != nil {
			return false
		}
	}

	fd, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return false
	}
	defer fd.Close()

	if size > 1 {
		_, _ = fd.Seek(size-1, 0)
		_, _ = fd.Write([]byte{0})
	}

	return true
}

// Rename 重命名一个文件或目录
func (kf *LkkFile) Rename(oldname, newname string) error {
	return os.Rename(oldname, newname)
}

// Unlink 删除文件
func (kf *LkkFile) Unlink(filename string) error {
	return os.Remove(filename)
}

// CopyFile 拷贝源文件到目标文件,cover为枚举(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY)
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
	defer sourceFile.Close()

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
	defer destFile.Close()

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

			if _, err := destFile.Write(buf[:n]); err != nil {
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

// FastCopy 快速拷贝源文件到目标文件,不做安全检查
func (kf *LkkFile) FastCopy(source string, dest string) (int64, error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return 0, err
	}

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

		if _, err := destFile.Write(buf[:n]); err != nil {
			return int64(nBytes), err
		}

		nBytes += n
	}

	return int64(nBytes), err
}

// CopyLink 拷贝链接
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

// CopyDir 拷贝源目录到目标目录,cover为枚举(FILE_COVER_ALLOW、FILE_COVER_IGNORE、FILE_COVER_DENY)
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
	defer directory.Close()

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
				if cover != FILE_COVER_ALLOW {
					continue
				} else if os.SameFile(obj, destFileInfo) {
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

// Img2Base64 读取图片文件,并转换为base64字符串
func (kf *LkkFile) Img2Base64(path string) (string, error) {
	if !kf.IsImg(path) {
		return "", fmt.Errorf("%s is not a image", path)
	}

	imgBuffer, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	ext := kf.GetExt(path)
	return fmt.Sprintf("data:image/%s;base64,%s", ext, base64.StdEncoding.EncodeToString(imgBuffer)), nil
}

// DelDir 删除目录;delRoot为true时连该目录一起删除;为false时只清空该目录
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

// FileTree 获取目录的文件树列表;
// ftype为枚举(FILE_TREE_ALL、FILE_TREE_DIR、FILE_TREE_FILE);
// recursive为是否递归;
// filters为一个或多个文件过滤器函数,FileFilter类型
func (kf *LkkFile) FileTree(path string, ftype LkkFileTree, recursive bool, filters ...FileFilter) []string {
	var trees []string

	if kf.IsFile(path) || kf.IsLink(path) {
		if ftype != FILE_TREE_DIR {
			trees = append(trees, path)
		}
		return trees
	}

	path = strings.TrimRight(path, "/")
	files, err := filepath.Glob(filepath.Join(path, "*"))
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

// FormatDir 格式化路径,将"\","//"替换为"/",且以"/"结尾
func (kf *LkkFile) FormatDir(path string) string {
	if path == "" {
		return ""
	}

	re := regexp.MustCompile(`(/){2,}|(\\){1,}`)
	str := re.ReplaceAllString(path, "/")
	return strings.TrimRight(str, "/") + "/"
}

// Md5 获取文件md5值,length指定结果长度32/16
func (kf *LkkFile) Md5(path string, length uint8) (string, error) {
	var res string
	f, err := os.Open(path)
	if err != nil {
		return res, err
	}
	defer f.Close()

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

// ShaX 计算文件的 shaX 散列值,x为1/256/512
func (kf *LkkFile) ShaX(path string, x uint16) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(shaXStr(data, x)), nil
}

// Pathinfo 获取文件路径的信息,options的值为-1: all; 1: dirname; 2: basename; 4: extension; 8: filename
func (kf *LkkFile) Pathinfo(path string, options int) map[string]string {
	if options == -1 {
		options = 1 | 2 | 4 | 8
	}
	info := make(map[string]string)
	if (options & 1) == 1 {
		info["dirname"] = filepath.Dir(path)
	}
	if (options & 2) == 2 {
		info["basename"] = filepath.Base(path)
	}
	if ((options & 4) == 4) || ((options & 8) == 8) {
		basename := ""
		if (options & 2) == 2 {
			basename, _ = info["basename"]
		} else {
			basename = filepath.Base(path)
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

// Basename 返回路径中的文件名部分
func (kf *LkkFile) Basename(path string) string {
	return filepath.Base(path)
}

// Dirname 返回路径中的目录部分,注意空路径或无目录的返回"."
func (kf *LkkFile) Dirname(path string) string {
	return filepath.Dir(path)
}

// Filemtime 取得文件修改时间
func (kf *LkkFile) Filemtime(filename string) (int64, error) {
	fileinfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileinfo.ModTime().Unix(), nil
}

// Glob 寻找与模式匹配的文件路径
func (kf *LkkFile) Glob(pattern string) ([]string, error) {
	return filepath.Glob(pattern)
}

// TarGz 打包压缩tar.gz;src为源文件或目录,dstTar为打包的路径名,ignorePatterns为要忽略的文件正则
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

	src = kf.Realpath(src)
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
	defer fw.Close()

	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

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
			defer fr.Close()

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

// UnTarGz 将tar.gz文件解压缩;srcTar为压缩包,dstDir为解压目录
func (kf *LkkFile) UnTarGz(srcTar, dstDir string) (bool, error) {
	fr, err := os.Open(srcTar)
	if err != nil {
		return false, err
	}
	defer fr.Close()

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
