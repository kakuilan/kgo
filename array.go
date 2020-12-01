package kgo

import (
	"bytes"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// InArray 元素needle 是否在数组haystack(切片/字典)内.
func (ka *LkkArray) InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if reflect.DeepEqual(needle, val.Index(i).Interface()) {
				return true
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if reflect.DeepEqual(needle, val.MapIndex(k).Interface()) {
				return true
			}
		}
	default:
		panic("[InArray]haystack type must be array, slice or map")
	}

	return false
}

// InIntSlice 是否在整型数组/切片内.
func (ka *LkkArray) InIntSlice(i int, list []int) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}

// InInt64Slice 是否在64位整型数组/切片内.
func (ka *LkkArray) InInt64Slice(i int64, list []int64) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}

// InStringSlice 是否在字符串数组/切片内.
func (ka *LkkArray) InStringSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

// ArrayFill 用给定的值value填充数组,num为插入元素的数量.
func (ka *LkkArray) ArrayFill(value interface{}, num int) []interface{} {
	if num <= 0 {
		return nil
	}

	var res = make([]interface{}, num)
	for i := 0; i < num; i++ {
		res[i] = value
	}

	return res
}

// ArrayFlip 交换数组(切片/字典)中的键和值.
func (ka *LkkArray) ArrayFlip(arr interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() != nil && fmt.Sprintf("%v", val.Index(i).Interface()) != "" {
				res[val.Index(i).Interface()] = i
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if val.MapIndex(k).Interface() != nil && fmt.Sprintf("%v", val.MapIndex(k).Interface()) != "" {
				res[val.MapIndex(k).Interface()] = k
			}
		}
	default:
		panic("[ArrayFlip]arr type must be array, slice or map")
	}

	return res
}

// ArrayKeys 返回数组(切片/字典)中所有的键名.
func (ka *LkkArray) ArrayKeys(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	res := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			res[i] = i
		}
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = k
		}
	default:
		panic("[ArrayValues]arr type must be array, slice or map")
	}

	return res
}

// ArrayValues 返回数组(切片/字典)中所有的值.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素.
func (ka *LkkArray) ArrayValues(arr interface{}, filterNil bool) []interface{} {
	return arrayValues(arr, filterNil)
}

// MergeSlice 合并一个或多个数组/切片.
// filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素;ss是元素为数组/切片的数组.
func (ka *LkkArray) MergeSlice(filterNil bool, ss ...interface{}) []interface{} {
	var res []interface{}
	switch len(ss) {
	case 0:
		break
	default:
		n := 0
		for i, v := range ss {
			chkLen := isArrayOrSlice(v, 3)
			if chkLen == -1 {
				msg := fmt.Sprintf("[MergeSlice]ss type must be array or slice, but [%d]th item not is.", i)
				panic(msg)
			} else {
				n += chkLen
			}
		}
		res = make([]interface{}, 0, n)
		var item interface{}
		for _, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Array, reflect.Slice:
				for i := 0; i < val.Len(); i++ {
					item = val.Index(i).Interface()
					if !filterNil || (filterNil && item != nil && fmt.Sprintf("%v", item) != "") {
						res = append(res, item)
					}
				}
			}
		}
	}
	return res
}

// MergeMap 合并字典.
// 相同的键名时,后面的值将覆盖前一个值;key2Str是否将键转换为字符串;ss是元素为字典的数组.
func (ka *LkkArray) MergeMap(key2Str bool, ss ...interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	switch len(ss) {
	case 0:
		break
	default:
		for i, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Map:
				for _, k := range val.MapKeys() {
					if key2Str {
						res[k.String()] = val.MapIndex(k).Interface()
					} else {
						res[k] = val.MapIndex(k).Interface()
					}
				}
			default:
				msg := fmt.Sprintf("[MergeMap]ss type must be map, but [%d]th item not is.", i)
				panic(msg)
			}
		}
	}
	return res
}

