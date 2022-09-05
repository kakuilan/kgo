package kgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestFile_GetExt(t *testing.T) {
	var ext string

	ext = KFile.GetExt(fileGo)
	assert.Equal(t, "go", ext)

	ext = KFile.GetExt(fileGitkee)
	assert.Equal(t, "gitkeep", ext)

	ext = KFile.GetExt(fileSongs)
	assert.Equal(t, "txt", ext)

	ext = KFile.GetExt(fileNone)
	assert.Empty(t, ext)
}

func BenchmarkFile_GetExt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetExt(fileMd)
	}
}

func TestFile_ReadFile(t *testing.T) {
	var bs []byte
	var err error

	bs, err = KFile.ReadFile(fileMd)
	assert.NotEmpty(t, bs)
	assert.Nil(t, err)

	//不存在的文件
	bs, err = KFile.ReadFile(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_ReadFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadFile(fileMd)
	}
}

func TestFile_ReadInArray(t *testing.T) {
	var sl []string
	var err error

	sl, err = KFile.ReadInArray(fileDante)
	assert.Equal(t, 19568, len(sl))

	//不存在的文件
	sl, err = KFile.ReadInArray(fileNone)
	assert.NotNil(t, err)
}

func BenchmarkFile_ReadInArray(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ReadInArray(fileMd)
	}
}

func TestFile_ReadFirstLine(t *testing.T) {
	var res []byte

	res = KFile.ReadFirstLine(fileDante)
	assert.NotEmpty(t, res)

	//不存在的文件
	res = KFile.ReadFirstLine(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_ReadFirstLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.ReadFirstLine(fileMd)
	}
}

func TestFile_ReadLastLine(t *testing.T) {
	var res []byte

	res = KFile.ReadLastLine(changLog)
	assert.NotEmpty(t, res)

	//不存在的文件
	res = KFile.ReadLastLine(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_ReadLastLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.ReadLastLine(fileMd)
	}
}

func TestFile_WriteFile(t *testing.T) {
	var err error

	err = KFile.WriteFile(putfile, bytsHello)
	assert.Nil(t, err)

	//设置权限
	err = KFile.WriteFile(putfile, bytsHello, 0777)
	assert.Nil(t, err)

	//无权限写
	err = KFile.WriteFile(rootFile1, bytsHello)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.NotNil(t, err)
	}

	//空路径
	err = KFile.WriteFile("", bytsHello, 0777)
	assert.NotNil(t, err)
}

func BenchmarkFile_WriteFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		filename := fmt.Sprintf("./testdata/file/putfile_%d", i)
		_ = KFile.WriteFile(filename, bytsHello)
	}
}

func TestFile_AppendFile(t *testing.T) {
	var err error

	//创建
	err = KFile.AppendFile(apndfile, bytsHello)
	assert.Nil(t, err)

	//追加
	err = KFile.AppendFile(apndfile, bytsHello)
	assert.Nil(t, err)

	//空路径
	err = KFile.AppendFile("", bytsHello)
	assert.NotNil(t, err)

	//权限不足
	err = KFile.AppendFile(rootFile1, bytsHello)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.NotNil(t, err)
	}
}

func BenchmarkFile_AppendFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KFile.AppendFile(apndfile, bytsHello)
	}
}

func TestFile_GetMime(t *testing.T) {
	var res string

	res = KFile.GetMime(imgPng, false)
	assert.NotEmpty(t, res)

	res = KFile.GetMime(fileDante, true)
	if KOS.IsWindows() {
		assert.NotEmpty(t, res)
	}

	res = KFile.GetMime(imgJpg, true)

	//空文件
	res = KFile.GetMime(gitkeep, false)
	assert.Empty(t, res)

	//不存在的文件
	res = KFile.GetMime(fileNone, true)
	assert.Empty(t, res)
}

func BenchmarkFile_GetMime_Fast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetMime(fileMd, true)
	}
}

