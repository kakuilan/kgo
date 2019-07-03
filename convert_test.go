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
