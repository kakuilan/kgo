package kgo

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"unicode"
)

// SystemInfo 系统信息
type SystemInfo struct {
	ServerName   string  `json:"server_name"`    //服务器名称
	SystemArch   string  `json:"system_arch"`    //系统架构
	SystemOs     string  `json:"system_os"`      //操作系统名称
	Runtime      uint64  `json:"run_time"`       //服务运行时间,纳秒
	Uptime       uint64  `json:"up_time"`        //操作系统运行时间,秒
	GoroutineNum int     `json:"goroutine_num"`  //goroutine数量
	CpuNum       int     `json:"cpu_num"`        //cpu核数
	CpuUser      float64 `json:"cpu_user"`       //cpu用户态比率
	CpuFree      float64 `json:"cpu_free"`       //cpu空闲比率
	DiskUsed     uint64  `json:"disk_used"`      //已用磁盘空间,字节数
	DiskFree     uint64  `json:"disk_free"`      //可用磁盘空间,字节数
	DiskTotal    uint64  `json:"disk_total"`     //总磁盘空间,字节数
	MemUsed      uint64  `json:"mem_used"`       //已用内存,字节数
	MemSys       uint64  `json:"mem_sys"`        //系统内存占用量,字节数
	MemFree      uint64  `json:"mem_free"`       //剩余内存,字节数
	MemTotal     uint64  `json:"mem_total"`      //总内存,字节数
	AllocGolang  uint64  `json:"alloc_golang"`   //golang内存使用量,字节数
	AllocTotal   uint64  `json:"alloc_total"`    //总分配的内存,字节数
	Lookups      uint64  `json:"lookups"`        //指针查找次数
	Mallocs      uint64  `json:"mallocs"`        //内存分配次数
	Frees        uint64  `json:"frees"`          //内存释放次数
	LastGCTime   uint64  `json:"last_gc_time"`   //上次GC时间,纳秒
	NextGC       uint64  `json:"next_gc"`        //下次GC内存回收量,字节数
	PauseTotalNs uint64  `json:"pause_total_ns"` //GC暂停时间总量,纳秒
	PauseNs      uint64  `json:"pause_ns"`       //上次GC暂停时间,纳秒
}

// BiosInfo BIOS信息
type BiosInfo struct {
	Vendor  string `json:"vendor"`
	Version string `json:"version"`
	Date    string `json:"date"`
}

// BoardInfo Board信息
type BoardInfo struct {
	Name     string `json:"name"`
	Vendor   string `json:"vendor"`
	Version  string `json:"version"`
	Serial   string `json:"serial"`
	AssetTag string `json:"assettag"`
}

// CpuInfo CPU信息
type CpuInfo struct {
	Vendor  string `json:"vendor"`
	Model   string `json:"model"`
	Speed   string `json:"speed"`   // CPU clock rate in MHz
	Cache   uint   `json:"cache"`   // CPU cache size in KB
	Cpus    uint   `json:"cpus"`    // number of physical CPUs
	Cores   uint   `json:"cores"`   // number of physical CPU cores
	Threads uint   `json:"threads"` // number of logical (HT) CPU cores
}

var (
	cpuRegTwoColumns = regexp.MustCompile("\t+: ")
	cpuRegExtraSpace = regexp.MustCompile(" +")
	cpuRegCacheSize  = regexp.MustCompile(`^(\d+) KB$`)
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

// PrivateCIDR 获取私有网段的CIDR(无类别域间路由).
func (ko *LkkOS) PrivateCIDR() []*net.IPNet {
	maxCidrBlocks := []string{
		"127.0.0.1/8",    // localhost
		"10.0.0.0/8",     // 24-bit block
		"172.16.0.0/12",  // 20-bit block
		"192.168.0.0/16", // 16-bit block
		"169.254.0.0/16", // link local address
		"::1/128",        // localhost IPv6
		"fc00::/7",       // unique local address IPv6
		"fe80::/10",      // link local address IPv6
	}

	res := make([]*net.IPNet, len(maxCidrBlocks))
	for i, maxCidrBlock := range maxCidrBlocks {
		_, cidr, _ := net.ParseCIDR(maxCidrBlock)
		res[i] = cidr
	}

	return res
}

// IsPrivateIp 是否私有IP地址(ipv4/ipv6).
func (ko *LkkOS) IsPrivateIp(str string) (bool, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return false, errors.New("[IsPrivateIp]`str is not valid ip")
	}

	if kPrivCidrs == nil {
		kPrivCidrs = ko.PrivateCIDR()
	}
	for i := range kPrivCidrs {
		if kPrivCidrs[i].Contains(ip) {
			return true, nil
		}
	}

	return false, nil
}

