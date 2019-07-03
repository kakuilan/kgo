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
