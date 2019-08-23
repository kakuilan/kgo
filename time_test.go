package kgo

import (
	"fmt"
	"testing"
	"time"
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
	for i := 0; i < b.N; i++ {
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
	for i := 0; i < b.N; i++ {
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
	for i := 0; i < b.N; i++ {
		KTime.MicroTime()
	}
}

func TestStrtotime(t *testing.T) {
	ti, err := KTime.Strtotime("2019-07-11 10:11:23")
	if err != nil || ti <= 0 {
		t.Error("Strtotime fail")
		return
	}

	_, err2 := KTime.Strtotime("02/01/2016 15:04:05")
	if err2 == nil {
		t.Error("Strtotime fail")
		return
	}
}

func BenchmarkStrtotime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.Strtotime("2019-07-11 10:11:23")
	}
}

func TestDate(t *testing.T) {
	date1 := KTime.Date("Y-m-d H:i:s", 1562811851)
	date2 := KTime.Date("y-n-j H:i:s", int64(1562811851))
	date3 := KTime.Date("m/d/y h-i-s", time.Now())
	if date1 == "" || date2 == "" || date3 == "" {
		t.Error("Date fail")
		return
	}

	date4 := KTime.Date("Y-m-d H:i:s")
	date5 := KTime.Date("Y-m-d H:i:s", "hello")
	if date4 == "" || date5 != "" {
		t.Error("Date fail")
		return
	}
}

func BenchmarkDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Date("Y-m-d H:i:s", 1562811851)
	}
}

func TestCheckDate(t *testing.T) {
	chk1 := KTime.CheckDate(7, 31, 2019)
	chk2 := KTime.CheckDate(2, 31, 2019)
	if !chk1 || chk2 {
		t.Error("CheckDate fail")
		return
	}
	KTime.CheckDate(0, 31, 2019)
	KTime.CheckDate(4, 31, 2019)
	KTime.CheckDate(2, 30, 2008)
}

func BenchmarkCheckDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.CheckDate(7, 31, 2019)
	}
}

func TestSleep(t *testing.T) {
	ti1 := KTime.Time()
	KTime.Sleep(1)
	ti2 := KTime.Time()
	diff := ti2 - ti1
	if diff != 1 {
		t.Error("Sleep fail")
		return
	}
}

func TestUsleep(t *testing.T) {
	ti1 := KTime.MicroTime()
	KTime.Usleep(100)
	ti2 := KTime.MicroTime()
	diff := ti2 - ti1
	if diff < 100 {
		t.Error("Usleep fail")
		return
	}
}

func TestServiceStartime(t *testing.T) {
	res := KTime.ServiceStartime()
	println(KTime.Date("Y-m-d H:i:s", res))
	if res <= 0 {
		t.Error("ServiceStartime fail")
		return
	}
}

func BenchmarkServiceStartime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceStartime()
	}
}

func TestServiceUptime(t *testing.T) {
	res := KTime.ServiceUptime()
	println("ServiceUptime", res)
	if res <= 0 {
		t.Error("ServiceUptime fail")
		return
	}
}

func BenchmarkServiceUptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceUptime()
	}
}
