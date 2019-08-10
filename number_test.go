package kgo

import (
	"fmt"
	"testing"
)

func TestNumberFormat(t *testing.T) {
	num1 := 123.4567890
	num2 := -123.4567890
	num3 := 123456789.1234567890
	res1 := KNum.NumberFormat(num1, 3, ".", "")
	res2 := KNum.NumberFormat(num2, 0, ".", "")
	res3 := KNum.NumberFormat(num3, 6, ".", ",")
	if res1 != "123.457" || res2 != "-123" || res3 != "123,456,789.123457" {
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

func TestRange(t *testing.T) {
	min := 1
	max := 5
	res := KNum.Range(min, max)
	if len(res) != max-min+1 {
		t.Error("Range fail")
		return
	}

}

func BenchmarkRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Range(0, 9)
	}
}

func TestAbs(t *testing.T) {
	num := -123.456
	res := KNum.Abs(num)
	if res < 0 {
		t.Error("Abs fail")
		return
	}
}

func BenchmarkAbs(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Abs(num)
	}
}

func TestRand(t *testing.T) {
	min := 1
	max := 66666
	res := KNum.Rand(min, max)

	if res < min || res > max {
		t.Error("Rand fail")
		return
	}
	KNum.Rand(5, 5)
}

func TestRandPanicMin(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.Rand(5, 1)
}

func TestRandPanicMax(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.Rand(1, 2147483648)
}

func BenchmarkRand(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Abs(num)
	}
}

func TestRound(t *testing.T) {
	num1 := 0.3
	num2 := 0.6
	res1 := KNum.Round(num1)
	res2 := KNum.Round(num2)
	if int(res1) != 0 || int(res2) != 1 {
		t.Error("Round fail")
		return
	}
}

func BenchmarkRound(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Round(num)
	}
}

func TestFloor(t *testing.T) {
	num1 := 0.3
	num2 := 0.6
	res1 := KNum.Floor(num1)
	res2 := KNum.Floor(num2)
	if int(res1) != 0 || int(res2) != 0 {
		t.Error("Floor fail")
		return
	}
}

func BenchmarkFloor(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Floor(num)
	}
}

func TestCeil(t *testing.T) {
	num1 := 0.3
	num2 := 1.6
	res1 := KNum.Ceil(num1)
	res2 := KNum.Ceil(num2)
	if int(res1) != 1 || int(res2) != 2 {
		t.Error("Ceil fail")
		return
	}
}

func BenchmarkCeil(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Ceil(num)
	}
}

func TestPi(t *testing.T) {
	res := KNum.Pi()
	num := KNum.NumberFormat(res, 2, ".", "")
	if num != "3.14" {
		t.Error("Pi fail")
		return
	}
}

func BenchmarkPi(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Pi()
	}
}

func TestMax(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := KNum.Max(nums...)
	if int(res) != 9 {
		t.Error("Max fail")
		return
	}
}

func TestMaxPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.Max(3)
}

func BenchmarkMax(b *testing.B) {
	b.ResetTimer()
	nums := []float64{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.Max(nums...)
	}
}

func TestMin(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := KNum.Min(nums...)
	if int(res) != -4 {
		t.Error("Min fail")
		return
	}
}

func TestMinPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.Min(3)
}

func BenchmarkMin(b *testing.B) {
	b.ResetTimer()
	nums := []float64{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.Min(nums...)
	}
}

func TestExp(t *testing.T) {
	res := KNum.Exp(1.2)
	if res < 1 {
		t.Error("Exp fail")
		return
	}
}

func BenchmarkExp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Exp(2.34)
	}
}

func TestExpm1(t *testing.T) {
	res := KNum.Expm1(0.1)
	if res < 0 {
		t.Error("Exp fail")
		return
	}
}

func BenchmarkExpm1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Expm1(0.01)
	}
}

func TestPow(t *testing.T) {
	res := KNum.Pow(10, 2)
	if int(res) != 100 {
		t.Error("Pow fail")
		return
	}
}

func BenchmarkPow(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Pow(10, 2)
	}
}
