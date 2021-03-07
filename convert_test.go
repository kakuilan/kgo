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
	gofakeit.Struct(&p1)
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
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "json")
	}
}

func BenchmarkConvert_Struct2Map_NoTag(b *testing.B) {
	b.ResetTimer()
	var p1 sPerson
	gofakeit.Struct(&p1)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Struct2Map(p1, "")
	}
}

func TestConver_Int2Str(t *testing.T) {
	var res string

	res = KConv.Int2Str(0)
	assert.NotEmpty(t, res)

	res = KConv.Int2Str(31.4)
	assert.Empty(t, res)

	res = KConv.Int2Str(PKCS_SEVEN)
	assert.Equal(t, "7", res)
}

func BenchmarkConver_Int2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Int2Str(123456789)
	}
}

func TestConver_Float2Str(t *testing.T) {
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

func BenchmarkConver_Float2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Float2Str(flPi2, 3)
	}
}

func TestConver_Bool2Str(t *testing.T) {
	var res string

	res = KConv.Bool2Str(true)
	assert.Equal(t, "true", res)

	res = KConv.Bool2Str(false)
	assert.Equal(t, "false", res)
}

func BenchmarkConver_Bool2Str(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Str(true)
	}
}

func TestConver_Bool2Int(t *testing.T) {
	var res int

	res = KConv.Bool2Int(true)
	assert.Equal(t, 1, res)

	res = KConv.Bool2Int(false)
	assert.Equal(t, 0, res)
}

func BenchmarkConver_Bool2Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bool2Int(true)
	}
}

func TestConver_Str2Int(t *testing.T) {
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

func BenchmarkConver_Str2Int_Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("TRUE")
	}
}

func BenchmarkConver_Str2Int_Float(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("1234.567")
	}
}

func BenchmarkConver_Str2Int_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int("1234567")
	}
}

func TestConver_Str2Int8(t *testing.T) {
	var res int8

	res = KConv.Str2Int8("99")
	assert.Equal(t, int8(99), res)

	res = KConv.Str2Int8(nowNanoStr)
	assert.Equal(t, int8(127), res)
}

func BenchmarkConver_Str2Int8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int8("99")
	}
}

func TestConver_Str2Int16(t *testing.T) {
	var res int16

	res = KConv.Str2Int16("99")
	assert.Equal(t, int16(99), res)

	res = KConv.Str2Int16(nowNanoStr)
	assert.Equal(t, int16(32767), res)
}

func BenchmarkConver_Str2Int16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int16("99")
	}
}

func TestConver_Str2Int32(t *testing.T) {
	var res int32

	res = KConv.Str2Int32("99")
	assert.Equal(t, int32(99), res)

	res = KConv.Str2Int32(nowNanoStr)
	assert.Equal(t, int32(2147483647), res)
}

func BenchmarkConver_Str2Int32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int32("99")
	}
}

func TestConver_Str2Int64(t *testing.T) {
	var res int64

	res = KConv.Str2Int64("99")
	assert.Equal(t, int64(99), res)

	res = KConv.Str2Int64(nowNanoStr)
	assert.Greater(t, res, int64(2147483648))
}

func BenchmarkConver_Str2Int64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Int64("99")
	}
}

func TestConver_Str2Uint(t *testing.T) {
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

func BenchmarkConver_Str2Uint_Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("TRUE")
	}
}

func BenchmarkConver_Str2Uint_Float(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("1234.567")
	}
}

func BenchmarkConver_Str2Uint_Int(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint("1234567")
	}
}

func TestConver_Str2Uint8(t *testing.T) {
	var res uint8

	res = KConv.Str2Uint8("99")
	assert.Equal(t, uint8(99), res)

	res = KConv.Str2Uint8(nowNanoStr)
	assert.Equal(t, uint8(255), res)
}

func BenchmarkConver_Str2Uint8(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint8("99")
	}
}

func TestConver_Str2Uint16(t *testing.T) {
	var res uint16

	res = KConv.Str2Uint16("99")
	assert.Equal(t, uint16(99), res)

	res = KConv.Str2Uint16(nowNanoStr)
	assert.Equal(t, uint16(65535), res)
}

