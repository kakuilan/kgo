package kgo

import (
	"strings"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	res1 := KDbug.GetFuncName(nil, false)
	res2 := KDbug.GetFuncName(nil, true)
	res3 := KDbug.GetFuncName(KArr.ArrayDiff, false) // ...ArrayDiff-fm
	res4 := KDbug.GetFuncName(KArr.ArrayDiff, true)  // ArrayDiff-fm

	if !strings.Contains(res1, "TestGetFuncName") || res2 != "TestGetFuncName" || !strings.Contains(res3, "ArrayDiff") || !strings.HasPrefix(res4, "ArrayDiff") {
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

func TestGetFuncFileDir(t *testing.T) {
	res1 := KDbug.GetFuncFile()
	res2 := KDbug.GetFuncDir()
	if res1 == "" {
		t.Error("GetFuncFile fail")
		return
	} else if res2 == "" {
		t.Error("GetFuncDir fail")
		return
	}
}

func BenchmarkGetFuncFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncFile()
	}
}

func BenchmarkGetFuncDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncDir()
	}
}

func TestGetFuncPackage(t *testing.T) {
	res1 := KDbug.GetFuncPackage()
	res2 := KDbug.GetFuncPackage(KDbug.GetFuncFile())
	res3 := KDbug.GetFuncPackage("test")

	if res1 != "kgo" || res1 != res2 || res3 != "" {
		t.Error("GetFuncPackage fail")
		return
	}
}

func BenchmarkGetFuncPackage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncPackage()
	}
}

func TestDumpStacks(t *testing.T) {
	KDbug.DumpStacks()
}

func BenchmarkDumpStacks(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < 100; i++ {
		KDbug.DumpStacks()
	}
}
