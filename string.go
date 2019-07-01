package gohelper

import (
	"strings"
	"regexp"
)

// Nl2br Inserts HTML line breaks before all newlines in a string.
func (ks *LkkString) Nl2br(html string) string {
	if html == "" {
		return ""
	}else{
		return strings.Replace(html, "\n", "<br />", -1)
	}
}

// StripTags Strip HTML and PHP tags from a string
func (ks *LkkString) StripTags(html string) string {
	if html == "" {
		return ""
	}else{
		re := regexp.MustCompile(`<(.|\n)*?>`)
		return re.ReplaceAllString(html,"")
	}
}

// IsBinary determines whether the content is a binary file content.
func (kf *LkkString) IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}
	return false
}