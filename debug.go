package kgo

import "fmt"

// DumpPrint 调试打印变量.
func (ks *LkkDebug) DumpPrint(v interface{}) {
	fmt.Printf("%+v\n", v)
}
