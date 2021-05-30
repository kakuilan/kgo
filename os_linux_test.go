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

func TestOS_Linux_CpuUsage(t *testing.T) {
	var user, idle, total uint64
	user, idle, total = KOS.CpuUsage()
	assert.Greater(t, int(user), 1)
	assert.Greater(t, int(idle), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Linux_CpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.CpuUsage()
	}
}

func TestOS_Linux_DiskUsage(t *testing.T) {
	var used, free, total uint64
	used, free, total = KOS.DiskUsage("/")
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Linux_DiskUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.DiskUsage("/")
	}
}

func TestOS_Linux_Uptime(t *testing.T) {
	res, err := KOS.Uptime()
	assert.Greater(t, int(res), 1)
	assert.Nil(t, err)
}

func BenchmarkOS_Linux_Uptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Uptime()
	}
}

func TestOS_Linux_GetBiosInfo(t *testing.T) {
	res := KOS.GetBiosInfo()
	assert.NotNil(t, res)
}

func BenchmarkOS_Linux_GetBiosInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBiosInfo()
	}
}

func TestOS_Linux_GetBoardInfo(t *testing.T) {
	res := KOS.GetBoardInfo()
	//dumpPrint("------------GetBoardInfo", res)
	assert.NotNil(t, res)
}

func BenchmarkOS_Linux_GetBoardInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBoardInfo()
	}
}
