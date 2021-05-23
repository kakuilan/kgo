package kgo

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"os/user"
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

func TestOS_IsPrivateIp(t *testing.T) {
	var res bool
	var err error

	res, err = KOS.IsPrivateIp(lanIp)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = KOS.IsPrivateIp(publicIp2)
	assert.Nil(t, err)
	assert.False(t, res)

	//非IP
	res, err = KOS.IsPrivateIp(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_IsPrivateIp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.IsPrivateIp(lanIp)
	}
}

func TestOS_IsPublicIP(t *testing.T) {
	var res bool
	var err error

	res, err = KOS.IsPublicIP(localIp)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = KOS.IsPublicIP(lanIp)
	assert.Nil(t, err)
	assert.False(t, res)

	res, err = KOS.IsPublicIP(googleIpv4)
	assert.Nil(t, err)
	assert.True(t, res)

	res, err = KOS.IsPublicIP(googleIpv6)
	assert.Nil(t, err)
	assert.True(t, res)

	//非IP
	res, err = KOS.IsPublicIP(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_IsPublicIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.IsPublicIP(publicIp1)
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

func TestOS_GetMacAddrs(t *testing.T) {
	res := KOS.GetMacAddrs()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_GetMacAddrs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetMacAddrs()
	}
}

func TestOS_Hostname_GetIpByHostname(t *testing.T) {
	var res string
	var err error

	res, err = KOS.Hostname()
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpByHostname(res)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpByHostname(tesIp2)
	assert.Empty(t, res)

	res, err = KOS.GetIpByHostname(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_Hostname(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Hostname()
	}
}

func BenchmarkOS_GetIpByHostname(b *testing.B) {
	b.ResetTimer()
	host, _ := KOS.Hostname()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpByHostname(host)
	}
}

func TestOS_GetIpsByDomain(t *testing.T) {
	var res []string
	var err error

	res, err = KOS.GetIpsByDomain(tesDomain30)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetIpsByDomain(strHello)
	assert.NotNil(t, err)
}

func BenchmarkOS_GetIpsByDomain(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpsByDomain(tesDomain30)
	}
}

func TestOS_GetHostByIp(t *testing.T) {
	var res string
	var err error

	res, err = KOS.GetHostByIp(localIp)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)

	res, err = KOS.GetHostByIp(strHello)
	assert.NotNil(t, err)
	assert.Empty(t, res)
}

func BenchmarkOS_GetHostByIp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetHostByIp(localIp)
	}
}

func TestOS_Setenv_Getenv_Unsetenv(t *testing.T) {
	var res string
	var err error

	err = KOS.Setenv(helloEngICase, helloOther)
	assert.Nil(t, err)

	res = KOS.Getenv(helloEngICase)
	assert.Equal(t, res, helloOther)

	err = KOS.Unsetenv(helloEngICase)
	assert.Nil(t, err)

	res = KOS.Getenv(helloEngICase, helloOther2)
	assert.Equal(t, res, helloOther2)
}

func BenchmarkOS_Setenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KOS.Setenv(helloEngICase, helloOther)
	}
}

func BenchmarkOS_Getenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Getenv(helloEngICase)
	}
}

func BenchmarkOS_Unsetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KOS.Unsetenv(helloEngICase)
	}
}

func TestOS_GetEndian_IsLittleEndian(t *testing.T) {
	res := KOS.GetEndian()
	chk := KOS.IsLittleEndian()
	if chk {
		assert.Equal(t, res, binary.LittleEndian)
	} else {
		assert.Equal(t, res, binary.BigEndian)
	}
}

func BenchmarkOS_GetEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetEndian()
	}
}

func BenchmarkOS_IsLittleEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLittleEndian()
	}
}

func TestOS_Chmod_Chown(t *testing.T) {
	var res bool

	res = KOS.Chmod(dirDoc, 0777)
	assert.True(t, res)

	usr, _ := user.Current()
	uid := toInt(usr.Uid)
	guid := toInt(usr.Gid)
	res = KOS.Chown(dirDoc, uid, guid)
	assert.True(t, res)
}

func BenchmarkOS_Chmod(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Chmod(dirDoc, 0777)
	}
}

func BenchmarkOS_Chown(b *testing.B) {
	b.ResetTimer()
	usr, _ := user.Current()
	uid := toInt(usr.Uid)
	guid := toInt(usr.Gid)
	for i := 0; i < b.N; i++ {
		KOS.Chown(dirDoc, uid, guid)
	}
}

func TestOS_GetTempDir(t *testing.T) {
	res := KOS.GetTempDir()
	assert.NotEmpty(t, res)
}

func BenchmarkOS_GetTempDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetTempDir()
	}
}
