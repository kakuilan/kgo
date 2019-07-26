package kgo

import "time"

// Time 获取当前Unix时间戳(秒)
func (kt *LkkTime) Time() int64 {
	return time.Now().Unix()
}

// MilliTime 获取当前Unix时间戳(毫秒)
func (kt *LkkTime) MilliTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// MicroTime 获取当前Unix时间戳(微秒)
func (kt *LkkTime) MicroTime() int64 {
	return time.Now().UnixNano() / int64(time.Microsecond)
}

// Strtotime 字符串转时间戳
// 例如KTime.Strtotime("2019-07-11 10:11:23") == 1562839883
func (kt *LkkTime) Strtotime(str string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}

// Date 格式化时间戳
// 例如KTime.Date(1562811851, "2006-01-02 15:04:05") == "2019-07-11 10:24:11"
func (kt *LkkTime) Date(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

// CheckDate 检查是否正常的日期
func (kt *LkkTime) CheckDate(month, day, year int) bool {
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

// Sleep 延缓执行,秒
func (kt *LkkTime) Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// Usleep 以指定的微秒数延迟执行
func (kt *LkkTime) Usleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
