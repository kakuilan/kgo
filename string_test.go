package kgo

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestString_Md5Byte_Md5_IsMd5(t *testing.T) {
	var res1, res2 []byte
	var res3, res4 string
	var chk bool

	res1 = KStr.Md5Byte(bytsHello, 16)
	assert.Equal(t, len(res1), 16)

	res1 = KStr.Md5Byte(bytsHello, 0)
	res2 = KStr.Md5Byte(bytsHello, 32)
	assert.Equal(t, res1, res2)

	res3 = KStr.Md5(strHello, 0)
	res4 = KStr.Md5(strHello, 32)
	assert.Equal(t, res3, res4)

	res2 = KStr.Md5Byte(bytsHello)
	res4 = KStr.Md5(strHello)
	assert.Equal(t, string(res2), res4)

	res3 = KStr.Md5(strHello, 16)
	chk = KStr.IsMd5(res3)
	assert.False(t, chk)

	chk = KStr.IsMd5(res4)
	assert.True(t, chk)
}

func BenchmarkString_Md5Byte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Md5Byte(bytsHello)
	}
}

func BenchmarkString_Md5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Md5(strHello)
	}
}

func BenchmarkString_IsMd5(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMd5(strHello)
	}
}

func TestString_ShaXByte_ShaX_IsSha1_IsSha256_IsSha512(t *testing.T) {
	var res1, res2 []byte
	var res3, res4 string
	var chk bool

	res1 = KStr.ShaXByte(bytsHello, 1)
	res3 = KStr.ShaX(strHello, 1)
	chk = KStr.IsSha1(res3)
	assert.Equal(t, res3, string(res1))
	assert.True(t, chk)

	res2 = KStr.ShaXByte(bytsHello, 256)
	res4 = KStr.ShaX(strHello, 256)
	chk = KStr.IsSha256(res4)
	assert.Equal(t, res4, string(res2))
	assert.True(t, chk)

	res1 = KStr.ShaXByte(bytsHello, 512)
	res3 = KStr.ShaX(strHello, 512)
	chk = KStr.IsSha512(res3)
	assert.Equal(t, res3, string(res1))
	assert.True(t, chk)
}

func TestString_ShaXByte_Panic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()
	KStr.ShaXByte(bytsHello, 32)
}

func TestString_ShaX_Panic(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()
	KStr.ShaX(strHello, 64)
}

func BenchmarkString_ShaXByte1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaXByte(bytsHello, 1)
	}
}

func BenchmarkString_ShaXByte256(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaXByte(bytsHello, 256)
	}
}

func BenchmarkString_ShaXByte512(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaXByte(bytsHello, 512)
	}
}

func BenchmarkString_ShaX1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaX(strHello, 1)
	}
}

func BenchmarkString_ShaX256(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaX(strHello, 256)
	}
}

func BenchmarkString_ShaX512(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ShaX(strHello, 512)
	}
}

func BenchmarkString_IsSha1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsSha1(strSha1)
	}
}

func BenchmarkString_strSha256(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsSha256(strSha256)
	}
}

func BenchmarkString_strSha512(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsSha512(strSha512)
	}
}

func TestString_AddslashesStripslashes(t *testing.T) {
	var res1, res2 string

	res1 = KStr.Addslashes(tesStr5)
	assert.Contains(t, res1, "\\")

	res2 = KStr.Stripslashes(res1)
	assert.Equal(t, res2, tesStr5)
	assert.NotContains(t, res2, "\\")

	res2 = KStr.Stripslashes(tesStr6)
	assert.NotContains(t, res2, '\\')
}

func BenchmarkString_Addslashes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Addslashes(tesStr5)
	}
}

func BenchmarkString_Stripslashes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Stripslashes(tesStr6)
	}
}

func TestString_JsonEncodeJsonDecode(t *testing.T) {
	var res1 []byte
	var res2 []interface{}
	var err error

	//编码
	res1, err = KStr.JsonEncode(personMps)
	assert.Nil(t, err)

	//解码
	err = KStr.JsonDecode(res1, &res2)
	assert.Nil(t, err)
	assert.Equal(t, string(res1), personsArrJson)
}

