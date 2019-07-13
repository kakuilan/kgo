package kgo

import (
	"testing"
)

func TestNumberFormat(t *testing.T) {
	num1 := 123.4567890
	num2 := -123.4567890
	res1 := KNum.NumberFormat(num1, 3, ".", "")
	res2 := KNum.NumberFormat(num2, 0, ".", "")
	if res1 != "123.457" || res2 != "-123" {
		t.Error("NumberFormat fail")
		return
	}
}

func BenchmarkNumberFormat(b *testing.B) {
	b.ResetTimer()
	num := 123.4567890
	for i := 0; i < b.N; i++ {
		KNum.NumberFormat(num, 3, ".", "")
	}
}
