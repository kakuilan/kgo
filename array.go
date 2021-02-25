package kgo

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// ArrayChunk 将一个数组/切片分割成多个,size为每个子数组的长度.
func (ka *LkkArray) ArrayChunk(arr interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("[ArrayChunk]`size cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}

		chunkNum := int(math.Ceil(float64(length) / float64(size)))
		var res [][]interface{}
		var item []interface{}
		var start int
		for i, end := 0, 0; chunkNum > 0; chunkNum-- {
			end = (i + 1) * size
			if end > length {
				end = length
			}

			item = nil
			start = i * size
			for ; start < end; start++ {
				item = append(item, val.Index(start).Interface())
			}
			if item != nil {
				res = append(res, item)
			}

			i++
		}

		return res
	default:
		panic("[ArrayChunk]`arr type must be array|slice")
	}
}

// ArrayColumn 返回数组(切片/字典/结构体)中元素指定的一列.
// arr的元素必须是字典;
// columnKey为元素的字段名;
// 该方法效率较低.
func (ka *LkkArray) ArrayColumn(arr interface{}, columnKey string) []interface{} {
	val := reflect.ValueOf(arr)
	var res []interface{}
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = GetFieldValue(val.Index(i).Interface(), columnKey)
			if item != nil {
				res = append(res, item)
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			item = GetFieldValue(val.Field(i).Interface(), columnKey)
			if item != nil {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = GetFieldValue(val.MapIndex(k).Interface(), columnKey)
			if item != nil {
				res = append(res, item)
			}
		}
	default:
		panic("[ArrayColumn]`arr type must be array|slice|struct|map; but : " + val.Kind().String())
	}

	return res
}

// SlicePush 将一个或多个元素压入切片的末尾(入栈),返回处理之后切片的元素个数.
func (ka *LkkArray) SlicePush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}

// SlicePop 弹出切片最后一个元素(出栈),并返回该元素.
func (ka *LkkArray) SlicePop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// SliceUnshift 在切片开头插入一个或多个元素,返回处理之后切片的元素个数.
func (ka *LkkArray) SliceUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

// SliceShift 将切片开头的元素移出,并返回该元素.
func (ka *LkkArray) SliceShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	e := (*s)[0]
	*s = (*s)[1:]
	return e
}

// ArrayKeyExists 检查arr(数组/切片/字典/结构体)里是否有key指定的键名(索引/字段).
func (ka *LkkArray) ArrayKeyExists(key interface{}, arr interface{}) bool {
	if key == nil {
		return false
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		var keyInt int
		var keyIsInt, ok bool
		if keyInt, ok = key.(int); ok {
			keyIsInt = true
		}

		length := val.Len()
		if keyIsInt && length > 0 && keyInt >= 0 && keyInt < length {
			return true
		}
	case reflect.Struct:
		field := val.FieldByName(fmt.Sprintf("%s", key))
		if field.IsValid() {
			return true
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if fmt.Sprintf("%s", key) == fmt.Sprintf("%s", k) || reflect.DeepEqual(key, k) {
				return true
			}
		}
	default:
		panic("[ArrayKeyExists]`arr type must be array|slice|struct|map; but : " + val.Kind().String())
	}
	return false
}

// ArrayReverse 返回单元顺序相反的数组(仅限数组和切片).
func (ka *LkkArray) ArrayReverse(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		res := make([]interface{}, length)
		i, j := 0, length-1
		for ; i < j; i, j = i+1, j-1 {
			res[i], res[j] = val.Index(j).Interface(), val.Index(i).Interface()
		}
		if length > 0 && res[j] == nil {
			res[j] = val.Index(j).Interface()
		}

		return res
	default:
		panic("[ArrayReverse]`arr type must be array|slice")
	}
}

// Implode 用delimiter将数组(数组/切片/字典/结构体)的值连接为一个字符串.
func (ka *LkkArray) Implode(delimiter string, arr interface{}) string {
	val := reflect.ValueOf(arr)
	var buf bytes.Buffer
	var length, j int
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length = val.Len()
		if length == 0 {
			return ""
		}
		j = length
		for i := 0; i < length; i++ {
			buf.WriteString(toStr(val.Index(i).Interface()))
			if j--; j > 0 {
				buf.WriteString(delimiter)
			}
		}
	case reflect.Struct:
		length = val.NumField()
		if length == 0 {
			return ""
		}
		j = length
		for i := 0; i < length; i++ {
			j--
			if val.Field(i).CanInterface() {
				buf.WriteString(toStr(val.Field(i).Interface()))
				if j > 0 {
					buf.WriteString(delimiter)
				}
			}
		}
	case reflect.Map:
		length = len(val.MapKeys())
		if length == 0 {
			return ""
		}
		for _, k := range val.MapKeys() {
			buf.WriteString(toStr(val.MapIndex(k).Interface()))
			if length--; length > 0 {
				buf.WriteString(delimiter)
			}
		}
	default:
		panic("[Implode]`arr type must be array|slice|struct|map")
	}

	return buf.String()
}

// JoinStrings 使用分隔符delimiter连接字符串切片strs.效率比Implode高.
func (ka *LkkArray) JoinStrings(delimiter string, strs []string) (res string) {
	length := len(strs)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	res = sb.String()

	return
}

// JoinInts 使用分隔符delimiter连接整数切片.
func (ka *LkkArray) JoinInts(delimiter string, ints []int) (res string) {
	length := len(ints)
	if length == 0 {
		return
	}

	var sb strings.Builder
	for _, num := range ints {
		sb.WriteString(strconv.Itoa(num))
		if length--; length > 0 {
			sb.WriteString(delimiter)
		}
	}
	res = sb.String()

	return
}

// UniqueInts 移除整数切片中的重复值.
func (ka *LkkArray) UniqueInts(ints []int) (res []int) {
	sort.Ints(ints)
	var last int
	for i, num := range ints {
		if i == 0 || num != last {
			res = append(res, num)
		}
		last = num
	}
	return
}

// Unique64Ints 移除64位整数切片中的重复值.
func (ka *LkkArray) Unique64Ints(ints []int64) (res []int64) {
	seen := make(map[int64]bool)
	for _, num := range ints {
		if _, ok := seen[num]; !ok {
			seen[num] = true
			res = append(res, num)
		}
	}
	return
}

// UniqueStrings 移除字符串切片中的重复值.
func (ka *LkkArray) UniqueStrings(strs []string) (res []string) {
	sort.Strings(strs)
	var last string
	for _, str := range strs {
		if str != last {
			res = append(res, str)
		}
		last = str
	}

	return
}
