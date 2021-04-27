package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
