package kgo

import (
	"fmt"
	"reflect"
	"strconv"
)

// Struct2Map 结构体转为字典;tagName为要导出的标签名,可以为空,为空时将导出所有字段.
func (kc *LkkConvert) Struct2Map(obj interface{}, tagName string) (map[string]interface{}, error) {
	return struct2Map(obj, tagName)
}

// Int2Str 将整数转换为字符串.
func (kc *LkkConvert) Int2Str(val interface{}) string {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	default:
		r := reflect.ValueOf(val)
		switch r.Kind() {
		case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
			return fmt.Sprintf("%d", r.Int())
		default:
			return ""
		}
	}
}

// Float2Str 将浮点数转换为字符串,decimal为小数位数.
func (kc *LkkConvert) Float2Str(val interface{}, decimal int) string {
	if decimal <= 0 {
		decimal = 2
	}

	switch val.(type) {
	case float32:
		return strconv.FormatFloat(float64(val.(float32)), 'f', decimal, 32)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', decimal, 64)
	default:
		r := reflect.ValueOf(val)
		switch r.Kind() {
		case reflect.Float32:
			return strconv.FormatFloat(r.Float(), 'f', decimal, 32)
		case reflect.Float64:
			return strconv.FormatFloat(r.Float(), 'f', decimal, 64)
		default:
			return ""
		}
	}
}

// Bool2Str 将布尔值转换为字符串.
func (kc *LkkConvert) Bool2Str(val bool) string {
	if val {
		return "true"
	}
	return "false"
}

// Bool2Int 将布尔值转换为整型.
func (kc *LkkConvert) Bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

// Str2Int 将字符串转换为int.其中"true", "TRUE", "True"为1;若为浮点字符串,则取整数部分.
func (kc *LkkConvert) Str2Int(val string) int {
	return str2Int(val)
}

// Str2Int8 将字符串转换为int8.
func (kc *LkkConvert) Str2Int8(val string) int8 {
	res, _ := strconv.ParseInt(val, 0, 8)
	return int8(res)
}

// Str2Int16 将字符串转换为int16.
func (kc *LkkConvert) Str2Int16(val string) int16 {
	res, _ := strconv.ParseInt(val, 0, 16)
	return int16(res)
}

// Str2Int32 将字符串转换为int32.
func (kc *LkkConvert) Str2Int32(val string) int32 {
	res, _ := strconv.ParseInt(val, 0, 32)
	return int32(res)
}

// Str2Int64 将字符串转换为int64.
func (kc *LkkConvert) Str2Int64(val string) int64 {
	res, _ := strconv.ParseInt(val, 0, 64)
	return res
}

// Str2Uint 将字符串转换为uint.其中"true", "TRUE", "True"为1;若为浮点字符串,则取整数部分;若为负值则取0.
func (kc *LkkConvert) Str2Uint(val string) uint {
	return str2Uint(val)
}

// Str2Uint8 将字符串转换为uint8.
func (kc *LkkConvert) Str2Uint8(val string) uint8 {
	res, _ := strconv.ParseUint(val, 0, 8)
	return uint8(res)
}

// Str2Uint16 将字符串转换为uint16.
func (kc *LkkConvert) Str2Uint16(val string) uint16 {
	res, _ := strconv.ParseUint(val, 0, 16)
	return uint16(res)
}

// Str2Uint32 将字符串转换为uint32.
func (kc *LkkConvert) Str2Uint32(val string) uint32 {
	res, _ := strconv.ParseUint(val, 0, 32)
	return uint32(res)
}

// Str2Uint64 将字符串转换为uint64.
func (kc *LkkConvert) Str2Uint64(val string) uint64 {
	res, _ := strconv.ParseUint(val, 0, 64)
	return res
}
