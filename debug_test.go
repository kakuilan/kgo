package kgo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDebug_DumpPrint(t *testing.T) {
	defer func() {
		r := recover()
		assert.Empty(t, r)
	}()

	KDbug.DumpPrint(Version)
}

func TestDebug_DumpStacks(t *testing.T) {
	defer func() {
		r := recover()
		assert.Empty(t, r)
	}()

	KDbug.DumpStacks()
}

func TestDebug_GetCallName(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res string

	res = KDbug.GetCallName(nil, false)
	assert.Contains(t, res, "TestDebug_GetCallName")

	res = KDbug.GetCallName(nil, true)
	assert.Equal(t, "TestDebug_GetCallName", res)

	res = KDbug.GetCallName("", false)
	assert.Empty(t, res)

	res = KDbug.GetCallName(KArr.ArrayRand, false)
	assert.Contains(t, res, "ArrayRand")

	res = KDbug.GetCallName(KArr.ArrayRand, true)
	assert.Equal(t, "ArrayRand-fm", res)

	//未实现的方法
	KDbug.GetCallName(itfObj.noRealize, false)
}

func BenchmarkDebug_GetCallName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallName(KArr.ArrayRand, false)
	}
}

func TestDebug_GetCallFile(t *testing.T) {
	res := KDbug.GetCallFile()
	assert.NotEmpty(t, res)
}

func BenchmarkDebug_GetCallFile(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallFile()
	}
}

func TestDebug_GetCallDir(t *testing.T) {
	res := KDbug.GetCallDir()
	assert.NotEmpty(t, res)
}

func BenchmarkDebug_GetCallDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallDir()
	}
}

func TestDebug_GetCallLine(t *testing.T) {
	res := KDbug.GetCallLine()
	assert.Greater(t, res, 1)
}

func BenchmarkDebug_GetCallLine(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallLine()
	}
}

func TestDebug_GetCallPackage(t *testing.T) {
	var res string

	res = KDbug.GetCallPackage()
	assert.Equal(t, "kgo", res)

	res = KDbug.GetCallPackage(KDbug.GetCallFile())
	assert.Equal(t, "kgo", res)

	res = KDbug.GetCallPackage(strHello)
	assert.Empty(t, res)
}

func BenchmarkDebug_GetCallPackage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallPackage()
	}
}

func TestDebug_HasMethod(t *testing.T) {
	var tests = []struct {
		input    interface{}
		method   string
		expected bool
	}{
		{intSpeedLight, "", false},
		{strHello, "toString", false},
		{KArr, "InIntSlice", true},
		{&KArr, "InInt64Slice", true},
		{&KArr, strHello, false},
	}
	for _, test := range tests {
		actual := KDbug.HasMethod(test.input, test.method)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkDebug_HasMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.HasMethod(KArr, "SliceFill")
	}
}

func TestDebug_GetMethod(t *testing.T) {
	var res interface{}

	res = KDbug.GetMethod(intSpeedLight, "")
	assert.Nil(t, res)

	res = KDbug.GetMethod(&KArr, "InIntSlice")
	assert.NotNil(t, res)

	res = KDbug.GetMethod(KArr, "InIntSlice")
	assert.NotNil(t, res)

	res = KDbug.GetMethod(KArr, strHello)
	assert.Nil(t, res)
}

func BenchmarkDebug_GetMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetMethod(&KArr, "InIntSlice")
	}
}

func TestDebug_CallMethod(t *testing.T) {
	var res interface{}
	var err error

	//调用不存在的方法
	res, err = KDbug.CallMethod(KArr, strHello)
	assert.NotNil(t, err)

	//有参数调用
	res, err = KDbug.CallMethod(&KConv, "BaseConvert", "123456", 10, 16)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	//无参数调用
	res, err = KDbug.CallMethod(&KOS, "Getcwd")
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkDebug_CallMethod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KDbug.CallMethod(&KOS, "Getcwd")
	}
}

func TestDebug_GetFuncNames(t *testing.T) {
	var res []string

	//空字符串
	res = KDbug.GetFuncNames("")
	assert.Empty(t, res)

	//空变量
	res = KDbug.GetFuncNames(nil)
	assert.Empty(t, res)

	//非指针变量
	res = KDbug.GetFuncNames(KStr)
	assert.NotEmpty(t, res)

	//指针变量
	res = KDbug.GetFuncNames(&KConv)
	assert.NotEmpty(t, res)

	n1 := KDbug.GetFuncNames(KFile)
	n2 := KDbug.GetFuncNames(KStr)
	n3 := KDbug.GetFuncNames(KNum)
	n4 := KDbug.GetFuncNames(KArr)
	n5 := KDbug.GetFuncNames(KTime)
	n6 := KDbug.GetFuncNames(KConv)
	n7 := KDbug.GetFuncNames(KOS)
	n8 := KDbug.GetFuncNames(KEncr)
	n9 := KDbug.GetFuncNames(KDbug)
	funTotal := len(n1) + len(n2) + len(n3) + len(n4) + len(n5) + len(n6) + len(n7) + len(n8) + len(n9)
	dumpPrint("the package function total:", funTotal)
}

func BenchmarkDebug_GetFuncNames(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetFuncNames(&KConv)
	}
}

func TestDebug_WrapError(t *testing.T) {
	var err = errors.New("a test error")

	err1 := KDbug.WrapError(nil)
	assert.NotNil(t, err1)
	assert.Contains(t, err1.Error(), "parameter error")

	err2 := KDbug.WrapError(err)
	assert.NotNil(t, err2)
	assert.Equal(t, err2, err)

	err3 := KDbug.WrapError(err, helloEng, utf8Hello)
	assert.NotNil(t, err3)
	assert.Contains(t, err3.Error(), helloEng)
	assert.Contains(t, err3.Error(), err.Error())
}

func BenchmarkDebug_WrapError(b *testing.B) {
	var err = errors.New("a test error")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KDbug.WrapError(err, helloEng, utf8Hello)
	}
}

func TestDebug_Stacks(t *testing.T) {
	var res1, res2, res3 []byte
	var len1, len2, len3 int

	res1 = KDbug.Stacks(-1)
	len1 = len(res1)
	assert.Greater(t, len1, 1)

	res2 = KDbug.Stacks(0)
	len2 = len(res2)
	assert.Greater(t, len1, len2)

	res3 = KDbug.Stacks(1)
	len3 = len(res3)
	assert.Greater(t, len2, len3)
}

func BenchmarkDebug_Stacks(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KDbug.Stacks(0)
	}
}