func BenchmarkFile_GetMime_NoFast(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetMime(fileMd, false)
	}
}

func TestFile_FileSize(t *testing.T) {
	var res int64

	res = KFile.FileSize(changLog)
	assert.Greater(t, res, int64(0))

	//不存在的文件
	res = KFile.FileSize(fileNone)
	assert.Equal(t, int64(-1), res)
}

func BenchmarkFile_FileSize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FileSize(fileMd)
	}
}

func TestFile_DirSize(t *testing.T) {
	var res int64

	res = KFile.DirSize(dirCurr)
	assert.Greater(t, res, int64(0))

	//不存在的目录
	res = KFile.DirSize(fileNone)
	assert.Equal(t, int64(0), res)
}

func BenchmarkFile_DirSize(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.DirSize(dirTdat)
	}
}

func TestFile_IsExist(t *testing.T) {
	var res bool

	res = KFile.IsExist(changLog)
	assert.True(t, res)

	res = KFile.IsExist(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsExist(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsExist(fileMd)
	}
}

func TestFile_IsReadable(t *testing.T) {
	var res bool

	res = KFile.IsReadable(dirTdat)
	assert.True(t, res)

	//不存在的目录
	res = KFile.IsReadable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsReadable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsReadable(dirTdat)
	}
}

func TestFile_IsWritable(t *testing.T) {
	var res bool

	res = KFile.IsWritable(dirTdat)
	assert.True(t, res)

	//不存在的目录
	res = KFile.IsWritable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsWritable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsWritable(dirTdat)
	}
}

func TestFile_IsExecutable(t *testing.T) {
	res := KFile.IsExecutable(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsExecutable(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsExecutable(fileMd)
	}
}

func TestFile_IsLink(t *testing.T) {
	//创建链接文件
	if !KFile.IsExist(fileLink) {
		_ = os.Symlink(filePubPem, fileLink)
	}

	var res bool

	res = KFile.IsLink(fileLink)
	assert.True(t, res)

	//非链接
	res = KFile.IsLink(changLog)
	assert.False(t, res)

	//不存在
	res = KFile.IsLink(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsLink(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsLink(fileLink)
	}
}

func TestFile_IsFile(t *testing.T) {
	tests := []struct {
		f        string
		t        LkkFileType
		expected bool
	}{
		{"", FILE_TYPE_ANY, false},
		{fileNone, FILE_TYPE_ANY, false},
		{fileGo, FILE_TYPE_ANY, true},
		{fileMd, FILE_TYPE_LINK, false},
		{fileLink, FILE_TYPE_LINK, true},
		{fileLink, FILE_TYPE_REGULAR, false},
		{fileGitkee, FILE_TYPE_REGULAR, true},
		{fileLink, FILE_TYPE_COMMON, true},
		{imgJpg, FILE_TYPE_COMMON, true},
	}
	for _, test := range tests {
		actual := KFile.IsFile(test.f, test.t)
		assert.Equal(t, test.expected, actual)
	}
}

func BenchmarkFile_IsFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsFile(fileMd, FILE_TYPE_ANY)
	}
}

func TestFile_IsDir(t *testing.T) {
	var res bool

	res = KFile.IsDir(fileMd)
	assert.False(t, res)

	res = KFile.IsDir(fileNone)
	assert.False(t, res)

	res = KFile.IsDir(dirTdat)
	assert.True(t, res)
}

func BenchmarkFile_IsDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsDir(dirTdat)
	}
}

func TestFile_IsBinary(t *testing.T) {
	var res bool
	res = KFile.IsBinary(changLog)
	assert.False(t, res)

	res = KFile.IsBinary(imgPng)
	assert.True(t, res)

	//不存在
	res = KFile.IsBinary(fileNone)
	assert.False(t, res)
}

func BenchmarkFile_IsBinary(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsBinary(changLog)
	}
}

