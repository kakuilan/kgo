//测试数据

package kgo

import "github.com/brianvoe/gofakeit/v6"

//类型-接口
type itfType interface {
	noRealize() //该方法不实现
	sayHello(name string) string
}

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

//接口对象
var itfObj itfType

//结构体-人员
var personS1, personS2, personS3, personS4, personS5 sPerson

//结构体-人群
var crowd sPersons

//结构体-组织
var orgS1 = new(sOrganization) //返回指针

//字典-普通人员
var personMp1 = map[string]interface{}{"age": 20, "name": "test1", "naction": "us", "tel": "13712345678"}
var personMp2 = map[string]interface{}{"age": 21, "name": "test2", "naction": "cn", "tel": "13712345679"}
var personMp3 = map[string]interface{}{"age": 22, "name": "test3", "naction": "en", "tel": "13712345670"}
var personMp4 = map[string]interface{}{"age": 23, "name": "test4", "naction": "fr", "tel": "13712345671"}
var personMp5 = map[string]interface{}{"age": 21, "name": "test5", "naction": "cn", "tel": "13712345672"}
var personMps = []interface{}{personMp1, personMp2, personMp3, personMp4, personMp5}

//字典-结构体人员
var perStuMps map[string]sPerson

//类型-圆周率
type fPi32 float32
type fPi64 float64

var flPi1 float32 = 3.141592456
var flPi2 float64 = 3.141592456
var flPi3 fPi32 = 3.141592456
var flPi4 fPi64 = 3.141592456
var bytPi5 = []byte{229, 10, 191, 57, 251, 33, 9, 64} //flPi2的字节切片
var strPi6 = "3.141592456"

//数值
var intSpeedLight int = 299792458            //光速
var intAstronomicalUnit int64 = 149597870660 //天文单位
var floSpeedLight float32 = 2.99792458
var bytAstronomicalUnit = []byte{0, 0, 0, 34, 212, 186, 90, 68} //intAstronomicalUnit的字节切片

var floAvogadro float64 = 6.02214129e23   // 阿伏伽德罗常数
var floPlanck float64 = 6.62606957e-34    // 普朗克常数
var floGravitional float64 = 6.673e-11    //重力常数
var floPermittivity float64 = 8.85419e-12 //真空介电常数

//复数
var cmplNum1 = complex(1, 2)
var cmplNum2 = complex(3, 4)

//字符串
var strHello = "Hello World! 你好！"
var b64Hello = "SGVsbG8gV29ybGQhIOS9oOWlve+8gQ=="
var strHelloHex = "48656c6c6f20576f726c642120e4bda0e5a5bdefbc81" //strHello的16进制
var strSpeedLight = "299792458"
var binAstronomicalUnit = "10001011010100101110100101101001000100" //intAstronomicalUnit的二进制
var hexAstronomicalUnit = "22d4ba5a44"                             //intAstronomicalUnit的16进制
var otcAstronomicalUnit = "2132456455104"                          //intAstronomicalUnit的8进制
var similarStr1 = "We love China,how are you?"
var similarStr2 = "Tom love you,he come from China."
var str2Code = "https://tool.google.com.net/encrypt?type=4Hello World! 你好！"
var b64UrlCode = "aHR0cHM6Ly90b29sLmdvb2dsZS5jb20ubmV0L2VuY3J5cHQ_dHlwZT00SGVsbG8gV29ybGQhIOS9oOWlve-8gQ"
var esyenCode = "23da39b4epjQzaJZuaPW0piFWEbvA0cJISjztw"
var strNoGbk = "月日は百代の過客にして、行かふ年も又旅人也。안녕.ＡＢＣＤＥＦＧＨＩＪＫ"

//当前时间
var nowNanoInt = Kuptime.UnixNano()
var nowNanoStr = toStr(Kuptime.UnixNano())

//IP
var noneIp = "0.0.0.0"
var localIp = "127.0.0.1"
var localIpInt uint32 = 2130706433
var lanIp = "192.168.0.1"
var lanIpInt uint32 = 3232235521
var dockerIp = "172.16.0.1"
var baiduIpv4 = "39.156.69.79"
var googleIpv4 = "172.217.26.142"
var googleIpv6 = "2404:6800:4005:80f::200e"
var publicIp1 = "199.232.96.133"
var publicIp2 = "140.82.114.3"

//自然数数组
var naturalArr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

//整数切片
var intSlc = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}
var int64Slc = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}

//浮点切片
var flo32Slc = []float32{1.23, 0.0, flPi1, floSpeedLight, 6.6260755, 1.60217733}
var flo64Slc = []float64{flPi2, floAvogadro, floPlanck, floGravitional, floPermittivity}

//布尔切片
var booSlc = []bool{true, true, false, true, false, true, true}

