package gohelper

import (
	"fmt"
	"strconv"
)

// 将整数转换为字符串
func (kc *LkkConvert) Int2Str(val interface{}) string {
	switch val.(type) {
	// Integers
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	// Type is not integers, return empty string
	default:
		return ""
	}
}

// 将浮点数转换为字符串,length为小数位数
func (kc *LkkConvert) Float2Str(val interface{}, length int) string {
	switch val.(type) {
	// Floats
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', length, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', length, 64)
	// Type is not floats, return empty string
	default:
		return ""
	}
}
