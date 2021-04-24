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

func TestTime_Str2Timestamp(t *testing.T) {
	var res int64
	var err error

	res, err = KTime.Str2Timestamp(strTime1)
	assert.Nil(t, err)
	assert.Greater(t, res, int64(1))

	res, err = KTime.Str2Timestamp(strTime3, "01/02/2006 15:04:05")
	assert.Nil(t, err)
	assert.Greater(t, res, int64(1))

	//时间格式错误
	res, err = KTime.Str2Timestamp(strTime2, "2006-01-02")
	assert.NotNil(t, err)
}

func BenchmarkTime_Str2Timestamp(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.Str2Timestamp(strTime3, "01/02/2006 15:04:05")
	}
}

func TestTime_Date(t *testing.T) {
	var res string

	res = KTime.Date("Y-m-d H:i:s")
	assert.NotEmpty(t, res)

	res = KTime.Date("Y-m-d H:i:s", intTime1)
	assert.NotEmpty(t, res)

	res = KTime.Date("y-n-j H:i:s", int64(intTime1))
	assert.NotEmpty(t, res)

	res = KTime.Date("m/d/y h-i-s", Kuptime)
	assert.NotEmpty(t, res)

	res = KTime.Date("Y-m-d H:i:s")
	assert.NotEmpty(t, res)

	res = KTime.Date("Y-m-d H:i:s", strHello)
	assert.Empty(t, res)
}

func BenchmarkTime_Date(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Date("Y-m-d H:i:s")
	}
}

func TestTime_CheckDate(t *testing.T) {
	var res bool

	res = KTime.CheckDate(2019, 7, 31)
	assert.True(t, res)

	res = KTime.CheckDate(2019, 2, 31)
	assert.False(t, res)

	res = KTime.CheckDate(2019, 0, 31)
	assert.False(t, res)

	res = KTime.CheckDate(2019, 4, 31)
	assert.False(t, res)

	res = KTime.CheckDate(2008, 2, 30)
	assert.False(t, res)
}

func BenchmarkTime_CheckDate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.CheckDate(2019, 7, 31)
	}
}
