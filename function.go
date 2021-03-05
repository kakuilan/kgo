package kgo

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// dumpPrint 打印调试变量,变量可多个.
func dumpPrint(vs ...interface{}) {
	for _, v := range vs {
		fmt.Printf("%+v\n", v)
		//fmt.Printf("%#v\n", v)
	}
}

// lenArrayOrSlice 获取数组/切片的长度.
// chkType为检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示变量不是数组或切片,>=0表示合法长度.
func lenArrayOrSlice(val interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		chkType = 3
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
			if !filterZero || (filterZero && !fieldVal.IsZero()) {
				res = append(res, fieldVal.Interface())
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			fieldVal = val.MapIndex(k)
			if !filterZero || (filterZero && !fieldVal.IsZero()) {
				res = append(res, fieldVal.Interface())
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			fieldVal = val.Field(i)
			if fieldVal.CanInterface() {
				if !filterZero || (filterZero && !fieldVal.IsZero()) {
					res = append(res, fieldVal.Interface())
				}
			}
		}
	default:
		panic("[arrayValues]`arr type must be array|slice|map|struct; but : " + val.Kind().String())
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

// reflect2Itf 将反射值转为接口(原值)
func reflect2Itf(r reflect.Value) (res interface{}) {
	switch r.Kind() {
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int, reflect.Int64:
		res = r.Int()
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint, reflect.Uint64:
		res = r.Uint()
	case reflect.Float32, reflect.Float64:
		res = r.Float()
	case reflect.String:
		res = r.String()
	case reflect.Bool:
		res = r.Bool()
	default:
		if r.CanInterface() {
			res = r.Interface()
		} else {
			res = r
		}
	}

	return
}

// structVal 获取结构体的反射值
func structVal(obj interface{}) (reflect.Value, error) {
	v := reflect.ValueOf(obj)

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return v, errors.New("[structVal]`obj type must be struct; but : " + v.Kind().String())
	}

	return v, nil
}

// structFields 获取结构体的字段;all是否包含所有字段(包括未导出的).
func structFields(obj interface{}, all bool) ([]reflect.StructField, error) {
	v, e := structVal(obj)
	if e != nil {
		return nil, e
	}

	var fs []reflect.StructField
	var t = v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		// 不能访问未导出的字段
		if !all && field.PkgPath != "" {
			continue
		}

		fs = append(fs, field)
	}

	return fs, nil
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
func compareConditionMap(condition map[string]interface{}, arr interface{}) (res interface{}) {
	val := reflect.ValueOf(arr)
	conNum := len(condition)
	if conNum > 0 {
		chkNum := 0

		switch val.Kind() {
		case reflect.Map:
			if conNum > 0 {
				for _, k := range val.MapKeys() {
					if condVal, ok := condition[k.String()]; ok && reflect.DeepEqual(val.MapIndex(k).Interface(), condVal) {
						chkNum++
					}
				}
			}
		case reflect.Struct:
			var field reflect.Value
			for k, v := range condition {
				field = val.FieldByName(k)

				if field.IsValid() && field.CanInterface() && reflect.DeepEqual(field.Interface(), v) {
					chkNum++
				}
			}
		default:
			panic("[compareConditionMap]`arr type must be map|struct; but : " + val.Kind().String())
		}

		if chkNum == conNum {
			res = arr
		}
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

// methodExists 检查val结构体中是否存在methodName方法.
func methodExists(val interface{}, methodName string) (bool, error) {
	valRef := reflect.ValueOf(val)

	if valRef.Type().Kind() != reflect.Ptr {
		valRef = reflect.New(reflect.TypeOf(val))
	}

	method := valRef.MethodByName(methodName)
	if !method.IsValid() {
		return false, fmt.Errorf("[methodExists] Method `%s` not exists in interface `%s`", methodName, valRef.Type())
	}

	return true, nil
}

// getMethod 获取val结构体的methodName方法.
func getMethod(val interface{}, methodName string) reflect.Value {
	m, b := reflect.TypeOf(val).MethodByName(methodName)
	if !b {
		return reflect.ValueOf(nil)
	}
	return m.Func
}

// camelCaseToLowerCase 驼峰转为小写.
func camelCaseToLowerCase(str string, connector rune) string {
	if len(str) == 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	var prev, r0, r1 rune
	var size int

	r0 = connector

	for len(str) > 0 {
		prev = r0
		r0, size = utf8.DecodeRuneInString(str)
		str = str[size:]

		switch {
		case r0 == utf8.RuneError:
			continue

		case unicode.IsUpper(r0):
			if prev != connector && !unicode.IsNumber(prev) {
				buf.WriteRune(connector)
			}

			buf.WriteRune(unicode.ToLower(r0))

			if len(str) == 0 {
				break
			}

			r0, size = utf8.DecodeRuneInString(str)
			str = str[size:]

			if !unicode.IsUpper(r0) {
				buf.WriteRune(r0)
				break
			}

			// find next non-upper-case character and insert connector properly.
			// it's designed to convert `HTTPServer` to `http_server`.
			// if there are more than 2 adjacent upper case characters in a word,
			// treat them as an abbreviation plus a normal word.
			for len(str) > 0 {
				r1 = r0
				r0, size = utf8.DecodeRuneInString(str)
				str = str[size:]

				if r0 == utf8.RuneError {
					buf.WriteRune(unicode.ToLower(r1))
					break
				}

				if !unicode.IsUpper(r0) {
					if isCaseConnector(r0) {
						r0 = connector

						buf.WriteRune(unicode.ToLower(r1))
					} else if unicode.IsNumber(r0) {
						// treat a number as an upper case rune
						// so that both `http2xx` and `HTTP2XX` can be converted to `http_2xx`.
						buf.WriteRune(unicode.ToLower(r1))
						buf.WriteRune(connector)
						buf.WriteRune(r0)
					} else {
						buf.WriteRune(connector)
						buf.WriteRune(unicode.ToLower(r1))
						buf.WriteRune(r0)
					}

					break
				}

				buf.WriteRune(unicode.ToLower(r1))
			}

			if len(str) == 0 || r0 == connector {
				buf.WriteRune(unicode.ToLower(r0))
			}

		case unicode.IsNumber(r0):
			if prev != connector && !unicode.IsNumber(prev) {
				buf.WriteRune(connector)
			}

			buf.WriteRune(r0)

		default:
			if isCaseConnector(r0) {
				r0 = connector
			}

			buf.WriteRune(r0)
		}
	}

	return buf.String()
}

// isCaseConnector 是否字符转换连接符.
func isCaseConnector(r rune) bool {
	return r == '-' || r == '_' || unicode.IsSpace(r)
}

// pkcs7Padding PKCS7填充.
// cipherText为密文;blockSize为分组长度;isZero是否零填充.
func pkcs7Padding(cipherText []byte, blockSize int, isZero bool) []byte {
	clen := len(cipherText)
	if cipherText == nil || clen == 0 || blockSize <= 0 {
		return nil
	}

	var padtext []byte
	padding := blockSize - clen%blockSize
	if isZero {
		padtext = bytes.Repeat([]byte{0}, padding)
	} else {
		padtext = bytes.Repeat([]byte{byte(padding)}, padding)
	}

	return append(cipherText, padtext...)
}

// pkcs7UnPadding PKCS7拆解.
// origData为源数据;blockSize为分组长度.
func pkcs7UnPadding(origData []byte, blockSize int) []byte {
	olen := len(origData)
	if origData == nil || olen == 0 || blockSize <= 0 || olen%blockSize != 0 {
		return nil
	}

	unpadding := int(origData[olen-1])
	if unpadding == 0 || unpadding > olen {
		return nil
	}

	return origData[:(olen - unpadding)]
}

// zeroPadding PKCS7使用0填充.
func zeroPadding(cipherText []byte, blockSize int) []byte {
	return pkcs7Padding(cipherText, blockSize, true)
}

// zeroUnPadding PKCS7-0拆解.
func zeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

// GetFieldValue 获取(字典/结构体的)字段值;fieldName为字段名,大小写敏感.
func GetFieldValue(arr interface{}, fieldName string) (res interface{}, err error) {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Map:
		for _, subKey := range val.MapKeys() {
			if fmt.Sprintf("%s", subKey) == fieldName {
				res = val.MapIndex(subKey).Interface()
				break
			}
		}
	case reflect.Struct:
		field := val.FieldByName(fieldName)
		if !field.IsValid() || !field.CanInterface() {
			break
		}
		res = field.Interface()
	default:
		err = errors.New("[GetFieldValue]`arr type must be map|struct; but : " + val.Kind().String())
	}

	return
}

// GetVariateType 获取变量类型.
func GetVariateType(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// ValidFunc 检查是否函数,并且参数个数、类型是否正确.
// 返回有效的函数、有效的参数.
func ValidFunc(f interface{}, args ...interface{}) (vf reflect.Value, vargs []reflect.Value, err error) {
	vf = reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[ValidFunc] %v is not the function", f)
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[ValidFunc] %d number of the argument is incorrect", _len)
	}

	vargs = make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, fmt.Errorf("[ValidFunc] %d-td argument`s type is incorrect", i+1)
		}
		vargs[i] = reflect.ValueOf(args[i])
	}
	return vf, vargs, nil
}

