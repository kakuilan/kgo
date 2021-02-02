package kgo

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray_ArrayChunk(t *testing.T) {
	size := 3
	res := KArr.ArrayChunk(ssSingle, size)
	assert.Equal(t, 4, len(res))

	item := res[0]
	assert.Equal(t, size, len(item))

	KArr.ArrayChunk([]int{}, 1)
}

func TestArray_ArrayChunk_PanicSize(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, "[ArrayChunk]`size cannot be less than 1", r)
	}()
	KArr.ArrayChunk(ssSingle, 0)
}

func TestArray_ArrayChunk_PanicType(t *testing.T) {
	defer func() {
		r := recover()
		assert.Equal(t, "[ArrayChunk]`arr type must be array or slice", r)
	}()
	KArr.ArrayChunk("hello", 2)
}

func BenchmarkArray_ArrayChunk(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayChunk(ssSingle, 3)
	}
}

func TestArray_ArrayColumn_Struct(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayColumn]`arr type must be")
	}()

	var p1, p2, p3, p4 sPerson
	gofakeit.Struct(&p1)
	gofakeit.Struct(&p2)
	gofakeit.Struct(&p3)
	gofakeit.Struct(&p4)

	var ps = make(sPersons, 4)
	var org = new(sOrganization)

	ps = append(ps, p1, p2, p3, p4)

	org.Leader = p1
	org.Assistant = p2
	org.Member = p3
	org.Substitute = p4

	var res []interface{}

	res = KArr.ArrayColumn(ps, "Name")
	assert.NotEmpty(t, res)

	res = KArr.ArrayColumn(*org, "Age")
	assert.NotEmpty(t, res)

	res = KArr.ArrayColumn(*org, "age")
	assert.Empty(t, res)

	// type err
	KArr.ArrayColumn(org, "Age")
}

func TestArray_ArrayColumn_Map(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[GetFieldValue]`arr type must be")
	}()

	var arr map[string]interface{}
	var res []interface{}

	_ = KStr.JsonDecode([]byte(personsJson), &arr)

	res = KArr.ArrayColumn(arr, "name")
	assert.NotEmpty(t, res)

	res = KArr.ArrayColumn(arr, "Name")
	assert.Empty(t, res)

	arr["person5"] = "hello"
	KArr.ArrayColumn(arr, "name")
}

func BenchmarkArray_ArrayColumn(b *testing.B) {
	var p1, p2, p3, p4 sPerson
	var ps = make(sPersons, 4)
	gofakeit.Struct(&p1)
	gofakeit.Struct(&p2)
	gofakeit.Struct(&p3)
	gofakeit.Struct(&p4)
	ps = append(ps, p1, p2, p3, p4)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = KArr.ArrayColumn(ps, "Name")
	}
}

func TestArray_SlicePush_SlicePop(t *testing.T) {
	var ss []interface{}
	var item interface{}

	lenght := KArr.SlicePush(&ss, slItf...)
	assert.Greater(t, lenght, 1)

	for i := 0; i < lenght; i++ {
		item = KArr.SlicePop(&ss)
		assert.NotEmpty(t, item)
	}
}

func BenchmarkArray_SlicePush(b *testing.B) {
	var ss []interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss = nil
		KArr.SlicePush(&ss, slItf...)
	}
}

func BenchmarkArray_SlicePop(b *testing.B) {
	var ss [][]interface{}
	var sub []interface{}
	for j := 0; j < 10000000; j++ {
		sub = nil
		copy(sub, slItf)
		ss = append(ss, sub)
	}

	b.ResetTimer()
	for _, item := range ss {
		for i := 0; i < len(item); i++ {
			KArr.SlicePop(&item)
		}
	}
}

func TestArray_SliceUnshift_SliceShift(t *testing.T) {
	var ss []interface{}
	var item interface{}
	lenght := KArr.SliceUnshift(&ss, slItf...)
	assert.Greater(t, lenght, 1)

	for i := 0; i < lenght; i++ {
		item = KArr.SliceShift(&ss)
		assert.NotEmpty(t, item)
	}
}

func BenchmarkArray_SliceUnshift(b *testing.B) {
	var ss []interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ss = nil
		KArr.SliceUnshift(&ss, slItf...)
	}
}

func BenchmarkArray_SliceShift(b *testing.B) {
	var ss [][]interface{}
	var sub []interface{}
	for j := 0; j < 10000000; j++ {
		sub = nil
		copy(sub, slItf)
		ss = append(ss, sub)
	}

	b.ResetTimer()
	for _, item := range ss {
		for i := 0; i < len(item); i++ {
			KArr.SliceShift(&item)
		}
	}
}

func TestArray_ArrayKeyExists(t *testing.T) {
	defer func() {
		r := recover()
		assert.Contains(t, r, "[ArrayKeyExists]`arr type must be")
	}()

	chk1 := KArr.ArrayKeyExists(len(naturalArr)-1, naturalArr)
	assert.True(t, chk1)

	chk2 := KArr.ArrayKeyExists(len(slItf)-1, slItf)
	assert.True(t, chk2)

	var person sPerson
	gofakeit.Struct(&person)
	chk3 := KArr.ArrayKeyExists("Name", person)
	chk4 := KArr.ArrayKeyExists("name", person)
	assert.True(t, chk3)
	assert.False(t, chk4)

	var persons map[string]interface{}
	_ = KStr.JsonDecode([]byte(personsJson), &persons)
	chk5 := KArr.ArrayKeyExists("person1", persons)
	chk6 := KArr.ArrayKeyExists("Age", persons)
	assert.True(t, chk5)
	assert.False(t, chk6)

	var key interface{}
	chk7 := KArr.ArrayKeyExists(key, persons)
	assert.False(t, chk7)

	KArr.ArrayKeyExists(1, nil)
}

func BenchmarkArray_ArrayKeyExists_Slice(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists(1, naturalArr)
	}
}

func BenchmarkArray_ArrayKeyExists_Struct(b *testing.B) {
	var person sPerson
	gofakeit.Struct(&person)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists("Name", person)
	}
}

func BenchmarkArray_ArrayKeyExists_Map(b *testing.B) {
	var persons map[string]interface{}
	_ = KStr.JsonDecode([]byte(personsJson), &persons)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KArr.ArrayKeyExists("person1", persons)
	}
}