// IsPublicIP 是否公网IPv4.
func (ko *LkkOS) IsPublicIP(str string) (bool, error) {
	ip := net.ParseIP(str)
	if ip == nil {
		return false, errors.New("[IsPublicIP]`str is not valid ip")
	}

	if ip.IsLoopback() || ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
		return false, nil
	}
	if ip4 := ip.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false, nil
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false, nil
		case ip4[0] == 192 && ip4[1] == 168:
			return false, nil
		}
	}

	return true, nil
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

// GetIpsByDomain 获取互联网域名/主机名对应的 IPv4 地址列表.
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

	_, _ = io.Copy(outWr, stdoutIn)
	_, _ = io.Copy(errWr, stderrIn)
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

// Chmod 改变文件模式.
func (ko *LkkOS) Chmod(filename string, mode os.FileMode) bool {
	return os.Chmod(filename, mode) == nil
}

// Chown 改变文件的所有者.
func (ko *LkkOS) Chown(filename string, uid, gid int) bool {
	return os.Chown(filename, uid, gid) == nil
}

// GetTempDir 返回用于临时文件的目录.
func (ko *LkkOS) GetTempDir() string {
	return os.TempDir()
}

// ClientIp 获取客户端真实IP,req为http请求.
func (ko *LkkOS) ClientIp(req *http.Request) string {
	// 获取头部信息,有可能是代理
	xRealIP := req.Header.Get("X-Real-Ip")
	xForwardedFor := req.Header.Get("X-Forwarded-For")

	// If both empty, return IP from remote address
	if xRealIP == "" && xForwardedFor == "" {
		var remoteIP string

		// If there are colon in remote address, remove the port number
		// otherwise, return remote address as is
		if strings.ContainsRune(req.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(req.RemoteAddr)
		} else {
			remoteIP = req.RemoteAddr
		}

		return remoteIP
	}

	// Check list of IP in X-Forwarded-For and return the first global address
	// X-Forwarded-For是逗号分隔的IP地址列表,如"10.0.0.1, 10.0.0.2, 10.0.0.3"
	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		isPrivate, err := ko.IsPrivateIp(address)
		if !isPrivate && err == nil {
			return address
		}
	}

	if xRealIP == "::1" {
		xRealIP = "127.0.0.1"
	}

	// If nothing succeed, return X-Real-IP
	return xRealIP
}

// IsPortOpen 检查主机端口是否开放.
// host为主机名;port为(整型/字符串)端口号;protocols为协议名称,可选,默认tcp.
func (ko *LkkOS) IsPortOpen(host string, port interface{}, protocols ...string) bool {
	if KStr.IsHost(host) && isPort(port) {
		// 默认tcp协议
		protocol := "tcp"
		if len(protocols) > 0 && len(protocols[0]) > 0 {
			protocol = strings.ToLower(protocols[0])
		}

		conn, _ := net.DialTimeout(protocol, net.JoinHostPort(host, KConv.ToStr(port)), CHECK_CONNECT_TIMEOUT)
		if conn != nil {
			_ = conn.Close()
			return true
		}
	}

	return false
}

// ForceGC 强制手动GC垃圾回收(阻塞).
func (ko *LkkOS) ForceGC() {
	runtime.GC()
	debug.FreeOSMemory()
}

// TriggerGC 触发GC(非阻塞).
func (ko *LkkOS) TriggerGC() {
	go func() {
		ko.ForceGC()
	}()
}