// CallFunc 动态调用函数.
func CallFunc(f interface{}, args ...interface{}) (results []interface{}, err error) {
	vf, vargs, _err := ValidFunc(f, args...)
	if _err != nil {
		return nil, _err
	}
	ret := vf.Call(vargs)
	_len := len(ret)
	results = make([]interface{}, _len)
	for i := 0; i < _len; i++ {
		results[i] = ret[i].Interface()
	}
	return
}

// str2Int 将字符串转换为int.其中"true", "TRUE", "True"为1;若为浮点字符串,则取整数部分.
func str2Int(val string) (res int) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1
		return
	} else if ok := RegFloat.MatchString(val); ok {
		fl, _ := strconv.ParseFloat(val, 1)
		res = int(fl)
		return
	}

	res, _ = strconv.Atoi(val)
	return
}

// str2Int 将字符串转换为uint.其中"true", "TRUE", "True"为1;若为浮点字符串,则取整数部分;若为负值则为0.
func str2Uint(val string) (res uint) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1
		return
	} else if ok := RegFloat.MatchString(val); ok {
		fl, _ := strconv.ParseFloat(val, 1)
		if fl > 0 {
			res = uint(fl)
		}

		return
	}

	n, e := strconv.Atoi(val)
	if e == nil && n > 0 {
		res = uint(n)
	}

	return
}

