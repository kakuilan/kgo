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

func TestArrayColumn(t *testing.T) {
	//数组切片
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wang5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr interface{}
	err := KStr.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode fail")
		return
	}

	res := KArr.ArrayColumn(arr, "name")
	if len(res) != 4 {
		t.Error("ArrayColumn fail")
		return
	}

	//字典
	jsonStr = `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
	err = KStr.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode fail")
		return
	}

	res = KArr.ArrayColumn(arr, "name")
	if len(res) != 4 {
		t.Error("ArrayColumn fail")
		return
	}
}

func TestArrayColumnPanicSlice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	//数组切片
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wang5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr []interface{}
	err := KStr.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode fail")
		return
	}

	arr = append(arr, "hello")
	KArr.ArrayColumn(arr, "name")
}

func TestArrayColumnPanicMap(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	//数组切片
	jsonStr := `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
	var arr map[string]interface{}
	err := KStr.JsonDecode([]byte(jsonStr), &arr)
	if err != nil {
		t.Error("JsonDecode fail")
		return
	}

	arr["person5"] = "hello"
	KArr.ArrayColumn(arr, "name")
}

func TestArrayColumnPanicType(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArrayColumn("hello", "name")
}

func BenchmarkArrayColumn(b *testing.B) {
	b.ResetTimer()
	jsonStr := `[{"name":"zhang3","age":23,"sex":1},{"name":"li4","age":30,"sex":1},{"name":"wang5","age":25,"sex":0},{"name":"zhao6","age":50,"sex":0}]`
	var arr interface{}
	_ = KStr.JsonDecode([]byte(jsonStr), &arr)
	for i := 0; i < b.N; i++ {
		KArr.ArrayColumn(arr, "name")
	}
}

func TestArrayPushPop(t *testing.T) {
	var arr []interface{}
	length := KArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")
	if length != 6 {
		t.Error("ArrayPush fail")
		return
	}

	last := KArr.ArrayPop(&arr)
	if fmt.Sprintf("%v", last) != "c" {
		t.Error("ArrayPop fail")
		return
	}
	arr = nil
	KArr.ArrayPop(&arr)
}

func BenchmarkArrayPush(b *testing.B) {
	b.ResetTimer()
	var arr []interface{}
	for i := 0; i < b.N; i++ {
		KArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")
	}
}

func BenchmarkArrayPop(b *testing.B) {
	b.ResetTimer()
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayPop(&arr)
	}
}

func TestArrayShiftUnshift(t *testing.T) {
	var arr []interface{}
	length := KArr.ArrayUnshift(&arr, 1, 2, 3, "a", "b", "c")
	if length != 6 {
		t.Error("ArrayUnshift fail")
		return
	}

	first := KArr.ArrayShift(&arr)
	if fmt.Sprintf("%v", first) != "1" {
		t.Error("ArrayPop fail")
		return
	}
	arr = nil
	KArr.ArrayShift(&arr)
}

func BenchmarkArrayUnshift(b *testing.B) {
	b.ResetTimer()
	var arr []interface{}
	for i := 0; i < b.N; i++ {
		KArr.ArrayUnshift(&arr, 1, 2, 3, "a", "b", "c")
	}
}

func BenchmarkArrayShift(b *testing.B) {
	b.ResetTimer()
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayShift(&arr)
	}
}

func TestArrayKeyExistsArr(t *testing.T) {
	var arr []interface{}
	KArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")

	chk1 := KArr.ArrayKeyExists(1, arr)
	chk2 := KArr.ArrayKeyExists(-1, arr)
	chk3 := KArr.ArrayKeyExists(6, arr)
	if !chk1 || chk2 || chk3 {
		t.Error("ArrayKeyExists fail")
		return
	}

	var key interface{}
	KArr.ArrayKeyExists(key, arr)

	arr = nil
	KArr.ArrayKeyExists(1, arr)
}

func TestArrayKeyExistsMap(t *testing.T) {
	jsonStr := `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
	var arr map[string]interface{}
	_ = KStr.JsonDecode([]byte(jsonStr), &arr)

	chk1 := KArr.ArrayKeyExists("person2", arr)
	chk2 := KArr.ArrayKeyExists("test", arr)
	if !chk1 || chk2 {
		t.Error("ArrayKeyExists fail")
		return
	}
}

func TestArrayKeyExistsPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArrayKeyExists("abc", "hello")
}

func BenchmarkArrayKeyExistsArr(b *testing.B) {
	b.ResetTimer()
	var arr []interface{}
	KArr.ArrayPush(&arr, 1, 2, 3, "a", "b", "c")
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists(3, arr)
	}
}

func BenchmarkArrayKeyExistsMap(b *testing.B) {
	b.ResetTimer()
	jsonStr := `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
	var arr map[string]interface{}
	_ = KStr.JsonDecode([]byte(jsonStr), &arr)
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists("person2", arr)
	}
}

func TestArrayReverse(t *testing.T) {
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	res := KArr.ArrayReverse(arr)

	if len(res) != 5 || fmt.Sprintf("%s", res[2]) != "c" {
		t.Error("ArrayReverse fail")
		return
	}

	var myslice []int
	KArr.ArrayReverse(myslice)
}

func TestArrayReversePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArrayReverse("hello")
}

func BenchmarkArrayReverse(b *testing.B) {
	b.ResetTimer()
	var arr = []interface{}{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayReverse(arr)
	}
}

func TestImplode(t *testing.T) {
	var arr = []string{"a", "b", "c", "d", "e"}
	res := KArr.Implode(",", arr)

	arr = nil
	res = KArr.Implode(",", arr)
	if res != "" {
		t.Error("Implode fail")
		return
	}
}

func TestImplodePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.Implode(",", "hello")
}

func BenchmarkImplode(b *testing.B) {
	b.ResetTimer()
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < b.N; i++ {
		KArr.Implode(",", sli)
	}
}

func TestRand(t *testing.T) {
	min := 1
	max := 66666
	res := KNum.Rand(min, max)

	if res < min || res > max {
		t.Error("Rand fail")
		return
	}
}
