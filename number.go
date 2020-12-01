package kgo

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// NumberFormat 以千位分隔符方式格式化一个数字.
// decimal为要保留的小数位数,point为小数点显示的字符,thousand为千位分隔符显示的字符.
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

// AbsFloat 浮点型取绝对值.
func (kn *LkkNumber) AbsFloat(number float64) float64 {
	return math.Abs(number)
}

// AbsInt 整型取绝对值.
func (kn *LkkNumber) AbsInt(number int64) int64 {
	r := number >> 63
	return (number ^ r) - r
}

// FloatEqual 比较两个浮点数是否相等.decimal为小数精确位数.
func (kn *LkkNumber) FloatEqual(f1 float64, f2 float64, decimal ...int) bool {
	var threshold float64
	var dec int
	if len(decimal) == 0 {
		dec = FLOAT_DECIMAL
	} else {
		dec = decimal[0]
	}

	//比较精度
	threshold = math.Pow10(-dec)

	return math.Abs(f1-f2) <= threshold
}

// RandInt 产生一个随机int整数.
func (kn *LkkNumber) RandInt(min, max int) int {
	if min > max {
		panic("[RandInt]: min cannot be greater than max")
	}

	if min == max {
		return min
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

// RandInt64 生产一个随机int64整数.
func (kn *LkkNumber) RandInt64(min, max int64) int64 {
	if min > max {
		panic("[RandInt64]: min cannot be greater than max")
	}

	if min == max {
		return min
	}

	//范围是否在边界内
	mMax := int64(math.MaxInt32)
	mMin := int64(math.MinInt32)
	inrang := (mMin <= min && max <= mMax) || (INT64_MIN <= min && max <= 0) || (0 <= min && max <= INT64_MAX)

	if !inrang {
		panic("[RandInt64]: min and max exceed capacity,the result should be overflows int64.")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

// RandFloat64 生产一个随机float64整数.
func (kn *LkkNumber) RandFloat64(min, max float64) float64 {
	if min > max {
		panic("[RandFloat64]: min cannot be greater than max")
	}

	//范围是否在边界内
	mMax := float64(math.MaxFloat32)
	mMin := -mMax
	inrang := (mMin <= min && max <= mMax) || (-math.MaxFloat64 <= min && max <= 0) || (0 <= min && max <= math.MaxFloat64)
	if !inrang {
		panic("[RandFloat64]: min and max exceed capacity,the result should be overflows float64.")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Float64()

	res := min + num*(max-min)
	return res
}

// Rand RandInt的别名.
func (kn *LkkNumber) Rand(min, max int) int {
	return kn.RandInt(min, max)
}

// Round 对浮点数(的整数)进行四舍五入.
func (kn *LkkNumber) Round(value float64) float64 {
	return math.Floor(value + 0.5)
}

// RoundPlus 对指定的小数位进行四舍五入.
// precision为小数位数.
func (kn *LkkNumber) RoundPlus(value float64, precision int8) float64 {
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

// Pi 得到圆周率值.
func (kn *LkkNumber) Pi() float64 {
	return math.Pi
}

// MaxInt 整数序列求最大值.
func (kn *LkkNumber) MaxInt(nums ...int) (res int) {
	if len(nums) < 1 {
		panic("[MaxInt]: the nums length is less than 1")
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
		panic("[MaxFloat64]: the nums length is less than 1")
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
		panic("[Max]: the nums length is less than 1")
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
		panic("[MinInt]: the nums length is less than 1")
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
		panic("[MinFloat64]: the nums length is less than 1")
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
		panic("[Min]: the nums length is less than 1")
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

// Expm1 返回 exp(number) - 1，甚至当 number 的值接近零也能计算出准确结果.
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
	var arr = []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB", "UnKnown"}
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

// IsNatural 数值是否为自然数.
func (kn *LkkNumber) IsNatural(value float64) bool {
	return kn.IsWhole(value) && kn.IsNonNegative(value)
}

// InRangeInt 数值是否在2个整数范围内.
func (kn *LkkNumber) InRangeInt(value, left, right int) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// InRangeFloat32 数值是否在2个32位浮点数范围内.
func (kn *LkkNumber) InRangeFloat32(value, left, right float32) bool {
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

// InRange 数值是否在某个范围内,将自动转换类型再比较.
func (kn *LkkNumber) InRange(value interface{}, left interface{}, right interface{}) bool {
	reflectValue := reflect.TypeOf(value).Kind()
	reflectLeft := reflect.TypeOf(left).Kind()
	reflectRight := reflect.TypeOf(right).Kind()

	if reflectValue == reflect.Int && reflectLeft == reflect.Int && reflectRight == reflect.Int {
		return kn.InRangeInt(value.(int), left.(int), right.(int))
	} else if reflectValue == reflect.Float32 && reflectLeft == reflect.Float32 && reflectRight == reflect.Float32 {
		return kn.InRangeFloat32(value.(float32), left.(float32), right.(float32))
	} else if reflectValue == reflect.Float64 && reflectLeft == reflect.Float64 && reflectRight == reflect.Float64 {
		return kn.InRangeFloat64(value.(float64), left.(float64), right.(float64))
	} else if KConv.IsInt(value) && KConv.IsInt(left) && KConv.IsInt(right) {
		return kn.InRangeInt(KConv.ToInt(value), KConv.ToInt(left), KConv.ToInt(right))
	} else if KConv.IsNumeric(value) && KConv.IsNumeric(left) && KConv.IsNumeric(right) {
		return kn.InRangeFloat64(KConv.ToFloat(value), KConv.ToFloat(left), KConv.ToFloat(right))
	}

	return false
}

// SumInt 整数求和.
func (kn *LkkNumber) SumInt(nums ...int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return sum
}

// SumFloat64 浮点数求和.
func (kn *LkkNumber) SumFloat64(nums ...float64) float64 {
	var sum float64
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Sum 对任意类型序列中的数值类型求和,忽略非数值的.
func (kn *LkkNumber) Sum(nums ...interface{}) (res float64) {
	var err error
	var val float64
	for _, v := range nums {
		val, err = numeric2Float(v)
		if err == nil {
			res += val
		}
	}

	return
}

// AverageInt 对整数序列求平均值.
func (kn *LkkNumber) AverageInt(nums ...int) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res = float64(nums[0])
	} else {
		total := kn.SumInt(nums...)
		res = float64(total) / float64(length)
	}

	return
}

// AverageFloat64 对浮点数序列求平均值.
func (kn *LkkNumber) AverageFloat64(nums ...float64) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res = nums[0]
	} else {
		total := kn.SumFloat64(nums...)
		res = total / float64(length)
	}

	return
}

// Average 对任意类型序列中的数值类型求平均值,忽略非数值的.
func (kn *LkkNumber) Average(nums ...interface{}) (res float64) {
	length := len(nums)
	if length == 0 {
		return
	} else if length == 1 {
		res, _ = numeric2Float(nums[0])
	} else {
		var count int
		var err error
		var val, total float64
		for _, v := range nums {
			val, err = numeric2Float(v)
			if err == nil {
				count++
				total += val
			}
		}

		res = total / float64(count)
	}

	return
}

// Percent 返回百分比(val/total).
func (kn *LkkNumber) Percent(val, total interface{}) float64 {
	t := KConv.ToFloat(total)
	if t == 0 {
		return float64(0)
	}

	v := KConv.ToFloat(val)

	return (v / t) * 100
}

// GeoDistance 获取地理距离/米.
// 参数分别为两点的经度和纬度.lat:-90~90,lng:-180~180.
func (kn *LkkNumber) GeoDistance(lng1, lat1, lng2, lat2 float64) float64 {
	//地球半径
	radius := 6371000.0
	rad := math.Pi / 180.0

	lng1 = lng1 * rad
	lat1 = lat1 * rad
	lng2 = lng2 * rad
	lat2 = lat2 * rad
	theta := lng2 - lng1

	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

// IsNan 是否为“非数值”.
func (kn *LkkNumber) IsNan(val interface{}) bool {
	if isFloat(val) {
		return math.IsNaN(KConv.ToFloat(val))
	}

	return !isNumeric(val)
}

// IsNaturalRange 是否连续的自然数数组/切片,如[0,1,2,3...],其中不能有间断.
// strict为是否严格检查元素的顺序.
func (kn *LkkNumber) IsNaturalRange(arr []int, strict bool) (res bool) {
	n := len(arr)
	if n == 0 {
		return
	}

	orig := kn.Range(0, n-1)
	ctyp := COMPARE_ONLY_VALUE

	if strict {
		ctyp = COMPARE_BOTH_KEYVALUE
	}

	diff := KArr.ArrayDiff(orig, arr, ctyp)

	res = len(diff) == 0
	return
}
