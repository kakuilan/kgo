package kgo

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// md5Str 计算字符串的 MD5 散列值
func md5Str(str []byte, length uint8) string {
	var res string
	hash := md5.New()
	hash.Write(str)

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

// sha1Str 计算字符串的 sha1 散列值
func sha1Str(str []byte) string {
	hash := sha1.New()
	hash.Write(str)
	return hex.EncodeToString(hash.Sum(nil))
}