func BenchmarkConver_Str2Uint16(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint16("99")
	}
}

func TestConver_Str2Uint32(t *testing.T) {
	var res uint32

	res = KConv.Str2Uint32("99")
	assert.Equal(t, uint32(99), res)

	res = KConv.Str2Uint32(nowNanoStr)
	assert.Equal(t, uint32(4294967295), res)
}

func BenchmarkConver_Str2Uint32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint32("99")
	}
}

func TestConver_Str2Uint64(t *testing.T) {
	var res uint64

	res = KConv.Str2Uint64("99")
	assert.Equal(t, uint64(99), res)

	res = KConv.Str2Uint64(nowNanoStr)
	assert.Greater(t, res, uint64(4294967295))
}

func BenchmarkConver_Str2Uint64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Uint64("99")
	}
}

func TestConver_Str2Float32(t *testing.T) {
	var res float32

	res = KConv.Str2Float32("true")
	assert.Equal(t, float32(1), res)

	res = KConv.Str2Float32("")
	assert.Equal(t, float32(0), res)

	res = KConv.Str2Float32("123.556")
	assert.Equal(t, float32(123.556), res)
}

func BenchmarkConver_Str2Float32(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float32("123.556")
	}
}

func TestConver_Str2Float64(t *testing.T) {
	var res float64

	res = KConv.Str2Float64("true")
	assert.Equal(t, float64(1), res)

	res = KConv.Str2Float64("")
	assert.Equal(t, float64(0), res)

	res = KConv.Str2Float64("123.556")
	assert.Equal(t, float64(123.556), res)
}

func BenchmarkConver_Str2Float64(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Float64("123.556")
	}
}

func TestConver_Str2Bool(t *testing.T) {
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

func BenchmarkConver_Str2Bool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Bool(strHello)
	}
}

func TestConver_Str2Bytes(t *testing.T) {
	var res []byte

	res = KConv.Str2Bytes("")
	assert.Empty(t, res)

	res = KConv.Str2Bytes(strHello)
	assert.Equal(t, len(strHello), len(res))
}

func BenchmarkConver_Str2Bytes(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2Bytes(strHello)
	}
}

func TestConver_Bytes2Str(t *testing.T) {
	var res string

	res = KConv.Bytes2Str([]byte{})
	assert.Equal(t, "", res)

	res = KConv.Bytes2Str([]byte(strHello))
	assert.NotEmpty(t, res)
}

func BenchmarkConver_Bytes2Str(b *testing.B) {
	var bs = []byte(strHello)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bytes2Str(bs)
	}
}

func TestConver_Str2BytesUnsafe(t *testing.T) {
	var res []byte

	res = KConv.Str2BytesUnsafe("")
	assert.Empty(t, res)

	res = KConv.Str2BytesUnsafe(strHello)
	assert.Equal(t, len(strHello), len(res))
}

func BenchmarkConver_Str2BytesUnsafe(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Str2BytesUnsafe(strHello)
	}
}

func TestConver_Bytes2StrUnsafe(t *testing.T) {
	var res string

	res = KConv.Bytes2StrUnsafe([]byte{})
	assert.Equal(t, "", res)

	res = KConv.Bytes2StrUnsafe([]byte(strHello))
	assert.NotEmpty(t, res)
}

func BenchmarkConver_Bytes2StrUnsafe(b *testing.B) {
	var bs = []byte(strHello)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Bytes2StrUnsafe(bs)
	}
}

func TestConver_Dec2Bin(t *testing.T) {
	var res string

	res = KConv.Dec2Bin(8)
	assert.Equal(t, "1000", res)

	res = KConv.Dec2Bin(16)
	assert.Equal(t, "10000", res)

	res = KConv.Dec2Bin(16)
}

func BenchmarkConver_Dec2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Bin(16)
	}
}

func TestConver_Bin2Dec(t *testing.T) {
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

func BenchmarkConver_Bin2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2Dec("10000")
	}
}

