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

//单字符切片
var ssSingle = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

