package kgo

import "fmt"

// DumpPrint 打印调试变量.
func (ks *LkkDebug) DumpPrint(v interface{}) {
	fmt.Printf("%+v\n", v)
}
