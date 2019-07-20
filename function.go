package kgo

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// md5Str 计算字符串的 MD5 散列值
func md5Str(str []byte, length uint8) []byte {
	var res []byte
	hash := md5.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	dst := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(dst, hBytes)
	if length > 0 && length < 32 {
		res = dst[:length]
	} else {
		res = dst
	}

	return res
}

// sha1Str 计算字符串的 sha1 散列值
func sha1Str(str []byte) []byte {
	hash := sha1.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}

// sha256Str 计算字符串的 sha256 散列值
func sha256Str(str []byte) []byte {
	hash := sha256.New()
	hash.Write(str)

	hBytes := hash.Sum(nil)
	res := make([]byte, hex.EncodedLen(len(hBytes)))
	hex.Encode(res, hBytes)
	return res
}