func TestFile_IsImg(t *testing.T) {
	var res bool

	res = KFile.IsImg(fileMd)
	assert.False(t, res)

	res = KFile.IsImg(imgSvg)
	assert.True(t, res)

	res = KFile.IsImg(imgPng)
	assert.True(t, res)
}

func BenchmarkFile_IsImg(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.IsImg(imgPng)
	}
}

func TestFile_Mkdir(t *testing.T) {
	err := KFile.Mkdir(dirNew, 0777)
	assert.Nil(t, err)
}

func BenchmarkFile_Mkdir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dname := fmt.Sprintf(dirNew+"/tmp_%d", i)
		_ = KFile.Mkdir(dname, 0777)
	}
}

func TestFile_AbsPath(t *testing.T) {
	var res string

	res = KFile.AbsPath(changLog)
	assert.NotEqual(t, '.', rune(res[0]))

	res = KFile.AbsPath(fileNone)
	assert.NotEmpty(t, res)

	//空路径
	res = KFile.AbsPath("")
	assert.NotEmpty(t, res)

	//windows下
	var bt = string(bytAstronomicalUnit)
	res = KFile.AbsPath(bt)
	assert.NotEmpty(t, res)
	if KOS.IsWindows() {
		chk := KStr.StartsWith(res, "\\", false)
		assert.True(t, chk)
	}
}

func BenchmarkFile_AbsPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.AbsPath(changLog)
	}
}

func TestFile_RealPath(t *testing.T) {
	var res string

	res = KFile.RealPath(fileMd)
	assert.NotEmpty(t, res)

	//不存在的路径
	res = KFile.RealPath(fileNone)
	assert.Empty(t, res)
}

func BenchmarkFile_RealPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.RealPath(fileMd)
	}
}

func TestFile_TouchRenameUnlink(t *testing.T) {
	var res bool
	var err error

	//创建不存在的文件
	res = KFile.Touch(touchfile, 2097152)
	assert.True(t, res)

	//创建已存在的文件
	res = KFile.Touch(fileGitkee, 256)
	assert.False(t, res)

	//重命名
	err = KFile.Rename(touchfile, renamefile)
	assert.Nil(t, err)

	//删除
	err = KFile.Unlink(renamefile)
	assert.Nil(t, err)
}

func BenchmarkFile_Touch(b *testing.B) {
	b.ResetTimer()
	var filename string
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf(dirTouch+"/zero_%d", i)
		KFile.Touch(filename, 0)
	}
}

func BenchmarkFile_Rename(b *testing.B) {
	b.ResetTimer()
	var f1, f2 string
	for i := 0; i < b.N; i++ {
		f1 = fmt.Sprintf(dirTouch+"/zero_%d", i)
		f2 = fmt.Sprintf(dirTouch+"/zero_re%d", i)
		_ = KFile.Rename(f1, f2)
	}
}

func BenchmarkFile_Unlink(b *testing.B) {
	b.ResetTimer()
	var filename string
	for i := 0; i < b.N; i++ {
		filename = fmt.Sprintf(dirTouch+"/zero_re%d", i)
		_ = KFile.Unlink(filename)
	}
}

func TestFile_CopyFile(t *testing.T) {
	var res int64
	var err error

	//源文件不存在
	res, err = KFile.CopyFile(fileNone, imgCopy, FILE_COVER_IGNORE)
	assert.NotNil(t, err)

	//忽略已存在的
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//覆盖已存在的-允许
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_ALLOW)
	assert.Greater(t, res, int64(0))

	//覆盖已存在的-忽略
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_IGNORE)
	assert.Nil(t, err)
	assert.Equal(t, int64(0), res)

	//覆盖已存在的-拒绝
	res, err = KFile.CopyFile(imgPng, imgCopy, FILE_COVER_DENY)
	assert.NotNil(t, err)

	//源和目标文件相同
	res, err = KFile.CopyFile(imgPng, imgPng, FILE_COVER_ALLOW)
	assert.Equal(t, int64(0), res)
	assert.Nil(t, err)

	//拷贝大文件
	KFile.Touch(touchfile, 2097152)
	res, err = KFile.CopyFile(touchfile, copyfile, FILE_COVER_ALLOW)

	//目标为空
	res, err = KFile.CopyFile(imgPng, "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)

	//源非正常文件
	res, err = KFile.CopyFile(".", "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func BenchmarkFile_CopyFile(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/diglett_copy_%d.png", i)
		_, _ = KFile.CopyFile(imgPng, des, FILE_COVER_ALLOW)
	}
}

