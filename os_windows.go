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

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGlobalMemoryStatusEx = kernel32.NewProc("GlobalMemoryStatusEx")
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
