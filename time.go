package kgo

import (
	"errors"
	"strings"
	"time"
)

// DateFormat pattern rules.
var datePatterns = []string{
	// year
	"Y", "2006", // A full numeric representation of a year, 4 digits   Examples: 1999 or 2003
	"y", "06", // A two digit representation of a year   Examples: 99 or 03

	// month
	"m", "01", // Numeric representation of a month, with leading zeros 01 through 12
	"n", "1", // Numeric representation of a month, without leading zeros   1 through 12
	"M", "Jan", // A short textual representation of a month, three letters Jan through Dec
	"F", "January", // A full textual representation of a month, such as January or March   January through December

	// day
	"d", "02", // Day of the month, 2 digits with leading zeros 01 to 31
	"j", "2", // Day of the month without leading zeros 1 to 31

	// week
	"D", "Mon", // A textual representation of a day, three letters Mon through Sun
	"l", "Monday", // A full textual representation of the day of the week  Sunday through Saturday

	// time
	"g", "3", // 12-hour format of an hour without leading zeros    1 through 12
	"G", "15", // 24-hour format of an hour without leading zeros   0 through 23
	"h", "03", // 12-hour format of an hour with leading zeros  01 through 12
	"H", "15", // 24-hour format of an hour with leading zeros  00 through 23

	"a", "pm", // Lowercase Ante meridiem and Post meridiem am or pm
	"A", "PM", // Uppercase Ante meridiem and Post meridiem AM or PM

	"i", "04", // Minutes with leading zeros    00 to 59
	"s", "05", // Seconds, with leading zeros   00 through 59

	// time zone
	"T", "MST",
	"P", "-07:00",
	"O", "-0700",

	// RFC 2822
	"r", time.RFC1123Z,
}

// UnixTime 获取当前Unix时间戳(秒,10位).
func (kt *LkkTime) UnixTime() int64 {
	return time.Now().Unix()
}

// MilliTime 获取当前Unix时间戳(毫秒,13位).
func (kt *LkkTime) MilliTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MicroTime 获取当前Unix时间戳(微秒,16位).
func (kt *LkkTime) MicroTime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// Str2Timestruct 将字符串转换为时间结构.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05" .
func (kt *LkkTime) Str2Timestruct(str string, format ...string) (time.Time, error) {
	var f string
	if len(format) > 0 {
		f = strings.Trim(format[0], " ")
	} else {
		f = "2006-01-02 15:04:05"
	}

	if len(str) != len(f) {
		return time.Now(), errors.New("[Str2Timestruct]`format error")
	}

	return time.ParseInLocation(f, str, kuptime.Location())
}

// Str2Timestamp 将字符串转换为时间戳,秒.
// str 为要转换的字符串;
// format 为该字符串的格式,默认为"2006-01-02 15:04:05" .
func (kt *LkkTime) Str2Timestamp(str string, format ...string) (int64, error) {
	tim, err := kt.Str2Timestruct(str, format...)
	if err != nil {
		return 0, err
	}

	return tim.Unix(), nil
}

// Date 格式化时间.
// format 格式,如"Y-m-d H:i:s".
// ts为int/int64类型时间戳或time.Time类型.
func (kt *LkkTime) Date(format string, ts ...interface{}) string {
	replacer := strings.NewReplacer(datePatterns...)
	format = replacer.Replace(format)

	var t time.Time
	if len(ts) > 0 {
		val := ts[0]
		if v, ok := val.(time.Time); ok {
			t = v
		} else if v, ok := val.(int); ok {
			t = time.Unix(int64(v), 0)
		} else if v, ok := val.(int64); ok {
			t = time.Unix(v, 0)
		} else {
			return ""
		}
	} else {
		t = time.Now()
	}

	return t.Format(format)
}

// CheckDate 检查是否正常的日期.
func (kt *LkkTime) CheckDate(year, month, day int) bool {
	if month < 1 || month > 12 || day < 1 || day > 31 || year < 1 || year > 32767 {
		return false
	}
	switch month {
	case 4, 6, 9, 11:
		if day > 30 {
			return false
		}
	case 2:
		// leap year
		if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
			if day > 29 {
				return false
			}
		} else if day > 28 {
			return false
		}
	}

	return true
}

// Sleep 延缓执行,秒.
func (kt *LkkTime) Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep 以指定的微秒数延迟执行.
func (kt *LkkTime) Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}

// ServiceStartime 获取当前服务启动时间戳,秒.
func (kt *LkkTime) ServiceStartime() int64 {
	return kuptime.Unix()
}

// ServiceUptime 获取当前服务运行时间,纳秒int64.
func (kt *LkkTime) ServiceUptime() time.Duration {
	return time.Since(kuptime)
}