// ArrayChunk 将一个数组/切片分割成多个,size为每个子数组的长度.
func (ka *LkkArray) ArrayChunk(arr interface{}, size int) [][]interface{} {
	if size < 1 {
		panic("[ArrayChunk]size: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}

		chunks := int(math.Ceil(float64(length) / float64(size)))
		var res [][]interface{}
		var item []interface{}
		var start int
		for i, end := 0, 0; chunks > 0; chunks-- {
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
		panic("[ArrayChunk]arr type must be array or slice")
	}
}

// ArrayPad 以指定长度将一个值item填充进arr数组/切片.
// 若 size 为正，则填补到数组的右侧，如果为负则从左侧开始填补;
// 若 size 的绝对值小于或等于 arr 数组的长度则没有任何填补.
func (ka *LkkArray) ArrayPad(arr interface{}, size int, item interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 && size > 0 {
			return ka.ArrayFill(item, size)
		}

		orig := make([]interface{}, length)
		for i := 0; i < length; i++ {
			orig[i] = val.Index(i).Interface()
		}

		if size == 0 || (size > 0 && size < length) || (size < 0 && size > -length) {
			return orig
		}

		n := size
		if size < 0 {
			n = -size
		}
		n -= length
		items := make([]interface{}, n)
		for i := 0; i < n; i++ {
			items[i] = item
		}

		if size > 0 {
			return append(orig, items...)
		}
		return append(items, orig...)
	default:
		panic("[ArrayPad]arr type must be array, slice")
	}
}

// ArraySlice 返回根据 offset 和 size 参数所指定的 arr 数组/切片中的一段切片.
func (ka *LkkArray) ArraySlice(arr interface{}, offset, size int) []interface{} {
	if size < 1 {
		panic("[ArraySlice]size: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 || (offset > 0 && offset > length-1) {
			return nil
		}

		items := make([]interface{}, length)
		for i := 0; i < val.Len(); i++ {
			items[i] = val.Index(i).Interface()
		}

		if offset < 0 {
			offset = offset%length + length
		}
		end := offset + size
		if end < length {
			return items[offset:end]
		}
		return items[offset:]
	default:
		panic("[ArraySlice]arr type must be array or slice")
	}
}

// ArrayRand 从数组/切片中随机取出num个单元.
func (ka *LkkArray) ArrayRand(arr interface{}, num int) []interface{} {
	if num < 1 {
		panic("[ArrayRand]num: cannot be less than 1")
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return nil
		}
		if num > length {
			num = length
		}
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		res := make([]interface{}, num)
		for i, v := range r.Perm(length) {
			if i < num {
				res[i] = val.Index(v).Interface()
			} else {
				break
			}
		}
		return res
	default:
		panic("[ArrayRand]arr type must be array or slice")
	}
}

// ArrayColumn 返回数组(切片/字典)中元素指定的一列.
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
			item = val.Index(i).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						res = append(res, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s slice item must be map")
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			itemVal := reflect.ValueOf(item)
			switch itemVal.Kind() {
			case reflect.Map:
				for _, subKey := range itemVal.MapKeys() {
					if fmt.Sprintf("%s", subKey) == columnKey {
						res = append(res, itemVal.MapIndex(subKey).Interface())
						break
					}
				}
			default:
				panic("[ArrayColumn]arr`s map item must be map")
			}
		}
	default:
		panic("[ArrayColumn]arr type must be array, slice or map")
	}

	return res
}

// ArrayPush 将一个或多个元素压入数组的末尾(入栈),返回处理之后数组的元素个数.
func (ka *LkkArray) ArrayPush(s *[]interface{}, elements ...interface{}) int {
	*s = append(*s, elements...)
	return len(*s)
}

// ArrayPop 弹出数组最后一个元素(出栈),并返回该元素.
func (ka *LkkArray) ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayUnshift 在数组开头插入一个或多个元素,返回处理之后数组的元素个数.
func (ka *LkkArray) ArrayUnshift(s *[]interface{}, elements ...interface{}) int {
	*s = append(elements, *s...)
	return len(*s)
}

// ArrayShift 将数组开头的元素移出数组,并返回该元素.
func (ka *LkkArray) ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	e := (*s)[0]
	*s = (*s)[1:]
	return e
}

// ArrayKeyExists 检查数组里是否有指定的键名或索引.
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
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if fmt.Sprintf("%s", key) == fmt.Sprintf("%s", k) || reflect.DeepEqual(key, k) {
				return true
			}
		}
	default:
		panic("[ArrayKeyExists]arr type must be array, slice or map")
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
		panic("[ArrayReverse]arr type must be array, slice")
	}
}

