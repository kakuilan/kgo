//go:build windows
// +build windows

package kgo

import (
	"errors"
	"fmt"
	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

//内存状态扩展
type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

type fileTime struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

type win32BIOS struct {
	InstallDate  *string
	Manufacturer *string
	Version      *string
}

type win32Baseboard struct {
	Manufacturer *string
	SerialNumber *string
	Tag          *string
	Version      *string
	Product      *string
}

type win32Processor struct {
	Manufacturer              *string
	Name                      *string
	NumberOfLogicalProcessors uint32
	NumberOfCores             uint32
	MaxClockSpeed             uint32
	L2CacheSize               uint32
	L3CacheSize               uint32
}

type Win32_Process struct {
	Name                  string
	ExecutablePath        *string
	CommandLine           *string
	Priority              uint32
	CreationDate          *time.Time
	ProcessId             uint32
	ThreadCount           uint32
	Status                *string
	ReadOperationCount    uint64
	ReadTransferCount     uint64
	WriteOperationCount   uint64
	WriteTransferCount    uint64
	CSCreationClassName   string
	CSName                string
	Caption               *string
	CreationClassName     string
	Description           *string
	ExecutionState        *uint16
	HandleCount           uint32
	KernelModeTime        uint64
	MaximumWorkingSetSize *uint32
	MinimumWorkingSetSize *uint32
	OSCreationClassName   string
	OSName                string
	OtherOperationCount   uint64
	OtherTransferCount    uint64
	PageFaults            uint32
	PageFileUsage         uint32
	ParentProcessID       uint32
	PeakPageFileUsage     uint32
	PeakVirtualSize       uint64
	PeakWorkingSetSize    uint32
	PrivatePageCount      uint64
	TerminationDate       *time.Time
	UserModeTime          uint64
	WorkingSetSize        uint64
}

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGlobalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
	procGetSystemTimes       = kernel32.NewProc("GetSystemTimes")
	procGetDiskFreeSpaceExW  = kernel32.NewProc("GetDiskFreeSpaceExW")
	procGetTickCount64       = kernel32.NewProc("GetTickCount64")
	procGetTickCount32       = kernel32.NewProc("GetTickCount")
)

// getProcessByPid 根据pid获取进程列表.
func getProcessByPid(pid int) (res []Win32_Process) {
	var dst []Win32_Process
	query := fmt.Sprintf("WHERE ProcessId = %d", pid)
	q := wmi.CreateQuery(&dst, query)
	err := wmi.Query(q, &dst)
	if err == nil && len(dst) > 0 {
		res = dst
	}

	return
}

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) (res string) {
	ps := getProcessByPid(pid)
	if len(ps) > 0 {
		res = *ps[0].ExecutablePath
	}

	return
}

// MemoryUsage 获取内存使用率,单位字节.
// 参数 virtual(仅支持linux),是否取虚拟内存.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) MemoryUsage(virtual bool) (used, free, total uint64) {
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem > 0 {
		total = memInfo.ullTotalPhys
		free = memInfo.ullAvailPhys
		used = memInfo.ullTotalPhys - memInfo.ullAvailPhys
	}

	return
}

// CpuUsage 获取CPU使用率(darwin系统必须使用cgo),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (ko *LkkOS) CpuUsage() (user, idle, total uint64) {
	var lpIdleTime fileTime
	var lpKernelTime fileTime
	var lpUserTime fileTime
	r, _, _ := procGetSystemTimes.Call(
		uintptr(unsafe.Pointer(&lpIdleTime)),
		uintptr(unsafe.Pointer(&lpKernelTime)),
		uintptr(unsafe.Pointer(&lpUserTime)))

	if r > 0 {
		LOT := float64(0.0000001)
		HIT := (LOT * 4294967296.0)
		tmpIdle := ((HIT * float64(lpIdleTime.DwHighDateTime)) + (LOT * float64(lpIdleTime.DwLowDateTime)))
		tmpUser := ((HIT * float64(lpUserTime.DwHighDateTime)) + (LOT * float64(lpUserTime.DwLowDateTime)))
		tmpKernel := ((HIT * float64(lpKernelTime.DwHighDateTime)) + (LOT * float64(lpKernelTime.DwLowDateTime)))
		//tmpSystem := (tmpKernel - tmpIdle)

		user = uint64(tmpUser)
		idle = uint64(tmpIdle)
		total = user + idle + uint64(tmpKernel)
	}

	return
}

