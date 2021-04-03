package kgo

import (
	"bytes"
	"github.com/json-iterator/go"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strings"
	"unicode/utf8"
)

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
func (ks *LkkString) IsUtf8(str string) bool {
	return str != "" && utf8.ValidString(str)
}

// IsGbk 字符串是否GBK编码.
func (ks *LkkString) IsGbk(data []byte) (res bool) {
	length := len(data)
	var i, j int
	for i < length {
		j = i + 1
		//大于127的使用双字节编码,且落在gbk编码范围内的字符
		//GBK中每个汉字包含两个字节，第一个字节(首字节)的范围是0x81-0xFE(即129-254),第二个字节(尾字节)的范围是0x40-0xFE(即64-254)
		if data[i] > 0x7f && j < length {
			if data[i] >= 0x81 &&
				data[i] <= 0xfe &&
				data[j] >= 0x40 &&
				data[j] <= 0xfe {
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
