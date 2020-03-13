package kgo

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os/user"
	"strings"
	"testing"
	"time"
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

func TestGetcwd(t *testing.T) {
	res, err := KOS.Getcwd()
	if err != nil || res == "" {
		t.Error("Getcwd fail")
		return
	}
}

func BenchmarkGetcwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.Getcwd()
	}
}

func TestChdir(t *testing.T) {
	err := KOS.Chdir("./testdata")
	if err != nil {
		println(err.Error())
		t.Error("Chdir fail")
		return
	}

	err = KOS.Chdir("../")
	if err != nil {
		println(err.Error())
		t.Error("Chdir fail")
		return
	}
	_ = KOS.Chdir("")
}

func BenchmarkChdir(b *testing.B) {
	b.ResetTimer()
	dir := KOS.Pwd()
	for i := 0; i < b.N; i++ {
		_ = KOS.Chdir(dir)
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
	KOS.IsPublicIP(net.ParseIP("127.0.0.1"))
	KOS.IsPublicIP(net.ParseIP("172.16.0.1"))
	KOS.IsPublicIP(net.ParseIP("192.168.0.1"))
	//google
	KOS.IsPublicIP(net.ParseIP("172.217.26.142"))
	//google IPv6
	KOS.IsPublicIP(net.ParseIP("2404:6800:4005:80f::200e"))
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
	_, _ = KOS.GetIpByHostname("hello")
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

	_, _ = KOS.GetHostByIp("192.168.1.1")
}

func BenchmarkGetHostByIp(b *testing.B) {
	b.ResetTimer()
	ip := "127.0.0.1"
	for i := 0; i < b.N; i++ {
		_, _ = KOS.GetHostByIp(ip)
	}
}

func TestGoMemory(t *testing.T) {
	mem := KOS.GoMemory()
	if mem == 0 {
		t.Error("GoMemory fail")
		return
	}
}

func BenchmarkGoMemory(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GoMemory()
	}
}

func TestMemoryUsage(t *testing.T) {
	// 虚拟内存
	used1, free1, total1 := KOS.MemoryUsage(true)
	//usedRate1 := float64(used1) / float64(total1)
	if used1 <= 0 || free1 <= 0 || total1 <= 0 {
		t.Error("MemoryUsage(true) fail")
		return
	}

	// 真实物理内存
	used2, free2, total2 := KOS.MemoryUsage(false)
	//usedRate2 := float64(used2) / float64(total2)
	if used2 <= 0 || free2 <= 0 || total2 <= 0 {
		t.Error("MemoryUsage(false) fail")
		return
	}
}

func BenchmarkMemoryUsageVirtual(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(true)
	}
}

func BenchmarkMemoryUsagePhysic(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.MemoryUsage(false)
	}
}

func TestCpuUsage(t *testing.T) {
	usr, idle, total := KOS.CpuUsage()

	// 注意: usr/total,两个整数相除，结果是整数，取整数部分为0
	usedRate := float64(usr) / float64(total)
	freeRate := float64(idle) / float64(total)

	if usedRate == 0 || freeRate == 0 {
		t.Error("CpuUsage fail")
		return
	}
}

func BenchmarkCpuUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.CpuUsage()
	}
}

func TestDiskUsage(t *testing.T) {
	used, free, total := KOS.DiskUsage("/")
	if used <= 0 || free <= 0 || total <= 0 {
		t.Error("DiskUsage fail")
		return
	}
}

func BenchmarkDiskUsage(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.DiskUsage("/")
	}
}

func BenchmarkSetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KOS.Setenv("HELLO", "world")
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

	val3 := KOS.Getenv("admusr", "zhang3")
	if val3 != "zhang3" {
		t.Error("Getenv fail")
		return
	}
}

func BenchmarkGetenv(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Getenv("HELLO")
	}
}

func TestGetEndian_IsLittleEndian(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	endi := KOS.GetEndian()
	isLit := KOS.IsLittleEndian()

	if fmt.Sprintf("%v", endi) == "" {
		t.Error("GetEndian fail")
		return
	} else if isLit && fmt.Sprintf("%v", endi) != "LittleEndian" {
		t.Error("IsLittleEndian fail")
		return
	}
}

func BenchmarkGetEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetEndian()
	}
}

func BenchmarkIsLittleEndian(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLittleEndian()
	}
}

