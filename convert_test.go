package kgo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInt2Str(t *testing.T) {
	tim := KConv.Int2Str(KTime.Time())
	if fmt.Sprint(reflect.TypeOf(tim)) != "string" {
		t.Error("Int2Str fail")
		return
	}

	//非整型的转为空
	res := KConv.Int2Str(1.23)
	if res != "" {
		t.Error("Int2Str fail")
		return
	}
}

func BenchmarkInt2Str(b *testing.B) {
	b.ResetTimer()
	tim := KTime.Time()
	for i := 0; i < b.N; i++ {
		KConv.Int2Str(tim)
	}
}

func TestIntFloat2Str(t *testing.T) {
	fl := float32(1234.567890)
	f2 := float64(1234.567890)
	res1 := KConv.Float2Str(fl, 4)
	res2 := KConv.Float2Str(f2, 8)
	if fmt.Sprint(reflect.TypeOf(res1)) != fmt.Sprint(reflect.TypeOf(res2)) {
		t.Error("Int2Str fail")
		return
	}

	//非浮点的转为空
	res := KConv.Float2Str(123, 2)
	if res != "" {
		t.Error("Float2Str fail")
		return
	}
}

func Benchmark32Float2Str(b *testing.B) {
	b.ResetTimer()
	fl := float32(1234.567890)
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(fl, 4)
	}
}

func Benchmark64Float2Str(b *testing.B) {
	b.ResetTimer()
	f2 := float64(1234.567890)
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(f2, 8)
	}
}

func TestBool2Str(t *testing.T) {
	res1 := KConv.Bool2Str(true)
	res2 := KConv.Bool2Str(false)
	if res1 != "true" {
		t.Error("Bool2Str fail")
		return
	} else if res2 != "false" {
		t.Error("Bool2Str fail")
		return
	}
}

func BenchmarkBool2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Str(true)
	}
}

func TestBool2Int(t *testing.T) {
	res1 := KConv.Bool2Int(true)
	res2 := KConv.Bool2Int(false)
	if res1 != 1 || res2 != 0 {
		t.Error("Bool2Int fail")
		return
	}
}

func BenchmarkBool2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Int(true)
	}
}

func TestStr2IntStrict(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := KConv.Str2IntStrict("abc123", 8, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "int8" {
		t.Error("Str2IntStrict fail")
		return
	}
}

func TestStr2Int(t *testing.T) {
	res := KConv.Str2Int("123")
	if fmt.Sprint(reflect.TypeOf(res)) != "int" {
		t.Error("Str2Int fail")
		return
	}

	var tests = []struct {
		param    string
		expected int
	}{
		{"", 0},
		{"123", 123},
		{"123.45", 0},
		{"True", 1},
		{"false", 0},
	}

	for _, test := range tests {
		actual := KConv.Str2Int(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToBool(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkStr2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("-123")
	}
}

func TestStr2Int8(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Int8(tim)
	if res > 127 {
		t.Error("Str2Int8 fail")
		return
	}
}

func BenchmarkStr2Int8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int8("128")
	}
}

func TestStr2Int16(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Int16(tim)
	if res > 32767 {
		t.Error("Str2Int16 fail")
		return
	}
}

func BenchmarkStr2Int16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int16("32768")
	}
}

func TestStr2Int32(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Int32(tim)
	if res > 2147483647 {
		t.Error("Str2Int32 fail")
		return
	}
}

func BenchmarkStr2Int32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int32("2147483647")
	}
}

func TestStr2Int64(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Int64(tim)
	if res > INT64_MAX {
		t.Error("Str2Int64 fail")
		return
	}
}

func BenchmarkStr2Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int64("9223372036854775808")
	}
}

func TestStrictStr2Uint(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()
	res := KConv.StrictStr2Uint("abc123", 8, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "uint8" {
		t.Error("StrictStr2Uint fail")
		return
	}
}

func TestStr2Uint(t *testing.T) {
	res := KConv.Str2Uint("-123")
	if fmt.Sprint(reflect.TypeOf(res)) != "uint" {
		t.Error("Str2Uint fail")
		return
	}
}

func BenchmarkStr2Uint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("123")
	}
}

func TestStr2Uint8(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Uint8(tim)
	if res > 255 {
		t.Error("Str2Uint8 fail")
		return
	}
}

