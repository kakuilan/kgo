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

// Str2Float32 将字符串转换为float32;其中"true", "TRUE", "True"为1.0 .
func (kc *LkkConvert) Str2Float32(val string) float32 {
	return str2Float32(val)
}

// Str2Float64 将字符串转换为float64;其中"true", "TRUE", "True"为1.0 .
func (kc *LkkConvert) Str2Float64(val string) float64 {
	return str2Float64(val)
}

// Str2Bool 将字符串转换为布尔值.
// 1, t, T, TRUE, true, True 等字符串为真;
// 0, f, F, FALSE, false, False 等字符串为假.
func (kc *LkkConvert) Str2Bool(val string) bool {
	return str2Bool(val)
}

// Str2Bytes 将字符串转换为字节切片.
func (kc *LkkConvert) Str2Bytes(val string) []byte {
	return str2Bytes(val)
}

// Bytes2Str 将字节切片转换为字符串.
func (kc *LkkConvert) Bytes2Str(val []byte) string {
	return bytes2Str(val)
}

// Str2BytesUnsafe (非安全的)将字符串转换为字节切片.
// 该方法零拷贝,但不安全.它直接转换底层指针,两者指向的相同的内存,改一个另外一个也会变.
// 仅当临时需将长字符串转换且不长时间保存时可以使用.
// 转换之后若没做其他操作直接改变里面的字符,则程序会崩溃.
// 如 b:=Str2BytesUnsafe("xxx"); b[1]='d'; 程序将panic.
func (kc *LkkConvert) Str2BytesUnsafe(val string) []byte {
	return str2BytesUnsafe(val)
}

// Bytes2StrUnsafe (非安全的)将字节切片转换为字符串.
// 零拷贝,不安全.效率是string([]byte{})的百倍以上,且转换量越大效率优势越明显.
func (kc *LkkConvert) Bytes2StrUnsafe(val []byte) string {
	return bytes2StrUnsafe(val)
}

// Dec2Bin 将十进制转换为二进制字符串.
func (kc *LkkConvert) Dec2Bin(num int64) string {
	return dec2Bin(num)
}

// Bin2Dec 将二进制字符串转换为十进制.
func (kc *LkkConvert) Bin2Dec(str string) (int64, error) {
	return bin2Dec(str)
}

// Hex2Bin 将十六进制字符串转换为二进制字符串.
func (kc *LkkConvert) Hex2Bin(str string) (string, error) {
	return hex2Bin(str)
}

// Bin2Hex 将二进制字符串转换为十六进制字符串.
func (kc *LkkConvert) Bin2Hex(str string) (string, error) {
	return bin2Hex(str)
}

// Dec2Hex 将十进制转换为十六进制.
func (kc *LkkConvert) Dec2Hex(num int64) string {
	return dec2Hex(num)
}

// Hex2Dec 将十六进制转换为十进制.
func (kc *LkkConvert) Hex2Dec(str string) (int64, error) {
	return hex2Dec(str)
}
