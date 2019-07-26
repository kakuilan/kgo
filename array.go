package kgo

import (
	"fmt"
	"reflect"
)

// IsArrayOrSlice 检查变量是否数组或切片;chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (ka *LkkArray) IsArrayOrSlice(data interface{}, chkType uint8) int {
	return isArrayOrSlice(data, chkType)
}

// IsMap 检查变量是否字典
func (ka *LkkArray) IsMap(data interface{}) bool {
	return isMap(data)
}

// InArray 元素是否在数组(切片/字典)内
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
		panic("[InArray]haystack type muset be array, slice or map")
	}

	return false
}

// ArrayFill 用给定的值填充数组
func (ka *LkkArray) ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	m := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		m[startIndex] = value
		startIndex++
	}
	return m
}

// ArrayFlip 交换数组中的键和值
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
			res[val.MapIndex(k).Interface()] = k
		}
	default:
		panic("[ArrayFlip]arr type muset be array, slice or map")
	}

	return res
}

// ArrayKeys 返回数组中所有的键名
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
		panic("[ArrayValues]arr type muset be array, slice or map")
	}

	return res
}

// ArrayValues 返回数组中所有的值
func (ka *LkkArray) ArrayValues(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	res := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			res[i] = val.Index(i).Interface()
		}
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = val.MapIndex(k).Interface()
		}
	default:
		panic("[ArrayValues]arr type muset be array, slice or map")
	}

	return res
}

// SliceMerge 合并一个或多个数组/切片;filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素
func (ka *LkkArray) SliceMerge(filterNil bool, ss ...interface{}) []interface{} {
	var res []interface{}
	switch len(ss) {
	case 0:
		break
	default:
		n := 0
		for i, v := range ss {
			chkLen := isArrayOrSlice(v, 3)
			if chkLen == -1 {
				msg := fmt.Sprintf("[SliceMerge]ss type muset be array or slice, but [%d]th item not is.", i)
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

// MapMerge 合并字典,相同的键名后面的值将覆盖前一个值;key2Str是否将键转换为字符串
func (ka *LkkArray) MapMerge(key2Str bool, ss ...interface{}) map[interface{}]interface{} {
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
				msg := fmt.Sprintf("[MapMerge]ss type muset be map, but [%d]th item not is.", i)
				panic(msg)
			}
		}
	}
	return res
}
