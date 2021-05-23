package kgo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"
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

// GetIpsByHost 获取互联网域名/主机名对应的 IPv4 地址列表.
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

// GetHostByIp 获取指定的IP地址对应的主机名.
func (ko *LkkOS) GetHostByIp(ipAddress string) (string, error) {
	names, err := net.LookupAddr(ipAddress)
	if names != nil {
		return strings.TrimRight(names[0], "."), nil
	}
	return "", err
}

// Setenv 设置一个环境变量的值.
func (ko *LkkOS) Setenv(varname, data string) error {
	return os.Setenv(varname, data)
}

// Getenv 获取一个环境变量的值.defvalue为默认值.
func (ko *LkkOS) Getenv(varname string, defvalue ...string) string {
	val := os.Getenv(varname)
	if val == "" && len(defvalue) > 0 {
		val = defvalue[0]
	}

	return val
}

// Unsetenv 删除一个环境变量.
func (ko *LkkOS) Unsetenv(varname string) error {
	return os.Unsetenv(varname)
}

// GetEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian .
func (ko *LkkOS) GetEndian() binary.ByteOrder {
	return getEndian()
}

// IsLittleEndian 系统字节序类型是否小端存储.
func (ko *LkkOS) IsLittleEndian() bool {
	return isLittleEndian()
}

// Exec 执行一个外部命令.
// retInt为1时失败,为0时成功;outStr为执行命令的输出;errStr为错误输出.
// 命令如
// "ls -a"
// "/bin/bash -c \"ls -a\""
func (ko *LkkOS) Exec(command string) (retInt int, outStr, errStr []byte) {
	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})

	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}

	var stdout, stderr bytes.Buffer
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		retInt = 1 //失败
		stderr.WriteString(err.Error())
		errStr = stderr.Bytes()
	} else {
		retInt = 0 //成功
		outStr, errStr = stdout.Bytes(), stderr.Bytes()
	}

	return
}

// System 与Exec相同,但会同时打印标准输出和标准错误.
func (ko *LkkOS) System(command string) (retInt int, outStr, errStr []byte) {
	// split command
	q := rune(0)
	parts := strings.FieldsFunc(command, func(r rune) bool {
		switch {
		case r == q:
			q = rune(0)
			return false
		case q != rune(0):
			return false
		case unicode.In(r, unicode.Quotation_Mark):
			q = r
			return false
		default:
			return unicode.IsSpace(r)
		}
	})

	// remove the " and ' on both sides
	for i, v := range parts {
		f, l := v[0], len(v)
		if l >= 2 && (f == '"' || f == '\'') {
			parts[i] = v[1 : l-1]
		}
	}

	var stdout, stderr bytes.Buffer
	var err error

	cmd := exec.Command(parts[0], parts[1:]...)
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	outWr := io.MultiWriter(os.Stdout, &stdout)
	errWr := io.MultiWriter(os.Stderr, &stderr)

	err = cmd.Start()
	if err != nil {
		retInt = 1 //失败
		stderr.WriteString(err.Error())
		fmt.Printf("%s\n", stderr.Bytes())
		return
	}

	go func() {
		_, _ = io.Copy(outWr, stdoutIn)
	}()
	go func() {
		_, _ = io.Copy(errWr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		stderr.WriteString(err.Error())
		fmt.Println(stderr.Bytes())
		retInt = 1 //失败
	} else {
		retInt = 0 //成功
	}
	outStr, errStr = stdout.Bytes(), stderr.Bytes()

	return
}
