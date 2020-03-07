package kgo

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// IsLetters 字符串是否全(英文)字母组成.
func (ks *LkkString) IsLetters(str string) bool {
	for _, r := range str {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return str != ""
}

// IsEmpty 字符串是否为空(包括空格).
func (ks *LkkString) IsEmpty(str string) bool {
	if str == "" || len(ks.Trim(str)) == 0 {
		return true
	}

	return false
}

// IsUpper 字符串是否全部大写.
func (ks *LkkString) IsUpper(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}

// IsLower 字符串是否全部小写.
func (ks *LkkString) IsLower(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return str != ""
}

// HasLetter 字符串是否含有(英文)字母.
func (ks *LkkString) HasLetter(str string) bool {
	for _, r := range str {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return true
		}
	}
	return false
}

// IsUtf8 字符串是否UTF-8编码.
func (ks *LkkString) IsUtf8(str string) bool {
	return str != "" && utf8.ValidString(str)
}

// IsASCII 是否IsASCII字符串.
func (ks *LkkString) IsASCII(str string) bool {
	//return str != "" && RegAscii.MatchString(str)
	n := len(str)
	for i := 0; i < n; i++ {
		if str[i] > 127 {
			return false
		}
	}

	return str != ""
}

// IsMultibyte 字符串是否含有多字节字符.
func (ks *LkkString) IsMultibyte(str string) bool {
	return str != "" && RegMultiByte.MatchString(str)
}

// HasFullWidth 是否含有全角字符.
func (ks *LkkString) HasFullWidth(str string) bool {
	return str != "" && RegFullWidth.MatchString(str)
}

// HasHalfWidth 是否含有半角字符.
func (ks *LkkString) HasHalfWidth(str string) bool {
	return str != "" && RegHalfWidth.MatchString(str)
}

// IsEnglish 字符串是否纯英文.letterCase是否检查大小写,枚举值(CASE_NONE,CASE_LOWER,CASE_UPPER).
func (ks *LkkString) IsEnglish(str string, letterCase LkkCaseSwitch) bool {
	switch letterCase {
	case CASE_NONE:
		return ks.IsLetters(str)
	case CASE_LOWER:
		return str != "" && RegAlphaLower.MatchString(str)
	case CASE_UPPER:
		return str != "" && RegAlphaUpper.MatchString(str)
	default:
		return ks.IsLetters(str)
	}
}

// HasEnglish 是否含有英文字符,HasLetter的别名.
func (ks *LkkString) HasEnglish(str string) bool {
	return ks.HasLetter(str)
}

// HasChinese 字符串是否含有中文.
func (ks *LkkString) HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}

// IsChinese 字符串是否全部中文.
func (ks *LkkString) IsChinese(str string) bool {
	return str != "" && RegChineseAll.MatchString(str)
}

// IsChineseName 字符串是否中文名称.
func (ks *LkkString) IsChineseName(str string) bool {
	return str != "" && RegChineseName.MatchString(str)
}

// IsWord 是否词语(不以下划线开头的中文、英文、数字、下划线).
func (ks *LkkString) IsWord(str string) bool {
	return str != "" && RegWord.MatchString(str)
}

// HasSpecialChar 字符串是否含有特殊字符.
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

// IsJSON 字符串是否合法的json格式.
func (ks *LkkString) IsJSON(str string) bool {
	length := len(str)
	if length == 0 {
		return false
	} else if (str[0] != '{' || str[length-1] != '}') && (str[0] != '[' || str[length-1] != ']') {
		return false
	}

	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// IsIP 检查字符串是否IP地址.
func (ks *LkkString) IsIP(str string) bool {
	return str != "" && net.ParseIP(str) != nil
}

// IsIPv4 检查字符串是否IPv4地址.
func (ks *LkkString) IsIPv4(str string) bool {
	ipAddr := net.ParseIP(str)
	// 不是合法的IP地址
	if ipAddr == nil {
		return false
	}

	return ipAddr.To4() != nil && strings.ContainsRune(str, '.')
}

// IsIPv6 检查字符串是否IPv6地址.
func (ks *LkkString) IsIPv6(str string) bool {
	ipAddr := net.ParseIP(str)
	return ipAddr != nil && strings.ContainsRune(str, ':')
}

// IsPort 字符串或数字是否端口号.
func (ks *LkkString) IsPort(val interface{}) bool {
	if KConv.IsInt(val) {
		port := KConv.ToInt(val)
		if port > 0 && port < 65536 {
			return true
		}
	}

	return false
}

// IsDNSName 是否DNS名称.
func (ks *LkkString) IsDNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return !ks.IsIP(str) && RegDNSname.MatchString(str)
}

