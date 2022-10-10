//go:build darwin
// +build darwin

package kgo

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net"
	"testing"
	"time"
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

func TestOS_Darwin_GetPidByPort(t *testing.T) {
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

func BenchmarkOS_Darwin_GetPidByPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetPidByPort(8899)
	}
}
