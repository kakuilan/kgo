// +build darwin
// +build !cgo

package kgo

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOS_Darwin_Nocgo_CpuUsage(t *testing.T) {
	var user, idle, total uint64
	user, idle, total = KOS.CpuUsage()
	assert.Equal(t, int(user), 0)
	assert.Equal(t, int(idle), 0)
	assert.Equal(t, int(total), 0)
}

func BenchmarkOS_Darwin_Nocgo_CpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.CpuUsage()
	}
}

func TestOS_Darwin_Nocgo_GetProcessExecPath(t *testing.T) {
	var res string

	pid := os.Getpid()
	res = KOS.GetProcessExecPath(pid)
	dumpPrint("-------------TestOS_Darwin_Nocgo_GetProcessExecPath res:", pid, res)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Darwin_Nocgo_GetProcessExecPath(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.GetProcessExecPath(pid)
	}
}
