// +build windows

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_Windows_IsWindows(t *testing.T) {
	res := KOS.IsWindows()
	assert.True(t, res)
}

func BenchmarkOS_Windows_IsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsWindows()
	}
}

func TestOS_Windows_HomeDir(t *testing.T) {
	res, err := KOS.HomeDir()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Windows_HomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestOS_Windows_Exec(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.Exec(tesCommand03)
	assert.Equal(t, ret, 0)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)
}

func BenchmarkOS_Windows_Exec_Windows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.Exec(tesCommand03)
	}
}

func TestOS_Windows_System(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.System(tesCommand03)
	assert.Equal(t, ret, 0)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)
}

func BenchmarkOS_Windows_System(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.System(tesCommand03)
	}
}

func TestOS_Windows_MemoryUsage(t *testing.T) {
	var used, free, total uint64

	used, free, total = KOS.MemoryUsage(true)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Windows_MemoryUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(true)
	}
}
