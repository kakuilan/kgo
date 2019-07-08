package gohelper

import (
	"testing"
)

func TestIsWindows(t *testing.T) {
	res := KOS.IsWindows()
	if res {
		t.Error("IsWindows fail")
		return
	}
}

func BenchmarkIsWindows(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsWindows()
	}
}

func TestIsLinux(t *testing.T) {
	res := KOS.IsLinux()
	if !res {
		t.Error("IsLinux fail")
		return
	}
}

func BenchmarkIsLinux(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsLinux()
	}
}

func TestIsMac(t *testing.T) {
	res := KOS.IsMac()
	if res {
		t.Error("IsMac fail")
		return
	}
}

func BenchmarkIsMac(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.IsMac()
	}
}

func TestPwd(t *testing.T) {
	res := KOS.Pwd()
	if res == "" {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkPwd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.Pwd()
	}
}

func TestHomeDir(t *testing.T) {
	_, err := KOS.HomeDir()
	if err != nil {
		t.Error("Pwd fail")
		return
	}
}

func BenchmarkHomeDir(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KOS.HomeDir()
	}
}

func TestHomeUnix(t *testing.T) {
	_, err := homeUnix()
	if err != nil {
		t.Error("homeUnix fail")
		return
	}
}

func TestHomeWindows(t *testing.T) {
	_, err := homeWindows()
	if err == nil {
		t.Error("homeWindows fail")
		return
	}
}
