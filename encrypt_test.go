package kgo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncrypt_Base64Encode(t *testing.T) {
	var res []byte

	res = KEncr.Base64Encode(bytEmp)
	assert.Nil(t, res)

	res = KEncr.Base64Encode(bytsHello)
	assert.Contains(t, string(res), "=")
}

func BenchmarkEncrypt_Base64Encode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.Base64Encode(bytsHello)
	}
}

func TestEncrypt_Base64Decode(t *testing.T) {
	var res []byte
	var err error

	res, err = KEncr.Base64Decode(bytEmp)
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

	res = KEncr.Base64UrlEncode(bytEmp)
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

	res, err = KEncr.Base64UrlDecode(bytEmp)
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

func TestEncrypt_AuthCode(t *testing.T) {
	var res, res2 []byte
	var exp int64

	res, exp = KEncr.AuthCode(bytsHello, bytSpeedLight, true, 0)
	assert.NotNil(t, res)

	res2, exp = KEncr.AuthCode(res, bytSpeedLight, false, 0)
	assert.Equal(t, string(bytsHello), string(res2))

	//过期
	res, exp = KEncr.AuthCode(bytsHello, bytSpeedLight, true, -3600)
	res2, exp = KEncr.AuthCode(res, bytSpeedLight, false, 0)
	assert.Empty(t, res2)
	assert.Greater(t, exp, int64(0))

	//空串
	res, exp = KEncr.AuthCode(bytEmp, bytSpeedLight, true, 0)
	assert.Empty(t, res)

	//空密钥
	res, exp = KEncr.AuthCode(bytsHello, bytEmp, true, 0)
	assert.NotEmpty(t, res)

	//不合法
	KEncr.AuthCode([]byte("7caeNfPt/N1zHdj5k/7i7pol6NHsVs0Cji7c15h4by1RYcrBoo7EEw=="), bytSpeedLight, false, 0)
	KEncr.AuthCode([]byte("7caeNfPt/N1zHdj5k/7i7pol6N"), bytSpeedLight, false, 0)
	KEncr.AuthCode([]byte("123456"), bytEmp, false, 0)
	KEncr.AuthCode([]byte("1234#iu3498r"), bytEmp, false, 0)
}

func BenchmarkEncrypt_AuthCode_Encode(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(bytsHello, bytSpeedLight, true, 0)
	}
}

func BenchmarkEncrypt_AuthCode_Decode(b *testing.B) {
	b.ResetTimer()
	bs := []byte("b0140641v309wJW2_-MvoovhaHKtHLBvZ2JFsvirqYQK5144m-wQJlez8XBfHkCohr3clxPR")
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(bs, bytSpeedLight, false, 0)
	}
}

func TestEncrypt_PasswordHash(t *testing.T) {
	var res []byte
	var err error

	//空密码
	res, err = KEncr.PasswordHash(bytEmp)
	assert.NotNil(t, res)
	assert.Nil(t, err)

	res, err = KEncr.PasswordHash(bytsHello)
	assert.NotEmpty(t, res)

	//慎用20以上,太耗时
	_, _ = KEncr.PasswordHash(bytsHello, 1)
	_, _ = KEncr.PasswordHash(bytsHello, 33)
}

func BenchmarkEncrypt_PasswordHash(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.PasswordHash(bytsHello)
	}
}

func TestEncrypt_PasswordVerify(t *testing.T) {
	var res bool

	res = KEncr.PasswordVerify(bytsHello, bytsPasswd)
	assert.True(t, res)

	res = KEncr.PasswordVerify(bytSpeedLight, bytsPasswd)
	assert.False(t, res)
}

func BenchmarkEncrypt_PasswordVerify(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.PasswordVerify(bytsHello, bytsPasswd)
	}
}

func TestEncrypt_EasyEncryptDecrypt(t *testing.T) {
	var enc, dec []byte

	enc = KEncr.EasyEncrypt(bytsHello, bytSpeedLight)
	assert.NotEmpty(t, enc)

	dec = KEncr.EasyDecrypt(enc, bytSpeedLight)
	assert.Equal(t, bytsHello, dec)

	//空字符串
	enc = KEncr.EasyEncrypt(bytEmp, bytSpeedLight)
	assert.Empty(t, enc)

	//空密钥
	enc = KEncr.EasyEncrypt(bytsHello, bytEmp)
	assert.NotEmpty(t, enc)

	//密钥错误
	dec = KEncr.EasyDecrypt(enc, bytSpeedLight)
	assert.Empty(t, dec)

	//密文错误
	dec = KEncr.EasyDecrypt(bytsHello, bytSpeedLight)
	assert.Empty(t, dec)
	dec = KEncr.EasyDecrypt([]byte("1234#iu3498r"), bytSpeedLight)
	assert.Empty(t, dec)
}

