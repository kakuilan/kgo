// +build linux darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOS_Unix_NotWindows(t *testing.T) {
	res := KOS.IsWindows()
	assert.False(t, res)
}

func TestOS_Unix_HomeDir(t *testing.T) {
	res, err := KOS.HomeDir()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Unix_HomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestOS_Unix_Exec(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.Exec(tesCommand01)
	assert.Equal(t, ret, 0)
	assert.NotEmpty(t, res)
	assert.Empty(t, err)

	//错误的命令
	ret, res, err = KOS.Exec(tesCommand02)
	assert.Equal(t, ret, 1)
	assert.Empty(t, res)
	assert.NotEmpty(t, err)
}

func BenchmarkOS_Unix_Exec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.Exec(tesCommand01)
	}
}

func TestOS_Unix_System(t *testing.T) {
	var ret int
	var res []byte
	var err []byte

	ret, res, err = KOS.System(tesCommand01)
	assert.Equal(t, ret, 0)
	assert.Empty(t, err)
	assert.NotEmpty(t, res)

	//错误的命令
	ret, res, err = KOS.System(tesCommand02)
	assert.Equal(t, ret, 1)
	assert.NotEmpty(t, err)
	assert.Empty(t, res)
}

func BenchmarkOS_Unix_System(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.System(tesCommand01)
	}
}

func TestOS_Unix_IsProcessExists(t *testing.T) {
	var res bool

	pid := os.Getpid()
	res = KOS.IsProcessExists(pid)
	assert.True(t, res)

	res = KOS.IsProcessExists(-1)
	assert.False(t, res)
}

func BenchmarkOS_Unix_IsProcessExists(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.IsProcessExists(pid)
	}
}
