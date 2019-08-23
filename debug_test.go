package kgo

import (
	"strings"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	res1 := KDbug.GetFuncName(nil, false)
	res2 := KDbug.GetFuncName(nil, true)
	res3 := KDbug.GetFuncName(KArr.ArrayDiff, true) //ArrayDiff-fm
	res4 := KDbug.GetFuncName(KArr.ArrayDiff, true) //ArrayDiff-fm

	if !strings.Contains(res1, "TestGetFuncName") || res2 != "TestGetFuncName" || !strings.Contains(res3, "ArrayDiff") || !strings.Contains(res4, "ArrayDiff") {
		t.Error("GetFuncName fail")
		return
	}
}

func BenchmarkGetFuncName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncName(nil, true)
	}
}

func TestGetFuncLine(t *testing.T) {
	res := KDbug.GetFuncLine()
	if res <= 0 {
		t.Error("GetFuncLine fail")
		return
	}
}

func BenchmarkGetFuncLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncLine()
	}
}