func BenchmarkStr2Uint8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint8("256")
	}
}

func TestStr2Uint16(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Uint16(tim)
	if res > 65535 {
		t.Error("Str2Uint16 fail")
		return
	}
}

func BenchmarkStr2Uint16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint16("65536")
	}
}

func TestStr2Uint32(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Uint32(tim)
	if res > 4294967295 {
		t.Error("Str2Uint32 fail")
		return
	}
}

func BenchmarkStr2Uint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint32("4294967296")
	}
}

func TestStr2Uint64(t *testing.T) {
	tim := KConv.Int2Str(KTime.MicroTime())
	res := KConv.Str2Uint64(tim)
	if res > UINT64_MAX {
		t.Error("Str2Uint64 fail")
		return
	}
}

func BenchmarkStr2Uint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint64("9223372036854775808")
	}
}

func TestStr2FloatStrict(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := KConv.Str2FloatStrict("abc123", 32, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "float32" {
		t.Error("Str2FloatStrict fail")
		return
	}
}

func TestStr2Float32(t *testing.T) {
	res := KConv.Str2Float32("123.456")
	if fmt.Sprint(reflect.TypeOf(res)) != "float32" {
		t.Error("Str2Float32 fail")
		return
	}
}

func BenchmarkStr2Float32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float32("123.456")
	}
}

func TestStr2Float64(t *testing.T) {
	res := KConv.Str2Float64("123.456")
	if fmt.Sprint(reflect.TypeOf(res)) != "float64" {
		t.Error("Str2Float64 fail")
		return
	}
}

func BenchmarkStr2Float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float64("123.456")
	}
}

func TestStr2Bool(t *testing.T) {
	res1 := KConv.Str2Bool("true")
	res2 := KConv.Str2Bool("True")
	res3 := KConv.Str2Bool("TRUE")
	res4 := KConv.Str2Bool("Hello")

	if !res1 || !res2 || !res3 {
		t.Error("Str2Bool fail")
		return
	} else if res4 {
		t.Error("Str2Bool fail")
		return
	}
}

func BenchmarkStr2Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Bool("123.456")
	}
}

func TestStr2ByteSlice(t *testing.T) {
	str := `hello world!`
	res := KConv.Str2ByteSlice(str)
	if fmt.Sprint(reflect.TypeOf(res)) != "[]uint8" {
		t.Error("Str2ByteSlice fail")
		return
	}
}

func BenchmarkStr2ByteSlice(b *testing.B) {
	b.ResetTimer()
	str := `hello world!
// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
`
	for i := 0; i < b.N; i++ {
		KConv.Str2ByteSlice(str)
	}
}

func TestByteSlice2Str(t *testing.T) {
	sli := []byte("hello world!")
	res := KConv.ByteSlice2Str(sli)
	if fmt.Sprint(reflect.TypeOf(res)) != "string" {
		t.Error("ByteSlice2Str fail")
		return
	}
}

func BenchmarkByteSlice2Str(b *testing.B) {
	b.ResetTimer()
	sli := []byte(`hello world!
// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
`)
	for i := 0; i < b.N; i++ {
		KConv.ByteSlice2Str(sli)
	}
}

func TestDec2Bin(t *testing.T) {
	var num int64 = 8
	res := KConv.Dec2Bin(num)
	if res != "1000" {
		t.Error("Dec2Bin fail")
		return
	}
}

func BenchmarkDec2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Bin(10)
	}
}

func TestBin2dec(t *testing.T) {
	res, err := KConv.Bin2dec("1000")
	if err != nil || res != 8 {
		t.Error("Bin2dec fail")
		return
	}
	_, _ = KConv.Bin2dec("hello")
}

func BenchmarkBin2dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2dec("1000")
	}
}

func TestHex2Bin(t *testing.T) {
	_, err := KConv.Hex2Bin("123abff")
	if err != nil {
		t.Error("Hex2Bin fail")
		return
	}
	_, _ = KConv.Hex2Bin("hello")
}

func BenchmarkHex2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Bin("123abff")
	}
}

func TestBin2Hex(t *testing.T) {
	_, err := KConv.Bin2Hex("1001000111010101111111111")
	if err != nil {
		t.Error("Bin2Hex fail")
		return
	}
	_, _ = KConv.Bin2Hex("hello")
}

func BenchmarkBin2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2Hex("1001000111010101111111111")
	}
}

