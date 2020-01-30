package kgo

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"net"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
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

// Bool2Int 将布尔值转换为整型.
func (kc *LkkConvert) Bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

// Str2IntStrict 严格将字符串转换为有符号整型.
// bitSize为类型位数,strict为是否严格检查.
func (kc *LkkConvert) Str2IntStrict(val string, bitSize int, strict bool) int64 {
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
	return int8(kc.Str2IntStrict(val, 8, false))
}

// Str2Int16 将字符串转换为int16
func (kc *LkkConvert) Str2Int16(val string) int16 {
	return int16(kc.Str2IntStrict(val, 16, false))
}

// Str2Int32 将字符串转换为int32
func (kc *LkkConvert) Str2Int32(val string) int32 {
	return int32(kc.Str2IntStrict(val, 32, false))
}

// Str2Int64 将字符串转换为int64
func (kc *LkkConvert) Str2Int64(val string) int64 {
	return kc.Str2IntStrict(val, 64, false)
}

// StrictStr2Uint 严格将字符串转换为无符号整型,bitSize为类型位数,strict为是否严格检查
func (kc *LkkConvert) StrictStr2Uint(val string, bitSize int, strict bool) uint64 {
	res, err := strconv.ParseUint(val, 0, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// Str2Uint 将字符串转换为uint
func (kc *LkkConvert) Str2Uint(val string) uint {
	return uint(kc.StrictStr2Uint(val, 0, false))
}

// Str2Uint8 将字符串转换为uint8
func (kc *LkkConvert) Str2Uint8(val string) uint8 {
	return uint8(kc.StrictStr2Uint(val, 8, false))
}

// Str2Uint16 将字符串转换为uint16
func (kc *LkkConvert) Str2Uint16(val string) uint16 {
	return uint16(kc.StrictStr2Uint(val, 16, false))
}

// Str2Uint32 将字符串转换为uint32
func (kc *LkkConvert) Str2Uint32(val string) uint32 {
	return uint32(kc.StrictStr2Uint(val, 32, false))
}

// Str2Uint64 将字符串转换为uint64
func (kc *LkkConvert) Str2Uint64(val string) uint64 {
	return uint64(kc.StrictStr2Uint(val, 64, false))
}

// Str2FloatStrict 严格将字符串转换为浮点型.
// bitSize为类型位数,strict为是否严格检查
func (kc *LkkConvert) Str2FloatStrict(val string, bitSize int, strict bool) float64 {
	res, err := strconv.ParseFloat(val, bitSize)
	if err != nil {
		if strict {
			panic(err)
		}
	}
	return res
}

// Str2Float32 将字符串转换为float32
func (kc *LkkConvert) Str2Float32(val string) float32 {
	return float32(kc.Str2FloatStrict(val, 32, false))
}

// Str2Float64 将字符串转换为float64
func (kc *LkkConvert) Str2Float64(val string) float64 {
	return float64(kc.Str2FloatStrict(val, 64, false))
}

// Str2Bool 将字符串转换为布尔值
func (kc *LkkConvert) Str2Bool(val string) bool {
	if val == "true" || val == "True" || val == "TRUE" {
		return true
	}
	return false
}

// Int2Bool 将整数转换为布尔值
func (kc *LkkConvert) Int2Bool(val interface{}) bool {
	switch val.(type) {
	case int:
		return (val.(int) > 0)
	case int8:
		return (val.(int8) > 0)
	case int16:
		return (val.(int16) > 0)
	case int32:
		return (val.(int32) > 0)
	case int64:
		return (val.(int64) > 0)
	case uint:
		return (val.(uint) > 0)
	case uint8:
		return (val.(uint8) > 0)
	case uint16:
		return (val.(uint16) > 0)
	case uint32:
		return (val.(uint32) > 0)
	case uint64:
		return (val.(uint64) > 0)
	default:
		return false
	}
}

// Str2ByteSlice 将字符串转换为字节切片;该方法零拷贝,但不安全,仅当临时需将长字符串转换且不长时间保存时可以使用.
func (kc *LkkConvert) Str2ByteSlice(val string) []byte {
	pSliceHeader := &reflect.SliceHeader{}
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
	pSliceHeader.Data = strHeader.Data
	pSliceHeader.Len = strHeader.Len
	pSliceHeader.Cap = strHeader.Len
	return *(*[]byte)(unsafe.Pointer(pSliceHeader))
}

// ByteSlice2Str 将字节切片转换为字符串,零拷贝
func (kc *LkkConvert) ByteSlice2Str(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

// Dec2bin 将十进制转换为二进制
func (kc *LkkConvert) Dec2bin(number int64) string {
	return strconv.FormatInt(number, 2)
}

// Bin2dec 将二进制转换为十进制
func (kc *LkkConvert) Bin2dec(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// Hex2bin 将十六进制字符串转换为二进制字符串
func (kc *LkkConvert) Hex2bin(data string) (string, error) {
	i, err := strconv.ParseInt(data, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// Bin2hex 将二进制字符串转换为十六进制字符串
func (kc *LkkConvert) Bin2hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// Dec2hex 将十进制转换为十六进制
func (kc *LkkConvert) Dec2hex(number int64) string {
	return strconv.FormatInt(number, 16)
}

// Hex2dec 将十六进制转换为十进制
func (kc *LkkConvert) Hex2dec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽
	return strconv.ParseInt(str[start:], 16, 0)
}

// Dec2oct 将十进制转换为八进制
func (kc *LkkConvert) Dec2oct(number int64) string {
	return strconv.FormatInt(number, 8)
}

// Oct2dec 将八进制转换为十进制
func (kc *LkkConvert) Oct2dec(str string) (int64, error) {
	start := 0
	if len(str) > 1 && str[0:1] == "0" {
		start = 1
	}

	return strconv.ParseInt(str[start:], 8, 0)
}

// BaseConvert 进制转换,在任意进制之间转换数字
func (kc *LkkConvert) BaseConvert(number string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(number, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// Ip2long 将 IPV4 的字符串互联网协议转换成长整型数字
func (kc *LkkConvert) Ip2long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// Long2ip 将长整型转化为字符串形式带点的互联网标准格式地址（IPV4）
func (kc *LkkConvert) Long2ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

// Gettype 获取变量类型
func (kc *LkkConvert) Gettype(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// ToStr 强制将变量转换为字符串
func (kc *LkkConvert) ToStr(val interface{}) (res string) {
	switch val.(type) {
	case []byte:
		res = fmt.Sprintf("%s", (string(val.([]byte))))
	default:
		res = fmt.Sprintf("%v", val)
	}
	return
}

// ToInt 强制将变量转换为整型;其中true或"true"为1
func (kc *LkkConvert) ToInt(val interface{}) int {
	str := strings.ToLower(fmt.Sprintf("%v", val))
	if str == "true" {
		return 1
	} else {
		return kc.Str2Int(str)
	}
}

// ToFloat 强制将变量转换为浮点型;其中true或"true"为1.0
func (kc *LkkConvert) ToFloat(val interface{}) float64 {
	str := strings.ToLower(fmt.Sprintf("%v", val))
	if str == "true" {
		return 1.0
	} else {
		return kc.Str2Float64(str)
	}
}

// Float64ToByte 64位浮点数转字节切片
func (kc *LkkConvert) Float64ToByte(val float64) []byte {
	bits := math.Float64bits(val)
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, bits)

	return res
}

// ByteToFloat64 字节切片转64位浮点数
func (kc *LkkConvert) ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

// Int64ToByte 64位整型转字节切片
func (kc *LkkConvert) Int64ToByte(val int64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, uint64(val))

	return res
}

// ByteToInt64 字节切片转64位整型
func (kc *LkkConvert) ByteToInt64(val []byte) int64 {
	return int64(binary.BigEndian.Uint64(val))
}

// Byte2Hex 字节切片转16进制字符串
func (kc *LkkConvert) Byte2Hex(val []byte) string {
	return hex.EncodeToString(val)
}

// Hex2Byte 16进制字符串转字节切片
func (kc *LkkConvert) Hex2Byte(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}

// GetPointerAddrInt 获取变量指针地址整型值.variable为变量.
func (kc *LkkConvert) GetPointerAddrInt(variable interface{}) int64 {
	res, _ := kc.Hex2dec(fmt.Sprintf("%p", &variable))
	return res
}