// DiskUsage 获取磁盘(目录)使用情况,单位字节.参数path为路径.
// used为已用,
// free为空闲,
// total为总数.
func (ko *LkkOS) DiskUsage(path string) (used, free, total uint64) {
	lpFreeBytesAvailable := int64(0)
	lpTotalNumberOfBytes := int64(0)
	lpTotalNumberOfFreeBytes := int64(0)
	diskret, _, _ := procGetDiskFreeSpaceExW.Call(
		uintptr(unsafe.Pointer(windows.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))

	if diskret > 0 {
		total = uint64(lpTotalNumberOfBytes)
		free = uint64(lpTotalNumberOfFreeBytes)
		used = uint64(lpTotalNumberOfBytes - lpTotalNumberOfFreeBytes)
	}

	return
}

// Uptime 获取系统运行时间,秒.
func (ko *LkkOS) Uptime() (uint64, error) {
	var res uint64
	var err error

	procGetTickCount := procGetTickCount64
	chkErr := procGetTickCount64.Find()
	if chkErr != nil {
		// handle WinXP, but keep in mind that "the time will wrap around to zero if the system is run continuously for 49.7 days." from MSDN
		procGetTickCount = procGetTickCount32
	}

	ret, _, errno := syscall.Syscall(procGetTickCount.Addr(), 0, 0, 0, 0)
	if errno == 0 {
		res = uint64((time.Duration(ret) * time.Millisecond).Seconds())
	} else {
		err = errors.New(errno.Error())
	}

	return res, err
}

// GetBiosInfo 获取BIOS信息.
// 注意:Mac机器没有BIOS信息,它使用EFI.
func (ko *LkkOS) GetBiosInfo() *BiosInfo {
	res := &BiosInfo{
		Vendor:  "",
		Version: "",
		Date:    "",
	}

	// Getting data from WMI
	var win32BIOSDescriptions []win32BIOS
	if err := wmi.Query("SELECT InstallDate, Manufacturer, Version FROM CIM_BIOSElement", &win32BIOSDescriptions); err == nil {
		if len(win32BIOSDescriptions) > 0 {
			res.Vendor = *win32BIOSDescriptions[0].Manufacturer
			res.Version = *win32BIOSDescriptions[0].Version
			res.Date = *win32BIOSDescriptions[0].InstallDate
		}
	}

	return res
}

// GetBoardInfo 获取Board信息.
func (ko *LkkOS) GetBoardInfo() *BoardInfo {
	res := &BoardInfo{
		Name:     "",
		Vendor:   "",
		Version:  "",
		Serial:   "",
		AssetTag: "",
	}

	// Getting data from WMI
	var win32BaseboardDescriptions []win32Baseboard
	if err := wmi.Query("SELECT Manufacturer, SerialNumber, Tag, Version, Product FROM Win32_BaseBoard", &win32BaseboardDescriptions); err == nil {
		if len(win32BaseboardDescriptions) > 0 {
			res.Name = *win32BaseboardDescriptions[0].Product
			res.Vendor = *win32BaseboardDescriptions[0].Manufacturer
			res.Version = *win32BaseboardDescriptions[0].Version
			res.Serial = *win32BaseboardDescriptions[0].SerialNumber
			res.AssetTag = *win32BaseboardDescriptions[0].Tag
		}
	}

	return res
}

// GetCpuInfo 获取CPU信息.
func (ko *LkkOS) GetCpuInfo() *CpuInfo {
	var res = &CpuInfo{
		Vendor:  "",
		Model:   "",
		Speed:   "",
		Cache:   0,
		Cpus:    0,
		Cores:   0,
		Threads: 0,
	}

	// Getting info from WMI
	var win32descriptions []win32Processor
	if err := wmi.Query("SELECT Manufacturer, Name, NumberOfLogicalProcessors, NumberOfCores, MaxClockSpeed, L2CacheSize, L3CacheSize FROM Win32_Processor", &win32descriptions); err == nil {
		var cores, threads uint
		for _, description := range win32descriptions {
			if res.Vendor == "" {
				res.Vendor = *description.Manufacturer
			}
			if res.Model == "" {
				res.Model = *description.Name
			}
			if res.Speed == "" {
				res.Speed = toStr(description.MaxClockSpeed)
			}
			if res.Cache == 0 {
				res.Cache = uint(description.L2CacheSize + description.L3CacheSize)
			}

			cores += uint(description.NumberOfCores)
			threads += uint(description.NumberOfLogicalProcessors)
		}

		res.Cpus = uint(len(win32descriptions))
		res.Cores = cores
		res.Threads = threads
	}

	return res
}

// IsProcessExists 进程是否存在.
func (ko *LkkOS) IsProcessExists(pid int) (res bool) {
	if pid <= 0 {
		return
	} else if pid%4 != 0 {
		// OpenProcess will succeed even on non-existing pid here https://devblogs.microsoft.com/oldnewthing/20080606-00/?p=22043
		// so we list every pid just to be sure and be future-proof
		ps := getProcessByPid(pid)
		if len(ps) > 0 && ps[0].ProcessId > 0 {
			return true
		}
	} else {
		var still_active uint32 = 259 // https://docs.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-getexitcodeprocess
		h, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, uint32(pid))
		if err != nil {
			if err == windows.ERROR_ACCESS_DENIED {
				res = true
			}
			return
		}

		defer func() {
			_ = syscall.CloseHandle(syscall.Handle(h))
		}()
		var exitCode uint32
		_ = windows.GetExitCodeProcess(h, &exitCode)
		res = exitCode == still_active
	}

	return
}

// GetPidByPort 根据端口号获取监听的进程PID.
// linux可能要求root权限;
// darwin依赖lsof;
// windows依赖netstat.
func (ko *LkkOS) GetPidByPort(port int) (pid int) {
	command := fmt.Sprintf("cmd /C netstat -ano | findstr %d", port)
	_, out, _ := ko.System(command)
	lines := strings.Split(string(out), "\r\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 5 && isNumeric(fields[4]) {
			p := toInt(fields[4])
			if p > 0 {
				pid = p
				break
			}
		}
	}

	return
}
