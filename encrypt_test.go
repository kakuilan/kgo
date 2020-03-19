package kgo

import (
	"crypto/aes"
	"fmt"
	"strings"
	"testing"
)

func TestBase64Encode(t *testing.T) {
	str := []byte("This is an string to encod")
	res := KEncr.Base64Encode(str)
	if !strings.HasSuffix(res, "=") {
		t.Error("Base64Encode fail")
		return
	}
}

func BenchmarkBase64Encode(b *testing.B) {
	b.ResetTimer()
	str := []byte("This is an string to encod")
	for i := 0; i < b.N; i++ {
		KEncr.Base64Encode(str)
	}
}

func TestBase64Decode(t *testing.T) {
	str := "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
	_, err := KEncr.Base64Decode(str)
	if err != nil {
		t.Error("Base64Decode fail")
		return
	}
	_, err = KEncr.Base64Decode("#iu3498r")
	if err == nil {
		t.Error("Base64Decode fail")
		return
	}
	_, err = KEncr.Base64Decode("VGhpcy")
	_, err = KEncr.Base64Decode("VGhpcyB")
}

func BenchmarkBase64Decode(b *testing.B) {
	b.ResetTimer()
	str := "VGhpcyBpcyBhbiBlbmNvZGVkIHN0cmluZw=="
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64Decode(str)
	}
}

func TestBase64UrlEncodeDecode(t *testing.T) {
	str := []byte("This is an string to encod")
	res := KEncr.Base64UrlEncode(str)
	if strings.HasSuffix(res, "=") {
		t.Error("Base64UrlEncode fail")
		return
	}

	_, err := KEncr.Base64UrlDecode(res)
	if err != nil {
		t.Error("Base64UrlDecode fail")
		return
	}
}

func BenchmarkBase64UrlEncode(b *testing.B) {
	b.ResetTimer()
	str := []byte("This is an string to encod")
	for i := 0; i < b.N; i++ {
		KEncr.Base64UrlEncode(str)
	}
}

func BenchmarkBase64UrlDecode(b *testing.B) {
	b.ResetTimer()
	str := "VGhpcyBpcyBhbiBzdHJpbmcgdG8gZW5jb2Q"
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.Base64UrlDecode(str)
	}
}

func TestAuthCode(t *testing.T) {
	key := "123456"
	str := "hello world"

	res, _ := KEncr.AuthCode(str, key, true, 0)
	if res == "" {
		t.Error("AuthCode Encode fail")
		return
	}

	res2, _ := KEncr.AuthCode(res, key, false, 0)
	if res2 == "" {
		t.Error("AuthCode Decode fail")
		return
	}

	res, _ = KEncr.AuthCode(str, key, true, -3600)
	KEncr.AuthCode(res, key, false, 0)
	KEncr.AuthCode("", key, true, 0)
	KEncr.AuthCode("", "", true, 0)
	KEncr.AuthCode("7caeNfPt/N1zHdj5k/7i7pol6NHsVs0Cji7c15h4by1RYcrBoo7EEw==", key, false, 0)
	KEncr.AuthCode("7caeNfPt/N1zHdj5k/7i7pol6N", key, false, 0)
	KEncr.AuthCode("123456", "", false, 0)
	KEncr.AuthCode("1234#iu3498r", "", false, 0)
}

func BenchmarkAuthCodeEncode(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "hello world"
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(str, key, true, 0)
	}
}

func BenchmarkAuthCodeDecode(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "a79b5do3C9nbaZsAz5j3NQRj4e/L6N+y5fs2U9r1mO0LinOWtxmscg=="
	for i := 0; i < b.N; i++ {
		KEncr.AuthCode(str, key, false, 0)
	}
}

func TestPasswordHashVerify(t *testing.T) {
	pwd := []byte("123456")
	has, err := KEncr.PasswordHash(pwd)
	if err != nil {
		t.Error("PasswordHash fail")
		return
	}

	chk := KEncr.PasswordVerify(pwd, has)
	if !chk {
		t.Error("PasswordVerify fail")
		return
	}

	_, _ = KEncr.PasswordHash(pwd, 1)
	//慎用20以上,太耗时
	_, _ = KEncr.PasswordHash(pwd, 15)
	_, _ = KEncr.PasswordHash(pwd, 33)
}

