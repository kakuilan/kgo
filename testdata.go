//测试数据

package kgo

import "github.com/brianvoe/gofakeit/v6"

//类型-人员
type sPerson struct {
	secret string `json:"secret"`
	Name   string `fake:"{name}" json:"name"`
	Addr   string `fake:"{city}" json:"city"`
	Age    int    `fake:"{number:1,99}" json:"age"`
	Gender bool   `fake:"{bool}" json:"gender"`
	other  int    `json:"other"`
	none   bool
}

//类型-人群
type sPersons []sPerson

//类型-组织
type sOrganization struct {
	Leader     sPerson  //领导
	Assistant  sPerson  //副手
	Substitute sPerson  //候补
	Members    sPersons //成员
}

//结构体-人员
var personS1, personS2, personS3, personS4, personS5 sPerson

//结构体-人群
var crowd = make(sPersons, 1)

//结构体-组织
var orgS1 = new(sOrganization)

//字典-普通人员
var personMp1 = map[string]interface{}{"age": 20, "name": "test1", "naction": "us", "tel": "13712345678"}
var personMp2 = map[string]interface{}{"age": 21, "name": "test2", "naction": "cn", "tel": "13712345679"}
var personMp3 = map[string]interface{}{"age": 22, "name": "test3", "naction": "en", "tel": "13712345670"}
var personMp4 = map[string]interface{}{"age": 23, "name": "test4", "naction": "fr", "tel": "13712345671"}
var personMp5 = map[string]interface{}{"age": 21, "name": "test5", "naction": "cn", "tel": "13712345672"}
var personMps = []interface{}{personMp1, personMp2, personMp3, personMp4, personMp5}

//字典-结构体人员
var perStuMps map[string]sPerson

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
var strMp2 = map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": "", "2": "cc"}
var strMpEmp = make(map[string]string)
var colorMp = map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red", "c": "yellow"}

//字符串切片
var strSl1 = []string{"aa", "bb", "cc", "dd", "ee", "", "hh", "ii"}
var strSl2 = []string{"bb", "cc", "ff", "gg", "ee", "", "gg"}
var strSlEmp = []string{}

//字符串
var strHello = "Hello World!"

func init() {
	gofakeit.Struct(&personS1)
	gofakeit.Struct(&personS2)
	gofakeit.Struct(&personS3)
	gofakeit.Struct(&personS4)
	gofakeit.Struct(&personS5)

	crowd = append(crowd, personS1, personS2, personS3, personS4, personS5)

	orgS1.Leader = personS1
	orgS1.Assistant = personS2
	orgS1.Substitute = personS3
	orgS1.Members = sPersons{personS4, personS5}

	perStuMps = map[string]sPerson{"a": personS1, "b": personS2, "c": personS3, "d": personS4, "e": personS5}
}
