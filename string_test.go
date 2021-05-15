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

func TestString_HasFullWidth(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{tesStr11, false},
		{strSpeedLight, false},
		{tesStr5, false},
		{strPunctuation2, false},
		{strJap, true},
		{strKor, true},
		{strHello, true},
		{tesStr15, true},
		{tesStr16, true},
	}
	for _, test := range tests {
		actual := KStr.HasFullWidth(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_HasFullWidth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasFullWidth(strHello)
	}
}

func TestString_HasHalfWidth(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{tesStr11, true},
		{strSpeedLight, true},
		{tesStr5, true},
		{strPunctuation2, true},
		{strJap, false},
		{strKor, false},
		{strHello, true},
		{tesStr15, true},
		{tesStr16, false},
	}
	for _, test := range tests {
		actual := KStr.HasHalfWidth(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_HasHalfWidth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasHalfWidth(strHello)
	}
}

func TestString_IsEnglish(t *testing.T) {
	var tests = []struct {
		str      string
		cas      LkkCaseSwitch
		expected bool
	}{
		{"", CASE_NONE, false},
		{strPi6, CASE_NONE, false},
		{strHello, CASE_NONE, false},
		{b64Hello, CASE_NONE, false},
		{helloEngICase, CASE_NONE, true},
		{helloEngICase, 9, true},
		{helloEngICase, CASE_LOWER, false},
		{helloEngICase, CASE_UPPER, false},
		{helloEngLower, CASE_LOWER, true},
		{helloEngUpper, CASE_UPPER, true},
	}
	for _, test := range tests {
		actual := KStr.IsEnglish(test.str, test.cas)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsEnglish(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsEnglish(helloEngICase, CASE_NONE)
	}
}

func TestString_HasEnglish(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strPi6, false},
		{utf8Hello, false},
		{strHello, true},
		{helloEngICase, true},
	}
	for _, test := range tests {
		actual := KStr.HasEnglish(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func TestString_HasChinese(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strPi6, false},
		{helloEngICase, false},
		{strKor, false},
		{utf8Hello, true},
		{strHello, true},
	}
	for _, test := range tests {
		actual := KStr.HasChinese(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_HasChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasChinese(strHello)
	}
}

func TestString_IsChinese(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strPi6, false},
		{helloEngICase, false},
		{strKor, false},
		{utf8Hello, false},
		{helloCn, true},
	}
	for _, test := range tests {
		actual := KStr.IsChinese(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsChinese(helloCn)
	}
}

func TestString_IsChineseName(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strPi6, false},
		{strKor, false},
		{helloEngICase, false},
		{utf8Hello, false},
		{helloCn, true},
		{tesChineseName1, true},
		{tesChineseName2, false},
		{tesChineseName3, true},
		{tesChineseName4, true},
		{tesChineseName5, true},
		{tesChineseName6, true},
		{tesChineseName7, true},
	}
	for _, test := range tests {
		actual := KStr.IsChineseName(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsChineseName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsChineseName(tesChineseName3)
	}
}

func TestString_IsWord(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"", false},
		{tesStr19, false},
		{tesStr20, false},
		{tesStr21, false},
		{tesStr12, false},
		{helloCn, true},
		{tesStr13, true},
		{tesStr22, true},
		{tesStr23, true},
		{tesStr24, true},
		{tesStr25, false},
	}
	for _, test := range tests {
		actual := KStr.IsWord(test.str)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsWord(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsWord(helloCn)
	}
}

func TestString_HasSpecialChar(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"", false},
		{helloCn, false},
		{helloEngICase, false},
		{tesStr15, false},
		{tesStr16, false},
		{strHello, true},
		{tesStr12, true},
		{tesStr19, true},
		{tesStr20, true},
		{tesStr26, true},
		{strPunctuation3, true},
	}
	for _, test := range tests {
		actual := KStr.HasSpecialChar(test.str)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_HasSpecialChar(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasSpecialChar(strPunctuation3)
	}
}

