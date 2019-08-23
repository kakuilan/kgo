package kgo

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

// GetFunctionName 获取调用方法的名称;f为目标方法;onlyFun为true时仅返回方法,不包括包名
func (kd *LkkDebug) GetFunctionName(f interface{}, onlyFun ...bool) string {
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
