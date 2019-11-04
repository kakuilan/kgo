package kgo

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

// NumberFormat 以千位分隔符方式格式化一个数字;decimals为要保留的小数位数,decPoint为小数点显示的字符,thousandsSep为千位分隔符显示的字符
func (kn *LkkNumber) NumberFormat(number float64, decimals uint8, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
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
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

// Range 根据范围创建数组，包含指定的元素
func (kn *LkkNumber) Range(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// Abs 取绝对值
func (kn *LkkNumber) Abs(number float64) float64 {
	return math.Abs(number)
}

// Rand 产生一个随机整数,范围: [0, 2147483647]
func (kn *LkkNumber) Rand(min, max int) int {
	if min > max {
		panic("[Rand]: min cannot be greater than max")
	}
	// PHP: getrandmax()
	if int31 := 1<<31 - 1; max > int31 {
		panic("[Rand]: max can not be greater than " + strconv.Itoa(int31))
	}
	if min == max {
		return min
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max+1-min) + min
}

// Round 对浮点数进行四舍五入
func (kn *LkkNumber) Round(value float64) float64 {
	return math.Floor(value + 0.5)
}

// Floor 向下取整
func (kn *LkkNumber) Floor(value float64) float64 {
	return math.Floor(value)
}

// Ceil 向上取整
func (kn *LkkNumber) Ceil(value float64) float64 {
	return math.Ceil(value)
}

// Pi 得到圆周率值
func (kn *LkkNumber) Pi() float64 {
	return math.Pi
}

// Max 取出最大值
func (kn *LkkNumber) Max(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("[Max]: the nums length is less than 2")
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		max = math.Max(max, nums[i])
	}
	return max
}

// Min 取出最小值
func (kn *LkkNumber) Min(nums ...float64) float64 {
	if len(nums) < 2 {
		panic("[Min]: the nums length is less than 2")
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		min = math.Min(min, nums[i])
	}
	return min
}

// Exp 计算 e 的指数
func (kn *LkkNumber) Exp(x float64) float64 {
	return math.Exp(x)
}

// Expm1 返回 exp(number) - 1，甚至当 number 的值接近零也能计算出准确结果
func (kn *LkkNumber) Expm1(x float64) float64 {
	return math.Exp(x) - 1
}

// Pow 指数表达式
func (kn *LkkNumber) Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

// ByteFormat 格式化文件比特大小.size为文件大小,decimals为要保留的小数位数.
func (kn *LkkNumber) ByteFormat(size float64, decimals uint8) string {
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

	return fmt.Sprintf("%."+strconv.Itoa(int(decimals))+"f%s", j, arr[pos])
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
func (kn *LkkNumber) NumSign(value float64) float64 {
	if value > 0 {
		return 1
	} else if value < 0 {
		return -1
	} else {
		return 0
	}
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
	return kn.IsWhole(value) && kn.IsPositive(value)
}

// InRangeInt 数值是否在2个整数范围内,将自动转换为整数再比较.
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
