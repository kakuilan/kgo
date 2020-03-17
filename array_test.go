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

func TestInIntSlice(t *testing.T) {
	tests := []struct {
		list     []int
		find     int
		expected bool
	}{
		{[]int{42}, 42, true},
		{[]int{42}, 4, false},
		{[]int{42, 666, 324523}, 666, true},
		{[]int{42, 796, 141259}, 0, false},
		{[]int{}, 0, false},
	}

	for _, test := range tests {
		actual := KArr.InIntSlice(test.find, test.list)
		if actual != test.expected {
			t.Errorf("Expected InIntSlice(%d, %v) to be %v, got %v", test.find, test.list, test.expected, actual)
			return
		}
	}
}

func BenchmarkInIntSlice(b *testing.B) {
	b.ResetTimer()
	arr := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		KArr.InIntSlice(7, arr)
	}
}

func TestInInt64Slice(t *testing.T) {
	tests := []struct {
		list     []int64
		find     int64
		expected bool
	}{
		{[]int64{42}, 42, true},
		{[]int64{42}, 4, false},
		{[]int64{42, 666, 324523}, 666, true},
		{[]int64{42, 796, 141259}, 0, false},
		{[]int64{}, 0, false},
	}

	for _, test := range tests {
		actual := KArr.InInt64Slice(test.find, test.list)
		if actual != test.expected {
			t.Errorf("Expected InInt64Slice(%d, %v) to be %v, got %v", test.find, test.list, test.expected, actual)
			return
		}
	}
}

func BenchmarkInInt64Slice(b *testing.B) {
	b.ResetTimer()
	arr := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		KArr.InInt64Slice(7, arr)
	}
}

func TestInStringSlice(t *testing.T) {
	tests := []struct {
		list     []string
		find     string
		expected bool
	}{
		{[]string{}, "", false},
		{[]string{""}, "", true},
		{[]string{"aa"}, "bb", false},
		{[]string{"aa", "bb", "cc", "hello"}, "ee", false},
		{[]string{"aa", "bb", "cc", "hello"}, "hello", true},
	}

	for _, test := range tests {
		actual := KArr.InStringSlice(test.find, test.list)
		if actual != test.expected {
			t.Errorf("Expected InStringSlice(%s, %v) to be %v, got %v", test.find, test.list, test.expected, actual)
			return
		}
	}
}

