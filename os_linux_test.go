// +build linux

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_IsLinux(t *testing.T) {
	res := KOS.IsLinux()
	assert.True(t, res)
}

func BenchmarkOS_IsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLinux()
	}
}
