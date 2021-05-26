// +build linux

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_Linux_IsLinux(t *testing.T) {
	res := KOS.IsLinux()
	assert.True(t, res)
}

func BenchmarkOS_Linux_IsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLinux()
	}
}

func TestOS_Linux_MemoryUsage(t *testing.T) {
	var used, free, total uint64

	// 虚拟内存
	used, free, total = KOS.MemoryUsage(true)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)

	// 真实物理内存
	used, free, total = KOS.MemoryUsage(false)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Linux_MemoryUsage_Virtual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(true)
	}
}

func BenchmarkOS_Linux_MemoryUsage_Physic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(false)
	}
}