func TestString_IsJSON_Jsonp2Json(t *testing.T) {
	var res string
	var chk bool
	var err error

	chk = KStr.IsJSON("")
	assert.False(t, chk)

	chk = KStr.IsJSON(strHello)
	assert.False(t, chk)

	chk = KStr.IsJSON(strJson5)
	assert.True(t, chk)

	chk = KStr.IsJSON(strJson6)
	assert.True(t, chk)

	res, err = KStr.Jsonp2Json(strJson1)
	chk = KStr.IsJSON(res)
	assert.True(t, chk)
	assert.Nil(t, err)

	res, err = KStr.Jsonp2Json(strJson2)
	chk = KStr.IsJSON(res)
	assert.True(t, chk)
	assert.Nil(t, err)

	//错误格式
	res, err = KStr.Jsonp2Json("")
	chk = KStr.IsJSON(res)
	assert.False(t, chk)
	assert.NotNil(t, err)

	res, err = KStr.Jsonp2Json(strHello)
	chk = KStr.IsJSON(res)
	assert.False(t, chk)
	assert.NotNil(t, err)

	res, err = KStr.Jsonp2Json(strJson3)
	chk = KStr.IsJSON(res)
	assert.False(t, chk)
	assert.NotNil(t, err)
}

func BenchmarkString_IsJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsJSON(strJson6)
	}
}

func BenchmarkString_Jsonp2Json(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.Jsonp2Json(strJson4)
	}
}

func TestString_IsNumeric(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"", false},
		{helloCn, false},
		{helloEngICase, false},
		{strSpeedLight, true},
		{strPi6, true},
	}
	for _, test := range tests {
		actual := KStr.IsNumeric(test.str)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsNumeric(strPi6)
	}
}

func TestString_IsAlphaNumeric(t *testing.T) {
	var tests = []struct {
		str      string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{helloCn, false},
		{strPi6, false},
		{helloEngICase, true},
		{strSpeedLight, true},
		{tesStr27, true},
	}
	for _, test := range tests {
		actual := KStr.IsAlphaNumeric(test.str)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsAlphaNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsAlphaNumeric(tesStr27)
	}
}

func TestString_IsIP(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{localIp, true},
		{noneIp, true},
		{lanIp, true},
		{dockerIp, true},
		{publicIp1, true},
		{publicIp2, true},
		{tesIp1, true},
		{tesIp2, true},
		{tesIp3, true},
		{tesIp4, false},
	}
	for _, test := range tests {
		actual := KStr.IsIP(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIP(lanIp)
	}
}

func TestString_IsIPv4(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{localIp, true},
		{noneIp, true},
		{lanIp, true},
		{baiduIpv4, true},
		{googleIpv4, true},
		{googleIpv6, false},
		{tesIp2, false},
		{tesIp4, false},
		{tesIp5, false},
		{tesIp6, false},
		{tesIp7, false},
	}
	for _, test := range tests {
		actual := KStr.IsIPv4(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsIPv4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIPv4(googleIpv4)
	}
}

