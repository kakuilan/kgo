package kgo

//测试数据

//类型-人员
type sPerson struct {
	Name   string `fake:"{name}"`
	Addr   string `fake:"{city}"`
	Age    int    `fake:"{number:1,99}"`
	Gender bool   `fake:"{bool}"`
}

//类型-人群
type sPersons []sPerson

//类型-组织
type sOrganization struct {
	Leader     sPerson //领导
	Assistant  sPerson //副手
	Member     sPerson //成员
	Substitute sPerson //候补
}

//自然数数组
var naturalArr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

//整数切片
var intSlc = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}
var int64Slc = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}

//单字符切片
var ssSingle = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

//接口切片
var slItf = []interface{}{1, 2, 3, 3.14, 6.67428, true, 'a', 'b', 'c', "hello", "你好"}

//persons JSON串
var personsJson = `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`

//字符串map
var strMp1 = map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "", "2": "cc", "3": "no"}

//var strMp2 = map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": "", "2": "cc"}