func TestExec(t *testing.T) {
	cmd := " ls -a -h"
	ret, _, _ := KOS.Exec(cmd)
	if ret == 1 {
		t.Error("Exec fail")
		return
	}

	cmd = " ls -a\"\" -h 'hehe'"
	_, _, _ = KOS.Exec(cmd)
}

func BenchmarkExec(b *testing.B) {
	b.ResetTimer()
	cmd := " ls -a -h"
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.Exec(cmd)
	}
}

func TestSystem(t *testing.T) {
	cmd := " ls -a -h"
	ret, _, _ := KOS.System(cmd)
	if ret == 1 {
		t.Error("System fail")
		return
	}

	cmd = "123"
	_, _, _ = KOS.System(cmd)

	cmd = " ls -a\"\" -h 'hehe'"
	_, _, _ = KOS.System(cmd)

	cmd = "ls -a /root/"
	_, _, _ = KOS.System(cmd)

	filename := ""
	for i := 0; i < 10000; i++ {
		filename = fmt.Sprintf("./testdata/empty/zero_%d", i)
		KFile.Touch(filename, 0)
	}

	cmd = "ls -a -h ./testdata/empty"
	_, _, _ = KOS.System(cmd)
	_, _, _ = KOS.System(cmd)
	_, _, _ = KOS.System(cmd)

	cmd = "touch /root/hello"
	_, _, _ = KOS.System(cmd)
	_ = KFile.DelDir("./testdata/empty", false)
}

func BenchmarkSystem(b *testing.B) {
	b.ResetTimer()
	cmd := " ls -a -h"
	for i := 0; i < b.N; i++ {
		_, _, _ = KOS.System(cmd)
	}
}

func TestChmodChown(t *testing.T) {
	file := "./testdata"
	res1 := KOS.Chmod(file, 0777)

	usr, _ := user.Current()
	uid := KConv.Str2Int(usr.Uid)
	guid := KConv.Str2Int(usr.Gid)

	res2 := KOS.Chown(file, uid, guid)

	if !res1 || !res2 {
		t.Error("Chmod fail")
		return
	}
}

func BenchmarkChmod(b *testing.B) {
	b.ResetTimer()
	file := "./testdata"
	for i := 0; i < b.N; i++ {
		KOS.Chmod(file, 0777)
	}
}

func BenchmarkChown(b *testing.B) {
	b.ResetTimer()
	file := "./testdata"
	usr, _ := user.Current()
	uid := KConv.Str2Int(usr.Uid)
	guid := KConv.Str2Int(usr.Gid)
	for i := 0; i < b.N; i++ {
		KOS.Chown(file, uid, guid)
	}
}

func TestGetTempDir(t *testing.T) {
	res := KOS.GetTempDir()
	if res == "" {
		t.Error("GetTempDir fail")
		return
	}
}

func BenchmarkGetTempDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetTempDir()
	}
}

func TestPrivateCIDR(t *testing.T) {
	res := KOS.PrivateCIDR()
	if len(res) == 0 {
		t.Error("PrivateCIDR fail")
		return
	}
}

func BenchmarkPrivateCIDR(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.PrivateCIDR()
	}
}

func TestIsPrivateIp(t *testing.T) {
	//无效Ip
	res, err := KOS.IsPrivateIp("hello")
	if res || err == nil {
		t.Error("IsPrivateIp fail")
		return
	}

	//KPrivCidrs未初始化数据
	if len(KPrivCidrs) != 0 {
		t.Error("IsPrivateIp fail")
		return
	}

	//docker ip
	res, err = KOS.IsPrivateIp("172.17.0.1")
	if !res || err != nil {
		t.Error("IsPrivateIp fail")
		return
	}

	//外网ip
	res, err = KOS.IsPrivateIp("220.181.38.148")
	if res || err != nil {
		t.Error("IsPrivateIp fail")
		return
	}

	//KPrivCidrs已初始化数据
	if len(KPrivCidrs) == 0 {
		t.Error("IsPrivateIp fail")
		return
	}
}

func BenchmarkIsPrivateIp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.IsPrivateIp("172.17.0.1")
	}
}

