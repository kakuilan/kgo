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

func TestArrayFill(t *testing.T) {
	num := uint(4)
	res := KArr.ArrayFill(0, num, "abc")
	if len(res) != int(num) {
		t.Error("InArray fail")
		return
	}
}

func BenchmarkArrayFill(b *testing.B) {
	b.ResetTimer()
	num := uint(4)
	for i := 0; i < b.N; i++ {
		KArr.ArrayFill(0, num, "abc")
	}
}
