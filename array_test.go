package kgo

import (
	"fmt"
	"testing"
)

func TestIsArrayOrSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"

	res1 := KArr.IsArrayOrSlice(arr, 1)
	res2 := KArr.IsArrayOrSlice(arr, 2)
	res3 := KArr.IsArrayOrSlice(arr, 3)
	if res1 != 10 || res2 != -1 || res3 != 10 {
		t.Error("IsArrayOrSlice fail")
		return
	}

	res4 := KArr.IsArrayOrSlice(sli, 1)
	res5 := KArr.IsArrayOrSlice(sli, 2)
	res6 := KArr.IsArrayOrSlice(sli, 3)
	if res4 != -1 || res5 != 5 || res6 != 5 {
		t.Error("IsArrayOrSlice fail")
		return
	}

	KArr.IsArrayOrSlice(sli, 6)
}

func BenchmarkIsArrayOrSlice(b *testing.B) {
	b.ResetTimer()
	var arr = [10]int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		KArr.IsArrayOrSlice(arr, 1)
	}
}

func TestInArray(t *testing.T) {
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

func TestArrayFlip(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := KArr.ArrayFlip(mp)
	if val, ok := res[1]; !ok || fmt.Sprintf("%v", val) != "a" {
		t.Error("ArrayFlip fail")
		return
	}

	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = KArr.ArrayFlip(sli)

	KArr.ArrayFlip("hello")
}

func BenchmarkArrayFlip(b *testing.B) {
	b.ResetTimer()
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		KArr.ArrayFlip(mp)
	}
}

func TestArrayKeys(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := KArr.ArrayKeys(mp)
	if len(res) != 3 {
		t.Error("ArrayKeys fail")
		return
	}

	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = KArr.ArrayKeys(sli)
	if len(res) != 5 {
		t.Error("ArrayKeys fail")
		return
	}

	KArr.ArrayKeys("hello")
}

func BenchmarkArrayKeys(b *testing.B) {
	b.ResetTimer()
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeys(mp)
	}
}

func TestArrayValues(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	res := KArr.ArrayValues(mp)
	if len(res) != 3 {
		t.Error("ArrayValues fail")
		return
	}

	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = KArr.ArrayValues(sli)
	if len(res) != 5 {
		t.Error("ArrayValues fail")
		return
	}

	KArr.ArrayValues("hello")
}

func BenchmarkArrayValues(b *testing.B) {
	b.ResetTimer()
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		KArr.ArrayValues(mp)
	}
}

func TestSliceMerge(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"

	res1 := KArr.SliceMerge(false, arr, sli)
	if len(res1) != 15 {
		t.Error("SliceMerge fail")
		return
	}

	res2 := KArr.SliceMerge(true, arr, sli)
	if len(res2) != 13 {
		t.Error("SliceMerge fail")
		return
	}
	KArr.SliceMerge(true)
	KArr.SliceMerge(false, "hellow")
}

func BenchmarkSliceMerge(b *testing.B) {
	b.ResetTimer()
	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	for i := 0; i < b.N; i++ {
		KArr.SliceMerge(false, arr, sli)
	}
}