func BenchmarkString_JsonEncode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.JsonEncode(personMps)
	}
}

func BenchmarkString_JsonDecode(b *testing.B) {
	b.ResetTimer()
	var res []interface{}
	for i := 0; i < b.N; i++ {
		_ = KStr.JsonDecode([]byte(personsArrJson), &res)
	}
}

func TestString_Utf8ToGbkGbkToUtf8_IsUtf8IsGbk(t *testing.T) {
	var res1, res2 []byte
	var chk1, chk2 bool
	var err error

	//utf8 -> gbk
	chk1 = KStr.IsUtf8(bytsUtf8Hello)
	res1, err = KStr.Utf8ToGbk(bytsUtf8Hello)
	assert.True(t, chk1)
	assert.Nil(t, err)

	//gbk -> utf8
	chk2 = KStr.IsGbk(res1)
	res2, err = KStr.GbkToUtf8(res1)
	assert.True(t, chk2)
	assert.Nil(t, err)

	assert.Equal(t, res1, bytsGbkHello)
	assert.Equal(t, res2, bytsUtf8Hello)
}

func BenchmarkString_Utf8ToGbk(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Utf8ToGbk(bytsUtf8Hello)
	}
}

func BenchmarkString_GbkToUtf8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.GbkToUtf8(bytsGbkHello)
	}
}

func BenchmarkString_IsUtf8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUtf8(bytsUtf8Hello)
	}
}

func BenchmarkString_IsGbk(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsGbk(bytsGbkHello)
	}
}

func TestString_Nl2br_Br2nl(t *testing.T) {
	var res1, res2 string

	res1 = KStr.Nl2br(tesStr7)
	assert.Contains(t, res1, "<br />")

	res2 = KStr.Br2nl(res1)
	assert.Equal(t, res2, tesStr7)

	res2 = KStr.Br2nl(tesStr8)
	assert.NotContains(t, res2, "br")
	assert.NotContains(t, res2, "BR")
}

func BenchmarkString_Nl2br(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Nl2br(tesStr7)
	}
}

func BenchmarkString_Br2nl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Br2nl(tesStr8)
	}
}

func TestString_RemoveSpace(t *testing.T) {
	var res string

	//移除所有空格
	res = KStr.RemoveSpace(tesStr9, true)
	assert.NotContains(t, res, " ")

	//移除连续空格
	res = KStr.RemoveSpace(tesStr9, false)
	assert.NotContains(t, res, "  ")
}

func BenchmarkString_RemoveSpace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.RemoveSpace(tesStr9, true)
	}
}

func TestString_StripTags(t *testing.T) {
	var res string

	res = KStr.StripTags(tesStr10)
	assert.NotContains(t, res, "script>")
}

func BenchmarkString_StripTags(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.StripTags(tesStr10)
	}
}

func TestString_Html2Text(t *testing.T) {
	var res string

	res = KStr.Html2Text(tesHtmlDoc)
	assert.NotEmpty(t, res)
	assert.NotContains(t, res, "<")
	assert.NotContains(t, res, ">")
}

func BenchmarkString_Html2Text(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Html2Text(tesHtmlDoc)
	}
}

func TestString_ParseStr(t *testing.T) {
	var res map[string]interface{}
	var err error

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri1, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri2, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri3, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri4, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri5, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri6, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri7, res)
	assert.Nil(t, err)

	//将不合法的参数名转换
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri8, res)
	assert.Nil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri9, res)
	assert.NotNil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri10, res)
	assert.NotNil(t, err)

	//错误的
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri11, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri12, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri13, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri14, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri15, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri16, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri17, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri18, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri19, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri20, res)
	assert.Nil(t, err)

	//key nvalid URL escape "%"
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri21, res)
	assert.NotNil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri22, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri23, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri24, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri25, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri26, res)
	assert.Nil(t, err)

	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri27, res)
	assert.Nil(t, err)

	//key nvalid URL escape "%"
	res = KArr.NewStrMapItf()
	err = KStr.ParseStr(tesUri28, res)
	assert.NotNil(t, err)
}

