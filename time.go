package gohelper

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
