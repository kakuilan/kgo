// +build windows

package kgo

import "strings"

// formatPath 格式化路径
func formatPath(fpath string) string {
	//替换特殊字符
	fpath = strings.NewReplacer(`|`, "", "", `<`, "", `>`, "", `?`, "", `\`, "/").Replace(fpath)
	//替换连续斜杠
	fpath = RegFormatDir.ReplaceAllString(fpath, "/")
	return fpath
}
