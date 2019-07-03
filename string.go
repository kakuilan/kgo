package gohelper

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

// Nl2br 将换行符转换为br标签
func (ks *LkkString) Nl2br(html string) string {
	if html == "" {
		return ""
	}
	return strings.Replace(html, "\n", "<br />", -1)
}

// StripTags 过滤html和php标签
func (ks *LkkString) StripTags(html string) string {
	if html == "" {
		return ""
	}
	re := regexp.MustCompile(`<(.|\n)*?>`)
	return re.ReplaceAllString(html, "")
}

// IsBinary 是否二进制的字符串
func (ks *LkkString) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}

// Md5 获取字符串md5值,length指定结果长度32/16
func (ks *LkkString) Md5(str string, length uint8) string {
	var res string
	hash := md5.New()
	hash.Write([]byte(str))

	hashInBytes := hash.Sum(nil)
	if length > 0 && length < 32 {
		dst := make([]byte, hex.EncodedLen(len(hashInBytes)))
		hex.Encode(dst, hashInBytes)
		res = string(dst[:length])
	} else {
		res = hex.EncodeToString(hashInBytes)
	}

	return res
}