func BenchmarkPasswordHash(b *testing.B) {
	b.ResetTimer()
	pwd := []byte("123456")
	for i := 0; i < b.N; i++ {
		//太耗时,只测试少量的
		if i > 10 {
			break
		}
		_, _ = KEncr.PasswordHash(pwd)
	}
}

func BenchmarkPasswordVerify(b *testing.B) {
	b.ResetTimer()
	pwd := []byte("123456")
	has := []byte("$2a$10$kCv6ljsVuTSI54oPkWulreEmUNTW/zj0Dgh6qF4Vz0w4C3gVf/w7a")
	for i := 0; i < b.N; i++ {
		//太耗时,只测试少量的
		if i > 10 {
			break
		}
		KEncr.PasswordVerify(pwd, has)
	}
}

func TestEasyEncryptDecrypt(t *testing.T) {
	key := "123456"
	str := "hello world你好!hello world你好!hello world你好!hello world你好!"
	enc := KEncr.EasyEncrypt(str, key)
	if enc == "" {
		t.Error("EasyEncrypt fail")
		return
	}

	dec := KEncr.EasyDecrypt(enc, key)
	if dec != str {
		t.Error("EasyDecrypt fail")
		return
	}

	dec = KEncr.EasyDecrypt("你好，世界！", key)
	if dec != "" {
		t.Error("EasyDecrypt fail")
		return
	}

	KEncr.EasyEncrypt("", key)
	KEncr.EasyEncrypt("", "")
	KEncr.EasyDecrypt(enc, "1qwer")
	KEncr.EasyDecrypt("123", key)
	KEncr.EasyDecrypt("1234#iu3498r", key)
}

func BenchmarkEasyEncrypt(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "hello world你好"
	for i := 0; i < b.N; i++ {
		KEncr.EasyEncrypt(str, key)
	}
}

func BenchmarkEasyDecrypt(b *testing.B) {
	b.ResetTimer()
	key := "123456"
	str := "e10azZaczdODqqimpcY"
	for i := 0; i < b.N; i++ {
		KEncr.EasyDecrypt(str, key)
	}
}

func TestHmacShaX(t *testing.T) {
	str := []byte("hello world")
	key := []byte("123456")
	res1 := KEncr.HmacShaX(str, key, 1)
	res2 := KEncr.HmacShaX(str, key, 256)
	res3 := KEncr.HmacShaX(str, key, 512)
	if res1 == "" || res2 == "" || res3 == "" {
		t.Error("HmacShaX fail")
		return
	}
}

func TestHmacShaXPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	str := []byte("hello world")
	key := []byte("123456")
	KEncr.HmacShaX(str, key, 4)
}

func BenchmarkHmacShaX(b *testing.B) {
	b.ResetTimer()
	str := []byte("hello world")
	key := []byte("123456")
	for i := 0; i < b.N; i++ {
		KEncr.HmacShaX(str, key, 256)
	}
}

