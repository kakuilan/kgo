package kgo

import (
	"fmt"
	"runtime"
)

// DumpPrint 打印调试变量.
func (ks *LkkDebug) DumpPrint(vs ...interface{}) {
	dumpPrint(vs)
}

// DumpStacks 打印堆栈信息.
func (kd *LkkDebug) DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
