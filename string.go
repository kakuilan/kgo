package kgo

import (
	"bytes"
	"github.com/json-iterator/go"
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
