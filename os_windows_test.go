//go:build windows
// +build windows

package kgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"os"
	"testing"
	"time"
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

func TestOS_Windows_CpuUsage(t *testing.T) {
	var user, idle, total uint64
	user, idle, total = KOS.CpuUsage()
	assert.Greater(t, int(user), 1)
	assert.Greater(t, int(idle), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Windows_CpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.CpuUsage()
	}
}

func TestOS_Windows_DiskUsage(t *testing.T) {
	var used, free, total uint64
	used, free, total = KOS.DiskUsage("C:")
	assert.Greater(t, int(used), 1)
	assert.Greater(t, int(free), 1)
	assert.Greater(t, int(total), 1)
}

func BenchmarkOS_Windows_DiskUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.DiskUsage("C:")
	}
}

func TestOS_Windows_Uptime(t *testing.T) {
	res, err := KOS.Uptime()
	assert.Greater(t, int(res), 1)
	assert.Nil(t, err)
}

func BenchmarkOS_Windows_Uptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Uptime()
	}
}

func TestOS_Windows_GetBiosInfo(t *testing.T) {
	res := KOS.GetBiosInfo()
	assert.NotNil(t, res)
}

func BenchmarkOS_Windows_GetBiosInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBiosInfo()
	}
}

func TestOS_Windows_GetBoardInfo(t *testing.T) {
	res := KOS.GetBoardInfo()
	assert.NotNil(t, res)
}

func BenchmarkOS_Windows_GetBoardInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBoardInfo()
	}
}

func TestOS_Windows_GetCpuInfo(t *testing.T) {
	res := KOS.GetCpuInfo()
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Vendor)
	assert.NotEmpty(t, res.Model)
}

func BenchmarkOS_Windows_GetCpuInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetCpuInfo()
	}
}

func TestOS_Windows_IsProcessExists(t *testing.T) {
	var res bool

	pid := os.Getpid()
	res = KOS.IsProcessExists(pid)
	assert.True(t, res)

	res = KOS.IsProcessExists(5)
	assert.False(t, res)

	res = KOS.IsProcessExists(-1)
	assert.False(t, res)
}

func BenchmarkOS_Windows_IsProcessExists(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.IsProcessExists(pid)
	}
}

func TestOS_Windows_GetProcessExecPath(t *testing.T) {
	var res string

	pid := os.Getpid()
	res = KOS.GetProcessExecPath(pid)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Windows_GetProcessExecPath(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.GetProcessExecPath(pid)
	}
}

func TestOS_Windows_GetPidByPort(t *testing.T) {
	time.AfterFunc(time.Millisecond*250, func() {
		res := KOS.GetPidByPort(8899)
		assert.Greater(t, res, 1)

		KOS.GetPidByPort(80)
	})

	//发送消息
	time.AfterFunc(time.Millisecond*500, func() {
		conn, err := net.Dial("tcp", ":8899")
		assert.Nil(t, err)

		defer func() {
			_ = conn.Close()
		}()

		_, err = fmt.Fprintf(conn, helloEng)
		assert.Nil(t, err)
	})

	//开启监听端口
	l, err := net.Listen("tcp", ":8899")
	assert.Nil(t, err)
	defer func() {
		_ = l.Close()
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer func() {
			_ = conn.Close()
		}()

		//接收
		buf, err := io.ReadAll(conn)
		assert.Nil(t, err)

		msg := string(buf[:])
		assert.Equal(t, msg, helloEng)
		return
	}
}

func BenchmarkOS_Windows_GetPidByPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetPidByPort(8899)
	}
}
