package kgo

import (
	"fmt"
	"math"
	"reflect"
)

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

// ArrayFill 用给定的值value填充数组,num为插入元素的数量
func (ka *LkkArray) ArrayFill(value interface{}, num int) []interface{} {
	if num <= 0 {
		return nil
	}

	var res []interface{} = make([]interface{}, num)
	for i := 0; i < num; i++ {
		res[i] = value
	}

	return res
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

// SliceMerge 合并一个或多个数组/切片;filterNil是否过滤空元素(nil,''),true时排除空元素,false时保留空元素;ss是元素为数组/切片的切片
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

// MapMerge 合并字典,相同的键名后面的值将覆盖前一个值;key2Str是否将键转换为字符串;ss是元素为数组/切片的切片
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

// ArrayChunk 将一个数组分割成多个,size为每个子数组的长度
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
		panic("[ArrayChunk]arr type muset be array or slice")
	}
}

// ArrayPad 以指定长度将一个值item填充进数组
// 如果 size 为正，则填补到数组的右侧，如果为负则从左侧开始填补。如果 size 的绝对值小于或等于 arr 数组的长度则没有任何填补
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
		panic("[ArrayPad]arr type muset be array, slice")
	}
}
