package kgo

import (
	"encoding/base64"
)

// Base64Encode 使用 MIME base64 对数据进行编码
func (ke *LkkEncrypt) Base64Encode(str []byte) string {
	return base64.StdEncoding.EncodeToString(str)
}

// Base64Decode 对使用 MIME base64 编码的数据进行解码
func (ke *LkkEncrypt) Base64Decode(str string) ([]byte, error) {
	switch len(str) % 4 {
	case 2:
		str += "=="
	case 3:
		str += "="
	}

	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return data, nil
}
