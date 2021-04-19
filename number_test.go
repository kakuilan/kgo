package kgo

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestNumber_NumberFormat(t *testing.T) {
	var res string

	res = KNum.NumberFormat(floNum1, 10, ".", "")
	assert.Equal(t, "12345.1234567890", res)

	res = KNum.NumberFormat(floNum2, 6, ".", ",")
	assert.Equal(t, "12,345,678.123457", res)

	res = KNum.NumberFormat(floNum3, 0, ".", "")
	assert.Equal(t, "-123", res)

	res = KNum.NumberFormat(math.Pi, 15, ".", "")
	assert.Equal(t, "3.141592653589793", res)
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

func TestNumber_AbsFloat(t *testing.T) {
	var res float64

	res = KNum.AbsFloat(floNum3)
	assert.Greater(t, res, 0.0)
}

func BenchmarkNumber_AbsFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AbsFloat(floNum3)
	}
}

func TestNumber_AbsInt(t *testing.T) {
	var res int64

	res = KNum.AbsInt(-123)
	assert.Greater(t, res, int64(0))
}

func BenchmarkNumber_AbsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AbsInt(-123)
	}
}

func TestNumber_FloatEqual(t *testing.T) {
	var res bool

	//默认小数位
	res = KNum.FloatEqual(floNum1, floNum4)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 0)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 11)
	assert.True(t, res)

	res = KNum.FloatEqual(floNum1, floNum4, 12)
	assert.False(t, res)
}

func BenchmarkNumber_FloatEqual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.FloatEqual(floNum1, floNum4)
	}
}

func TestNumber_RandInt64(t *testing.T) {
	var min, max, res int64

	min, max = -9, 9
	res = KNum.RandInt64(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	//最小最大值调换
	min, max = 9, -9
	res = KNum.RandInt64(min, max)
	assert.GreaterOrEqual(t, res, max)
	assert.LessOrEqual(t, res, min)

	res = KNum.RandInt64(max, max)
	assert.Equal(t, res, max)

	KNum.RandInt64(INT64_MIN, INT64_MAX)
}

func BenchmarkNumber_RandInt64(b *testing.B) {
	b.ResetTimer()
	var min, max int64
	min, max = 9, -9
	for i := 0; i < b.N; i++ {
		KNum.RandInt64(min, max)
	}
}

func TestNumber_RandInt(t *testing.T) {
	var res int
	min, max := -9, 9
	res = KNum.RandInt(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	//最小最大值调换
	min, max = 9, -9
	res = KNum.RandInt(min, max)
	assert.GreaterOrEqual(t, res, max)
	assert.LessOrEqual(t, res, min)

	res = KNum.RandInt(max, max)
	assert.Equal(t, res, max)

	KNum.RandInt(INT_MIN, INT_MAX)
}

func BenchmarkNumber_RandInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.RandInt(-9, 9)
	}
}

func TestNumber_Rand(t *testing.T) {
	var res int
	min, max := -9, 9
	res = KNum.Rand(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	res = KNum.Rand(max, max)
	assert.Equal(t, res, max)
}

func BenchmarkNumber_Rand(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Rand(-9, 9)
	}
}

func TestNumber_RandFloat64(t *testing.T) {
	var res float64

	min, max := floNum3, floNum1
	res = KNum.RandFloat64(min, max)
	assert.GreaterOrEqual(t, res, min)
	assert.LessOrEqual(t, res, max)

	//最小最大值调换
	min, max = floNum1, floNum3
	res = KNum.RandFloat64(min, max)
	assert.GreaterOrEqual(t, res, max)
	assert.LessOrEqual(t, res, min)

	KNum.RandFloat64(max, max)
	KNum.RandFloat64(-math.MaxFloat64, math.MaxFloat64)
}

func BenchmarkNumber_RandFloat64(b *testing.B) {
	b.ResetTimer()
	min, max := floNum3, floNum1
	for i := 0; i < b.N; i++ {
		KNum.RandFloat64(min, max)
	}
}

func TestNumber_Round(t *testing.T) {
	var tests = []struct {
		num      float64
		expected int
	}{
		{0.3, 0},
		{0.6, 1},
		{1.55, 2},
		{-2.4, -2},
		{-3.6, -4},
	}
	for _, test := range tests {
		actual := KNum.Round(test.num)
		assert.Equal(t, test.expected, int(actual))
	}
}

func BenchmarkNumber_Round(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Round(floNum3)
	}
}

func TestNumber_RoundPlus(t *testing.T) {
	var res float64

	res = KNum.RoundPlus(floNum1, 2)
	assert.True(t, KNum.FloatEqual(res, 12345.12, 2))

	res = KNum.RoundPlus(floNum1, 4)
	assert.True(t, KNum.FloatEqual(res, 12345.1235, 4))

	res = KNum.RoundPlus(floNum3, 4)
	assert.True(t, KNum.FloatEqual(res, -123.4568, 4))
}

func BenchmarkNumber_RoundPlus(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.RoundPlus(floNum1, 2)
	}
}

