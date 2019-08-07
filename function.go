package kgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"reflect"
	"unsafe"
)

// md5Str 计算字符串的 MD5 散列值
func md5Str(str []byte, length uint8) []byte {
	var res []byte
	hash := md5.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(dst, hBytes)
	if length > 0 && length < 32 {
		res = dst[:length]
	} else {
		res = dst
	}

	return res
}

// sha1Str 计算字符串的 sha1 散列值
func sha1Str(str []byte) []byte {
	hash := sha1.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}

// sha256Str 计算字符串的 sha256 散列值
func sha256Str(str []byte) []byte {
	hash := sha256.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}

// isArrayOrSlice 检查变量是否数组或切片;chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func isArrayOrSlice(data interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		msg := fmt.Sprintf("[isArrayOrSlice]chkType value muset in (1, 2, 3), but it`s %d", chkType)
		panic(msg)
	}

	var res int = -1
	val := reflect.ValueOf(data)
	switch val.Kind() {
	case reflect.Array:
		if chkType == 1 || chkType == 3 {
			res = val.Len()
		}
	case reflect.Slice:
		if chkType == 2 || chkType == 3 {
			res = val.Len()
		}
	}

	return res
}

// isMap 检查变量是否字典
func isMap(data interface{}) bool {
	return reflect.ValueOf(data).Kind() == reflect.Map
}

// getEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian
func getEndian() (nativeEndian binary.ByteOrder) {
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		nativeEndian = binary.BigEndian
	default:
		panic("[getEndian] could not determine native endianness")
	}

	return
}

// isLittleEndian 系统字节序类型是否小端存储
func isLittleEndian() bool {
	var i int32 = 0x01020304

	// 将int32类型的指针转换为byte类型的指针
	u := unsafe.Pointer(&i)
	pb := (*byte)(u)

	// 取得pb位置对应的值
	b := *pb

	// 由于b是byte类型的,最多保存8位,那么只能取得开始的8位
	// 小端: 04 (03 02 01)
	// 大端: 01 (02 03 04)
	return (b == 0x04)
}
