package kgo

import (
	"fmt"
	"math"
	"strings"
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
	var start, end int = 1, 5
	res0 := KNum.Range(start, end)
	if len(res0) != 5 || res0[0] != start {
		t.Error("Range fail")
		return
	}

	start, end = 5, 1
	res1 := KNum.Range(start, end)
	if len(res1) != 5 || res1[0] != start {
		t.Error("Range fail")
		return
	}

	start, end = 3, 3
	res2 := KNum.Range(start, end)
	if len(res2) != 1 {
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

func TestAbsFloat(t *testing.T) {
	num := -123.456
	res := KNum.AbsFloat(num)
	if res <= 0 {
		t.Error("AbsFloat fail")
		return
	}
}

func BenchmarkAbsFloat(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.AbsFloat(num)
	}
}

func TestAbsInt(t *testing.T) {
	var num int64 = -123
	res := KNum.AbsInt(num)
	if res <= 0 {
		t.Error("AbsInt fail")
		return
	}
}

func BenchmarkAbsInt(b *testing.B) {
	b.ResetTimer()
	var num int64 = -123
	for i := 0; i < b.N; i++ {
		KNum.AbsInt(num)
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
	KNum.Rand(-2147483648, 2147483647)
}

func TestRandPanicMin(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.Rand(5, 1)
}

func BenchmarkRand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Rand(-2147483648, 2147483647)
	}
}

func TestRandInt64(t *testing.T) {
	var min int64 = -9999
	var max int64 = 6666

	res := KNum.RandInt64(min, max)

	if res < min || res > max {
		t.Error("RandInt64 fail")
		return
	}

	res = KNum.RandInt64(5, 5)
	if res != 5 {
		t.Error("RandInt64 fail")
		return
	}
}

func TestRandInt64PanicMin(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.RandInt64(int64(5), int64(1))
}

func TestRandInt64PanicOverflow(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.RandInt64(int64(-9223372036854775808), int64(9223372036854775807))
}

func BenchmarkRandInt64(b *testing.B) {
	b.ResetTimer()
	min := int64(-2147483648)
	max := int64(-2147483647)
	for i := 0; i < b.N; i++ {
		KNum.RandInt64(min, max)
	}
}

func TestRandFloat64(t *testing.T) {
	var min float64 = -9999.0
	var max float64 = 6666.0

	res := KNum.RandFloat64(min, max)
	if res < min || res > max {
		t.Error("Rand fail")
		return
	}
}

func TestRandFloat64PanicMin(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.RandFloat64(float64(5), float64(1))
}

func TestRandFloat64PanicOverflow(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KNum.RandFloat64(-math.MaxFloat64, math.MaxFloat64)
}

func BenchmarkRandFloat64(b *testing.B) {
	b.ResetTimer()
	var min float64 = -9999.0
	var max float64 = 6666.0
	for i := 0; i < b.N; i++ {
		KNum.RandFloat64(min, max)
	}
}

func TestRound(t *testing.T) {
	var tests = []struct {
		num      float64
		expected int
	}{
		{0.3, 0},
		{0.6, 1},
		{-2.4, -2},
		{-3.6, -4},
	}
	for _, test := range tests {
		actual := KNum.Round(test.num)
		if int(actual) != test.expected {
			t.Errorf("Expected KNum.Round(%f) , got %v", test.num, actual)
		}
	}
}

func BenchmarkRound(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.Round(num)
	}
}

func TestRoundPlus(t *testing.T) {
	var tests = []struct {
		num      float64
		pre      int8
		expected float64
	}{
		{0.334, 2, 0.33},
		{0.6467, 2, 0.65},
		{7.258, 2, 7.26},
		{-2.42439, 3, -2.424},
		{-3.611504, 3, -3.612},
	}
	for _, test := range tests {
		actual := KNum.RoundPlus(test.num, test.pre)
		if !KNum.FloatEqual(actual, test.expected, 2) {
			t.Errorf("Expected KNum.RoundPlus(%f, %d) , got %v, not %v", test.num, test.pre, actual, test.expected)
		}
	}
}

