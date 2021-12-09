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

	res = KTime.Date("m/d/y h-i-s", kuptime)
	assert.NotEmpty(t, res)

	res = KTime.Date("Y-m-d H:i:s")
	assert.NotEmpty(t, res)

	res = KTime.Date("Y-m-d H:i:s", strHello)
	assert.Empty(t, res)

	//时间戳为0的时间点
	res = KTime.Date("Y-m-d H:i:s", 0)
	assert.NotEmpty(t, res)
	//东八区
	//assert.Equal(t, res, "1970-01-01 08:00:00")
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

func TestTime_Sleep(t *testing.T) {
	var t0, t1, t2, res int64
	t0 = 1
	t1 = KTime.UnixTime()
	KTime.Sleep(t0)
	t2 = KTime.UnixTime()
	res = t2 - t1
	assert.GreaterOrEqual(t, res, t0)
}

func TestTime_Usleep(t *testing.T) {
	var t0, t1, t2, res int64
	t0 = 100
	t1 = KTime.MicroTime()
	KTime.Usleep(t0)
	t2 = KTime.MicroTime()
	res = t2 - t1
	assert.GreaterOrEqual(t, res, t0)
}

func TestTime_ServiceStartime(t *testing.T) {
	var res int64
	res = KTime.ServiceStartime()
	assert.Greater(t, res, int64(1))
}

func BenchmarkTime_ServiceStartime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceStartime()
	}
}

func TestTime_ServiceUptime(t *testing.T) {
	var res time.Duration

	res = KTime.ServiceUptime()
	assert.Greater(t, int64(res), int64(1))
}

func BenchmarkTime_ServiceUptime(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.ServiceUptime()
	}
}

func TestTime_GetMonthDays(t *testing.T) {
	var res int
	var tests = []struct {
		month    int
		year     int
		expected int
	}{
		{3, 1970, 31},
		{1, 2009, 31},
		{0, 2009, 0},
		{2, 2009, 28},
		{2, 2016, 29},
		{2, 1900, 28},
		{4, 2020, 30},
		{2, 1600, 29},
	}
	for _, test := range tests {
		actual := KTime.GetMonthDays(test.month, test.year)
		assert.Equal(t, actual, test.expected)
	}

	res = KTime.GetMonthDays(2)
	assert.Greater(t, res, 0)
}

func BenchmarkTime_GetMonthDays(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.GetMonthDays(3, 1970)
	}
}

func TestTime_YearMonthDay(t *testing.T) {
	var y1, m1, d1, y2, m2, d2 int

	y1 = KTime.Year()
	m1 = KTime.Month()
	d1 = KTime.Day()

	y2 = KTime.Year(kuptime)
	m2 = KTime.Month(kuptime)
	d2 = KTime.Day(kuptime)

	assert.Equal(t, y1, y2)
	assert.Equal(t, m1, m2)
	assert.Equal(t, d1, d2)
}

func BenchmarkTime_Year(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Year(kuptime)
	}
}

func BenchmarkTime_Month(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Month(kuptime)
	}
}

func BenchmarkTime_Day(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Year(kuptime)
	}
}

func TestTime_HourMinuteSecond(t *testing.T) {
	var h1, m1, s1, h2, m2, s2 int

	h1 = KTime.Hour()
	m1 = KTime.Minute()
	s1 = KTime.Second()
	assert.GreaterOrEqual(t, h1, 0)
	assert.GreaterOrEqual(t, m1, 0)
	assert.GreaterOrEqual(t, s1, 0)

	h2 = KTime.Hour(myDate1)
	m2 = KTime.Minute(myDate1)
	s2 = KTime.Second(myDate1)
	assert.Equal(t, h2, 23)
	assert.Equal(t, m2, 4)
	assert.Equal(t, s2, 35)
}

func BenchmarkTime_Hour(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Hour(kuptime)
	}
}

func BenchmarkTime_Minute(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Minute(kuptime)
	}
}

func BenchmarkTime_Second(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.Second(kuptime)
	}
}

func TestTime_StartOfDay(t *testing.T) {
	var res time.Time

	res = KTime.StartOfDay(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-10 00:00:00")
}

func BenchmarkTime_StartOfDay(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.StartOfDay(myDate1)
	}
}

func TestTime_EndOfDay(t *testing.T) {
	var res time.Time

	res = KTime.EndOfDay(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-10 23:59:59")
}

func BenchmarkTime_EndOfDay(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.EndOfDay(myDate1)
	}
}

func TestTime_StartOfMonth(t *testing.T) {
	var res time.Time

	res = KTime.StartOfMonth(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-01 00:00:00")
}

func BenchmarkTime_StartOfMonth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.StartOfMonth(myDate1)
	}
}

func TestTime_EndOfMonth(t *testing.T) {
	var res time.Time

	res = KTime.EndOfMonth(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-31 23:59:59")
}

func BenchmarkTime_EndOfMonth(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.EndOfMonth(myDate1)
	}
}

func TestTime_StartOfYear(t *testing.T) {
	var res time.Time

	res = KTime.StartOfYear(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-01-01 00:00:00")
}

func BenchmarkTime_StartOfYear(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.StartOfYear(myDate1)
	}
}

func TestTime_EndOfYear(t *testing.T) {
	var res time.Time

	res = KTime.EndOfYear(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-12-31 23:59:59")
}

func BenchmarkTime_EndOfYear(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.EndOfYear(myDate1)
	}
}

func TestTime_StartOfWeek(t *testing.T) {
	var res time.Time

	res = KTime.StartOfWeek(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-09 00:00:00")

	res = KTime.StartOfWeek(myDate2, time.Tuesday)
	str = KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-03 00:00:00")
}

func BenchmarkTime_StartOfWeek(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.StartOfWeek(myDate1)
	}
}

func TestTime_EndOfWeek(t *testing.T) {
	var res time.Time

	res = KTime.EndOfWeek(myDate1)
	str := KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-15 23:59:59")

	res = KTime.EndOfWeek(myDate2, time.Tuesday)
	str = KTime.Date("Y-m-d H:i:s", res)
	assert.Equal(t, str, "2020-03-09 23:59:59")
}

func BenchmarkTime_EndOfWeek(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.EndOfWeek(myDate1)
	}
}

func TestTime_DaysBetween(t *testing.T) {
	days := KTime.DaysBetween(myDate1, myDate3)
	assert.Equal(t, days, 107)

	days = KTime.DaysBetween(myDate3, myDate1)
	assert.Equal(t, days, -107)
}

func BenchmarkTime_DaysBetween(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KTime.DaysBetween(myDate1, myDate3)
	}
}

func TestTime_IsDate2time(t *testing.T) {
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"hello", false},
		{"0000", true},
		{"1970", true},
		{"1970-01-01", true},
		{"1970-01-01 00:00:01", true},
		{"1971-01-01 00:00:01", true},
		{"1990-01", true},
		{"1990/01", true},
		{"1990-01-02", true},
		{"1990/01/02", true},
		{"1990-01-02 03", true},
		{"1990/01/02 03", true},
		{"1990-01-02 03:14", true},
		{"1990/01/02 03:14", true},
		{"1990-01-02 03:14:59", true},
		{"1990/01/02 03:14:59", true},
		{"2990-00-00 03:14:59", false},
	}
	for _, test := range tests {
		actual, tim := KTime.IsDate2time(test.param)
		assert.Equal(t, actual, test.expected)
		if actual {
			if toInt(test.param) > 1970 {
				assert.Greater(t, tim, int64(0))
			}
		}
	}
}

func BenchmarkTime_IsDate2time(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KTime.IsDate2time(strTime7)
	}
}