//字节切片
var bytsHello = []byte(strHello)
var runesHello = []rune(strHello)
var bytSpeedLight = []byte(strSpeedLight)
var bytsPasswd = []byte("$2a$10$j3WOP6rP2I7skNoxiFdNdOh6OhPxP0Sp3Wmeuekh90oeF3D1EQQBK")
var bytCryptKey = []byte("1234567890123456")

//单字符切片
var ssSingle = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

//persons JSON串
var personsJson = `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`

//字符串map
var strMp1 = map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "", "2": "cc", "3": "no"}
var strMp2 = map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": "", "2": "cc"}
var strMpEmp = make(map[string]string)
var colorMp = map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red", "c": "yellow", "n": ""}

//字符串切片
var strSl1 = []string{"aa", "bb", "cc", "dd", "ee", "", "hh", "ii"}
var strSl2 = []string{"bb", "cc", "ff", "gg", "ee", "", "gg"}
var strSlEmp = []string{}

//接口切片
var slItf = []interface{}{99, 0, 1, 2, 0.0, 3, false, 3.14, 6.67428, true, 'a', "", 'b', nil, 'c', intSpeedLight, "hello", nowNanoInt, strSlEmp, "你好"}

//回调函数
var fnCb1 CallBack
var fnPtr1 = &fnCb1

//rsa相关
//错误的公钥
var rsaPublicErrStr = `-----BEGIN RSA PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDteXRcRyppm5sOVvteo37Dmaid
bx6YrV6QWZ0L9mGfCmSW1a/Ad61kT6OoU0Z3DyId7vA9TtvULucEUpywPpSoP/r+
820UHFihdyhcb1iy8Z3v6KUcarWzUOZpo0mc+o4hW2O1VnzNxLcXmhQOA9NdEOV/
-----END RSA PUBLIC KEY`

//错误的私钥
var rsaPrivateErrStr = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDteXRcRyppm5sOVvteo37Dmaidbx6YrV6QWZ0L9mGfCmSW1a/A
d61kT6OoU0Z3DyId7vA9TtvULucEUpywPpSoP/r+820UHFihdyhcb1iy8Z3v6KUc
arWzUOZpo0mc+o4hW2O1VnzNxLcXmhQOA9NdEOV/M+zxubFKo4VsY0ti9QIDAQAB
AoGAZuD/MBsEnMv02LmGHPHnsQWYrtu8/ZfeJ9sq1kve7u+ptE7O3Sr7y0FVPU8W
b+32cdFZ8rV/NuU63/yKNTBnZcbPwwGV9DmNpXy9YCdjwXkxfjYiDqUX9Fsxth1M
EqMb0PRO85akxCKxxtMagHDHNWkQaVThLagG31sh5d38SwECQQDuVsbRTbEz/H/j
Ip1NNU+8XERwMv1ac0LE9GhSRlqzUWDhukQ1gp9DmoKic8QMr6DS+JYvTCq38J8t
LHMNmzcpAkEA/xJHH/MwRlUSHsfP+DGXBuue2cAyw3NVLgusNV222kIgDOLcVxLl
8YOAgnheD5iI8+/GIVB4cXIfXKgqvzMC7QJAPUg8uMaEQLy02V8mGRsTFHiY9Ex4
DlDCo0fApx8F5UOQaJnvPd8HOme5HTIs/6IM9RIL879e4IrTMtdSAfad+QJBANAc
Opmv0mBgAnPItT8cPsvvrGCfdwuO6x2xemTkPE9hikLZSctlaOUfVNeem6f/3SWi
-----END RSA PRIVATE KEY-----`

//文件
var rootDir = "/root"
var rootFile1 = "/root/hello/world"
var admDir = `C:\Users\Administrator`
var dirCurr = "./"
var dirTdat = "./testdata"
var dirNew = "./testdata/new/hello"
var dirTouch = "./testdata/touch"
var dirCopy = "./testdata/copy"
var dirLink = "./testdata/link"
var changLog = "./docs/changelog.md"
var fileMd = "./README.md"
var fileGo = "./file.go"
var fileSongs = "./testdata/诗经.txt"
var fileDante = "./testdata/dante.txt"
var filePubPem = "./testdata/rsa/public_key.pem"
var filePriPem = "./testdata/rsa/private_key.pem"
var fileGitkee = "./testdata/.gitkeep"
var fileLink = "./testdata/lnk"
var copyLink = "./testdata/lnk_copy"
var fileNone = "./testdata/none"
var imgPng = "./testdata/diglett.png"
var imgJpg = "./testdata/gopher10th-small.jpg"
var imgSvg = "./testdata/jetbrains.svg"
var putfile = "./testdata/putfile"
var apndfile = "./testdata/appendfile"
var touchfile = "./testdata/touchfile"
var renamefile = "./testdata/renamefile"
var copyfile = "./testdata/copyfile"
var fastcopyfile = dirCopy + "/fast/fastcopyfile"
var imgCopy = dirCopy + "/diglett_copy.png"

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