func TestDec2Hex(t *testing.T) {
	res := KConv.Dec2Hex(1234567890)
	if res != "499602d2" {
		t.Error("Dec2Hex fail")
		return
	}
}

func BenchmarkDec2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Hex(1234567890)
	}
}

func TestHex2Dec(t *testing.T) {
	res1, err := KConv.Hex2Dec("123abf")
	res2, _ := KConv.Hex2Dec("0x123abf")
	if err != nil {
		t.Error("Hex2Dec fail")
		return
	} else if res1 != res2 {
		t.Error("Hex2Dec fail")
		return
	}
}

func BenchmarkHex2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Dec("123abf")
	}
}

func TestDec2Oct(t *testing.T) {
	res := KConv.Dec2Oct(123456789)
	if res != "726746425" {
		t.Error("Dec2Oct fail")
		return
	}
}

func BenchmarkDec2Oct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Oct(123456789)
	}
}

func TestOct2dec(t *testing.T) {
	res1, err := KConv.Oct2dec("726746425")
	res2, _ := KConv.Oct2dec("0726746425")
	if err != nil {
		t.Error("Oct2dec fail")
		return
	} else if res1 != res2 {
		t.Error("Oct2dec fail")
		return
	}
}

func BenchmarkOct2dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Oct2dec("726746425")
	}
}

func TestBaseConvert(t *testing.T) {
	_, err := KConv.BaseConvert("726746425", 10, 16)
	if err != nil {
		t.Error("BaseConvert fail")
		return
	}
	_, _ = KConv.BaseConvert("hello", 10, 16)
}

func BenchmarkBaseConvert(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.BaseConvert("726746425", 10, 16)
	}
}

func TestIp2Long(t *testing.T) {
	res := KConv.Ip2Long("127.0.0.1")
	if res == 0 {
		t.Error("Ip2Long fail")
		return
	}
	KConv.Ip2Long("1")
}

func BenchmarkIp2Long(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Ip2Long("127.0.0.1")
	}
}

func TestLong2Ip(t *testing.T) {
	res := KConv.Long2Ip(2130706433)
	if res != "127.0.0.1" {
		t.Error("Long2Ip fail")
		return
	}
}

func BenchmarkLong2Ip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Long2Ip(2130706433)
	}
}

func TestGettype(t *testing.T) {
	res1 := KConv.Gettype(1)
	res2 := KConv.Gettype("hello")
	res3 := KConv.Gettype(false)
	if res1 != "int" || res2 != "string" || res3 != "bool" {
		t.Error("Gettype fail")
		return
	}
}

func BenchmarkGettype(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Gettype("hello")
	}
}

func TestToStr(t *testing.T) {
	res1 := KConv.ToStr(1)
	res2 := KConv.ToStr(false)
	res3 := KConv.ToStr(UINT64_MAX)
	res4 := KConv.ToStr([]byte("hello"))
	if res1 != "1" || res2 != "false" || res3 != "18446744073709551615" || res4 != "hello" {
		t.Error("ToStr fail")
		return
	}
}

func BenchmarkToStr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToStr(UINT64_MAX)
	}
}

func TestToInt(t *testing.T) {
	var fn CallBack
	var tests = []struct {
		param    interface{}
		expected int
	}{
		{int(-1), -1},
		{int8(0), 0},
		{int16(1), 1},
		{int32(2), 2},
		{int64(3), 3},
		{uint(0), 0},
		{uint8(0), 0},
		{uint16(0), 0},
		{uint32(0), 0},
		{uint64(0), 0},
		{float32(0), 0},
		{float64(0), 0},
		{[]byte{}, 0},
		{"1", 1},
		{"2.1", 0},
		{"TRUE", 1},
		{true, 1},
		{false, 0},
		{fn, 0},
	}

	for _, test := range tests {
		actual := KConv.ToInt(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToBool(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToInt("123")
	}
}

func TestToFloat(t *testing.T) {
	res1 := KConv.ToFloat("")
	res2 := KConv.ToFloat(true)
	res3 := KConv.ToFloat(UINT64_MAX)
	res4 := KConv.ToFloat("123")
	if res1 != 0 || res2 != 1 || res3 < 1 || res4 != 123.0 {
		t.Error("ToFloat fail")
		return
	}
}

func BenchmarkToFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToFloat("123")
	}
}

