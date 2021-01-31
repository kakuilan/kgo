package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var strSli = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

func TestArray_ArrayChunk(t *testing.T) {
	size := 3
	res := KArr.ArrayChunk(strSli, size)
	assert.Equal(t, 4, len(res))

	item := res[0]
	assert.Equal(t, size, len(item))

	KArr.ArrayChunk([]int{}, 1)
}

func TestArray_ArrayChunk_PanicSize(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, "[ArrayChunk]`size cannot be less than 1", r)
	}()
	KArr.ArrayChunk(strSli, 0)
}

func TestArray_ArrayChunk_PanicType(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, "[ArrayChunk]`arr type must be array or slice", r)
	}()
	KArr.ArrayChunk("hello", 2)
}

func BenchmarkArray_ArrayChunk(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayChunk(strSli, 3)
	}
}
