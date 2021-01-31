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
