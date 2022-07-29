package kgo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestConvert_Struct2Map(t *testing.T) {
	//结构体
	var p1 sPerson
	_ = gofakeit.Struct(&p1)
	mp1, _ := KConv.Struct2Map(p1, "json")
	mp2, _ := KConv.Struct2Map(p1, "")

	var ok bool

	_, ok = mp1["name"]
	assert.True(t, ok)

	_, ok = mp1["none"]
	assert.False(t, ok)

	_, ok = mp2["Age"]
	assert.True(t, ok)

	_, ok = mp2["none"]
	assert.True(t, ok)
}

func BenchmarkConvert_Struct2Map_UseTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	_ = gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "json")
	}
}

func BenchmarkConvert_Struct2Map_NoTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	_ = gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "")
	}
}

func TestConvert_Int2Str(t *testing.T) {
	var res string

	res = KConv.Int2Str(0)
	assert.NotEmpty(t, res)

	res = KConv.Int2Str(31.4)
	assert.Empty(t, res)

	res = KConv.Int2Str(PKCS_SEVEN)
	assert.Equal(t, "7", res)
}

func BenchmarkConvert_Int2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int2Str(123456789)
	}
}

func TestConvert_Float2Str(t *testing.T) {
	var res string

	//小数位为负数
	res = KConv.Float2Str(flPi1, -2)
	assert.Equal(t, 4, len(res))

	res = KConv.Float2Str(flPi2, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi3, 3)
	assert.Equal(t, 5, len(res))

	res = KConv.Float2Str(flPi4, 9)
	assert.Equal(t, 11, len(res))

	res = KConv.Float2Str(true, 9)
	assert.Empty(t, res)
}

func BenchmarkConvert_Float2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(flPi2, 3)
	}
}

func TestConvert_Bool2Str(t *testing.T) {
	var res string

	res = KConv.Bool2Str(true)
	assert.Equal(t, "true", res)

	res = KConv.Bool2Str(false)
	assert.Equal(t, "false", res)
}

func BenchmarkConvert_Bool2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Str(true)
	}
}

func TestConvert_Bool2Int(t *testing.T) {
	var res int

	res = KConv.Bool2Int(true)
	assert.Equal(t, 1, res)

	res = KConv.Bool2Int(false)
	assert.Equal(t, 0, res)
}

func BenchmarkConvert_Bool2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Int(true)
	}
}

func TestConvert_Str2Int(t *testing.T) {
	var res int

	res = KConv.Str2Int("123")
	assert.Equal(t, 123, res)

	res = KConv.Str2Int("TRUE")
	assert.Equal(t, 1, res)

	res = KConv.Str2Int("")
	assert.Equal(t, 0, res)

	res = KConv.Str2Int(strHello)
	assert.Equal(t, 0, res)

	res = KConv.Str2Int("123.456")
	assert.Equal(t, 123, res)

	res = KConv.Str2Int("123.678")
	assert.Equal(t, 123, res)
}

func BenchmarkConvert_Str2Int_Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("TRUE")
	}
}

func BenchmarkConvert_Str2Int_Float(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("1234.567")
	}
}

func BenchmarkConvert_Str2Int_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("1234567")
	}
}

func TestConvert_Str2Int8(t *testing.T) {
	var res int8

	res = KConv.Str2Int8("99")
	assert.Equal(t, int8(99), res)

	res = KConv.Str2Int8(nowNanoStr)
	assert.Equal(t, int8(127), res)
}

func BenchmarkConvert_Str2Int8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int8("99")
	}
}

func TestConvert_Str2Int16(t *testing.T) {
	var res int16

	res = KConv.Str2Int16("99")
	assert.Equal(t, int16(99), res)

	res = KConv.Str2Int16(nowNanoStr)
	assert.Equal(t, int16(32767), res)
}

func BenchmarkConvert_Str2Int16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int16("99")
	}
}