func TestFile_FastCopy(t *testing.T) {
	var res int64
	var err error

	res, err = KFile.FastCopy(imgJpg, fastcopyfile)
	assert.Greater(t, res, int64(0))

	//源文件不存在
	res, err = KFile.FastCopy(fileNone, fastcopyfile)
	assert.NotNil(t, err)

	//目标为空
	res, err = KFile.FastCopy(imgJpg, "")
	assert.NotNil(t, err)
}

func BenchmarkFile_FastCopy(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/fast_copy_%d", i)
		_, _ = KFile.FastCopy(imgJpg, des)
	}
}

func TestFile_CopyLink(t *testing.T) {
	var err error

	//创建链接文件
	if !KFile.IsExist(fileLink) {
		_ = os.Symlink(filePubPem, fileLink)
	}

	//源链接不存在
	err = KFile.CopyLink(fileNone, copyLink, FILE_COVER_IGNORE)
	assert.NotNil(t, err)

	//源和目标相同
	err = KFile.CopyLink(fileLink, fileLink, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//目标为空
	err = KFile.CopyLink(fileLink, "", FILE_COVER_IGNORE)
	assert.NotNil(t, err)

	//成功拷贝
	err = KFile.CopyLink(fileLink, copyLink, FILE_COVER_ALLOW)
	assert.Nil(t, err)
	err = KFile.CopyLink(fileLink, copyLink2, FILE_COVER_ALLOW)
	assert.Nil(t, err)

	//已存在-忽略
	err = KFile.CopyLink(fileLink, copyLink, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//已存在-拒绝
	err = KFile.CopyLink(fileLink, copyLink, FILE_COVER_DENY)
	assert.NotNil(t, err)

	//已存在-允许
	err = KFile.CopyLink(fileLink, copyLink, FILE_COVER_ALLOW)
	assert.Nil(t, err)
}

func BenchmarkFile_CopyLink(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirLink+"/lnk_%d.copy", i)
		_ = KFile.CopyLink(fileLink, des, FILE_COVER_ALLOW)
	}
}

func TestFile_CopyDir(t *testing.T) {
	var res int64
	var err error

	//源和目标相同
	res, err = KFile.CopyDir(dirVendor, dirVendor, FILE_COVER_ALLOW)
	assert.Equal(t, int64(0), res)
	assert.Nil(t, err)

	//源不存在
	res, err = KFile.CopyDir(fileNone, dirTdat, FILE_COVER_IGNORE)
	assert.NotNil(t, err)

	//源不是目录
	res, err = KFile.CopyDir(fileMd, dirTdat, FILE_COVER_ALLOW)
	assert.NotNil(t, err)

	//忽略已存在的
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_IGNORE)
	assert.Nil(t, err)

	//覆盖已存在的
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_ALLOW)
	assert.Nil(t, err)

	//禁止覆盖
	res, err = KFile.CopyDir(dirVendor, dirTdat, FILE_COVER_DENY)
	assert.Equal(t, int64(0), res)

	//目标为空
	res, err = KFile.CopyDir(dirVendor, "", FILE_COVER_ALLOW)
	assert.NotNil(t, err)
}