func TestConver_Hex2Bin(t *testing.T) {
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

func BenchmarkConver_Hex2Bin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Bin(hexAstronomicalUnit)
	}
}

func TestConver_Bin2Hex(t *testing.T) {
	var res string
	var err error

	res, err = KConv.Bin2Hex(binAstronomicalUnit)
	assert.Equal(t, hexAstronomicalUnit, res)

	//不合法
	_, err = KConv.Bin2Hex(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConver_Bin2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Bin2Hex(binAstronomicalUnit)
	}
}

func TestConver_Dec2Hex(t *testing.T) {
	var res string

	res = KConv.Dec2Hex(intAstronomicalUnit)
	assert.Equal(t, hexAstronomicalUnit, res)

	res = KConv.Dec2Hex(0)
	assert.Equal(t, "0", res)
}

func BenchmarkConver_Dec2Hex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Hex(intAstronomicalUnit)
	}
}

func TestConver_Hex2Dec(t *testing.T) {
	var res int64
	var err error

	res, err = KConv.Hex2Dec(hexAstronomicalUnit)
	assert.Equal(t, intAstronomicalUnit, res)

	//不合法
	_, err = KConv.Hex2Dec(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConver_Hex2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Hex2Dec(hexAstronomicalUnit)
	}
}

func TestConver_Dec2Oct(t *testing.T) {
	var res string

	res = KConv.Dec2Oct(intAstronomicalUnit)
	assert.Equal(t, otcAstronomicalUnit, res)

	res = KConv.Dec2Oct(0)
	assert.Equal(t, "0", res)
}

func BenchmarkConver_Dec2Oct(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Dec2Oct(intAstronomicalUnit)
	}
}

func TestConver_Oct2Dec(t *testing.T) {
	var res int64
	var err error

	res, err = KConv.Oct2Dec(otcAstronomicalUnit)
	assert.Equal(t, intAstronomicalUnit, res)

	//不合法
	_, err = KConv.Oct2Dec(strHello)
	assert.NotNil(t, err)
}

func BenchmarkConver_Oct2Dec(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KConv.Oct2Dec(otcAstronomicalUnit)
	}
}

func TestConver_BaseConvert(t *testing.T) {
	var res string
	var err error

	res, err = KConv.BaseConvert(toStr(intAstronomicalUnit), 10, 16)
	assert.Equal(t, hexAstronomicalUnit, res)

	//不合法
	_, err = KConv.BaseConvert(strHello, 10, 8)
	assert.NotNil(t, err)
}

func BenchmarkConver_BaseConvert(b *testing.B) {
	b.ResetTimer()
	s := toStr(intAstronomicalUnit)
	for i := 0; i < b.N; i++ {
		_, _ = KConv.BaseConvert(s, 10, 16)
	}
}

func TestConver_Ip2Long(t *testing.T) {
	var res uint32

	res = KConv.Ip2Long(localIp)
	assert.Equal(t, localIpInt, res)

	res = KConv.Ip2Long(lanIp)
	assert.Equal(t, lanIpInt, res)

	res = KConv.Ip2Long("")
	assert.Equal(t, uint32(0), res)
}

func BenchmarkConver_Ip2Long(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Ip2Long(localIp)
	}
}

func TestConver_Long2Ip(t *testing.T) {
	var res string

	res = KConv.Long2Ip(localIpInt)
	assert.Equal(t, localIp, res)

	res = KConv.Long2Ip(lanIpInt)
	assert.Equal(t, lanIp, res)

	res = KConv.Long2Ip(0)
	assert.Equal(t, noneIp, res)
}

func BenchmarkConver_Long2Ip(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.Long2Ip(localIpInt)
	}
}

func TestConver_ToStr(t *testing.T) {
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
		{bytSlcHello, strHello},
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

func TestConver_ToBool(t *testing.T) {
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
		{bytSlcHello, true},
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

func BenchmarkConver_ToBool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToBool(intSpeedLight)
	}
}

func TestConver_ToInt(t *testing.T) {
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

func BenchmarkConver_ToInt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KConv.ToInt(intSpeedLight)
	}
}