func TestNumber_Floor(t *testing.T) {
	var res float64

	res = KNum.Floor(flPi2)
	assert.Equal(t, int(res), 3)

	res = KNum.Floor(float64(floSpeedLight))
	assert.Equal(t, int(res), 2)

	res = KNum.Floor(floNum3)
	assert.Equal(t, int(res), -124)
}

func BenchmarkNumber_Floor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Floor(flPi2)
	}
}

func TestNumber_Ceil(t *testing.T) {
	var res float64

	res = KNum.Ceil(flPi2)
	assert.Equal(t, int(res), 4)

	res = KNum.Ceil(float64(floSpeedLight))
	assert.Equal(t, int(res), 3)

	res = KNum.Ceil(floNum3)
	assert.Equal(t, int(res), -123)
}

func BenchmarkNumber_Ceil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Ceil(flPi2)
	}
}

func TestNumber_MaxInt(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res int

	res = KNum.MaxInt(intSlc...)
	assert.Equal(t, res, 15)

	//无输入
	KNum.MaxInt()
}

func BenchmarkNumber_MaxInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.MaxInt(intSlc...)
	}
}

func TestNumber_MaxFloat64(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res float64

	res = KNum.MaxFloat64(flo64Slc...)
	assert.Equal(t, res, floAvogadro)

	//无输入
	KNum.MaxFloat64()
}

func BenchmarkNumber_MaxFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.MaxFloat64(flo64Slc...)
	}
}

func TestNumber_Max(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res float64

	res = KNum.Max(slItf...)
	assert.Equal(t, res, floAvogadro)

	//非数值输入
	res = KNum.Max(strHello, admTesDir)
	assert.Equal(t, res, 0.0)

	//无输入
	KNum.Max()
}

func BenchmarkNumber_Max(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Max(slItf...)
	}
}

func TestNumber_MinInt(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res int

	res = KNum.MinInt(intSlc...)
	assert.Equal(t, res, 0)

	//无输入
	KNum.MinInt()
}

func BenchmarkNumber_MinInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.MinInt(intSlc...)
	}
}

func TestNumber_MinFloat64(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res float64

	res = KNum.MinFloat64(flo64Slc...)
	assert.Equal(t, res, floPlanck)

	//无输入
	KNum.MinFloat64()
}

func BenchmarkNumber_MinFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.MinFloat64(flo64Slc...)
	}
}

func TestNumber_Min(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res float64

	res = KNum.Min(slItf...)
	assert.Equal(t, res, floNum3)

	//非数值输入
	res = KNum.Min(strHello, admTesDir)
	assert.Equal(t, res, 0.0)

	//无输入
	KNum.Min()
}

func BenchmarkNumber_Min(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Min(slItf...)
	}
}

func TestNumber_Exp(t *testing.T) {
	var res float64

	res = KNum.Exp(1.1)
	assert.Greater(t, res, 1.0)
}

func BenchmarkNumber_Exp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Exp(1.1)
	}
}

func TestNumber_Expm1(t *testing.T) {
	var res float64

	res = KNum.Expm1(1.1)
	assert.Greater(t, res, 0.0)
}

func BenchmarkNumber_Expm1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Expm1(1.1)
	}
}

func TestNumber_Pow(t *testing.T) {
	var res float64

	res = KNum.Pow(10, 2)
	assert.Equal(t, res, 100.0)
}

func BenchmarkNumber_Equal(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Pow(10, 2)
	}
}

func TestNumber_Log(t *testing.T) {
	var res float64

	res = KNum.Log(100, 10)
	assert.Equal(t, res, 2.0)
}

func BenchmarkNumber_Log(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.Log(100, 10)
	}
}

func TestNumber_ByteFormat(t *testing.T) {
	var res string

	res = KNum.ByteFormat(0, 0, "")
	assert.Equal(t, res, "0B")

	res = KNum.ByteFormat(floNum5, 4, " ")
	assert.NotEmpty(t, res)

	res = KNum.ByteFormat(floNum6, 4, "")
	assert.Contains(t, res, Unknown)
}

func BenchmarkNumber_ByteFormat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.ByteFormat(floNum6, 4, "")
	}
}

