// +build windows

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_IsWindows(t *testing.T) {
	res := KOS.IsWindows()
	assert.True(t, res)
}

func BenchmarkOS_IsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsWindows()
	}
}

func TestOS_HomeDir(t *testing.T) {
	res, err := KOS.HomeDir()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_HomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestOS_Exec_Windows(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.Exec(tesCommand03)
	assert.Equal(t, ret, 0)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)
}

func BenchmarkOS_Exec_Windows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.Exec(tesCommand03)
	}
}

func TestOS_System_Windows(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.System(tesCommand03)
	assert.Equal(t, ret, 0)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)
}

func BenchmarkOS_System_Windows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.System(tesCommand03)
	}
}