// GoMemory 获取当前go程序的内存使用,返回字节数.
func (ko *LkkOS) GoMemory() uint64 {
	stat := new(runtime.MemStats)
	runtime.ReadMemStats(stat)
	return stat.Alloc
}

// GetSystemInfo 获取系统运行信息.
func (ko *LkkOS) GetSystemInfo() *SystemInfo {
	//运行时信息
	mstat := &runtime.MemStats{}
	runtime.ReadMemStats(mstat)

	//CPU信息
	cpuUser, cpuIdel, cpuTotal := ko.CpuUsage()
	cpuUserRate := float64(cpuUser) / float64(cpuTotal)
	cpuFreeRate := float64(cpuIdel) / float64(cpuTotal)

	//磁盘空间信息
	var diskUsed, diskFree, diskTotal uint64
	if runtime.GOOS == "windows" {
		//TODO 待修改
		diskUsed, diskFree, diskTotal = ko.DiskUsage("C:")
	} else {
		diskUsed, diskFree, diskTotal = ko.DiskUsage("/")
	}

	//内存使用信息
	memUsed, memFree, memTotal := ko.MemoryUsage(true)

	serverName, _ := os.Hostname()
	uptime, _ := ko.Uptime()

	return &SystemInfo{
		ServerName:   serverName,
		SystemArch:   runtime.GOARCH,
		SystemOs:     runtime.GOOS,
		Runtime:      uint64(KTime.ServiceUptime()),
		Uptime:       uptime,
		GoroutineNum: runtime.NumGoroutine(),
		CpuNum:       runtime.NumCPU(),
		CpuUser:      cpuUserRate,
		CpuFree:      cpuFreeRate,
		DiskUsed:     diskUsed,
		DiskFree:     diskFree,
		DiskTotal:    diskTotal,
		MemUsed:      memUsed,
		MemSys:       mstat.Sys,
		MemFree:      memFree,
		MemTotal:     memTotal,
		AllocGolang:  mstat.Alloc,
		AllocTotal:   mstat.TotalAlloc,
		Lookups:      mstat.Lookups,
		Mallocs:      mstat.Mallocs,
		Frees:        mstat.Frees,
		LastGCTime:   mstat.LastGC,
		NextGC:       mstat.NextGC,
		PauseTotalNs: mstat.PauseTotalNs,
		PauseNs:      mstat.PauseNs[(mstat.NumGC+255)%256],
	}
}

// GetProcessExecPath 根据PID获取进程的执行路径.
func (ko *LkkOS) GetProcessExecPath(pid int) string {
	return getProcessPathByPid(pid)
}

// HomeDir 获取当前用户的主目录.
func (ko *LkkOS) HomeDir() (string, error) {
	return os.UserHomeDir()
}

// DownloadFile 下载文件.其中
// url 为文件网址;
// savePath 为保存路径;
// cover 是否覆盖已有文件;
// client 为自定义的请求客户端,不提供时默认使用http.DefaultClient.
func (ko *LkkOS) DownloadFile(url string, savePath string, cover bool, client *http.Client) (written int64, err error) {
	if !isUrl(url) {
		return 0, errors.New("url is not a valid URL")
	}

	savePath = trim(savePath)
	if savePath == "" {
		return 0, errors.New("savePath cannot be empty")
	}

	//已存在
	if !cover && KFile.IsExist(savePath) {
		return 0, nil
	}

	if client == nil {
		client = http.DefaultClient
	}

	//父级目录
	dir := KFile.Dirname(savePath)
	if err = os.MkdirAll(dir, 0766); err != nil {
		return 0, err
	}

	var resp *http.Response
	resp, err = client.Get(url)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	//创建文件
	var out *os.File
	out, err = os.Create(savePath)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = out.Close()
	}()

	//添加缓冲 bufio 是通过缓冲来提高效率。
	wt := bufio.NewWriter(out)
	written, err = io.Copy(wt, resp.Body)
	if err == nil {
		//将缓存的数据写入到文件中
		err = wt.Flush()
	}

	return written, err
}
