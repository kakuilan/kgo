package kgo

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// DumpPrint 打印调试变量.
func (ks *LkkDebug) DumpPrint(vs ...interface{}) {
	dumpPrint(vs)
}

// DumpStacks 打印堆栈信息.
func (kd *LkkDebug) DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===\n\n", buf)
}

// GetCallName 获取调用的方法名称;f为目标方法;onlyFun为true时仅返回方法,不包括包名.
func (kd *LkkDebug) GetCallName(f interface{}, onlyFun bool) string {
	var funcObj *runtime.Func
	r := reflectPtr(reflect.ValueOf(f))
	switch r.Kind() {
	case reflect.Invalid:
		// Skip this function, and fetch the PC and file for its parent
		pc, _, _, _ := runtime.Caller(1)
		// Retrieve a Function object this functions parent
		funcObj = runtime.FuncForPC(pc)
	case reflect.Func:
		funcObj = runtime.FuncForPC(r.Pointer())
	default:
		return ""
	}

	name := funcObj.Name()
	if onlyFun {
		// extract just the function name (and not the module path)
		return strings.TrimPrefix(filepath.Ext(name), ".")
	}

	return name
}

// GetCallFile 获取调用方法的文件路径.
func (kd *LkkDebug) GetCallFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// GetCallLine 获取调用方法的行号.
func (kd *LkkDebug) GetCallLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}
