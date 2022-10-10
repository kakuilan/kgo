//go:build linux
// +build linux

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
	assert.NotNil(t, res)
}

func BenchmarkOS_Linux_GetBoardInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBoardInfo()
	}
}

func TestOS_Linux_GetCpuInfo(t *testing.T) {
	res := KOS.GetCpuInfo()
	assert.NotNil(t, res)
	assert.NotEmpty(t, res.Vendor)
	assert.NotEmpty(t, res.Model)
}

func BenchmarkOS_Linux_GetCpuInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetCpuInfo()
	}
}

func TestOS_Linux_GetProcessExecPath(t *testing.T) {
	var res string

	pid := os.Getpid()
	res = KOS.GetProcessExecPath(pid)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Linux_GetProcessExecPath(b *testing.B) {
	b.ResetTimer()
	pid := os.Getpid()
	for i := 0; i < b.N; i++ {
		KOS.GetProcessExecPath(pid)
	}
}

func TestOS_Linux_GetPidByPort(t *testing.T) {
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

func BenchmarkOS_Linux_GetPidByPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetPidByPort(8899)
	}
}
