package kgo

import "encoding/base64"

// Base64Encode 使用 MIME base64 对数据进行编码.
func (ke *LkkEncrypt) Base64Encode(str []byte) []byte {
	l := len(str)
	if l > 0 {
		buf := make([]byte, base64.StdEncoding.EncodedLen(len(str)))
		base64.StdEncoding.Encode(buf, str)
		return buf
	}

	return nil
}