func BenchmarkFile_CopyDir(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/copydir_%d", i)
		_, _ = KFile.CopyDir(dirDoc, des, FILE_COVER_ALLOW)
	}
}

func TestFile_DelDir(t *testing.T) {
	var err error
	var chk bool

	//非目录
	err = KFile.DelDir(fileMd, false)
	assert.NotNil(t, err)

	//清空目录
	err = KFile.DelDir(dirCopy, false)
	chk = KFile.IsDir(dirCopy)
	assert.Nil(t, err)
	assert.True(t, chk)

	//删除目录
	err = KFile.DelDir(dirNew, true)
	chk = KFile.IsDir(dirNew)
	assert.Nil(t, err)
	assert.False(t, chk)
}

func BenchmarkFile_DelDir(b *testing.B) {
	b.ResetTimer()
	var des string
	for i := 0; i < b.N; i++ {
		des = fmt.Sprintf(dirCopy+"/copydir_%d", i)
		_ = KFile.DelDir(des, true)
	}
}

func TestFile_Img2Base64(t *testing.T) {
	var res string
	var err error

	//非图片
	res, err = KFile.Img2Base64(fileMd)
	assert.NotNil(t, err)

	//图片不存在
	res, err = KFile.Img2Base64(imgNone)
	assert.NotNil(t, err)

	//png
	res, err = KFile.Img2Base64(imgPng)
	assert.Nil(t, err)
	assert.Contains(t, res, "png")

	//jpg
	res, err = KFile.Img2Base64(imgJpg)
	assert.Nil(t, err)
	assert.Contains(t, res, "jpg")
}

func BenchmarkFile_Img2Base64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Img2Base64(imgPng)
	}
}

func TestFile_FileTree(t *testing.T) {
	var res []string

	//显示全部
	res = KFile.FileTree(dirVendor, FILE_TREE_ALL, true)
	assert.NotEmpty(t, res)

	//仅目录
	res = KFile.FileTree(dirVendor, FILE_TREE_DIR, true)
	assert.NotEmpty(t, res)

	//仅文件
	res = KFile.FileTree(dirTdat, FILE_TREE_FILE, true)
	assert.NotEmpty(t, res)
	assert.GreaterOrEqual(t, len(res), 10)

	//不递归
	res = KFile.FileTree(dirCurr, FILE_TREE_DIR, false)
	assert.GreaterOrEqual(t, len(res), 4)

	//文件过滤
	res = KFile.FileTree(dirCurr, FILE_TREE_FILE, true, func(s string) bool {
		ext := KFile.GetExt(s)
		return ext == "go"
	})
	assert.NotEmpty(t, res)
}

func BenchmarkFile_FileTree(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FileTree(dirCurr, FILE_TREE_ALL, false)
	}
}

func TestFile_FormatDir(t *testing.T) {
	var res string

	res = KFile.FormatDir(pathTes3)
	assert.NotContains(t, res, "\\")

	//win格式
	res = KFile.FormatDir(pathTes2)
	assert.Equal(t, 1, strings.Count(res, ":"))

	//空目录
	res = KFile.FormatDir("")
	assert.Empty(t, res)
}

func BenchmarkFile_FormatDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FormatDir(pathTes3)
	}
}

func TestFile_FormatPath(t *testing.T) {
	var res string

	res = KFile.FormatPath(pathTes1)
	assert.NotContains(t, res, ":")

	res = KFile.FormatPath(fileGmod)
	assert.Equal(t, res, fileGmod)

	res = KFile.FormatPath(fileGo)
	assert.Equal(t, res, fileGo)

	res = KFile.FormatPath(pathTes3)
	assert.NotContains(t, res, "\\")

	//win格式
	res = KFile.FormatPath(pathTes2)
	assert.Equal(t, 1, strings.Count(res, ":"))

	//空路径
	res = KFile.FormatPath("")
	assert.Empty(t, res)
}

func BenchmarkFile_FormatPath(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.FormatPath(pathTes2)
	}
}

