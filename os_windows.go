// +build windows

package kgo

import (
	"errors"
	"golang.org/x/sys/windows"
	"os"
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

type FILETIME struct {
	DwLowDateTime  uint32
	DwHighDateTime uint32
}

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGlobalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
	procGetSystemTimes       = kernel32.NewProc("GetSystemTimes")
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
	var lpIdleTime FILETIME
	var lpKernelTime FILETIME
	var lpUserTime FILETIME
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
