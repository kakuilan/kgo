package kgo

import (
	"net"
	"time"
)

type (
	// LkkFile is the receiver of file utilities
	LkkFile byte
	// LkkString is the receiver of string utilities
	LkkString byte
	// LkkNumber is the receiver of number utilities
	LkkNumber byte
	// LkkArray is the receiver of array utilities
	LkkArray byte
	// LkkTime is the receiver of time utilities
	LkkTime byte
	// LkkConvert is the receiver of convert utilities
	LkkConvert byte
	// LkkOS is the receiver of OS utilities
	LkkOS byte
	// LkkUrl is the receiver of url utilities
	LkkUrl byte
	// LkkEncrypt is the receiver of encrypt utilities
	LkkEncrypt byte
	// LkkDebug is the receiver of debug utilities
	LkkDebug byte

	// LkkFileCover 枚举类型,文件是否覆盖
	LkkFileCover int8
	// LkkFileTree 枚举类型,文件树查找类型
	LkkFileTree int8
	// LkkRandString 枚举类型,随机字符串类型
	LkkRandString uint8

	// LkkCaseSwitch 枚举类型,大小写开关
	LkkCaseSwitch uint8

	// FileFilter 文件过滤函数
	FileFilter func(string) bool

	// CallBack 回调执行函数,无参数且无返回值
	CallBack func()
)

const (
	// Version 版本号
	Version = "0.0.1"

	//UINT_MAX 无符号整型uint最大值
	UINT_MAX = ^uint(0)
	//UINT_MIN 无符号整型uint最小值
	UINT_MIN uint = 0

	//UINT64_MAX 无符号整型uint64最大值, 18446744073709551615
	UINT64_MAX = ^uint64(0)
	//UINT64_MIN 无符号整型uint64最小值
	UINT64_MIN uint64 = 0

	//INT_MAX 有符号整型int最大值
	INT_MAX = int(^uint(0) >> 1)
	//INT_MIN 有符号整型int最小值
	INT_MIN = ^INT_MAX

	//INT64_MAX 有符号整型int64最大值, 9223372036854775807
	INT64_MAX = int64(^uint(0) >> 1)
	//INT64_MIN 有符号整型int64最小值, -9223372036854775808
	INT64_MIN = ^INT64_MAX

	// FILE_COVER_ALLOW 文件覆盖,允许
	FILE_COVER_ALLOW LkkFileCover = 1
	// FILE_COVER_IGNORE 文件覆盖,忽略
	FILE_COVER_IGNORE LkkFileCover = 0
	// FILE_COVER_DENY 文件覆盖,禁止
	FILE_COVER_DENY LkkFileCover = -1

	// FILE_TREE_ALL 文件树,查找所有(包括目录和文件)
	FILE_TREE_ALL LkkFileTree = 3
	// FILE_TREE_DIR 文件树,仅查找目录
	FILE_TREE_DIR LkkFileTree = 2
	// FILE_TREE_FILE 文件树,仅查找文件
	FILE_TREE_FILE LkkFileTree = 1

	// RAND_STRING_ALPHA 随机字符串类型,字母
	RAND_STRING_ALPHA LkkRandString = 0
	// RAND_STRING_NUMERIC 随机字符串类型,数值
	RAND_STRING_NUMERIC LkkRandString = 1
	// RAND_STRING_ALPHANUM 随机字符串类型,字母+数值
	RAND_STRING_ALPHANUM LkkRandString = 2
	// RAND_STRING_SPECIAL 随机字符串类型,字母+数值+特殊字符
	RAND_STRING_SPECIAL LkkRandString = 3
	// RAND_STRING_CHINESE 随机字符串类型,仅中文
	RAND_STRING_CHINESE LkkRandString = 4

	// CASE_NONE 忽略大小写
	CASE_NONE LkkCaseSwitch = 0
	// CASE_LOWER 检查小写
	CASE_LOWER LkkCaseSwitch = 1
	// CASE_UPPER 检查大写
	CASE_UPPER LkkCaseSwitch = 2

	//检查连接超时的时间
	CHECK_CONNECT_TIMEOUT = time.Second * 5

	// 正则模式-全中文
	PATTERN_ALL_CHINESE = "^[\u4e00-\u9fa5]+$"

	// 正则模式-浮点数
	PATTERN_FLOAT = `^(-?\d+)(\.\d+)?`

	// 正则模式-邮箱
	PATTERN_EMAIL = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	// 正则模式-大陆手机号
	PATTERN_MOBILE = `^1[3-9]\d{9}$`

	// 正则模式-固定电话
	PATTERN_TEL_FIX = `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$`

	// 正则模式-400或800
	PATTERN_TEL_4800 = `^[48]00\d?(-?\d{3,4}){2}$`

	// 正则模式-座机号(固定电话或400或800)
	PATTERN_TELEPHONE = `(` + PATTERN_TEL_FIX + `)|(` + PATTERN_TEL_4800 + `)`

	// 正则模式-电话(手机或固话)
	PATTERN_PHONE = `(` + PATTERN_MOBILE + `)|(` + PATTERN_TEL_FIX + `)`

	// 正则模式-日期时间
	PATTERN_DATETIME = `^[0-9]{4}(|\-[0-9]{2}(|\-[0-9]{2}(|\s+[0-9]{2}(|:[0-9]{2}(|:[0-9]{2})))))$`

	// 正则模式-身份证号码,18位或15位
	PATTERN_CREDIT_NO = `(^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`

	// 正则模式-小写英文
	PATTERN_ALPHA_LOWER = `^[a-z]+$`

	// 正则模式-大写英文
	PATTERN_ALPHA_UPPER = `^[A-Z]+$`
)

var (
	//Kuptime 当前服务启动时间
	Kuptime = time.Now()

	// KFile utilities
	KFile LkkFile

	// KStr utilities
	KStr LkkString

	// KNum utilities
	KNum LkkNumber

	// KArr utilities
	KArr LkkArray

	// KTime utilities
	KTime LkkTime

	// KConv utilities
	KConv LkkConvert

	// KOS utilities
	KOS LkkOS

	// KUrl utilities
	KUrl LkkUrl

	// KEncr utilities
	KEncr LkkEncrypt

	// KDbug utilities
	KDbug LkkDebug

	// KPrivCidrs 私有网段的CIDR数组
	KPrivCidrs []*net.IPNet

	// 身份证区域
	CreditArea = map[string]string{"11": "北京", "12": "天津", "13": "河北", "14": "山西", "15": "内蒙古", "21": "辽宁", "22": "吉林", "23": "黑龙江", " 31": "上海", "32": "江苏", "33": "浙江", "34": "安徽", "35": "福建", "36": "江西", "37": "山东", "41": "河南", "42": "湖北", "43": "湖南", "44": "广东", "45": "广西", "46": "海南", "50": "重庆", "51": "四川", "52": "贵州", "53": "云南", "54": "西藏", "61": "陕西", "62": "甘肃", "63": "青海", "64": "宁夏", "65": "新疆", "71": "台湾", "81": "香港", "82": "澳门", "91": "国外"}
)
