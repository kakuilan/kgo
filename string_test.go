package kgo

import (
	"fmt"
	"io/ioutil"
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

func TestBr2nl(t *testing.T) {
	html := `
hello world<br>
hello world<br/>
你好，世界<br />
hello world<BR>
hello world<BR/>
你好，世界<BR />
the end.
`
	res := KStr.Br2nl(html)
	if strings.Contains(res, "br") || strings.Contains(res, "BR") {
		t.Error("Br2nl fail")
		return
	}
}

func BenchmarkBr2nl(b *testing.B) {
	b.ResetTimer()
	html := `
hello world<br>
hello world<br/>
你好，世界<br />
hello world<BR>
hello world<BR/>
你好，世界<BR />
the end.
`
	for i := 0; i < b.N; i++ {
		KStr.Br2nl(html)
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

func BenchmarkStringMd5Str16(b *testing.B) {
	b.ResetTimer()
	str := []byte("hello world!")
	for i := 0; i < b.N; i++ {
		md5Str(str, 16)
	}
}

func BenchmarkStringMd5Str32(b *testing.B) {
	b.ResetTimer()
	str := []byte("hello world!")
	for i := 0; i < b.N; i++ {
		md5Str(str, 32)
	}
}

func TestStringShaX(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	str := "apple"

	res1 := KStr.ShaX(str, 1)
	if res1 != "d0be2dc421be4fcd0172e5afceea3970e2f3d940" {
		t.Error("String ShaX[1] fail")
		return
	}

	res2 := KStr.ShaX(str, 256)
	if res2 != "3a7bd3e2360a3d29eea436fcfb7e44c735d117c42d1c1835420b6b9942dd4f1b" {
		t.Error("String ShaX[256] fail")
		return
	}

	res3 := KStr.ShaX(str, 512)
	if res3 != "844d8779103b94c18f4aa4cc0c3b4474058580a991fba85d3ca698a0bc9e52c5940feb7a65a3a290e17e6b23ee943ecc4f73e7490327245b4fe5d5efb590feb2" {
		t.Error("String ShaX[512] fail")
		return
	}
	KStr.ShaX(str, 16)
}

func BenchmarkStringShaX(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. (can you hear me?)"
	for i := 0; i < b.N; i++ {
		KStr.ShaX(str, 256)
	}
}

func TestRandomAlpha(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_ALPHA)
	if !KStr.IsLetters(res) {
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
	if !KConv.IsNumeric(str) {
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

func TestStringIndex(t *testing.T) {
	var tests = []struct {
		str        string
		sub        string
		ignoreCase bool
		expected   int
	}{
		{"", "", false, -1},
		{"Hello 你好, World 世界！", "hello", false, -1},
		{"Hello 你好, World 世界！", "Hello", false, 0},
		{"Hello 你好, World 世界！", "hello", true, 0},
		{"Hello 你好, World 世界！", "world 世", true, 14},
	}
	for _, test := range tests {
		actual := KStr.Index(test.str, test.sub, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.Index(%q, %q, %t) , got %v", test.str, test.sub, test.ignoreCase, actual)
		}
	}
}

func BenchmarkStringIndex(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Index(str, "World", true)
	}
}

func TestStringLastIndex(t *testing.T) {
	var tests = []struct {
		str        string
		sub        string
		ignoreCase bool
		expected   int
	}{
		{"", "", false, -1},
		{"Hello 你好, World 世界！", "world", false, -1},
		{"Hello 你好, World 世界！", "World", false, 14},
		{"Hello 你好, World 世界！", "world", true, 14},
		{"Hello 你好, World 世界！", "world 世", true, 14},
	}
	for _, test := range tests {
		actual := KStr.LastIndex(test.str, test.sub, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.LastIndex(%q, %q, %t) , got %v", test.str, test.sub, test.ignoreCase, actual)
		}
	}
}

func BenchmarkStringLastIndex(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.LastIndex(str, "World", true)
	}
}

func TestStrpos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strpos(str, "world", 0)
	res2 := KStr.Strpos(str, "World", 0)
	if res1 < 0 || res2 > 0 {
		t.Error("Strpos fail")
		return
	}
	KStr.Strpos("", "world", 0)
	KStr.Strpos(str, "world", -1)
}

func BenchmarkStrpos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strpos(str, "world", 0)
	}
}

func TestStripos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Stripos(str, "world", 0)
	res2 := KStr.Stripos(str, "World", 0)
	if res1 < 0 || res2 < 0 {
		t.Error("Stripos fail")
		return
	}
	KStr.Stripos("", "world", 0)
	KStr.Stripos(str, "world", -1)
	KStr.Stripos(str, "haha", 0)
}

func BenchmarkStripos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Stripos(str, "World", 0)
	}
}

func TestStrrpos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strrpos(str, "world", 1)
	res2 := KStr.Strrpos(str, "World", 0)
	if res1 < 0 || res2 > 0 {
		t.Error("Strrpos fail")
		return
	}
	KStr.Strrpos("", "world", 0)
	KStr.Strrpos(str, "world", -1)
}

func BenchmarkStrrpos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strrpos(str, "world", 0)
	}
}

func TestStrripos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strripos(str, "world", 1)
	res2 := KStr.Strripos(str, "World", 2)
	if res1 < 0 || res2 < 0 {
		t.Error("Strripos fail")
		return
	}
	KStr.Strripos("", "world", 0)
	KStr.Strripos(str, "world", -1)
	KStr.Strripos(str, "haha", 0)
}

func BenchmarkStrripos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strripos(str, "World", 0)
	}
}

func TestUcfirst(t *testing.T) {
	str := "hello world!"
	res := KStr.Ucfirst(str)
	if res[0] != 'H' {
		t.Error("Ucfirst fail")
		return
	}
	KStr.Ucfirst("")
}

func BenchmarkUcfirst(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Ucfirst(str)
	}
}

func TestLcfirst(t *testing.T) {
	str := "HELLOW WORLD!"
	res := KStr.Lcfirst(str)
	if res[0] != 'h' {
		t.Error("Lcfirst fail")
		return
	}
	KStr.Lcfirst("")
}

func BenchmarkLcfirst(b *testing.B) {
	b.ResetTimer()
	str := "HELLOW WORLD!"
	for i := 0; i < b.N; i++ {
		KStr.Lcfirst(str)
	}
}

func TestSubstr(t *testing.T) {
	KStr.Substr("", 0)
	KStr.Substr("abcdef", 0)

	var tests = []struct {
		param    string
		start    int
		length   int
		expected string
	}{
		{"abcdef01", 0, 4, "abcd"},
		{"abcdef02", -2, 4, "02"},
		{"abcdef03", 0, -2, "abcdef"},
		{"abcdef04", -9, 8, ""},
		{"abcdef05", 5, 10, "f05"},
	}

	for _, test := range tests {
		actual := KStr.Substr(test.param, test.start, test.length)
		if actual != test.expected {
			t.Errorf("Expected Substr(%q, %d, %d) to be %v, got %v", test.param, test.start, test.length, test.expected, actual)
			return
		}
	}
}

