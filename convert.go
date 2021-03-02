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
