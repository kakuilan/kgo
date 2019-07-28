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

	err1 := KUrl.ParseStr(str1, arr1)
	err2 := KUrl.ParseStr(str2, arr2)
	err3 := KUrl.ParseStr(str3, arr3)
	err4 := KUrl.ParseStr(str4, arr4)
	err5 := KUrl.ParseStr(str5, arr5)
	err6 := KUrl.ParseStr(str6, arr6)
	err7 := KUrl.ParseStr(str7, arr7)
	err8 := KUrl.ParseStr(str8, arr8)
	err9 := KUrl.ParseStr(str9, arr9)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err9 != nil {
		t.Error("ParseStr fail")
		return
	} else if err8 == nil {
		t.Error("ParseStr fail")
		return
	}

	err9 = KUrl.ParseStr("f=n&f[a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}
	err9 = KUrl.ParseStr("f=n&f[][a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}

	arr9 = make(map[string]interface{})
	_ = KUrl.ParseStr("f[][a]=&f[][b]=", arr9)
	_ = KUrl.ParseStr("?first=value&arr[]=foo+bar&arr[]=baz&arr[][a]=aaa", arr9)
	_ = KUrl.ParseStr("f[][a]=m&f[][b]=h", arr9)
	_ = KUrl.ParseStr("he&a=1", arr9)
	_ = KUrl.ParseStr("he& =2", arr9)
	_ = KUrl.ParseStr("he& g=2", arr9)
	_ = KUrl.ParseStr("he&=3", arr9)
	_ = KUrl.ParseStr("he&[=4", arr9)
	_ = KUrl.ParseStr("he&]=5", arr9)
	_ = KUrl.ParseStr("f[a].=m&f=n&", arr9)
	_ = KUrl.ParseStr("f=n&f[a][]b=m&", arr9)
	_ = KUrl.ParseStr("f=n&f[a][]=m&", arr9)

	err4 = KUrl.ParseStr("f[a][]=1&f[a][]=c&f[a][]=&f[b][]=bb&f[]=3&f[]=4", arr4)
	err5 = KUrl.ParseStr("f[a][]=12&f[a][]=1.2&f[a][]=abc", arr5)
	err6 = KUrl.ParseStr("f[][b]=&f[][a]=12&f[][a]=1.2&f[][a]=abc", arr6)
	err7 = KUrl.ParseStr("f=n&f[a][]=m&", arr7)

	if err4 == nil || err5 == nil || err6 == nil || err7 == nil {
		t.Error("ParseStr fail")
		return
	}

	err8 = KUrl.ParseStr("%=%gg&b=4", arr9)  //key nvalid URL escape "%"
	err9 = KUrl.ParseStr("he&e=%&b=4", arr9) //value nvalid URL escape "%"
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
		_ = KUrl.ParseStr(str, arr)
	}
}

func TestParseUrl(t *testing.T) {
	url := `https://www.google.com/search?source=hp&ei=tDUwXejNGs6DoATYkqCYCA&q=golang&oq=golang&gs_l=psy-ab.3..35i39l2j0i67l8.1729.2695..2888...1.0..0.126.771.2j5......0....1..gws-wiz.....10..0.fFQmXkC_LcQ&ved=0ahUKEwjo9-H7jb7jAhXOAYgKHVgJCIMQ4dUDCAU&uact=5`
	_, err := KUrl.ParseUrl(url, -1)
	if err != nil {
		t.Error("ParseUrl fail")
		return
	}
	_, _ = KUrl.ParseUrl(url, 1)
	_, _ = KUrl.ParseUrl(url, 2)
	_, _ = KUrl.ParseUrl(url, 4)
	_, _ = KUrl.ParseUrl(url, 8)
	_, _ = KUrl.ParseUrl(url, 16)
	_, _ = KUrl.ParseUrl(url, 32)
	_, _ = KUrl.ParseUrl(url, 64)
	_, _ = KUrl.ParseUrl(url, 128)
	_, _ = KUrl.ParseUrl("123456789", -1)
	_, err = KUrl.ParseUrl("sg>g://asdf43123412341234", -1)
	if err == nil {
		t.Error("ParseUrl fail")
		return
	}
}

func BenchmarkParseUrl(b *testing.B) {
	b.ResetTimer()
	url := `https://www.google.com/search?source=hp&ei=tDUwXejNGs6DoATYkqCYCA&q=golang&oq=golang&gs_l=psy-ab.3..35i39l2j0i67l8.1729.2695..2888...1.0..0.126.771.2j5......0....1..gws-wiz.....10..0.fFQmXkC_LcQ&ved=0ahUKEwjo9-H7jb7jAhXOAYgKHVgJCIMQ4dUDCAU&uact=5`
	for i := 0; i < b.N; i++ {
		_, _ = KUrl.ParseUrl(url, -1)
	}
}

func TestUrlEncode(t *testing.T) {
	str := "'test-bla-bla-4>2-y-3<6'"
	res := KUrl.UrlEncode(str)
	if !strings.Contains(res, "%") {
		t.Error("UrlEncode fail")
		return
	}
}

func BenchmarkUrlEncode(b *testing.B) {
	b.ResetTimer()
	str := "'test-bla-bla-4>2-y-3<6'"
	for i := 0; i < b.N; i++ {
		KUrl.UrlEncode(str)
	}
}

func TestUrlUrlDecode(t *testing.T) {
	str := "one%20%26%20two"
	_, err := KUrl.UrlDecode(str)
	if err != nil {
		t.Error("UrlDecode fail")
		return
	}
}

func BenchmarkUrlDecode(b *testing.B) {
	b.ResetTimer()
	str := "one%20%26%20two"
	for i := 0; i < b.N; i++ {
		_, _ = KUrl.UrlDecode(str)
	}
}

func TestRawurlEncode(t *testing.T) {
	str := "'foo @+%/'你好"
	res := KUrl.RawurlEncode(str)
	if !strings.Contains(res, "%") {
		t.Error("UrlEncode fail")
		return
	}
}

func BenchmarkRawurlEncode(b *testing.B) {
	b.ResetTimer()
	str := "'foo @+%/'你好"
	for i := 0; i < b.N; i++ {
		KUrl.RawurlEncode(str)
	}
}

func TestRawurlDecode(t *testing.T) {
	str := "foo%20bar%40baz"
	_, err := KUrl.RawurlDecode(str)
	if err != nil {
		t.Error("RawurlDecode fail")
		return
	}
}

func BenchmarkRawurlDecode(b *testing.B) {
	b.ResetTimer()
	str := "foo%20bar%40baz"
	for i := 0; i < b.N; i++ {
		_, _ = KUrl.RawurlDecode(str)
	}
}

func TestHttpBuildQuery(t *testing.T) {
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")

	res := KUrl.HttpBuildQuery(params)
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
		KUrl.HttpBuildQuery(params)
	}
}
