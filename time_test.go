package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTime_UnixTime(t *testing.T) {
	var res int64

	res = KTime.UnixTime()
	assert.Equal(t, 10, len(toStr(res)))
}

func BenchmarkTime_UnixTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.UnixTime()
	}
}

func TestTime_MilliTime(t *testing.T) {
	var res int64

	res = KTime.MilliTime()
	assert.Equal(t, 13, len(toStr(res)))
}

func BenchmarkTime_MilliTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.MilliTime()
	}
}

func TestTime_MicroTime(t *testing.T) {
	var res int64

	res = KTime.MicroTime()
	assert.Equal(t, 16, len(toStr(res)))
}

func BenchmarkTime_MicroTime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.MicroTime()
	}
}

func TestTime_Str2Timestruct(t *testing.T) {
	var res time.Time
	var err error

	res, err = KTime.Str2Timestruct(strTime1)
	assert.Nil(t, err)
	assert.Equal(t, res.Year(), 2019)
	assert.Equal(t, int(res.Month()), 7)
	assert.Equal(t, res.Day(), 11)
}

func BenchmarkTime_Str2Timestruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.Str2Timestruct(strTime1)
	}
}