func TestFile_Md5File_Md5Reader(t *testing.T) {
	var res1, res2 string
	var err error

	fh1, _ := os.Open(fileMd)
	fh2, _ := os.Open(fileMd)
	defer func() {
		_ = fh1.Close()
		_ = fh2.Close()
	}()

	//16
	res1, err = KFile.Md5File(fileMd, 16)
	assert.NotEmpty(t, res1)
	assert.Nil(t, err)

	res2, err = KFile.Md5Reader(fh1, 16)
	assert.NotEmpty(t, res2)
	assert.Nil(t, err)
	assert.Equal(t, res1, res2)

	//32
	res1, err = KFile.Md5File(fileMd, 32)
	assert.NotEmpty(t, res1)
	assert.Nil(t, err)

	res2, err = KFile.Md5Reader(fh2, 32)
	assert.NotEmpty(t, res2)
	assert.Nil(t, err)
	assert.Equal(t, res1, res2)

	//不存在的文件
	_, err = KFile.Md5File(fileNone, 32)
	assert.NotNil(t, err)
}

func BenchmarkFile_Md5File(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Md5File(fileMd, 32)
	}
}

func BenchmarkFile_Md5Reader(b *testing.B) {
	b.ResetTimer()
	fh, _ := os.Open(fileMd)
	defer func() {
		_ = fh.Close()
	}()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Md5Reader(fh, 32)
	}
}

func TestFile_ShaXFile_ShaXReader(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res1, res2 string
	var err error

	fh1, _ := os.Open(fileGmod)
	fh2, _ := os.Open(fileGmod)
	fh3, _ := os.Open(fileGmod)
	defer func() {
		_ = fh1.Close()
		_ = fh2.Close()
		_ = fh3.Close()
	}()

	//1
	res1, err = KFile.ShaXFile(fileGmod, 1)
	assert.NotEmpty(t, res1)
	assert.Nil(t, err)

	res2, err = KFile.ShaXReader(fh1, 1)
	assert.NotEmpty(t, res2)
	assert.Nil(t, err)
	assert.Equal(t, res1, res2)

	//256
	res1, err = KFile.ShaXFile(fileGmod, 256)
	assert.NotEmpty(t, res1)
	assert.Nil(t, err)

	res2, err = KFile.ShaXReader(fh2, 256)
	assert.NotEmpty(t, res2)
	assert.Nil(t, err)
	assert.Equal(t, res1, res2)

	//512
	res1, err = KFile.ShaXFile(fileGmod, 512)
	assert.NotEmpty(t, res1)
	assert.Nil(t, err)

	res2, err = KFile.ShaXReader(fh3, 512)
	assert.NotEmpty(t, res2)
	assert.Nil(t, err)
	assert.Equal(t, res1, res2)

	//文件不存在
	_, err = KFile.ShaXFile(fileNone, 512)
	assert.NotNil(t, err)

	//err x
	_, err = KFile.ShaXFile(fileGmod, 32)
}

func BenchmarkFile_ShaXFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ShaXFile(fileGmod, 256)
	}
}

func BenchmarkFile_ShaXReader(b *testing.B) {
	b.ResetTimer()
	fh, _ := os.Open(fileMd)
	defer func() {
		_ = fh.Close()
	}()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.ShaXReader(fh, 256)
	}
}

func TestFile_Pathinfo(t *testing.T) {
	var res map[string]string

	//所有信息
	res = KFile.Pathinfo(imgPng, -1)
	assert.Equal(t, 4, len(res))

	//仅目录
	res = KFile.Pathinfo(imgPng, 1)

	//仅基础名(文件+扩展)
	res = KFile.Pathinfo(imgPng, 2)

	//仅扩展名
	res = KFile.Pathinfo(imgPng, 4)

	//仅文件名
	res = KFile.Pathinfo(imgPng, 8)

	//目录+基础名
	res = KFile.Pathinfo(imgPng, 3)

	//特殊类型
	res = KFile.Pathinfo(fileGitkee, -1)
	assert.Empty(t, res["filename"])

	//文件名没有后缀
	res = KFile.Pathinfo(rootFile1, -1)
	assert.Equal(t, res["basename"], res["filename"])
}