func TestConvert_Str2Int32(t *testing.T) {
	var res int32

	res = KConv.Str2Int32("99")
	assert.Equal(t, int32(99), res)

	res = KConv.Str2Int32(nowNanoStr)
	assert.Equal(t, int32(2147483647), res)
}

func BenchmarkConvert_Str2Int32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int32("99")
	}
}

func TestConvert_Str2Int64(t *testing.T) {
	var res int64

	res = KConv.Str2Int64("99")
	assert.Equal(t, int64(99), res)

	res = KConv.Str2Int64(nowNanoStr)
	assert.Greater(t, res, int64(2147483648))
}

func BenchmarkConvert_Str2Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int64("99")
	}
}

func TestConvert_Str2Uint(t *testing.T) {
	var res uint

	res = KConv.Str2Uint("TRUE")
	assert.Equal(t, uint(1), res)

	res = KConv.Str2Uint("")
	assert.Equal(t, uint(0), res)

	res = KConv.Str2Uint(strHello)
	assert.Equal(t, uint(0), res)

	res = KConv.Str2Uint("123.456")
	assert.Equal(t, uint(123), res)

	//不合法的
	res = KConv.Str2Uint(" 123.456")
	assert.Equal(t, uint(0), res)

	res = KConv.Str2Uint("123.678")
	assert.Equal(t, uint(123), res)

	res = KConv.Str2Uint("125")
	assert.Equal(t, uint(125), res)

	res = KConv.Str2Uint("-125")
	assert.Equal(t, uint(0), res)
}

func BenchmarkConvert_Str2Uint_Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("TRUE")
	}
}

func BenchmarkConvert_Str2Uint_Float(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("1234.567")
	}
}

func BenchmarkConvert_Str2Uint_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("1234567")
	}
}

func TestConvert_Str2Uint8(t *testing.T) {
	var res uint8

	res = KConv.Str2Uint8("99")
	assert.Equal(t, uint8(99), res)

	res = KConv.Str2Uint8(nowNanoStr)
	assert.Equal(t, uint8(255), res)
}

func BenchmarkConvert_Str2Uint8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint8("99")
	}
}

func TestConvert_Str2Uint16(t *testing.T) {
	var res uint16

	res = KConv.Str2Uint16("99")
	assert.Equal(t, uint16(99), res)

	res = KConv.Str2Uint16(nowNanoStr)
	assert.Equal(t, uint16(65535), res)
}

func BenchmarkConvert_Str2Uint16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint16("99")
	}
}

func TestConvert_Str2Uint32(t *testing.T) {
	var res uint32

	res = KConv.Str2Uint32("99")
	assert.Equal(t, uint32(99), res)

	res = KConv.Str2Uint32(nowNanoStr)
	assert.Equal(t, uint32(4294967295), res)
}

func BenchmarkConvert_Str2Uint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint32("99")
	}
}

func TestConvert_Str2Uint64(t *testing.T) {
	var res uint64

	res = KConv.Str2Uint64("99")
	assert.Equal(t, uint64(99), res)

	res = KConv.Str2Uint64(nowNanoStr)
	assert.Greater(t, res, uint64(4294967295))
}

func BenchmarkConvert_Str2Uint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint64("99")
	}
}

func TestConvert_Str2Float32(t *testing.T) {
	var res float32

	res = KConv.Str2Float32("true")
	assert.Equal(t, float32(1), res)

	res = KConv.Str2Float32("")
	assert.Equal(t, float32(0), res)

	res = KConv.Str2Float32("123.556")
	assert.Equal(t, float32(123.556), res)
}

func BenchmarkConvert_Str2Float32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float32("123.556")
	}
}

func TestConvert_Str2Float64(t *testing.T) {
	var res float64

	res = KConv.Str2Float64("true")
	assert.Equal(t, float64(1), res)

	res = KConv.Str2Float64("")
	assert.Equal(t, float64(0), res)

	res = KConv.Str2Float64("123.556")
	assert.Equal(t, float64(123.556), res)
}

func BenchmarkConvert_Str2Float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float64("123.556")
	}
}

