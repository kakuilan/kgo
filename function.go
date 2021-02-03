package kgo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

// isArrayOrSlice 检查变量是否数组或切片.
// chkType为检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是(数组长度).
func isArrayOrSlice(val interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		panic(fmt.Sprintf("[isArrayOrSlice]`chkType refValue must in [1, 2, 3]; but: %d", chkType))
	}

	var res = -1
	refVal := reflect.ValueOf(val)
	switch refVal.Kind() {
	case reflect.Array:
		if chkType == 1 || chkType == 3 {
			res = refVal.Len()
		}
	case reflect.Slice:
		if chkType == 2 || chkType == 3 {
			res = refVal.Len()
		}
	}

	return res
}

// isMap 检查变量是否字典.
func isMap(val interface{}) bool {
	return reflect.ValueOf(val).Kind() == reflect.Map
}

// isStruct 检查变量是否结构体.
func isStruct(val interface{}) bool {
	return reflect.ValueOf(val).Kind() == reflect.Struct
}

// isInt 变量是否整型数值.
func isInt(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.Atoi(str)
		return err == nil
	}

	return false
}

// isFloat 变量是否浮点数值.
func isFloat(val interface{}) bool {
	switch val.(type) {
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}

		if ok := RegFloat.MatchString(str); ok {
			return true
		}
	}

	return false
}

// isNumeric 变量是否数值(不包含复数和科学计数法).
func isNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		_, err := strconv.ParseFloat(str, 64)
		return err == nil
	}

	return false
}

// md5Byte 计算字节切片的 MD5 散列值.
func md5Byte(str []byte, length uint8) []byte {
	var res []byte
	h := md5.New()
	_, _ = h.Write(str)

	hBytes := h.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(dst, hBytes)
	if length > 0 && length < 32 {
		res = dst[:length]
	} else {
		res = dst
	}

	return res
}

// GetFieldValue 获取(字典/结构体的)字段值;fieldName为字段名,大小写敏感.
func GetFieldValue(arr interface{}, fieldName string) interface{} {
	var res interface{}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Struct:
		field := val.FieldByName(fieldName)
		if !field.IsValid() {
			break
		} else if !field.CanInterface() {
			break
		}
		res = field.Interface()
	case reflect.Map:
		for _, subKey := range val.MapKeys() {
			if fmt.Sprintf("%s", subKey) == fieldName {
				res = val.MapIndex(subKey).Interface()
				break
			}
		}
	default:
		panic("[GetFieldValue]`arr type must be struct or map; but : " + val.Kind().String())
	}

	return res
}

// GetVariateType 获取变量类型.
func GetVariateType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
