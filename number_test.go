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
