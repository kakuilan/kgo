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
	cont, _ := KFile.GetContents("./file.go")
	if KStr.IsBinary(cont) {
		t.Error("str isn`t binary")
		return
	}
	_, _ = KFile.GetContents("")
}

func BenchmarkStringIsBinary(b *testing.B) {
	b.ResetTimer()
	str := "hello"
	for i := 0; i < b.N; i++ {
		KFile.IsBinary(str)
	}
}

func TestIsLetter(t *testing.T) {
	res := KStr.IsLetter("hello")
	if !res {
		t.Error("IsLetter fail")
		return
	}
	KStr.IsLetter("")
	KStr.IsLetter("123")
}

func BenchmarkIsLetter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsLetter("hello")
	}
}

func TestIsNumeric(t *testing.T) {
	res := KStr.IsNumeric("123.456")
	res2 := KStr.IsNumeric("123")
	if !res || !res2 {
		t.Error("IsNumeric fail")
		return
	}
	KStr.IsNumeric("")
}

func BenchmarkIsNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsNumeric("123.456")
	}
}

func TestIsInt(t *testing.T) {
	res := KStr.IsInt("123.456")
	res2 := KStr.IsInt("123")
	if res || !res2 {
		t.Error("IsInt fail")
		return
	}
	KStr.IsInt("")
}

func TestHasChinese(t *testing.T) {
	res := KStr.HasChinese("123.456")
	res2 := KStr.HasChinese("hello你好")
	if res || !res2 {
		t.Error("HasChinese fail")
		return
	}
	KStr.HasChinese("")
}

func BenchmarkHasChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasChinese("hello你好")
	}
}

func TestIsChinese(t *testing.T) {
	res := KStr.IsChinese("hello你好")
	res2 := KStr.IsChinese("你好世界")
	if res || !res2 {
		t.Error("IsChinese fail")
		return
	}
	KStr.IsChinese("")
}

func BenchmarkIsChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsChinese("你好世界")
	}
}

func BenchmarkIsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsInt("123")
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

func TestRandomAlpha(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_ALPHA)
	if !KStr.IsLetter(res) {
		t.Error("RandomAlpha fail")
		return
	}
	KStr.Random(0, RAND_STRING_ALPHA)
	KStr.Random(1, 99)
}

func BenchmarkRandomAlpha(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_ALPHA)
	}
}

func TestRandomNumeric(t *testing.T) {
	str := KStr.Random(8, RAND_STRING_NUMERIC)
	if !KStr.IsNumeric(str) {
		t.Error("RandomNumeric fail")
		return
	}
}

func BenchmarkRandomNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_NUMERIC)
	}
}

func TestRandomAlphanum(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_ALPHANUM)
	if len(res) != 8 {
		t.Error("RandomAlphanum fail")
		return
	}
}

func BenchmarkRandomAlphanum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_ALPHANUM)
	}
}

func TestRandomSpecial(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_SPECIAL)
	if len(res) != 8 {
		t.Error("RandomSpecial fail")
		return
	}
}

func BenchmarkRandomSpecial(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_SPECIAL)
	}
}

func TestRandomChinese(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_CHINESE)
	println(res)
}

func BenchmarkRandomChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_CHINESE)
	}
}