func BenchmarkSubstr(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Substr(str, 5, 10)
	}
}

func TestMbSubstr(t *testing.T) {
	KStr.MbSubstr("", 0)
	KStr.MbSubstr("abcdef", 0)

	var tests = []struct {
		param    string
		start    int
		length   int
		expected string
	}{
		{"ab你好世界cdef01", 0, 4, "ab你好"},
		{"ab你好世界cdef02", -2, 4, "02"},
		{"ab你好世界cdef03", 0, -2, "ab你好世界cdef"},
		{"ab你好世界cdef04", -20, 8, ""},
		{"ab你好世界cdef05", 5, 50, "界cdef05"},
	}

	for _, test := range tests {
		actual := KStr.MbSubstr(test.param, test.start, test.length)
		if actual != test.expected {
			t.Errorf("Expected MbSubstr(%q, %d, %d) to be %v, got %v", test.param, test.start, test.length, test.expected, actual)
			return
		}
	}
}

func BenchmarkMbSubstr(b *testing.B) {
	b.ResetTimer()
	str := "hello world你好世界!"
	for i := 0; i < b.N; i++ {
		KStr.MbSubstr(str, 6, 10)
	}
}

func TestSubstrCount(t *testing.T) {
	str := "hello world!welcome to golang,go go go!"
	res := KStr.SubstrCount(str, "go")
	if res != 4 {
		t.Error("SubstrCount fail")
		return
	}
}

func BenchmarkSubstrCount(b *testing.B) {
	b.ResetTimer()
	str := "hello world!welcome to golang,go go go!"
	for i := 0; i < b.N; i++ {
		KStr.SubstrCount(str, "go")
	}
}

func TestStrReverse(t *testing.T) {
	str := "hello,world"
	res := KStr.Reverse(str)
	if res != "dlrow,olleh" {
		t.Error("String Reverse fail")
		return
	}
}

func BenchmarkStrReverse(b *testing.B) {
	b.ResetTimer()
	str := "hello world,你好，世界.hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Reverse(str)
	}
}

func TestChunkSplit(t *testing.T) {
	str := "Yar?m kilo ?ay, yar?m kilo ?eker"
	res := KStr.ChunkSplit(str, 4, "\r\n")
	if len(res) == 0 {
		t.Error("ChunkSplit fail")
		return
	}
	_ = KStr.ChunkSplit(str, 5, "")
	_ = KStr.ChunkSplit("a", 4, "")
	_ = KStr.ChunkSplit("ab", 64, "")
	_ = KStr.ChunkSplit("abc", 1, "")
}

func BenchmarkChunkSplit(b *testing.B) {
	b.ResetTimer()
	str := "Yar?m kilo ?ay, yar?m kilo ?eker"
	for i := 0; i < b.N; i++ {
		KStr.ChunkSplit(str, 4, "")
	}
}

func TestStrlen(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.Strlen(str)
	if res != 28 {
		t.Error("Strlen fail")
		return
	}
}

func BenchmarkStrlen(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.Strlen(str)
	}
}

func TestMbStrlen(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.MbStrlen(str)
	if res != 18 {
		t.Error("MbStrlen fail")
		return
	}
}

func BenchmarkMbStrlen(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.MbStrlen(str)
	}
}

func TestMbStrShuffle(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.Shuffle(str)
	if res == str {
		t.Error("StrShuffle fail")
		return
	}

	KStr.Shuffle("")
}

func BenchmarkStrShuffle(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.Shuffle(str)
	}
}

func TestTrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Trim(str)
	if res[0] != 'h' {
		t.Error("Trim fail")
		return
	}

	res = KStr.Trim("\v\t 0.0.0\f\n ")
	if res != "0.0.0" {
		t.Error("Trim fail")
		return
	}

	KStr.Trim(str, "\n")
}

func BenchmarkTrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Trim(str)
	}
}

func TestLtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Ltrim(str)
	if res[0] != 'h' {
		t.Error("Ltrim fail")
		return
	}
	KStr.Ltrim(str, "\n")
}

func BenchmarkLtrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Ltrim(str)
	}
}

func TestRtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Rtrim(str, "　")
	if strings.HasSuffix(res, "　") {
		t.Error("Rtrim fail")
		return
	}
	KStr.Rtrim(str)
}

func BenchmarkRtrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Rtrim(str)
	}
}

func TestChr(t *testing.T) {
	res := KStr.Chr(65)
	if res != "A" {
		t.Error("Chr fail")
		return
	}
}

func BenchmarkChr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Chr(int(i))
	}
}

func TestOrd(t *testing.T) {
	res := KStr.Ord("b")
	if res != 98 {
		t.Error("Ord fail")
		return
	}
}

func BenchmarkOrd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Ord("c")
	}
}

func TestJsonEncodeDecode(t *testing.T) {
	obj := make(map[string]interface{})
	obj["k1"] = "abc"
	obj["k2"] = 123
	obj["k3"] = false
	jstr, err := KStr.JsonEncode(obj)
	if err != nil {
		t.Error("JsonEncode fail")
		return
	}

	mp := make(map[string]interface{})
	err2 := KStr.JsonDecode(jstr, &mp)
	if err2 != nil {
		t.Error("JsonDecode fail")
		return
	}
}

func BenchmarkJsonEncode(b *testing.B) {
	b.ResetTimer()
	obj := make(map[string]interface{})
	obj["k1"] = "abc"
	obj["k2"] = 123
	obj["k3"] = false
	for i := 0; i < b.N; i++ {
		_, _ = KStr.JsonEncode(obj)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	b.ResetTimer()
	str := []byte(`{"k1":"abc","k2":123,"k3":false}`)
	mp := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		_ = KStr.JsonDecode(str, &mp)
	}
}

func TestAddslashesStripslashes(t *testing.T) {
	str := "Is your name O'reilly?"
	res1 := KStr.Addslashes(str)
	if !strings.Contains(res1, "\\") {
		t.Error("Addslashes fail")
		return
	}

	res2 := KStr.Stripslashes(res1)
	if strings.Contains(res2, "\\") {
		t.Error("Stripslashes fail")
		return
	}
	KStr.Stripslashes(`Is \ your \\name O\'reilly?`)
}

func BenchmarkAddslashes(b *testing.B) {
	b.ResetTimer()
	str := "Is your name O'reilly?"
	for i := 0; i < b.N; i++ {
		KStr.Addslashes(str)
	}
}

func BenchmarkStripslashes(b *testing.B) {
	b.ResetTimer()
	str := `Is your name O\'reilly?`
	for i := 0; i < b.N; i++ {
		KStr.Stripslashes(str)
	}
}

