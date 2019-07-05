package gohelper

import (
	"strings"
	"testing"
)

func TestNl2br(t *testing.T) {
	str := `hello
world!
你好！`
	res := KStr.Nl2br(str)
	if !strings.Contains(res, "<br />") {
		t.Error("Nl2br fail")
		return
	}
	_ = KStr.Nl2br("")
}

func BenchmarkNl2br(b *testing.B) {
	b.ResetTimer()
	str := `hello
world!
你好！`
	for i := 0; i < b.N; i++ {
		_ = KStr.Nl2br(str)
	}
}

func TestStripTags(t *testing.T) {
	str := `
<h1>Hello world!</h1>
<script>alert('你好！')</scripty>
`
	res := KStr.StripTags(str)
	if strings.Contains(res, "<script>") {
		t.Error("StripTags fail")
		return
	}
	_ = KStr.StripTags("")
}

func BenchmarkStripTags(b *testing.B) {
	b.ResetTimer()
	str := `
<h1>Hello world!</h1>
<script>alert('你好！')</scripty>
`
	for i := 0; i < b.N; i++ {
		_ = KStr.StripTags(str)
	}
}

func TestStringIsBinary(t *testing.T) {
	filename := "./file.go"
	if KFile.IsBinary(filename) {
		t.Error("file isn`t binary")
		return
	}

	file2 := "/usr/lib/golang/bin/go"
	if !KFile.IsBinary(file2) {
		t.Error("file isn`t binary")
		return
	}
}

func BenchmarkStringIsBinary(b *testing.B) {
	b.ResetTimer()
	filename := "./README.md"
	for i := 0; i < b.N; i++ {
		KFile.IsBinary(filename)
	}
}

func TestStringMd5(t *testing.T) {
	str := ""
	res1 := KStr.Md5(str, 32)
	res2 := KStr.Md5(str, 16)
	if res1 != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Error("string Md5 fail")
		return
	}
	if !strings.Contains(res1, res2) {
		t.Error("string Md5 fail")
		return
	}
}

func BenchmarkStringMd5(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		_ = KStr.Md5(str, 32)
	}
}