// GetMonthDays 获取指定月份的天数.year年份,可选,默认当前年份.
func (kt *LkkTime) GetMonthDays(month int, year ...int) (days int) {
	if month < 1 || month > 12 {
		return
	}

	var yr int
	if len(year) == 0 {
		yr = time.Now().Year()
	} else {
		yr = year[0]
	}

	if month != 2 {
		if month == 4 || month == 6 || month == 9 || month == 11 {
			days = 30
		} else {
			days = 31
		}
	} else {
		if ((yr%4) == 0 && (yr%100) != 0) || (yr%400) == 0 {
			days = 29
		} else {
			days = 28
		}
	}

	return
}

// Year 获取年份.
func (kt *LkkTime) Year(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return tm.Year()
}

// Month 获取月份.
func (kt *LkkTime) Month(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return int(tm.Month())
}

// Day 获取日份.
func (kt *LkkTime) Day(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return tm.Day()
}

// Hour 获取小时.
func (kt *LkkTime) Hour(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return tm.Hour()
}

// Minute 获取分钟.
func (kt *LkkTime) Minute(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return tm.Minute()
}

// Second 获取秒数.
func (kt *LkkTime) Second(t ...time.Time) int {
	var tm time.Time
	if len(t) > 0 {
		tm = t[0]
	} else {
		tm = time.Now()
	}
	return tm.Second()
}

// StartOfDay 获取日期中当天的开始时间.
func (kt *LkkTime) StartOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

// EndOfDay 获取日期中当天的结束时间.
func (kt *LkkTime) EndOfDay(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, int(time.Second-time.Nanosecond), date.Location())
}

// StartOfMonth 获取日期中当月的开始时间.
func (kt *LkkTime) StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

// EndOfMonth 获取日期中当月的结束时间.
func (kt *LkkTime) EndOfMonth(date time.Time) time.Time {
	return kt.StartOfMonth(date).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// StartOfYear 获取日期中当年的开始时间.
func (kt *LkkTime) StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), 1, 1, 0, 0, 0, 0, date.Location())
}

// EndOfYear 获取日期中当年的结束时间.
func (kt *LkkTime) EndOfYear(date time.Time) time.Time {
	return kt.StartOfYear(date).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// StartOfWeek 获取日期中当周的开始时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func (kt *LkkTime) StartOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	weekstart := time.Monday
	if len(weekStartDay) > 0 {
		weekstart = weekStartDay[0]
	}

	// 当前是周几
	weekday := int(date.Weekday())
	if weekstart != time.Sunday {
		weekStartDayInt := int(weekstart)

		if weekday < weekStartDayInt {
			weekday = weekday + 7 - weekStartDayInt
		} else {
			weekday = weekday - weekStartDayInt
		}
	}

	return time.Date(date.Year(), date.Month(), date.Day()-weekday, 0, 0, 0, 0, date.Location())
}

// EndOfWeek 获取日期中当周的结束时间;
// weekStartDay 周几作为周的第一天,本库默认周一.
func (kt *LkkTime) EndOfWeek(date time.Time, weekStartDay ...time.Weekday) time.Time {
	return kt.StartOfWeek(date, weekStartDay...).AddDate(0, 0, 7).Add(-time.Nanosecond)
}

// DaysBetween 获取两个日期的间隔天数.
// fromDate 为开始时间,toDate 为终点时间.
func (kt *LkkTime) DaysBetween(fromDate, toDate time.Time) int {
	return int(toDate.Sub(fromDate) / (24 * time.Hour))
}

// IsDate2time 检查字符串是否日期格式,并转换为时间戳.注意,时间戳可能为负数(小于1970年时).
// 匹配如:
//
//	0000
//	0000-00
//	0000/00
//	0000-00-00
//	0000/00/00
//	0000-00-00 00
//	0000/00/00 00
//	0000-00-00 00:00
//	0000/00/00 00:00
//	0000-00-00 00:00:00
//	0000/00/00 00:00:00
//
// 等日期格式.
func (kt *LkkTime) IsDate2time(str string) (bool, int64) {
	if str == "" {
		return false, 0
	} else if strings.ContainsRune(str, '/') {
		str = strings.Replace(str, "/", "-", -1)
	}

	chk := RegDatetime.MatchString(str)
	if !chk {
		return false, 0
	}

	leng := len(str)
	if leng < 19 {
		reference := "1970-01-01 00:00:00"
		str = str + reference[leng:19]
	}

	tim, err := KTime.Str2Timestamp(str)
	if err != nil {
		return false, 0
	}

	return true, tim
}