// Implode 用delimiter将数组(数组/切片/字典)的值连接为一个字符串.
func (ka *LkkArray) Implode(delimiter string, arr interface{}) string {
	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		length := val.Len()
		if length == 0 {
			return ""
		}
		var buf bytes.Buffer
		j := length
		for i := 0; i < length; i++ {
			buf.WriteString(fmt.Sprintf("%s", val.Index(i)))
			if j--; j > 0 {
				buf.WriteString(delimiter)
			}
		}
		return buf.String()
	case reflect.Map:
		length := len(val.MapKeys())
		if length == 0 {
			return ""
		}
		var buf bytes.Buffer
		for _, k := range val.MapKeys() {
			buf.WriteString(fmt.Sprintf("%s", val.MapIndex(k).Interface()))
			if length--; length > 0 {
				buf.WriteString(delimiter)
			}
		}

		return buf.String()
	default:
		panic("[Implode]arr type must be array, slice")
	}
}

// JoinStrings 使用分隔符delimiter连接字符串数组.效率比Implode高.
func (ka *LkkArray) JoinStrings(strs []string, delimiter string) (res string) {
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

// JoinInts 使用分隔符delimiter连接整数数组.
func (ka *LkkArray) JoinInts(ints []int, delimiter string) (res string) {
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

// UniqueInts 移除整数数组中的重复值.
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

// Unique64Ints 移除64位整数数组中的重复值.
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

// UniqueStrings 移除字符串数组中的重复值.
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

// ArrayDiff 计算数组(数组/切片/字典)的交集,返回在 arr1 中但不在 arr2 里的元素,注意会同时返回键.
// compareType为两个数组的比较方式,枚举类型,有:
// COMPARE_ONLY_VALUE 根据元素值比较, 返回在 arr1 中但是不在arr2 里的值;
// COMPARE_ONLY_KEY 根据 arr1 中的键名和 arr2 进行比较,返回不同键名的项;
// COMPARE_BOTH_KEYVALUE 同时比较键和值.
func (ka *LkkArray) ArrayDiff(arr1, arr2 interface{}, compareType LkkArrCompareType) map[interface{}]interface{} {
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if (typA != reflect.Array && typA != reflect.Slice && typA != reflect.Map) || (typB != reflect.Array && typB != reflect.Slice && typB != reflect.Map) {
		panic("[ArrayDiff] arr1, arr2 type must be array, slice or map")
	}

	lenA := valA.Len()
	lenB := valB.Len()
	if lenA == 0 {
		return nil
	}

	resMap := make(map[interface{}]interface{})
	var iteA interface{}
	var chkKey bool
	var chkVal bool
	var chkRes bool

	if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Array || typB == reflect.Slice) {
		//两者都是数组/切片
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = true

			if compareType == COMPARE_BOTH_KEYVALUE {
				if i < lenB {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(i).Interface())
				}
			} else if compareType == COMPARE_ONLY_KEY {
				chkRes = lenB > 0 && i >= lenB
			} else if compareType == COMPARE_ONLY_VALUE {
				for j := 0; j < lenB; j++ {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(j).Interface())
					if !chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Map) {
		//A是数组/切片,B是字典
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = true

			for _, k := range valB.MapKeys() {
				chkKey = isInt(k.Interface()) && KConv.ToInt(k.Interface()) == i
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k).Interface())

				if compareType == COMPARE_ONLY_KEY && chkKey {
					chkRes = false
					break
				} else if compareType == COMPARE_ONLY_VALUE && chkVal {
					chkRes = false
					break
				} else if compareType == COMPARE_BOTH_KEYVALUE && (chkKey && chkVal) {
					chkRes = false
					break
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Array || typB == reflect.Slice) {
		//A是字典,B是数组/切片
		var kv int
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = true

			if isInt(k.Interface()) {
				kv = KConv.ToInt(k.Interface())
			} else {
				kv = -1
			}

			if compareType == COMPARE_BOTH_KEYVALUE {
				if kv >= 0 && kv < lenB {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(kv).Interface())
				}
			} else if compareType == COMPARE_ONLY_KEY {
				chkRes = (kv < 0 || kv >= lenB)
			} else if compareType == COMPARE_ONLY_VALUE {
				for i := 0; i < lenB; i++ {
					chkRes = !reflect.DeepEqual(iteA, valB.Index(i).Interface())
					if !chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Map) {
		//两者都是字典
		var kv string
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = true
			kv = KConv.ToStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == KConv.ToStr(k2.Interface())
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k2).Interface())

				if compareType == COMPARE_ONLY_KEY && chkKey {
					chkRes = false
					break
				} else if compareType == COMPARE_ONLY_VALUE && chkVal {
					chkRes = false
					break
				} else if compareType == COMPARE_BOTH_KEYVALUE && (chkKey || chkVal) {
					chkRes = false
					break
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	}

	return resMap
}

// ArrayIntersect 计算数组(数组/切片/字典)的交集,返回在 arr1 中且在 arr2 里的元素,注意会同时返回键.
// compareType为两个数组的比较方式,枚举类型,有:
// COMPARE_ONLY_VALUE 根据元素值比较, 返回在 arr1 中且在arr2 里的值;
// COMPARE_ONLY_KEY 根据 arr1 中的键名和 arr2 进行比较,返回相同键名的项;
// COMPARE_BOTH_KEYVALUE 同时比较键和值.
func (ka *LkkArray) ArrayIntersect(arr1, arr2 interface{}, compareType LkkArrCompareType) map[interface{}]interface{} {
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if (typA != reflect.Array && typA != reflect.Slice && typA != reflect.Map) || (typB != reflect.Array && typB != reflect.Slice && typB != reflect.Map) {
		panic("[ArrayIntersect] arr1, arr2 type must be array, slice or map")
	}

	lenA := valA.Len()
	lenB := valB.Len()
	if lenA == 0 || lenB == 0 {
		return nil
	}

	resMap := make(map[interface{}]interface{})
	var iteA interface{}
	var chkKey bool
	var chkVal bool
	var chkRes bool

	if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Array || typB == reflect.Slice) {
		//两者都是数组/切片
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = false

			if compareType == COMPARE_BOTH_KEYVALUE {
				if i < lenB {
					chkRes = reflect.DeepEqual(iteA, valB.Index(i).Interface())
				}
			} else if compareType == COMPARE_ONLY_KEY {
				chkRes = i < lenB
			} else if compareType == COMPARE_ONLY_VALUE {
				for j := 0; j < lenB; j++ {
					chkRes = reflect.DeepEqual(iteA, valB.Index(j).Interface())
					if chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Array || typA == reflect.Slice) && (typB == reflect.Map) {
		//A是数组/切片,B是字典
		for i := 0; i < lenA; i++ {
			iteA = valA.Index(i).Interface()
			chkRes = false

			for _, k := range valB.MapKeys() {
				chkKey = isInt(k.Interface()) && KConv.ToInt(k.Interface()) == i
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k).Interface())

				if compareType == COMPARE_ONLY_KEY && chkKey {
					chkRes = true
					break
				} else if compareType == COMPARE_ONLY_VALUE && chkVal {
					chkRes = true
					break
				} else if compareType == COMPARE_BOTH_KEYVALUE && (chkKey && chkVal) {
					chkRes = true
					break
				}
			}

			if chkRes {
				resMap[i] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Array || typB == reflect.Slice) {
		//A是字典,B是数组/切片
		var kv int
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = false

			if isInt(k.Interface()) {
				kv = KConv.ToInt(k.Interface())
			} else {
				kv = -1
			}

			if compareType == COMPARE_BOTH_KEYVALUE {
				if kv >= 0 && kv < lenB {
					chkRes = reflect.DeepEqual(iteA, valB.Index(kv).Interface())
				}
			} else if compareType == COMPARE_ONLY_KEY {
				chkRes = kv >= 0 && kv < lenB
			} else if compareType == COMPARE_ONLY_VALUE {
				for i := 0; i < lenB; i++ {
					chkRes = reflect.DeepEqual(iteA, valB.Index(i).Interface())
					if chkRes {
						break
					}
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	} else if (typA == reflect.Map) && (typB == reflect.Map) {
		//两者都是字典
		var kv string
		for _, k := range valA.MapKeys() {
			iteA = valA.MapIndex(k).Interface()
			chkRes = false
			kv = KConv.ToStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == KConv.ToStr(k2.Interface())
				chkVal = reflect.DeepEqual(iteA, valB.MapIndex(k2).Interface())

				if compareType == COMPARE_ONLY_KEY && chkKey {
					chkRes = true
					break
				} else if compareType == COMPARE_ONLY_VALUE && chkVal {
					chkRes = true
					break
				} else if compareType == COMPARE_BOTH_KEYVALUE && (chkKey && chkVal) {
					chkRes = true
					break
				}
			}

			if chkRes {
				resMap[k.Interface()] = iteA
			}
		}
	}

	return resMap
}

// ArrayUnique 移除数组中重复的值.
func (ka *LkkArray) ArrayUnique(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	var res []interface{}
	var item interface{}
	var str, key string
	mp := make(map[string]interface{})
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(md5Byte([]byte(str), 32))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(md5Byte([]byte(str), 32))
			if _, ok := mp[key]; !ok {
				mp[key] = true
				res = append(res, item)
			}
		}
	default:
		panic("[ArrayUnique]arr type must be array, slice or map")
	}

	return res
}

// ArraySearchItem 从数组中搜索对应元素(单个).
// arr为要查找的数组,元素必须为字典;condition为条件字典.
func (ka *LkkArray) ArraySearchItem(arr interface{}, condition map[string]interface{}) (res interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			res = compareConditionMap(condition, val.Index(i).Interface())
			if res != nil {
				return
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			res = compareConditionMap(condition, val.MapIndex(k).Interface())
			if res != nil {
				return
			}
		}
	default:
		panic("[ArraySearchItem]arr type must be array, slice or map")
	}

	return
}