func BenchmarkRoundPlus(b *testing.B) {
	b.ResetTimer()
	num := -123.456
	for i := 0; i < b.N; i++ {
		KNum.RoundPlus(num, 2)
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

func TestMaxInt(t *testing.T) {
	nums := []int{-4, 0, 3, 9}
	res := KNum.MaxInt(nums...)
	if res != 9 {
		t.Error("MaxInt fail")
		return
	}

	res = KNum.MaxInt(-1)
	if res != -1 {
		t.Error("MaxInt fail")
		return
	}
}

func TestMaxIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.MaxInt()
}

func BenchmarkMaxInt(b *testing.B) {
	b.ResetTimer()
	nums := []int{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.MaxInt(nums...)
	}
}

func TestMaxFloat64(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := KNum.MaxFloat64(nums...)
	if int(res) != 9 {
		t.Error("MaxFloat64 fail")
		return
	}

	res = KNum.MaxFloat64(-1)
	if int(res) != -1 {
		t.Error("MaxFloat64 fail")
		return
	}
}

func TestMaxFloat64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.MaxFloat64()
}

func BenchmarkMaxFloat64(b *testing.B) {
	b.ResetTimer()
	nums := []float64{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.MaxFloat64(nums...)
	}
}

func TestMax(t *testing.T) {
	nums := []interface{}{-1, 0, "18", true, nil, int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	res := KNum.Max(nums...)
	if int(res) != 18 {
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
	KNum.Max()
}

func BenchmarkMax(b *testing.B) {
	b.ResetTimer()
	nums := []interface{}{-4, 0, 3, 9, "18", true, nil}
	for i := 0; i < b.N; i++ {
		KNum.Max(nums...)
	}
}

func TestMinInt(t *testing.T) {
	nums := []int{0, 3, -4, 5, 9}
	res := KNum.MinInt(nums...)
	if res != -4 {
		t.Error("MinInt fail")
		return
	}

	res = KNum.MinInt(-1)
	if res != -1 {
		t.Error("MinInt fail")
		return
	}
}

func TestMinIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.MinInt()
}

func BenchmarkMinInt(b *testing.B) {
	b.ResetTimer()
	nums := []int{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.MinInt(nums...)
	}
}

func TestMaxMinFloat64(t *testing.T) {
	nums := []float64{-4, 0, 3, 9}
	res := KNum.MinFloat64(nums...)
	if int(res) != -4 {
		t.Error("MinFloat64 fail")
		return
	}

	res = KNum.MinFloat64(-1)
	if int(res) != -1 {
		t.Error("MinFloat64 fail")
		return
	}
}

func TestMinFloat64Panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	KNum.MinFloat64()
}

func BenchmarkMinFloat64(b *testing.B) {
	b.ResetTimer()
	nums := []float64{-4, 0, 3, 9}
	for i := 0; i < b.N; i++ {
		KNum.MinFloat64(nums...)
	}
}

func TestMin(t *testing.T) {
	nums := []interface{}{-1, 0, "18", true, nil, "hello", int8(1), int16(2), int32(3), int64(4), uint(5),
		uint8(6), uint16(7), uint32(8), uint64(9), float32(10.0), float64(11.1)}
	res := KNum.Min(nums...)
	if int(res) != -1 {
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
	KNum.Min()
}

func BenchmarkMin(b *testing.B) {
	b.ResetTimer()
	nums := []interface{}{-4, 0, 3, 9, "18", true, nil}
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

func TestLog(t *testing.T) {
	res1 := KNum.Log(100, 10)
	res2 := KNum.Log(16, 2)
	if int(res1) != 2 || int(res2) != 4 {
		t.Error("Log fail")
		return
	}
}

func BenchmarkLog(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Log(100, 10)
	}
}

func TestByteFormat(t *testing.T) {
	res1 := KNum.ByteFormat(0, 0, "")
	res2 := KNum.ByteFormat(1024000, 2, " ")
	res3 := KNum.ByteFormat(1024000000, 3, " ")
	res4 := KNum.ByteFormat(1024000000000, 4, " ")
	res5 := KNum.ByteFormat(1024000000000000000000000000000000000, 4, " ")

	if res1 == "" || res2 == "" || res3 == "" || res4 == "" || !strings.HasSuffix(res5, "UnKnown") {
		t.Error("ByteFormat fail")
		return
	}
}

func BenchmarkByteFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.ByteFormat(1024000000000, 4, "")
	}
}