func TestQuotemeta(t *testing.T) {
	str := "Hello world. (can you hear me?)"
	res := KStr.Quotemeta(str)
	if !strings.Contains(res, "\\") {
		t.Error("Quotemeta fail")
		return
	}
}

func BenchmarkQuotemeta(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. (can you hear me?)"
	for i := 0; i < b.N; i++ {
		KStr.Quotemeta(str)
	}
}

func TestHtmlentitiesEncodeDecode(t *testing.T) {
	str := "A 'quote' is <b>bold</b>"
	res1 := KStr.Htmlentities(str)
	if !strings.Contains(res1, "&") {
		t.Error("Htmlentities fail")
		return
	}

	res2 := KStr.HtmlentityDecode(res1)
	if res2 != str {
		t.Error("HtmlentityDecode fail")
		return
	}
}

func BenchmarkHtmlentities(b *testing.B) {
	b.ResetTimer()
	str := "A 'quote' is <b>bold</b>"
	for i := 0; i < b.N; i++ {
		KStr.Htmlentities(str)
	}
}

func BenchmarkHtmlentityDecode(b *testing.B) {
	b.ResetTimer()
	str := `A &#39;quote&#39; is &lt;b&gt;bold&lt;/b&gt;`
	for i := 0; i < b.N; i++ {
		KStr.HtmlentityDecode(str)
	}
}

func TestCrc32(t *testing.T) {
	str := "The quick brown fox jumped over the lazy dog"
	res := KStr.Crc32(str)
	if res <= 0 {
		t.Error("Crc32 fail")
		return
	}
}

func BenchmarkCrc32(b *testing.B) {
	b.ResetTimer()
	str := "The quick brown fox jumped over the lazy dog"
	for i := 0; i < b.N; i++ {
		KStr.Crc32(str)
	}
}

func TestSimilarText(t *testing.T) {
	str1 := "The quick brown fox jumped over the lazy dog"
	str2 := "The quick brown fox jumped over the lazy dog"
	var percent float64

	res := KStr.SimilarText(str1, str2, &percent)
	if res <= 0 || percent <= 0 {
		t.Error("Crc32 fail")
		return
	}
	KStr.SimilarText("PHP IS GREAT", "WITH MYSQL", &percent)
	KStr.SimilarText("", "", &percent)
}

func BenchmarkSimilarText(b *testing.B) {
	b.ResetTimer()
	str1 := "The quick brown fox jumped over the lazy dog"
	str2 := "The quick brown fox jumped over the lazy dog"
	var percent float64
	for i := 0; i < b.N; i++ {
		KStr.SimilarText(str1, str2, &percent)
	}
}

func TestExplode(t *testing.T) {
	res := KStr.Explode("hello,world;welcome golang", []string{",", " ", ";"}...)
	if len(res) != 4 {
		t.Error("Explode fail")
		return
	}
	KStr.Explode("")
	KStr.Explode("hello,world,welcome,golang")
	KStr.Explode("hello,world,welcome,golang", "")
	KStr.Explode("hello,world,welcome,golang", ",")
}

func BenchmarkExplode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Explode("hello,world;welcome golang", []string{",", " ", ";"}...)
	}
}

func TestUniqid(t *testing.T) {
	res := KStr.Uniqid("test_")
	if len(res) <= 5 {
		t.Error("Uniqid fail")
		return
	}
}

func BenchmarkUniqid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Uniqid("hello_")
	}
}

func TestVersionCompare(t *testing.T) {
	res1 := KStr.VersionCompare("", "", "=")
	res2 := KStr.VersionCompare("", "1.0", "=")
	res3 := KStr.VersionCompare("0.9", "", "=")

	if !res1 || res2 || res3 {
		t.Error("VersionCompare fail")
		return
	}

	KStr.VersionCompare("#09", "#10", "=")
	KStr.VersionCompare("0.9", "1.0", "=")
	KStr.VersionCompare("11.0", "2.0", "=")
	KStr.VersionCompare("dev11.0", "dev2.0", "=")
	KStr.VersionCompare("11.0", "dev2.0", "=")
	KStr.VersionCompare("a21.0", "2.0", "=")

	KStr.VersionCompare("dev-21.0", "1.0", "=")
	KStr.VersionCompare("dev-21.0", "1.0", "=")
	KStr.VersionCompare("dev-21.0.summer", "1.0", "=")
	KStr.VersionCompare("dev-12.0", "dev-12.0", "=")
	KStr.VersionCompare("beta-11.0", "dev-12.0", "=")

	res4 := KStr.VersionCompare("beta-12.0", "dev-12.0", "<")
	res5 := KStr.VersionCompare("beta-12.0", "dev-12.0", "<=")
	res6 := KStr.VersionCompare("beta-12.0", "dev-12.0", ">")
	res7 := KStr.VersionCompare("beta-12.0", "dev-12.0", ">=")
	res8 := KStr.VersionCompare("beta-12.0", "dev-12.0", "=")
	res9 := KStr.VersionCompare("beta-12.0", "dev-12.0", "!=")

	if res4 || res5 || !res6 || !res7 || res8 || !res9 {
		t.Error("VersionCompare fail")
		return
	}

	KStr.VersionCompare("dev11.-1200", "dev11.-1200", "=")
	KStr.VersionCompare("1.2.3-alpha", "1.2.3alph.123", "=")
	KStr.VersionCompare("1.2.3-alpha", "1.2.3alph.num", "=")
	KStr.VersionCompare("1.2.3alph.123", "1.2.3-alpha", "=")
	KStr.VersionCompare("1.2.3alph.sum", "1.2.3-alpha", "=")
	KStr.VersionCompare("1.2.3alph.sum", "1.2.3-alpha.", "=")
}

func TestVersionComparePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KStr.VersionCompare("1.0", "1.2", "dd")
}

func BenchmarkVersionCompare(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.VersionCompare("2.3.1", "2.1.3.4", ">=")
	}
}

