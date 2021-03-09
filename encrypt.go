package kgo

import "encoding/base64"

// Base64Encode 使用 MIME base64 对数据进行编码.
func (ke *LkkEncrypt) Base64Encode(str []byte) []byte {
	l := len(str)
	if l > 0 {
		buf := make([]byte, base64.StdEncoding.EncodedLen(l))
		base64.StdEncoding.Encode(buf, str)
		return buf
	}

	return nil
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码.
func (ke *LkkEncrypt) Base64Decode(str []byte) ([]byte, error) {
	l := len(str)
	if l > 0 {
		dbuf := make([]byte, base64.StdEncoding.DecodedLen(l))
		n, err := base64.StdEncoding.Decode(dbuf, str)
		return dbuf[:n], err
	}

	return nil, nil
}
