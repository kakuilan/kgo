package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumber_NumberFormat(t *testing.T) {
	var res string

	res = KNum.NumberFormat(floNum1, 3, ".", "")
	assert.Equal(t, "123456789.123", res)

	res = KNum.NumberFormat(floNum1, 6, ".", ",")
	assert.Equal(t, "123,456,789.123457", res)

	res = KNum.NumberFormat(floNum2, 0, ".", "")
	assert.Equal(t, "-123", res)
}

func BenchmarkNumber_NumberFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.NumberFormat(floNum1, 3, ".", "")
	}
}

func TestNumber_Range(t *testing.T) {
	var res []int
	var start, end int

	//升序
	start, end = 1, 5
	res = KNum.Range(start, end)
	assert.Equal(t, 5, len(res))
	assert.Equal(t, start, res[0])

	//降序
	start, end = 5, 1
	res = KNum.Range(start, end)
	assert.Equal(t, 5, len(res))
	assert.Equal(t, start, res[0])

	//起始和结尾相同
	start, end = 3, 3
	res = KNum.Range(start, end)
	assert.Equal(t, 1, len(res))
	assert.Equal(t, start, res[0])
}

func BenchmarkNumber_Range(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Range(0, 9)
	}
}
