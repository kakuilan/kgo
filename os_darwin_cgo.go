// +build darwin
// +build cgo

package kgo

import (
	"unsafe"
)

// #include <mach/mach_host.h>
// #include <mach/host_info.h>
import "C"

// 获取CPU使用率(darwin系统必须使用cgo),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (ko *LkkOS) CpuUsage() (user, idle, total uint64) {
	var cpuLoad C.host_cpu_load_info_data_t
	var count C.mach_msg_type_number_t = C.HOST_CPU_LOAD_INFO_COUNT
	ret := C.host_statistics(C.host_t(C.mach_host_self()), C.HOST_CPU_LOAD_INFO, C.host_info_t(unsafe.Pointer(&cpuLoad)), &count)
	if ret != C.KERN_SUCCESS {
		return
	}

	user = uint64(cpuLoad.cpu_ticks[C.CPU_STATE_USER])
	idle = uint64(cpuLoad.cpu_ticks[C.CPU_STATE_IDLE])
	total = user + idle + uint64(cpuLoad.cpu_ticks[C.CPU_STATE_SYSTEM]) + uint64(cpuLoad.cpu_ticks[C.CPU_STATE_NICE])

	return
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

	return res
}