func BenchmarkInStringSlice(b *testing.B) {
	b.ResetTimer()
	arr := []string{"aa", "bb", "cc", "hello"}
	for i := 0; i < b.N; i++ {
		KArr.InStringSlice("hello", arr)
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
	res := KArr.ArrayValues(mp, false)
	if len(res) != 3 {
		t.Error("ArrayValues fail")
		return
	}

	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	res = KArr.ArrayValues(sli, false)
	if len(res) != 5 {
		t.Error("ArrayValues fail")
		return
	}

	KArr.ArrayValues("hello", false)
}

func BenchmarkArrayValues(b *testing.B) {
	b.ResetTimer()
	mp := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < b.N; i++ {
		KArr.ArrayValues(mp, false)
	}
}

func TestMergeSlice(t *testing.T) {
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

	res1 := KArr.MergeSlice(false, arr, sli)
	if len(res1) != 15 {
		t.Error("MergeSlice fail")
		return
	}

	res2 := KArr.MergeSlice(true, arr, sli)
	if len(res2) != 13 {
		t.Error("MergeSlice fail")
		return
	}
	KArr.MergeSlice(true)
	KArr.MergeSlice(false, "hellow")
}

func BenchmarkMergeSlice(b *testing.B) {
	b.ResetTimer()
	var arr = [10]int{1, 2, 3, 4, 5, 6}
	var sli []string = make([]string, 5)
	sli[0] = "aaa"
	sli[2] = "ccc"
	sli[3] = "ddd"
	for i := 0; i < b.N; i++ {
		KArr.MergeSlice(false, arr, sli)
	}
}

func TestMergeMap(t *testing.T) {
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

	res := KArr.MergeMap(true, mp1, mp2)
	_, err := KStr.JsonEncode(res)
	if err != nil {
		t.Error("MergeMap fail")
		return
	}
	KArr.MergeMap(false)
	KArr.MergeMap(false, mp1, mp2)
	KArr.MergeMap(false, mp1, mp2, "hello")
}

func BenchmarkMergeMap(b *testing.B) {
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
		KArr.MergeMap(true, mp1, mp2)
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
		t.Error("Implode slice fail")
		return
	}

	//字典
	var mp1 = make(map[string]string)
	res = KArr.Implode(",", mp1)
	if res != "" {
		t.Error("Implode map fail")
		return
	}

	mp2 := map[string]string{
		"a": "aa",
		"b": "bb",
		"c": "cc",
		"d": "dd",
	}
	res = KArr.Implode(",", mp2)
	if res == "" {
		t.Error("Implode map fail")
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

func TestJoinStrings(t *testing.T) {
	var arr = []string{}

	res := KArr.JoinStrings(arr, ",")
	if res != "" {
		t.Error("JoinStrings fail")
		return
	}

	arr = append(arr, "a", "b", "c", "d", "e")
	KArr.JoinStrings(arr, ",")
}

func BenchmarkJoinStrings(b *testing.B) {
	b.ResetTimer()
	var arr = []string{"a", "b", "c", "d", "e"}
	for i := 0; i < b.N; i++ {
		KArr.JoinStrings(arr, ",")
	}
}

func TestJoinJoinInts(t *testing.T) {
	var arr = []int{}

	res := KArr.JoinInts(arr, ",")
	if res != "" {
		t.Error("JoinStrings fail")
		return
	}

	arr = append(arr, 1, 2, 3, 4, 5, 6)
	KArr.JoinInts(arr, ",")
}

func BenchmarkJoinInts(b *testing.B) {
	b.ResetTimer()
	var arr = []int{1, 2, 3, 4, 5, 6}
	for i := 0; i < b.N; i++ {
		KArr.JoinInts(arr, ",")
	}
}

func TestArrayDiff(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	ar1 := []string{"aa", "bb", "cc", "dd", ""}
	ar2 := []string{"bb", "cc", "ff", "gg", ""}
	mp1 := map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": ""}
	mp2 := map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": ""}

	var ar3 []string
	var mp3 = make(map[string]string)

	res1 := KArr.ArrayDiff(ar1, ar2)
	res2 := KArr.ArrayDiff(mp1, mp2)
	if len(res1) != len(res2) {
		t.Error("ArrayDiff fail")
		return
	}

	res5 := KArr.ArrayDiff(ar3, ar1)
	res6 := KArr.ArrayDiff(ar1, ar3)
	if len(res5) != 0 || len(res6) != 4 {
		t.Error("ArrayDiff fail")
		return
	}

	res7 := KArr.ArrayDiff(mp3, mp1)
	res8 := KArr.ArrayDiff(mp1, mp3)
	if len(res7) != 0 || len(res8) != 4 {
		t.Error("ArrayDiff fail")
		return
	}

	res9 := KArr.ArrayDiff(ar3, mp1)
	res10 := KArr.ArrayDiff(ar1, mp3)
	res11 := KArr.ArrayDiff(ar1, mp1)
	if len(res9) != 0 || len(res10) != len(res11) {
		t.Error("ArrayDiff fail")
		return
	}

	res12 := KArr.ArrayDiff(mp3, ar1)
	res13 := KArr.ArrayDiff(mp1, ar3)
	res14 := KArr.ArrayDiff(mp1, ar1)
	if len(res12) != 0 || len(res13) != len(res14) {
		t.Error("ArrayDiff fail")
		return
	}

	KArr.ArrayDiff("hello", ar1)
}

func BenchmarkArrayDiff(b *testing.B) {
	b.ResetTimer()
	ar1 := []string{"aa", "bb", "cc", "dd", ""}
	ar2 := []string{"bb", "cc", "ff", "gg", ""}
	for i := 0; i < b.N; i++ {
		KArr.ArrayDiff(ar1, ar2)
	}
}

func TestArrayUnique(t *testing.T) {
	arr1 := map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red"}
	arr2 := []string{"aa", "bb", "cc", "", "bb", "aa"}
	res1 := KArr.ArrayUnique(arr1)
	res2 := KArr.ArrayUnique(arr2)
	if len(res1) == 0 || len(res2) == 0 {
		t.Error("ArrayUnique fail")
		return
	}
}

func TestArrayUniquePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	_ = KArr.ArrayUnique("hello")
}

func BenchmarkArrayUnique(b *testing.B) {
	b.ResetTimer()
	arr := map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red"}
	for i := 0; i < b.N; i++ {
		KArr.ArrayUnique(arr)
	}
}

func TestArraySearchItemNMutil(t *testing.T) {
	type titem map[string]interface{}

	var list []interface{}
	arrs := make(map[string]interface{})
	cond := make(map[string]interface{})

	res1 := KArr.ArraySearchItem(list, cond)
	res2 := KArr.ArraySearchItem(arrs, cond)
	if res1 != nil || res2 != nil {
		t.Error("ArraySearchItem fail")
	}

	mul1 := KArr.ArraySearchMutil(list, cond)
	mul2 := KArr.ArraySearchMutil(arrs, cond)
	if mul1 != nil || mul2 != nil {
		t.Error("ArraySearchMutil fail")
	}

	item1 := titem{"age": 20, "name": "test1", "naction": "us", "tel": "13712345678"}
	item2 := titem{"age": 21, "name": "test2", "naction": "cn", "tel": "13712345679"}
	item3 := titem{"age": 22, "name": "test3", "naction": "en", "tel": "13712345670"}
	item4 := titem{"age": 23, "name": "test4", "naction": "fr", "tel": "13712345671"}
	item5 := titem{"age": 21, "name": "test5", "naction": "cn", "tel": "13712345672"}

	list = append(list, item1, item2, item3, item4, item5, nil, "hello")
	arrs["a"] = item1
	arrs["b"] = item2
	arrs["c"] = item3
	arrs["c"] = item4
	arrs["d"] = nil
	arrs["d"] = "world"
	arrs["e"] = item5

	cond1 := map[string]interface{}{"age": 23}
	cond2 := map[string]interface{}{"age": 21, "naction": "cn"}
	cond3 := map[string]interface{}{"age": 22, "naction": "cn", "tel": "13712345671"}

	res3 := KArr.ArraySearchItem(list, cond1)
	res4 := KArr.ArraySearchItem(arrs, cond1)
	if res3 == nil || res4 == nil {
		t.Error("ArraySearchItem fail")
	}

	mul3 := KArr.ArraySearchMutil(list, cond1)
	mul4 := KArr.ArraySearchMutil(arrs, cond1)
	if mul3 == nil || mul4 == nil {
		t.Error("ArraySearchMutil fail")
	}

	res5 := KArr.ArraySearchItem(list, cond2)
	res6 := KArr.ArraySearchItem(arrs, cond2)
	if res5 == nil || res6 == nil {
		t.Error("ArraySearchItem fail")
	}

	mul5 := KArr.ArraySearchMutil(list, cond2)
	mul6 := KArr.ArraySearchMutil(arrs, cond2)
	if mul5 == nil || mul6 == nil {
		t.Error("ArraySearchMutil fail")
	}

	res7 := KArr.ArraySearchItem(list, cond3)
	res8 := KArr.ArraySearchItem(arrs, cond3)
	if res7 != nil || res8 != nil {
		t.Error("ArraySearchItem fail")
	}

	mul7 := KArr.ArraySearchMutil(list, cond3)
	mul8 := KArr.ArraySearchMutil(arrs, cond3)
	if mul7 != nil || mul8 != nil {
		t.Error("ArraySearchMutil fail")
	}
}

func TestArraySearchItemPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArraySearchItem("hello", map[string]interface{}{"a": 1})
}

func TestArraySearchMutilPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	KArr.ArraySearchMutil("hello", map[string]interface{}{"a": 1})
}

func BenchmarkArraySearchItem(b *testing.B) {
	b.ResetTimer()
	type titem map[string]interface{}
	var list []interface{}

	item1 := titem{"age": 20, "name": "test1", "naction": "us", "tel": "13712345678"}
	item2 := titem{"age": 21, "name": "test2", "naction": "cn", "tel": "13712345679"}
	item3 := titem{"age": 22, "name": "test3", "naction": "en", "tel": "13712345670"}
	item4 := titem{"age": 23, "name": "test4", "naction": "fr", "tel": "13712345671"}
	list = append(list, item1, item2, item3, item4, nil, "hello")
	cond := map[string]interface{}{"age": 21, "naction": "cn"}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchItem(list, cond)
	}
}

func BenchmarkArraySearchMutil(b *testing.B) {
	b.ResetTimer()
	type titem map[string]interface{}
	var list []interface{}

	item1 := titem{"age": 20, "name": "test1", "naction": "us", "tel": "13712345678"}
	item2 := titem{"age": 21, "name": "test2", "naction": "cn", "tel": "13712345679"}
	item3 := titem{"age": 22, "name": "test3", "naction": "en", "tel": "13712345670"}
	item4 := titem{"age": 23, "name": "test4", "naction": "fr", "tel": "13712345671"}
	item5 := titem{"age": 21, "name": "test5", "naction": "cn", "tel": "13712345672"}
	list = append(list, item1, item2, item3, item4, nil, "hello", item5)
	cond := map[string]interface{}{"age": 21, "naction": "cn"}
	for i := 0; i < b.N; i++ {
		KArr.ArraySearchMutil(list, cond)
	}
}

func TestUniqueInts(t *testing.T) {
	arr := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	res := KArr.UniqueInts(arr)
	if len(arr) == len(res) {
		t.Error("UniqueInts fail")
		return
	}
}

