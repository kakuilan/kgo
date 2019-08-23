package kgo

import (
	"strings"
	"testing"
)

func TestGetFunctionName(t *testing.T) {
	res1 := KDbug.GetFunctionName(nil, false)
	res2 := KDbug.GetFunctionName(nil, true)
	res3 := KDbug.GetFunctionName(KArr.ArrayDiff, true) //ArrayDiff-fm
	res4 := KDbug.GetFunctionName(KArr.ArrayDiff, true) //ArrayDiff-fm

	if !strings.Contains(res1, "TestGetFunctionName") || res2 != "TestGetFunctionName" || !strings.Contains(res3, "ArrayDiff") || !strings.Contains(res4, "ArrayDiff") {
		t.Error("GetFunctionName fail")
		return
	}
}

func BenchmarkGetFunctionName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFunctionName(nil, true)
	}
}
