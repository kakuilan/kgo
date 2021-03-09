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

func TestDebug_GetCallFuncName(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res string

	res = KDbug.GetCallFuncName(nil, false)
	assert.Contains(t, res, "TestDebug_GetCallFuncName")

	res = KDbug.GetCallFuncName(nil, true)
	assert.Equal(t, "TestDebug_GetCallFuncName", res)

	res = KDbug.GetCallFuncName("", false)
	assert.Empty(t, res)

	res = KDbug.GetCallFuncName(KArr.ArrayRand, false)
	assert.Contains(t, res, "ArrayRand")

	res = KDbug.GetCallFuncName(KArr.ArrayRand, true)
	assert.Equal(t, "ArrayRand-fm", res)

	//未实现的方法
	KDbug.GetCallFuncName(itfObj.noRealize, false)
}

func BenchmarkDebug_GetCallFuncName(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KDbug.GetCallFuncName(KArr.ArrayRand, false)
	}
}
