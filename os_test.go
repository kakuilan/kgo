package kgo

import (
	"net"
	"testing"
)

func TestIsWindows(t *testing.T) {
	res := KOS.IsWindows()
	if res {
		t.Error("IsWindows fail")
		return
	}
}

func BenchmarkIsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsWindows()
	}
}

func TestIsLinux(t *testing.T) {
	res := KOS.IsLinux()
	if !res {
		t.Error("IsLinux fail")
		return
	}
}

func BenchmarkIsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLinux()
	}
}

func TestIsMac(t *testing.T) {
	res := KOS.IsMac()
	if res {
		t.Error("IsMac fail")
		return
	}
}

func BenchmarkIsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}

func TestPwd(t *testing.T) {
	res := KOS.Pwd()
	if res == "" {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkPwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestHomeDir(t *testing.T) {
	_, err := KOS.HomeDir()
	if err != nil {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkHomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestLocalIP(t *testing.T) {
	_, err := KOS.LocalIP()
	if err != nil {
		t.Error("LocalIP fail")
		return
	}
}

func BenchmarkLocalIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.LocalIP()
	}
}

func TestOutboundIP(t *testing.T) {
	_, err := KOS.OutboundIP()
	if err != nil {
		t.Error("OutboundIP fail")
		return
	}
}

func BenchmarkOutboundIP(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.OutboundIP()
	}
}

func TestIsPublicIP(t *testing.T) {
	ipStr, _ := KOS.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	if KOS.IsPublicIP(ipAddr) {
		t.Error("IsPublicIP fail")
		return
	}
}

func BenchmarkIsPublicIP(b *testing.B) {
	b.ResetTimer()
	ipStr, _ := KOS.LocalIP()
	ipAddr := net.ParseIP(ipStr)
	for i := 0; i < b.N; i++ {
		KOS.IsPublicIP(ipAddr)
	}
}

func TestGetIPs(t *testing.T) {
	ips := KOS.GetIPs()
	if len(ips) == 0 {
		t.Error("GetIPs fail")
		return
	}
}

func BenchmarkGetIPs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetIPs()
	}
}

func TestGetMacAddrs(t *testing.T) {
	macs := KOS.GetMacAddrs()
	//fmt.Printf("%v", macs)
	if len(macs) == 0 {
		t.Error("GetMacAddrs fail")
		return
	}
}

func BenchmarkGetMacAddrs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetMacAddrs()
	}
}

func TestHostname(t *testing.T) {
	res, err := KOS.Hostname()
	if err != nil || res == "" {
		t.Error("Hostname fail")
		return
	}
}

func BenchmarkHostname(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Hostname()
	}
}

func TestGetIpByHostname(t *testing.T) {
	name := "localhost"
	ip, err := KOS.GetIpByHostname(name)
	if err != nil || ip != "127.0.0.1" {
		t.Error("GetIpByHostname fail")
		return
	}

	_, _ = KOS.GetIpByHostname("::1")
	_, err = KOS.GetIpByHostname("hello")
	if err != nil {
		t.Error("GetIpByHostname fail")
		return
	}

}

func BenchmarkGetIpByHostname(b *testing.B) {
	b.ResetTimer()
	name := "localhost"
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpByHostname(name)
	}
}

func TestGetIpsByDomain(t *testing.T) {
	name := "google.com"
	ips, err := KOS.GetIpsByDomain(name)
	if err != nil || len(ips) == 0 {
		t.Error("GetIpsByDomain fail")
		return
	}

	ips, err = KOS.GetIpsByDomain("hello")
	if err == nil || len(ips) > 0 {
		t.Error("GetIpsByDomain fail")
		return
	}
}

func BenchmarkGetIpsByDomain(b *testing.B) {
	b.ResetTimer()
	name := "google.com"
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetIpsByDomain(name)
	}
}

func TestGetHostByIp(t *testing.T) {
	ip := "127.0.0.1"
	host, err := KOS.GetHostByIp(ip)
	if err != nil || host == "" {
		t.Error("GetHostByIp fail")
		return
	}

	ip, err = KOS.GetHostByIp("192.168.1.1")
	if err != nil {
		println(err.Error())
	}
}

func BenchmarkGetHostByIp(b *testing.B) {
	b.ResetTimer()
	ip := "127.0.0.1"
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetHostByIp(ip)
	}
}

func TestMemoryUsage(t *testing.T) {
	mem := KOS.MemoryUsage()
	if mem == 0 {
		t.Error("MemoryUsage fail")
		return
	}
}

func BenchmarkMemoryUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage()
	}
}

func TestSetenvGetenv(t *testing.T) {
	name1 := "HELLO"
	name2 := "HOME"

	err := KOS.Setenv(name1, "world")
	if err != nil {
		t.Error("Setenv fail")
		return
	}

	val1 := KOS.Getenv(name1)
	val2 := KOS.Getenv(name2)
	if val1 != "world" || val2 == "" {
		t.Error("Getenv fail")
		return
	}
}

func BenchmarkSetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KOS.Setenv("HELLO", "world")
	}
}

func BenchmarkGetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Getenv("HELLO")
	}
}
