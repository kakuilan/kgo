package kgo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvert_Struct2Map(t *testing.T) {
	//结构体
	var p1 sPerson
	gofakeit.Struct(&p1)
	mp1, _ := KConv.Struct2Map(p1, "json")
	mp2, _ := KConv.Struct2Map(p1, "")

	var ok bool

	_, ok = mp1["name"]
	assert.True(t, ok)

	_, ok = mp1["none"]
	assert.False(t, ok)

	_, ok = mp2["Age"]
	assert.True(t, ok)

	_, ok = mp2["none"]
	assert.True(t, ok)
}

func BenchmarkConvert_Struct2Map_UseTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "json")
	}
}

func BenchmarkConvert_Struct2Map_NoTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "")
	}
}

func TestConver_Int2Str(t *testing.T) {
	var res string

	res = KConv.Int2Str(0)
	assert.NotEmpty(t, res)

	res = KConv.Int2Str(31.4)
	assert.Empty(t, res)

	res = KConv.Int2Str(PKCS_SEVEN)
	assert.Equal(t, "7", res)
}

func BenchmarkConver_Int2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int2Str(123456789)
	}
}

func TestConver_Float2Str(t *testing.T) {
	var res string

	//小数位为负数
	res = KConv.Float2Str(flPi1, -2)
	assert.Equal(t, 4, len(res))

	res = KConv.Float2Str(flPi2, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi3, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi4, 9)
	assert.Equal(t, 11, len(res))

	res = KConv.Float2Str(true, 9)
	assert.Empty(t, res)
}

func BenchmarkConver_Float2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(flPi2, 3)
	}
}

func TestConver_Bool2Str(t *testing.T) {
	var res string

	res = KConv.Bool2Str(true)
	assert.Equal(t, "true", res)

	res = KConv.Bool2Str(false)
	assert.Equal(t, "false", res)
}

func BenchmarkConver_Bool2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Str(true)
	}
}

func TestConver_Bool2Int(t *testing.T) {
	var res int

	res = KConv.Bool2Int(true)
	assert.Equal(t, 1, res)

	res = KConv.Bool2Int(false)
	assert.Equal(t, 0, res)
}

func BenchmarkConver_Bool2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Int(true)
	}
}

func TestConver_Str2Int(t *testing.T) {
	var res int

	res = KConv.Str2Int("123")
	assert.Equal(t, 123, res)

	res = KConv.Str2Int("TRUE")
	assert.Equal(t, 1, res)

	res = KConv.Str2Int("")
	assert.Equal(t, 0, res)

	res = KConv.Str2Int(strHello)
	assert.Equal(t, 0, res)

	res = KConv.Str2Int("123.456")
	assert.Equal(t, 123, res)

	res = KConv.Str2Int("123.678")
	assert.Equal(t, 123, res)
}

func BenchmarkConver_Str2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("1234567")
	}
}

func TestConver_Str2Int8(t *testing.T) {
	var res int8

	res = KConv.Str2Int8("99")
	assert.Equal(t, int8(99), res)

	res = KConv.Str2Int8(nowNanoStr)
	assert.Equal(t, int8(127), res)
}

func BenchmarkConver_Str2Int8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int8("99")
	}
}

func TestConver_Str2Int16(t *testing.T) {
	var res int16

	res = KConv.Str2Int16("99")
	assert.Equal(t, int16(99), res)

	res = KConv.Str2Int16(nowNanoStr)
	assert.Equal(t, int16(32767), res)
}

func BenchmarkConver_Str2Int16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int16("99")
	}
}

func TestConver_Str2Int32(t *testing.T) {
	var res int32

	res = KConv.Str2Int32("99")
	assert.Equal(t, int32(99), res)

	res = KConv.Str2Int32(nowNanoStr)
	assert.Equal(t, int32(2147483647), res)
}

func BenchmarkConver_Str2Int32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int32("99")
	}
}
