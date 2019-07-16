package kgo

import "testing"

func TestStringIsBinary(t *testing.T) {
	cont, _ := KFile.GetContents("./file.go")
	if KStr.IsBinary(string(cont)) {
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

func BenchmarkIsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsInt("123")
	}
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
