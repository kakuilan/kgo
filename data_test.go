//测试数据

package kgo

import (
	"bytes"
	"github.com/brianvoe/gofakeit/v6"
	"time"
)

//类型-接口
type itfType interface {
	noRealize() //该方法不实现
	sayHello(name string) string
}

//类型-人员
type sPerson struct {
	secret string ``
	Name   string `fake:"{name}" json:"name"`
	Addr   string `fake:"{city}" json:"city"`
	Age    int    `fake:"{number:1,99}" json:"age"`
	Gender bool   `fake:"{bool}" json:"gender"`
	other  int    ``
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

type userAccount struct {
	ID       uint32 `json:"id"`
	Status   bool   `json:"status"`
	Type     uint8  `json:"type"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"-"`
	Avatar   string `json:"avatar"`
}

type userAccountJson struct {
	ID       uint32 `json:"id"`
	Type     uint8  `json:"type"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	sPerson
}

//接口对象
var itfObj itfType

//结构体-人员
var personS1, personS2, personS3, personS4, personS5 sPerson

//结构体-人群
var crowd sPersons

//结构体-组织
var orgS1 = new(sOrganization) //返回指针

//结构体-用户账号
var account1 userAccount

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
var intTen = 10
var floSpeedLight float32 = 2.99792458
var bytAstronomicalUnit = []byte{0, 0, 0, 34, 212, 186, 90, 68} //intAstronomicalUnit的字节切片

var floAvogadro float64 = 6.02214129e23   // 阿伏伽德罗常数
var floPlanck float64 = 6.62606957e-34    // 普朗克常数
var floGravitional float64 = 6.673e-11    //重力常数
var floPermittivity float64 = 8.85419e-12 //真空介电常数
var floTen = 10.0
var floNum1 = 12345.12345678901231
var floNum2 = 12345678.12345678901231
var floNum3 = -123.4567890
var floNum4 float64 = 12345.12345678901252
var floNum5 = 1024000000000.0
var floNum6 = 1024000000000000000000000000000000000.0
var floNum7 = -10e-12

//复数
var cmplNum1 = complex(1, 2)
var cmplNum2 = complex(3, 4)

//字符串
var strHello = "Hello World! 你好！"
var b64Hello = "SGVsbG8gV29ybGQhIOS9oOWlve+8gQ=="
var strHelloHex = "48656c6c6f20576f726c642120e4bda0e5a5bdefbc81" //strHello的16进制
var utf8Hello = "你好，世界！"
var helloCn = "你好世界"
var helloEng = "hello world!"
var helloWidth = "ｈｅｌｌｏ　ｗｏｒｌｄ！"
var helloEngICase = "HelloWorld"
var helloEngUpper = "HELLOWORLD"
var helloEngLower = "helloworld"
var helloOther = "Hello world. 你好，世界。I`m use golang, python, and so on."
var helloOther2 = "Hello 你好, World 世界！"
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
var strJap = "ひらがな・カタカナ、．漢字"
var strKor = "안녕하세요"
var strSha1 = "82c9c0b34622756f6ef9731fbd8fbcef168a907f"
var strSha256 = "dcad188403ba3a4931288076f8398283abed9a90d1955364b3b5beeb551f0062"
var strSha512 = "057e65f970c85399b3953059b059c58c5b4eeeb572c741adb13af2fe2696f1ca3edc3757005aa801ea2bedc29529ba0c638e945fd95341d4dfbb6b693c3f6dfb"
var uuidNamespaceDNS = "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
var uuidNamespaceURL = "6ba7b811-9dad-11d1-80b4-00c04fd430c8"
var tesStr1 = "'test-bla-bla-4>2-y-3<6'"
var tesStr2 = "one%20%26%20two"
var tesStr3 = "'foo @+%/'你好"
var tesStr4 = `%27foo%20%40%2B%25%2F%27%E4%BD%A0%E5%A5%BD`
var tesStr5 = "Is your name O'reilly?"
var tesStr6 = `Is \ your \\name O\'reilly?`
var tesStr7 = `hello
world!
你好！`
var tesStr8 = `
hello world<br>
hello world<br/>
你好，世界<br />
hello world<BR>
hello world<BR/>
你好，世界<BR />
the end.
`
var tesStr9 = "hello World. Hello  \t \n world!   Text   \f\n\t\v\r\fMore \014\012\011\013\015here      \t\n\t Hello,\tWorld\n!\n\t"
var tesStr10 = `
<h1>Hello world!</h1>
<script>alert('你好！')</scripty>
`
var tesStr11 = "LeBronJames"
var tesStr12 = "Hello 你好, World 世界！"
var tesStr13 = "HELLO"
var tesStr14 = "world"
var tesStr15 = "ｆｏｏbar"
var tesStr16 = "ｘｙｚ０９８"
var tesStr17 = "１２３456"
var tesStr18 = "foobar"
var tesStr19 = "_Football"
var tesStr20 = "-Football"
var tesStr21 = " 3.124"
var tesStr22 = "作品T"
var tesStr23 = "8point"
var tesStr24 = "hello_Kitty2"
var tesStr25 = "hello-Kitty2"
var tesStr26 = "Hello ៉៊់៌៍！"
var tesStr27 = "pi314159"
var tesStr28 = "    "
var tesStr29 = "  \n  "
var tesStr30 = "\014\012\011\013\015"
var tesStr31 = "\014\012\011\013 abc  \015"
var tesStr32 = "\f\n\t\v\r\f"
var tesStr33 = "x\n\t\t\t\t"
var tesStr34 = "\f\n\t  \n\n\n   \v\r\f"
var tesStr35 = "Hi jac. $a=3*5, (can you hear me?)"
var tesStr36 = "A 'quote' is <b>bold</b>"
var tesStr37 = "A &#39;quote&#39; is &lt;b&gt;bold&lt;/b&gt;"
var tesStr38 = "The quick brown fox jumped over the lazy dog"
var tesStr39 = "中国"
var tesStr40 = "中华人民共和国"
var tesStr41 = "中华"
var tesStr42 = "000000"
var tesStr43 = "3.0.504"
var tesStr44 = "-3.14159"
var tesStr45 = "+3.14159"

//中文名
var tesChineseName1 = "李四"
var tesChineseName2 = "张三a"
var tesChineseName3 = "赵武灵王"
var tesChineseName4 = "南宫先生"
var tesChineseName5 = "吉乃•阿衣·依扎嫫"
var tesChineseName6 = "古丽莎•卡迪尔"
var tesChineseName7 = "迪丽热巴.迪力木拉提"

//公司名
var tesCompName1 = "北京搜狗科技公司"
var tesCompName2 = "北京搜狗科技发展有限公司"
var tesCompName3 = "工商发展银行深圳南山科苑梅龙路支行"

//标点符号、特殊字符
var strPunctuation1 = "<>@;.-="
var strPunctuation2 = "!\"#$%&()<>/+=-_? ~^|.,@`{}[]"
var strPunctuation3 = "`~!@#$%^&*()_+-=:'|<>?,./\""

//json
var strJson1 = `JsonpCallbackFn_abc123etc({"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}})`
var strJson2 = `myFunc([{"Name":"Bob","Age":32,"Company":"IBM","Engineer":true},{"Name":"John","Age":20,"Company":"Oracle","Engineer":false},{"Name":"Henry","Age":45,"Company":"Microsoft","Engineer":false}]);`
var strJson3 = "call)hello world(done"
var strJson4 = `JsonpCallbackFn_abc123etc({"meta":{"Status":200,"Content-Type":"application/json","Content-Length":"19","etc":"etc"},"data":{"name":"yummy"}})`
var strJson5 = `{"id":"1"}`
var strJson6 = `[{"key1":"value1"},{"key2":"value2"}]`
var strJson7 = `{"message_code":["bb9041bcfd55be4be20243b8e051963b","e5d94d692a4af45397a04c403d89bc3a"],"send_to":"tester","create_time":1641201974,"expire_time":4102415999}`

//email
var tesEmail1 = "test@example.com"
var tesEmail2 = "a@b.c"
var tesEmail3 = "hello-world@c"
var tesEmail4 = "ç$€§/az@gmail.com"
var tesEmail5 = "email@unkown_none_asdf_domain.com"
var tesEmail6 = "copyright@github.com"
var tesEmail7 = "abc@abc123.com"
var tesEmail8 = "test@163.com"

//手机号
var tesMobilecn1 = "13712345678"
var tesMobilecn2 = "17796325759"
var tesMobilecn3 = "15204810099"
var tesMobilecn4 = "18088664423"
var tesMobilecn5 = "12345678901"

//电话
var tesTel01 = "10086"
var tesTel02 = "010-88888888"
var tesTel03 = "021-87888822"
var tesTel04 = "0511-4405222"
var tesTel05 = "021-44055520-555"
var tesTel06 = "020-89571800-125"
var tesTel07 = "400-020-9800"
var tesTel08 = "400-999-0000"
var tesTel09 = "4006-589-589"
var tesTel10 = "4007005606"
var tesTel11 = "4000631300"
var tesTel12 = "400-6911195"
var tesTel13 = "800-4321"
var tesTel14 = "8004-321"
var tesTel15 = "8004321999"
var tesTel16 = "8008676014"

//身份证
var tesCredno01 = "123123123"
var tesCredno02 = "510723198006202551"
var tesCredno03 = "34052419800101001x"
var tesCredno04 = "511028199507215915"
var tesCredno05 = "511028199502315915"
var tesCredno06 = "53010219200508011X"
var tesCredno07 = "99010219200508011X"
var tesCredno08 = "130503670401001"
var tesCredno09 = "370986890623212"
var tesCredno10 = "370725881105149"
var tesCredno11 = "370725881105996"
var tesCredno12 = "35051419930513051X"
var tesCredno13 = "44141419900430157X"
var tesCredno14 = "110106209901012141"
var tesCredno15 = "513436200011013606"
var tesCredno16 = "51343620180101646X"

//颜色值
var tesColor01 = "#ff"
var tesColor02 = "fff0"
var tesColor03 = "#ff12FG"
var tesColor04 = "CCccCC"
var tesColor05 = "fff"
var tesColor06 = "#f00"
var tesColor07 = "#FAFAFA"
var tesColor08 = "#83C129"
var tesColor09 = "rgb(0,31,255)"
var tesColor10 = "rgb(0,  31, 255)"
var tesColor11 = "rgb(131, 193, 41)"
var tesColor12 = "rgb(1,349,275)"
var tesColor13 = "rgb(01,31,255)"
var tesColor14 = "rgb(0.6,31,255)"
var tesColor15 = "rgba(0,31,255)"

//base64
var tesBase64_01 = "Vml2YW11cyBmZXJtZtesting123" //false
var tesBase64_02 = "TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4="
var tesBase64_03 = "Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg=="
var tesBase64_04 = "U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw=="
var tesBase64_05 = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
	"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
	"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
	"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
	"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
	"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB"
var tesBase64_06 = "data:image/png;base6412345"
var tesBase64_07 = "data:image/png;base64,12345"
var tesBase64_08 = "data:text/plain;base64," + tesBase64_03
var tesBase64_09 = "data:image/png;base64," + tesBase64_02
var tesBase64_10 = "image/gif;base64," + tesBase64_04
var tesBase64_11 = "data:image/gif;base64," + tesBase64_05
var tesBase64_12 = "data:text,:;base85," + tesBase64_04

//html
var tesHtmlDoc = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>This is page title</title>
    <link rel="shortcut icon" href="/favicon.ico">
    <link href="/assets/css/frontend.min.css?v=0.0.1" rel="stylesheet">
    <link href="/assets/css/all.css?v=0.0.1" rel="stylesheet">
    <!--[if lt IE 9]>
    <script src="https://cdn.staticfile.org/html5shiv/r29/html5.min.js"></script>
    <script src="https://cdn.staticfile.org/respond.js/1.4.2/respond.min.js"></script>
    <![endif]-->
    <style>
        a{
            color: red;
        }
        span{
            margin: 5px;
        }
    </style>
</head>
<body>
    <div>
        <img src="/assets/img/nf.jpg" alt="this is image" class="fleft">
        <div class="fleft">最新公告</div>
        <div class="fright">
            <a href="logout" class="logoutBtn" style="display: none">退出</a>
            <a href="javascript:;" class="loginPwdBtn">登录</a>
            <a href="javascript:;" class="regisBtn">注册</a>
        </div>
        <h1>This is H1 title.</h1>
        <div>
            <p>
                Hello world!
                <span>TEXT <b>I</b> WANT</span>
            </p>
            <ul>
                <li><a href="foo">Foo</a><li>
                <a href="/bar/baz">BarBaz</a>
            </ul>

            <form name="query" action="http://www.example.net" method="post">
                <input type="text" value="123" />
                <textarea type="text" name="nameiknow">The text I want</textarea>
                <select>
                    <option value="111">111</option>
                    <option value="222">222</option>
                </select>
                <canvas>hello</canvas>
                <div id="button">
                    <input type="submit" value="Submit" />
                    <button>提交按钮</button>
                </div>
            </form>
        </div>
        <div>
            <iframe src="http://google.com"></iframe>
        </div>
    </div>
    <script type="text/javascript">
        var require = {
            config: {
                "modulename": "index",
                "controllername": "index",
                "actionname": "index",
                "jsname": "index",
                "moduleurl": "demo",
                "language": "zh-cn",
                "__PUBLIC__": "/",
                "__ROOT__": "/",
                "__CDN__": ""
            }
        };
        /* <![CDATA[ */
        var post_notif_widget_ajax_obj = {"ajax_url":"http:\/\/site.com\/wp-admin\/admin-ajax.php","nonce":"9b8270e2ef","processing_msg":"Processing..."};
        /* ]]> */
    </script>
    <script src="/assets/js/require.min.js" data-main="/assets/js/require-frontend.min.js?v=0.0.1"></script>
</body>
</html>
`

//时间
var strTime1 = "2019-07-11 10:11:23"
var strTime2 = "2020-02-01 13:39:36"
var strTime3 = "02/01/2016 15:04:05"
var strTime4 = "2020-03-10 23:04:35"
var strTime5 = "2020-03-08 23:04:35"
var strTime6 = "2020-06-25 23:59:59"
var strTime7 = "1990-01-02 03:14:59"
var intTime1 = 1562811851
var myDate1, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime4, time.Local)
var myDate2, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime5, time.Local)
var myDate3, _ = time.ParseInLocation("2006-01-02 15:04:05", strTime6, time.Local)

//当前时间
var nowNanoInt = kuptime.UnixNano()
var nowNanoStr = toStr(kuptime.UnixNano())

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
var tesIp1 = "255.255.255.255"
var tesIp2 = "::1"
var tesIp3 = "2001:db8:0000:1:1:1:1:1"
var tesIp4 = "300.0.0.0"
var tesIp5 = "192.168.0.1:80"
var tesIp6 = "::FFFF:C0A8:1"
var tesIp7 = "fe80::2c04:f7ff:feaa:33b7"
var tesIp8 = "8.8.8.8:8080"

//domain
var localHost = "localhost"
var tesDomain01 = "lÖcalhost"
var tesDomain02 = "localhost/"
var tesDomain03 = "a.bc"
var tesDomain04 = "a.b."
var tesDomain05 = "a.b.."
var tesDomain06 = "localhost.local"
var tesDomain07 = "localhost.localdomain.intern"
var tesDomain08 = "localhost.localdomain.intern:65535"
var tesDomain09 = "l.local.intern"
var tesDomain10 = "ru.link.n.svpncloud.com"
var tesDomain11 = "-localhost"
var tesDomain12 = "_localhost"
var tesDomain13 = "localhost.-localdomain"
var tesDomain14 = "localhost._localdomain"
var tesDomain15 = "localhost.localdomain.-int"
var tesDomain16 = "localhost.localdomain._int"
var tesDomain17 = "localhost.lÖcaldomain"
var tesDomain18 = "localhost.localdomain.üntern"
var tesDomain19 = "__"
var tesDomain20 = "[::1]"
var tesDomain21 = "www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de"
var tesDomain22 = "www.google.com"
var tesDomain23 = "localhost:80"
var tesDomain24 = "127.0.0.1:30000"
var tesDomain25 = "[::1]:80"
var tesDomain26 = "[1200::AB00:1234::2552:7777:1313]:22"
var tesDomain27 = "localhost.loc:100000"
var tesDomain28 = "漢字汉字:2"
var tesDomain29 = tesDomain21 + ":2000"
var tesDomain30 = "baidu.com"
var tesDomain31 = "golang.google.cn"
var tesDomain32 = "www.baidu.com"

//mac地址
var tesMac01 = "3D-F2-C9-A6-B3:4F"       //false
var tesMac02 = "fe80::5054:ff:fe4d:77d3" //false
var tesMac03 = "01:23:45:67:89:ab"
var tesMac04 = "01:23:45:67:89:ab:cd:ef"
var tesMac05 = "01-23-45-67-89-ab"
var tesMac06 = "01-23-45-67-89-ab-cd-ef"
var tesMac07 = "0123.4567.89ab"
var tesMac08 = "0123.4567.89ab.cdef"
var tesMac09 = "3D:F2:C9:A6:B3:4F"
var tesMac10 = "08:00:27:88:0f:fd"
var tesMac11 = "00:e0:66:07:5c:97:00:00"
var tesMac12 = "08:00:27:00:d8:94:00:00"
var tesMac13 = "02:42:b5:38:df:5a"
var tesMac14 = "0A-00-27-00-00-0E"

//bom字符
var tesBom1 = "\xEF\xBB\xBF"
var tesBom2 = bomChars + "hello"
var tesBom3 = tesBom1 + "world"

//自然数数组
var naturalArr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, intTen}

//整数切片
var intSlc = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}
var intSlEmp = []int{}
var int64Slc = []int64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9, 8, 7, 6, 5, 4, 11, 12, 13, 14, 15}

//浮点切片
var flo32Slc = []float32{1.23, 0.0, flPi1, floSpeedLight, 6.6260755, 1.60217733}
var flo64Slc = []float64{flPi2, floAvogadro, floPlanck, floGravitional, floPermittivity, floTen}
var flo64Slc2 = []float64{flPi2, floNum1, floNum2, floNum3, floNum4}

//布尔切片
var booSlc = []bool{true, true, false, true, false, true, true}

//字节切片
var bytsHello = []byte(strHello)
var runesHello = []rune(strHello)
var bytSpeedLight = []byte(strSpeedLight)
var bytsPasswd = []byte("$2a$10$j3WOP6rP2I7skNoxiFdNdOh6OhPxP0Sp3Wmeuekh90oeF3D1EQQBK")
var bytCryptKey = []byte("1234567890123456")
var bytsUtf8Hello = []byte(utf8Hello)
var bytsGbkHello = []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
var bytsUuidNamespaceDNS = bytes.Replace([]byte(uuidNamespaceDNS), bytMinus, bytEmpty, -1)
var bytsUuidNamespaceUrl = bytes.Replace([]byte(uuidNamespaceURL), bytMinus, bytEmpty, -1)

//单字符切片
var ssSingle = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

//字符串map
var strMp1 = map[string]string{"a": "1", "b": "2", "c": "3", "d": "4", "e": "", "2": "cc", "3": "no"}
var strMp2 = map[string]string{"a": "0", "b": "2", "c": "4", "g": "4", "h": "", "2": "cc"}
var strMpEmp = make(map[string]string)
var colorMp = map[string]string{"a": "green", "0": "red", "b": "green", "1": "blue", "2": "red", "c": "yellow", "n": ""}

//字符串切片
var strSl1 = []string{"aa", "bb", "cc", "dd", "ee", "", "hh", "ii"}
var strSl2 = []string{"bb", "cc", "ff", "gg", "ee", "", "gg"}
var strSl3 = []string{"hehe,php lang", "Hello,go language", "HeLlo,python!", "haha,java", "I`m going."}
var strSlEmp = []string{}

//接口切片
var slItf = []interface{}{99, 0, 1, 2, 0.0, 3, false, 3.14, 6.67428, true, 'a', "", 'b', nil, 'c', intSpeedLight, "hello", nowNanoInt, floAvogadro, strSlEmp, "你好", floNum3}
var slItf2 = []interface{}{1, 0, 1.2, -3, false, nil, "4"}

//persons JSON串
var personsMapJson = `{"person1":{"name":"zhang3","age":23,"sex":1},"person2":{"name":"li4","age":30,"sex":1},"person3":{"name":"wang5","age":25,"sex":0},"person4":{"name":"zhao6","age":50,"sex":0}}`
var personsArrJson = `[{"age":20,"naction":"us","name":"test1","tel":"13712345678"},{"age":21,"naction":"cn","name":"test2","tel":"13712345679"},{"age":22,"naction":"en","name":"test3","tel":"13712345670"},{"age":23,"naction":"fr","name":"test4","tel":"13712345671"},{"age":21,"naction":"cn","name":"test5","tel":"13712345672"}]`

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

//RSA公钥
//正确的
var tesRsaPubKey01 = `MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuu
XwKYLq0DKUE3t/HHsNdowfD9+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9BmMEcI3uoKbeXCbJRI
HoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzTUmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZ
B7ucimFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUvbQIDAQAB`
var tesRsaPubKey02 = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvncDCeibmEkabJLmFec7
x9y86RP6dIvkVxxbQoOJo06E+p7tH6vCmiGHKnuuXwKYLq0DKUE3t/HHsNdowfD9
+NH8caLzmXqGBx45/Dzxnwqz0qYq7idK+Qff34qrk/YFoU7498U1Ee7PkKb7/VE9
BmMEcI3uoKbeXCbJRIHoTp8bUXOpNTSUfwUNwJzbm2nsHo2xu6virKtAZLTsJFzT
UmRd11MrWCvj59lWzt1/eIMN+ekjH8aXeLOOl54CL+kWp48C+V9BchyKCShZB7uc
imFvjHTtuxziXZQRO7HlcsBOa0WwvDJnRnskdyoD31s4F4jpKEYBJNWTo63v6lUv
bQIDAQAB
-----END PUBLIC KEY-----`

//错误的
var tesRsaPubKey03 = `-----BEGIN PRIVATE KEY-----
MIICdQIBADANBgkqhkiG9w0BAQEFAASCAl8wggJbAgEAAoGBAKn4X6phG2ZsKjof
ytRsM8zC7VTZmQSi9hr7ZqHxsIe+UeGToXLSqfJ8ikWWMg15N8PTbzIG11GTexyd
QH/u+zPAS//qrf0HbCXjICt741A8qMipMHIG409PYLQWjfnrjusLt51dY84llj9C
7BzXlHvWqowBGU5jCEaQTBAHPRutAgMBAAECgYAYNdeylihn+2o8Y0Dp5wut0+oo
VuJT5b52c27YDGwfub1CC1xI1bb9Yj3z0YQJpUWLMDe7gXv0E7TKi5+fWXQQXJWt
ejTBtbf0hE14x6OqTzazess99UAxKIdsk7trzVRlPkE4NpJ5jAGTzPqHPlkuaFb3
IK3dyQGLas5QriFnAQJBANagrgmfxygmwH+i7QacffZ6yTu+rhyAcdeUSu6ekPUu
ITv8mOA/bT2m9sIGinW3gjf8KMfz9JH11TasZVsL8e0CQQDKu/bc9oTI0f2jRupY
vmrc31rmOdPq4C4Z6Uj00Ui/FicdywUnGF0bvA+jlCUTLEqBYerl3EEHeLiyZsbT
E5jBAkBVhIZz/T78h5xR/xgUd0xVZo1CCfMUFjXGISdONs4pcyz42ugLChq74wgV
PUf0KZ9wMUAKk/DSK7K96ykjgvntAkBwmqBOMLqmFETN2Mi3S+RtE74YXAxBzAyv
Jaz5FflS8Yn+eVI+WcD1c6o4EEPbd2FWpb1juMeBz+K+bGmIubzBAkB61Sd8LvfF
fDA7MDOGRtIcWq+7bPPw3y44RYIKA35ocMAlzHFhXw7RtSLCl6xgzIpkIfW4ilCP
oCbhuSHBcPnj
-----END PRIVATE KEY-----`
var tesRsaPubKey04 = "JXU4RkQ5JXU5MUNDJXU2NjJGJXU4OTgxJXU1MkEwJXU1QkM2JXU3Njg0JXU1MTg1JXU1QkI5JXVGRjAx"
var tesRsaPubKey05 = `-----BEGIN PUBLIC KEY-----
MIIDRzCCAjkGByqGSM44BAEwggIsAoIBAQCYBeAV/nYFehIyAJqGBSl6Kqthllr5
25iJYG7R9V+/wG5oaVtFJSow/vexBaQ0D5fLQZHJhOPPd+QkEQeMWXVh1mLv0a/V
tbVzA/X5nPrh6qf3SK1fO3cM19Z2YFqCE9sXtrDfroi/DR9Ze1uDT/HVDJ23iZZ7
x7f8cegQN23jOv1APz2d4OEqGe1s85RcS0RPoRrBe1e5itaM1EU0eCCaUjozYt4H
dLZ/VhYZlTG5k814EqrAX+4aWFXUKW1X374a6cvfXirGzZfYr90pL/8VAHATbR2O
P6R0VrdZ0W1hfwPkPb9zBZMaV3+A1HewCjsuheXIKLxnIG+SbceMyYizAiEAkr9Q
R4mvyGhvC79HoQxjRJZRYYqf1O92Yn1dixROC+sCggEAL0rHy4qOIW3g4l/FFh4y
uzzXXePBooCc2jpdYlGXa9g9B5ueX2GQ5+f/QB0VoXvGeYaXefo2YTW5B45IHn7W
9ceX9yme3n9tl8H1dK3sjyqQKxAhyynM1wJaBaALhYT0NzuCXEoBq3kn7On3rU8d
/LM+1UoDwJ0iPqooI9xDW5UX8xd+iYV2FzMtc+SWu4YWmH57EKjcOgC9MqPzCpIn
1Cgo7nSexzSCYIXGDVOqJ0hjeHlL54CMOON2EkUg0e3J/mcneTT8YbP8zPMuBrEX
vwPWNk8wJr2rtxpjhny/sj8BCJY5hhKQFHL1kive7i16AQJv3gJn42eGFJgBsdYa
lgOCAQYAAoIBAQCFyXq2x1BWFxj8qQrbGl5bojxO4r8+gnIoCIbzaxJbiK+eo+JT
BiJNQlludq8f1+0SZ9Paiv1qLaH5p1qxw7mz4ZU8HO4+9grDIb1tuWld/RyhH9PJ
NIoXIVT1J6lK8DqpjnIIoIjqHh5kSJNnXw6XQrA5nlcdZfokVl9oXjH0tGl3McdZ
TQ3WVV0EekGzoIrPw7BkGgb71UBedEt9AqkLSnW6KzQ1A1ILokX8Yq9oWLASea3F
9UxJXpPlCRz3FYgvuR+Q07thgm/z3VQ/+Uq0PFsGFB7Cern0vOKZ+E4673jYK9nq
xVZ+SCC8Wd6nIK4FyZbYaa3Jz7GkqHdMelsl
-----END PUBLIC KEY-----`

//文件
var rootDir = "/root"
var rootDir2 = "/root/hello/directory"
var rootFile1 = "/root/hello/world"
var rootFile2 = "/root/hello/ok.zip"
var rootFile3 = "/root/tar/test2.tar.gz"
var admDir = `C:\Users\Administrator`
var admTesDir = admDir + `\Test`
var dirCurr = "./"
var dirDoc = "./docs"
var dirTdat = "./testdata"
var dirNew = "./testdata/new/hello"
var dirTouch = "./testdata/touchs"
var dirCopy = "./testdata/copys"
var dirLink = "./testdata/links"
var dirChmod = "./testdata/chmod"
var dirVendor = "./vendor"
var changLog = "./docs/changelog.md"
var fileMd = "./README.md"
var fileGo = "./file.go"
var fileGmod = "go.mod"
var fileSongs = "./testdata/诗经.txt"
var fileDante = "./testdata/dante.txt"
var filePubPem = "./testdata/rsa/public_key1024.pem"
var filePriPem = "./testdata/rsa/private_key1024.pem"
var filePubPem2048 = "./testdata/rsa/public_key2048.pem"
var filePriPem2048 = "./testdata/rsa/private_key2048.pem"
var fileGitkee = "./testdata/.gitkeep"
var fileNone = "./testdata/none"
var fileLink = "./testdata/lnk"
var copyLink = "./testdata/lnk_copy"
var copyLink2 = "./vendor/lnk_copy"
var imgPng = "./testdata/diglett.png"
var imgJpg = "./testdata/gopher10th-small.jpg"
var imgSvg = "./testdata/jetbrains.svg"
var imgNone = "./testdata/none-image.jpeg"
var gitkeep = "./testdata/.gitkeep"
var putfile = "./testdata/putfile"
var apndfile = "./testdata/appendfile"
var touchfile = "./testdata/touchfile"
var renamefile = "./testdata/renamefile"
var copyfile = "./testdata/copyfile"
var chownfile = "./testdata/chownfile"
var fastcopyfile = dirCopy + "/fast/fastcopyfile"
var imgCopy = dirCopy + "/diglett_copy.png"
var pathTes1 = `/usr|///tmp:\\\123/\abc:d<|\hello>\/%world?\\how$\\are@#test.png`
var pathTes2 = `C:\Users\/Administrator/\AppData\:Local`
var pathTes3 = `/usr\bin\\golang//fmt/\test\/hehe`
var pathTes4 = `123456789     '_-?ASDF@£$%£%^é.html`
var pathTes5 = `file:///c:/test.go`
var pathTes6 = `../../../Hello World!.txt`
var targzfile1 = "./testdata/targz/test1.tar.gz"
var targzfile2 = "./testdata/targz/test2.tar.gz"
var untarpath1 = "./testdata/targz/un1"
var zipfile1 = "./testdata/zip/test1.zip"
var zipfile2 = "./testdata/zip/test2.zip"
var unzippath1 = "./testdata/zip/un1"

//uri
var tesUri1 = `?first=value&arr[]=foo+bar&arr[]=baz`
var tesUri2 = `f1=m&f2=n`
var tesUri3 = `f[a]=m&f[b]=n`
var tesUri4 = `f[a][a]=m&f[a][b]=n`
var tesUri5 = `f[]=m&f[]=n`
var tesUri6 = `f[a][]=m&f[a][]=n`
var tesUri7 = `f[][]=m&f[][]=n`
var tesUri8 = `a .[[b=c`
var tesUri9 = `f=m&f[a]=n`
var tesUri10 = `f=n&f[a]=m&`
var tesUri11 = `f=n&f[][a]=m&`
var tesUri12 = `f[][a]=&f[][b]=`
var tesUri13 = `f[][a]=m&f[][b]=h`
var tesUri14 = `f=n&f[a][]=m&`
var tesUri15 = `f=n&f[a][]b=m&`
var tesUri16 = `f[][b]=&f[][a]=12&f[][a]=1.2&f[][a]=abc`
var tesUri17 = `f[a].=m&f=n&`
var tesUri18 = `f[a][]=1&f[a][]=c&f[a][]=&f[b][]=bb&f[]=3&f[]=4`
var tesUri19 = `f[a][]=12&f[a][]=1.2&f[a][]=abc`
var tesUri20 = `?first=value&arr[]=foo+bar&arr[]=baz&arr[][a]=aaa`
var tesUri21 = `%=%gg&b=4`
var tesUri22 = `he& =2`
var tesUri23 = `he& g=2`
var tesUri24 = `he&=3`
var tesUri25 = `he&[=4`
var tesUri26 = `he&]=5`
var tesUri27 = `he&a=1`
var tesUri28 = `he&e=%&b=4`

//url
var tesUrl01 = `https://www.google.com/search?source=hp&ei=tDUwXejNGs6DoATYkqCYCA&q=golang&oq=golang&gs_l=psy-ab.3..35i39l2j0i67l8.1729.2695..2888...1.0..0.126.771.2j5......0....1..gws-wiz.....10..0.fFQmXkC_LcQ&ved=0ahUKEwjo9-H7jb7jAhXOAYgKHVgJCIMQ4dUDCAU&uact=5`
var tesUrl02 = `sg>g://asdf43123412341234`
var tesUrl03 = "abc.com"
var tesUrl04 = "abc.com/hello?a=1"
var tesUrl05 = `http://login.localhost:3000\/ab//cd/ef///hi\\12/33\`
var tesUrl06 = "https://play.golang.com:3000/p/3R1TPyk8qck"
var tesUrl07 = "https://www.siongui.github.io/pali-chanting/zh/archives.html"
var tesUrl08 = "http://foobar.中文网/"
var tesUrl09 = "foobar.com/abc/efg/h=1"
var tesUrl10 = "https://github.com/kakuilan/kgo"
var tesUrl11 = "////google.com/test?name=hello"
var tesUrl12 = "google.com/test?name=hello////"
var tesUrl13 = ".com"
var tesUrl14 = "ftp://foobar.ru/"
var tesUrl15 = "http://127.0.0.1/"
var tesUrl16 = "http://duckduckgo.com/?q=%2F"
var tesUrl17 = "http://foo.bar/#com"
var tesUrl18 = "http://foobar.coffee/"
var tesUrl19 = "http://foobar.com"
var tesUrl20 = "http://foobar.com/#baz=qux"
var tesUrl21 = "http://foobar.com/?foo=bar#baz=qux"
var tesUrl22 = "http://foobar.com/t$-_.+!*\\'(),"
var tesUrl23 = "http://foobar.com?foo=bar"
var tesUrl24 = "http://foobar.org:8080/"
var tesUrl25 = "http://localhost:3000/"
var tesUrl26 = "http://user:pass@www.foobar.com/"
var tesUrl27 = "http://www.-foobar.com/"
var tesUrl28 = "http://www.foo---bar.com/"
var tesUrl29 = "http://www.foo_bar.com/"
var tesUrl30 = "http://www.foobar.com/~foobar"
var tesUrl31 = "http://www.xn--froschgrn-x9a.net/"
var tesUrl32 = "https://foobar.com"
var tesUrl33 = "https://foobar.org/"
var tesUrl34 = "invalid."
var tesUrl35 = "irc://irc.server.org/channel"
var tesUrl36 = "mailto:someone@example.com"
var tesUrl37 = "rtmp://foobar.com"
var tesUrl38 = "xyz://foobar.com"
var tesUrl39 = "https://www.baidu.com/"
var tesUrl40 = "https://www.w3.org/"

//下载文件
var downloadfile01 = "./testdata/download/test001/file001"

//命令
var tesCommand01 = " ls -a -h"
var tesCommand02 = " ls -a\"\" -h 'hehe'"
var tesCommand03 = "cmd /C dir "

//等式
var equationStr01 = "190000017056834?utm_source=tag-newest "
var equationStr02 = `String str = "AB==2LSKF=5!@!=$%^()==AD=";`
var equationStr03 = `    | | |   {
    | | |     "IOUserClientCreator" = "pid 195, loginwindow"
    | | |   }
    
  +-o VMware7,1  <class IOPlatformExpertDevice, id 0x100000112, registered, matched, active, busy 0 (33207 ms), retain 27>
    | {
    |   "compatible" = <"VMware7,1">
    |   "version" = <"None">
    |   "board-id" = <"440BX Desktop Reference Platform">
    |   "IOInterruptSpecifiers" = (<0900000007000000>)
    |   "IOPolledInterface" = "SMCPolledInterface is not serializable"
    |   "serial-number" = <764f445a000000000000000000564d54464d475a71764f445a000000000000000000000000000000000000>
    |   "IOInterruptControllers" = ("io-apic-0")
    |   "IOPlatformUUID" = "4203018E-580F-C1B5-9525-B745CECA79EB"
    |   "target-type" = <"Mac">
    |   "clock-frequency" = <00e1f505>
    |   "manufacturer" = <"VMware, Inc.">
    |   "IOPlatformSerialNumber" = "VMTFMGZqvODZ"
    |   "product-name" = <"VMware7,1">
    |   "IOBusyInterest" = "IOCommand is not serializable"
    |   "model" = <"VMware7,1">
    |   "name" = <"/">
    | }
    | 
    +-o AppleACPIPlatformExpert  <class AppleACPIPlatformExpert, id 0x100000113, registered, matched, active, busy 0 (33168 ms), retain 30>
    | | {
    | |   "IOClass" = "AppleACPIPlatformExpert"
    | |   "CFBundleIdentifier" = "com.apple.driver.AppleACPIPlatform"
    | |   "IOProviderClass" = "IOPlatformExpertDevice"
    | |   "IOProbeScore" = 10000
    | |   "IONameMatch" = "ACPI"
    | |   "acpi-mmcfg-seg0" = 3758096384
    | |   "IOMatchCategory" = "IODefaultMatchCategory"
    | |   "IOPolledInterface" = "AppleACPIEventPoller is not serializable"
    | |   "IOPlatformMaxBusDelay" = (18446744073709551615,0)
    | |   "IONameMatched" = "ACPI"
    | |   "Platform Memory Ranges" = (0,4294967295)
    | |   "IOPlatformMaxInterruptDelay" = (18446744073709551615,0)
    | |   "CFBundleIdentifierKernel" = "com.apple.driver.AppleACPIPlatform"
    | |   "ACPI Statistics" = {"MethodCount"=412,"SciCount"=0,"GpeCount"=0,"FixedEventCount"=0}
    | | }
    | |     `

//表情符
var tesEmoji1 = `Lorem ipsum 🥊dolor 🤒sit amet, consectetur adipiscing 🍂 elit. 🍁🍃🍂🌰🍁🌿🌾🌼🌻سلام تست شد hell中文
😀😁😂😃😄😅😆😉😊😋😎😍😘😗😙😚☺😇😐😑😶😏😣😥😮😯😪😫😴😌😛😜😝😒😓😔😕😲😷😖😞😟😤😢😭😦😧😨😬😰😱😳😵😡😠
👦👧👨👩👴👵👶👱👮👲👳👷👸💂🎅👰👼💆💇🙍🙎🙅🙆💁🙋🙇🙌🙏👤👥🚶🏃👯💃👫👬👭💏💑👪
💪👈👉☝👆👇✌✋👌👍👎✊👊👋👏👐✍
👣👀👂👃👅👄💋👓👔👕👖👗👘👙👚👛👜👝🎒💼👞👟👠👡👢👑👒🎩🎓💄💅💍🌂
📱📲📶📳📴☎📞📟📠
♻🏧🚮🚰♿🚹🚺🚻🚼🚾⚠🚸⛔🚫🚳🚭🚯🚱🚷🔞💈
🙈🙉🙊🐵🐒🐶🐕🐩🐺🐱😺😸😹😻😼😽🙀😿😾🐈🐯🐅🐆🐴🐎🐮🐂🐃🐄🐷🐖🐗🐽🐏🐑🐐🐪🐫🐘🐭🐁🐀🐹🐰🐇🐻🐨🐼🐾🐔🐓🐣🐤🐥🐦🐧🐸🐊🐢🐍🐲🐉🐳🐋🐬🐟🐠🐡🐙🐚🐌🐛🐜🐝🐞
💐🌸💮🌹🌺🌻🌼🌷🌱🌲🌳🌴🌵🌾🌿🍀🍁🍂🍃
🌍🌎🌏🌐🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜☀🌝🌞⭐🌟🌠☁⛅☔⚡❄🔥💧🌊
🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓🍅🍆🌽🍄🌰🍞🍖🍗🍔🍟🍕🍳🍲🍱🍘🍙🍚🍛🍜🍝🍠🍢🍣🍤🍥🍡🍦🍧🍨🍩🍪🎂🍰🍫🍬🍭🍮🍯🍼☕🍵🍶🍷🍸🍹🍺🍻🍴
🎪🎭🎨🎰🚣🛀🎫🏆⚽⚾🏀🏈🏉🎾🎱🎳⛳🎣🎽🎿🏂🏄🏇🏊🚴🚵🎯🎮🎲🎷🎸🎺🎻🎬
😈👿👹👺💀☠👻👽👾💣
🌋🗻🏠🏡🏢🏣🏤🏥🏦🏨🏩🏪🏫🏬🏭🏯🏰💒🗼🗽⛪⛲🌁🌃🌆🌇🌉🌌🎠🎡🎢🚂🚃🚄🚅🚆🚇🚈🚉🚊🚝🚞🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚚🚛🚜🚲⛽🚨🚥🚦🚧⚓⛵🚤🚢✈💺🚁🚟🚠🚡🚀🎑🗿🛂🛃🛄🛅
💌💎🔪💈🚪🚽🚿🛁⌛⏳⌚⏰🎈🎉🎊🎎🎏🎐🎀🎁📯📻📱📲☎📞📟📠🔋🔌💻💽💾💿📀🎥📺📷📹📼🔍🔎🔬🔭📡💡🔦🏮📔📕📖📗📘📙📚📓📃📜📄📰📑🔖💰💴💵💶💷💸💳✉📧📨📩📤📥📦📫📪📬📭📮✏✒📝📁📂📅📆📇📈📉📊📋📌📍📎📏📐✂🔒🔓🔏🔐🔑🔨🔫🔧🔩🔗💉💊🚬🔮🚩🎌💦💨
♠♥♦♣🀄🎴🔇🔈🔉🔊📢📣💤💢💬💭♨🌀🔔🔕✡✝🔯📛🔰🔱⭕✅☑✔✖❌❎➕➖➗➰➿〽✳✴❇‼⁉❓❔❕❗©®™🎦🔅🔆💯🔠🔡🔢🔣🔤🅰🆎🅱🆑🆒🆓ℹ🆔Ⓜ🆕🆖🅾🆗🅿🆘🆙🆚🈁🈂🈷🈶🈯🉐🈹🈚🈲🉑🈸🈴🈳㊗㊙🈺🈵▪▫◻◼◽◾⬛⬜🔶🔷🔸🔹🔺🔻💠🔲🔳⚪⚫🔴🔵
🐁🐂🐅🐇🐉🐍🐎🐐🐒🐓🐕🐖
♈♉♊♋♌♍♎♏♐♑♒♓⛎
🕛🕧🕐🕜🕑🕝🕒🕞🕓🕟🕔🕠🕕🕡🕖🕢🕗🕣🕘🕤🕙🕥🕚🕦⌛⏳⌚⏰⏱⏲🕰
💘❤💓💔💕💖💗💙💚💛💜💝💞💟❣
💐🌸💮🌹🌺🌻🌼🌷🌱🌿🍀
🌿🍀🍁🍂🍃
🌑🌒🌓🌔🌕🌖🌗🌘🌙🌚🌛🌜🌝
🍇🍈🍉🍊🍋🍌🍍🍎🍏🍐🍑🍒🍓
💴💵💶💷💰💸💳
🚂🚃🚄🚅🚆🚇🚈🚉🚊🚝🚞🚋🚌🚍🚎🚏🚐🚑🚒🚓🚔🚕🚖🚗🚘🚚🚛🚜🚲⛽🚨🚥🚦🚧⚓⛵🚣🚤🚢✈💺🚁🚟🚠🚡🚀
🏠🏡🏢🏣🏤🏥🏦🏨🏩🏪🏫🏬🏭🏯🏰💒🗼🗽⛪🌆🌇🌉
📱📲☎📞📟📠🔋🔌💻💽💾💿📀🎥📺📷📹📼🔍🔎🔬🔭📡📔📕📖📗📘📙📚📓📃📜📄📰📑🔖💳✉📧📨📩📤📥📦📫📪📬📭📮✏✒📝📁📂📅📆📇📈📉📊📋📌📍📎📏📐✂🔒🔓🔏🔐🔑
⬆↗➡↘⬇↙⬅↖↕↔↩↪⤴⤵🔃🔄🔙🔚🔛🔜🔝`
var tesEmoji2 = `Hi!😀👽😀☂❤华み원❤This is a string 😄 🐷 with some 👍🏻 🙈 emoji! 🐷 🏃🏿‍♂️`

func init() {
	_ = gofakeit.Struct(&personS1)
	_ = gofakeit.Struct(&personS2)
	_ = gofakeit.Struct(&personS3)
	_ = gofakeit.Struct(&personS4)
	_ = gofakeit.Struct(&personS5)
	_ = gofakeit.Struct(&account1)

	crowd = append(crowd, personS1, personS2, personS3, personS4, personS5)

	orgS1.Leader = personS1
	orgS1.Assistant = personS2
	orgS1.Substitute = personS3
	orgS1.Members = sPersons{personS4, personS5}

	perStuMps = map[string]sPerson{"a": personS1, "b": personS2, "c": personS3, "d": personS4, "e": personS5}
}
