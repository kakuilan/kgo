package gohelper

import (
	"fmt"
	"testing"
)

func TestTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.Time())
	if len(ti) != 10 {
		t.Error("Time fail")
		return
	}
}

func BenchmarkTime(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		KTime.Time()
	}
}

func TestMilliTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.MilliTime())
	if len(ti) != 13 {
		t.Error("MilliTime fail")
		return
	}
}

func BenchmarkMilliTime(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		KTime.MilliTime()
	}
}

func TestMicroTime(t *testing.T) {
	ti := fmt.Sprintf("%d", KTime.MicroTime())
	if len(ti) != 16 {
		t.Error("MicroTime fail")
		return
	}
}

func BenchmarkMicroTime(b *testing.B) {
	b.ResetTimer()
	for i:=0;i<b.N;i++{
		KTime.MicroTime()
	}
}