func TestString_IsIPv6(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{localIp, false},
		{noneIp, false},
		{lanIp, false},
		{baiduIpv4, false},
		{googleIpv4, false},
		{googleIpv6, true},
		{tesIp2, true},
		{tesIp4, false},
		{tesIp5, false},
		{tesIp6, true},
		{tesIp7, true},
	}
	for _, test := range tests {
		actual := KStr.IsIPv6(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsIPv6(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsIPv6(googleIpv6)
	}
}

func TestString_IsDNSName(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{strHello, false},
		{localIp, false},
		{localHost, true},
		{tesDomain01, false},
		{tesDomain02, false},
		{tesDomain03, true},
		{tesDomain04, true},
		{tesDomain05, false},
		{tesDomain06, true},
		{tesDomain07, true},
		{tesDomain08, false},
		{tesDomain09, true},
		{tesDomain10, true},
		{tesDomain11, false},
		{tesDomain12, true},
		{tesDomain13, false},
		{tesDomain14, true},
		{tesDomain15, false},
		{tesDomain16, true},
		{tesDomain17, false},
		{tesDomain18, false},
		{tesDomain19, true},
		{tesDomain20, false},
		{tesDomain21, false},
		{tesDomain22, true},
	}

	for _, test := range tests {
		actual := KStr.IsDNSName(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsDNSName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsDNSName(tesDomain22)
	}
}

func TestString_IsDialAddr(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{localHost, false},
		{tesDomain23, true},
		{tesDomain24, true},
		{tesDomain25, true},
		{tesDomain26, false},
		{tesDomain27, false},
		{tesDomain28, false},
		{tesDomain29, false},
	}

	for _, test := range tests {
		actual := KStr.IsDialAddr(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsDialAddr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsDialAddr(tesDomain23)
	}
}

func TestString_IsMACAddr(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{helloEngICase, false},
		{tesMac01, false},
		{tesMac02, false},
		{tesMac03, true},
		{tesMac04, true},
		{tesMac05, true},
		{tesMac06, true},
		{tesMac07, true},
		{tesMac08, true},
		{tesMac09, true},
		{tesMac10, true},
		{tesMac11, true},
		{tesMac12, true},
		{tesMac13, true},
		{tesMac14, true},
	}
	for _, test := range tests {
		actual := KStr.IsMACAddr(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsMACAddr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMACAddr(tesMac14)
	}
}

func TestString_IsHost(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{localIp, true},
		{localHost, true},
		{tesDomain06, true},
		{tesIp3, true},
		{tesIp2, true},
		{tesDomain22, true},
		{tesDomain08, false},
		{tesDomain13, false},
		{tesDomain20, false},
		{tesDomain28, false},
	}
	for _, test := range tests {
		actual := KStr.IsHost(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsHost(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsHost(localHost)
	}
}

func TestString_IsEmail(t *testing.T) {
	var res bool
	var err error

	//长度验证
	res, _ = KStr.IsEmail(tesEmail2, false)
	assert.False(t, res)
	res, _ = KStr.IsEmail(tesEmail3, false)
	assert.False(t, res)

	//无效的格式
	res, _ = KStr.IsEmail(tesEmail4, false)
	assert.False(t, res)

	//不验证主机
	res, _ = KStr.IsEmail(tesEmail1, false)
	assert.True(t, res)
	res, _ = KStr.IsEmail(tesEmail7, false)
	assert.True(t, res)

	//有效的账号
	res, err = KStr.IsEmail(tesEmail6, true)
	assert.True(t, res)
	assert.Nil(t, err)
	res, err = KStr.IsEmail(tesEmail8, true)
	assert.True(t, res)
	assert.Nil(t, err)

	//无效的域名
	_, _ = KStr.IsEmail(tesEmail5, true)
}

func BenchmarkString_IsEmail(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.IsEmail(tesEmail1, false)
	}
}

func TestString_Random(t *testing.T) {
	var res string
	var chk bool

	res = KStr.Random(0, RAND_STRING_ALPHA)
	assert.Empty(t, res)

	//字母
	res = KStr.Random(6, RAND_STRING_ALPHA)
	chk = KStr.IsLetters(res)
	assert.True(t, chk)

	res = KStr.Random(6, 90)
	chk = KStr.IsLetters(res)
	assert.True(t, chk)

	//数字
	res = KStr.Random(6, RAND_STRING_NUMERIC)
	chk = KStr.IsNumeric(res)
	assert.True(t, chk)

	//字母数字
	res = KStr.Random(6, RAND_STRING_ALPHANUM)
	chk = KStr.IsAlphaNumeric(res)
	assert.True(t, chk)

	//有特殊字符
	res = KStr.Random(32, RAND_STRING_SPECIAL)
	chk = KStr.IsAlphaNumeric(res)
	if !chk {
		chk = KStr.HasSpecialChar(res)
		assert.True(t, chk)
	}

	//中文
	res = KStr.Random(6, RAND_STRING_CHINESE)
	chk = KStr.IsChinese(res)
	assert.True(t, chk)
}

func BenchmarkString_Random_Alpha(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(6, RAND_STRING_ALPHA)
	}
}