func TestPkcs7PaddingUnPadding(t *testing.T) {
	var emp1 []byte
	var emp2 = []byte("")
	key1 := []byte("1234")
	dat1 := []byte{49, 50, 51, 52, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12}
	dat2 := []byte{49, 50, 51, 52, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	var tests = []struct {
		cipher    []byte
		orig      []byte
		size      int
		zero      bool
		expected1 []byte
		expected2 []byte
	}{
		{nil, nil, aes.BlockSize, false, nil, nil},
		{emp1, emp1, aes.BlockSize, false, nil, nil},
		{emp2, emp2, aes.BlockSize, false, nil, nil},
		{key1, key1, 0, false, nil, nil},
		{key1, dat1, aes.BlockSize, false, dat1, key1},
		{key1, dat2, aes.BlockSize, true, dat2, nil},
		{key1, dat2, aes.BlockSize, false, dat1, emp1},
	}

	for _, test := range tests {
		actual1 := pkcs7Padding(test.cipher, test.size, test.zero)
		if !KArr.IsEqualArray(actual1, test.expected1) {
			t.Errorf("Expected pkcs7Padding(%v, %d, %t) to be %v, got %v", test.cipher, test.size, test.zero, test.expected1, actual1)
		}

		actual2 := pkcs7UnPadding(test.orig, test.size)
		if !KArr.IsEqualArray(actual2, test.expected2) {
			t.Errorf("Expected pkcs7UnPadding(%v, %d) to be %v, got %v", test.orig, test.size, test.expected2, actual2)
		}
	}
}

func BenchmarkPkcs7Padding(b *testing.B) {
	b.ResetTimer()
	str := []byte("1234")
	for i := 0; i < b.N; i++ {
		pkcs7Padding(str, 16, false)
	}
}

func BenchmarkPkcs7UnPadding(b *testing.B) {
	b.ResetTimer()
	data := []byte{49, 50, 51, 52, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12}
	for i := 0; i < b.N; i++ {
		pkcs7UnPadding(data, 16)
	}
}

func TestZeroPaddingUnPadding(t *testing.T) {
	key := []byte("hello")
	ori := zeroPadding(key, 16)
	res := zeroUnPadding(ori)

	if ori == nil {
		t.Error("zeroPadding fail")
		return
	}

	if !KArr.IsEqualArray(key, res) {
		t.Error("zeroUnPadding fail")
		return
	}
}

func BenchmarkZeroPadding(b *testing.B) {
	b.ResetTimer()
	key := []byte("hello")
	for i := 0; i < b.N; i++ {
		zeroPadding(key, 16)
	}
}

func BenchmarkZeroUnPadding(b *testing.B) {
	b.ResetTimer()
	ori := []byte{104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < b.N; i++ {
		zeroUnPadding(ori)
	}
}

func TestAesCBCEncryptDecrypt(t *testing.T) {
	ori := []byte("hello")
	key := []byte("1234567890123456")
	emp := []byte("")
	var err error
	var enc, des []byte

	_, err = KEncr.AesCBCEncrypt(ori, []byte("123"))
	if err == nil {
		t.Error("AesCBCEncrypt fail")
		return
	}

	enc, err = KEncr.AesCBCEncrypt(ori, key)
	des, err = KEncr.AesCBCDecrypt(enc, key)
	if !KArr.IsEqualArray(ori, des) {
		t.Error("AesCBCEncrypt fail")
		return
	}

	enc, err = KEncr.AesCBCEncrypt(ori, key, PKCS_SEVEN)
	des, err = KEncr.AesCBCDecrypt(enc, key, PKCS_SEVEN)
	if !KArr.IsEqualArray(ori, des) {
		t.Error("AesCBCEncrypt fail")
		return
	}

	enc, err = KEncr.AesCBCEncrypt(emp, key, PKCS_SEVEN)
	des, err = KEncr.AesCBCDecrypt(enc, key, PKCS_SEVEN)
	if !KArr.IsEqualArray(emp, des) {
		t.Error("AesCBCEncrypt fail")
		return
	}

	enc, err = KEncr.AesCBCEncrypt(ori, key, PKCS_ZERO)
	des, err = KEncr.AesCBCDecrypt(enc, key, PKCS_ZERO)
	if !KArr.IsEqualArray(ori, des) {
		t.Error("AesCBCEncrypt fail")
		return
	}

	enc = []byte{83, 28, 170, 254, 29, 174, 21, 129, 241, 233, 243, 84, 1, 250, 95, 122, 104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	des, err = KEncr.AesCBCDecrypt(enc, key, PKCS_ZERO)
	if err == nil {
		t.Error("AesCBCDecrypt fail")
		return
	}

	_, err = KEncr.AesCBCDecrypt(enc, []byte("1"))
	if err == nil {
		t.Error("AesCBCDecrypt fail")
		return
	}

	_, err = KEncr.AesCBCDecrypt([]byte("1234"), key)
	if err == nil {
		t.Error("AesCBCDecrypt fail")
		return
	}

}

func BenchmarkAesCBCEncrypt(b *testing.B) {
	b.ResetTimer()
	ori := []byte("hello")
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCEncrypt(ori, key)
	}
}

func BenchmarkAesCBCDecrypt(b *testing.B) {
	b.ResetTimer()
	enc := []byte{214, 214, 97, 208, 185, 68, 246, 40, 124, 3, 155, 58, 5, 84, 136, 10, 104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCBCDecrypt(enc, key, PKCS_ZERO)
	}
}

func TestAesCFBEncryptDecrypt(t *testing.T) {
	ori := []byte("hello")
	key := []byte("1234567890123456")
	emp := []byte("")
	var err error
	var enc, des []byte

	_, err = KEncr.AesCFBEncrypt(ori, []byte("123"))
	if err == nil {
		t.Error("AesCFBEncrypt fail")
		return
	}

	enc, err = KEncr.AesCFBEncrypt(ori, key)
	des, err = KEncr.AesCFBDecrypt(enc, key)
	if !KArr.IsEqualArray(ori, des) {
		t.Error("AesCFBEncrypt fail")
		return
	}

	enc, err = KEncr.AesCFBEncrypt(emp, key)
	des, err = KEncr.AesCFBDecrypt(enc, key)
	if !KArr.IsEqualArray(emp, des) {
		t.Error("AesCFBEncrypt fail")
		return
	}

	enc = []byte{83, 28, 170, 254, 29, 174, 21, 129, 241, 233, 243, 84, 1, 250, 95, 122, 104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	des, err = KEncr.AesCFBDecrypt(enc, key)
	if err == nil {
		t.Error("AesCFBDecrypt fail")
		return
	}

	_, err = KEncr.AesCFBDecrypt(enc, []byte("1"))
	if err == nil {
		t.Error("AesCFBDecrypt fail")
		return
	}

	_, err = KEncr.AesCFBDecrypt([]byte("1234"), key)
	if err == nil {
		t.Error("AesCFBDecrypt fail")
		return
	}

}

func BenchmarkAesCFBEncrypt(b *testing.B) {
	b.ResetTimer()
	ori := []byte("hello")
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCFBEncrypt(ori, key)
	}
}

func BenchmarkAesCFBDecrypt(b *testing.B) {
	b.ResetTimer()
	enc := []byte{20, 193, 94, 23, 183, 173, 65, 237, 222, 161, 169, 129, 125, 200, 110, 132, 104, 101, 108, 108, 111, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11}
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCFBDecrypt(enc, key)
	}
}

func TestAesCTREncryptDecrypt(t *testing.T) {
	ori := []byte("hello")
	key := []byte("1234567890123456")
	emp := []byte("")
	var err error
	var enc, des []byte

	_, err = KEncr.AesCTREncrypt(ori, []byte("123"))
	if err == nil {
		t.Error("AesCTREncrypt fail")
		return
	}

	enc, err = KEncr.AesCTREncrypt(ori, key)
	des, err = KEncr.AesCTRDecrypt(enc, key)
	if !KArr.IsEqualArray(ori, des) {
		t.Error("AesCTREncrypt fail")
		return
	}

	enc, err = KEncr.AesCTREncrypt(emp, key)
	des, err = KEncr.AesCTRDecrypt(enc, key)
	if !KArr.IsEqualArray(emp, des) {
		t.Error("AesCTREncrypt fail")
		return
	}

	enc = []byte{83, 28, 170, 254, 29, 174, 21, 129, 241, 233, 243, 84, 1, 250, 95, 122, 104, 101, 108, 108, 111, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	des, err = KEncr.AesCTRDecrypt(enc, key)
	if err == nil {
		t.Error("AesCTRDecrypt fail")
		return
	}

	_, err = KEncr.AesCTRDecrypt(enc, []byte("1"))
	if err == nil {
		t.Error("AesCTRDecrypt fail")
		return
	}

	_, err = KEncr.AesCTRDecrypt([]byte("1234"), key)
	if err == nil {
		t.Error("AesCTRDecrypt fail")
		return
	}
	//fmt.Printf("%v \n", enc)
	//fmt.Printf("%v \n", des)
	//fmt.Printf("%s \n", des)
}

func BenchmarkAesCTREncrypt(b *testing.B) {
	b.ResetTimer()
	ori := []byte("hello")
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCTREncrypt(ori, key)
	}
}

func BenchmarkAesCTRDecrypt(b *testing.B) {
	b.ResetTimer()
	enc := []byte{120, 134, 68, 163, 146, 37, 245, 79, 12, 237, 58, 77, 188, 123, 24, 77, 104, 101, 108, 108, 111, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11, 11}
	key := []byte("1234567890123456")
	for i := 0; i < b.N; i++ {
		_, _ = KEncr.AesCTRDecrypt(enc, key)
	}
}
