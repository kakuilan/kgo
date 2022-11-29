package kgo

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"os/user"
	"strings"
	"testing"
	"time"
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

	KFile.Touch(chownfile, 128)
	res = KOS.Chmod(chownfile, 0777)
	assert.True(t, res)

	usr, _ := user.Current()
	uid := toInt(usr.Uid)
	guid := toInt(usr.Gid)
	res = KOS.Chown(chownfile, uid, guid)
	if KOS.IsWindows() {
		assert.False(t, res)
	} else {
		assert.True(t, res)
	}
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

func TestOS_ClientIp(t *testing.T) {
	// Create type and function for testing
	type testIP struct {
		name     string
		request  *http.Request
		expected string
	}

	newRequest := func(remoteAddr, xRealIP string, xForwardedFor ...string) *http.Request {
		h := http.Header{}
		h.Set("X-Real-IP", xRealIP)
		for _, address := range xForwardedFor {
			h.Set("X-Forwarded-For", address)
		}

		return &http.Request{
			RemoteAddr: remoteAddr,
			Header:     h,
		}
	}

	// Create test data
	testData := []testIP{
		{
			name:     "No header,no port",
			request:  newRequest(publicIp1, ""),
			expected: publicIp1,
		}, {
			name:     "No header,has port",
			request:  newRequest(tesIp8, ""),
			expected: tesIp8,
		}, {
			name:     "Has X-Forwarded-For",
			request:  newRequest("", "", publicIp1),
			expected: publicIp1,
		}, {
			name:     "Has multiple X-Forwarded-For",
			request:  newRequest("", "", localIp, publicIp1, publicIp2),
			expected: publicIp2,
		}, {
			name:     "Has X-Real-IP",
			request:  newRequest("", publicIp1),
			expected: publicIp1,
		}, {
			name:     "Local ip",
			request:  newRequest("", tesIp2),
			expected: tesIp2,
		},
	}

	// Run test
	var actual string
	for _, v := range testData {
		actual = KOS.ClientIp(v.request)
		if v.expected == "::1" {
			assert.Equal(t, actual, localIp)
		} else {
			if strings.Contains(v.expected, ":") {
				ip, _, _ := net.SplitHostPort(v.expected)
				assert.Equal(t, actual, ip)
			} else {
				assert.Equal(t, actual, v.expected)
			}
		}
	}
}

func BenchmarkOS_ClientIp(b *testing.B) {
	b.ResetTimer()
	req := &http.Request{
		RemoteAddr: baiduIpv4,
	}
	for i := 0; i < b.N; i++ {
		KOS.ClientIp(req)
	}
}

func TestOS_IsPortOpen(t *testing.T) {
	var tests = []struct {
		host     string
		port     interface{}
		protocol string
		expected bool
	}{
		{"", 23, "", false},
		{localHost, 0, "", false},
		{localIp, 23, "", false},
		{tesDomain31, 80, "udp", true},
		{tesDomain31, 80, "tcp", true},
		{tesDomain32, "443", "tcp", true},
	}
	for _, test := range tests {
		actual := KOS.IsPortOpen(test.host, test.port, test.protocol)
		assert.Equal(t, actual, test.expected)
	}

	//默认协议
	chk := KOS.IsPortOpen(lanIp, 80)
	assert.False(t, chk)
}

func BenchmarkOS_IsPortOpen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsPortOpen(localIp, 80, "tcp")
	}
}

func TestOS_ForceGC(t *testing.T) {
	KOS.ForceGC()
}

func BenchmarkOS_ForceGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.ForceGC()
	}
}

func TestOS_TriggerGC(t *testing.T) {
	KOS.TriggerGC()
}

func BenchmarkOS_TriggerGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.TriggerGC()
	}
}

func TestOS_GoMemory(t *testing.T) {
	res := KOS.GoMemory()
	assert.Greater(t, int(res), 1)
}

func BenchmarkOS_GoMemory(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GoMemory()
	}
}

func TestOS_GetSystemInfo(t *testing.T) {
	res := KOS.GetSystemInfo()
	assert.NotNil(t, res)
}

func BenchmarkOS_GetSystemInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetSystemInfo()
	}
}

func TestOS_DownloadFile(t *testing.T) {
	var written int64
	var err error

	//非url
	written, err = KOS.DownloadFile(strHello, "", false, nil)
	assert.NotNil(t, err)
	assert.Equal(t, written, int64(0))

	//空路径
	written, err = KOS.DownloadFile(tesUrl39, "", false, nil)
	assert.NotNil(t, err)
	assert.Equal(t, written, int64(0))

	//使用默认客户端
	written, err = KOS.DownloadFile(tesUrl39, downloadfile01, false, nil)
	assert.Nil(t, err)
	assert.Greater(t, written, int64(0))

	//已存在文件
	written, err = KOS.DownloadFile(tesUrl39, downloadfile01, false, nil)
	assert.Nil(t, err)
	assert.Equal(t, written, int64(0))

	//自定义客户端,覆盖已存在文件
	client := &http.Client{}
	client.Timeout = 6 * time.Second
	written, err = KOS.DownloadFile(tesUrl39, downloadfile01, true, client)
	assert.Nil(t, err)
	assert.Greater(t, written, int64(0))
}
