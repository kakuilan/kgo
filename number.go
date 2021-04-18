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

// RoundPlus 对指定的小数位进行四舍五入.
// precision为小数位数.
func (kn *LkkNumber) RoundPlus(value float64, precision uint8) float64 {
	shift := math.Pow(10, float64(precision))
	return kn.Round(value*shift) / shift
}

// Floor 向下取整.
func (kn *LkkNumber) Floor(value float64) float64 {
	return math.Floor(value)
}

// Ceil 向上取整.
func (kn *LkkNumber) Ceil(value float64) float64 {
	return math.Ceil(value)
}

// MaxInt 整数序列求最大值.
func (kn *LkkNumber) MaxInt(nums ...int) (res int) {
	if len(nums) < 1 {
		panic("[MaxInt]` nums length is less than 1")
	}

	res = nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}

	return
}

// MaxFloat64 64位浮点数序列求最大值.
func (kn *LkkNumber) MaxFloat64(nums ...float64) (res float64) {
	if len(nums) < 1 {
		panic("[MaxFloat64]` nums length is less than 1")
	}

	res = nums[0]
	for _, v := range nums {
		res = math.Max(res, v)
	}

	return
}

// Max 取出任意类型中数值类型的最大值,无数值类型则为0.
func (kn *LkkNumber) Max(nums ...interface{}) (res float64) {
	if len(nums) < 1 {
		panic("[Max]` nums length is less than 1")
	}

	var err error
	var val float64
	res, _ = numeric2Float(nums[0])
	for _, v := range nums {
		val, err = numeric2Float(v)
		if err == nil {
			res = math.Max(res, val)
		}
	}

	return
}

// MinInt 整数序列求最小值.
func (kn *LkkNumber) MinInt(nums ...int) (res int) {
	if len(nums) < 1 {
		panic("[MinInt]` nums length is less than 1")
	}
	res = nums[0]
	for _, v := range nums {
		if v < res {
			res = v
		}
	}

	return
}

// MinFloat64 64位浮点数序列求最小值.
func (kn *LkkNumber) MinFloat64(nums ...float64) (res float64) {
	if len(nums) < 1 {
		panic("[MinFloat64]` nums length is less than 1")
	}
	res = nums[0]
	for _, v := range nums {
		res = math.Min(res, v)
	}

	return
}

// Min 取出任意类型中数值类型的最小值,无数值类型则为0.
func (kn *LkkNumber) Min(nums ...interface{}) (res float64) {
	if len(nums) < 1 {
		panic("[Min]` nums length is less than 1")
	}

	var err error
	var val float64
	res, _ = numeric2Float(nums[0])
	for _, v := range nums {
		val, err = numeric2Float(v)
		if err == nil {
			res = math.Min(res, val)
		}
	}

	return
}

// Exp 计算 e 的指数.
func (kn *LkkNumber) Exp(x float64) float64 {
	return math.Exp(x)
}

// Expm1 返回 exp(x) - 1.
func (kn *LkkNumber) Expm1(x float64) float64 {
	return math.Exp(x) - 1
}

// Pow 指数表达式,求x的y次方.
func (kn *LkkNumber) Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// Log 对数表达式,求以y为底x的对数.
func (kn *LkkNumber) Log(x, y float64) float64 {
	return math.Log(x) / math.Log(y)
}

// ByteFormat 格式化文件比特大小.
// size为文件大小,decimal为要保留的小数位数,delimiter为数字和单位间的分隔符.
func (kn *LkkNumber) ByteFormat(size float64, decimal uint8, delimiter string) string {
	var arr = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", Unknown}
	var pos int = 0
	var j float64 = float64(size)
	for {
		if size >= 1024 {
			size = size / 1024
			j = j / 1024
			pos++
		} else {
			break
		}
	}
	if pos >= len(arr) { // fixed out index bug
		pos = len(arr) - 1
	}

	return fmt.Sprintf("%."+strconv.Itoa(int(decimal))+"f%s%s", j, delimiter, arr[pos])
}

// IsOdd 变量是否奇数.
func (kn *LkkNumber) IsOdd(val int) bool {
	return val%2 != 0
}

// IsEven 变量是否偶数.
func (kn *LkkNumber) IsEven(val int) bool {
	return val%2 == 0
}

// NumSign 返回数值的符号.值>0为1,<0为-1,其他为0.
func (kn *LkkNumber) NumSign(value float64) (res int8) {
	if value > 0 {
		res = 1
	} else if value < 0 {
		res = -1
	} else {
		res = 0
	}

	return
}

// IsNegative 数值是否为负数.
func (kn *LkkNumber) IsNegative(value float64) bool {
	return value < 0
}

// IsPositive 数值是否为正数.
func (kn *LkkNumber) IsPositive(value float64) bool {
	return value > 0
}

// IsNonNegative 数值是否为非负数.
func (kn *LkkNumber) IsNonNegative(value float64) bool {
	return value >= 0
}

// IsNonPositive 数值是否为非正数.
func (kn *LkkNumber) IsNonPositive(value float64) bool {
	return value <= 0
}

// IsWhole 数值是否为整数.
func (kn *LkkNumber) IsWhole(value float64) bool {
	return math.Remainder(value, 1) == 0
}

// IsNatural 数值是否为自然数(包括0).
func (kn *LkkNumber) IsNatural(value float64) bool {
	return kn.IsNonNegative(value) && kn.IsWhole(value)
}

// InRangeInt 数值是否在2个整数范围内.
func (kn *LkkNumber) InRangeInt(value, left, right int) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRangeFloat64 数值是否在2个64位浮点数范围内.
func (kn *LkkNumber) InRangeFloat64(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}
