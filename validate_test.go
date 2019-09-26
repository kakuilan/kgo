package kgo

import (
	"fmt"
	"math"
	"testing"
)

func TestStringIsBinary(t *testing.T) {
	cont, _ := KFile.GetContents("./file.go")
	if KConv.IsBinary(string(cont)) {
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

func TestIsLetters(t *testing.T) {
	res := KStr.IsLetters("hello")
	if !res {
		t.Error("IsLetters fail")
		return
	}
	KStr.IsLetters("")
	KStr.IsLetters("123")
}

func BenchmarkIsLetters(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsLetters("hello")
	}
}

func TestIsNumeric(t *testing.T) {
	res1 := KConv.IsNumeric(123)
	res2 := KConv.IsNumeric("123.456")
	res3 := KConv.IsNumeric("-0.56")
	res4 := KConv.IsNumeric(45.678)
	if !res1 || !res2 || !res3 || !res4 {
		t.Error("IsNumeric fail")
		return
	}

	var sli []int
	KConv.IsNumeric("")
	KConv.IsNumeric(sli)
}

func BenchmarkIsNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsNumeric("123.456")
	}
}

func TestIsInt(t *testing.T) {
	res1 := KConv.IsInt(123)
	res2 := KConv.IsInt("123")
	res3 := KConv.IsInt("-45")
	if !res1 || !res2 || !res3 {
		t.Error("IsInt fail")
		return
	}
	var sli []int
	KConv.IsInt("")
	KConv.IsInt(sli)
}

func BenchmarkIsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsInt("123")
	}
}

func TestIsFloat(t *testing.T) {
	res1 := KConv.IsFloat(123.0)
	res2 := KConv.IsFloat("123.4")
	res3 := KConv.IsFloat("-45.6")
	if !res1 || !res2 || !res3 {
		t.Error("IsFloat IsFloat")
		return
	}

	var sli []int
	KConv.IsFloat("")
	KConv.IsFloat(sli)
}

func BenchmarkIsFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsFloat("123.45")
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

func TestIsJSON(t *testing.T) {
	chk1 := KStr.IsJSON("hello你好")
	chk2 := KStr.IsJSON(`{"id":"1"}`)
	if chk1 || !chk2 {
		t.Error("IsJSON fail")
		return
	}
	KStr.IsJSON("")
}

func BenchmarkIsJSON(b *testing.B) {
	b.ResetTimer()
	str := `{"key1": "value1"}, {"key2": "value2"}`
	for i := 0; i < b.N; i++ {
		KStr.IsJSON(str)
	}
}

func TestIsIPv4(t *testing.T) {
	res1 := KStr.IsIPv4("")
	res2 := KStr.IsIPv4("8.9.10.11")
	res3 := KStr.IsIPv4("192.168.0.1:80")
	res4 := KStr.IsIPv4("::FFFF:C0A8:1")
	res5 := KStr.IsIPv4("fe80::2c04:f7ff:feaa:33b7")

	if res1 || !res2 || res3 || res4 || res5 {
		t.Error("IsIPv4 fail")
		return
	}
}

func BenchmarkIsIPv4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIPv4("8.9.10.11")
	}
}

func TestIsIPv6(t *testing.T) {
	res1 := KStr.IsIPv6("")
	res2 := KStr.IsIPv6("8.9.10.11")
	res3 := KStr.IsIPv6("192.168.0.1:80")
	res4 := KStr.IsIPv6("::FFFF:C0A8:1")
	res5 := KStr.IsIPv6("fe80::2c04:f7ff:feaa:33b7")

	if res1 || res2 || res3 || !res4 || !res5 {
		t.Error("IsIPv6 fail")
		return
	}
}

func BenchmarkIsIPv6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIPv6("fe80::2c04:f7ff:feaa:33b7")
	}
}

func TestIsArrayOrSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"

	res1 := KArr.IsArrayOrSlice(arr, 1)
	res2 := KArr.IsArrayOrSlice(arr, 2)
	res3 := KArr.IsArrayOrSlice(arr, 3)
	if res1 != 10 || res2 != -1 || res3 != 10 {
		t.Error("IsArrayOrSlice fail")
		return
	}

	res4 := KArr.IsArrayOrSlice(sli, 1)
	res5 := KArr.IsArrayOrSlice(sli, 2)
	res6 := KArr.IsArrayOrSlice(sli, 3)
	if res4 != -1 || res5 != 5 || res6 != 5 {
		t.Error("IsArrayOrSlice fail")
		return
	}

	KArr.IsArrayOrSlice(sli, 6)
}

func BenchmarkIsArrayOrSlice(b *testing.B) {
	b.ResetTimer()
	var arr = [10]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		KArr.IsArrayOrSlice(arr, 1)
	}
}

func TestIsMap(t *testing.T) {
	mp := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	res1 := KArr.IsMap(mp)
	res2 := KArr.IsMap(123)
	if !res1 || res2 {
		t.Error("IsMap fail")
		return
	}
}

func BenchmarkIsMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsMap("hello")
	}
}

func TestIsNan(t *testing.T) {
	res1 := KNum.IsNan(math.Acos(1.01))
	res2 := KNum.IsNan(123.456)

	if !res1 || res2 {
		t.Error("IsNan fail")
		return
	}
}

func BenchmarkIsNan(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNan(123.456)
	}
}

func TestIsEmpty(t *testing.T) {
	var sli []int
	mp := make(map[string]int)
	var i uint = 0
	var val1 interface{} = &sli

	type myStru struct {
		conv LkkFileCover
		name string
	}
	var val2 myStru

	res1 := KConv.IsEmpty(nil)
	res2 := KConv.IsEmpty("")
	res3 := KConv.IsEmpty(sli)
	res4 := KConv.IsEmpty(mp)
	res5 := KConv.IsEmpty(false)
	res6 := KConv.IsEmpty(0)
	res7 := KConv.IsEmpty(i)
	res8 := KConv.IsEmpty(0.0)
	res9 := KConv.IsEmpty(val1)
	res10 := KConv.IsEmpty(val2)

	if !res1 || !res2 || !res3 || !res4 || !res5 || !res6 || !res7 || !res8 || res9 || !res10 {
		t.Error("IsEmpty fail")
		return
	}
}

func BenchmarkIsEmpty(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsEmpty("")
	}
}

func TestIsBool(t *testing.T) {
	res1 := KConv.IsBool(1)
	res2 := KConv.IsBool("hello")
	res3 := KConv.IsBool(false)
	if res1 || res2 || !res3 {
		t.Error("IsBool fail")
		return
	}
}

func BenchmarkIsBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsBool("hello")
	}
}

func TestIsHex(t *testing.T) {
	num1 := KConv.Dec2hex(1234)
	num2 := "0x" + num1
	res1 := KConv.IsHex(num1)
	res2 := KConv.IsHex(num2)
	res3 := KConv.IsHex("hello")
	res4 := KConv.IsHex("")
	if !res1 || !res2 || res3 || res4 {
		t.Error("IsHex fail")
		return
	}
}

func BenchmarkIsHex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsHex("4d2")
	}
}

func TestIsString(t *testing.T) {
	chk1 := KConv.IsString(123)
	chk2 := KConv.IsString("hello")
	if chk1 || !chk2 {
		t.Error("IsString fail")
		return
	}
}

func BenchmarkIsString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsString("hello")
	}
}

func TestIsByte(t *testing.T) {
	chk1 := KConv.IsByte("hello")
	chk2 := KConv.IsByte([]byte("hello"))
	if chk1 || !chk2 {
		t.Error("IsByte fail")
		return
	}
}

func BenchmarkIsByte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsByte([]byte("hello"))
	}
}
