package kgo

import "fmt"

// DumpPrint 调试打印变量.
func DumpPrint(v interface{}) {
	fmt.Printf("%+v\n", v)
}
