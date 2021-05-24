// +build linux darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_NotWindows(t *testing.T) {
	res := KOS.IsWindows()
	assert.False(t, res)
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

func TestOS_Exec_Unix(t *testing.T) {
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

func BenchmarkOS_Exec_Unix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.Exec(tesCommand01)
	}
}

func TestOS_System_Unix(t *testing.T) {
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
}

func BenchmarkOS_System_Unix(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.System(tesCommand01)
	}
}
