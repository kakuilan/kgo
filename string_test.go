package kgo

import (
	"fmt"
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
	str := "hello world,welcome to golang!"
	res1 := KStr.Substr(str, 5, 10)
	res2 := KStr.Substr(str, 0, -5)
	res3 := KStr.Substr(str, 5, -1)
	res4 := KStr.Substr(str, 5, 0)

	if len(res1) != 10 || res2 != str || !strings.Contains(str, res3) || res4 != "" {
		t.Error("Substr fail")
		return
	}
	KStr.Substr(str, 10, 50)
}

func BenchmarkSubstr(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Substr(str, 5, 10)
	}
}

func TestMbSubstr(t *testing.T) {
	str := "hello world,你好世界welcome to golang!"
	res1 := KStr.MbSubstr(str, 6, 10)
	res2 := KStr.MbSubstr(str, 0, -5)
	res3 := KStr.MbSubstr(str, 6, -1)
	res4 := KStr.MbSubstr(str, 6, 0)

	if KStr.MbStrlen(res1) != 10 || res2 != str || !strings.Contains(str, res3) || res4 != "" {
		t.Error("MbSubstr fail")
		return
	}
	KStr.MbSubstr(str, 10, 50)
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

func TestStrrev(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strrev(str)
	res2 := KStr.Strrev(res1)
	if res2 != str {
		t.Error("Strrev fail")
		return
	}
}

func BenchmarkStrrev(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strrev(str)
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
	res := KStr.StrShuffle(str)
	if res == str {
		t.Error("StrShuffle fail")
		return
	}
}

func BenchmarkStrShuffle(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.StrShuffle(str)
	}
}

func TestTrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Trim(str)
	if res[0] != 'h' {
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
	res := KStr.Explode(",", "hello,world,welcome,golang")
	if len(res) != 4 {
		t.Error("Explode fail")
		return
	}
}

func BenchmarkExplode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Explode(",", "hello,world,welcome,golang")
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

func TestCamelNameUnderscoreName(t *testing.T) {
	str := "hello world learn_golang"
	res := KStr.CamelName(str)
	res1 := KStr.CamelName("device_id ")
	res2 := KStr.CamelName("create_time ")
	res3 := KStr.CamelName("location ")

	if strings.Contains(res, "_") || strings.Contains(res1, "_") || strings.Contains(res2, "_") || strings.Contains(res3, " ") {
		t.Error("CamelName fail")
		return
	}

	str = KStr.UnderscoreName(res)
	if !strings.Contains(str, "_") {
		t.Error("UnderscoreName fail")
		return
	}
}

func BenchmarkCamelName(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn_golang"
	for i := 0; i < b.N; i++ {
		KStr.CamelName(str)
	}
}

func BenchmarkUnderscoreName(b *testing.B) {
	b.ResetTimer()
	str := "HelloWorldLearnGolang"
	for i := 0; i < b.N; i++ {
		KStr.UnderscoreName(str)
	}
}

func TestRemoveBefore(t *testing.T) {
	str := "hello world learn golang"
	res1 := KStr.RemoveBefore(str, "world", false)
	res2 := KStr.RemoveBefore(str, "world", true)
	res3 := KStr.RemoveBefore(str, "World", false)
	if !strings.Contains(res1, "world") || strings.Contains(res2, "world") || res3 != str {
		t.Error("RemoveBefore fail")
		return
	}
}

func BenchmarkRemoveBefore(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn golang"
	for i := 0; i < b.N; i++ {
		KStr.RemoveBefore(str, "world", true)
	}
}

func TestRemoveAfter(t *testing.T) {
	str := "hello world learn golang"
	res1 := KStr.RemoveAfter(str, "learn", false)
	res2 := KStr.RemoveAfter(str, "learn", true)
	res3 := KStr.RemoveAfter(str, "Learn", false)
	if !strings.Contains(res1, "learn") || strings.Contains(res2, "learn") || res3 != str {
		t.Error("RemoveAfter fail")
		return
	}
}

func BenchmarkRemoveAfter(b *testing.B) {
	b.ResetTimer()
	str := "hello world learn golang"
	for i := 0; i < b.N; i++ {
		KStr.RemoveAfter(str, "learn", true)
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

func TestLowerUpperCaseFirstWords(t *testing.T) {
	str := "Hello world. I`m use Golang, python, and so on."
	res1 := KStr.LowerCaseFirstWords(str)
	res2 := KStr.UpperCaseFirstWords(str)

	if res1 == "" {
		t.Error("LowerCaseFirstWords fail")
		return
	} else if res2 == "" {
		t.Error("UpperCaseFirstWords fail")
		return
	}
}

func BenchmarkLowerCaseFirstWords(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. I`m use Golang, python, and so on."
	for i := 0; i < b.N; i++ {
		KStr.LowerCaseFirstWords(str)
	}
}

func BenchmarkUpperCaseFirstWords(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. I`m use Golang, python, and so on."
	for i := 0; i < b.N; i++ {
		KStr.UpperCaseFirstWords(str)
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
	println("res2:", res2)

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
