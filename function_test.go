package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
}

func BenchmarkConvert_GetVariatePointerAddr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetVariatePointerAddr(intSpeedLight)
	}
}
