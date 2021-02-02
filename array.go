package kgo

import (
	"math"
	"reflect"
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
		panic("[ArrayChunk]`arr type must be array or slice")
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
		panic("[ArrayColumn]`arr type must be array, slice, struct or map; but : " + val.Kind().String())
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
