package kgo

import (
	"fmt"
	"testing"
)

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
	num := 4
	res := KArr.ArrayFill("abc", num)
	if len(res) != num {
		t.Error("InArray fail")
		return
	}
	KArr.ArrayFill("abc", 0)
}

func BenchmarkArrayFill(b *testing.B) {
	b.ResetTimer()
	num := 10
	for i := 0; i < b.N; i++ {
		KArr.ArrayFill("abc", num)
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

func TestMapMerge(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	mp1 := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	mp2 := map[string]int{
		"h": 1,
		"i": 2,
		"j": 3,
	}

	res := KArr.MapMerge(true, mp1, mp2)
	_, err := KStr.JsonEncode(res)
	if err != nil {
		t.Error("MapMerge fail")
		return
	}
	KArr.MapMerge(false)
	KArr.MapMerge(false, mp1, mp2)
	KArr.MapMerge(false, mp1, mp2, "hello")
}

func BenchmarkMapMerge(b *testing.B) {
	b.ResetTimer()
	mp1 := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	mp2 := map[string]int{
		"h": 1,
		"i": 2,
		"j": 3,
	}
	for i := 0; i < b.N; i++ {
		KArr.MapMerge(true, mp1, mp2)
	}
}

func TestArrayChunk(t *testing.T) {
	size := 3
	var arr = [11]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res1 := KArr.ArrayChunk(arr, size)
	if len(res1) != 4 {
		t.Error("ArrayChunk fail")
		return
	}

	var myslice []int
	KArr.ArrayChunk(myslice, 1)
}

func TestArrayChunkPanicSize(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var arr = [11]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	KArr.ArrayChunk(arr, 0)
}

func TestArrayChunkPanicType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArrayChunk("hello", 1)
}

func BenchmarkArrayChunk(b *testing.B) {
	b.ResetTimer()
	size := 3
	var arr = [11]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayChunk(arr, size)
	}
}

func TestArrayPad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var sli []int
	var arr = [3]string{"a", "b", "c"}

	res1 := KArr.ArrayPad(sli, 5, 1)
	res2 := KArr.ArrayPad(arr, 6, "d")
	res3 := KArr.ArrayPad(arr, -6, "d")
	res4 := KArr.ArrayPad(arr, 2, "d")
	if len(res1) != 5 || len(res2) != 6 || fmt.Sprintf("%v", res3[0]) != "d" || len(res4) != 3 {
		t.Error("ArrayPad fail")
		return
	}

	KArr.ArrayPad("hello", 2, "d")
}

func BenchmarkArrayPad(b *testing.B) {
	b.ResetTimer()
	var arr = [3]string{"a", "b", "c"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayPad(arr, 10, "d")
	}
}

func TestArraySlice(t *testing.T) {
	var sli []int
	var arr = [6]string{"a", "b", "c", "d", "e", "f"}

	res1 := KArr.ArraySlice(sli, 0, 1)
	res2 := KArr.ArraySlice(arr, 1, 2)
	res3 := KArr.ArraySlice(arr, -3, 2)
	res4 := KArr.ArraySlice(arr, -3, 4)
	if len(res1) != 0 || len(res2) != 2 || len(res3) != 2 || len(res4) != 3 {
		t.Error("ArraySlice fail")
		return
	}
}

func TestArraySlicePanicSize(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var sli []int
	KArr.ArraySlice(sli, 0, 0)
}

func TestArraySlicePanicType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArraySlice("hello", 0, 2)
}

func BenchmarkArraySlice(b *testing.B) {
	b.ResetTimer()
	var arr = [6]string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < b.N; i++ {
		KArr.ArraySlice(arr, 1, 4)
	}
}

func TestArrayRand(t *testing.T) {
	var arr = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	var sli []int

	res1 := KArr.ArrayRand(sli, 1)
	res2 := KArr.ArrayRand(arr, 3)
	res3 := KArr.ArrayRand(arr, 9)

	if len(res1) != 0 || len(res2) != 3 || len(res3) != 8 {
		t.Error("ArraySlice fail")
		return
	}
}

func TestArrayRandPanicNum(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	var sli []int
	KArr.ArrayRand(sli, 0)
}

func TestArrayRandPanicType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArrayRand("hello", 2)
}

func BenchmarkArrayRand(b *testing.B) {
	b.ResetTimer()
	var arr = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayRand(arr, 6)
	}
}
