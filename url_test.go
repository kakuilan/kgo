package kgo

import (
	"net/url"
	"strings"
	"testing"
)

func TestParseStr(t *testing.T) {
	str1 := `first=value&arr[]=foo+bar&arr[]=baz`
	str2 := `f1=m&f2=n`
	str3 := `f[a]=m&f[b]=n`
	str4 := `f[a][a]=m&f[a][b]=n`
	str5 := `f[]=m&f[]=n`
	str6 := `f[a][]=m&f[a][]=n`
	str7 := `f[][]=m&f[][]=n`
	str8 := `f=m&f[a]=n`
	str9 := `a .[[b=c`
	arr1 := make(map[string]interface{})
	arr2 := make(map[string]interface{})
	arr3 := make(map[string]interface{})
	arr4 := make(map[string]interface{})
	arr5 := make(map[string]interface{})
	arr6 := make(map[string]interface{})
	arr7 := make(map[string]interface{})
	arr8 := make(map[string]interface{})
	arr9 := make(map[string]interface{})

	err1 := KStr.ParseStr(str1, arr1)
	err2 := KStr.ParseStr(str2, arr2)
	err3 := KStr.ParseStr(str3, arr3)
	err4 := KStr.ParseStr(str4, arr4)
	err5 := KStr.ParseStr(str5, arr5)
	err6 := KStr.ParseStr(str6, arr6)
	err7 := KStr.ParseStr(str7, arr7)
	err8 := KStr.ParseStr(str8, arr8)
	err9 := KStr.ParseStr(str9, arr9)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err9 != nil {
		t.Error("ParseStr fail")
		return
	} else if err8 == nil {
		t.Error("ParseStr fail")
		return
	}

	err9 = KStr.ParseStr("f=n&f[a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}
	err9 = KStr.ParseStr("f=n&f[][a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}

	arr9 = make(map[string]interface{})
	_ = KStr.ParseStr("f[][a]=&f[][b]=", arr9)
	_ = KStr.ParseStr("?first=value&arr[]=foo+bar&arr[]=baz&arr[][a]=aaa", arr9)
	_ = KStr.ParseStr("f[][a]=m&f[][b]=h", arr9)
	_ = KStr.ParseStr("he&a=1", arr9)
	_ = KStr.ParseStr("he& =2", arr9)
	_ = KStr.ParseStr("he& g=2", arr9)
	_ = KStr.ParseStr("he&=3", arr9)
	_ = KStr.ParseStr("he&[=4", arr9)
	_ = KStr.ParseStr("he&]=5", arr9)
	_ = KStr.ParseStr("f[a].=m&f=n&", arr9)
	_ = KStr.ParseStr("f=n&f[a][]b=m&", arr9)
	_ = KStr.ParseStr("f=n&f[a][]=m&", arr9)

	err4 = KStr.ParseStr("f[a][]=1&f[a][]=c&f[a][]=&f[b][]=bb&f[]=3&f[]=4", arr4)
	err5 = KStr.ParseStr("f[a][]=12&f[a][]=1.2&f[a][]=abc", arr5)
	err6 = KStr.ParseStr("f[][b]=&f[][a]=12&f[][a]=1.2&f[][a]=abc", arr6)
	err7 = KStr.ParseStr("f=n&f[a][]=m&", arr7)

	if err4 == nil || err5 == nil || err6 == nil || err7 == nil {
		t.Error("ParseStr fail")
		return
	}

	err8 = KStr.ParseStr("%=%gg&b=4", arr9)  //key nvalid URL escape "%"
	err9 = KStr.ParseStr("he&e=%&b=4", arr9) //value nvalid URL escape "%"
	if err8 == nil || err9 == nil {
		t.Error("ParseStr fail")
		return
	}
}

func BenchmarkParseStr(b *testing.B) {
	b.ResetTimer()
	str := `first=value&arr[]=foo+bar&arr[]=baz`
	arr := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		_ = KStr.ParseStr(str, arr)
	}
}

func TestParseUrl(t *testing.T) {
	url := `https://www.google.com/search?source=hp&ei=tDUwXejNGs6DoATYkqCYCA&q=golang&oq=golang&gs_l=psy-ab.3..35i39l2j0i67l8.1729.2695..2888...1.0..0.126.771.2j5......0....1..gws-wiz.....10..0.fFQmXkC_LcQ&ved=0ahUKEwjo9-H7jb7jAhXOAYgKHVgJCIMQ4dUDCAU&uact=5`
	_, err := KStr.ParseUrl(url, -1)
	if err != nil {
		t.Error("ParseUrl fail")
		return
	}
	_, _ = KStr.ParseUrl(url, 1)
	_, _ = KStr.ParseUrl(url, 2)
	_, _ = KStr.ParseUrl(url, 4)
	_, _ = KStr.ParseUrl(url, 8)
	_, _ = KStr.ParseUrl(url, 16)
	_, _ = KStr.ParseUrl(url, 32)
	_, _ = KStr.ParseUrl(url, 64)
	_, _ = KStr.ParseUrl(url, 128)
	_, _ = KStr.ParseUrl("123456789", -1)
	_, err = KStr.ParseUrl("sg>g://asdf43123412341234", -1)
	if err == nil {
		t.Error("ParseUrl fail")
		return
	}
}

func BenchmarkParseUrl(b *testing.B) {
	b.ResetTimer()
	url := `https://www.google.com/search?source=hp&ei=tDUwXejNGs6DoATYkqCYCA&q=golang&oq=golang&gs_l=psy-ab.3..35i39l2j0i67l8.1729.2695..2888...1.0..0.126.771.2j5......0....1..gws-wiz.....10..0.fFQmXkC_LcQ&ved=0ahUKEwjo9-H7jb7jAhXOAYgKHVgJCIMQ4dUDCAU&uact=5`
	for i := 0; i < b.N; i++ {
		_, _ = KStr.ParseUrl(url, -1)
	}
}

func TestUrlEncode(t *testing.T) {
	str := "'test-bla-bla-4>2-y-3<6'"
	res := KStr.UrlEncode(str)
	if !strings.Contains(res, "%") {
		t.Error("UrlEncode fail")
		return
	}
}

func BenchmarkUrlEncode(b *testing.B) {
	b.ResetTimer()
	str := "'test-bla-bla-4>2-y-3<6'"
	for i := 0; i < b.N; i++ {
		KStr.UrlEncode(str)
	}
}

func TestUrlUrlDecode(t *testing.T) {
	str := "one%20%26%20two"
	_, err := KStr.UrlDecode(str)
	if err != nil {
		t.Error("UrlDecode fail")
		return
	}
}

func BenchmarkUrlDecode(b *testing.B) {
	b.ResetTimer()
	str := "one%20%26%20two"
	for i := 0; i < b.N; i++ {
		_, _ = KStr.UrlDecode(str)
	}
}

func TestRawurlEncode(t *testing.T) {
	str := "'foo @+%/'你好"
	res := KStr.RawurlEncode(str)
	if !strings.Contains(res, "%") {
		t.Error("UrlEncode fail")
		return
	}
}

func BenchmarkRawurlEncode(b *testing.B) {
	b.ResetTimer()
	str := "'foo @+%/'你好"
	for i := 0; i < b.N; i++ {
		KStr.RawurlEncode(str)
	}
}

func TestRawurlDecode(t *testing.T) {
	str := "foo%20bar%40baz"
	_, err := KStr.RawurlDecode(str)
	if err != nil {
		t.Error("RawurlDecode fail")
		return
	}
}

func BenchmarkRawurlDecode(b *testing.B) {
	b.ResetTimer()
	str := "foo%20bar%40baz"
	for i := 0; i < b.N; i++ {
		_, _ = KStr.RawurlDecode(str)
	}
}

func TestHttpBuildQuery(t *testing.T) {
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")

	res := KStr.HttpBuildQuery(params)
	if !strings.Contains(res, "&") {
		t.Error("HttpBuildQuery fail")
		return
	}
}

func BenchmarkHttpBuildQuery(b *testing.B) {
	b.ResetTimer()
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")
	for i := 0; i < b.N; i++ {
		KStr.HttpBuildQuery(params)
	}
}

func TestFormatUrl(t *testing.T) {
	res1 := KStr.FormatUrl("")
	res2 := KStr.FormatUrl("abc.com")
	res3 := KStr.FormatUrl("abc.com/hello?a=1")
	res4 := KStr.FormatUrl(`http://login.localhost:3000\/ab//cd/ef///hi\\12/33\`)

	if res1 != "" || res2 == "" || res3 == "" || strings.ContainsRune(res4, '\\') {
		t.Error("FormatUrl fail")
	}
}

func BenchmarkFormatUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.FormatUrl(`http://login.localhost:3000\/ab//cd/ef///hi\\12/33\`)
	}
}

func TestGetDomain(t *testing.T) {
	var tests = []struct {
		param    string
		isMain   bool
		expected string
	}{
		{"", false, ""},
		{"hello world", false, ""},
		{"http://login.localhost:3000", false, "login.localhost"},
		{"https://play.golang.com:3000/p/3R1TPyk8qck", false, "play.golang.com"},
		{"https://www.siongui.github.io/pali-chanting/zh/archives.html", true, "github.io"},
		{"http://foobar.中文网/", false, "foobar.中文网"},
		{"foobar.com/abc/efg/h=1", false, "foobar.com"},
		{"127.0.0.1", false, "127.0.0.1"},
	}
	for _, test := range tests {
		actual := KStr.GetDomain(test.param, test.isMain)
		if actual != test.expected {
			t.Errorf("Expected GetDomain(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	KStr.GetDomain("123456")
}

func BenchmarkGetDomain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.GetDomain("https://github.com/abc")
	}
}

func TestClearUrlPrefix(t *testing.T) {
	var tests = []struct {
		url      string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{"https://github.com/abc", "https://", "github.com/abc"},
		{"////google.com/test?name=hello", "/", "google.com/test?name=hello"},
	}
	for _, test := range tests {
		actual := KStr.ClearUrlPrefix(test.url, test.prefix)
		if actual != test.expected {
			t.Errorf("Expected ClearUrlPrefix(%q, %q) to be %v, got %v", test.url, test.prefix, test.expected, actual)
		}
	}

	KStr.ClearUrlPrefix("")
}

func BenchmarkClearUrlPrefix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ClearUrlPrefix("//github.com/abc")
	}
}

func TestClearUrlSuffix(t *testing.T) {
	var tests = []struct {
		url      string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{"https://github.com/abc/abc", "/abc", "https://github.com"},
		{"google.com/test?name=hello////", "/", "google.com/test?name=hello"},
	}
	for _, test := range tests {
		actual := KStr.ClearUrlSuffix(test.url, test.prefix)
		if actual != test.expected {
			t.Errorf("Expected ClearUrlSuffix(%q, %q) to be %v, got %v", test.url, test.prefix, test.expected, actual)
		}
	}

	KStr.ClearUrlPrefix("")
}

func BenchmarkClearUrlSuffix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ClearUrlSuffix("github.com/abc///")
	}
}
