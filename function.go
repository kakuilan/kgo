package kgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

// isArrayOrSlice 检查变量是否数组或切片.
// chkType为检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是(数组长度).
func isArrayOrSlice(val interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		panic(fmt.Sprintf("[isArrayOrSlice]`chkType must in [1, 2, 3]; but: %d", chkType))
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

// isLittleEndian 系统字节序类型是否小端存储.
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

// getEndian 获取系统字节序类型,小端返回binary.LittleEndian,大端返回binary.BigEndian .
func getEndian() binary.ByteOrder {
	var nativeEndian binary.ByteOrder = binary.BigEndian
	buf := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buf[0])) = uint16(0xABCD)

	switch buf {
	case [2]byte{0xCD, 0xAB}:
		nativeEndian = binary.LittleEndian
		//case [2]byte{0xAB, 0xCD}:
		//	nativeEndian = binary.BigEndian
	}

	return nativeEndian
}

// numeric2Float 将数值转换为float64.
func numeric2Float(val interface{}) (res float64, err error) {
	switch val.(type) {
	case int:
		res = float64(val.(int))
	case int8:
		res = float64(val.(int8))
	case int16:
		res = float64(val.(int16))
	case int32:
		res = float64(val.(int32))
	case int64:
		res = float64(val.(int64))
	case uint:
		res = float64(val.(uint))
	case uint8:
		res = float64(val.(uint8))
	case uint16:
		res = float64(val.(uint16))
	case uint32:
		res = float64(val.(uint32))
	case uint64:
		res = float64(val.(uint64))
	case float32:
		res = float64(val.(float32))
	case float64:
		res = val.(float64)
	case string:
		str := val.(string)
		res, err = strconv.ParseFloat(str, 64)
	}
	return
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

// shaXByte 计算字节切片的 shaX 散列值,x为1/256/512.
func shaXByte(str []byte, x uint16) []byte {
	var h hash.Hash
	switch x {
	case 1:
		h = sha1.New()
		break
	case 256:
		h = sha256.New()
		break
	case 512:
		h = sha512.New()
		break
	default:
		panic(fmt.Sprintf("[shaXByte]`x must be in [1, 256, 512]; but: %d", x))
	}

	h.Write(str)

	hBytes := h.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}

// arrayValues 返回arr(数组/切片/字典/结构体)中所有的值.
// filterZero 是否过滤零值元素(nil,false,0,'',[]),true时排除零值元素,false时保留零值元素.
func arrayValues(arr interface{}, filterZero bool) []interface{} {
	var res []interface{}
	var fieldVal reflect.Value
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			fieldVal = val.Index(i)
			if !filterZero || (filterZero && !fieldVal.IsNil() && !fieldVal.IsZero()) {
				res = append(res, fieldVal.Interface())
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			fieldVal = val.Field(i)
			if fieldVal.IsValid() && fieldVal.CanInterface() {
				if !filterZero || (filterZero && !fieldVal.IsNil() && !fieldVal.IsZero()) {
					res = append(res, fieldVal.Interface())
				}
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			fieldVal = val.MapIndex(k)
			if !filterZero || (filterZero && !fieldVal.IsNil() && !fieldVal.IsZero()) {
				res = append(res, fieldVal.Interface())
			}
		}
	default:
		panic("[arrayValues]`arr type must be array|slice|struct|map; but : " + val.Kind().String())
	}

	return res
}

// reflectPtr 获取反射的指向.
func reflectPtr(r reflect.Value) reflect.Value {
	// 如果是指针,则获取其所指向的元素
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return r
}

// creditChecksum 计算身份证校验码,其中id为身份证号码.
func creditChecksum(id string) byte {
	//∑(ai×Wi)(mod 11)
	// 加权因子
	factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	// 校验位对应值
	code := []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}

	leng := len(id)
	sum := 0
	for i, char := range id[:leng-1] {
		num, _ := strconv.Atoi(string(char))
		sum += num * factor[i]
	}

	return code[sum%11]
}

// compareConditionMap 比对数组是否匹配条件.condition为条件字典,arr为要比对的数据(字典/结构体).
func compareConditionMap(condition map[interface{}]interface{}, arr interface{}) (res interface{}) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Map:
		condLen := len(condition)
		chkNum := 0
		if condLen > 0 {
			for _, k := range val.MapKeys() {
				if condVal, ok := condition[k.String()]; ok && reflect.DeepEqual(val.MapIndex(k).Interface(), condVal) {
					chkNum++
				}
			}
		}

		if chkNum == condLen {
			res = arr
		}
	default:
		return
	}

	return
}

// getTrimMask 获取要修剪的字符串集合,masks为要屏蔽的字符切片.
func getTrimMask(masks []string) string {
	var str string
	if len(masks) == 0 {
		str = " \t\n\r\v\f\x00　"
	} else {
		str = strings.Join(masks, "")
	}
	return str
}

func methodExists() {
	//TODO 方法是否存在
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
		panic("[GetFieldValue]`arr type must be struct|map; but : " + val.Kind().String())
	}

	return res
}

// GetVariateType 获取变量类型.
func GetVariateType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}
