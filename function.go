package kgo

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"hash"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

// md5Str 计算字符串的 MD5 散列值.
func md5Str(str []byte, length uint8) []byte {
	var res []byte
	h := md5.New()
	h.Write(str)

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

// shaXStr 计算字符串的 shaX 散列值,x为1/256/512.
func shaXStr(str []byte, x uint16) []byte {
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
		panic("[shaXStr] x must be in [1, 256, 512]")
	}

	h.Write(str)

	hBytes := h.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}

// isArrayOrSlice 检查变量是否数组或切片.
// chkType为检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是.
func isArrayOrSlice(data interface{}, chkType uint8) int {
	if chkType != 1 && chkType != 2 && chkType != 3 {
		panic(fmt.Sprintf("[isArrayOrSlice] chkType value muset in (1, 2, 3), but it`s %d", chkType))
	}

	var res = -1
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

// isMap 检查变量是否字典.
func isMap(data interface{}) bool {
	return reflect.ValueOf(data).Kind() == reflect.Map
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

// arrayValues 返回数组/切片/字典中所有的值.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素.
func arrayValues(arr interface{}, filterNil bool) []interface{} {
	var res []interface{}
	var item interface{}
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
				res = append(res, item)
			}
		}
	default:
		panic("[arrayValues] arr type must be array, slice or map")
	}

	return res
}

// getTrimMask 去除mask字符.
func getTrimMask(characterMask []string) string {
	var mask string
	if len(characterMask) == 0 {
		mask = " \t\n\r\v\f\x00　"
	} else {
		mask = strings.Join(characterMask, "")
	}
	return mask
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

// compareConditionMap 比对数组是否匹配条件.condition为条件字典,arr为要比对的数据数组.
func compareConditionMap(condition map[string]interface{}, arr interface{}) (res interface{}) {
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

// getMethod 获取对象的方法.
func getMethod(t interface{}, method string) reflect.Value {
	m, b := reflect.TypeOf(t).MethodByName(method)
	if !b {
		return reflect.ValueOf(nil)
	}
	return m.Func
}

// ValidFunc 检查是否函数,并且参数个数、类型是否正确.
// 返回有效的函数、有效的参数.
func ValidFunc(f interface{}, args ...interface{}) (vf reflect.Value, vargs []reflect.Value, err error) {
	vf = reflect.ValueOf(f)
	if vf.Kind() != reflect.Func {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %v is not the function", f)
	}

	tf := vf.Type()
	_len := len(args)
	if tf.NumIn() != _len {
		return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %d number of the argument is incorrect", _len)
	}

	vargs = make([]reflect.Value, _len)
	for i := 0; i < _len; i++ {
		typ := tf.In(i).Kind()
		if (typ != reflect.Interface) && (typ != reflect.TypeOf(args[i]).Kind()) {
			return reflect.ValueOf(nil), nil, fmt.Errorf("[validFunc] %d-td argument`s type is incorrect", i+1)
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

// getPidByInode 根据套接字的inode获取PID.须root权限.
func getPidByInode(inode string, procDirs []string) (pid int) {
	if len(procDirs) == 0 {
		procDirs, _ = filepath.Glob("/proc/[0-9]*/fd/[0-9]*")
	}

	re := regexp.MustCompile(inode)
	for _, item := range procDirs {
		path, _ := os.Readlink(item)
		out := re.FindString(path)
		if len(out) != 0 {
			pid, _ = strconv.Atoi(strings.Split(item, "/")[2])
			break
		}
	}

	return pid
}

// getProcessPathByPid 根据PID获取进程的执行路径.
func getProcessPathByPid(pid int) string {
	exe := fmt.Sprintf("/proc/%d/exe", pid)
	path, _ := os.Readlink(exe)
	return path
}
