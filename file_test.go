package gohelper

import (
	"testing"
)

func TestGetExt(t *testing.T) {
	filename := "./file.go"
	if File.GetExt(filename) !="go" {
		t.Error("file extension error")
		return
	}
}

func BenchmarkGetExt(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.GetExt(filename)
	}
}

func TestGetSize(t *testing.T)  {
	filename := "./file.go"
	if File.GetSize(filename) <=0 {
		t.Error("file size error")
		return
	}
}

func BenchmarkGetSize(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.GetSize(filename)
	}
}

func TestIsExist(t *testing.T) {
	filename := "./file.go"
	if !File.IsExist(filename) {
		t.Error("file not exist")
		return
	}
}

func BenchmarkIsExist(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsExist(filename)
	}
}

func TestIsWritable(t *testing.T) {
	filename := "./README.md"
	if !File.IsWritable(filename) {
		t.Error("file can not write")
		return
	}
}

func BenchmarkIsWritable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsWritable(filename)
	}
}

func TestIsReadable(t *testing.T) {
	filename := "./README.md"
	if !File.IsReadable(filename) {
		t.Error("file can not read")
		return
	}
}

func BenchmarkIsReadable(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsReadable(filename)
	}
}

func TestIsFile(t *testing.T) {
	filename := "./file.go"
	if !File.IsFile(filename) {
		t.Error("isn`t a file")
		return
	}
}

func BenchmarkIsFile(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsFile(filename)
	}
}

func TestIsDir(t *testing.T) {
	dirname := "./"
	if !File.IsDir(dirname) {
		t.Error("isn`t a dir")
		return
	}
}

func BenchmarkIsDir(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsDir(filename)
	}
}

func TestIsBinary(t *testing.T) {
	filename := "./file.go"
	if File.IsBinary(filename) {
		t.Error("file isn`t binary")
		return
	}
}

func BenchmarkIsBinary(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i:=0;i<b.N;i++{
		File.IsBinary(filename)
	}
}

func TestIsImg(t *testing.T) {
	filename := "./testdata/diglett.png"
	if !File.IsImg(filename) {
		t.Error("file isn`t img")
		return
	}
}

func BenchmarkIsImg(b *testing.B) {
	b.ResetTimer()
	filename := "./testdata/diglett.png"
	for i:=0;i<b.N;i++{
		File.IsImg(filename)
	}
}