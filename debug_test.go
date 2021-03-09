package kgo

import (
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

	//KDbug.DumpStacks()
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
