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

// ArrayKeys 返回数组(切片/字典/结构体)中所有的键名;如果是结构体,只返回公开的字段.
func (ka *LkkArray) ArrayKeys(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	typ := val.Kind()
	if typ != reflect.Array && typ != reflect.Slice && typ != reflect.Struct && typ != reflect.Map {
		panic("[ArrayKeys]`arr type must be array|slice|map|struct; but : " + typ.String())
	}

	var res []interface{}
	switch typ {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			res = append(res, i)
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			res = append(res, k)
		}
	case reflect.Struct:
		var t = val.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			// 不能访问未导出的字段
			if field.PkgPath != "" {
				continue
			}
			res = append(res, field.Name)
		}
	}

	return res
}

// ArrayValues 返回arr(数组/切片/字典/结构体)中所有的值;如果是结构体,只返回公开字段的值.
// filterZero 是否过滤零值元素(nil,false,0,"",[]),true时排除零值元素,false时保留零值元素.
func (ka *LkkArray) ArrayValues(arr interface{}, filterZero bool) []interface{} {
	return arrayValues(arr, filterZero)
}

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
	case reflect.Map:
		for _, k := range val.MapKeys() {
			item, err = GetFieldValue(val.MapIndex(k).Interface(), columnKey)
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
	default:
		panic("[ArrayColumn]`arr type must be array|slice|map|struct; but : " + val.Kind().String())
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
		panic("[ArrayKeyExists]`arr type must be array|slice|map|struct; but : " + typ.String())
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
	case reflect.Map:
		for _, k := range val.MapKeys() {
			if fmt.Sprintf("%s", key) == fmt.Sprintf("%s", k) || reflect.DeepEqual(key, k) {
				return true
			}
		}
	case reflect.Struct:
		field := val.FieldByName(fmt.Sprintf("%s", key))
		if field.IsValid() {
			return true
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
	default:
		panic("[Implode]`arr type must be array|slice|map|struct; but : " + val.Kind().String())
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

	if len(condition) > 0 {
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

	if len(condition) > 0 {
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

// IsEqualArray 两个数组/切片是否相同(不管元素顺序),且不会检查元素类型;
// arr1, arr2 是要比较的数组/切片.
func (ka *LkkArray) IsEqualArray(arr1, arr2 interface{}) bool {
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if (typA != reflect.Array && typA != reflect.Slice) || (typB != reflect.Array && typB != reflect.Slice) {
		panic("[IsEqualArray]`arr1,arr2 type must be array|slice; but : " + typA.String() + "/" + typB.String())
	}

	length := valA.Len()
	if length != valB.Len() {
		return false
	}

	var itmAStr, itmBstr string
	var format string = "%#v"
	expectedMap := make(map[string]bool)
	actualMap := make(map[string]bool)
	for i := 0; i < length; i++ {
		itmAStr = fmt.Sprintf(format, valA.Index(i).Interface())
		itmBstr = fmt.Sprintf(format, valB.Index(i).Interface())
		expectedMap[itmAStr] = true
		actualMap[itmBstr] = true
	}

	return reflect.DeepEqual(expectedMap, actualMap)
}

// IsEqualMap 两个字典是否相同(不管键顺序),且不会严格检查元素类型;
// arr1, arr2 是要比较的字典.
func (ka *LkkArray) IsEqualMap(arr1, arr2 interface{}) bool {
	valA := reflect.ValueOf(arr1)
	valB := reflect.ValueOf(arr2)
	typA := valA.Kind()
	typB := valB.Kind()

	if typA != reflect.Map || typB != reflect.Map {
		panic("[IsEqualMap]`arr1,arr2 type must be array|slice; but : " + typA.String() + "/" + typB.String())
	}

	length := valA.Len()
	if length != valB.Len() {
		return false
	}

	var key string
	expectedMap := make(map[string]interface{})
	actualMap := make(map[string]interface{})

	for _, k := range valA.MapKeys() {
		key = fmt.Sprintf("%v", k)
		expectedMap[key] = fmt.Sprintf("%#v", valA.MapIndex(k).Interface())
	}
	for _, k := range valB.MapKeys() {
		key = fmt.Sprintf("%v", k)
		actualMap[key] = fmt.Sprintf("%#v", valB.MapIndex(k).Interface())
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

// IsMap 变量是否字典.
func (ka *LkkArray) IsMap(val interface{}) bool {
	return isMap(val)
}

// IsStruct 变量是否结构体.
func (ka *LkkArray) IsStruct(val interface{}) bool {
	return isStruct(val)
}

// DeleteSliceItems 删除数组/切片的元素,返回一个新切片.
// ids为多个元素的索引(0~len(val)-1);
// del为删除元素的数量.
func (ka *LkkArray) DeleteSliceItems(val interface{}, ids ...int) (res []interface{}, del int) {
	sl := reflect.ValueOf(val)
	styp := sl.Kind()

	if styp != reflect.Array && styp != reflect.Slice {
		panic("[DeleteSliceItems]`val type must be array|slice")
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

// InArray 元素needle是否在数组haystack(切片/字典)内.
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
		panic("[InArray]`haystack type must be array|slice|map")
	}

	return false
}

// InIntSlice 是否在整型切片内.
func (ka *LkkArray) InIntSlice(i int, list []int) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}

// InInt64Slice 是否在64位整型切片内.
func (ka *LkkArray) InInt64Slice(i int64, list []int64) bool {
	for _, item := range list {
		if item == i {
			return true
		}
	}
	return false
}

// InStringSlice 是否在字符串切片内.
func (ka *LkkArray) InStringSlice(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}

// SliceFill 用给定的值val填充切片,num为插入元素的数量.
func (ka *LkkArray) SliceFill(val interface{}, num int) []interface{} {
	if num <= 0 {
		return nil
	}

	var res = make([]interface{}, num)
	for i := 0; i < num; i++ {
		res[i] = val
	}

	return res
}

// ArrayFlip 交换数组(切片/字典)中的键和值.
func (ka *LkkArray) ArrayFlip(arr interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	val := reflect.ValueOf(arr)
	var key interface{}
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			key = val.Index(i).Interface()
			if key != nil && fmt.Sprintf("%v", key) != "" {
				res[key] = i
			}
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			key = val.MapIndex(k).Interface()
			if key != nil && fmt.Sprintf("%v", key) != "" {
				res[key] = k
			}
		}
	default:
		panic("[ArrayFlip]`arr type must be array|slice|map")
	}

	return res
}

// MergeSlice 合并一个或多个数组/切片.
// filterZero 是否过滤零值元素(nil,false,0,'',[]),true时排除零值元素,false时保留零值元素.
// ss是元素为数组/切片的切片.
func (ka *LkkArray) MergeSlice(filterZero bool, ss ...interface{}) []interface{} {
	var res []interface{}
	if len(ss) > 0 {
		n := 0
		for i, v := range ss {
			chkLen := lenArrayOrSlice(v, 3)
			if chkLen == -1 {
				msg := fmt.Sprintf("[MergeSlice]`ss type must be array|slice, but [%d]th item not is.", i)
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
					if !filterZero || (filterZero && !val.Index(i).IsZero()) {
						res = append(res, item)
					}
				}
			}
		}
	}

	return res
}

// MergeMap 合并字典,相同的键名时,后面的值将覆盖前一个值.
// ss是元素为字典的切片.
func (ka *LkkArray) MergeMap(ss ...interface{}) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})
	if len(ss) > 0 {
		for i, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Map:
				for _, k := range val.MapKeys() {
					res[reflect2Itf(k)] = val.MapIndex(k).Interface()
				}
			default:
				panic(fmt.Sprintf("[MergeMap]`ss type must be map, but [%d]th item not is.", i))
			}
		}
	}

	return res
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
			return ka.SliceFill(item, size)
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
		panic("[ArrayPad]`arr type must be array|slice")
	}
}

