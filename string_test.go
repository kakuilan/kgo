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
	if !KStr.IsChinese(res) {
		t.Error("RandomChinese fail")
		return
	}
}

func BenchmarkRandomChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_CHINESE)
	}
}
