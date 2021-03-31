package kgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
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
	err = KFile.WriteFile(rootFile1, bytsHello, 0777)
	if KOS.IsLinux() || KOS.IsMac() {
		assert.NotNil(t, err)
	}
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
	var res bool

	res = KFile.IsExecutable(fileNone)
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

	res = KFile.IsLink(changLog)
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

	//TODO true
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
	var err error

	err = KFile.Mkdir(dirNew, 0655)
	assert.Nil(t, err)
}

func BenchmarkFile_Mkdir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		dname := fmt.Sprintf(dirNew+"/tmp_%d", i)
		_ = KFile.Mkdir(dname, 0655)
	}
}
