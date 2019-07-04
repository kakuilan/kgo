package gohelper

import (
	"fmt"
	"strconv"
)

// Int2Str 将整数转换为字符串
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

// Float2Str 将浮点数转换为字符串,length为小数位数
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

// Bool2Str 将布尔值转换为字符串
func (kc *LkkConvert) Bool2Str(val bool) string {
	if val {
		return "true"
	}
	return "false"
}

// StrictStr2Int 严格将字符串转换为有符号整型,bitSize为整型位数,strict为是否严格检查
func (kc *LkkConvert) StrictStr2Int(val string, bitSize int, strict bool) int64 {
	res, err := strconv.ParseInt(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// Str2Int 将字符串转换为int
func (kc *LkkConvert) Str2Int(val string) int {
	res, _ := strconv.Atoi(val)
	return res
}

// Str2Int8 将字符串转换为int8
func (kc *LkkConvert) Str2Int8(val string) int8 {
	return int8(kc.StrictStr2Int(val, 8, false))
}

// Str2Int16 将字符串转换为int16
func (kc *LkkConvert) Str2Int16(val string) int16 {
	return int16(kc.StrictStr2Int(val, 16, false))
}

// Str2Int32 将字符串转换为int32
func (kc *LkkConvert) Str2Int32(val string) int32 {
	return int32(kc.StrictStr2Int(val, 32, false))
}

// Str2Int64 将字符串转换为int64
func (kc *LkkConvert) Str2Int64(val string) int64 {
	return kc.StrictStr2Int(val, 64, false)
}
