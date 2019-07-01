package gohelper

import (
	"testing"
)

func TestNl2br(t *testing.T) {
	str := `hello
world!
你好！`
	res := KStr.Nl2br(str)
	println(res)
}

func BenchmarkNl2br(b *testing.B) {
	b.ResetTimer()
	str := `hello
world!
你好！`
	for i:=0;i<b.N;i++{
		_ = KStr.Nl2br(str)
	}
}