func TestToCamelCase(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"", ""},
		{"some_words", "SomeWords"},
		{"http_server", "HttpServer"},
		{"no_https", "NoHttps"},
		{"_complex__case_", "_Complex_Case_"},
		{"some words", "SomeWords"},
		{"sayHello", "SayHello"},
		{"SayHello", "SayHello"},
		{"SayHelloWorld", "SayHelloWorld"},
		{"DOYouOK", "DoYouOk"},
		{"AReYouOK", "AreYouOk"},
	}

	for _, test := range tests {
		actual := KStr.ToCamelCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToCamelCase(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToCamelCase(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn_golang"
	for i := 0; i < b.N; i++ {
		KStr.ToCamelCase(str)
	}
}

func TestToSnakeCase(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"", ""},
		{"FirstName", "first_name"},
		{"HTTPServer", "http_server"},
		{"NoHTTPS", "no_https"},
		{"GO_PATH", "go_path"},
		{"GO PATH", "go_path"},
		{"GO-PATH", "go_path"},
		{"HTTP2XX", "http_2xx"},
		{"http2xx", "http_2xx"},
		{"HTTP20xOK", "http_20x_ok"},
	}

	for _, test := range tests {
		actual := KStr.ToSnakeCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToSnakeCase(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToSnakeCase(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn_golang go-go"
	for i := 0; i < b.N; i++ {
		KStr.ToSnakeCase(str)
	}
}

func TestToKebabCase(t *testing.T) {
	var tests = []struct {
		param    string
		expected string
	}{
		{"", ""},
		{"�helloWorld", "hello-world"},
		{"A", "a"},
		{"HellOW�orld", "hell-oworld"},
		{"-FirstName", "-first-name"},
		{"FirstName", "first-name"},
		{"HTTPServer", "http-server"},
		{"NoHTTPS", "no-https"},
		{"GO_PATH", "go-path"},
		{"GO PATH", "go-path"},
		{"GO-PATH", "go-path"},
		{"HTTP2XX", "http-2xx"},
		{"http2xx", "http-2xx"},
		{"HTTP20xOK", "http-20x-ok"},
	}

	for _, test := range tests {
		actual := KStr.ToKebabCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToKebabCase(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToKebabCase(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn_golang go-go"
	for i := 0; i < b.N; i++ {
		KStr.ToKebabCase(str)
	}
}

func TestRemoveBefore(t *testing.T) {
	var tests = []struct {
		str        string
		sub        string
		include    bool
		ignoreCase bool
		expected   string
	}{
		{"", "", false, false, ""},
		{"hello world", "", false, false, "hello world"},
		{"Hello 你好, World 世界！", "world", false, false, "Hello 你好, World 世界！"},
		{"Hello 你好, World 世界！", "World", false, false, "World 世界！"},
		{"Hello 你好, World 世界！", "World", true, false, " 世界！"},
		{"Hello 你好, World 世界！", "world", false, true, "World 世界！"},
		{"Hello 你好, World 世界！", "world 世", false, true, "World 世界！"},
		{"Hello 你好, World 世界！", "world 世", true, true, "界！"},
	}
	for _, test := range tests {
		actual := KStr.RemoveBefore(test.str, test.sub, test.include, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.RemoveBefore(%q, %q, %t, %t) , got %v", test.str, test.sub, test.include, test.ignoreCase, actual)
		}
	}
}

func BenchmarkRemoveBefore(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn golang"
	for i := 0; i < b.N; i++ {
		KStr.RemoveBefore(str, "world", true, true)
	}
}

func TestRemoveAfter(t *testing.T) {
	var tests = []struct {
		str        string
		sub        string
		include    bool
		ignoreCase bool
		expected   string
	}{
		{"", "", false, false, ""},
		{"hello world", "", false, false, "hello world"},
		{"Hello 你好, World 世界！", "world", false, false, "Hello 你好, World 世界！"},
		{"Hello 你好, World 世界！", "World", false, false, "Hello 你好, World"},
		{"Hello 你好, World 世界！", "World", true, false, "Hello 你好, "},
		{"Hello 你好, World 世界！", "world", false, true, "Hello 你好, World"},
		{"Hello 你好, World 世界！", "world 世", false, true, "Hello 你好, World 世"},
		{"Hello 你好, World 世界！", "world 世", true, true, "Hello 你好, "},
	}
	for _, test := range tests {
		actual := KStr.RemoveAfter(test.str, test.sub, test.include, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.RemoveAfter(%q, %q, %t, %t) , got %v", test.str, test.sub, test.include, test.ignoreCase, actual)
		}
	}
}

func BenchmarkRemoveAfter(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn golang"
	for i := 0; i < b.N; i++ {
		KStr.RemoveAfter(str, "learn", true, true)
	}
}

func TestDBC2SBC(t *testing.T) {
	str := "hello world!"
	res := KStr.DBC2SBC(str)
	for i := 0; i < len(str); i++ {
		ch := str[i] //此处是数字而非字符
		if strings.Contains(res, string(ch)) {
			t.Error("DBC2SBC fail")
			return
		}
	}
}

func BenchmarkDBC2SBC(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.DBC2SBC(str)
	}
}

func TestSBC2DBC(t *testing.T) {
	str := "１２３４５６７８９ａｂｃ！"
	res := KStr.SBC2DBC(str)
	for i := 0; i < len(str); i++ {
		ch := str[i] //此处是数字而非字符
		if strings.Contains(res, string(ch)) {
			t.Error("SBC2DBC fail")
			return
		}
	}
}

func BenchmarkSBC2DBC(b *testing.B) {
	b.ResetTimer()
	str := "１２３４５６７８９ａｂｃ！"
	for i := 0; i < b.N; i++ {
		KStr.SBC2DBC(str)
	}
}

func TestLevenshtein(t *testing.T) {
	s1 := "frederick"
	s2 := "fredelstick"

	res1 := KStr.Levenshtein(&s1, &s2)
	res2 := KStr.Levenshtein(&s2, &s1)
	res3 := KStr.Levenshtein(&s1, &s1)

	if res1 != res2 || res3 != 0 {
		t.Error("Levenshtein fail")
		return
	}

	s3 := "中国"
	s4 := "中华人民共和国"
	s5 := "中华"
	s6 := ""
	s7 := strings.Repeat(s4, 15)
	res4 := KStr.Levenshtein(&s3, &s4)
	res5 := KStr.Levenshtein(&s4, &s5)
	res6 := KStr.Levenshtein(&s5, &s6)
	res7 := KStr.Levenshtein(&s5, &s7)

	if res4 != res5 || res6 <= 0 || res7 != -1 {
		t.Error("Levenshtein fail")
		return
	}
}

func BenchmarkLevenshtein(b *testing.B) {
	b.ResetTimer()
	s1 := "Asheville"
	s2 := "Arizona"
	for i := 0; i < b.N; i++ {
		KStr.Levenshtein(&s1, &s2)
	}
}

func TestClosestWord(t *testing.T) {
	word := "hello,golang"
	searchs := []string{"hehe,php lang", "Hello,go language", "HeLlo,python!", "haha,java", "I`m going."}
	res, dis := KStr.ClosestWord(word, searchs)
	if res == "" || dis == 0 {
		t.Error("ClosestWord fail")
		return
	}

	searchs = append(searchs, word)
	res2, dis2 := KStr.ClosestWord(word, searchs)
	if res2 != word || dis2 != 0 {
		t.Error("ClosestWord fail")
		return
	}
}

func BenchmarkClosestWord(b *testing.B) {
	b.ResetTimer()
	word := "hello,golang"
	searchs := []string{"hehe,php lang", "Hello,go language", "HeLlo,python!", "haha,java", "I`m going."}
	for i := 0; i < b.N; i++ {
		KStr.ClosestWord(word, searchs)
	}
}

func TestUtf8GbkTrans(t *testing.T) {
	// 测试utf-8和gbk编码互转
	str := "你好，世界！"
	gbk, err0 := KStr.Utf8ToGbk([]byte(str))
	if err0 != nil {
		t.Error("Utf8ToGbk fail")
		return
	}

	// "你好，世界！"的GBK编码
	gbkBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	utf1, err1 := KStr.GbkToUtf8(gbk)
	utf2, err2 := KStr.GbkToUtf8(gbkBytes)
	if err1 != nil || err2 != nil || string(utf1) != str || string(utf2) != str {
		t.Error("GbkToUtf8 fail")
		return
	}
}

func BenchmarkUtf8ToGbk(b *testing.B) {
	b.ResetTimer()
	str := []byte("你好，世界！")
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Utf8ToGbk(str)
	}
}

func BenchmarkGbkToUtf8(b *testing.B) {
	b.ResetTimer()
	gbk, _ := KStr.Utf8ToGbk([]byte("你好，世界！"))
	for i := 0; i < b.N; i++ {
		_, _ = KStr.GbkToUtf8(gbk)
	}
}

func TestUtf8Big5Trans(t *testing.T) {
	// 测试utf-8和big5编码互转
	str := "你好，世界！"
	big, err1 := KStr.Utf8ToBig5([]byte(str))
	if err1 != nil {
		t.Error("Utf8ToBig5 fail")
		return
	}

	utf, err2 := KStr.Big5ToUtf8(big)
	if err2 != nil || string(utf) != str {
		t.Error("Big5ToUtf8 fail")
		return
	}
}

func BenchmarkUtf8ToBig5(b *testing.B) {
	b.ResetTimer()
	str := []byte("你好，世界！")
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Utf8ToBig5(str)
	}
}

func BenchmarkBig5ToUtf8(b *testing.B) {
	b.ResetTimer()
	gbk, _ := KStr.Utf8ToBig5([]byte("你好，世界！"))
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Big5ToUtf8(gbk)
	}
}

func TestFirstLetter(t *testing.T) {
	str1 := "hello world"
	str2 := "你好，世界"
	str3 := "hello，世界"
	str4 := "啊哈，world"

	res1 := KStr.FirstLetter(str1)
	res2 := KStr.FirstLetter(str2)
	res3 := KStr.FirstLetter(str3)
	res4 := KStr.FirstLetter(str4)
	res5 := KStr.FirstLetter("")
	res6 := KStr.FirstLetter("~！@")

	if res1 != "h" || res2 != "N" || res3 != "h" || res4 != "A" || res5 != "" || res6 != "" {
		t.Error("FirstLetter fail")
		return
	}

	//其他
	KStr.FirstLetter("布料")
	KStr.FirstLetter("从来")
	KStr.FirstLetter("到达")
	KStr.FirstLetter("饿了")
	KStr.FirstLetter("发展")
	KStr.FirstLetter("改革")
	KStr.FirstLetter("好啊")
	KStr.FirstLetter("将来")
	KStr.FirstLetter("开心")
	KStr.FirstLetter("里面")
	KStr.FirstLetter("名字")
	KStr.FirstLetter("哪里")
	KStr.FirstLetter("欧洲")
	KStr.FirstLetter("品尝")
	KStr.FirstLetter("前进")
	KStr.FirstLetter("人类")
	KStr.FirstLetter("是的")
	KStr.FirstLetter("天天")
	KStr.FirstLetter("问题")
	KStr.FirstLetter("西安")
	KStr.FirstLetter("用途")
	KStr.FirstLetter("这里")
}

func BenchmarkFirstLetter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.FirstLetter("你好")
	}
}