func TestIsOddIsEven(t *testing.T) {
	res1 := KNum.IsOdd(-1)
	res2 := KNum.IsOdd(0)
	res3 := KNum.IsEven(2)
	res4 := KNum.IsEven(-3)

	if !res1 || res2 {
		t.Error("IsOdd fail")
		return
	} else if !res3 || res4 {
		t.Error("IsEven fail")
		return
	}
}

func BenchmarkIsOdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsOdd(-1)
	}
}

func BenchmarkIsEven(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsEven(2)
	}
}

func TestSign(t *testing.T) {
	var tests = []struct {
		param    float64
		expected int8
	}{
		{0, 0},
		{-1, -1},
		{10, 1},
		{3.14, 1},
		{-96, -1},
		{-10e-12, -1},
	}
	for _, test := range tests {
		actual := KNum.NumSign(test.param)
		if actual != test.expected {
			t.Errorf("Expected NumSign(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkSign(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.NumSign(-2)
	}
}

func TestIsNegative(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, false},
		{-1, true},
		{10, false},
		{3.14, false},
		{-96, true},
		{-10e-12, true},
	}
	for _, test := range tests {
		actual := KNum.IsNegative(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNegative(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsNegative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNegative(9)
	}
}

func TestIsPositive(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, false},
		{-1, false},
		{10, true},
		{3.14, true},
		{-96, false},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := KNum.IsPositive(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsPositive(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsPositive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsPositive(-2)
	}
}

func TestIsNonNegative(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, false},
		{10, true},
		{3.14, true},
		{-96, false},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := KNum.IsNonNegative(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNonNegative(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsNonNegative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNonNegative(8)
	}
}

func TestIsNonPositive(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, true},
		{10, false},
		{3.14, false},
		{-96, true},
		{-10e-12, true},
	}
	for _, test := range tests {
		actual := KNum.IsNonPositive(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNonPositive(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsNonPositive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNonPositive(34)
	}
}

func TestIsWhole(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, true},
		{10, true},
		{3.14, false},
		{-96, true},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := KNum.IsWhole(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsWhole(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsWhole(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsWhole(1.2)
	}
}

func TestIsNatural(t *testing.T) {
	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, false},
		{10, true},
		{3.14, false},
		{96, true},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := KNum.IsNatural(test.param)
		if actual != test.expected {
			t.Errorf("Expected IsNatural(%v) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func BenchmarkIsNatural(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNatural(-3)
	}
}

func TestInRangeInt(t *testing.T) {
	var testAsInts = []struct {
		param    int
		left     int
		right    int
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsInts {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func BenchmarkInRangeInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeInt(6, -1, 9)
	}
}

func TestInRangeFloat32(t *testing.T) {
	var tests = []struct {
		param    float32
		left     float32
		right    float32
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range tests {
		actual := KNum.InRangeFloat32(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeFloat32(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func BenchmarkInRangeFloat32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeFloat32(3.14, 0.1, 65.27)
	}
}

func TestInRangeFloat64(t *testing.T) {
	var tests = []struct {
		param    float64
		left     float64
		right    float64
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range tests {
		actual := KNum.InRangeFloat64(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeFloat64(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func BenchmarkInRangeFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeFloat64(3.14, 0.1, 65.27)
	}
}

func TestInRange(t *testing.T) {
	var testsInt = []struct {
		param    int
		left     int
		right    int
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range testsInt {
		actual := KNum.InRange(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRange(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testsFloat32 = []struct {
		param    float32
		left     float32
		right    float32
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range testsFloat32 {
		actual := KNum.InRange(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRange(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testsFloat64 = []struct {
		param    float64
		left     float64
		right    float64
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range testsFloat64 {
		actual := KNum.InRange(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRange(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testsTypeMix = []struct {
		param    int
		left     float64
		right    float64
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range testsTypeMix {
		actual := KNum.InRange(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRange(%v, %v, %v) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}

	res := KNum.InRange(89, -1.2, 999.123)
	if !res {
		t.Error("InRange fail")
	}
	KNum.InRange("1", 0, 3)
	KNum.InRange("hello", []byte{}, 3)
}

func BenchmarkInRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRange(89, -1.2, 999.123)
	}
}

func TestSumInt(t *testing.T) {
	sum := KNum.SumInt(0, 1, -2, 3, 5)
	if sum != 7 {
		t.Error("SumInt fail")
		return
	}
}

func BenchmarkSumInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.SumInt(0, 1, -2, 3, 5)
	}
}

func TestSumFloat64(t *testing.T) {
	sum := KNum.SumFloat64(0.0, 1.1, -2.2, 3.30, 5.55)
	if KNum.NumberFormat(sum, 2, ".", "") != "7.75" {
		t.Error("SumFloat64 fail")
		return
	}
}

func BenchmarkSumFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.SumInt(0, 1, -2, 3, 5)
	}
}

func TestSumAny(t *testing.T) {
	sum := KNum.Sum(1, 0, 1.2, -3, false, nil, "4")
	if KNum.NumberFormat(sum, 2, ".", "") != "3.20" {
		t.Error("Sum fail")
		return
	}

	sum = KNum.Sum(true, false, "false", "true")
	if KNum.NumberFormat(sum, 2, ".", "") != "0.00" {
		t.Error("Sum fail")
		return
	}
}

func BenchmarkSumAny(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Sum(1, 0, 1.2, -3, false, nil, "4")
	}
}

func TestAverageInt(t *testing.T) {
	var res1, res2, res3 float64

	res1 = KNum.AverageInt()
	res2 = KNum.AverageInt(1)
	res3 = KNum.AverageInt(1, 2, 3, 4, 5, 6)

	if res1 != 0 || int(res2) != 1 || KNum.NumberFormat(res3, 2, ".", "") != "3.50" {
		t.Error("AverageInt fail")
		return
	}
}

func BenchmarkAverageInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AverageInt(1, 2, 3, 4, 5, 6)
	}
}

func TestAverageFloat64(t *testing.T) {
	var res1, res2, res3 float64

	res1 = KNum.AverageFloat64()
	res2 = KNum.AverageFloat64(1)
	res3 = KNum.AverageFloat64(1, 2, 3, 4, 5, 6)

	if res1 != 0 || int(res2) != 1 || KNum.NumberFormat(res3, 2, ".", "") != "3.50" {
		t.Error("AverageFloat64 fail")
		return
	}
}

func BenchmarkAverageFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AverageFloat64(1, 2, 3, 4, 5, 6)
	}
}

func TestAverage(t *testing.T) {
	var res1, res2, res3 float64

	res1 = KNum.Average()
	res2 = KNum.Average(1)
	res3 = KNum.Average(1, 2.0, "3", 4.0, 5, "6")

	if res1 != 0 || int(res2) != 1 || KNum.NumberFormat(res3, 2, ".", "") != "3.50" {
		t.Error("Average fail")
		return
	}
}

func BenchmarkAverage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Average(1, 2.0, "3", 4.0, 5, "6")
	}
}

func TestFloatEqual(t *testing.T) {
	var f1, f2 float64

	f1 = 1.2345678
	f2 = 1.2345679

	chk1 := KNum.FloatEqual(f1, f2)
	chk2 := KNum.FloatEqual(f1, f2, 6)

	if chk1 || !chk2 {
		t.Error("FloatEqual fail")
		return
	}
}

func BenchmarkFloatEqual(b *testing.B) {
	b.ResetTimer()
	f1 := 1.2345678
	f2 := 1.2345679
	for i := 0; i < b.N; i++ {
		KNum.FloatEqual(f1, f2)
	}
}

func TestGeoDistance(t *testing.T) {
	var lat1, lng1, lat2, lng2, res1, res2 float64

	lat1, lng1 = 30.0, 45.0
	lat2, lng2 = 40.0, 90.0
	res1 = KNum.GeoDistance(lng1, lat1, lng2, lat2)

	lat1, lng1 = 390.0, 405.0
	lat2, lng2 = -320.0, 90.0
	res2 = KNum.GeoDistance(lng1, lat1, lng2, lat2)

	if res1 <= 0 || res2 <= 0 || !KNum.FloatEqual(res1, res2) {
		t.Error("GeoDistance fail")
		return
	}
}

func BenchmarkGeoDistance(b *testing.B) {
	b.ResetTimer()
	lat1, lng1 := 30.0, 45.0
	lat2, lng2 := 40.0, 90.0
	for i := 0; i < b.N; i++ {
		KNum.GeoDistance(lng1, lat1, lng2, lat2)
	}
}

func TestIsNan(t *testing.T) {
	res1 := KNum.IsNan(math.Acos(1.01))
	res2 := KNum.IsNan(123.456)
	res3 := KNum.IsNan("4.367")
	res4 := KNum.IsNan("hello")

	if !res1 || res2 || res3 || !res4 {
		t.Error("IsNan fail")
		return
	}
}

func BenchmarkIsNan(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNan(123.456)
	}
}

func TestPercent(t *testing.T) {
	tests := []struct {
		val      interface{}
		total    interface{}
		expected float64
	}{
		{0, "", 0},
		{0, 0, 0},
		{"1", "20", 5},
		{2.5, 10, 25},
		{3.46, 12.24, 28.2679},
	}

	for _, test := range tests {
		actual := KNum.Percent(test.val, test.total)
		if !KNum.FloatEqual(actual, test.expected, 4) {
			t.Errorf("Expected Percent(%v, %v) to be %v, got %v", test.val, test.total, test.expected, actual)
			return
		}
	}

}

func BenchmarkPercent(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Percent(23.456, 68)
	}
}

func TestIsNaturalRange(t *testing.T) {
	arr1 := []int{1, 2, 3}
	arr2 := []int{0, 3, 1, 2}
	arr3 := []int{0, 1, 2, 3}
	arr4 := []int{0, 1, 3, 4}

	res1 := KNum.IsNaturalRange(arr1, false)
	if res1 {
		t.Error("IsNaturalRange fail")
		return
	}

	res2 := KNum.IsNaturalRange(arr2, false)
	res3 := KNum.IsNaturalRange(arr2, true)
	if !res2 || res3 {
		t.Error("IsNaturalRange fail")
		return
	}

	res4 := KNum.IsNaturalRange(arr3, false)
	res5 := KNum.IsNaturalRange(arr3, true)
	if !res4 || !res5 {
		t.Error("IsNaturalRange fail")
		return
	}

	res6 := KNum.IsNaturalRange(arr4, false)
	if res6 {
		t.Error("IsNaturalRange fail")
		return
	}
}

func BenchmarkIsNaturalRange(b *testing.B) {
	b.ResetTimer()
	arr := []int{0, 1, 2, 3}
	for i := 0; i < b.N; i++ {
		KNum.IsNaturalRange(arr, false)
	}
}
