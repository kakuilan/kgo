package kgo

import (
	"github.com/stretchr/testify/assert"
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