func TestDstrpos(t *testing.T) {
	var str string
	var arr []string

	chk1, itm1 := KStr.Dstrpos(str, arr, false)
	if chk1 || itm1 != "" {
		t.Error("Dstrpos fail")
		return
	}

	str = "Hello 你好, World 世界！"
	arr = []string{"he", "好", "world"}

	chk2, itm2 := KStr.Dstrpos(str, arr, false)
	if !chk2 || itm2 == "" {
		t.Error("Dstrpos fail")
		return
	}

	chk3, itm3 := KStr.Dstrpos(str, arr, true)
	if !chk3 || itm3 != "好" {
		t.Error("Dstrpos fail")
		return
	}

	arr = []string{"呵呵", "时间", "gogo"}
	chk4, itm4 := KStr.Dstrpos(str, arr, true)
	if chk4 || itm4 != "" {
		t.Error("Dstrpos fail")
		return
	}
}

func BenchmarkDstrpos(b *testing.B) {
	b.ResetTimer()
	str := "Hello 你好, World 世界！"
	arr := []string{"he", "好", "world"}
	for i := 0; i < b.N; i++ {
		KStr.Dstrpos(str, arr, false)
	}
}

func TestUcwordsLcwords(t *testing.T) {
	str := "Hello world. 你好，世界。I`m use Golang, python, and so on."
	res1 := KStr.Lcwords(str)
	res2 := KStr.Ucwords(str)

	if res1 != "hello world. 你好，世界。i`m use golang, python, and so on." {
		t.Error("Lcwords fail")
		return
	}
	if res2 != "Hello World. 你好，世界。I`M Use Golang, Python, And So On." {
		t.Error("Ucwords fail")
		return
	}
}

func BenchmarkLcwords(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. 你好，世界。I`m use Golang, python, and so on."
	for i := 0; i < b.N; i++ {
		KStr.Lcwords(str)
	}
}

func BenchmarkUcwords(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. 你好，世界。I`m use Golang, python, and so on."
	for i := 0; i < b.N; i++ {
		KStr.Ucwords(str)
	}
}

func TestRemoveSpace(t *testing.T) {
	str := "hello World. Hello  \t \n world!   Text   \f\n\t\v\r\fMore \014\012\011\013\015here      \t\n\t Hello,\tWorld\n!\n\t"
	res1 := KStr.RemoveSpace(str, true)
	res2 := KStr.RemoveSpace(str, false)

	if strings.Contains(res1, " ") {
		t.Error("RemoveSpace fail")
	} else if !strings.Contains(res2, " ") {
		t.Error("RemoveSpace fail")
	}
}

func BenchmarkRemoveSpace(b *testing.B) {
	b.ResetTimer()
	str := "hello World. Hello  \t \n world!   Text   \f\n\t\v\r\fMore \014\012\011\013\015here      \t\n\t Hello,\tWorld\n!\n\t"
	for i := 0; i < b.N; i++ {
		KStr.RemoveSpace(str, true)
	}
}

var htmlDoc = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>This is page title</title>
    <link rel="shortcut icon" href="/favicon.ico">
    <link href="/assets/css/frontend.min.css?v=0.0.1" rel="stylesheet">
    <link href="/assets/css/all.css?v=0.0.1" rel="stylesheet">
    <!--[if lt IE 9]>
    <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
    <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <style>
        a{
            color: red;
        }
        span{
            margin: 5px;
        }
    </style>
