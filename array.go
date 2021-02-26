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
		panic("[ArrayChunk]`arr type must be array|slice; but : " + val.Kind().String())
	}
}

// ArrayColumn 返回数组(切片/字典/结构体)中元素指定的一列.
// arr的元素必须是字典;
// columnKey为元素的字段名;
// 该方法效率较低.
func (ka *LkkArray) ArrayColumn(arr interface{}, columnKey string) []interface{} {
	val := reflect.ValueOf(arr)
	var res []interface{}
	var err error
	var item interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item, err = GetFieldValue(val.Index(i).Interface(), columnKey)
			if item != nil && err == nil {
				res = append(res, item)
			}
		}
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			item, err = GetFieldValue(reflect2Itf(val.Field(i)), columnKey)
			if item != nil && err == nil {
				res = append(res, item)
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item, err = GetFieldValue(val.MapIndex(k).Interface(), columnKey)
			if item != nil && err == nil {
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
	val := reflect.ValueOf(arr)
	typ := val.Kind()

	if typ != reflect.Array && typ != reflect.Slice && typ != reflect.Struct && typ != reflect.Map {
		panic("[ArrayKeyExists]`arr type must be array|slice|struct|map; but : " + typ.String())
	}

	if key == nil {
		return false
	}

	switch typ {
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
		panic("[ArrayReverse]`arr type must be array|slice; but : " + val.Kind().String())
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
		panic("[Implode]`arr type must be array|slice|struct|map; but : " + val.Kind().String())
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
		panic("[ArrayDiff]`arr1,arr2 type must be array|slice|map; but : " + typA.String() + "/" + typB.String())
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
				chkKey = isInt(k.Interface()) && toInt(k.Interface()) == i
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
				kv = toInt(k.Interface())
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
			kv = toStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == toStr(k2.Interface())
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
		panic("[ArrayIntersect]`arr1,arr2 type must be array|slice|map; but : " + typA.String() + "/" + typB.String())
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
				chkKey = isInt(k.Interface()) && toInt(k.Interface()) == i
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
				kv = toInt(k.Interface())
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
			kv = toStr(k.Interface())

			for _, k2 := range valB.MapKeys() {
				chkKey = kv == toStr(k2.Interface())
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

// ArrayUnique 移除数组(切片/字典)中重复的值,返回字典,保留键名.
func (ka *LkkArray) ArrayUnique(arr interface{}) map[interface{}]interface{} {
	var item interface{}
	var str, key string
	val := reflect.ValueOf(arr)
	chkMp := make(map[string]interface{})
	res := make(map[interface{}]interface{})
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			item = val.Index(i).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(md5Byte([]byte(str), 32))
			if _, ok := chkMp[key]; !ok {
				chkMp[key] = true
				res[i] = item
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item = val.MapIndex(k).Interface()
			str = fmt.Sprintf("%+v", item)
			key = string(md5Byte([]byte(str), 32))
			if _, ok := chkMp[key]; !ok {
				chkMp[key] = true
				res[reflect2Itf(k)] = item
			}
		}
	default:
		panic("[ArrayUnique]`arr type must be array|slice|map; but : " + val.Kind().String())
	}

	return res
}

// ArraySearchItem 从数组(切片/字典)中搜索对应元素(单个).
// arr为要查找的数组,元素必须为字典/结构体;condition为条件字典.
func (ka *LkkArray) ArraySearchItem(arr interface{}, condition map[string]interface{}) (res interface{}) {
	val := reflect.ValueOf(arr)
	typ := val.Kind()
	if typ != reflect.Array && typ != reflect.Slice && typ != reflect.Map {
		panic("[ArraySearchItem]`arr type must be array|slice|map; but : " + typ.String())
	}

	// 条件为空
	if len(condition) == 0 {
		return
	}

	switch typ {
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
	}

	return
}

// ArraySearchMutil 从数组(切片/字典)中搜索对应元素(多个).
// arr为要查找的数组,元素必须为字典/结构体;condition为条件字典.
func (ka *LkkArray) ArraySearchMutil(arr interface{}, condition map[string]interface{}) (res []interface{}) {
	val := reflect.ValueOf(arr)
	typ := val.Kind()
	if typ != reflect.Array && typ != reflect.Slice && typ != reflect.Map {
		panic("[ArraySearchMutil]`arr type must be array|slice|map; but : " + typ.String())
	}

	// 条件为空
	if len(condition) == 0 {
		return
	}

	var item interface{}
	switch typ {
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
	}

	return
}

// ArrayShuffle 打乱数组/切片排序.
func (ka *LkkArray) ArrayShuffle(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	typ := val.Kind()
	if typ != reflect.Array && typ != reflect.Slice {
		panic("[ArrayShuffle]`arr type must be array|slice; but : " + typ.String())
	}

	num := val.Len()
	res := make([]interface{}, num)

	for i := 0; i < num; i++ {
		res[i] = val.Index(i).Interface()
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(num, func(i, j int) {
		res[i], res[j] = res[j], res[i]
	})

	return res
}

// IsEqualArray 两个数组/切片是否相同(不管元素顺序);expected, actual 是要比较的数组/切片.
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
		panic("[IsEqualArray]`expected type must be array|slice")
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
		panic("[IsEqualArray]`actual type must be array|slice")
	}

	return reflect.DeepEqual(expectedMap, actualMap)
}

// Length 获取数组/切片的长度;结果为-1表示变量不是数组或切片.
func (ka *LkkArray) Length(val interface{}) int {
	return lenArrayOrSlice(val, 3)
}

// IsArray 变量是否数组.
func (ka *LkkArray) IsArray(val interface{}) bool {
	l := lenArrayOrSlice(val, 1)
	return l >= 0
}

// IsSlice 变量是否切片.
func (ka *LkkArray) IsSlice(val interface{}) bool {
	l := lenArrayOrSlice(val, 2)
	return l >= 0
}

// IsArrayOrSlice 变量是否数组或切片.
func (ka *LkkArray) IsArrayOrSlice(val interface{}) bool {
	l := lenArrayOrSlice(val, 3)
	return l >= 0
}
