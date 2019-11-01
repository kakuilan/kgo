package kgo

import (
	"encoding/json"
	"fmt"
	"math"
	"net"
	"net/http"
	"net/smtp"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// IsLetters 字符串是否全(英文)字母组成
func (ks *LkkString) IsLetters(str string) bool {
	if str == "" {
		return false
	}
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// HasLetter 字符串是否含有(英文)字母
func (ks *LkkString) HasLetter(str string) bool {
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

// IsUtf8 字符串是否UTF-8编码
func (ks *LkkString) IsUtf8(str string) bool {
	return utf8.ValidString(str)
}

// IsEnglish 字符串是否纯英文.letterCase是否检查大小写,枚举值(CASE_NONE,CASE_LOWER,CASE_UPPER).
func (ks *LkkString) IsEnglish(str string, letterCase LkkCaseSwitch) bool {
	switch letterCase {
	case CASE_NONE:
		return ks.IsLetters(str)
	case CASE_LOWER:
		return str != "" && regexp.MustCompile(PATTERN_ALPHA_LOWER).MatchString(str)
	case CASE_UPPER:
		return str != "" && regexp.MustCompile(PATTERN_ALPHA_UPPER).MatchString(str)
	default:
		return ks.IsLetters(str)
	}
}

// HasEnglish 是否含有英文字符,HasLetter的别名.
func (ks *LkkString) HasEnglish(str string) bool {
	return ks.HasLetter(str)
}

// HasChinese 字符串是否含有中文
func (ks *LkkString) HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

// IsChinese 字符串是否全部中文
func (ks *LkkString) IsChinese(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_ALL_CHINESE).MatchString(str)
}

// HasSpecialChar 字符串是否含有特殊字符
func (ks *LkkString) HasSpecialChar(str string) (res bool) {
	if str == "" {
		return
	}

	for _, r := range str {
		// IsPunct 判断 r 是否为一个标点字符 (类别 P)
		// IsSymbol 判断 r 是否为一个符号字符
		// IsMark 判断 r 是否为一个 mark 字符 (类别 M)
		if unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsMark(r) {
			res = true
			return
		}
	}

	return
}

// IsJSON 字符串是否合法的json格式
func (ks *LkkString) IsJSON(str string) bool {
	if str == "" {
		return false
	} else if str[0] != '{' || str[len(str)-1] != '}' {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsIPv4 检查字符串是否IPv4地址
func (ks *LkkString) IsIPv4(str string) bool {
	ipAddr := net.ParseIP(str)
	// 不是合法的IP地址
	if ipAddr == nil {
		return false
	}

	return ipAddr.To4() != nil && strings.ContainsRune(str, '.')
}

// IsIPv6 检查字符串是否IPv6地址
func (ks *LkkString) IsIPv6(str string) bool {
	ipAddr := net.ParseIP(str)
	return ipAddr != nil && strings.ContainsRune(str, ':')
}

// IsEmail 检查字符串是否邮箱.参数validateTrue,是否验证邮箱的真实性.
func (ks *LkkString) IsEmail(email string, validateTrue bool) (bool, error) {
	//验证邮箱格式
	chkFormat := regexp.MustCompile(PATTERN_EMAIL).MatchString(email)
	if !chkFormat {
		return false, fmt.Errorf("invalid email format")
	}

	//验证真实性
	if validateTrue {
		i := strings.LastIndexByte(email, '@')
		host := email[i+1:]

		// MX records
		mx, err := net.LookupMX(host)
		if err != nil {
			return false, err
		}

		server := fmt.Sprintf("%s:%d", mx[0].Host, 25)
		conn, err := net.DialTimeout("tcp", server, CHECK_CONNECT_TIMEOUT)
		if err != nil {
			return false, err
		}

		t := time.AfterFunc(CHECK_CONNECT_TIMEOUT, func() { conn.Close() })
		defer t.Stop()

		client, err := smtp.NewClient(conn, host)
		if err != nil {
			return false, err
		}
		defer client.Close()

		err = client.Hello("checkmail.me")
		if err != nil {
			return false, err
		}
		err = client.Mail("kakuilan@gmail.com")
		if err != nil {
			return false, err
		}
		err = client.Rcpt(email)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// IsMobile 检查字符串是否手机号
func (ks *LkkString) IsMobile(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_MOBILE).MatchString(str)
}

// IsTel 是否固定电话或400/800电话.
func (ks *LkkString) IsTel(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_TELEPHONE).MatchString(str)
}

// IsPhone 是否电话号码(手机或固话).
func (ks *LkkString) IsPhone(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_PHONE).MatchString(str)
}

// IsCreditNo 检查是否(15或18位)身份证号码,并返回经校验的号码.
func (ks *LkkString) IsCreditNo(str string) (bool, string) {
	chk := str != "" && regexp.MustCompile(PATTERN_CREDIT_NO).MatchString(str)
	if !chk {
		return false, ""
	}

	// 检查省份代码
	if _, chk = CreditArea[str[0:2]]; !chk {
		return false, ""
	}

	// 将15位身份证升级到18位
	leng := len(str)
	if leng == 15 {
		// 先转为17位,如果身份证顺序码是996 997 998 999,这些是为百岁以上老人的特殊编码
		if chk, _ = ks.Dstrpos(str[12:], []string{"996", "997", "998", "999"}, false); chk {
			str = str[0:6] + "18" + str[6:]
		} else {
			str = str[0:6] + "19" + str[6:]
		}

		// 再加上校验码
		code := append([]byte{}, creditChecksum(str))
		str += string(code)
	}

	// 检查生日
	birthday := str[6:10] + "-" + str[10:12] + "-" + str[12:14]
	chk, tim := KTime.IsDate2time(birthday)
	now := KTime.Time()
	if !chk {
		return false, ""
	} else if tim >= now {
		return false, ""
	}

	// 18位身份证需要验证最后一位校验位
	if leng == 18 {
		str = strings.ToUpper(str)
		if str[17] != creditChecksum(str) {
			return false, ""
		}
	}

	return true, str
}

// IsAlphaNumeric 是否字母或数字.
func (ks *LkkString) IsAlphaNumeric(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_ALPHA_NUMERIC).MatchString(str)
}

// IsHexcolor 检查是否十六进制颜色,并返回带"#"的修正值.
func (ks *LkkString) IsHexcolor(str string) (bool, string) {
	chk := str != "" && regexp.MustCompile(PATTERN_HEXCOLOR).MatchString(str)
	if chk && !strings.ContainsRune(str, '#') {
		str = "#" + strings.ToUpper(str)
	}
	return chk, str
}

// IsRGBcolor 检查字符串是否RGB颜色格式.
func (ks *LkkString) IsRGBcolor(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_RGBCOLOR).MatchString(str)
}

// IsWhitespaces 是否全部空白字符.
func (ks *LkkString) IsWhitespaces(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_WHITESPACE_ALL).MatchString(str)
}

