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
