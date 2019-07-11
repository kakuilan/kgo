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
	date1 := KTime.Date(1562811851, "2006-01-02 15:04:05")
	date2 := KTime.Date(1562811851, "02/01/2006 15:04:05")
	if date1 == "" || date2 == "" {
		t.Error("Date fail")
		return
	}
}

func BenchmarkDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Date(1562811851, "2006-01-02 15:04:05")
	}
}

func TestCheckdate(t *testing.T) {
	chk1 := KTime.Checkdate(7, 31, 2019)
	chk2 := KTime.Checkdate(2, 31, 2019)
	if !chk1 || chk2 {
		t.Error("Checkdate fail")
		return
	}
}

func BenchmarkCheckdate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Checkdate(7, 31, 2019)
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
