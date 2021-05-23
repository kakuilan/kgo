package kgo

import (
	"net"
	"os"
	"path/filepath"
	"runtime"
)

// IsWindows 当前操作系统是否Windows.
func (ko *LkkOS) IsWindows() bool {
	return "windows" == runtime.GOOS
}

// IsLinux 当前操作系统是否Linux.
func (ko *LkkOS) IsLinux() bool {
	return "linux" == runtime.GOOS
}

// IsMac 当前操作系统是否Mac OS/X.
func (ko *LkkOS) IsMac() bool {
	return "darwin" == runtime.GOOS
}

// Pwd 获取当前程序运行所在的路径,注意和Getwd有所不同.
// 若当前执行的是链接文件,则会指向真实二进制程序的所在目录.
func (ko *LkkOS) Pwd() string {
	var dir, ex string
	var err error
	ex, err = os.Executable()
	if err == nil {
		exReal, _ := filepath.EvalSymlinks(ex)
		exReal, _ = filepath.Abs(exReal)
		dir = filepath.Dir(exReal)
	}

	return dir
}

// Getcwd 取得当前工作目录(程序可能在任务中进行多次目录切换).
func (ko *LkkOS) Getcwd() (string, error) {
	dir, err := os.Getwd()
	return dir, err
}

// Chdir 改变/进入新的工作目录.
func (ko *LkkOS) Chdir(dir string) error {
	return os.Chdir(dir)
}

// LocalIP 获取本机第一个NIC's IP.
func (ko *LkkOS) LocalIP() (string, error) {
	res := ""
	addrs, err := net.InterfaceAddrs()
	if len(addrs) > 0 {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if nil != ipnet.IP.To4() {
					res = ipnet.IP.String()
					break
				}
			}
		}
	}

	return res, err
}

// OutboundIP 获取本机的出口IP.
func (ko *LkkOS) OutboundIP() (string, error) {
	res := ""
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if conn != nil {
		addr := conn.LocalAddr().(*net.UDPAddr)
		res = addr.IP.String()
		_ = conn.Close()
	}

	return res, err
}

// IsPublicIP 是否公网IPv4.
func (ko *LkkOS) IsPublicIPv4(ip net.IP) bool {
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

// GetIPs 获取本机的IP列表.
func (ko *LkkOS) GetIPs() (ips []string) {
	interfaceAddrs, _ := net.InterfaceAddrs()
	if len(interfaceAddrs) > 0 {
		for _, addr := range interfaceAddrs {
			ipNet, isValidIpNet := addr.(*net.IPNet)
			if isValidIpNet && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					ips = append(ips, ipNet.IP.String())
				}
			}
		}
	}

	return
}

// GetMacAddrs 获取本机的Mac网卡地址列表.
func (ko *LkkOS) GetMacAddrs() (macAddrs []string) {
	netInterfaces, _ := net.Interfaces()
	if len(netInterfaces) > 0 {
		for _, netInterface := range netInterfaces {
			macAddr := netInterface.HardwareAddr.String()
			if len(macAddr) == 0 {
				continue
			}
			macAddrs = append(macAddrs, macAddr)
		}
	}

	return
}

// Hostname 获取主机名.
func (ko *LkkOS) Hostname() (string, error) {
	return os.Hostname()
}

// GetIpByHostname 返回主机名对应的 IPv4地址.
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