func TestConvert_Str2Bool(t *testing.T) {
	var res bool

	//true
	res = KConv.Str2Bool("1")
	assert.True(t, res)

	res = KConv.Str2Bool("t")
	assert.True(t, res)

	res = KConv.Str2Bool("T")
	assert.True(t, res)

	res = KConv.Str2Bool("TRUE")
	assert.True(t, res)

	res = KConv.Str2Bool("true")
	assert.True(t, res)

	res = KConv.Str2Bool("True")
	assert.True(t, res)

	//false
	res = KConv.Str2Bool("0")
	assert.False(t, res)

	res = KConv.Str2Bool("f")
	assert.False(t, res)

	res = KConv.Str2Bool("F")
	assert.False(t, res)

	res = KConv.Str2Bool("FALSE")
	assert.False(t, res)

	res = KConv.Str2Bool("false")
	assert.False(t, res)

	res = KConv.Str2Bool("False")
	assert.False(t, res)

	//other
	res = KConv.Str2Bool("2")
	assert.False(t, res)

	res = KConv.Str2Bool(strHello)
	assert.False(t, res)
}

func BenchmarkConvert_Str2Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Bool(strHello)
	}
}

func TestConvert_Str2Bytes(t *testing.T) {
	var res []byte

	res = KConv.Str2Bytes("")
	assert.Empty(t, res)

	res = KConv.Str2Bytes(strHello)
	assert.Equal(t, len(strHello), len(res))
}

func BenchmarkConvert_Str2Bytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Bytes(strHello)
	}
}

func TestConvert_Bytes2Str(t *testing.T) {
	var res string

	res = KConv.Bytes2Str([]byte{})
	assert.Equal(t, "", res)

	res = KConv.Bytes2Str(bytsHello)
	assert.NotEmpty(t, res)
}

func BenchmarkConvert_Bytes2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bytes2Str(bytsHello)
	}
}

func TestConvert_Str2BytesUnsafe(t *testing.T) {
	var res []byte

	res = KConv.Str2BytesUnsafe("")
	assert.Empty(t, res)

	res = KConv.Str2BytesUnsafe(strHello)
	assert.Equal(t, len(strHello), len(res))
}

func BenchmarkConvert_Str2BytesUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2BytesUnsafe(strHello)
	}
}

func TestConvert_Bytes2StrUnsafe(t *testing.T) {
	var res string

	res = KConv.Bytes2StrUnsafe([]byte{})
	assert.Equal(t, "", res)

	res = KConv.Bytes2StrUnsafe(bytsHello)
	assert.NotEmpty(t, res)
}

func BenchmarkConvert_Bytes2StrUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bytes2StrUnsafe(bytsHello)
	}
}

func TestConvert_Dec2Bin(t *testing.T) {
	var res string

	res = KConv.Dec2Bin(8)
	assert.Equal(t, "1000", res)

	res = KConv.Dec2Bin(16)
	assert.Equal(t, "10000", res)

	res = KConv.Dec2Bin(16)
}

func BenchmarkConvert_Dec2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Bin(16)
	}
}

func TestConvert_Bin2Dec(t *testing.T) {
	var res int64
	var err error

	res, _ = KConv.Bin2Dec("1000")
	assert.Equal(t, int64(8), res)

	res, _ = KConv.Bin2Dec("10000")
	assert.Equal(t, int64(16), res)

	//不合法
	_, err = KConv.Bin2Dec(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConvert_Bin2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2Dec("10000")
	}
}

func TestConvert_Hex2Bin(t *testing.T) {
	var res string
	var dec int64
	var err error

	res, err = KConv.Hex2Bin("123abff")
	assert.Nil(t, err)

	res, err = KConv.Hex2Bin(hexAstronomicalUnit)
	dec, _ = KConv.Bin2Dec(res)
	assert.Equal(t, intAstronomicalUnit, dec)

	//不合法
	_, err = KConv.Hex2Bin(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConvert_Hex2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Bin(hexAstronomicalUnit)
	}
}

