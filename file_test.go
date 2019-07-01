package gohelper

import (
	"testing"
	"fmt"
	"os/exec"
)

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if KFile.GetExt(filename) !="go" {
		t.Error("file extension error")
		return
	}
}

func BenchmarkGetExt(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.GetExt(filename)
	}
}

func TestGetGetContents(t *testing.T) {
	filename := "./file.go"
	cont,_ := KFile.GetContents(filename)
	if  cont=="" {
		t.Error("file get contents error")
		return
	}
}

func BenchmarkGetContents(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		_,_ = KFile.GetContents(filename)
	}
}

func TestFileSize(t *testing.T)  {
	filename := "./file.go"
	if KFile.FileSize(filename) <=0 {
		t.Error("file size error")
		return
	}
}

func BenchmarkFileSize(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.FileSize(filename)
	}
}

func TestDirSize(t *testing.T)  {
	dirpath := "./"
	size := KFile.DirSize(dirpath)
	if size==0 {
		t.Error("dir size error")
		return
	}
}

func BenchmarkDirSize(b *testing.B) {
	b.ResetTimer()
	dirpath := "./"
	for i:=0;i<b.N;i++{
		_ = KFile.DirSize(dirpath)
	}
}

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkIsExist(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsExist(filename)
	}
}

func TestIsWritable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsWritable(filename) {
		t.Error("file can not write")
		return
	}
}

func BenchmarkIsWritable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsWritable(filename)
	}
}

func TestIsReadable(t *testing.T) {
	filename := "./README.md"
	if !KFile.IsReadable(filename) {
		t.Error("file can not read")
		return
	}
}

func BenchmarkIsReadable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsReadable(filename)
	}
}

func TestIsFile(t *testing.T) {
	filename := "./file.go"
	if !KFile.IsFile(filename) {
		t.Error("isn`t a file")
		return
	}
}

func BenchmarkIsFile(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsFile(filename)
	}
}

func TestIsDir(t *testing.T) {
	dirname := "./"
	if !KFile.IsDir(dirname) {
		t.Error("isn`t a dir")
		return
	}
}

func BenchmarkIsDir(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsDir(filename)
	}
}

func TestFileIsBinary(t *testing.T) {
	filename := "./file.go"
	if KFile.IsBinary(filename) {
		t.Error("file isn`t binary")
		return
	}
}

func BenchmarkFileIsBinary(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		KFile.IsBinary(filename)
	}
}

func TestIsImg(t *testing.T) {
	filename := "./testdata/diglett.png"
	if !KFile.IsImg(filename) {
		t.Error("file isn`t img")
		return
	}
}

func BenchmarkIsImg(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i:=0;i<b.N;i++{
		KFile.IsImg(filename)
	}
}

func TestAbsPath(t *testing.T) {
	filename := "./testdata/diglett.png"
	abspath := KFile.AbsPath(filename)
	if !KFile.IsExist(abspath) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkAbsPath(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i:=0;i<b.N;i++{
		KFile.AbsPath(filename)
	}
}

func TestCopyFile(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"
	num, err := KFile.CopyFile(src, des, FCOVER_ALLOW)
	if err != nil || num ==0 {
		t.Error("copy file fail")
		return
	}
}

func BenchmarkCopyFile(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i:=0;i<b.N;i++{
		des = fmt.Sprintf("./testdata/diglett_copy_%d.png", i)
		_,_ = KFile.CopyFile(src, des, FCOVER_ALLOW)
	}
}

func TestFastCopy(t *testing.T) {
	src := "./testdata/diglett.png"
	des := "./testdata/diglett_copy.png"

	num, err := KFile.FastCopy(src, des)
	if err != nil || num ==0 {
		t.Error("copy file fail")
		return
	}
}

func BenchmarkFastCopy(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett.png"
	des := ""
	for i:=0;i<b.N;i++{
		des = fmt.Sprintf("./testdata/diglett_fast_%d.png", i)
		_,_ = KFile.FastCopy(src, des)
	}
}

func TestCopyLink(t *testing.T) {
	cmd := exec.Command("/bin/bash", "-c", "ln -sf ./testdata/diglett.png ./testdata/diglett-lnk")
	_ = cmd.Run()

	src := "./testdata/diglett-lnk"
	des := "./testdata/diglett-lnk.copy"

	err := KFile.CopyLink(src, des)
	if err != nil {
		t.Error("copy link fail" + err.Error())
		return
	}
}

func BenchmarkCopyLink(b *testing.B) {
	b.ResetTimer()
	src := "./testdata/diglett-lnk"
	des := ""
	for i:=0;i<b.N;i++{
		des = fmt.Sprintf("./testdata/diglett-lnk_%d.copy", i)
		_ = KFile.CopyLink(src, des)
	}
}

func TestCopyDir(t *testing.T) {
	src := "./testdata"
	des := "./test"

	num, err := KFile.CopyDir(src, des, FCOVER_ALLOW)
	if err != nil || num ==0 {
		t.Error("copy directory fail")
		return
	}
}

func BenchmarkCopyDir(b *testing.B) {
	b.ResetTimer()
	src := "./testdata"
	des := "./test"
	for i:=0;i<b.N;i++{
		des = fmt.Sprintf("./test/copy_%d", i)
		_,_ = KFile.CopyDir(src, des, FCOVER_ALLOW)
	}
}