func BenchmarkString_ParseStr(b *testing.B) {
	b.ResetTimer()
	res := KArr.NewStrMapItf()
	for i := 0; i < b.N; i++ {
		_ = KStr.ParseStr(tesUri1, res)
	}
}

func TestString_ParseUrl(t *testing.T) {
	var res map[string]string
	var err error
	var chk bool

	res, err = KStr.ParseUrl(tesUrl01, -1)
	assert.Nil(t, err)

	res, err = KStr.ParseUrl(strHello, -1)
	assert.Nil(t, err)

	//错误的URL
	res, err = KStr.ParseUrl(tesUrl02, -1)
	assert.NotNil(t, err)
	assert.Empty(t, res)

	res, err = KStr.ParseUrl(tesUrl01, 1)
	_, chk = res["scheme"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 2)
	_, chk = res["host"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 4)
	_, chk = res["port"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 8)
	_, chk = res["user"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 16)
	_, chk = res["pass"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 32)
	_, chk = res["path"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 64)
	_, chk = res["query"]
	assert.True(t, chk)

	res, err = KStr.ParseUrl(tesUrl01, 128)
	_, chk = res["fragment"]
	assert.True(t, chk)
}

func BenchmarkString_ParseUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.ParseUrl(tesUrl01, -1)
	}
}

func TestString_UrlEncodeUrlDecode(t *testing.T) {
	var res1, res2 string
	var err error

	res1 = KStr.UrlEncode(tesStr1)
	res2, err = KStr.UrlDecode(res1)
	assert.Equal(t, res2, tesStr1)
	assert.Nil(t, err)
}

func BenchmarkString_UrlEncode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.UrlEncode(tesStr1)
	}
}

func BenchmarkString_UrlDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.UrlDecode(tesStr2)
	}
}

func TestString_RawUrlEncodeRawUrlDecode(t *testing.T) {
	var res1, res2 string
	var err error

	res1 = KStr.RawUrlEncode(tesStr3)
	res2, err = KStr.RawUrlDecode(res1)
	assert.Equal(t, res2, tesStr3)
	assert.Nil(t, err)
}

func BenchmarkString_RawUrlEncode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.RawUrlEncode(tesStr3)
	}
}

func BenchmarkString_RawUrlDecode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.RawUrlDecode(tesStr4)
	}
}

func TestString_HttpBuildQuery(t *testing.T) {
	var res string
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")

	res = KStr.HttpBuildQuery(params)
	assert.Contains(t, res, "&")
}

func BenchmarkString_HttpBuildQuery(b *testing.B) {
	b.ResetTimer()
	params := url.Values{}
	params.Add("a", "abc")
	params.Add("b", "123")
	params.Add("c", "你好")
	for i := 0; i < b.N; i++ {
		KStr.HttpBuildQuery(params)
	}
}

func TestString_FormatUrl(t *testing.T) {
	var res string

	res = KStr.FormatUrl("")
	assert.Empty(t, res)

	res = KStr.FormatUrl(tesUrl03)
	assert.Contains(t, res, "://")

	res = KStr.FormatUrl(tesUrl04)
	assert.Contains(t, res, "://")

	res = KStr.FormatUrl(tesUrl05)
	assert.NotContains(t, res, '\\')
}

func BenchmarkString_FormatUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.FormatUrl(tesUrl05)
	}
}