// str2Float32 将字符串转换为float32;其中"true", "TRUE", "True"为1.0 .
func str2Float32(val string) (res float32) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1.0
	} else {
		r, _ := strconv.ParseFloat(val, 32)
		res = float32(r)
	}

	return
}

// str2Float64 将字符串转换为float64;其中"true", "TRUE", "True"为1.0 .
func str2Float64(val string) (res float64) {
	if val == "true" || val == "TRUE" || val == "True" {
		res = 1.0
	} else {
		res, _ = strconv.ParseFloat(val, 64)
	}

	return
}

// str2Bool 将字符串转换为布尔值.
// 1, t, T, TRUE, true, True 等字符串为真;
// 0, f, F, FALSE, false, False 等字符串为假.
func str2Bool(val string) (res bool) {
	if val != "" {
		res, _ = strconv.ParseBool(val)
	}

	return
}

// bool2Int 将布尔值转换为整型.
func bool2Int(val bool) int {
	if val {
		return 1
	}
	return 0
}

// str2Bytes 将字符串转换为字节切片.
func str2Bytes(val string) []byte {
	return []byte(val)
}

// bytes2Str 将字节切片转换为字符串.
func bytes2Str(val []byte) string {
	return string(val)
}

// str2BytesUnsafe (非安全的)将字符串转换为字节切片.
// 该方法零拷贝,但不安全.它直接转换底层指针,两者指向的相同的内存,改一个另外一个也会变.
// 仅当临时需将长字符串转换且不长时间保存时可以使用.
// 转换之后若没做其他操作直接改变里面的字符,则程序会崩溃.
// 如 b:=str2BytesUnsafe("xxx"); b[1]='d'; 程序将panic.
func str2BytesUnsafe(val string) []byte {
	psHeader := &reflect.SliceHeader{}
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&val))
	psHeader.Data = strHeader.Data
	psHeader.Len = strHeader.Len
	psHeader.Cap = strHeader.Len
	return *(*[]byte)(unsafe.Pointer(psHeader))
}

