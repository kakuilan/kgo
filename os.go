package kgo

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// IsWindows 当前操作系统是否Windows
func (ko *LkkOS) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux 当前操作系统是否Linux
func (ko *LkkOS) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsMac 当前操作系统是否Mac OS/X
func (ko *LkkOS) IsMac() bool {
	return "darwin" == runtime.GOOS
}

// Pwd 获取当前所在路径
func (ko *LkkOS) Pwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)

	return filepath.Dir(pwd)
}

// HomeDir 获取当前用户的主目录(仅支持Unix-like system)
func (ko *LkkOS) HomeDir() (string, error) {
	usr, err := user.Current()
	if nil == err {
		return usr.HomeDir, nil
	}

	// Unix-like system, so just assume Unix
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

// LocalIP 获取本机第一个NIC's IP
func (ko *LkkOS) LocalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if nil != err {
		return "", err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if nil != ipnet.IP.To4() {
				return ipnet.IP.String(), nil
			}
		}
	}

	return "", errors.New("can't get local IP")
}

// OutboundIP 获取本机的出口IP
func (ko *LkkOS) OutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", nil
	}
	defer conn.Close()

	addr := conn.LocalAddr().(*net.UDPAddr)
	return addr.IP.String(), nil
}

// IsPublicIP 是否公网IP
func (ko *LkkOS) IsPublicIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

// GetIPs 获取本机的IP列表
func (ko *LkkOS) GetIPs() (ips []string) {
	interfaceAddr, err := net.InterfaceAddrs()
	if err != nil {
		return
	}

	for _, address := range interfaceAddr {
		ipNet, isValidIpNet := address.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}

	return
}

// GetMacAddrs 获取本机的Mac网卡地址列表
func (ko *LkkOS) GetMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macAddrs = append(macAddrs, macAddr)
	}

	return
}

// Hostname 获取主机名
func (ko *LkkOS) Hostname() (string, error) {
	return os.Hostname()
}

// GetIpByHostname 返回主机名对应的 IPv4地址
func (ko *LkkOS) GetIpByHostname(hostname string) (string, error) {
	ips, err := net.LookupIP(hostname)
	if ips != nil {
		for _, v := range ips {
			if v.To4() != nil {
				return v.String(), nil
			}
		}
		return "", nil
	}
	return "", err
}

// GetIpsByHost 获取互联网域名/主机名对应的 IPv4 地址列表
func (ko *LkkOS) GetIpsByDomain(domain string) ([]string, error) {
	ips, err := net.LookupIP(domain)
	if ips != nil {
		var ipstrs []string
		for _, v := range ips {
			if v.To4() != nil {
				ipstrs = append(ipstrs, v.String())
			}
		}
		return ipstrs, nil
	}
	return nil, err
}

// GetHostByIp 获取指定的IP地址对应的主机名
func (ko *LkkOS) GetHostByIp(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// MemoryGetUsage 获取当前程序的内存使用率,返回字节数
func (ko *LkkOS) MemoryUsage() uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

// Setenv 设置一个环境变量的值
func (ko *LkkOS) Setenv(varname, data string) error {
	return os.Setenv(varname, data)
}

// Getenv 获取一个环境变量的值
func (ko *LkkOS) Getenv(varname string) string {
	return os.Getenv(varname)
}

// GetEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian
func (ko *LkkOS) GetEndian() binary.ByteOrder {
	return getEndian()
}

// IsLittleEndian 系统字节序类型是否小端存储
func (ko *LkkOS) IsLittleEndian() bool {
	return isLittleEndian()
}