func TestConvert_Bin2Hex(t *testing.T) {
	var res string
	var err error

	res, err = KConv.Bin2Hex(binAstronomicalUnit)
	assert.Equal(t, hexAstronomicalUnit, res)

	//不合法
	_, err = KConv.Bin2Hex(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConvert_Bin2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2Hex(binAstronomicalUnit)
	}
}

func TestConvert_Dec2Hex(t *testing.T) {
	var res string

	res = KConv.Dec2Hex(intAstronomicalUnit)
	assert.Equal(t, hexAstronomicalUnit, res)

	res = KConv.Dec2Hex(0)
	assert.Equal(t, "0", res)
}

func BenchmarkConvert_Dec2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Hex(intAstronomicalUnit)
	}
}

func TestConvert_Hex2Dec(t *testing.T) {
	var res int64
	var err error

	res, err = KConv.Hex2Dec(hexAstronomicalUnit)
	assert.Equal(t, intAstronomicalUnit, res)

	//不合法
	_, err = KConv.Hex2Dec(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConvert_Hex2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Dec(hexAstronomicalUnit)
	}
}

func TestConvert_Dec2Oct(t *testing.T) {
	var res string

	res = KConv.Dec2Oct(intAstronomicalUnit)
	assert.Equal(t, otcAstronomicalUnit, res)

	res = KConv.Dec2Oct(0)
	assert.Equal(t, "0", res)
}

func BenchmarkConvert_Dec2Oct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Oct(intAstronomicalUnit)
	}
}

func TestConvert_Oct2Dec(t *testing.T) {
	var res int64
	var err error

	res, err = KConv.Oct2Dec(otcAstronomicalUnit)
	assert.Equal(t, intAstronomicalUnit, res)

	//不合法
	_, err = KConv.Oct2Dec(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConvert_Oct2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Oct2Dec(otcAstronomicalUnit)
	}
}

func TestConvert_BaseConvert(t *testing.T) {
	var res string
	var err error

	res, err = KConv.BaseConvert(toStr(intAstronomicalUnit), 10, 16)
	assert.Equal(t, hexAstronomicalUnit, res)

	//不合法
	_, err = KConv.BaseConvert(strHello, 10, 8)
	assert.NotNil(t, err)
}

func BenchmarkConvert_BaseConvert(b *testing.B) {
	b.ResetTimer()
	s := toStr(intAstronomicalUnit)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.BaseConvert(s, 10, 16)
	}
}

func TestConvert_Ip2Long(t *testing.T) {
	var res uint32

	res = KConv.Ip2Long(localIp)
	assert.Equal(t, localIpInt, res)

	res = KConv.Ip2Long(lanIp)
	assert.Equal(t, lanIpInt, res)

	res = KConv.Ip2Long("")
	assert.Equal(t, uint32(0), res)
}

func BenchmarkConvert_Ip2Long(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Ip2Long(localIp)
	}
}

func TestConvert_Long2Ip(t *testing.T) {
	var res string

	res = KConv.Long2Ip(localIpInt)
	assert.Equal(t, localIp, res)

	res = KConv.Long2Ip(lanIpInt)
	assert.Equal(t, lanIp, res)

	res = KConv.Long2Ip(0)
	assert.Equal(t, noneIp, res)
}

func BenchmarkConvert_Long2Ip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Long2Ip(localIpInt)
	}
}

func TestConvert_ToStr(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected string
	}{
		{nil, ""},
		{true, "true"},
		{strHello, strHello},
		{intSpeedLight, strSpeedLight},
		{localIpInt, "2130706433"},
		{floSpeedLight, "2.99792458"},
		{flPi2, "3.141592456"},
		{fnCb1, "<nil>"},
		{fnPtr1, ""},
		{bytsHello, strHello},
		{bytColon, ":"},
		{int64(INT64_MAX), "9223372036854775807"},
		{uint64(UINT64_MAX), "18446744073709551615"},
		{float32(math.Pi), "3.1415927"},
		{float64(math.Pi), "3.141592653589793"},
		{strMpEmp, "{}"},
		{strMp2, `{"2":"cc","a":"0","b":"2","c":"4","g":"4","h":""}`},
	}

	for _, test := range tests {
		actual := KConv.ToStr(test.input)

		if reflect.DeepEqual(test.input, floSpeedLight) { //32位浮点会损失精度
			str := longestSameString("2.99792458", actual)
			assert.Less(t, 5, len(str))
		} else {
			assert.Equal(t, test.expected, actual)
		}
	}
}

