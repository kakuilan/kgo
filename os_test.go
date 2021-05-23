package kgo

import (
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestOS_Pwd(t *testing.T) {
	res := KOS.Pwd()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_Pwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestOS_Getcwd_Chdir(t *testing.T) {
	var ori, res string
	var err error

	ori, err = KOS.Getcwd()
	assert.Nil(t, err)
	assert.NotEmpty(t, ori)

	//切换目录
	err = KOS.Chdir(dirTdat)
	assert.Nil(t, err)

	res, err = KOS.Getcwd()
	assert.Nil(t, err)

	//返回原来目录
	err = KOS.Chdir(ori)
	assert.Nil(t, err)
	assert.Equal(t, KFile.AbsPath(res), KFile.AbsPath(dirTdat))
}

func BenchmarkOS_Getcwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Getcwd()
	}
}

func BenchmarkOS_Chdir(b *testing.B) {
	b.ResetTimer()
	dir := KOS.Pwd()
	for i := 0; i < b.N; i++ {
		_ = KOS.Chdir(dir)
	}
}

func TestOS_LocalIP(t *testing.T) {
	res, err := KOS.LocalIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_LocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.LocalIP()
	}
}

func TestOS_OutboundIP(t *testing.T) {
	res, err := KOS.OutboundIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
}

func BenchmarkOS_OutboundIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.OutboundIP()
	}
}

func TestOS_IsPublicIPv4(t *testing.T) {
	var res bool

	res = KOS.IsPublicIPv4(net.ParseIP(localIp))
	assert.False(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(lanIp))
	assert.False(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(googleIpv4))
	assert.True(t, res)

	res = KOS.IsPublicIPv4(net.ParseIP(googleIpv6))
	assert.False(t, res)
}

func BenchmarkOS_IsPublicIPv4(b *testing.B) {
	b.ResetTimer()
	ip := net.ParseIP(publicIp1)
	for i := 0; i < b.N; i++ {
		KOS.IsPublicIPv4(ip)
	}
}

func TestOS_GetIPs(t *testing.T) {
	res := KOS.GetIPs()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_GetIPs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetIPs()
	}
}