// ArrayRand 从数组(切片/字典)中随机取出num个元素.
func (ka *LkkArray) ArrayRand(arr interface{}, num int) []interface{} {
	val := reflect.ValueOf(arr)
	typ := val.Kind()
	if typ != reflect.Array && typ != reflect.Slice && typ != reflect.Map {
		panic("[ArrayRand]`arr type must be array|slice|map")
	}
	if num < 1 {
		panic("[ArrayRand]`num cannot be less than 1")
	}

	length := val.Len()
	if length == 0 {
		return nil
	}
	if num > length {
		num = length
	}
	res := make([]interface{}, num)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	switch typ {
	case reflect.Array, reflect.Slice:
		for i, v := range r.Perm(length) {
			if i < num {
				res[i] = val.Index(v).Interface()
			} else {
				break
			}
		}
	case reflect.Map:
		var ks []reflect.Value
		ks = append(ks, val.MapKeys()...)
		for i, v := range r.Perm(length) {
			if i < num {
				res[i] = val.MapIndex(ks[v]).Interface()
			} else {
				break
			}
		}
	}

	return res
}

// CutSlice 裁剪切片,返回根据offset(起始位置)和size(数量)参数所指定的arr(数组/切片)中的一段切片.
func (ka *LkkArray) CutSlice(arr interface{}, offset, size int) []interface{} {
	if size < 1 {
		panic("[CutSlice]`size cannot be less than 1")
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
		panic("[CutSlice]`arr type must be array|slice")
	}
}

// NewStrMapItf 新建[字符-接口]字典.
func (ka *LkkArray) NewStrMapItf() map[string]interface{} {
	return make(map[string]interface{})
}

// NewStrMapStr 新建[字符-字符]字典.
func (ka *LkkArray) NewStrMapStr() map[string]string {
	return make(map[string]string)
}

// CopyStruct 将resources的值拷贝到dest目标结构体;
// 要求dest必须是结构体指针,resources为多个源结构体;若resources存在多个相同字段的元素,结果以最后的为准;
// 只简单核对字段名,无错误处理,需开发自行检查dest和resources字段类型才可操作.
func (ka *LkkArray) CopyStruct(dest interface{}, resources ...interface{}) interface{} {
	dVal := reflect.ValueOf(dest)
	dTyp := reflect.TypeOf(dest)

	if dTyp.Kind() != reflect.Ptr {
		return nil
	}

	//dest是指针,需要.Elem()取得指针指向的value
	dVal = dVal.Elem()
	dTyp = dTyp.Elem()

	// 非结构体
	if dVal.Kind() != reflect.Struct {
		return nil
	}

	//目标结构体可导出的字段
	var dFields = make(map[string]reflect.Type, dTyp.NumField())
	reflectTypesMap(dTyp, dFields)

	var field string
	for _, resource := range resources {
		rVal := reflect.ValueOf(resource)
		rTyp := rVal.Type()
		if rTyp.Kind() == reflect.Ptr {
			rVal = rVal.Elem()
			rTyp = rTyp.Elem()
		}

		if rVal.Kind() == reflect.Struct {
			for i := 0; i < rTyp.NumField(); i++ {
				field = rTyp.Field(i).Name
				if typ, ok := dFields[field]; ok {
					dVal.FieldByName(field).Set(rVal.FieldByName(field).Convert(typ))
				}
			}
		}
	}

	return dest
}
