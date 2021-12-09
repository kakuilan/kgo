package kgo

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"net"
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

// Dec2Oct 将十进制转换为八进制.
func (kc *LkkConvert) Dec2Oct(num int64) string {
	return dec2Oct(num)
}

// Oct2Dec 将八进制转换为十进制.
func (kc *LkkConvert) Oct2Dec(str string) (int64, error) {
	return oct2Dec(str)
}

// Runes2Bytes 将[]rune转为[]byte.
func (kc *LkkConvert) Runes2Bytes(rs []rune) []byte {
	return runes2Bytes(rs)
}

// BaseConvert 进制转换,在任意进制之间转换数字.
// num为输入数值,frombase为原进制,tobase为结果进制.
func (kc *LkkConvert) BaseConvert(num string, frombase, tobase int) (string, error) {
	i, err := strconv.ParseInt(num, frombase, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, tobase), nil
}

// Ip2Long 将 IPV4 的字符串互联网协议转换成长整型数字.
func (kc *LkkConvert) Ip2Long(ipAddress string) uint32 {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip.To4())
}

// Long2Ip 将长整型转化为字符串形式带点的互联网标准格式地址(IPV4).
func (kc *LkkConvert) Long2Ip(properAddress uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, properAddress)
	ip := net.IP(ipByte)
	return ip.String()
}

// ToStr 强制将变量转换为字符串.
func (kc *LkkConvert) ToStr(val interface{}) string {
	return toStr(val)
}

// ToBool 强制将变量转换为布尔值.
// 数值类型将检查值是否>0;
// 字符串将使用Str2Bool;
// 数组、切片、字典、通道类型将检查它们的长度是否>0;
// 指针、结构体类型为true,其他为false.
func (kc *LkkConvert) ToBool(val interface{}) bool {
	return toBool(val)
}

// ToInt 强制将变量转换为整型.
// 数值类型将转为整型;
// 字符串类型将使用Str2Int;
// 布尔型的true为1,false为0;
// 数组、切片、字典、通道类型将取它们的长度;
// 指针、结构体类型为1,其他为0.
func (kc *LkkConvert) ToInt(val interface{}) int {
	return toInt(val)
}

// ToFloat 强制将变量转换为浮点型.
// 数值类型将转为浮点型;
// 字符串将使用Str2Float64;
// 布尔型的true为1.0,false为0;
// 数组、切片、字典、通道类型将取它们的长度;
// 指针、结构体类型为1.0,其他为0.
func (kc *LkkConvert) ToFloat(val interface{}) (res float64) {
	return toFloat(val)
}

// Float64ToByte 64位浮点数转字节切片.
func (kc *LkkConvert) Float64ToByte(val float64) []byte {
	bits := math.Float64bits(val)
	res := make([]byte, 8)
	binary.LittleEndian.PutUint64(res, bits)

	return res
}

// Byte2Float64 字节切片转64位浮点数.
func (kc *LkkConvert) Byte2Float64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)

	return math.Float64frombits(bits)
}

// Int64ToByte 64位整型转字节切片.
func (kc *LkkConvert) Int64ToByte(val int64) []byte {
	res := make([]byte, 8)
	binary.BigEndian.PutUint64(res, uint64(val))

	return res
}

// Byte2Int64 字节切片转64位整型.
func (kc *LkkConvert) Byte2Int64(val []byte) int64 {
	return int64(binary.BigEndian.Uint64(val))
}

// Byte2Hex 字节切片转16进制字符串.
func (kc *LkkConvert) Byte2Hex(val []byte) string {
	return hex.EncodeToString(val)
}

// Byte2Hexs 字节切片转16进制切片.
func (kc *LkkConvert) Byte2Hexs(val []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(val)))
	hex.Encode(dst, val)
	return dst
}

// Hex2Byte 16进制字符串转字节切片.
func (kc *LkkConvert) Hex2Byte(str string) []byte {
	h, _ := hex2Byte(str)
	return h
}

// Hexs2Byte 16进制切片转byte切片.
func (kc *LkkConvert) Hexs2Byte(val []byte) []byte {
	dst := make([]byte, hex.DecodedLen(len(val)))
	_, err := hex.Decode(dst, val)

	if err != nil {
		return nil
	}
	return dst
}

// IsString 变量是否字符串.
func (kc *LkkConvert) IsString(val interface{}) bool {
	return isString(val)
}

// IsBinary 字符串是否二进制.
func (kc *LkkConvert) IsBinary(s string) bool {
	return isBinary(s)
}

// IsNumeric 变量是否数值(不包含复数和科学计数法).
func (kc *LkkConvert) IsNumeric(val interface{}) bool {
	return isNumeric(val)
}

// IsInt 变量是否整型数值.
func (kc *LkkConvert) IsInt(val interface{}) bool {
	return isInt(val)
}

// IsFloat 变量是否浮点数值.
func (kc *LkkConvert) IsFloat(val interface{}) bool {
	return isFloat(val)
}

// IsEmpty 变量是否为空.
func (kc *LkkConvert) IsEmpty(val interface{}) bool {
	return isEmpty(val)
}

// IsNil 变量是否nil.
func (kc *LkkConvert) IsNil(val interface{}) bool {
	return isNil(val)
}

// IsBool 是否布尔值.
func (kc *LkkConvert) IsBool(val interface{}) bool {
	return isBool(val)
}

// IsHex 是否十六进制字符串.
func (kc *LkkConvert) IsHex(str string) bool {
	return isHex(str)
}

// IsByte 变量是否字节切片.
func (kc *LkkConvert) IsByte(val interface{}) bool {
	return isByte(val)
}

// IsStruct 变量是否结构体.
func (kc *LkkConvert) IsStruct(val interface{}) bool {
	return isStruct(val)
}

// IsInterface 变量是否接口.
func (kc *LkkConvert) IsInterface(val interface{}) bool {
	return isInterface(val)
}

// IsPort 变量值是否端口号(1~65535).
func (kc *LkkConvert) IsPort(val interface{}) bool {
	return isPort(val)
}

// ToInterfaces 强制将变量转为接口切片;和KArr.ArrayValues相同.
// 其中val类型必须是数组/切片/字典/结构体.
func (kc *LkkConvert) ToInterfaces(val interface{}) []interface{} {
	return arrayValues(val, false)
}
