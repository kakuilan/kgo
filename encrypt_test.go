package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt_Base64Encode(t *testing.T) {
	var res []byte

	res = KEncr.Base64Encode([]byte(""))
	assert.Nil(t, res)

	res = KEncr.Base64Encode(btysHello)
	assert.Contains(t, string(res), "=")
}

func BenchmarkEncrypt_Base64Encode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.Base64Encode(btysHello)
	}
}

func TestEncrypt_Base64Decode(t *testing.T) {
	var res []byte
	var err error

	res, err = KEncr.Base64Decode([]byte(""))
	assert.Nil(t, res)
	assert.Nil(t, err)

	res, err = KEncr.Base64Decode([]byte(b64Hello))
	assert.Equal(t, strHello, string(res))

	//不合法
	_, err = KEncr.Base64Decode([]byte("#iu3498r"))
	assert.NotNil(t, err)
}

func BenchmarkEncrypt_Base64Decode(b *testing.B) {
	b.ResetTimer()
	bs := []byte(b64Hello)
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64Decode(bs)
	}
}

func TestEncrypt_Base64UrlEncode(t *testing.T) {
	var res []byte

	res = KEncr.Base64UrlEncode([]byte(""))
	assert.Nil(t, res)

	res = KEncr.Base64UrlEncode([]byte(str2Code))
	assert.Equal(t, b64UrlCode, string(res))
}

func BenchmarkEncrypt_Base64UrlEncode(b *testing.B) {
	b.ResetTimer()
	bs := []byte(str2Code)
	for i := 0; i < b.N; i++ {
		KEncr.Base64UrlEncode(bs)
	}
}

func TestEncrypt_Base64UrlDecode(t *testing.T) {
	var res []byte
	var err error

	res, err = KEncr.Base64UrlDecode([]byte(""))
	assert.Nil(t, res)
	assert.Nil(t, err)

	res, err = KEncr.Base64UrlDecode([]byte(b64UrlCode))
	assert.Equal(t, str2Code, string(res))

	//不合法
	_, err = KEncr.Base64UrlDecode([]byte("#iu3498r"))
	assert.NotNil(t, err)
}

func BenchmarkEncrypt_Base64UrlDecode(b *testing.B) {
	b.ResetTimer()
	bs := []byte(b64UrlCode)
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64UrlDecode(bs)
	}
}