func TestString_GetDomain(t *testing.T) {
	var tests = []struct {
		param    string
		isMain   bool
		expected string
	}{
		{"", false, ""},
		{strHello, false, ""},
		{strSpeedLight, false, ""},
		{tesUrl05, false, "login.localhost"},
		{tesUrl06, false, "play.golang.com"},
		{tesUrl07, true, "github.io"},
		{tesUrl08, false, "foobar.中文网"},
		{tesUrl09, false, "foobar.com"},
		{localIp, false, "127.0.0.1"},
	}
	for _, test := range tests {
		actual := KStr.GetDomain(test.param, test.isMain)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_GetDomain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.GetDomain(tesUrl10)
	}
}

func TestString_ClearUrlPrefix(t *testing.T) {
	var tests = []struct {
		url      string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{tesUrl10, "https://", "github.com/kakuilan/kgo"},
		{tesUrl11, "/", "google.com/test?name=hello"},
	}
	for _, test := range tests {
		actual := KStr.ClearUrlPrefix(test.url, test.prefix)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_ClearUrlPrefix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ClearUrlPrefix(tesUrl10)
	}
}

func TestString_ClearUrlSuffix(t *testing.T) {
	var tests = []struct {
		url      string
		prefix   string
		expected string
	}{
		{"", "", ""},
		{tesUrl10, "/kgo", "https://github.com/kakuilan"},
		{tesUrl12, "/", "google.com/test?name=hello"},
	}
	for _, test := range tests {
		actual := KStr.ClearUrlSuffix(test.url, test.prefix)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_ClearUrlSuffix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ClearUrlSuffix(tesUrl12)
	}
}

func TestString_IsEmpty(t *testing.T) {
	var res bool

	res = KStr.IsEmpty("")
	assert.True(t, res)

	res = KStr.IsEmpty("  ")
	assert.True(t, res)

	res = KStr.IsEmpty(strHello)
	assert.False(t, res)
}

func BenchmarkString_IsEmpty(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsEmpty(strHello)
	}
}

func TestString_IsLetters(t *testing.T) {
	var res bool

	res = KStr.IsLetters(tesStr11)
	assert.True(t, res)

	res = KStr.IsLetters(tesStr12)
	assert.False(t, res)

	res = KStr.IsLetters("")
	assert.False(t, res)
}

func BenchmarkString_IsLetters(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsLetters(tesStr11)
	}
}

func TestString_IsUpper(t *testing.T) {
	var res bool

	res = KStr.IsUpper(tesStr13)
	assert.True(t, res)

	res = KStr.IsUpper(strHello)
	assert.False(t, res)

	res = KStr.IsUpper("")
	assert.False(t, res)
}

func BenchmarkString_IsUpper(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUpper(tesStr13)
	}
}

func TestString_IsLower(t *testing.T) {
	var res bool

	res = KStr.IsLower(tesStr14)
	assert.True(t, res)

	res = KStr.IsLower(strHello)
	assert.False(t, res)

	res = KStr.IsLower("")
	assert.False(t, res)
}

func BenchmarkString_IsLower(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsLower(tesStr14)
	}
}

func TestString_HasLetter(t *testing.T) {
	var res bool

	res = KStr.HasLetter(strHello)
	assert.True(t, res)

	res = KStr.HasLetter(strSpeedLight)
	assert.False(t, res)
}

func BenchmarkString_HasLetter(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasLetter(strHello)
	}
}

func TestString_IsASCII(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{tesStr15, false},
		{tesStr16, false},
		{tesStr17, false},
		{utf8Hello, false},
		{tesStr18, true},
		{otcAstronomicalUnit, true},
		{tesEmail1, true},
		{strHelloHex, true},
	}
	for _, test := range tests {
		actual := KStr.IsASCII(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsASCII(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsASCII(tesStr11)
	}
}

func TestString_IsMultibyte(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{tesStr11, false},
		{strSpeedLight, false},
		{strPunctuation1, false},
		{tesEmail1, false},
		{strKor, true},
		{strNoGbk, true},
		{strJap, true},
		{strHello, true},
		{tesStr16, true},
	}
	for _, test := range tests {
		actual := KStr.IsMultibyte(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsMultibyte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMultibyte(strNoGbk)
	}
}