func TestConvert_ToBool(t *testing.T) {
	//并行测试
	t.Parallel()

	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{int(-1), false},
		{int8(0), false},
		{int16(1), true},
		{int32(2), true},
		{int64(3), true},
		{uint(0), false},
		{uint8(0), false},
		{uint16(0), false},
		{uint32(0), false},
		{uint64(0), false},
		{float32(0), false},
		{float64(0), false},
		{[]byte{}, false},
		{bytsHello, true},
		{"1", true},
		{"2.1", false},
		{"TRUE", true},
		{false, false},
		{fnCb1, false},
		{nil, false},
		{personS1, true},
	}

	for _, test := range tests {
		actual := KConv.ToBool(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_ToBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToBool(intSpeedLight)
	}
}

func TestConvert_ToInt(t *testing.T) {
	//并行测试
	t.Parallel()

	var tests = []struct {
		input    interface{}
		expected int
	}{
		{int(-1), -1},
		{int8(0), 0},
		{int16(1), 1},
		{int32(2), 2},
		{int64(3), 3},
		{uint(0), 0},
		{uint8(0), 0},
		{uint16(0), 0},
		{uint32(0), 0},
		{uint64(0), 0},
		{float32(1.234), 1},
		{float64(4.5678), 4},
		{[]byte{}, 0},
		{"1", 1},
		{"2.1", 2},
		{"TRUE", 1},
		{true, 1},
		{false, 0},
		{fnCb1, 0},
		{nil, 0},
		{personS1, 1},
		{crowd, 5},
	}
	for _, test := range tests {
		actual := KConv.ToInt(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_ToInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToInt(intSpeedLight)
	}
}

func TestConvert_ToFloat(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected float64
	}{
		{int(-1), -1.0},
		{int8(0), 0.0},
		{int16(1), 1.0},
		{int32(2), 2.0},
		{int64(3), 3.0},
		{uint(0), 0.0},
		{uint8(0), 0.0},
		{uint16(0), 0.0},
		{uint32(0), 0.0},
		{uint64(0), 0.0},
		{float32(0), 0.0},
		{float64(0), 0.0},
		{[]byte{}, 0.0},
		{"1", 1.0},
		{"2.1", 2.1},
		{"TRUE", 1.0},
		{true, 1.0},
		{false, 0},
		{fnCb1, 0},
		{nil, 0},
		{personS1, 1.0},
		{crowd, 5.0},
	}
	for _, test := range tests {
		actual := KConv.ToFloat(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_ToFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToFloat(intSpeedLight)
	}
}

func TestConvert_Float64ToByte(t *testing.T) {
	res := KConv.Float64ToByte(flPi2)
	assert.NotEmpty(t, res)
}

func BenchmarkConvert_Float64ToByte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Float64ToByte(flPi2)
	}
}

func TestConvert_Byte2Float64(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	res := KConv.Byte2Float64(bytPi5)
	assert.Equal(t, flPi2, res)

	//不合法
	KConv.Byte2Float64([]byte{0, 1, 2, 3})
}

func BenchmarkConvert_Byte2Float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Byte2Float64(bytPi5)
	}
}

func TestConvert_Int64ToByte(t *testing.T) {
	res := KConv.Int64ToByte(intAstronomicalUnit)
	assert.NotEmpty(t, res)
}

func BenchmarkConvert_Int64ToByte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int64ToByte(intAstronomicalUnit)
	}
}