func TestNumber_IsOdd(t *testing.T) {
	var tests = []struct {
		num      int
		expected bool
	}{
		{-4, false},
		{-1, true},
		{0, false},
		{3, true},
	}

	var actual bool
	for _, test := range tests {
		actual = KNum.IsOdd(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsOdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsOdd(1)
	}
}

func TestNumber_IsEven(t *testing.T) {
	var tests = []struct {
		num      int
		expected bool
	}{
		{-4, true},
		{-1, false},
		{0, true},
		{3, false},
	}

	var actual bool
	for _, test := range tests {
		actual = KNum.IsEven(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsEven(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsEven(2)
	}
}

func TestNumber_NumSign(t *testing.T) {
	var tests = []struct {
		num      float64
		expected int8
	}{
		{0, 0},
		{floNum1, 1},
		{math.Pi, 1},
		{floNum3, -1},
		{floNum7, -1},
	}
	var actual int8
	for _, test := range tests {
		actual = KNum.NumSign(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func TestNumber_IsNegative(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, false},
		{floNum1, false},
		{math.Pi, false},
		{floNum3, true},
		{floNum7, true},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsNegative(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsNegative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNegative(floNum1)
	}
}

func TestNumber_IsPositive(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, false},
		{floNum1, true},
		{math.Pi, true},
		{floNum3, false},
		{floNum7, false},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsPositive(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsPositive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsPositive(floNum1)
	}
}

func TestNumber_IsNonNegative(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, true},
		{floNum1, true},
		{math.Pi, true},
		{floNum3, false},
		{floNum7, false},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsNonNegative(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsNonNegative(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNonNegative(floNum1)
	}
}

func TestNumber_IsNonPositive(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, true},
		{floNum1, false},
		{math.Pi, false},
		{floNum3, true},
		{floNum7, true},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsNonPositive(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsNonPositive(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNonPositive(floNum1)
	}
}

func TestNumber_IsWhole(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, true},
		{10, true},
		{math.Pi, false},
		{floNum3, false},
		{floNum7, false},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsWhole(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsWhole(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsWhole(floNum5)
	}
}

func TestNumber_IsNatural(t *testing.T) {
	var tests = []struct {
		num      float64
		expected bool
	}{
		{0, true},
		{10, true},
		{-1, false},
		{math.Pi, false},
		{floNum3, false},
		{floNum7, false},
	}
	var actual bool
	for _, test := range tests {
		actual = KNum.IsNatural(test.num)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_IsNatural(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.IsNatural(9)
	}
}

func TestNumber_InRangeInt(t *testing.T) {
	var testAsInts = []struct {
		num      int
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
		{1, 0, 5, true},
	}
	var actual bool
	for _, test := range testAsInts {
		actual = KNum.InRangeInt(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_InRangeInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeInt(5, 1, 9)
	}
}

func TestNumber_InRangeFloat64(t *testing.T) {
	var testAsInts = []struct {
		num      float64
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
		{1, 0, 5, true},
	}
	var actual bool
	for _, test := range testAsInts {
		actual = KNum.InRangeFloat64(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_InRangeFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeFloat64(5, 1, 9)
	}
}

func TestNumber_InRangeFloat32(t *testing.T) {
	var testAsInts = []struct {
		num      float32
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
		{1, 0, 5, true},
	}
	var actual bool
	for _, test := range testAsInts {
		actual = KNum.InRangeFloat32(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkNumber_InRangeFloat32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRangeFloat32(5, 1, 9)
	}
}

func TestNumber_InRange(t *testing.T) {
	var actual bool

	//int
	var testsInt = []struct {
		num      int
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
		actual = KNum.InRange(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}

	//float32
	var testsFloat32 = []struct {
		num      float32
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
		actual = KNum.InRange(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}

	//float64
	var testsFloat64 = []struct {
		num      float64
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
		actual = KNum.InRange(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}

	//mix
	var testsTypeMix = []struct {
		num      int
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
		actual = KNum.InRange(test.num, test.left, test.right)
		assert.Equal(t, actual, test.expected)
	}

	//other
	KNum.InRange("1", 0, 3)
	KNum.InRange("hello", []byte{}, 3)
}

func BenchmarkNumber_InRange(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.InRange(89, -1.2, 999.123)
	}
}

func TestNumber_SumInt(t *testing.T) {
	var res int

	res = KNum.SumInt(naturalArr[:]...)
	assert.Equal(t, res, 55)
}

func BenchmarkNumber_SumInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.SumInt(intSlc...)
	}
}

func TestNumber_SumFloat64(t *testing.T) {
	var res float64

	res = KNum.SumFloat64(flo64Slc2...)
	assert.Equal(t, "12370248.06", KNum.NumberFormat(res, 2, ".", ""))
}

func BenchmarkNumber_SumFloat64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.SumFloat64(flo64Slc2...)
	}
}

func TestNumber_Sum(t *testing.T) {
	var res float64
	var nums []interface{}

	nums = KArr.ArrayShuffle(naturalArr)
	res = KNum.Sum(nums...)
	assert.Equal(t, res, 55.0)

	nums = KArr.ArrayShuffle(flo64Slc2)
	res = KNum.Sum(nums...)
	assert.Equal(t, "12370248.06", KNum.NumberFormat(res, 2, ".", ""))
}

func BenchmarkNumber_Sum(b *testing.B) {
	b.ResetTimer()
	nums := KArr.ArrayShuffle(naturalArr)
	for i := 0; i < b.N; i++ {
		KNum.Sum(nums...)
	}
}

func TestNumber_AverageInt(t *testing.T) {
	var res float64

	res = KNum.AverageInt()
	assert.Equal(t, 0.0, res)

	res = KNum.AverageInt(intTen)
	assert.Equal(t, float64(intTen), res)

	res = KNum.AverageInt(naturalArr[:]...)
	assert.Equal(t, 5.0, res)
}

func BenchmarkNumber_AverageInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KNum.AverageInt(intSlc...)
	}
}
