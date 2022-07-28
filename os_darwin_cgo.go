//go:build darwin && cgo
// +build darwin,cgo

package kgo

// #include <stdlib.h>
// #include <libproc.h>
// #include <mach/mach_host.h>
// #include <mach/host_info.h>
import "C"
import (
	"unsafe"
)

// 获取CPU使用率(darwin系统必须使用cgo),单位jiffies(节拍数).
// user为用户态(用户进程)的运行时间,
// idle为空闲时间,
// total为累计时间.
func (ko *LkkOS) CpuUsage() (user, idle, total uint64) {
	var cpuLoad C.host_cpu_load_info_data_t
	var count C.mach_msg_type_number_t = C.HOST_CPU_LOAD_INFO_COUNT
	ret := C.host_statistics(C.host_t(C.mach_host_self()), C.HOST_CPU_LOAD_INFO, C.host_info_t(unsafe.Pointer(&cpuLoad)), &count)
	if ret == C.KERN_SUCCESS {
		user = uint64(cpuLoad.cpu_ticks[C.CPU_STATE_USER])
		idle = uint64(cpuLoad.cpu_ticks[C.CPU_STATE_IDLE])
		total = user + idle + uint64(cpuLoad.cpu_ticks[C.CPU_STATE_SYSTEM]) + uint64(cpuLoad.cpu_ticks[C.CPU_STATE_NICE])
	}

	return
}

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) (res string) {
	var c C.char // need a var for unsafe.Sizeof need a var
	const bufsize = C.PROC_PIDPATHINFO_MAXSIZE * unsafe.Sizeof(c)
	buffer := (*C.char)(C.malloc(C.size_t(bufsize)))
	defer C.free(unsafe.Pointer(buffer))

	ret, err := C.proc_pidpath(C.int(pid), unsafe.Pointer(buffer), C.uint32_t(bufsize))
	if err == nil && ret > 0 {
		res = C.GoString(buffer)
	}

	return
}