</head>
<body>
    <div>
        <img src="/assets/img/nf.jpg" alt="this is image" class="fleft">
        <div class="fleft">最新公告</div>
        <div class="fright">
            <a href="logout" class="logoutBtn" style="display: none">退出</a>
            <a href="javascript:;" class="loginPwdBtn">登录</a>
            <a href="javascript:;" class="regisBtn">注册</a>
        </div>
        <h1>This is H1 title.</h1>
        <div>
            <p>
                Hello world!
                <span>TEXT <b>I</b> WANT</span>
            </p>
            <ul>
                <li><a href="foo">Foo</a><li>
                <a href="/bar/baz">BarBaz</a>
            </ul>

            <form name="query" action="http://www.example.net" method="post">
                <input type="text" value="123" />
                <textarea type="text" name="nameiknow">The text I want</textarea>
                <select>
                    <option value="111">111</option>
                    <option value="222">222</option>
                </select>
                <canvas>hello</canvas>
                <div id="button">
                    <input type="submit" value="Submit" />
                    <button>提交按钮</button>
                </div>
            </form>
        </div>
        <div>
            <iframe src="http://google.com"></iframe>
        </div>
    </div>
    <script type="text/javascript">
        var require = {
            config: {
                "modulename": "index",
                "controllername": "index",
                "actionname": "index",
                "jsname": "index",
                "moduleurl": "demo",
                "language": "zh-cn",
                "__PUBLIC__": "/",
                "__ROOT__": "/",
                "__CDN__": ""
            }
        };
        /* <![CDATA[ */
        var post_notif_widget_ajax_obj = {"ajax_url":"http:\/\/site.com\/wp-admin\/admin-ajax.php","nonce":"9b8270e2ef","processing_msg":"Processing..."};
        /* ]]> */
    </script>
    <script src="/assets/js/require.min.js" data-main="/assets/js/require-frontend.min.js?v=0.0.1"></script>