func BenchmarkFile_Pathinfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.Pathinfo(imgPng, -1)
	}
}

func TestFile_Basename(t *testing.T) {
	var res string

	res = KFile.Basename(fileMd)
	assert.Equal(t, "README.md", res)

	res = KFile.Basename(fileNone)
	assert.Equal(t, "none", res)

	res = KFile.Basename("")
	assert.NotEmpty(t, res)
	assert.Equal(t, ".", res)
}

func BenchmarkFile_Basename(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.Basename(fileDante)
	}
}

func TestFile_Dirname(t *testing.T) {
	var res string

	res = KFile.Dirname(changLog)
	assert.Equal(t, "docs", res)

	res = KFile.Dirname("")
	assert.NotEmpty(t, res)
	assert.Equal(t, ".", res)
}

func BenchmarkFile_Dirname(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.Dirname(fileSongs)
	}
}

func TestFile_GetModTime(t *testing.T) {
	var res int64

	res = KFile.GetModTime(fileMd)
	assert.Greater(t, res, int64(0))

	//不存在的文件
	res = KFile.GetModTime(fileNone)
	assert.Equal(t, res, int64(0))

	//空路径
	res = KFile.GetModTime(fileNone)
	assert.Equal(t, res, int64(0))
}

func BenchmarkFile_GetModTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.GetModTime(fileMd)
	}
}

func TestFile_Glob(t *testing.T) {
	var res []string
	var err error

	res, err = KFile.Glob("*test.go")
	assert.NotEmpty(t, res)
	assert.Nil(t, err)
}

func BenchmarkFile_Glob(b *testing.B) {
	b.ResetTimer()
	pattern := "*test.go"
	for i := 0; i < b.N; i++ {
		_, _ = KFile.Glob(pattern)
	}
}

func TestFile_SafeFileName(t *testing.T) {
	var res string

	res = KFile.SafeFileName(pathTes4)
	assert.Equal(t, `123456789-ASDF.html`, res)

	res = KFile.SafeFileName(pathTes5)
	assert.Equal(t, `test.go`, res)

	res = KFile.SafeFileName(pathTes6)
	assert.Equal(t, `Hello-World.txt`, res)

	res = KFile.SafeFileName("")
	assert.Equal(t, ".", res)
}

func BenchmarkFile_SafeFileName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.SafeFileName(pathTes4)
	}
}

func TestFile_TarGzUnTarGz(t *testing.T) {
	var res1, res2 bool
	var err1, err2 error

	//打包
	patterns := []string{".*.md", ".*.yml", ".*_test.go"}
	res1, err1 = KFile.TarGz(dirVendor, targzfile1, patterns...)
	assert.True(t, res1)
	assert.Nil(t, err1)

	//解压
	res2, err2 = KFile.UnTarGz(targzfile1, untarpath1)
	assert.True(t, res2)
	assert.Nil(t, err2)

	//打包不存在的目录
	res1, err1 = KFile.TarGz(fileNone, targzfile2)
	assert.False(t, res1)
	assert.NotNil(t, err1)

	//解压非tar格式的文件
	res2, err2 = KFile.UnTarGz(fileDante, untarpath1)
	assert.False(t, res2)
	assert.NotNil(t, err2)

	//解压不存在的文件
	res2, err2 = KFile.UnTarGz(fileNone, untarpath1)
	assert.False(t, res2)
	assert.NotNil(t, err2)
}

func BenchmarkFile_TarGz(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := fmt.Sprintf(dirTdat+"/targz/test_%d.tar.gz", i)
		_, _ = KFile.TarGz(dirDoc, dst)
	}
}

