package kgo

import (
	"fmt"
	"reflect"
)

// InArray 元素是否在数组(切片/字典)内
func (ka *LkkArray) InArray(needle interface{}, haystack interface{}) bool {
	val := reflect.ValueOf(haystack)
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
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
		panic("[InArray]haystack type muset be slice, array or map")
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
	case reflect.Slice, reflect.Array:
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
		panic("[ArrayFlip]arr type muset be slice, array or map")
	}

	return res
}

// ArrayKeys 返回数组中所有的键名
func (ka *LkkArray) ArrayKeys(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	res := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res[i] = i
		}
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = k
		}
	default:
		panic("[ArrayValues]arr type muset be slice, array or map")
	}

	return res
}

// ArrayValues 返回数组中所有的值
func (ka *LkkArray) ArrayValues(arr interface{}) []interface{} {
	val := reflect.ValueOf(arr)
	res := make([]interface{}, val.Len())
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			res[i] = val.Index(i).Interface()
		}
	case reflect.Map:
		for i, k := range val.MapKeys() {
			res[i] = val.MapIndex(k).Interface()
		}
	default:
		panic("[ArrayValues]arr type muset be slice, array or map")
	}

	return res
}

// SliceMerge 合并一个或多个数组/切片
func (ka *LkkArray) SliceMerge0(ss ...[]interface{}) []interface{} {
	var res []interface{}
	switch len(ss) {
	case 0:
		break
	case 1:
		res = ss[0]
	default:
		n := 0
		for _, v := range ss {
			n += len(v)
		}
		res = make([]interface{}, 0, n)
		for _, v := range ss {
			res = append(res, v...)
		}
	}

	return res
}

func (ka *LkkArray) SliceMerge(ss ...interface{}) []interface{} {
	var res []interface{}
	switch len(ss) {
	case 0:
		break
	case 1:
		if isArrayOrSlice(ss[0], 3) == -1 {
			panic("[SliceMerge]ss type muset be array or slice")
		} else {
			res = append(res, ss[0])
		}
	default:
		n := 0
		for i, v := range ss {
			chkLen := isArrayOrSlice(v, 3)
			if chkLen == -1 {
				msg := fmt.Sprintf("[SliceMerge]ss type muset be array or slice, but [%d] item not is.", i)
				panic(msg)
			} else {
				n += chkLen
			}
		}
		res = make([]interface{}, 0, n)
		for _, v := range ss {
			val := reflect.ValueOf(v)
			switch val.Kind() {
			case reflect.Slice, reflect.Array:
				//TODO
			}
		}
	}
	return res
}
