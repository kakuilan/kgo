package kgo

import (
	"fmt"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionName 获取调用方法的名称;f为目标方法;onlyFun为true时仅返回方法,不包括包名
func (kd *LkkDebug) GetFuncName(f interface{}, onlyFun ...bool) string {
	var funcObj *runtime.Func
	if f == nil {
		// Skip this function, and fetch the PC and file for its parent
		pc, _, _, _ := runtime.Caller(1)
		// Retrieve a Function object this functions parent
		funcObj = runtime.FuncForPC(pc)
	} else {
		funcObj = runtime.FuncForPC(reflect.ValueOf(f).Pointer())
	}

	name := funcObj.Name()
	if len(onlyFun) > 0 && onlyFun[0] == true {
		// extract just the function name (and not the module path)
		return strings.TrimPrefix(filepath.Ext(name), ".")
	}

	return name
}

// GetFuncLine 获取调用方法的行号
func (kd *LkkDebug) GetFuncLine() int {
	// Skip this function, and fetch the PC and file for its parent
	_, _, line, _ := runtime.Caller(1)
	return line
}

// GetFuncFile 获取调用方法的文件路径
func (kd *LkkDebug) GetFuncFile() string {
	_, file, _, _ := runtime.Caller(1)
	return file
}

// GetFuncDir 获取调用方法的文件目录
func (kd *LkkDebug) GetFuncDir() string {
	return filepath.Dir(kd.GetFuncFile())
}

// GetFuncPackage 获取调用方法或源文件的包名.funcFile为源文件路径.
func (kd *LkkDebug) GetFuncPackage(funcFile ...string) string {
	var sourceFile string
	if len(funcFile) == 0 {
		sourceFile = kd.GetFuncFile()
	} else {
		sourceFile = funcFile[0]
	}

	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, sourceFile, nil, parser.PackageClauseOnly)
	if err != nil || astFile.Name == nil {
		return ""
	}

	return astFile.Name.Name
}

// DumpStacks 打印堆栈信息
func (kd *LkkDebug) DumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)
}
