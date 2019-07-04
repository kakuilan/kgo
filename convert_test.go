package gohelper

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
		t.Error("Str2Int fail")
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
	println(UINT_MAX)
	println(UINT_MIN)
	println(INT_MAX)
	println(INT_MIN)
	if res > 9223372036854775807 {
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