func BenchmarkString_Random_Numeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(6, RAND_STRING_NUMERIC)
	}
}

func BenchmarkString_Random_Alphanum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(6, RAND_STRING_ALPHANUM)
	}
}

func BenchmarkString_Random_Special(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(6, RAND_STRING_SPECIAL)
	}
}

func BenchmarkString_Random_Chinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(6, RAND_STRING_CHINESE)
	}
}

func TestString_IsMobilecn(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesMobilecn1, true},
		{tesMobilecn2, true},
		{tesMobilecn3, true},
		{tesMobilecn4, true},
		{tesMobilecn5, false},
	}
	for _, test := range tests {
		actual := KStr.IsMobilecn(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsMobilecn(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsMobilecn(tesMobilecn1)
	}
}

func TestString_IsTel(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesTel01, false},
		{tesTel02, true},
		{tesTel03, true},
		{tesTel04, true},
		{tesTel05, true},
		{tesTel06, true},
		{tesTel07, true},
		{tesTel08, true},
		{tesTel09, true},
		{tesTel10, true},
		{tesTel11, true},
		{tesTel12, true},
		{tesTel13, false},
		{tesTel14, false},
		{tesTel15, true},
		{tesTel16, true},
	}
	for _, test := range tests {
		actual := KStr.IsTel(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsTel(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsTel(tesTel02)
	}
}

func TestString_IsPhone(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesTel01, false},
		{tesTel02, true},
		{tesMobilecn1, true},
	}
	for _, test := range tests {
		actual := KStr.IsPhone(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsPhone(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsPhone(tesTel02)
	}
}

func TestString_IsCreditNo(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesCredno01, false},
		{tesCredno02, true},
		{tesCredno03, true},
		{tesCredno04, true},
		{tesCredno05, false},
		{tesCredno06, true},
		{tesCredno07, false},
		{tesCredno08, true},
		{tesCredno09, true},
		{tesCredno10, true},
		{tesCredno11, true},
		{tesCredno12, false},
		{tesCredno13, false},
		{tesCredno14, false},
		{tesCredno15, true},
		{tesCredno16, true},
	}
	for _, test := range tests {
		chk, _ := KStr.IsCreditNo(test.param)
		assert.Equal(t, chk, test.expected)
	}
}

func BenchmarkString_IsCreditNo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsCreditNo(tesCredno02)
	}
}

func TestString_IsHexcolor(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesColor01, false},
		{tesColor02, false},
		{tesColor03, false},
		{tesColor04, true},
		{tesColor05, true},
		{tesColor06, true},
		{tesColor07, true},
		{tesColor08, true},
	}
	for _, test := range tests {
		actual, _ := KStr.IsHexcolor(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsHexcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KStr.IsHexcolor(tesColor08)
	}
}

func TestString_IsRGBcolor(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesColor09, true},
		{tesColor10, true},
		{tesColor11, true},
		{tesColor12, false},
		{tesColor13, false},
		{tesColor14, false},
		{tesColor15, false},
	}
	for _, test := range tests {
		actual := KStr.IsRGBcolor(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsRGBcolor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsRGBcolor(tesColor11)
	}
}