func TestFloat64ToByte(t *testing.T) {
	var num float64 = 12345.6
	res := KConv.Float64ToByte(num)
	if len(res) == 0 {
		t.Error("Float64ToByte fail")
		return
	}
}

func BenchmarkFloat64ToByte(b *testing.B) {
	b.ResetTimer()
	var num float64 = 12345.6
	for i := 0; i < b.N; i++ {
		KConv.Float64ToByte(num)
	}
}

func TestByteToFloat64(t *testing.T) {
	bs := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	res := KConv.ByteToFloat64(bs)
	if res != 12345.6 {
		t.Error("ByteToFloat64 fail")
		return
	}
}

func BenchmarkByteToFloat64(b *testing.B) {
	b.ResetTimer()
	bs := []byte{205, 204, 204, 204, 204, 28, 200, 64}
	for i := 0; i < b.N; i++ {
		KConv.ByteToFloat64(bs)
	}
}

func TestInt64ToByte(t *testing.T) {
	var num int64 = 12345
	res := KConv.Int64ToByte(num)
	if len(res) == 0 {
		t.Error("Int64ToByte fail")
		return
	}
}

func BenchmarkInt64ToByte(b *testing.B) {
	b.ResetTimer()
	var num int64 = 12345
	for i := 0; i < b.N; i++ {
		KConv.Int64ToByte(num)
	}
}

func TestByteToInt64(t *testing.T) {
	bs := []byte{0, 0, 0, 0, 0, 0, 48, 57}
	res := KConv.ByteToInt64(bs)
	if res != 12345 {
		t.Error("ByteToFloat64 fail")
		return
	}
}

func BenchmarkByteToInt64(b *testing.B) {
	b.ResetTimer()
	bs := []byte{0, 0, 0, 0, 0, 0, 48, 57}
	for i := 0; i < b.N; i++ {
		KConv.ByteToInt64(bs)
	}
}

func TestByte2Hex(t *testing.T) {
	bs := []byte("hello")
	res := KConv.Byte2Hex(bs)
	if res != "68656c6c6f" {
		t.Error("Byte2Hex fail")
		return
	}
}

func BenchmarkByte2Hex(b *testing.B) {
	b.ResetTimer()
	bs := []byte("hello")
	for i := 0; i < b.N; i++ {
		KConv.Byte2Hex(bs)
	}
}

func TestHex2Byte(t *testing.T) {
	str := "68656c6c6f"
	res := KConv.Hex2Byte(str)
	if string(res) != "hello" {
		t.Error("Hex2Byte fail")
		return
	}
}

func BenchmarkHex2Byte(b *testing.B) {
	b.ResetTimer()
	str := "68656c6c6f"
	for i := 0; i < b.N; i++ {
		KConv.Hex2Byte(str)
	}
}

func TestGetPointerAddrInt(t *testing.T) {
	v1 := 1
	v2 := []byte("hello")

	res1 := KConv.GetPointerAddrInt(v1)
	res2 := KConv.GetPointerAddrInt(v2)
	if res1 <= 0 || res2 <= 0 {
		t.Error("GetPointerAddrInt fail")
		return
	}
}

func BenchmarkGetPointerAddrInt(b *testing.B) {
	b.ResetTimer()
	v := []byte("hello")
	for i := 0; i < b.N; i++ {
		KConv.GetPointerAddrInt(v)
	}
}

func TestToBool(t *testing.T) {
	//并行测试
	t.Parallel()

	var fn CallBack

	var tests = []struct {
		param    interface{}
		expected bool
	}{
		{int(-1), false},
		{int8(0), false},
		{int16(1), true},
		{int32(2), true},
		{int64(3), true},
		{uint(0), false},
		{uint8(0), false},
		{uint16(0), false},
		{uint32(0), false},
		{uint64(0), false},
		{float32(0), false},
		{float64(0), false},
		{[]byte{}, false},
		{"1", true},
		{"2.1", false},
		{"TRUE", true},
		{false, false},
		{fn, false},
	}

	for _, test := range tests {
		actual := KConv.ToBool(test.param)
		if actual != test.expected {
			t.Errorf("Expected ToBool(%q) to be %v, got %v", test.param, test.expected, actual)
			return
		}
	}
}

func BenchmarkToBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToBool(1)
	}
}