func TestConvert_Byte2Int64(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	res := KConv.Byte2Int64(bytAstronomicalUnit)
	assert.Equal(t, intAstronomicalUnit, res)

	//不合法
	KConv.Byte2Int64([]byte{0, 1, 2, 3})
}

func BenchmarkConvert_Byte2Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Byte2Int64(bytAstronomicalUnit)
	}
}

func TestConvert_Byte2Hex(t *testing.T) {
	var res string

	res = KConv.Byte2Hex(bytsHello)
	assert.Equal(t, strHelloHex, res)

	res = KConv.Byte2Hex(bytEmpty)
	assert.Empty(t, res)
}

func BenchmarkConvert_Byte2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Byte2Hex(bytsHello)
	}
}

func TestConvert_Byte2Hexs(t *testing.T) {
	var res []byte

	res = KConv.Byte2Hexs(bytsHello)
	assert.Equal(t, strHelloHex, string(res))

	res = KConv.Byte2Hexs(bytEmpty)
	assert.Empty(t, res)
}

func BenchmarkConvert_Byte2Hexs(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Byte2Hexs(bytsHello)
	}
}

func TestConvert_Hex2Byte(t *testing.T) {
	var res []byte

	res = KConv.Hex2Byte(strHelloHex)
	assert.Equal(t, strHello, string(res))

	res = KConv.Hex2Byte("")
	assert.Empty(t, res)
}

func BenchmarkConvert_Hex2Byte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Hex2Byte(strHelloHex)
	}
}

func TestConvert_Hexs2Byte(t *testing.T) {
	var res []byte

	bs := KConv.Byte2Hexs(bytsHello)
	res = KConv.Hexs2Byte(bs)
	assert.Equal(t, strHello, string(res))

	res = KConv.Hexs2Byte([]byte{})
	assert.Empty(t, res)

	//非16进制
	res = KConv.Hexs2Byte([]byte(utf8Hello))
	assert.Empty(t, res)
}

func BenchmarkConvert_Hexs2Byte(b *testing.B) {
	b.ResetTimer()
	bs := KConv.Byte2Hexs(bytsHello)
	for i := 0; i < b.N; i++ {
		KConv.Hexs2Byte(bs)
	}
}

func TestConvert_Runes2Bytes(t *testing.T) {
	res := KConv.Runes2Bytes(runesHello)
	assert.Equal(t, bytsHello, res)
}

func BenchmarkConvert_Runes2Bytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Runes2Bytes(runesHello)
	}
}

func TestConvert_IsString(t *testing.T) {
	var res bool

	res = KConv.IsString(intSpeedLight)
	assert.False(t, res)

	res = KConv.IsString(strHello)
	assert.True(t, res)
}

func BenchmarkConvert_IsString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsString(strHello)
	}
}

func TestConvert_IsBinary(t *testing.T) {
	var res bool
	var cont []byte

	cont, _ = KFile.ReadFile(imgPng)
	res = KConv.IsBinary(string(cont))
	assert.True(t, res)

	cont, _ = KFile.ReadFile(fileDante)
	res = KConv.IsBinary(string(cont))
	assert.False(t, res)
}

func BenchmarkConvert_IsBinary(b *testing.B) {
	b.ResetTimer()
	cont, _ := KFile.ReadFile(imgPng)
	for i := 0; i < b.N; i++ {
		KConv.IsBinary(string(cont))
	}
}

func TestConvert_IsNumeric(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{intSpeedLight, true},
		{flPi1, true},
		{strSpeedLight, true},
		{strHello, false},
		{crowd, false},
		{"", false},
	}
	for _, test := range tests {
		actual := KConv.IsNumeric(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsNumeric(intSpeedLight)
	}
}

func TestConvert_IsInt(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{intSpeedLight, true},
		{strSpeedLight, true},
		{tesStr42, true},
		{flPi1, false},
		{strPi6, false},
		{strHello, false},
		{strHelloHex, false},
		{crowd, false},
		{tesStr17, false},
		{"", false},
	}
	for _, test := range tests {
		actual := KConv.IsInt(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsInt(intSpeedLight)
	}
}