func BenchmarkEncrypt_EasyEncrypt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.EasyEncrypt(bytsHello, bytSpeedLight)
	}
}

func BenchmarkEncrypt_EasyDecrypt(b *testing.B) {
	b.ResetTimer()
	bs := []byte(esyenCode)
	for i := 0; i < b.N; i++ {
		KEncr.EasyDecrypt(bs, bytSpeedLight)
	}
}

func TestEncrypt_HmacShaX(t *testing.T) {
	defer func() {
		r := recover()
		assert.NotEmpty(t, r)
	}()

	var res []byte

	res = KEncr.HmacShaX(bytsHello, bytSpeedLight, 1)
	assert.NotEmpty(t, res)

	res = KEncr.HmacShaX(bytsHello, bytSpeedLight, 256)
	assert.NotEmpty(t, res)

	res = KEncr.HmacShaX(bytsHello, bytSpeedLight, 512)
	assert.NotEmpty(t, res)

	//不合法
	KEncr.HmacShaX(bytsHello, bytSpeedLight, 44)
}

func BenchmarkEncrypt_HmacShaX(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KEncr.HmacShaX(bytsHello, bytSpeedLight, 256)
	}
}

func TestEncrypt_AesCBCEncryptDecrypt(t *testing.T) {
	var err error
	var enc, des []byte

	//加密
	enc, err = KEncr.AesCBCEncrypt(bytsHello, bytCryptKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, enc)

	//解密
	des, err = KEncr.AesCBCDecrypt(enc, bytCryptKey)
	assert.Equal(t, bytsHello, des)

	//密钥不合法
	_, err = KEncr.AesCBCEncrypt(bytsHello, bytSpeedLight)
	assert.NotNil(t, err)

	//错误的密钥
	_, err = KEncr.AesCBCDecrypt(enc, []byte("1234561234567890"))
	assert.NotNil(t, err)

	//填充方式-PKCS_SEVEN
	enc, _ = KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_SEVEN)
	des, _ = KEncr.AesCBCDecrypt(enc, bytCryptKey, PKCS_SEVEN)
	assert.NotEmpty(t, enc)
	assert.Equal(t, bytsHello, des)

	//填充方式-PKCS_ZERO
	enc, _ = KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_ZERO)
	des, _ = KEncr.AesCBCDecrypt(enc, bytCryptKey, PKCS_ZERO)
	assert.NotEmpty(t, enc)
	assert.Equal(t, bytsHello, des)

	//空字符串
	enc, err = KEncr.AesCBCEncrypt(bytEmp, bytCryptKey)
	des, err = KEncr.AesCBCDecrypt(enc, bytCryptKey)
	assert.NotEmpty(t, enc)
	assert.Empty(t, des)

	//错误的加密串
	enc = []byte{83, 28, 170, 254, 29, 174, 21, 129, 241, 233, 243, 84, 1, 250, 95, 122, 104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	des, err = KEncr.AesCBCDecrypt(enc, bytCryptKey, PKCS_ZERO)
	assert.NotNil(t, err)
}

func BenchmarkEncrypt_AesCBCEncrypt_PKCS_SEVEN(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_SEVEN)
	}
}

func BenchmarkEncrypt_AesCBCDecrypt_PKCS_SEVEN(b *testing.B) {
	b.ResetTimer()
	bs, _ := KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_SEVEN)
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCDecrypt(bs, bytCryptKey, PKCS_SEVEN)
	}
}

func BenchmarkEncrypt_AesCBCEncrypt_PKCS_ZERO(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_ZERO)
	}
}

func BenchmarkEncrypt_AesCBCDecrypt_PKCS_ZERO(b *testing.B) {
	b.ResetTimer()
	bs, _ := KEncr.AesCBCEncrypt(bytsHello, bytCryptKey, PKCS_ZERO)
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCDecrypt(bs, bytCryptKey, PKCS_ZERO)
	}
}
