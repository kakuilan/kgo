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

func TestHasLetter(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"12345", false},
		{"http://none.localhost/", true},
		{"hello,world", true},
		{"PI:314159", true},
		{"你好，世界", false},
	}
	for _, test := range tests {
		actual := KStr.HasLetter(test.param)
		if actual != test.expected {
			t.Errorf("Expected HasLetter(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkHasLetter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasLetter("hello,world")
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

func TestIsEnglish(t *testing.T) {
	res1 := KStr.IsEnglish("", CASE_NONE)
	res2 := KStr.IsEnglish("1234", CASE_NONE)
	res3 := KStr.IsEnglish("hellWorld", CASE_NONE)
	res4 := KStr.IsEnglish("hellWorld", CASE_LOWER)
	res5 := KStr.IsEnglish("hellWorld", CASE_UPPER)
	res6 := KStr.IsEnglish("hellworld", CASE_LOWER)
	res7 := KStr.IsEnglish("HELLOWORLD", CASE_UPPER)
	res8 := KStr.IsEnglish("hehe", 9)

	if res1 || res2 || !res3 || res4 || res5 || !res6 || !res7 || !res8 {
		t.Error("IsEnglish fail")
		return
	}
}

func BenchmarkIsEnglishCASE_NONE(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsEnglish("hellWorld", CASE_LOWER)
	}
}

func BenchmarkIsEnglishCASE_LOWER(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsEnglish("hellWorld", CASE_LOWER)
	}
}

func BenchmarkIsEnglishCASE_UPPER(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsEnglish("hellWorld", CASE_UPPER)
	}
}

