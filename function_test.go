package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFun_lenArrayOrSlice(t *testing.T) {
	var res int

	res = lenArrayOrSlice(naturalArr, 3)
	assert.Greater(t, res, 0)

	res = lenArrayOrSlice(naturalArr, 1)
	assert.Greater(t, res, 0)

	res = lenArrayOrSlice(naturalArr, 2)
	assert.Equal(t, res, -1)

	res = lenArrayOrSlice(naturalArr, 9)
	assert.Greater(t, res, 0)
}

func TestFun_isNil(t *testing.T) {
	var s []int
	var i interface{}
	var res bool

	res = isNil(s)
	assert.True(t, res)

	res = isNil(&s)
	assert.False(t, res)

	res = isNil(i)
	assert.True(t, res)
}

func TestFun_numeric2Float(t *testing.T) {
	var tests = []struct {
		num      interface{}
		expected error
	}{
		{123, nil},
		{int8(12), nil},
		{int16(12), nil},
		{int32(12), nil},
		{int64(12), nil},
		{uint(123), nil},
		{uint8(13), nil},
		{uint16(13), nil},
		{uint32(13), nil},
		{uint64(13), nil},
		{float32(3.1415), nil},
		{float64(3.1415), nil},
		{"6145", nil},
	}
	for _, test := range tests {
		_, actual := numeric2Float(test.num)
		assert.Equal(t, test.expected, actual)
	}
}

func TestFun_GetVariateType(t *testing.T) {
	var res string

	res = GetVariateType(1)
	assert.Equal(t, "int", res)

	res = GetVariateType(intAstronomicalUnit)
	assert.Equal(t, "int64", res)

	res = GetVariateType(flPi1)
	assert.Equal(t, "float32", res)

	res = GetVariateType(floAvogadro)
	assert.Equal(t, "float64", res)

	res = GetVariateType(strHello)
	assert.Equal(t, "string", res)

	res = GetVariateType(true)
	assert.Equal(t, "bool", res)

	res = GetVariateType(rune('你'))
	assert.Equal(t, "int32", res)

	res = GetVariateType('你')
	assert.Equal(t, "int32", res)

	res = GetVariateType([]byte("你好"))
	assert.Equal(t, "[]uint8", res)
}

func BenchmarkFun_GetVariateType(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetVariateType(intAstronomicalUnit)
	}
}

func TestFun_GetVariatePointerAddr(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected float64
	}{
		{intSpeedLight, 0},
		{strHello, 0},
		{crowd, 0},
	}
	for _, test := range tests {
		actual := GetVariatePointerAddr(test.input)
		assert.Greater(t, actual, int64(test.expected))
	}

	res := GetVariatePointerAddr(&tests)
	assert.Greater(t, res, int64(0))
}

func BenchmarkFun_GetVariatePointerAddr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GetVariatePointerAddr(intSpeedLight)
	}
}

func TestFun_IsPointer(t *testing.T) {
	var chk bool

	//非指针
	chk = IsPointer(itfObj, false)
	assert.False(t, chk)

	//指针
	chk = IsPointer(orgS1, false)
	assert.True(t, chk)

	//非nil指针
	chk = IsPointer(orgS1, true)
	assert.True(t, chk)

	//空指针
	chk = IsPointer(itfObj, true)
	assert.False(t, chk)
}

func BenchmarkFun_IsPointer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPointer(orgS1, true)
	}
}

func TestFun_VerifyFunc_CallFunc(t *testing.T) {
	var res []interface{}
	var err error

	//不存在的对象方法调用
	res, err = CallFunc(strHello, helloEngICase)
	assert.NotNil(t, err)
	assert.Empty(t, res)

	//方法存在,但参数数量错误
	fn := getMethod(&KConv, "BaseConvert")
	res, err = CallFunc(fn)
	assert.NotNil(t, err)
	assert.Empty(t, res)

	//方法存在,参数数量无误,但参数类型错误
	res, err = CallFunc(fn, strHello, false, 1)
	assert.NotNil(t, err)
	assert.Empty(t, res)
}