func TestClientIp(t *testing.T) {
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
	publicAddr1 := "144.12.54.87"
	publicAddr2 := "119.14.55.11"
	publicAddr3 := "8.8.8.8:8080"
	localAddr1 := "127.0.0.0"
	localAddr2 := "::1"

	testData := []testIP{
		{
			name:     "No header,no port",
			request:  newRequest(publicAddr1, ""),
			expected: publicAddr1,
		}, {
			name:     "No header,has port",
			request:  newRequest(publicAddr3, ""),
			expected: publicAddr3,
		}, {
			name:     "Has X-Forwarded-For",
			request:  newRequest("", "", publicAddr1),
			expected: publicAddr1,
		}, {
			name:     "Has multiple X-Forwarded-For",
			request:  newRequest("", "", localAddr1, publicAddr1, publicAddr2),
			expected: publicAddr2,
		}, {
			name:     "Has X-Real-IP",
			request:  newRequest("", publicAddr1),
			expected: publicAddr1,
		}, {
			name:     "Local ip",
			request:  newRequest("", localAddr2),
			expected: localAddr2,
		},
	}

	// Run test
	var actual string
	for _, v := range testData {
		actual = KOS.ClientIp(v.request)
		if v.expected == "::1" {
			if actual != "127.0.0.1" {
				t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
			}
		} else {
			if strings.Contains(v.expected, ":") {
				ip, _, _ := net.SplitHostPort(v.expected)
				if ip != actual {
					t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
				}
			} else {
				if v.expected != actual {
					t.Errorf("%s: expected %s but get %s", v.name, v.expected, actual)
				}
			}
		}
	}
}

func BenchmarkClientIp(b *testing.B) {
	b.ResetTimer()
	req := &http.Request{
		RemoteAddr: "216.58.199.14",
	}
	for i := 0; i < b.N; i++ {
		KOS.ClientIp(req)
	}
}

func TestGetSystemInfo(t *testing.T) {
	info := KOS.GetSystemInfo()
	fmt.Printf("%+v\n", info)
}

func BenchmarkGetSystemInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetSystemInfo()
	}
}

func TestIsPortOpen(t *testing.T) {
	var tests = []struct {
		host     string
		port     interface{}
		protocol string
		expected bool
	}{
		{"", 23, "", false},
		{"localhost", 0, "", false},
		{"127.0.0.1", 23, "", false},
		{"golang.org", 80, "udp", true},
		{"golang.org", 80, "tcp", true},
		{"www.google.com", "443", "tcp", true},
	}
	for _, test := range tests {
		actual := KOS.IsPortOpen(test.host, test.port, test.protocol)
		if actual != test.expected {
			t.Errorf("Expected IsChineseName(%q, %v, %q) to be %v, got %v", test.host, test.port, test.protocol, test.expected, actual)
		}
	}

	KOS.IsPortOpen("127.0.0.1", 80, "tcp")
	KOS.IsPortOpen("::", 80, "tcp")
	KOS.IsPortOpen("::", 80, "")
	KOS.IsPortOpen("::", 80)
}

func BenchmarkIsPortOpen(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsPortOpen("127.0.0.1", 80, "tcp")
	}
}

func TestGetPidByPortGetProcessExeByPid(t *testing.T) {
	message := "Hi there!\n"

	time.AfterFunc(time.Millisecond*200, func() {
		getPidByInode("1234", nil)
		KOS.GetPidByPort(22)
		KOS.GetPidByPort(25)
		KOS.GetPidByPort(1999)
		res := KOS.GetPidByPort(2020)
		exepath := KOS.GetProcessExeByPid(res)
		if res == 0 {
			t.Error("GetPidByPort fail")
			return
		}
		if exepath == "" {
			t.Error("getProcessExeByPid fail")
			return
		}
	})

	time.AfterFunc(time.Millisecond*500, func() {
		conn, err := net.Dial("tcp", ":2020")
		if err != nil {
			t.Fatal(err)
		}
		defer conn.Close()

		if _, err := fmt.Fprintf(conn, message); err != nil {
			t.Fatal(err)
		}
	})

	l, err := net.Listen("tcp", ":2020")
	if err != nil {
		t.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		buf, err := ioutil.ReadAll(conn)
		if err != nil {
			t.Fatal(err)
		}

		if msg := string(buf[:]); msg != message {
			t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
		}
		return // Done
	}
}

func BenchmarkGetPidByPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetPidByPort(2020)
	}
}

func BenchmarkGetProcessExeByPid(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetProcessExeByPid(2020)
	}
}

func TestForceGC(t *testing.T) {
	KOS.ForceGC()
}

func TestTriggerGC(t *testing.T) {
	KOS.TriggerGC()
}

func BenchmarkForceGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.ForceGC()
	}
}

func BenchmarkTriggerGC(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.TriggerGC()
	}
}
