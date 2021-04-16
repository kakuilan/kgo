package kgo

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// AbsFloat 浮点型取绝对值.
func (kn *LkkNumber) AbsFloat(number float64) float64 {
	return math.Abs(number)
}

// AbsInt 整型取绝对值.
func (kn *LkkNumber) AbsInt(number int64) int64 {
	r := number >> 63
	return (number ^ r) - r
}

// Range 根据范围创建数组,包含指定的元素.
// start为起始元素值,end为末尾元素值.若start<end,返回升序的数组;若start>end,返回降序的数组.
func (kn *LkkNumber) Range(start, end int) []int {
	res := make([]int, kn.AbsInt(int64(end-start))+1)
	for i := range res {
		if end > start {
			res[i] = start + i
		} else {
			res[i] = start - i
		}
	}
	return res
}

// NumberFormat 以千位分隔符方式格式化一个数字.
// decimal为要保留的小数位数,point为小数点显示的字符,thousand为千位分隔符显示的字符.
// 有效数值是长度(包括小数点)为17位之内的数值,最后一位会四舍五入.
func (kn *LkkNumber) NumberFormat(number float64, decimal uint8, point, thousand string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimal)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousand)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += point + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// FloatEqual 比较两个浮点数是否相等.decimal为小数精确位数,默认为 FLOAT_DECIMAL .
// 有效数值是长度(包括小数点)为17位之内的数值,最后一位会四舍五入.
func (kn *LkkNumber) FloatEqual(f1 float64, f2 float64, decimal ...uint8) (res bool) {
	var threshold float64
	var dec uint8
	if len(decimal) == 0 {
		dec = FLOAT_DECIMAL
	} else {
		dec = decimal[0]
	}

	//比较精度
	threshold = math.Pow10(-int(dec))
	var diff float64
	if f1 > f2 {
		diff = f1 - f2
	} else {
		diff = f2 - f1
	}

	//diff := math.Abs(f1 - f2)
	res = diff <= threshold

	return
}

// RandInt64 生成一个min~max范围内的随机int64整数.
func (kn *LkkNumber) RandInt64(min, max int64) int64 {
	if min > max {
		min, max = max, min
	} else if min == max {
		return min
	}

	//范围是否在边界内
	mMax := int64(math.MaxInt32)
	mMin := int64(math.MinInt32)
	inrang := (mMin <= min && max <= mMax) || (INT64_MIN <= min && max <= 0) || (0 <= min && max <= INT64_MAX)
	if !inrang {
		min, max = mMin, mMax
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

// RandInt 生成一个min~max范围内的随机int整数.
func (kn *LkkNumber) RandInt(min, max int) int {
	if min > max {
		min, max = max, min
	} else if min == max {
		return min
	}

	//范围是否在边界内
	mMax := int(math.MaxInt32)
	mMin := int(math.MinInt32)
	inrang := (mMin <= min && max <= mMax) || (INT_MIN <= min && max <= 0) || (0 <= min && max <= INT_MAX)
	if !inrang {
		min, max = mMin, mMax
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// Rand RandInt的别名.
func (kn *LkkNumber) Rand(min, max int) int {
	return kn.RandInt(min, max)
}

// RandFloat64 生成一个min~max范围内的随机float64浮点数.
func (kn *LkkNumber) RandFloat64(min, max float64) float64 {
	if min > max {
		min, max = max, min
	}

	//范围是否在边界内
	mMax := float64(math.MaxFloat32)
	mMin := -mMax
	inrang := (mMin <= min && max <= mMax) || (-math.MaxFloat64 <= min && max <= 0) || (0 <= min && max <= math.MaxFloat64)
	if !inrang {
		min, max = mMin, mMax
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Float64()

	res := min + num*(max-min)
	return res
}

// Round 对浮点数(的整数)进行四舍五入.
func (kn *LkkNumber) Round(value float64) float64 {
	return math.Floor(value + 0.5)
}
