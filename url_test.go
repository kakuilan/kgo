package kgo

import (
	"testing"
)

func TestParseStr(t *testing.T) {
	str1 := `first=value&arr[]=foo+bar&arr[]=baz`
	str2 := `f1=m&f2=n`
	str3 := `f[a]=m&f[b]=n`
	str4 := `f[a][a]=m&f[a][b]=n`
	str5 := `f[]=m&f[]=n`
	str6 := `f[a][]=m&f[a][]=n`
	str7 := `f[][]=m&f[][]=n`
	str8 := `f=m&f[a]=n`
	str9 := `a .[[b=c`
	arr1 := make(map[string]interface{})
	arr2 := make(map[string]interface{})
	arr3 := make(map[string]interface{})
	arr4 := make(map[string]interface{})
	arr5 := make(map[string]interface{})
	arr6 := make(map[string]interface{})
	arr7 := make(map[string]interface{})
	arr8 := make(map[string]interface{})
	arr9 := make(map[string]interface{})

	err1 := KUrl.ParseStr(str1, arr1)
	err2 := KUrl.ParseStr(str2, arr2)
	err3 := KUrl.ParseStr(str3, arr3)
	err4 := KUrl.ParseStr(str4, arr4)
	err5 := KUrl.ParseStr(str5, arr5)
	err6 := KUrl.ParseStr(str6, arr6)
	err7 := KUrl.ParseStr(str7, arr7)
	err8 := KUrl.ParseStr(str8, arr8)
	err9 := KUrl.ParseStr(str9, arr9)
	//fmt.Printf("%+v\n", arr1)
	//fmt.Printf("%+v\n", arr2)
	//fmt.Printf("%+v\n", arr3)
	//fmt.Printf("%+v\n", arr4)
	//fmt.Printf("%+v\n", arr5)
	//fmt.Printf("%+v\n", arr6)
	//fmt.Printf("%+v\n", arr7)
	//fmt.Printf("%+v\n", arr9)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || err6 != nil || err7 != nil || err9 != nil {
		t.Error("ParseStr fail")
		return
	} else if err8 == nil {
		t.Error("ParseStr fail")
		return
	}

	err9 = KUrl.ParseStr("f=n&f[a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}
	err9 = KUrl.ParseStr("f=n&f[][a]=m&", arr9)
	if err9 == nil {
		t.Error("ParseStr fail")
		return
	}

	_ = KUrl.ParseStr("?first=value&arr[]=foo+bar&arr[]=baz&arr[][a]=aaa", arr9)
	_ = KUrl.ParseStr("f[][a]=m&f[][b]=h", arr9)
	_ = KUrl.ParseStr("f[][a]=&f[][b]=", arr9)
	_ = KUrl.ParseStr("he&a=1", arr9)
	_ = KUrl.ParseStr("he& =2", arr9)
	_ = KUrl.ParseStr("he& g=2", arr9)
	_ = KUrl.ParseStr("he&=3", arr9)
	_ = KUrl.ParseStr("he&[=4", arr9)
	_ = KUrl.ParseStr("he&]=5", arr9)
	_ = KUrl.ParseStr("f[a].=m&f=n&", arr9)
	_ = KUrl.ParseStr("f=n&f[a][]b=m&", arr9)
	err8 = KUrl.ParseStr("f=n&f[a][]=m&", arr9)
	println("err8 111", err8.Error())

	err8 = KUrl.ParseStr("%=%gg&b=4", arr9)  //key nvalid URL escape "%"
	err9 = KUrl.ParseStr("he&e=%&b=4", arr9) //value nvalid URL escape "%"
	if err8 == nil || err9 == nil {
		t.Error("ParseStr fail")
		return
	}

}

func BenchmarkParseStr(b *testing.B) {
	b.ResetTimer()
	str := `first=value&arr[]=foo+bar&arr[]=baz`
	arr := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		_ = KUrl.ParseStr(str, arr)
	}
}
