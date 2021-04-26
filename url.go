package kgo

import (
	"net/url"
	"strings"
)

// ParseStr 将URI查询字符串转换为字典.
func (ks *LkkString) ParseStr(encodedString string, result map[string]interface{}) error {
	// split encodedString.
	if encodedString[0] == '?' {
		encodedString = strings.TrimLeft(encodedString, "?")
	}

	parts := strings.Split(encodedString, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			return err
		}
		for key[0] == ' ' && key[1:] != "" {
			key = key[1:]
		}
		if key == "" || key[0] == '[' {
			continue
		}
		value, err := url.QueryUnescape(part[pos+1:])
		if err != nil {
			return err
		}

		// split into multiple keys
		var keys []string
		left := 0
		for i, k := range key {
			if k == '[' && left == 0 {
				left = i
			} else if k == ']' {
				if left > 0 {
					if len(keys) == 0 {
						keys = append(keys, key[:left])
					}
					keys = append(keys, key[left+1:i])
					left = 0
					if i+1 < len(key) && key[i+1] != '[' {
						break
					}
				}
			}
		}
		if len(keys) == 0 {
			keys = append(keys, key)
		}
		// first key
		first := ""
		for i, chr := range keys[0] {
			if chr == ' ' || chr == '.' || chr == '[' {
				first += "_"
			} else {
				first += string(chr)
			}
			if chr == '[' {
				first += keys[0][i+1:]
				break
			}
		}
		keys[0] = first

		// build nested map
		if err := buildQueryMap(result, keys, value); err != nil {
			return err
		}
	}

	return nil
}

// ParseUrl 解析URL,返回其组成部分.
// component为需要返回的组成;
// -1: all; 1: scheme; 2: host; 4: port; 8: user; 16: pass; 32: path; 64: query; 128: fragment .
func (ks *LkkString) ParseUrl(str string, component int16) (map[string]string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return nil, err
	}
	if component == -1 {
		component = 1 | 2 | 4 | 8 | 16 | 32 | 64 | 128
	}
	var res = make(map[string]string)
	if (component & 1) == 1 {
		res["scheme"] = u.Scheme
	}
	if (component & 2) == 2 {
		res["host"] = u.Hostname()
	}
	if (component & 4) == 4 {
		res["port"] = u.Port()
	}
	if (component & 8) == 8 {
		res["user"] = u.User.Username()
	}
	if (component & 16) == 16 {
		res["pass"], _ = u.User.Password()
	}
	if (component & 32) == 32 {
		res["path"] = u.Path
	}
	if (component & 64) == 64 {
		res["query"] = u.RawQuery
	}
	if (component & 128) == 128 {
		res["fragment"] = u.Fragment
	}
	return res, nil
}

// UrlEncode 编码 URL 字符串.
func (ks *LkkString) UrlEncode(str string) string {
	return url.QueryEscape(str)
}

// UrlDecode 解码已编码的 URL 字符串.
func (ks *LkkString) UrlDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}