func BenchmarkUniqueInts(b *testing.B) {
	b.ResetTimer()
	arr := []int{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		KArr.UniqueInts(arr)
	}
}

func TestUnique64Ints(t *testing.T) {
	arr := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	res := KArr.Unique64Ints(arr)
	if len(arr) == len(res) {
		t.Error("Unique64Ints fail")
		return
	}
}

func BenchmarkUnique64Ints(b *testing.B) {
	b.ResetTimer()
	arr := []int64{-3, 9, -5, 0, 5, -3, 0, 7}
	for i := 0; i < b.N; i++ {
		KArr.Unique64Ints(arr)
	}
}

func TestUniqueStrings(t *testing.T) {
	arr1 := []string{}
	res1 := KArr.UniqueStrings(arr1)
	if len(res1) != 0 {
		t.Error("UniqueStrings fail")
		return
	}

	arr2 := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	res2 := KArr.UniqueStrings(arr2)
	if len(arr2) == len(res2) {
		t.Error("UniqueStrings fail")
		return
	}
}

func BenchmarkUniqueStrings(b *testing.B) {
	b.ResetTimer()
	arr := []string{"", "hello", "world", "hello", "你好", "world", "1234"}
	for i := 0; i < b.N; i++ {
		KArr.UniqueStrings(arr)
	}
}

func TestIsEqualArray(t *testing.T) {
	s1 := []string{"a", "b"}
	s2 := []string{"b", "a"}
	chk1 := KArr.IsEqualArray(s1, s2)
	if !chk1 {
		t.Error("IsEqualArray fail")
		return
	}

	s3 := [2]int{3, 8}
	s4 := []int{8, 3}
	chk2 := KArr.IsEqualArray(s3, s4)
	if !chk2 {
		t.Error("IsEqualArray fail")
		return
	}

	s5 := []string{"3", "8"}
	chk3 := KArr.IsEqualArray(s3, s5)
	if chk3 {
		t.Error("IsEqualArray fail")
		return
	}

	type TestType struct {
		StrKey   string
		IntSlice []int
	}
	s6 := []TestType{
		{StrKey: "key1", IntSlice: []int{1, 2}},
		{StrKey: "key2", IntSlice: []int{3, 4, 5}},
	}
	s7 := []TestType{
		{StrKey: "key2", IntSlice: []int{3, 4, 5}},
		{StrKey: "key1", IntSlice: []int{1, 2}},
	}
	s8 := []TestType{
		{StrKey: "key2", IntSlice: []int{3, 4, 5}},
		{StrKey: "key3", IntSlice: []int{6, 7}},
	}
	chk4 := KArr.IsEqualArray(s6, s7)
	if !chk4 {
		t.Error("IsEqualArray fail")
		return
	}
	chk5 := KArr.IsEqualArray(s7, s8)
	if chk5 {
		t.Error("IsEqualArray fail")
		return
	}
}

func TestIsEqualArrayPanicExpected(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	s := []string{"a", "b"}
	KArr.IsEqualArray("hello", s)
}

func TestIsEqualArrayPanicActual(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover...:", r)
		}
	}()

	s := []string{"a", "b"}
	KArr.IsEqualArray(s, "hello")
}

func BenchmarkIsEqualArray(b *testing.B) {
	b.ResetTimer()
	s1 := []string{"a", "b"}
	s2 := []string{"b", "a"}
	for i := 0; i < b.N; i++ {
		KArr.IsEqualArray(s1, s2)
	}
}

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

func TestIsMap(t *testing.T) {
	mp := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	res1 := KArr.IsMap(mp)
	res2 := KArr.IsMap(123)
	if !res1 || res2 {
		t.Error("IsMap fail")
		return
	}
}

func BenchmarkIsMap(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.IsMap("hello")
	}
}

func TestGetBiosBoardCpuInfo(t *testing.T) {
	res1 := KOS.GetBiosInfo()
	res2 := KOS.GetBoardInfo()
	res3 := KOS.GetCpuInfo()

	//fmt.Printf("%+v\n", res1)

	if res1.Version == "" {
		t.Error("GetBiosInfo fail")
		return
	}

	if res2.Version == "" {
		t.Error("GetBoardInfo fail")
		return
	}

	if res3.Model == "" {
		t.Error("GetCpuInfo fail")
		return
	}

}

func BenchmarkGetBiosInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBiosInfo()
	}
}

func BenchmarkGetBoardInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetBoardInfo()
	}
}

func BenchmarkGetCpuInfo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KOS.GetCpuInfo()
	}
}
