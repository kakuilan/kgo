// +build darwin

package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOS_Darwin_IsMac(t *testing.T) {
	res := KOS.IsMac()
	assert.True(t, res)
}

func BenchmarkOS_Darwin_IsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}

func TestOS_Darwin_MemoryUsage(t *testing.T) {
	var used, free, total uint64

	used, free, total = KOS.MemoryUsage(true)
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Darwin_MemoryUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(true)
	}
}

func TestOS_Darwin_CpuUsage(t *testing.T) {
	var user, idle, total uint64
	user, idle, total = KOS.CpuUsage()
	assert.GreaterOrEqual(t, int(user), 0)
	assert.GreaterOrEqual(t, int(idle), 0)
	assert.GreaterOrEqual(t, int(total), 0)
}

func BenchmarkOS_Darwin_CpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.CpuUsage()
	}
}

func TestOS_Darwin_DiskUsage(t *testing.T) {
	var used, free, total uint64
	used, free, total = KOS.DiskUsage("/")
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Darwin_DiskUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.DiskUsage("/")
	}
}

func TestOS_Darwin_Uptime(t *testing.T) {
	res, err := KOS.Uptime()
	assert.Greater(t, int(res), 1)
	assert.Nil(t, err)
}

func BenchmarkOS_Darwin_Uptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Uptime()
	}
}

func TestOS_Darwin_GetBiosInfo(t *testing.T) {
	res := KOS.GetBiosInfo()
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Version)
}

func BenchmarkOS_Darwin_GetBiosInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBiosInfo()
	}
}

func TestOS_Darwin_GetBoardInfo(t *testing.T) {
	res := KOS.GetBoardInfo()
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Name)
	assert.NotEmpty(t, res.Version)
	assert.NotEmpty(t, res.Serial)
}

func BenchmarkOS_Darwin_GetBoardInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBoardInfo()
	}
}

func TestOS_Darwin_GetCpuInfo(t *testing.T) {
	res := KOS.GetCpuInfo()
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Vendor)
	assert.NotEmpty(t, res.Model)
}

func BenchmarkOS_Darwin_GetCpuInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetCpuInfo()
	}
}

func TestOS_Darwin_GetProcessExecPath(t *testing.T) {
	var res string

	pid := os.Getpid()
	res = KOS.GetProcessExecPath(pid)
	dumpPrint("-------------TestOS_Darwin_GetProcessExecPath res:", pid, res)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Darwin_GetProcessExecPath(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.GetProcessExecPath(pid)
	}
}
