package kgo

// DumpPrint 打印调试变量.
func (ks *LkkDebug) DumpPrint(vs ...interface{}) {
	dumpPrint(vs)
}
