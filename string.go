package gohelper

import (
	"strings"
	"regexp"
)

// Nl2br 将换行符转换为br标签
func (ks *LkkString) Nl2br(html string) string {
	if html == "" {
		return ""
	}else{
		return strings.Replace(html, "\n", "<br />", -1)
	}
}

// StripTags 过滤html和php标签
func (ks *LkkString) StripTags(html string) string {
	if html == "" {
		return ""
	}else{
		re := regexp.MustCompile(`<(.|\n)*?>`)
		return re.ReplaceAllString(html,"")
	}
}

// IsBinary 是否二进制的字符串
func (kf *LkkString) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}