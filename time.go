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

// Time 获取当前Unix时间戳(秒,10位).
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

// Str2Time 将字符串转换为时间结构.
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

	return time.ParseInLocation(f, str, Kuptime.Location())
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
			t = time.Unix(int64(v), 0)
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
	return Kuptime.Unix()
}
