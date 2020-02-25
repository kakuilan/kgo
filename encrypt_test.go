package kgo

import (
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
