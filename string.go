package kgo

import (
	"bytes"
	"github.com/json-iterator/go"
	xhtml "golang.org/x/net/html"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
	"unicode/utf8"
)

// Md5 获取字节切片md5值.
// length指定结果长度,默认32.
func (ks *LkkString) Md5Byte(str []byte, length ...uint8) []byte {
	var l uint8 = 32
	if len(length) > 0 {
		l = length[0]
	}

	return md5Byte(str, l)
}

// Md5 获取字符串md5值.
// length指定结果长度,默认32.
func (ks *LkkString) Md5(str string, length ...uint8) string {
	var l uint8 = 32
	if len(length) > 0 {
		l = length[0]
	}

	return string(md5Byte([]byte(str), l))
}

// IsMd5 是否md5值.
func (ks *LkkString) IsMd5(str string) bool {
	return str != "" && RegMd5.MatchString(str)
}

// ShaXByte 计算字节切片的 shaX 散列值,x为1/256/512 .
func (ks *LkkString) ShaXByte(str []byte, x uint16) []byte {
	return shaXByte(str, x)
}

// ShaX 计算字符串的 shaX 散列值,x为1/256/512 .
func (ks *LkkString) ShaX(str string, x uint16) string {
	return string(shaXByte([]byte(str), x))
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

// Addslashes 使用反斜线引用字符串.
func (ks *LkkString) Addslashes(str string) string {
	var buf bytes.Buffer
	for _, char := range str {
		switch char {
		case '\'', '"', '\\':
			buf.WriteRune('\\')
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// Stripslashes 反引用一个引用字符串.
func (ks *LkkString) Stripslashes(str string) string {
	var buf bytes.Buffer
	l, skip := len(str), false
	for i, char := range str {
		if skip {
			skip = false
		} else if char == '\\' {
			if i+1 < l && str[i+1] == '\\' {
				skip = true
			}
			continue
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

// JsonEncode 对val变量进行 JSON 编码.
// 依赖库github.com/json-iterator/go.
func (ks *LkkString) JsonEncode(val interface{}) ([]byte, error) {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Marshal(val)
}

// JsonDecode 对 JSON 格式的str字符串进行解码,注意res使用指针.
// 依赖库github.com/json-iterator/go.
func (ks *LkkString) JsonDecode(str []byte, res interface{}) error {
	var jsons = jsoniter.ConfigCompatibleWithStandardLibrary
	return jsons.Unmarshal(str, res)
}

// Utf8ToGbk UTF-8转GBK编码.
func (ks *LkkString) Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := io.ReadAll(reader)
	return d, e
}

// GbkToUtf8 GBK转UTF-8编码.
func (ks *LkkString) GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := io.ReadAll(reader)
	return d, e
}

// IsUtf8 字符串是否UTF-8编码.
func (ks *LkkString) IsUtf8(s []byte) bool {
	return utf8.Valid(s)
}

// IsGbk 字符串是否GBK编码.
func (ks *LkkString) IsGbk(s []byte) (res bool) {
	length := len(s)
	var i, j int
	for i < length {
		j = i + 1
		//大于127的使用双字节编码,且落在gbk编码范围内的字符
		//GBK中每个汉字包含两个字节，第一个字节(首字节)的范围是0x81-0xFE(即129-254),第二个字节(尾字节)的范围是0x40-0xFE(即64-254)
		if s[i] > 0x7f && j < length {
			if s[i] >= 0x81 &&
				s[i] <= 0xfe &&
				s[j] >= 0x40 &&
				s[j] <= 0xfe {
				i += 2
				res = true
			} else {
				res = false
				break
			}
		} else {
			i++
		}
	}

	return
}

// DetectEncoding 匹配字符编码,TODO.
func (ks *LkkString) DetectEncoding() {
	//TODO 检查字符编码
}

// Img2Base64 将图片字节转换为base64字符串.ext为图片扩展名,默认jpg.
func (ks *LkkString) Img2Base64(content []byte, ext ...string) string {
	var imgType string = "jpg"
	if len(ext) > 0 {
		imgType = strings.ToLower(ext[0])
	}

	return img2Base64(content, imgType)
}

// Index 查找子串sub在字符串str中第一次出现的位置,不存在则返回-1;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) Index(str, sub string, ignoreCase bool) int {
	if str == "" || sub == "" {
		return -1
	}

	if ignoreCase {
		str = strings.ToLower(str)
		sub = strings.ToLower(sub)
	}

	return strings.Index(str, sub)
}

// LastIndex 查找子串sub在字符串str中最后一次出现的位置,不存在则返回-1;
// ignoreCase为是否忽略大小写.
func (ks *LkkString) LastIndex(str, sub string, ignoreCase bool) int {
	if str == "" || sub == "" {
		return -1
	}

	if ignoreCase {
		str = strings.ToLower(str)
		sub = strings.ToLower(sub)
	}

	return strings.LastIndex(str, sub)
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

// Trim 去除字符串首尾处的空白字符（或者其他字符）.
func (ks *LkkString) Trim(str string, characterMask ...string) string {
	mask := getTrimMask(characterMask)
	return strings.Trim(str, mask)
}

// Strpos 查找字符串首次出现的位置,找不到时返回-1.
// haystack在该字符串中进行查找,needle要查找的字符串;
// offset起始位置,为负数时时,搜索会从字符串结尾指定字符数开始.
func (ks *LkkString) Strpos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := strings.Index(haystack[offset:], needle)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Stripos  查找字符串首次出现的位置（不区分大小写）,找不到时返回-1.
// haystack在该字符串中进行查找,needle要查找的字符串;
// offset起始位置,为负数时时,搜索会从字符串结尾指定字符数开始.
func (ks *LkkString) Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	if offset < 0 {
		offset += length
	}
	pos := ks.Index(haystack[offset:], needle, true)
	if pos == -1 {
		return -1
	}
	return pos + offset
}

// Dstrpos 检查字符串str是否包含数组arr的元素之一,返回检查结果和匹配的字符串.
// chkCase为是否检查大小写.
func (ks *LkkString) Dstrpos(str string, arr []string, chkCase bool) (bool, string) {
	if len(str) == 0 || len(arr) == 0 {
		return false, ""
	}

	for _, v := range arr {
		if (chkCase && ks.Strpos(str, v, 0) != -1) || (!chkCase && ks.Stripos(str, v, 0) != -1) {
			return true, v
		}
	}

	return false, ""
}

// Nl2br 将换行符转换为br标签.
func (ks *LkkString) Nl2br(str string) string {
	return strings.Replace(str, "\n", "<br />", -1)
}

// Br2nl 将br标签转换为换行符.
func (ks *LkkString) Br2nl(str string) string {
	// <br> , <br /> , <br/>
	// <BR> , <BR /> , <BR/>

	l := len(str)
	buf := make([]byte, 0, l) //prealloca

	for i := 0; i < l; i++ {
		switch str[i] {
		case 60: //<
			if l >= i+3 {
				/*
					b = 98
					B = 66
					r = 82
					R = 114
					SPACE = 32
					/ = 47
					> = 62
				*/

				if l >= i+3 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 62) { // <br> || <BR>
					buf = append(buf, bytLinefeed...)
					i += 3
					continue
				}

				if l >= i+4 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 47 && str[i+4] == 62) { // <br/> || <BR/>
					buf = append(buf, bytLinefeed...)
					i += 4
					continue
				}

				if l >= i+5 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 32 && str[i+4] == 47 && str[i+5] == 62) { // <br /> || <BR />
					buf = append(buf, bytLinefeed...)
					i += 5
					continue
				}
			}
		default:
			buf = append(buf, str[i])
		}
	}

	return string(buf)
}

// RemoveSpace 移除字符串中的空白字符.
// all为true时移除全部空白,为false时只替换连续的空白字符为一个空格.
func (ks *LkkString) RemoveSpace(str string, all bool) string {
	if all && str != "" {
		return strings.Join(strings.Fields(str), "")
	} else if str != "" {
		//先将2个以上的连续空白符转为空格
		str = RegWhitespaceDuplicate.ReplaceAllString(str, " ")
		//再将[\t\n\f\r]等转为空格
		str = RegWhitespace.ReplaceAllString(str, " ")
	}

	return strings.TrimSpace(str)
}

// StripTags 过滤html标签.
func (ks *LkkString) StripTags(str string) string {
	return RegHtmlTag.ReplaceAllString(str, "")
}

// Html2Text 将html转换为纯文本.
func (ks *LkkString) Html2Text(str string) string {
	domDoc := xhtml.NewTokenizer(strings.NewReader(str))
	previousStartToken := domDoc.Token()
	var text string
loopDom:
	for {
		nx := domDoc.Next()
		switch {
		case nx == xhtml.ErrorToken:
			break loopDom // End of the document
		case nx == xhtml.StartTagToken:
			previousStartToken = domDoc.Token()
		case nx == xhtml.TextToken:
			if chk, _ := ks.Dstrpos(previousStartToken.Data, TextHtmlExcludeTags, false); chk {
				continue
			}

			text += " " + strings.TrimSpace(xhtml.UnescapeString(string(domDoc.Text())))
		}
	}

	return ks.RemoveSpace(text, false)
}