// ArraySearchMutil 从数组中搜索对应元素(多个).
// arr为要查找的数组,元素必须为字典;condition为条件字典.
func (ka *LkkArray) ArraySearchMutil(arr interface{}, condition map[string]interface{}) (res []interface{}) {
	// 条件为空
	if len(condition) == 0 {
		return
	}

	val := reflect.ValueOf(arr)
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = compareConditionMap(condition, val.Index(i).Interface())
			if item != nil {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = compareConditionMap(condition, val.MapIndex(k).Interface())
			if item != nil {
				res = append(res, item)
			}
		}
	default:
		panic("[ArraySearchMutil]arr type must be array, slice or map")
	}

	return
}

// IsEqualArray 两个数组/切片是否相同(不管元素顺序).
// expected, actual 是要比较的数组/切片.
func (ka *LkkArray) IsEqualArray(expected, actual interface{}) bool {
	var itmAsStr string

	expectedMap := map[string]bool{}
	val := reflect.ValueOf(expected)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			itmAsStr = fmt.Sprintf("%#v", val.Index(i).Interface())
			expectedMap[itmAsStr] = true
		}
	default:
		panic("[IsEqualArray]expected type must be array or slice")
	}

	actualMap := map[string]bool{}
	val = reflect.ValueOf(actual)
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			itmAsStr = fmt.Sprintf("%#v", val.Index(i).Interface())
			actualMap[itmAsStr] = true
		}
	default:
		panic("[IsEqualArray]actual type must be array or slice")
	}

	return reflect.DeepEqual(expectedMap, actualMap)
}

// IsArrayOrSlice 检查变量是否数组或切片.
// chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (ka *LkkArray) IsArrayOrSlice(val interface{}, chkType uint8) int {
	return isArrayOrSlice(val, chkType)
}

// IsMap 检查变量是否字典.
func (ka *LkkArray) IsMap(val interface{}) bool {
	return isMap(val)
}

// DeleteSliceItems 删除数组/切片的元素,返回一个新切片.
// ids为多个元素的索引(0~len(val)-1);
// del为删除元素的数量.
func (ka *LkkArray) DeleteSliceItems(val interface{}, ids ...int) (res []interface{}, del int) {
	sl := reflect.ValueOf(val)
	styp := sl.Kind()

	if styp != reflect.Array && styp != reflect.Slice {
		panic("[DeleteSlice] val type must be array or slice")
	}

	slen := sl.Len()
	if slen == 0 {
		return
	}

	var item interface{}
	var chk bool
	for i := 0; i < slen; i++ {
		item = sl.Index(i).Interface()
		chk = true

		for _, v := range ids {
			if i == v {
				del++
				chk = false
				break
			}
		}

		if chk {
			res = append(res, item)
		}
	}

	return
}