func TestString_IsBlank(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{blankChars, true},
		{"0", false},
		{strHello, false},
	}
	for _, test := range tests {
		actual := KStr.IsBlank(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsBlank(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsBlank(blankChars)
	}
}

func TestString_IsWhitespaces(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{strHello, false},
		{"", false},
		{tesStr28, true},
		{tesStr29, true},
		{tesStr30, true},
		{tesStr31, false},
		{tesStr32, true},
		{tesStr33, false},
		{tesStr34, true},
	}
	for _, test := range tests {
		actual := KStr.IsWhitespaces(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsWhitespaces(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsWhitespaces(tesStr30)
	}
}

func TestString_HasWhitespace(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{strHello, true},
		{helloEngICase, false},
		{"", false},
		{tesStr28, true},
		{tesStr29, true},
		{tesStr30, true},
		{tesStr31, true},
		{tesStr32, true},
		{tesStr33, true},
		{tesStr34, true},
	}
	for _, test := range tests {
		actual := KStr.HasWhitespace(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_HasWhitespace(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.HasWhitespace(strHello)
	}
}

func TestString_IsBase64(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesBase64_01, false},
		{tesBase64_02, true},
		{tesBase64_03, true},
		{tesBase64_04, true},
		{tesBase64_05, true},
	}
	for _, test := range tests {
		actual := KStr.IsBase64(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsBase64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsBase64(tesBase64_02)
	}
}

func TestString_IsBase64Image(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesBase64_06, false},
		{tesBase64_07, false},
		{tesBase64_08, false},
		{tesBase64_09, true},
		{tesBase64_10, false},
		{tesBase64_11, true},
		{tesBase64_12, false},
	}
	for _, test := range tests {
		actual := KStr.IsBase64Image(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsBase64Image(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsBase64Image(tesBase64_11)
	}
}

func TestString_IsRsaPublicKey(t *testing.T) {
	var tests = []struct {
		rsastr   string
		keylen   uint16
		expected bool
	}{
		{strHello, 2048, false},
		{tesRsaPubKey01, 2048, true},
		{tesRsaPubKey01, 1024, false},
		{tesRsaPubKey02, 4096, false},
		{tesRsaPubKey03, 1024, false},
		{tesRsaPubKey04, 2048, false},
		{tesRsaPubKey05, 2048, false},
	}
	for _, test := range tests {
		actual := KStr.IsRsaPublicKey(test.rsastr, test.keylen)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsRsaPublicKey(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsRsaPublicKey(tesRsaPubKey01, 2048)
	}
}

func TestString_IsUrl(t *testing.T) {
	//并行测试
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesUrl01, true},
		{tesUrl02, false},
		{tesUrl04, false},
		{tesUrl05, false},
		{tesUrl06, true},
		{tesUrl07, true},
		{tesUrl08, true},
		{tesUrl10, true},
		{tesUrl11, false},
		{tesUrl13, false},
		{tesUrl14, true},
		{tesUrl15, true},
		{tesUrl16, true},
		{tesUrl17, true},
		{tesUrl18, true},
		{tesUrl19, true},
		{tesUrl20, true},
		{tesUrl21, true},
		{tesUrl22, true},
		{tesUrl23, true},
		{tesUrl24, true},
		{tesUrl25, true},
		{tesUrl26, true},
		{tesUrl27, true},
		{tesUrl28, true},
		{tesUrl29, true},
		{tesUrl30, true},
		{tesUrl31, true},
		{tesUrl32, true},
		{tesUrl33, true},
		{tesUrl34, false},
		{tesUrl35, true},
		{tesUrl36, true},
		{tesUrl37, true},
		{tesUrl38, true},
	}
	for _, test := range tests {
		actual := KStr.IsUrl(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsUrl(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUrl(tesUrl01)
	}
}

func TestString_IsUrlExists(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{strHello, false},
		{tesUrl05, false},
		{tesUrl39, true},
	}
	for _, test := range tests {
		actual := KStr.IsUrlExists(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_IsUrlExists(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.IsUrlExists(tesUrl10)
	}
}

func TestString_Strrpos(t *testing.T) {
	var tests = []struct {
		str      string
		needle   string
		offset   int
		expected int
	}{
		{"", "world", 0, -1},
		{helloEng, "world", 0, 6},
		{helloEng, "world", 1, 6},
		{helloEng, "world", -1, 6},
		{helloEng, "World", 0, -1},
	}
	for _, test := range tests {
		actual := KStr.Strrpos(test.str, test.needle, test.offset)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_Strrpos(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Strrpos(helloEng, "world", 0)
	}
}

func TestString_Strripos(t *testing.T) {
	var tests = []struct {
		str      string
		needle   string
		offset   int
		expected int
	}{
		{"", "world", 0, -1},
		{helloEng, "world", 0, 6},
		{helloEng, "world", 1, 6},
		{helloEng, "world", -1, 6},
		{helloEng, "World", 0, 6},
		{helloEng, "haha", 0, -1},
	}
	for _, test := range tests {
		actual := KStr.Strripos(test.str, test.needle, test.offset)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_Strripos(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Strripos(helloEng, "World", 0)
	}
}

func TestString_Ucfirst(t *testing.T) {
	var res string

	res = KStr.Ucfirst("")
	assert.Empty(t, res)

	res = KStr.Ucfirst(helloEng)
	assert.Equal(t, string(res[0]), "H")
}

func BenchmarkString_Ucfirst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Ucfirst(helloEng)
	}
}

func TestString_Lcfirst(t *testing.T) {
	var res string

	res = KStr.Lcfirst("")
	assert.Empty(t, res)

	res = KStr.Lcfirst(helloEngUpper)
	assert.Equal(t, string(res[0]), "h")
}

func BenchmarkString_Lcfirst(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Lcfirst(helloEngUpper)
	}
}

func TestString_Ucwords_Lcwords(t *testing.T) {
	var res1, res2 string

	res1 = KStr.Ucwords(helloOther)
	res2 = KStr.Lcwords(helloOther)

	assert.Equal(t, string(res1[0]), "H")
	assert.Equal(t, string(res2[0]), "h")
}

func BenchmarkString_Ucwords(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Ucwords(helloOther)
	}
}

func BenchmarkString_Lcwords(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Lcwords(helloOther)
	}
}

func TestString_Substr(t *testing.T) {
	var res string

	res = KStr.Substr("", 0)
	assert.Empty(t, res)

	res = KStr.Substr(helloEng, 0)
	assert.Equal(t, res, helloEng)

	var tests = []struct {
		param    string
		start    int
		length   int
		expected string
	}{
		{helloEngICase, 0, 4, "Hell"},
		{helloEngICase, -2, 4, "ld"},
		{helloEngICase, 0, -2, "HelloWor"},
		{helloEngICase, -11, 8, ""},
		{helloEngICase, 5, 16, "World"},
	}
	for _, test := range tests {
		actual := KStr.Substr(test.param, test.start, test.length)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_Substr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Substr(helloEngICase, 5, 10)
	}
}

func TestString_MbSubstr(t *testing.T) {
	var res string

	res = KStr.MbSubstr("", 0)
	assert.Empty(t, res)

	res = KStr.MbSubstr(helloOther, 0)
	assert.Equal(t, res, helloOther)

	var tests = []struct {
		param    string
		start    int
		length   int
		expected string
	}{
		{helloOther, 0, 15, "Hello world. 你好"},
		{helloOther, -3, 4, "on."},
		{helloOther, 0, -37, "Hello world. 你好，"},
		{helloOther, -40, 9, "你好，世界。I`m"},
		{helloOther, 6, 16, "world. 你好，世界。I`m"},
	}
	for _, test := range tests {
		actual := KStr.MbSubstr(test.param, test.start, test.length)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_MbSubstr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.MbSubstr(helloOther, 6, 16)
	}
}

func TestString_SubstrCount(t *testing.T) {
	var res int

	res = KStr.SubstrCount(tesStr9, "world")
	assert.Equal(t, res, 1)

	res = KStr.SubstrCount(tesStr9, "World")
	assert.Equal(t, res, 2)

	res = KStr.SubstrCount(tesStr9, "ello")
	assert.Equal(t, res, 3)
}

func BenchmarkString_SubstrCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.SubstrCount(tesStr9, "ello")
	}
}

