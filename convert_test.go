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

func TestStrictStr2Int(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := KConv.StrictStr2Int("abc123", 8, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "int8" {
		t.Error("StrictStr2Int fail")
		return
	}
}

func TestStr2Int(t *testing.T) {
	res := KConv.Str2Int("123")
	if fmt.Sprint(reflect.TypeOf(res)) != "int" {
		t.Error("Str2Int fail")
		return
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

func TestStrictStr2Float(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	res := KConv.StrictStr2Float("abc123", 32, true)
	if fmt.Sprint(reflect.TypeOf(res)) != "float32" {
		t.Error("StrictStr2Float fail")
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

func TestInt2Bool(t *testing.T) {
	var (
		it1 int   = -1
		it2 int8  = 0
		it3 int16 = 1
		it4 int32 = 2
		it5 int64 = 3

		ui1 uint   = 0
		ui2 uint8  = 1
		ui3 uint16 = 2
		ui4 uint32 = 3
		ui5 uint64 = 4

		it6 string = "1"
	)

	res1 := KConv.Int2Bool(it1)
	res2 := KConv.Int2Bool(it2)
	res3 := KConv.Int2Bool(it3)
	res4 := KConv.Int2Bool(it4)
	res5 := KConv.Int2Bool(it5)

	res6 := KConv.Int2Bool(ui1)
	res7 := KConv.Int2Bool(ui2)
	res8 := KConv.Int2Bool(ui3)
	res9 := KConv.Int2Bool(ui4)
	res10 := KConv.Int2Bool(ui5)
	res11 := KConv.Int2Bool(it6)

	if res1 || res2 || res6 || res11 {
		t.Error("Int2Bool fail")
		return
	}

	if !(res3 && res4 && res5 && res7 && res8 && res9 && res10) {
		t.Error("Int2Bool fail")
		return
	}

}

func BenchmarkInt2Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int2Bool(1)
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

func TestBytesSlice2Str(t *testing.T) {
	sli := []byte("hello world!")
	res := KConv.BytesSlice2Str(sli)
	if fmt.Sprint(reflect.TypeOf(res)) != "string" {
		t.Error("BytesSlice2Str fail")
		return
	}
}

func BenchmarkBytesSlice2Str(b *testing.B) {
	b.ResetTimer()
	sli := []byte(`hello world!
// Convert different types to byte slice using types and functions in unsafe and reflect package. 
// It has higher performance, but notice that it may be not safe when garbage collection happens.
// Use it when you need to temporary convert a long string to a byte slice and won't keep it for long time.
`)
	for i := 0; i < b.N; i++ {
		KConv.BytesSlice2Str(sli)
	}
}

func TestDecbin(t *testing.T) {
	var num int64 = 8
	res := KConv.Decbin(num)
	if res != "1000" {
		t.Error("Decbin fail")
		return
	}
}

func BenchmarkDecbin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Decbin(10)
	}
}

func TestBindec(t *testing.T) {
	res, err := KConv.Bindec("1000")
	if err != nil || res != 8 {
		t.Error("Bindec fail")
		return
	}
	_, _ = KConv.Bindec("hello")
}

func BenchmarkBindec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bindec("1000")
	}
}