</body>
</html>
`

func TestHtml2Text(t *testing.T) {
	res1 := KStr.Html2Text("")
	res2 := KStr.Html2Text(htmlDoc)

	if res1 != "" || res2 == "" {
		t.Error("Html2Text fail")
	}
}

func BenchmarkHtml2Text(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Html2Text(htmlDoc)
	}
}

func TestHideCard(t *testing.T) {
	res0 := KStr.HideCard("")
	res1 := KStr.HideCard("12345")
	res2 := KStr.HideCard("123456789")
	res3 := KStr.HideCard("123456789012345")
	if res0 == "" || res1 == "" || res2 == "" || res3 == "" {
		t.Error("HideCard fail")
	}
}

func BenchmarkHideCard(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HideCard("123456789012345")
	}
}

func TestHideMobile(t *testing.T) {
	res0 := KStr.HideMobile("")
	res1 := KStr.HideMobile("12345")
	res2 := KStr.HideMobile("13712345678")
	if res0 == "" || res1 == "" || res2 == "" {
		t.Error("HideCard fail")
	}
}

func BenchmarkHideMobile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HideMobile("13712345678")
	}
}

func TestHideTrueName(t *testing.T) {
	var tests = []struct {
		param string
	}{
		{""},
		{"李四"},
		{"张三丰"},
		{"公孙先生"},
		{"helloWorld"},
		{"北京搜狗科技公司"},
		{"北京搜狗科技发展有限公司"},
		{"工商发展银行深圳南山科苑梅龙路支行"},
	}
	for _, test := range tests {
		actual := KStr.HideTrueName(test.param)
		if actual == "" {
			t.Errorf("Expected HideTrueName(%q) , got %v", test.param, actual)
		}
	}

}

func BenchmarkHideTrueName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HideTrueName("公孙先生")
	}
}

func TestCountBase64Byte(t *testing.T) {
	img := "./testdata/diglett.png"
	str, _ := KFile.Img2Base64(img)
	res1 := KStr.CountBase64Byte(str)
	res2 := KStr.CountBase64Byte("hello")

	if res1 == 0 || res2 != 0 {
		t.Error("CountBase64Byte fail")
	}
}

func BenchmarkCountBase64Byte(b *testing.B) {
	b.ResetTimer()
	img := "./testdata/diglett.png"
	str, _ := KFile.Img2Base64(img)
	for i := 0; i < b.N; i++ {
		KStr.CountBase64Byte(str)
	}
}

func TestStrpadding(t *testing.T) {
	str1 := "hello,world"
	str2 := "你好,世界"

	chk1 := KStr.Strpad(str1, "-", 1, PAD_BOTH)
	chk2 := KStr.Strpad(str2, "", 10, PAD_BOTH)
	if chk1 != str1 || chk2 != str2 {
		t.Error("Strpad fail")
	}

	res1 := KStr.StrpadLeft(str1, "-", 30)
	res2 := KStr.StrpadLeft(str2, "。", 30)
	if KStr.MbStrlen(res1) != 30 || KStr.MbStrlen(res2) != 30 {
		t.Error("StrpadLeft fail")
	}

	res3 := KStr.StrpadRight(str1, "-", 30)
	res4 := KStr.StrpadRight(str2, "。", 30)
	if KStr.MbStrlen(res3) != 30 || KStr.MbStrlen(res4) != 30 {
		t.Error("StrpadLeft fail")
	}

	res5 := KStr.StrpadBoth(str1, "-", 30)
	res6 := KStr.StrpadBoth(str2, "。", 30)
	if KStr.MbStrlen(res5) != 30 || KStr.MbStrlen(res6) != 30 {
		t.Error("StrpadLeft fail")
	}

}

func BenchmarkStrpadLeft(b *testing.B) {
	b.ResetTimer()
	str := "hello,世界"
	for i := 0; i < b.N; i++ {
		KStr.StrpadLeft(str, "-。", 30)
	}
}

func BenchmarkStrpadRight(b *testing.B) {
	b.ResetTimer()
	str := "hello,世界"
	for i := 0; i < b.N; i++ {
		KStr.StrpadRight(str, "-。", 30)
	}
}

func BenchmarkStrpadBoth(b *testing.B) {
	b.ResetTimer()
	str := "hello,世界"
	for i := 0; i < b.N; i++ {
		KStr.StrpadBoth(str, "-。", 30)
	}
}

func TestStrImg2Base64(t *testing.T) {
	cont, _ := ioutil.ReadFile("testdata/diglett.png")
	img1 := KStr.Img2Base64(cont)
	img2 := KStr.Img2Base64(cont, "png")

	chk1 := KStr.IsBase64Image(img1)
	chk2 := KStr.IsBase64Image(img2)
	if !chk1 || !chk2 {
		t.Error("IsBase64Image fail")
		return
	}
}

func BenchmarkStrImg2Base64(b *testing.B) {
	b.ResetTimer()
	cont, _ := ioutil.ReadFile("testdata/diglett.png")
	for i := 0; i < b.N; i++ {
		KStr.Img2Base64(cont)
	}
}

func TestJsonp2Json(t *testing.T) {
	str := `JsonpCallbackFn_abc123etc({"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}})`
	res, _ := KStr.Jsonp2Json(str)
	if !KStr.IsJSON(res) {
		t.Error("Jsonp2Json fail")
		return
	}

	str = `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
	res, _ = KStr.Jsonp2Json(str)
	if !KStr.IsJSON(res) {
		t.Error("Jsonp2Json fail")
		return
	}

	str = "hello world"
	_, err := KStr.Jsonp2Json(str)
	if err == nil {
		t.Error("Jsonp2Json fail")
		return
	}

	str = "call)hello world(done"
	_, err = KStr.Jsonp2Json(str)
	if err == nil {
		t.Error("Jsonp2Json fail")
		return
	}
}

func BenchmarkJsonp2Json(b *testing.B) {
	b.ResetTimer()
	str := `JsonpCallbackFn_abc123etc({"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}})`
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Jsonp2Json(str)
	}
}

func TestCountWords(t *testing.T) {
	content, _ := KFile.ReadFile("./testdata/dante.txt")
	word_all, mp := KStr.CountWords(KConv.Bytes2Str(content))
	word_num := len(mp)

	if word_all == 0 || word_num == 0 || word_num > word_all {
		t.Error("CountWords fail")
		return
	}

	word_all, mp = KStr.CountWords("hello world,你好，世界.hello world!")
	word_num = len(mp)
	if word_all != 6 || word_num != 4 {
		t.Error("CountWords fail")
		return
	}
}

func BenchmarkCountWords(b *testing.B) {
	b.ResetTimer()
	str := "hello world,你好，世界.hello world!"
	for i := 0; i < b.N; i++ {
		KStr.CountWords(str)
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
			t.Errorf("Expected IsDNSName(%q) to be %v, got %v", test.param, test.expected, actual)
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
	for _, test := range tests {
		actual := KStr.IsRsaPublicKey(test.rsastr, test.keylen)
		if actual != test.expected {
			t.Errorf("Expected IsRsaPublicKey(%s, %d) to be %v, got %v %s", test.rsastr, test.keylen, test.expected, actual, test.rsastr)
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
	var tests = []struct {
		str        string
		sub        string
		ignoreCase bool
		expected   bool
	}{
		{"", "", false, false},
		{"hello world", "", false, false},
		{"Hello 你好, World 世界！", "hello", false, false},
		{"Hello 你好, World 世界！", "Hello", false, true},
		{"Hello 你好, World 世界！", "hello", true, true},
		{"Hello 你好, World 世界！", "Hello 你好", false, true},
		{"Hello 你好, World 世界！", "hello 你好", true, true},
		{"Hello 你好, World 世界！", "world 世", true, false},
	}
	for _, test := range tests {
		actual := KStr.StartsWith(test.str, test.sub, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.StartsWith(%q, %q, %t) , got %v", test.str, test.sub, test.ignoreCase, actual)
		}
	}
}

func BenchmarkStartsWith(b *testing.B) {
	b.ResetTimer()
	str := "hello world!welcome to golang,go go go!你好世界"
	for i := 0; i < b.N; i++ {
		KStr.StartsWith(str, "hello你好", false)
	}
}

func TestEndsWith(t *testing.T) {
	var tests = []struct {
		str        string
		sub        string
		ignoreCase bool
		expected   bool
	}{
		{"", "", false, false},
		{"hello world", "", false, false},
		{"Hello 你好, World 世界！", "World", false, false},
		{"Hello 你好, World 世界！", "World", true, false},
		{"Hello 你好, World 世界！", "World 世界！", false, true},
		{"Hello 你好, World 世界！", "world 世界！", true, true},
	}
	for _, test := range tests {
		actual := KStr.EndsWith(test.str, test.sub, test.ignoreCase)
		if actual != test.expected {
			t.Errorf("Expected KStr.EndsWith(%q, %q, %t) , got %v", test.str, test.sub, test.ignoreCase, actual)
		}
	}
}

func BenchmarkEndsWith(b *testing.B) {
	b.ResetTimer()
	str := "hello world!welcome to golang,go go go!你好世界"
	for i := 0; i < b.N; i++ {
		KStr.EndsWith(str, "hello你好", false)
	}
}

func TestIsWord(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"", false},
		{"_Football", false},
		{"-Football", false},
		{" 3.124", false},
		{"hello world.你好，世界！", false},
		{"世界", true},
		{"hello", true},
		{"作品T", true},
		{"8point", true},
		{"hello_Kitty2", true},
		{"hello-Kitty2", false},
	}
	for _, test := range tests {
		actual := KStr.IsWord(test.str)
		if actual != test.expected {
			t.Errorf("Expected KStr.IsWord(%q) , got %v", test.str, actual)
		}
	}
}

func BenchmarkIsWord(b *testing.B) {
	b.ResetTimer()
	str := "hello world,你好，世界.hello world!"
	for i := 0; i < b.N; i++ {
		KStr.IsWord(str)
	}
}

func TestRemoveEmoji(t *testing.T) {
	str := `Lorem ipsum 🥊dolor 🤒sit amet, consectetur adipiscing 🍂 elit. 🍁🍃🍂🌰🍁🌿🌾🌼🌻سلام تست شد hell中文
😀😁😂😃😄😅😆😉😊😋😎😍😘😗😙😚☺😇😐😑😶😏😣😥😮😯😪😫😴😌😛😜😝😒😓😔😕😲😷😖😞😟😤😢😭😦😧😨😬😰😱😳😵😡😠
👦👧👨👩👴👵👶👱👮👲👳👷👸💂🎅👰👼💆💇🙍🙎🙅🙆💁🙋🙇🙌🙏👤👥🚶🏃👯💃👫👬👭💏💑👪
💪👈👉☝👆👇✌✋👌👍👎✊👊👋👏👐✍
👣👀👂👃👅👄💋👓👔👕👖👗👘👙👚👛👜👝🎒💼👞👟👠👡👢👑👒🎩🎓💄💅💍🌂
📱📲📶📳📴☎📞📟📠
♻🏧🚮🚰♿🚹🚺🚻🚼🚾⚠🚸⛔🚫🚳🚭🚯🚱🚷🔞💈
🙈🙉🙊🐵🐒🐶🐕🐩🐺🐱😺😸😹😻😼😽🙀😿😾🐈🐯🐅🐆🐴🐎🐮🐂🐃🐄🐷🐖🐗🐽🐏🐑🐐🐪🐫🐘🐭🐁🐀🐹🐰🐇🐻🐨🐼🐾🐔🐓🐣🐤🐥🐦🐧🐸🐊🐢🐍🐲🐉🐳🐋🐬🐟🐠🐡🐙🐚🐌🐛🐜🐝🐞
💐🌸💮🌹🌺🌻🌼🌷🌱🌲🌳🌴🌵🌾🌿🍀🍁🍂🍃
🌍🌎🌏🌐🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜☀🌝🌞⭐🌟🌠☁⛅☔⚡❄🔥💧🌊
🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓🍅🍆🌽🍄🌰🍞🍖🍗🍔🍟🍕🍳🍲🍱🍘🍙🍚🍛🍜🍝🍠🍢🍣🍤🍥🍡🍦🍧🍨🍩🍪🎂🍰🍫🍬🍭🍮🍯🍼☕🍵🍶🍷🍸🍹🍺🍻🍴
🎪🎭🎨🎰🚣🛀🎫🏆⚽⚾🏀🏈🏉🎾🎱🎳⛳🎣🎽🎿🏂🏄🏇🏊🚴🚵🎯🎮🎲🎷🎸🎺🎻🎬
😈👿👹👺💀☠👻👽👾💣
🌋🗻🏠🏡🏢🏣🏤🏥🏦🏨🏩🏪🏫🏬🏭🏯🏰💒🗼🗽⛪⛲🌁🌃🌆🌇🌉🌌🎠🎡🎢🚂🚃🚄🚅🚆🚇🚈🚉🚊🚝🚞🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚚🚛🚜🚲⛽🚨🚥🚦🚧⚓⛵🚤🚢✈💺🚁🚟🚠🚡🚀🎑🗿🛂🛃🛄🛅
💌💎🔪💈🚪🚽🚿🛁⌛⏳⌚⏰🎈🎉🎊🎎🎏🎐🎀🎁📯📻📱📲☎📞📟📠🔋🔌💻💽💾💿📀🎥📺📷📹📼🔍🔎🔬🔭📡💡🔦🏮📔📕📖📗📘📙📚📓📃📜📄📰📑🔖💰💴💵💶💷💸💳✉📧📨📩📤📥📦📫📪📬📭📮✏✒📝📁📂📅📆📇📈📉📊📋📌📍📎📏📐✂🔒🔓🔏🔐🔑🔨🔫🔧🔩🔗💉💊🚬🔮🚩🎌💦💨
♠♥♦♣🀄🎴🔇🔈🔉🔊📢📣💤💢💬💭♨🌀🔔🔕✡✝🔯📛🔰🔱⭕✅☑✔✖❌❎➕➖➗➰➿〽✳✴❇‼⁉❓❔❕❗©®™🎦🔅🔆💯🔠🔡🔢🔣🔤🅰🆎🅱🆑🆒🆓ℹ🆔Ⓜ🆕🆖🅾🆗🅿🆘🆙🆚🈁🈂🈷🈶🈯🉐🈹🈚🈲🉑🈸🈴🈳㊗㊙🈺🈵▪▫◻◼◽◾⬛⬜🔶🔷🔸🔹🔺🔻💠🔲🔳⚪⚫🔴🔵
🐁🐂🐅🐇🐉🐍🐎🐐🐒🐓🐕🐖
♈♉♊♋♌♍♎♏♐♑♒♓⛎
🕛🕧🕐🕜🕑🕝🕒🕞🕓🕟🕔🕠🕕🕡🕖🕢🕗🕣🕘🕤🕙🕥🕚🕦⌛⏳⌚⏰⏱⏲🕰
💘❤💓💔💕💖💗💙💚💛💜💝💞💟❣
💐🌸💮🌹🌺🌻🌼🌷🌱🌿🍀
🌿🍀🍁🍂🍃
🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜🌝
🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓
💴💵💶💷💰💸💳
🚂🚃🚄🚅🚆🚇🚈🚉🚊🚝🚞🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚚🚛🚜🚲⛽🚨🚥🚦🚧⚓⛵🚣🚤🚢✈💺🚁🚟🚠🚡🚀
🏠🏡🏢🏣🏤🏥🏦🏨🏩🏪🏫🏬🏭🏯🏰💒🗼🗽⛪🌆🌇🌉
📱📲☎📞📟📠🔋🔌💻💽💾💿📀🎥📺📷📹📼🔍🔎🔬🔭📡📔📕📖📗📘📙📚📓📃📜📄📰📑🔖💳✉📧📨📩📤📥📦📫📪📬📭📮✏✒📝📁📂📅📆📇📈📉📊📋📌📍📎📏📐✂🔒🔓🔏🔐🔑
⬆↗➡↘⬇↙⬅↖↕↔↩↪⤴⤵🔃🔄🔙🔚🔛🔜🔝`

	res := KStr.RemoveEmoji(str)
	str1 := KStr.Trim(res)
	str2 := `Lorem ipsum dolor sit amet, consectetur adipiscing  elit. سلام تست شد hell中文`
	if str1 != str2 {
		t.Error("RemoveEmoji fail")
		return
	}
}

func BenchmarkRemoveEmoji(b *testing.B) {
	b.ResetTimer()
	str := `Hi!😀👽😀☂❤华み원❤This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️`
	for i := 0; i < b.N; i++ {
		KStr.RemoveEmoji(str)
	}
}

func TestUuidV4(t *testing.T) {
	res, err := KStr.UuidV4()

	if res == "" || res == "00000000-0000-0000-0000-000000000000" || err != nil {
		t.Error("UuidV4 fail")
		return
	}
}

func BenchmarkUuidV4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.UuidV4()
	}
}

func TestGravatar(t *testing.T) {
	res1 := KStr.Gravatar("", 100)
	res2 := KStr.Gravatar("hello@test.com", 150)

	if res1 == "" || res2 == "" {
		t.Error("Gravatar fail")
		return
	}
}

func BenchmarkGravatar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Gravatar("hello@test.com", 150)
	}
}

func TestAtWho(t *testing.T) {
	var tests = []struct {
		name     string
		leng     int
		expected []string
	}{
		{"", 0, []string{}},
		{"@hellowor", 3, []string{"hellowor"}},
		{"@hellowor", 5, []string{"hellowor"}},
		{" @hellowor", 5, []string{"hellowor"}},
		{"Hi, @hellowor", 5, []string{"hellowor"}},
		{"Hi,@hellowor", 5, []string{"hellowor"}},
		{"Hi, @hellowor, @tom", 3, []string{"tom"}},
		{"Hi, @hellowor and @tom and @hellowor again", 3, []string{"hellowor", "tom"}},
		{"@hellowor\nanother line @john", 3, []string{"hellowor", "john"}},
		{"hellowor@gmail.com", 0, []string{}},
		{"hellowor@gmail.com @test", 3, []string{"test"}},
	}
	for _, test := range tests {
		actual := KStr.AtWho(test.name, test.leng)
		if !KArr.IsEqualArray(test.expected, actual) {
			t.Errorf("Expected KStr.EndsWith(%q, %d) to be %v, got %v", test.name, test.leng, test.expected, actual)
		}
	}
}

func BenchmarkAtWho(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.AtWho("Hi, @hellowor", 5)
	}
}