func TestString_SubstriCount(t *testing.T) {
	var res int

	res = KStr.SubstriCount(tesStr9, "world")
	assert.Equal(t, res, 3)

	res = KStr.SubstriCount(tesStr9, "World")
	assert.Equal(t, res, 3)

	res = KStr.SubstriCount(tesStr9, "or")
	assert.Equal(t, res, 4)
}

func BenchmarkString_SubstriCount(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.SubstriCount(tesStr9, "or")
	}
}

func TestString_Reverse(t *testing.T) {
	var res string

	res = KStr.Reverse("")
	assert.Empty(t, res)

	res = KStr.Reverse(strHello)
	assert.Equal(t, res, "！好你 !dlroW olleH")
}

func BenchmarkString_Reverse(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Reverse(strHello)
	}
}

func TestString_ChunkSplit(t *testing.T) {
	var res string

	res = KStr.ChunkSplit(helloOther, 4, "")
	assert.Equal(t, res, helloOther)

	res = KStr.ChunkSplit(helloOther, 4, "\r\n")
	assert.Greater(t, len(res), len(helloOther))
}

func BenchmarkString_ChunkSplit(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.ChunkSplit(helloOther, 4, "\r\n")
	}
}

func TestString_Strlen(t *testing.T) {
	var tests = []struct {
		param    string
		expected int
	}{
		{"", 0},
		{strHello, 22},
		{utf8Hello, 18},
		{helloEng, 12},
		{helloOther, 65},
		{strNoGbk, 106},
		{strJap, 39},
		{strKor, 15},
	}
	for _, test := range tests {
		actual := KStr.Strlen(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_Strlen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Strlen(strHello)
	}
}

func TestString_MbStrlen(t *testing.T) {
	var tests = []struct {
		param    string
		expected int
	}{
		{"", 0},
		{strHello, 16},
		{utf8Hello, 6},
		{helloEng, 12},
		{helloOther, 53},
		{strNoGbk, 36},
		{strJap, 13},
		{strKor, 5},
	}
	for _, test := range tests {
		actual := KStr.MbStrlen(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkString_MbStrlen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.MbStrlen(strHello)
	}
}

func TestString_Shuffle(t *testing.T) {
	var res string

	res = KStr.Shuffle("")
	assert.Empty(t, res)

	res = KStr.Shuffle(strHello)
	assert.Equal(t, len(strHello), len(res))
	assert.NotEqual(t, res, strHello)
}

func BenchmarkString_Shuffle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Shuffle(strHello)
	}
}

