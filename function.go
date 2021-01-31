package kgo

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
)

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

// GetArrayFieldValue 获取数组(字典/结构体)的字段值,fieldName为字段名.
func GetArrayFieldValue(arr interface{}, fieldName string) interface{} {
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
		panic("[GetArrayFieldValue]`arr must be struct or map; but : " + val.Kind().String())
	}

	return res
}
