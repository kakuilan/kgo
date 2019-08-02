package kgo

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// NumberFormat 以千位分隔符方式格式化一个数字;decimals为要保留的小数位数,decPoint为小数点显示的字符,thousandsSep为千位分隔符显示的字符
func (kn *LkkNumber) NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
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
