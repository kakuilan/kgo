package kgo

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	//haystack类型不对
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	//数组
	arr := [5]int{1, 2, 3, 4, 5}
	it := 2
	if !KArr.InArray(it, arr) {
		t.Error("InArray fail")
		return
	}

	//字典
	mp := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	it2 := "a"
	it3 := "bb"
	if KArr.InArray(it2, mp) {
		t.Error("InArray fail")
		return
	} else if !KArr.InArray(it3, mp) {
		t.Error("InArray fail")
		return
	}

	if KArr.InArray(it2, "abc") {
		t.Error("InArray fail")
		return
	}
}

func BenchmarkInArray(b *testing.B) {
	b.ResetTimer()
	sli := []string{"a", "b", "c", "d", "e"}
	it := "d"
	for i := 0; i < b.N; i++ {
		KArr.InArray(it, sli)
	}
}

func TestEasyEncryptDecrypt(t *testing.T) {
	key := "123456"
	str := "hello world你好"
	enc := KEncr.EasyEncrypt(str, key)
	println("encode:", enc)
	if enc == "" {
		t.Error("EasyEncrypt fail")
		return
	}

	dec := KEncr.EasyDecrypt(enc, key)
	println("decode:", dec)
	if dec == "" {
		t.Error("EasyDecrypt fail")
		return
	}

	KEncr.EasyEncrypt("", key)
	KEncr.EasyEncrypt("", "")
	KEncr.EasyDecrypt(enc, "1qwer")
	KEncr.EasyDecrypt("123", key)
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