func TestHasEnglish(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"12345", false},
		{"http://none.localhost/", true},
		{"hello,world", true},
		{"PI:314159", true},
		{"你好，世界", false},
	}
	for _, test := range tests {
		actual := KStr.HasEnglish(test.param)
		if actual != test.expected {
			t.Errorf("Expected HasEnglish(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
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

func TestIsChineseName(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"hello world", false},
		{"赵武", true},
		{"赵武a", false},
		{"南宫先生", true},
		{"吉乃•阿衣·依扎嫫", true},
		{"古丽莎•卡迪尔", true},
		{"迪丽热巴.迪力木拉提", true},
	}
	for _, test := range tests {
		actual := KStr.IsChineseName(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsChineseName(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsChineseName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsChineseName("南宫先生")
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
	var res bool
	//长度验证
	res, _ = KStr.IsEmail("a@b.c", false)
	if res {
		t.Error("IsEmail fail")
		return
	}
	res, _ = KStr.IsEmail("hello-world@c", false)
	if res {
		t.Error("IsEmail fail")
		return
	}

	//无效的格式
	res, _ = KStr.IsEmail("ç$€§/az@gmail.com", false)
	if res {
		t.Error("IsEmail fail")
		return
	}

	//无效的域名
	res, _ = KStr.IsEmail("email@x-unkown-domain.com", true)
	if res {
		t.Error("IsEmail fail")
		return
	}

	//有效的账号
	res, err := KStr.IsEmail("copyright@github.com", true)
	if !res {
		t.Error("IsEmail fail")
		return
	} else if err != nil {
		t.Error("IsEmail fail:", err.Error())
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

func TestConvertIsEmpty(t *testing.T) {
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
		t.Error("Convert IsEmpty fail")
		return
	}
}

func BenchmarkConvertIsEmpty(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsEmpty("")
	}
}

func TestStrIsEmpty(t *testing.T) {
	chk1 := KStr.IsEmpty("")
	chk2 := KStr.IsEmpty("  ")
	chk3 := KStr.IsEmpty("hello")

	if !chk1 || !chk2 || chk3 {
		t.Error("String IsEmpty fail")
		return
	}
}

func BenchmarkStrIsEmpty(b *testing.B) {
	b.ResetTimer()
	str := "hello World"
	for i := 0; i < b.N; i++ {
		KStr.IsEmpty(str)
	}
}

func TestIsNil(t *testing.T) {
	var s []int
	chk1 := KConv.IsNil(nil)
	chk2 := KConv.IsNil(s)
	chk3 := KConv.IsNil("")

	if !chk1 || !chk2 || chk3 {
		t.Error("IsSha512 fail")
		return
	}
}

func BenchmarkIsNil(b *testing.B) {
	b.ResetTimer()
	var s []int
	for i := 0; i < b.N; i++ {
		KConv.IsNil(s)
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
	num1 := KConv.Dec2Hex(1234)
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
	//并行测试
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
		actual := KStr.IsUrl(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUrl(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUrl("http//google.com")
	}
}

func TestIsMobilecn(t *testing.T) {
	res1 := KStr.IsMobilecn("12345678901")
	res2 := KStr.IsMobilecn("13712345678")
	res3 := KStr.IsMobilecn("")
	res4 := KStr.IsMobilecn("hello")

	if res1 || !res2 || res3 || res4 {
		t.Error("IsMobilecn fail")
	}
}

func BenchmarkIsMobilecn(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMobilecn("13712345678")
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
		actual := KStr.IsUrlExists(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsUrlExists(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkIsUrlExists(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUrlExists("https://github.com/")
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

func TestIsAlphaNumeric(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"12345", true},
		{"helloworld", true},
		{"PI314159", true},
		{"你好，世界", false},
	}
	for _, test := range tests {
		actual := KStr.IsAlphaNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsAlphaNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkIsAlphaNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsAlphaNumeric("PI314159")
	}
}

func TestIsHexcolor(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"#ff", false},
		{"fff0", false},
		{"#ff12FG", false},
		{"CCccCC", true},
		{"fff", true},
		{"#f00", true},
	}
	for _, test := range tests {
		actual, color := KStr.IsHexcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHexcolor(%q) to be %v, got %v %s", test.param, test.expected, actual, color)
		}
	}
}

func BenchmarkIsHexcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.IsHexcolor("#ff12FG")
	}
}

func TestIsRGBcolor(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"rgb(0,31,255)", true},
		{"rgb(1,349,275)", false},
		{"rgb(01,31,255)", false},
		{"rgb(0.6,31,255)", false},
		{"rgba(0,31,255)", false},
		{"rgb(0,  31, 255)", true},
	}
	for _, test := range tests {
		actual := KStr.IsRGBcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsRGBcolor(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsRGBcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsRGBcolor("rgb(01,31,255)")
	}
}

func TestIsWhitespaces(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", false},
		{"    ", true},
		{"  \r\n  ", true},
		{"\014\012\011\013\015", true},
		{"\014\012\011\013 abc  \015", false},
		{"\f\n\t\v\r\f", true},
		{"x\n\t\t\t\t", false},
		{"\f\n\t  \n\n\n   \v\r\f", true},
	}
	for _, test := range tests {
		actual := KStr.IsWhitespaces(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsWhitespaces(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsWhitespaces(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsWhitespaces("\f\n\t\v\r\f")
	}
}

func TestHasWhitespace(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", false},
		{"    ", true},
		{"  \r\n  ", true},
		{"\014\012\011\013\015", true},
		{"\014\012\011\013 abc  \015", true},
		{"\f\n\t\v\r\f", true},
		{"x\n\t\t\t\t", true},
		{"\f\n\t  \n\n\n   \v\r\f", true},
	}
	for _, test := range tests {
		actual := KStr.HasWhitespace(test.param)
		if actual != test.expected {
			t.Errorf("Expected HasWhitespace(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkHasWhitespace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasWhitespace("x\n\t\t\t\t")
	}
}

func TestIsASCII(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"你好，世界", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
	}
	for _, test := range tests {
		actual := KStr.IsASCII(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsASCII(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsASCII("1234abcDEF")
	}
}

func TestIsMultibyte(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"你好，世界 foobar", true},
		{"test@＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"안녕하세요", true},
	}
	for _, test := range tests {
		actual := KStr.IsMultibyte(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsMultibyte(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsMultibyte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMultibyte("你好，世界 foobar")
	}
}

func TestHasFullWidth(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abc", false},
		{"123", false},
		{"hello world", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"你好，世界 foobar", true},
		{"test@＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"안녕하세요", true},
	}
	for _, test := range tests {
		actual := KStr.HasFullWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected HasFullWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkHasFullWidth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasFullWidth("test@＠example.com")
	}
}

func TestHasHalfWidth(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"abc", true},
		{"123", true},
		{"hello world", true},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", true},
		{"ひらがな・カタカナ、．漢字", false},
		{"你好，世界 foobar", true},
		{"test@＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"안녕하세요", false},
	}
	for _, test := range tests {
		actual := KStr.HasHalfWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected HasHalfWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkHasHalfWidth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasHalfWidth("test@＠example.com")
	}
}

func TestIsBase64(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", true},
		{"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"12345", false},
		{"", false},
		{"Vml2YW11cyBmZXJtZtesting123", false},
	}
	for _, test := range tests {
		actual := KStr.IsBase64(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsBase64(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsBase64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsBase64("Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==")
	}
}

func TestIsBase64Image(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"data:image/png;base6412345", false},
		{"data:image/png;base64,12345", false},
		{"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", false},
		{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
	}
	for _, test := range tests {
		actual := KStr.IsBase64Image(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsBase64Image(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsBase64Image(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsBase64Image("data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=")
	}
}

func TestIsIP(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := KStr.IsIP(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsIP(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIP("127.0.0.1")
	}
}

func TestIsPort(t *testing.T) {
	var tests = []struct {
		param    interface{}
		expected bool
	}{
		{"hello", false},
		{"1", true},
		{0, false},
		{100, true},
		{"65535", true},
		{"0", false},
		{"65536", false},
		{"65538", false},
	}

	for _, test := range tests {
		actual := KStr.IsPort(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPort(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsPort("65538")
	}
}

func TestIsDNSName(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"a.bc", true},
		{"a.b.", true},
		{"a.b..", false},
		{"localhost.local", true},
		{"localhost.localdomain.intern", true},
		{"l.local.intern", true},
		{"ru.link.n.svpncloud.com", true},
		{"-localhost", false},
		{"localhost.-localdomain", false},
		{"localhost.localdomain.-int", false},
		{"_localhost", true},
		{"localhost._localdomain", true},
		{"localhost.localdomain._int", true},
		{"lÖcalhost", false},
		{"localhost.lÖcaldomain", false},
		{"localhost.localdomain.üntern", false},
		{"__", true},
		{"localhost/", false},
		{"127.0.0.1", false},
		{"[::1]", false},
		{"50.50.50.50", false},
		{"localhost.localdomain.intern:65535", false},
		{"漢字汉字", false},
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de", false},
	}

	for _, test := range tests {
		actual := KStr.IsDNSName(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsDNS(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsDNSName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsDNSName("localhost.local")
	}
}

func TestIsDialString(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost.local:1", true},
		{"localhost.localdomain:9090", true},
		{"localhost.localdomain.intern:65535", true},
		{"127.0.0.1:30000", true},
		{"[::1]:80", true},
		{"[1200::AB00:1234::2552:7777:1313]:22", false},
		{"-localhost:1", false},
		{"localhost.-localdomain:9090", false},
		{"localhost.localdomain.-int:65535", false},
		{"localhost.loc:100000", false},
		{"漢字汉字:2", false},
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de:20000", false},
	}

	for _, test := range tests {
		actual := KStr.IsDialString(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsDialString(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsDialString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsDialString("127.0.0.1:30000")
	}
}

func TestIsMACAddr(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"123", false},
		{"abacaba", false},
		{"01:23:45:67:89:ab", true},
		{"01:23:45:67:89:ab:cd:ef", true},
		{"01-23-45-67-89-ab", true},
		{"01-23-45-67-89-ab-cd-ef", true},
		{"0123.4567.89ab", true},
		{"0123.4567.89ab.cdef", true},
		{"3D:F2:C9:A6:B3:4F", true},
		{"08:00:27:88:0f:fd", true},
		{"00:e0:66:07:5c:97:00:00", true},
		{"08:00:27:00:d8:94:00:00", true},
		{"3D-F2-C9-A6-B3:4F", false},
	}
	for _, test := range tests {
		actual := KStr.IsMACAddr(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsMACAddr(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsMACAddr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMACAddr("08:00:27:88:0f:fd")
	}
}

func TestIsHost(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"hello world", false},
		{"localhost", true},
		{"localhost.localdomain", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"::1", true},
		{"play.golang.org", true},
		{"localhost.localdomain.intern:65535", false},
		{"-[::1]", false},
		{"-localhost", false},
		{".localhost", false},
	}
	for _, test := range tests {
		actual := KStr.IsHost(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsHost(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsHost(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsHost("localhost.localdomain")
	}
}

func TestIsRsaPublicKey(t *testing.T) {
	var tests = []struct {
		rsastr   string
		keylen   int
		expected bool
	}{
		{`fubar`, 2048, false},
		{`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuu
XwKYLq0DKUE3t/HHsNdowfD9+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9BmMEcI3uoKbeXCbJRI
HoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzTUmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZ
B7ucimFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUvbQIDAQAB`, 2048, true},
		{`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuu
XwKYLq0DKUE3t/HHsNdowfD9+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9BmMEcI3uoKbeXCbJRI
HoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzTUmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZ
B7ucimFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUvbQIDAQAB`, 1024, false},
		{`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7
x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuuXwKYLq0DKUE3t/HHsNdowfD9
+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9
BmMEcI3uoKbeXCbJRIHoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzT
UmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZB7uc
imFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUv
bQIDAQAB
-----END PUBLIC KEY-----`, 2048, true},
		{`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7
x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuuXwKYLq0DKUE3t/HHsNdowfD9
+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9
BmMEcI3uoKbeXCbJRIHoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzT
UmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZB7uc
imFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUv
bQIDAQAB
-----END PUBLIC KEY-----`, 4096, false},
		{`-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKn4X6phG2ZsKjof
ytRsM8zC7VTZmQSi9hr7ZqHxsIe+UeGToXLSqfJ8ikWWMg15N8PTbzIG11GTexyd
QH/u+zPAS//qrf0HbCXjICt741A8qMipMHIG409PYLQWjfnrjusLt51dY84llj9C
7BzXlHvWqowBGU5jCEaQTBAHPRutAgMBAAECgYAYNdeylihn+2o8Y0Dp5wut0+oo
VuJT5b52c27YDGwfub1CC1xI1bb9Yj3z0YQJpUWLMDe7gXv0E7TKi5+fWXQQXJWt
ejTBtbf0hE14x6OqTzazess99UAxKIdsk7trzVRlPkE4NpJ5jAGTzPqHPlkuaFb3
IK3dyQGLas5QriFnAQJBANagrgmfxygmwH+i7QacffZ6yTu+rhyAcdeUSu6ekPUu
ITv8mOA/bT2m9sIGinW3gjf8KMfz9JH11TasZVsL8e0CQQDKu/bc9oTI0f2jRupY
vmrc31rmOdPq4C4Z6Uj00Ui/FicdywUnGF0bvA+jlCUTLEqBYerl3EEHeLiyZsbT
E5jBAkBVhIZz/T78h5xR/xgUd0xVZo1CCfMUFjXGISdONs4pcyz42ugLChq74wgV
PUf0KZ9wMUAKk/DSK7K96ykjgvntAkBwmqBOMLqmFETN2Mi3S+RtE74YXAxBzAyv
Jaz5FflS8Yn+eVI+WcD1c6o4EEPbd2FWpb1juMeBz+K+bGmIubzBAkB61Sd8LvfF
fDA7MDOGRtIcWq+7bPPw3y44RYIKA35ocMAlzHFhXw7RtSLCl6xgzIpkIfW4ilCP
oCbhuSHBcPnj
-----END PRIVATE KEY-----`, 1024, false},
		{`JXU4RkQ5JXU5MUNDJXU2NjJGJXU4OTgxJXU1MkEwJXU1QkM2JXU3Njg0JXU1MTg1JXU1QkI5JXVGRjAx`, 2048, false},
		// dsa public key
		{`-----BEGIN PUBLIC KEY-----
MIIDRzCCAjkGByqGSM44BAEwggIsAoIBAQCYBeAV/nYFehIyAJqGBSl6Kqthllr5
25iJYG7R9V+/wG5oaVtFJSow/vexBaQ0D5fLQZHJhOPPd+QkEQeMWXVh1mLv0a/V
tbVzA/X5nPrh6qf3SK1fO3cM19Z2YFqCE9sXtrDfroi/DR9Ze1uDT/HVDJ23iZZ7
x7f8cegQN23jOv1APz2d4OEqGe1s85RcS0RPoRrBe1e5itaM1EU0eCCaUjozYt4H
dLZ/VhYZlTG5k814EqrAX+4aWFXUKW1X374a6cvfXirGzZfYr90pL/8VAHATbR2O
P6R0VrdZ0W1hfwPkPb9zBZMaV3+A1HewCjsuheXIKLxnIG+SbceMyYizAiEAkr9Q
R4mvyGhvC79HoQxjRJZRYYqf1O92Yn1dixROC+sCggEAL0rHy4qOIW3g4l/FFh4y
uzzXXePBooCc2jpdYlGXa9g9B5ueX2GQ5+f/QB0VoXvGeYaXefo2YTW5B45IHn7W
9ceX9yme3n9tl8H1dK3sjyqQKxAhyynM1wJaBaALhYT0NzuCXEoBq3kn7On3rU8d
/LM+1UoDwJ0iPqooI9xDW5UX8xd+iYV2FzMtc+SWu4YWmH57EKjcOgC9MqPzCpIn
1Cgo7nSexzSCYIXGDVOqJ0hjeHlL54CMOON2EkUg0e3J/mcneTT8YbP8zPMuBrEX
vwPWNk8wJr2rtxpjhny/sj8BCJY5hhKQFHL1kive7i16AQJv3gJn42eGFJgBsdYa
lgOCAQYAAoIBAQCFyXq2x1BWFxj8qQrbGl5bojxO4r8+gnIoCIbzaxJbiK+eo+JT
BiJNQlludq8f1+0SZ9Paiv1qLaH5p1qxw7mz4ZU8HO4+9grDIb1tuWld/RyhH9PJ
NIoXIVT1J6lK8DqpjnIIoIjqHh5kSJNnXw6XQrA5nlcdZfokVl9oXjH0tGl3McdZ
TQ3WVV0EekGzoIrPw7BkGgb71UBedEt9AqkLSnW6KzQ1A1ILokX8Yq9oWLASea3F
9UxJXpPlCRz3FYgvuR+Q07thgm/z3VQ/+Uq0PFsGFB7Cern0vOKZ+E4673jYK9nq
xVZ+SCC8Wd6nIK4FyZbYaa3Jz7GkqHdMelsl
-----END PUBLIC KEY-----`, 2048, false},
	}
	for i, test := range tests {
		actual := KStr.IsRsaPublicKey(test.rsastr, test.keylen)
		if actual != test.expected {
			t.Errorf("Expected TestIsRsaPublicKey(%d, %d) to be %v, got %v %s", i, test.keylen, test.expected, actual, test.rsastr)
		}
	}
}

func BenchmarkIsRsaPublicKey(b *testing.B) {
	b.ResetTimer()
	str := `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuu
XwKYLq0DKUE3t/HHsNdowfD9+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9BmMEcI3uoKbeXCbJRI
HoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzTUmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZ
B7ucimFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUvbQIDAQAB`
	for i := 0; i < b.N; i++ {
		KStr.IsRsaPublicKey(str, 2048)
	}
}

func TestIsUpper(t *testing.T) {
	str1 := "HELLO"
	str2 := "world"
	str3 := "中文"

	res1 := KStr.IsUpper(str1)
	res2 := KStr.IsUpper(str2)
	res3 := KStr.IsUpper(str3)

	if !res1 || res2 || res3 {
		t.Error("IsUpper fail")
		return
	}
}

func BenchmarkIsUpper(b *testing.B) {
	b.ResetTimer()
	str := "HELLO WORLD"
	for i := 0; i < b.N; i++ {
		KStr.IsUpper(str)
	}
}

func TestIsLower(t *testing.T) {
	str1 := "HELLO"
	str2 := "world"
	str3 := "中文"

	res1 := KStr.IsLower(str1)
	res2 := KStr.IsLower(str2)
	res3 := KStr.IsLower(str3)

	if res1 || !res2 || res3 {
		t.Error("IsLower fail")
		return
	}
}

func BenchmarkIsLower(b *testing.B) {
	b.ResetTimer()
	str := "hello World"
	for i := 0; i < b.N; i++ {
		KStr.IsLower(str)
	}
}

func TestIsBlank(t *testing.T) {
	str1 := " 0"
	str2 := " \t\n\r\v\f　"
	str3 := "a1~"

	res1 := KStr.IsBlank(str1)
	res2 := KStr.IsBlank(str2)
	res3 := KStr.IsBlank(str3)
	if res1 || !res2 || res3 {
		t.Error("IsBlank fail")
		return
	}
}

func BenchmarkIsBlank(b *testing.B) {
	b.ResetTimer()
	str := "hello World"
	for i := 0; i < b.N; i++ {
		KStr.IsBlank(str)
	}
}

func TestIsMd5(t *testing.T) {
	str := "hello world"
	md5 := KStr.Md5(str, 32)

	res1 := KStr.IsMd5(md5)
	res2 := KStr.IsMd5("")
	res3 := KStr.IsMd5(str)
	if !res1 || res2 || res3 {
		t.Error("IsMd5 fail")
		return
	}
}

func BenchmarkIsMd5(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.IsMd5(str)
	}
}

func TestIsSha1(t *testing.T) {
	str := "hello world"
	has := KStr.ShaX(str, 1)

	res1 := KStr.IsSha1(has)
	res2 := KStr.IsSha1("")
	res3 := KStr.IsSha1(str)
	if !res1 || res2 || res3 {
		t.Error("IsSha1 fail")
		return
	}
}

func BenchmarkIsSha1(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.IsSha1(str)
	}
}

func TestIsSha256(t *testing.T) {
	str := "hello world"
	has := KStr.ShaX(str, 256)

	res1 := KStr.IsSha256(has)
	res2 := KStr.IsSha256("")
	res3 := KStr.IsSha256(str)
	if !res1 || res2 || res3 {
		t.Error("IsSha256 fail")
		return
	}
}

func BenchmarkIsSha256(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.IsSha256(str)
	}
}

func TestIsSha512(t *testing.T) {
	str := "hello world"
	has := KStr.ShaX(str, 512)

	res1 := KStr.IsSha512(has)
	res2 := KStr.IsSha512("")
	res3 := KStr.IsSha512(str)
	if !res1 || res2 || res3 {
		t.Error("IsSha512 fail")
		return
	}
}

func BenchmarkIsSha512(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.IsSha512(str)
	}
}

func TestStartsWith(t *testing.T) {
	str := "hello你好世界world"
	chk1 := KStr.StartsWith(str, "hello你好")
	chk2 := KStr.StartsWith(str, "gogo")
	if !chk1 || chk2 {
		t.Error("StartsWith fail")
		return
	}
}

func BenchmarkStartsWith(b *testing.B) {
	b.ResetTimer()
	str := "hello world!welcome to golang,go go go!你好世界"
	for i := 0; i < b.N; i++ {
		KStr.StartsWith(str, "hello你好")
	}
}

func TestEndsWith(t *testing.T) {
	str := "hello你好世界world"
	chk1 := KStr.EndsWith(str, "世界world")
	chk2 := KStr.EndsWith(str, "gogo")
	if !chk1 || chk2 {
		t.Error("EndsWith fail")
		return
	}
}

func BenchmarkEndsWith(b *testing.B) {
	b.ResetTimer()
	str := "hello world!welcome to golang,go go go!你好世界"
	for i := 0; i < b.N; i++ {
		KStr.EndsWith(str, "hello你好")
	}
}
