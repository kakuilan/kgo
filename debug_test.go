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

func TestHasMethod(t *testing.T) {
	var test = &KOS

	chk1 := KDbug.HasMethod(test, "IsLinux")
	chk2 := KDbug.HasMethod(test, "Hello")
	if !chk1 || chk2 {
		t.Error("HasMethod fail")
		return
	}
}

func BenchmarkHasMethod(b *testing.B) {
	b.ResetTimer()
	var test = &KOS
	for i := 0; i < b.N; i++ {
		KDbug.HasMethod(test, "IsLinux")
	}
}

func TestGetMethod(t *testing.T) {
	var test = &KOS

	fun1 := KDbug.GetMethod(test, "GoMemory")
	fun2 := KDbug.GetMethod(test, "Hello")

	if fun1 == nil || fun2 != nil {
		t.Error("GetMethod fail")
		return
	}
}

func BenchmarkGetMethod(b *testing.B) {
	b.ResetTimer()
	var test = &KOS
	for i := 0; i < b.N; i++ {
		KDbug.GetMethod(test, "GoMemory")
	}
}

func TestCallMethod(t *testing.T) {
	var test = &KOS

	//无参数调用
	res1, err1 := KDbug.CallMethod(test, "GoMemory")
	if res1 == nil || err1 != nil {
		t.Error("CallMethod fail")
		return
	}

	//调用不存在的方法
	res2, err2 := KDbug.CallMethod(test, "Hello")
	if res2 != nil || err2 == nil {
		t.Error("CallMethod fail")
		return
	}

	//有参数调用
	var conv = &KConv
	res3, err3 := KDbug.CallMethod(conv, "BaseConvert", "123456", 10, 16)
	//结果 [1e240 <nil>]
	if len(res3) != 2 || res3[0] != "1e240" || res3[1] != nil || err3 != nil {
		t.Error("CallMethod fail")
		return
	}
}

func BenchmarkCallMethod(b *testing.B) {
	b.ResetTimer()
	var test = &KOS
	for i := 0; i < b.N; i++ {
		KDbug.GetMethod(test, "GoMemory")
	}
}

func TestValidFunc(t *testing.T) {
	var err error
	var conv = &KConv
	method := KDbug.GetMethod(conv, "BaseConvert")

	//不存在的方法
	_, _, err = ValidFunc("test", "echo")
	if err == nil {
		t.Error("ValidFunc fail")
		return
	}

	//参数数量不足
	_, _, err = ValidFunc(method, "12345")
	if err == nil {
		t.Error("ValidFunc fail")
		return
	}

	//参数类型错误
	_, _, err = ValidFunc(method, 0, "12345", "10", 16)
	if err == nil {
		t.Error("ValidFunc fail")
		return
	}
}

func TestCallFunc(t *testing.T) {
	var err error
	var conv = &KConv
	method := KDbug.GetMethod(conv, "BaseConvert")

	_, err = CallFunc(method, 0, "12345", "10", 16)
	if err != nil {
		println(err.Error())
	}
}