func TestConvert_IsFloat(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{flPi1, true},
		{strPi6, true},
		{tesStr44, true},
		{tesStr45, true},
		{intSpeedLight, false},
		{strSpeedLight, false},
		{strHello, false},
		{crowd, false},
		{tesStr17, false},
		{tesStr21, false},
		{tesStr43, false},
		{"", false},
	}
	for _, test := range tests {
		actual := KConv.IsFloat(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsFloat(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsFloat(flPi1)
	}
}

func TestConvert_IsEmpty(t *testing.T) {
	var org sOrganization
	var itf interface{} = &strSlEmp
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{nil, true},
		{"", true},
		{strMpEmp, true},
		{false, true},
		{0, true},
		{uint(0), true},
		{0.0, true},
		{org, true},
		{itf, false},
	}
	for _, test := range tests {
		actual := KConv.IsEmpty(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsEmpty(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsEmpty("")
	}
}

func TestConvert_IsNil(t *testing.T) {
	var org sOrganization
	var itf interface{} = &strSlEmp
	var s []int
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{nil, true},
		{"", false},
		{strMpEmp, false},
		{false, false},
		{0, false},
		{uint(0), false},
		{0.0, false},
		{org, false},
		{itf, false},
		{s, true},
	}
	for _, test := range tests {
		actual := KConv.IsNil(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsNil(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsNil(nil)
	}
}

func TestConvert_IsBool(t *testing.T) {
	var res bool

	res = KConv.IsBool(false)
	assert.True(t, res)

	res = KConv.IsBool("true")
	assert.False(t, res)
}

func BenchmarkConvert_IsBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsBool(false)
	}
}

func TestConvert_IsHex(t *testing.T) {
	var str1 = KConv.Dec2Hex(intAstronomicalUnit)
	var str2 = "0x" + str1

	var tests = []struct {
		input    string
		expected bool
	}{
		{"", false},
		{str1, true},
		{str2, true},
		{strHelloHex, true},
		{strHello, false},
	}
	for _, test := range tests {
		actual := KConv.IsHex(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsHex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsHex(strHelloHex)
	}
}

func TestConvert_IsByte(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{"", false},
		{runesHello, false},
		{bytsHello, true},
	}
	for _, test := range tests {
		actual := KConv.IsByte(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsByte(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsByte(bytsHello)
	}
}

func TestConvert_IsStruct(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{strHello, false},
		{runesHello, false},
		{cmplNum1, false},
		{colorMp, false},
		{personS1, true},
		{&personS1, true},
		{orgS1, true},
		{&orgS1, false},
	}
	for _, test := range tests {
		actual := KConv.IsStruct(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsStruct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsStruct(personS1)
	}
}

func TestConvert_IsInterface(t *testing.T) {
	var tests = []struct {
		input    interface{}
		expected bool
	}{
		{strHello, false},
		{personS1, false},
		{itfObj, true},
	}
	for _, test := range tests {
		actual := KConv.IsInterface(test.input)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsInterface(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsInterface(itfObj)
	}
}

func TestConvert_IsPort(t *testing.T) {
	var tests = []struct {
		param    interface{}
		expected bool
	}{
		{"hello", false},
		{"1", true},
		{0, false},
		{100, true},
		{"65535", true},
		{"0", false},
		{"65536", false},
		{"65538.9", false},
	}

	for _, test := range tests {
		actual := KConv.IsPort(test.param)
		assert.Equal(t, actual, test.expected)
	}
}

func BenchmarkConvert_IsPort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.IsPort(80)
	}
}

func TestConvert_ToInterfaces(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[arrayValues]`arr type must be")
	}()

	res := KConv.ToInterfaces(ssSingle)
	assert.NotNil(t, res)
	assert.Equal(t, len(res), len(ssSingle))

	KConv.ToInterfaces(strHello)
}

func BenchmarkConvert_ToInterfaces(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToInterfaces(ssSingle)
	}
}