func BenchmarkFile_UnTarGz(b *testing.B) {
	b.ResetTimer()
	var src, dst string
	for i := 0; i < b.N; i++ {
		src = fmt.Sprintf(dirTdat+"/targz/test_%d.tar.gz", i)
		dst = fmt.Sprintf(dirTdat+"/targz/test_%d", i)
		_, _ = KFile.UnTarGz(src, dst)
	}
}

func TestFile_ChmodBatch(t *testing.T) {
	var res bool
	var tmp string

	for i := 0; i < 10; i++ {
		tmp = fmt.Sprintf(dirChmod+"/tmp_%d", i)
		KFile.Touch(tmp, 0)
	}

	res = KFile.ChmodBatch(dirChmod, 0766, 0755)
	assert.True(t, res)

	//不存在的路径
	res = KFile.ChmodBatch(fileNone, 0777, 0766)
	assert.False(t, res)
}

func BenchmarkFile_ChmodBatch(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KFile.ChmodBatch(dirDoc, 0777, 0766)
	}
}

func TestFile_CountLines(t *testing.T) {
	var res int
	var err error

	res, err = KFile.CountLines(fileDante, 0)
	assert.Equal(t, 19567, res)
	assert.Nil(t, err)

	//非文本文件
	res, err = KFile.CountLines(imgJpg, 8)
	assert.Greater(t, res, 0)
	assert.Nil(t, err)

	//不存在的文件
	res, err = KFile.CountLines(fileNone, 0)
	assert.Equal(t, -1, res)
	assert.NotNil(t, err)
}

func BenchmarkFile_CountLines(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KFile.CountLines(fileMd, 0)
	}
}

func TestFile_ZipIszipUnzip(t *testing.T) {
	var res bool
	var err error

	//空输入
	res, err = KFile.Zip(zipfile1)
	assert.False(t, res)
	assert.NotNil(t, err)

	//源文件不存在
	res, err = KFile.Zip(zipfile1, fileNone)
	assert.False(t, res)
	assert.NotNil(t, err)

	res, err = KFile.Zip(zipfile1, fileMd, fileGo, fileDante, dirDoc)
	assert.True(t, res)
	assert.Nil(t, err)

	//判断
	res, err = KFile.IsZip(zipfile1)
	assert.True(t, res)

	//后缀名不符
	res, err = KFile.IsZip(fileNone)
	assert.False(t, res)
	assert.Nil(t, err)

	//无权限检查
	res, err = KFile.IsZip(rootFile2)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.False(t, res)
		assert.NotNil(t, err)
	}

	//解压
	res, err = KFile.UnZip(zipfile1, unzippath1)
	assert.True(t, res)
	assert.Nil(t, err)

	//解压非zip文件
	res, err = KFile.UnZip(imgJpg, unzippath1)
	assert.False(t, res)
	assert.NotNil(t, err)

	//解压不存在文件
	res, err = KFile.UnZip(fileNone, unzippath1)
	assert.False(t, res)
	assert.NotNil(t, err)
}

func BenchmarkFile_Zip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := fmt.Sprintf(dirTdat+"/zip/test_%d.zip", i)
		_, _ = KFile.Zip(dst, dirDoc)
	}
}

func BenchmarkFile_IsZip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dst := fmt.Sprintf(dirTdat+"/zip/test_%d.zip", i)
		_, _ = KFile.IsZip(dst)
	}
}

func BenchmarkFile_UnZip(b *testing.B) {
	b.ResetTimer()
	var src, dst string
	for i := 0; i < b.N; i++ {
		src = fmt.Sprintf(dirTdat+"/zip/test_%d.zip", i)
		dst = fmt.Sprintf(dirTdat+"/zip/unzip/test_%d", i)
		_, _ = KFile.UnZip(src, dst)
	}
}
