package kgo

import (
	"fmt"
	"math"
	"strings"
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

func TestIsUtf8(t *testing.T) {
	str := "你好，世界！"
	chk1 := KStr.IsUtf8(str)

	if !chk1 {
		t.Error("IsUtf8 fail")
		return
	}

	gbk, _ := KStr.Utf8ToGbk([]byte(str))
	chk2 := KStr.IsUtf8(string(gbk))
	if chk2 {
		t.Error("IsUtf8 fail")
		return
	}
}

func BenchmarkIsUtf8(b *testing.B) {
	b.ResetTimer()
	str := "你好，世界！"
	for i := 0; i < b.N; i++ {
		KStr.IsUtf8(str)
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

func TestIsEmail(t *testing.T) {
	//无效的邮箱格式
	res1, _ := KStr.IsEmail("ç$€§/az@gmail.com", false)
	if res1 {
		t.Error("IsEmail fail")
		return
	}

	//有效的邮箱格式
	res2, _ := KStr.IsEmail("abc@abc123.com", false)
	if !res2 {
		t.Error("IsEmail fail")
		return
	}

	//无效的域名
	res3, _ := KStr.IsEmail("email@x-unkown-domain.com", true)
	if res3 {
		t.Error("IsEmail fail")
		return
	}

	//无效的账号
	res4, _ := KStr.IsEmail("unknown-user-123456789@gmail.com", true)
	if res4 {
		t.Error("IsEmail fail")
		return
	}

	//有效的账号
	res5, _ := KStr.IsEmail("copyright@github.com", true)
	host, _ := KOS.Hostname()
	//travis-ci不允许出站SMTP通信
	if !res5 && strings.Contains(host, "travis") == false {
		t.Error("IsEmail fail")
		return
	}
}

func BenchmarkIsEmail(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.IsEmail("abc@abc123.com", false)
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

func TestIsStruct(t *testing.T) {
	type sutTest struct {
		test string
	}
	sut := sutTest{test: "T"}

	chk1 := KConv.IsStruct("hello")
	chk2 := KConv.IsStruct(sut)
	chk3 := KConv.IsStruct(&sut)

	if chk1 || !chk2 || !chk3 {
		t.Error("IsStruct fail")
		return
	}
}

func BenchmarkIsStruct(b *testing.B) {
	b.ResetTimer()
	type sutTest struct {
		test string
	}
	sut := sutTest{test: "T"}
	for i := 0; i < b.N; i++ {
		KConv.IsStruct(&sut)
	}
}

func TestIsInterface(t *testing.T) {
	type inTest interface {
	}
	var in inTest

	chk1 := KConv.IsInterface("hello")
	chk2 := KConv.IsInterface(in)

	if chk1 || !chk2 {
		t.Error("IsInterface fail")
		return
	}
}

func BenchmarkIsInterface(b *testing.B) {
	b.ResetTimer()
	type inTest interface {
	}
	var in inTest
	for i := 0; i < b.N; i++ {
		KConv.IsInterface(in)
	}
}

func TestHasSpecialChar(t *testing.T) {
	str := "`~!@#$%^&*()_+-=:'|<>?,./\""
	res1 := KStr.HasSpecialChar(str)
	res2 := KStr.HasSpecialChar(str)
	// 掩码
	res3 := KStr.HasSpecialChar("Hello ៉៊់៌៍！")
	res4 := KStr.HasSpecialChar("hello world")
	res5 := KStr.HasSpecialChar("")

	if !res1 || !res2 || !res3 || res4 || res5 {
		t.Error("HasSpecialChar fail")
		return
	}
}

func BenchmarkHasSpecialChar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasSpecialChar("Hello ៉៊់៌៍！")
	}
}

func TestIsUrl(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar/#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"https://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
	}
	for _, test := range tests {
		actual := KUrl.IsUrl(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUrl(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KUrl.IsUrl("http//google.com")
	}
}

func TestIsMobile(t *testing.T) {
	res1 := KStr.IsMobile("12345678901")
	res2 := KStr.IsMobile("13712345678")
	res3 := KStr.IsMobile("")
	res4 := KStr.IsMobile("hello")

	if res1 || !res2 || res3 || res4 {
		t.Error("IsMobile fail")
	}
}

func BenchmarkIsMobile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMobile("13712345678")
	}
}

func TestIsTel(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"10086", false},
		{"010-88888888", true},
		{"021-87888822", true},
		{"0511-4405222", true},
		{"021-44055520-555", true},
		{"020-89571800-125", true},
		{"400-020-9800", true},
		{"400-999-0000", true},
		{"4006-589-589", true},
		{"4007005606", true},
		{"4000631300", true},
		{"400-6911195", true},
		{"800-4321", false},
		{"8004-321", false},
		{"8004321999", true},
		{"8008676014", true},
	}
	for _, test := range tests {
		actual := KStr.IsTel(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsTel(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkIsTel(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsTel("021-44055520-555")
	}
}

func TestIsPhone(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"10086", false},
		{"010-88888888", true},
		{"13712345678", true},
		{"hello", false},
	}
	for _, test := range tests {
		actual := KStr.IsPhone(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPhone(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkIsPhone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsPhone("13712345678")
	}
}

func TestIsDate2time(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"hello", false},
		{"0000", true},
		{"1970", true},
		{"1990-01", true},
		{"1990/01", true},
		{"1990-01-02", true},
		{"1990/01/02", true},
		{"1990-01-02 03", true},
		{"1990/01/02 03", true},
		{"1990-01-02 03:14", true},
		{"1990/01/02 03:14", true},
		{"1990-01-02 03:14:59", true},
		{"1990/01/02 03:14:59", true},
		{"2990-00-00 03:14:59", false},
	}
	for _, test := range tests {
		actual, tim := KTime.IsDate2time(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsDate2time(%q) to be %v, got %v %d", test.param, test.expected, actual, tim)
			return
		}
	}
}

func BenchmarkIsDate2time(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.IsDate2time("1990-01-02 03:14:59")
	}
}

func TestIsUrlExists(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"bing.com", false},
		{"http://none.localhost/", false},
		{"https://github.com/", true},
		{"https://github.com/dsfasdfasd/adsfasdfasdf", false},
	}
	for _, test := range tests {
		actual := KUrl.IsUrlExists(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUrlExists(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkIsUrlExists(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KUrl.IsUrlExists("https://github.com/")
	}
}

func TestIsCreditNo(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"hello", false},
		{"123123123", false},
		{"510723198006202551", true},
		{"34052419800101001x", true},
		{"511028199507215915", true},
		{"511028199502315915", false},
		{"53010219200508011X", true},
		{"99010219200508011X", false},
		{"130503670401001", true},
		{"370986890623212", true},
		{"370725881105149", true},
		{"370725881105996", true},
		{"35051419930513051X", false},
		{"44141419900430157X", false},
		{"110106209901012141", false},
		{"513436200011013606", true},
		{"51343620180101646X", true},
	}
	for _, test := range tests {
		chk, idNo := KStr.IsCreditNo(test.param)
		if chk != test.expected {
			t.Errorf("Expected IsCreditNo(%q) to be %v, got %v %s", test.param, test.expected, chk, idNo)
			return
		}
	}
}

func BenchmarkIsCreditNo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.IsCreditNo("51343620180101646X")
	}
}