// bytes2StrUnsafe (非安全的)将字节切片转换为字符串.
// 零拷贝,不安全.效率是string([]byte{})的百倍以上,且转换量越大效率优势越明显.
func bytes2StrUnsafe(val []byte) string {
	return *(*string)(unsafe.Pointer(&val))
}

// toStr 强制将变量转换为字符串.
func toStr(val interface{}) string {
	//先处理其他类型
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Invalid:
		return ""
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return v.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Float32:
		return strconv.FormatFloat(v.Float(), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.Float(), 'f', -1, 64)
	case reflect.Ptr, reflect.Struct, reflect.Map: //指针、结构体和字典
		b, err := json.Marshal(v.Interface())
		if err != nil {
			return ""
		}
		return string(b)
	}

	//再处理字节切片
	switch val.(type) {
	case []uint8:
		return string(val.([]uint8))
	}

	return fmt.Sprintf("%v", val)
}

// toInt 强制将变量转换为整型;其中true或"true"为1.
func toInt(val interface{}) int {
	switch val.(type) {
	case int:
		return val.(int)
	case int8:
		return int(val.(int8))
	case int16:
		return int(val.(int16))
	case int32:
		return int(val.(int32))
	case int64:
		return int(val.(int64))
	case uint:
		return int(val.(uint))
	case uint8:
		return int(val.(uint8))
	case uint16:
		return int(val.(uint16))
	case uint32:
		return int(val.(uint32))
	case uint64:
		return int(val.(uint64))
	case float32:
		return int(val.(float32))
	case float64:
		return int(val.(float64))
	case []uint8:
		return str2Int(string(val.([]uint8)))
	case string:
		return str2Int(val.(string))
	case bool:
		return bool2Int(val.(bool))
	default:
		return 0
	}
}

// toFloat 强制将变量转换为浮点型;其中true或"true"为1.0 .
func toFloat(val interface{}) (res float64) {
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
	case []uint8:
		res = str2Float64(string(val.([]uint8)))
	case string:
		res = str2Float64(val.(string))
	case bool:
		if val.(bool) {
			res = 1.0
		}
	}

	return
}

// struct2Map 结构体转为字典;tagName为要导出的标签名,可以为空,为空时将导出所有字段.
func struct2Map(obj interface{}, tagName string) (map[string]interface{}, error) {
	v, e := structVal(obj)
	if e != nil {
		return nil, e
	}

	t := v.Type()
	var res = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tagName != "" {
			if tagValue := field.Tag.Get(tagName); tagValue != "" {
				res[tagValue] = reflect2Itf(v.Field(i))
			}
		} else {
			res[field.Name] = reflect2Itf(v.Field(i))
		}
	}

	return res, nil
}

// dec2Bin 将十进制转换为二进制字符串.
func dec2Bin(num int64) string {
	return strconv.FormatInt(num, 2)
}

// bin2Dec 将二进制字符串转换为十进制.
func bin2Dec(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return 0, err
	}
	return i, nil
}

// hex2Bin 将十六进制字符串转换为二进制字符串.
func hex2Bin(str string) (string, error) {
	i, err := strconv.ParseInt(str, 16, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 2), nil
}

// bin2Hex 将二进制字符串转换为十六进制字符串.
func bin2Hex(str string) (string, error) {
	i, err := strconv.ParseInt(str, 2, 0)
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, 16), nil
}

// dec2Hex 将十进制转换为十六进制.
func dec2Hex(num int64) string {
	return strconv.FormatInt(num, 16)
}

// hex2Dec 将十六进制转换为十进制.
func hex2Dec(str string) (int64, error) {
	start := 0
	if len(str) > 2 && str[0:2] == "0x" {
		start = 2
	}

	// bitSize 表示结果的位宽（包括符号位），0 表示最大位宽
	return strconv.ParseInt(str[start:], 16, 0)
}

// dec2Oct 将十进制转换为八进制.
func dec2Oct(num int64) string {
	return strconv.FormatInt(num, 8)
}