// HasWhitespace 是否带有空白字符.
func (ks *LkkString) HasWhitespace(str string) bool {
	return str != "" && regexp.MustCompile(PATTERN_WHITESPACE_HAS).MatchString(str)
}

// IsUrl 检查字符串是否URL.
func (ku *LkkUrl) IsUrl(str string) bool {
	if str == "" || len(str) <= 3 || utf8.RuneCountInString(str) >= 2083 || strings.HasPrefix(str, ".") {
		return false
	}

	res, err := url.ParseRequestURI(str)
	if err != nil {
		return false //Couldn't even parse the url
	}
	if len(res.Scheme) == 0 {
		return false //No Scheme found
	}

	return true
}

// IsUrlExists 检查URL是否存在.
func (ku *LkkUrl) IsUrlExists(str string) bool {
	if !ku.IsUrl(str) {
		return false
	}

	client := &http.Client{}
	resp, err := client.Head(str)
	if err != nil {
		return false
	} else if resp.StatusCode == 404 {
		return false
	}

	return true
}

// IsArrayOrSlice 检查变量是否数组或切片;chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (ka *LkkArray) IsArrayOrSlice(data interface{}, chkType uint8) int {
	return isArrayOrSlice(data, chkType)
}

// IsMap 检查变量是否字典
func (ka *LkkArray) IsMap(data interface{}) bool {
	return isMap(data)
}

// IsDate2time 检查字符串是否日期格式,并转换为时间戳.注意,时间戳可能为负数(小于1970年时).
// 匹配如:
//	0000
//	0000-00
//	0000/00
//	0000-00-00
//	0000/00/00
//	0000-00-00 00
//	0000/00/00 00
//	0000-00-00 00:00
//	0000/00/00 00:00
//	0000-00-00 00:00:00
//	0000/00/00 00:00:00
// 等日期格式.
func (kt *LkkTime) IsDate2time(str string) (bool, int64) {
	if str == "" {
		return false, 0
	} else if strings.ContainsRune(str, '/') {
		str = strings.Replace(str, "/", "-", -1)
	}

	chk := regexp.MustCompile(PATTERN_DATETIME).MatchString(str)
	if !chk {
		return false, 0
	}

	leng := len(str)
	if leng < 19 {
		reference := "1970-01-01 00:00:00"
		str = str + reference[leng:19]
	}

	tim, err := KTime.Strtotime(str)
	if err != nil {
		return false, 0
	}

	return true, tim
}

// IsNan 是否为“非数值”
func (kn *LkkNumber) IsNan(val float64) bool {
	return math.IsNaN(val)
}

// IsString 变量是否字符串
func (kc *LkkConvert) IsString(v interface{}) bool {
	return kc.Gettype(v) == "string"
}

// IsBinary 字符串是否二进制
func (kc *LkkConvert) IsBinary(s string) bool {
	for _, b := range s {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsNumeric 变量是否数值(不包含复数和科学计数法)
func (kc *LkkConvert) IsNumeric(val interface{}) bool {
	return isNumeric(val)
}

// IsInt 变量是否整型数值
func (kc *LkkConvert) IsInt(val interface{}) bool {
	return isInt(val)
}

// IsFloat 变量是否浮点数值
func (kc *LkkConvert) IsFloat(val interface{}) bool {
	return isFloat(val)
}

// IsEmpty 检查一个变量是否为空
func (kc *LkkConvert) IsEmpty(val interface{}) bool {
	if val == nil {
		return true
	}
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.String, reflect.Array:
		return v.Len() == 0
	case reflect.Map, reflect.Slice:
		return v.Len() == 0 || v.IsNil()
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}

	return reflect.DeepEqual(val, reflect.Zero(v.Type()).Interface())
}

// IsBool 是否布尔值
func (kc *LkkConvert) IsBool(v interface{}) bool {
	return v == true || v == false
}

// IsHex 是否十六进制字符串
func (kc *LkkConvert) IsHex(str string) bool {
	_, err := kc.Hex2dec(str)
	return err == nil
}

// IsByte 变量是否字节切片
func (kc *LkkConvert) IsByte(v interface{}) bool {
	return kc.Gettype(v) == "[]uint8"
}

// IsStruct 变量是否结构体
func (kc *LkkConvert) IsStruct(v interface{}) bool {
	r := reflectPtr(reflect.ValueOf(v))
	return r.Kind() == reflect.Struct
}

// IsInterface 变量是否接口
func (kc *LkkConvert) IsInterface(v interface{}) bool {
	r := reflectPtr(reflect.ValueOf(v))
	return r.Kind() == reflect.Invalid
}
