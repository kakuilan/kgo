// +build windows

package kgo

import (
	"errors"
	"github.com/StackExchange/wmi"
	"golang.org/x/sys/windows"
	"os"
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

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGlobalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
	procGetSystemTimes       = kernel32.NewProc("GetSystemTimes")
	procGetDiskFreeSpaceExW  = kernel32.NewProc("GetDiskFreeSpaceExW")
	procGetTickCount64       = kernel32.NewProc("GetTickCount64")
	procGetTickCount32       = kernel32.NewProc("GetTickCount")
)

// HomeDir 获取当前用户的主目录.
func (ko *LkkOS) HomeDir() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// Prefer standard environment variable USERPROFILE
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		return "", errors.New("[HomeDir] HOMEDRIVE, HOMEPATH, or USERPROFILE are blank")
	}

	return home, nil
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
	if mem == 0 {
		return
	}

	total = memInfo.ullTotalPhys
	free = memInfo.ullAvailPhys
	used = memInfo.ullTotalPhys - memInfo.ullAvailPhys

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
	if r == 0 {
		return
	}

	LOT := float64(0.0000001)
	HIT := (LOT * 4294967296.0)
	tmpIdle := ((HIT * float64(lpIdleTime.DwHighDateTime)) + (LOT * float64(lpIdleTime.DwLowDateTime)))
	tmpUser := ((HIT * float64(lpUserTime.DwHighDateTime)) + (LOT * float64(lpUserTime.DwLowDateTime)))
	tmpKernel := ((HIT * float64(lpKernelTime.DwHighDateTime)) + (LOT * float64(lpKernelTime.DwLowDateTime)))
	//tmpSystem := (tmpKernel - tmpIdle)

	user = uint64(tmpUser)
	idle = uint64(tmpIdle)
	total = user + idle + uint64(tmpKernel)

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
	if diskret == 0 {
		return
	}
	total = uint64(lpTotalNumberOfBytes)
	free = uint64(lpTotalNumberOfFreeBytes)
	used = uint64(lpTotalNumberOfBytes - lpTotalNumberOfFreeBytes)

	return
}

// Uptime 获取系统运行时间,秒.
func (ko *LkkOS) Uptime() (uint64, error) {
	procGetTickCount := procGetTickCount64
	err := procGetTickCount64.Find()
	if err != nil {
		// handle WinXP, but keep in mind that "the time will wrap around to zero if the system is run continuously for 49.7 days." from MSDN
		procGetTickCount = procGetTickCount32
	}
	r1, _, lastErr := syscall.Syscall(procGetTickCount.Addr(), 0, 0, 0, 0)
	if lastErr != 0 {
		return 0, lastErr
	}
	return uint64((time.Duration(r1) * time.Millisecond).Seconds()), nil
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
	if err := wmi.Query("SELECT InstallDate, Manufacturer, Version FROM CIM_BIOSElement", &win32BIOSDescriptions); err != nil {
		return res
	}
	if len(win32BIOSDescriptions) > 0 {
		res.Vendor = *win32BIOSDescriptions[0].Manufacturer
		res.Version = *win32BIOSDescriptions[0].Version
		res.Date = *win32BIOSDescriptions[0].InstallDate
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
	if err := wmi.Query("SELECT Manufacturer, SerialNumber, Tag, Version, Product FROM Win32_BaseBoard", &win32BaseboardDescriptions); err != nil {
		return res
	}
	if len(win32BaseboardDescriptions) > 0 {
		res.Name = *win32BaseboardDescriptions[0].Product
		res.Vendor = *win32BaseboardDescriptions[0].Manufacturer
		res.Version = *win32BaseboardDescriptions[0].Version
		res.Serial = *win32BaseboardDescriptions[0].SerialNumber
		res.AssetTag = *win32BaseboardDescriptions[0].Tag
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
	if err := wmi.Query("SELECT Manufacturer, Name, NumberOfLogicalProcessors, NumberOfCores, MaxClockSpeed, L2CacheSize, L3CacheSize FROM Win32_Processor", &win32descriptions); err != nil {
		return res
	}

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

	return res
}