func TestString_Trim(t *testing.T) {
	var res string

	res = KStr.Trim(tesStr28)
	assert.Empty(t, res)

	res = KStr.Trim(tesStr29)
	assert.Empty(t, res)

	res = KStr.Trim(tesStr30)
	assert.Empty(t, res)

	res = KStr.Trim(tesStr32)
	assert.Empty(t, res)

	res = KStr.Trim(tesStr34)
	assert.Empty(t, res)

	res = KStr.Trim(tesStr31)
	assert.Equal(t, res, "abc")
}

func BenchmarkString_Trim(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Trim(tesStr31)
	}
}

func TestString_Ltrim(t *testing.T) {
	var res string

	res = KStr.Ltrim(tesStr28)
	assert.Empty(t, res)

	res = KStr.Ltrim(tesStr29)
	assert.Empty(t, res)

	res = KStr.Ltrim(tesStr30)
	assert.Empty(t, res)

	res = KStr.Ltrim(tesStr32)
	assert.Empty(t, res)

	res = KStr.Ltrim(tesStr34)
	assert.Empty(t, res)

	res = KStr.Ltrim(tesStr31)
	assert.Equal(t, string(res[0]), "a")
}

func BenchmarkString_Ltrim(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Ltrim(tesStr31)
	}
}
