package kgo

import (
	"fmt"
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

func TestByteFormat(t *testing.T) {
	res1 := KNum.ByteFormat(0, 0)
	res2 := KNum.ByteFormat(1024000, 2)
	res3 := KNum.ByteFormat(1024000000, 3)
	res4 := KNum.ByteFormat(1024000000000, 4)
	res5 := KNum.ByteFormat(1024000000000000000000000000000000000, 4)

	if res1 == "" || res2 == "" || res3 == "" || res4 == "" || !strings.HasSuffix(res5, "UnKnown") {
		t.Error("ByteFormat fail")
		return
	}
}

func BenchmarkByteFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.ByteFormat(1024000000000, 4)
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
		expected float64
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
		{0, false},
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

	var testAsInt8s = []struct {
		param    int8
		left     int8
		right    int8
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
	for _, test := range testAsInt8s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int8", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsInt16s = []struct {
		param    int16
		left     int16
		right    int16
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
	for _, test := range testAsInt16s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int16", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsInt32s = []struct {
		param    int32
		left     int32
		right    int32
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
	for _, test := range testAsInt32s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int32", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsInt64s = []struct {
		param    int64
		left     int64
		right    int64
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
	for _, test := range testAsInt64s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type int64", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsUInts = []struct {
		param    uint
		left     uint
		right    uint
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{0, 0, 1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsUInts {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type uint", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsUInt8s = []struct {
		param    uint8
		left     uint8
		right    uint8
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{0, 0, 1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsUInt8s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type uint", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsUInt16s = []struct {
		param    uint16
		left     uint16
		right    uint16
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{0, 0, 1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsUInt16s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type uint", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsUInt32s = []struct {
		param    uint32
		left     uint32
		right    uint32
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{0, 0, 1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsUInt32s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type uint", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsUInt64s = []struct {
		param    uint64
		left     uint64
		right    uint64
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{0, 0, 1, true},
		{0, 10, 5, false},
	}
	for _, test := range testAsUInt64s {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type uint", test.param, test.left, test.right, test.expected, actual)
		}
	}

	var testAsStrings = []struct {
		param    string
		left     string
		right    string
		expected bool
	}{
		{"0", "0", "0", true},
		{"1", "0", "0", false},
		{"-1", "0", "0", false},
		{"0", "-1", "1", true},
		{"0", "0", "1", true},
		{"0", "-1", "0", true},
		{"0", "0", "-1", true},
		{"0", "10", "5", false},
	}
	for _, test := range testAsStrings {
		actual := KNum.InRangeInt(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRangeInt(%v, %v, %v) to be %v, got %v using type string", test.param, test.left, test.right, test.expected, actual)
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
		{0, 0, 0, false},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, false},
		{0, 0, 1, false},
		{0, -1, 0, false},
		{0, 0, -1, false},
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
}

func BenchmarkInRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRange(89, -1.2, 999.123)
	}
}