// IsDialString 是否网络拨号字符串(形如127.0.0.1:80),用于net.Dial()检查.
func (ks *LkkString) IsDialString(str string) bool {
	h, p, err := net.SplitHostPort(str)
	if err == nil && h != "" && p != "" && (ks.IsDNSName(h) || ks.IsIP(h)) && ks.IsPort(p) {
		return true
	}

	return false
}

// IsMACAddr 是否MAC物理网卡地址.
func (ks *LkkString) IsMACAddr(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// IsHost 字符串是否主机名(IP或DNS名称).
func (ks *LkkString) IsHost(str string) bool {
	return ks.IsIP(str) || ks.IsDNSName(str)
}

// IsEmail 检查字符串是否邮箱.参数validateTrue,是否验证邮箱主机的真实性.
func (ks *LkkString) IsEmail(email string, validateHost bool) (bool, error) {
	//长度检查
	length := len(email)
	at := strings.LastIndexByte(email, '@')
	if (length < 6 || length > 254) || (at <= 0 || at > length-3) {
		return false, fmt.Errorf("invalid email length")
	}

	//验证邮箱格式
	chkFormat := RegEmail.MatchString(email)
	if !chkFormat {
		return false, fmt.Errorf("invalid email format")
	}

	//验证主机
	if validateHost {
		host := email[at+1:]
		if _, err := net.LookupMX(host); err != nil {
			//因无法确定mx主机的smtp端口,所以去掉Hello/Mail/Rcpt检查邮箱是否存在
			//仅检查主机是否有效
			//TODO 仅对国内几家大的邮件厂家进行检查
			if _, err := net.LookupIP(host); err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

// IsMobilecn 检查字符串是否中国大陆手机号.
func (ks *LkkString) IsMobilecn(str string) bool {
	return str != "" && RegMobilecn.MatchString(str)
}

// IsTel 是否固定电话或400/800电话.
func (ks *LkkString) IsTel(str string) bool {
	return str != "" && RegTelephone.MatchString(str)
}

// IsPhone 是否电话号码(手机或固话).
func (ks *LkkString) IsPhone(str string) bool {
	return str != "" && RegPhone.MatchString(str)
}

// IsCreditNo 检查是否(15或18位)身份证号码,并返回经校验的号码.
func (ks *LkkString) IsCreditNo(str string) (bool, string) {
	chk := str != "" && RegCreditno.MatchString(str)
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
	return str != "" && RegAlphaNumeric.MatchString(str)
}

// IsHexcolor 检查是否十六进制颜色,并返回带"#"的修正值.
func (ks *LkkString) IsHexcolor(str string) (bool, string) {
	chk := str != "" && RegHexcolor.MatchString(str)
	if chk && !strings.ContainsRune(str, '#') {
		str = "#" + strings.ToUpper(str)
	}
	return chk, str
}

// IsRGBcolor 检查字符串是否RGB颜色格式.
func (ks *LkkString) IsRGBcolor(str string) bool {
	return str != "" && RegRgbcolor.MatchString(str)
}

// IsBlank 是否空(空白)字符.
func (ks *LkkString) IsBlank(str string) bool {
	// Check length
	if len(str) > 0 {
		// Iterate string
		for i := range str {
			// Check about char different from whitespace
			// 227为全角空格
			if str[i] > 32 && str[i] != 227 {
				return false
			}
		}
	}
	return true
}

// IsWhitespaces 是否全部空白字符,不包括空字符串.
func (ks *LkkString) IsWhitespaces(str string) bool {
	return str != "" && RegWhitespaceAll.MatchString(str)
}

// HasWhitespace 是否带有空白字符.
func (ks *LkkString) HasWhitespace(str string) bool {
	return str != "" && RegWhitespaceHas.MatchString(str)
}

// IsBase64 是否base64字符串.
func (ks *LkkString) IsBase64(str string) bool {
	return str != "" && RegBase64.MatchString(str)
}

// IsBase64Image 是否base64编码的图片.
func (ks *LkkString) IsBase64Image(str string) bool {
	if str == "" || !strings.ContainsRune(str, ',') {
		return false
	}

	dataURI := strings.Split(str, ",")
	return RegBase64Image.MatchString(dataURI[0]) && RegBase64.MatchString(dataURI[1])
}

// IsRsaPublicKey 检查字符串是否RSA的公钥,keylen为密钥长度.
func (ks *LkkString) IsRsaPublicKey(str string, keylen int) bool {
	bb := bytes.NewBufferString(str)
	pemBytes, _ := ioutil.ReadAll(bb)

	// 获取公钥
	block, _ := pem.Decode(pemBytes)
	if block != nil && block.Type != "PUBLIC KEY" {
		return false
	}
	var der []byte
	var err error

	if block != nil {
		der = block.Bytes
	} else {
		der, err = base64.StdEncoding.DecodeString(str)
		if err != nil {
			return false
		}
	}

	key, err := x509.ParsePKIXPublicKey(der)
	if err != nil {
		return false
	}
	pubkey, ok := key.(*rsa.PublicKey)
	if !ok {
		return false
	}
	bitlen := len(pubkey.N.Bytes()) * 8
	return bitlen == int(keylen)
}

// IsUrl 检查字符串是否URL.
func (ks *LkkString) IsUrl(str string) bool {
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
func (ks *LkkString) IsUrlExists(str string) bool {
	if !ks.IsUrl(str) {
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

// IsMd5 是否md5值.
func (ks *LkkString) IsMd5(str string) bool {
	return str != "" && RegMd5.MatchString(str)
}

// IsSha1 是否Sha1值.
func (ks *LkkString) IsSha1(str string) bool {
	return str != "" && RegSha1.MatchString(str)
}

// IsSha256 是否Sha256值.
func (ks *LkkString) IsSha256(str string) bool {
	return str != "" && RegSha256.MatchString(str)
}

// IsSha512 是否Sha512值.
func (ks *LkkString) IsSha512(str string) bool {
	return str != "" && RegSha512.MatchString(str)
}

// StartsWith 字符串str是否以sub开头.
func (ks *LkkString) StartsWith(str, sub string, ignoreCase bool) bool {
	if str != "" && sub != "" {
		i := ks.Index(str, sub, ignoreCase)
		return i == 0
	}

	return false
}

// EndsWith 字符串str是否以sub结尾.
func (ks *LkkString) EndsWith(str, sub string, ignoreCase bool) bool {
	if str != "" && sub != "" {
		i := ks.LastIndex(str, sub, ignoreCase)
		return i != -1 && (len(str)-len(sub)) == i
	}

	return false
}

// IsArrayOrSlice 检查变量是否数组或切片.
// chkType检查类型,枚举值有(1仅数组,2仅切片,3数组或切片);结果为-1表示非,>=0表示是
func (ka *LkkArray) IsArrayOrSlice(val interface{}, chkType uint8) int {
	return isArrayOrSlice(val, chkType)
}

// IsMap 检查变量是否字典.
func (ka *LkkArray) IsMap(val interface{}) bool {
	return isMap(val)
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

	chk := RegDatetime.MatchString(str)
	if !chk {
		return false, 0
	}

	leng := len(str)
	if leng < 19 {
		reference := "1970-01-01 00:00:00"
		str = str + reference[leng:19]
	}

	tim, err := KTime.Str2Timestamp(str)
	if err != nil {
		return false, 0
	}

	return true, tim
}

// IsNan 是否为“非数值”.
func (kn *LkkNumber) IsNan(val float64) bool {
	return math.IsNaN(val)
}

// IsString 变量是否字符串.
func (kc *LkkConvert) IsString(val interface{}) bool {
	return kc.Gettype(val) == "string"
}

// IsBinary 字符串是否二进制.
func (kc *LkkConvert) IsBinary(s string) bool {
	for _, b := range s {
		if 0 == b {
			return true
		}
	}
	return false
}

// IsNumeric 变量是否数值(不包含复数和科学计数法).
func (kc *LkkConvert) IsNumeric(val interface{}) bool {
	return isNumeric(val)
}

// IsInt 变量是否整型数值.
func (kc *LkkConvert) IsInt(val interface{}) bool {
	return isInt(val)
}

// IsFloat 变量是否浮点数值.
func (kc *LkkConvert) IsFloat(val interface{}) bool {
	return isFloat(val)
}

// IsEmpty 检查变量是否为空.
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

// IsNil 检查变量是否空值.
func (kc *LkkConvert) IsNil(val interface{}) bool {
	if val == nil {
		return true
	}

	rv := reflect.ValueOf(val)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Slice, reflect.Interface:
		if rv.IsNil() {
			return true
		}
	}
	return false
}

// IsBool 是否布尔值.
func (kc *LkkConvert) IsBool(val interface{}) bool {
	return val == true || val == false
}

// IsHex 是否十六进制字符串.
func (kc *LkkConvert) IsHex(str string) bool {
	_, err := kc.Hex2Dec(str)
	return err == nil
}

// IsByte 变量是否字节切片.
func (kc *LkkConvert) IsByte(val interface{}) bool {
	return kc.Gettype(val) == "[]uint8"
}

// IsStruct 变量是否结构体.
func (kc *LkkConvert) IsStruct(val interface{}) bool {
	r := reflectPtr(reflect.ValueOf(val))
	return r.Kind() == reflect.Struct
}

// IsInterface 变量是否接口.
func (kc *LkkConvert) IsInterface(val interface{}) bool {
	r := reflectPtr(reflect.ValueOf(val))
	return r.Kind() == reflect.Invalid
}
