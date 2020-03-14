package kgo

import (
	"net"
	"regexp"
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
	// LkkEncrypt is the receiver of encrypt utilities
	LkkEncrypt byte
	// LkkDebug is the receiver of debug utilities
	LkkDebug byte

	// LkkFileCover 枚举类型,文件是否覆盖
	LkkFileCover int8
	// LkkFileType 枚举类型,文件类型
	LkkFileType uint8
	// LkkFileTree 枚举类型,文件树查找类型
	LkkFileTree uint8
	// LkkRandString 枚举类型,随机字符串类型
	LkkRandString uint8
	// LkkCaseSwitch 枚举类型,大小写开关
	LkkCaseSwitch uint8
	// LkkPadType 枚举类型,字符串填充类型
	LkkPadType uint8

	// FileFilter 文件过滤函数
	FileFilter func(string) bool

	// CallBack 回调执行函数,无参数且无返回值
	CallBack func()
)

const (
	// Version 版本号
	Version = "0.0.4"

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

	// FILE_TYPE_ANY 文件类型-任意
	FILE_TYPE_ANY LkkFileType = 0
	// FILE_TYPE_LINK 文件类型-链接文件
	FILE_TYPE_LINK LkkFileType = 1
	// FILE_TYPE_REGULAR 文件类型-常规文件(不包括链接)
	FILE_TYPE_REGULAR LkkFileType = 2
	// FILE_TYPE_COMMON 文件类型-普通文件(包括常规和链接)
	FILE_TYPE_COMMON LkkFileType = 3

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

	// PAD_LEFT 左侧填充
	PAD_LEFT LkkPadType = 0
	// PAD_RIGHT 右侧填充
	PAD_RIGHT LkkPadType = 1
	// PAD_BOTH 两侧填充
	PAD_BOTH LkkPadType = 2

	//默认浮点数精确小数位数
	FLOAT_DECIMAL = 10

	//AuthCode 动态密钥长度,须<32
	DYNAMIC_KEY_LEN = 8

	//检查连接超时的时间
	CHECK_CONNECT_TIMEOUT = time.Second * 5

	// 正则模式-全中文
	PATTERN_CHINESE_ALL = "^[\u4e00-\u9fa5]+$"

	// 正则模式-中文名称
	PATTERN_CHINESE_NAME = "^[\u4e00-\u9fa5][.•·\u4e00-\u9fa5]{0,30}[\u4e00-\u9fa5]$"

	// 正则模式-多字节字符
	PATTERN_MULTIBYTE = "[^\x00-\x7F]"

	// 正则模式-ASCII字符
	PATTERN_ASCII = "^[\x00-\x7F]+$"

	// 正则模式-全角字符
	PATTERN_FULLWIDTH = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"

	// 正则模式-半角字符
	PATTERN_HALFWIDTH = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"

	// 正则模式-词语,不以下划线开头的中文、英文、数字、下划线
	PATTERN_WORD = "^[a-zA-Z0-9\u4e00-\u9fa5][a-zA-Z0-9_\u4e00-\u9fa5]+$"

	// 正则模式-浮点数
	PATTERN_FLOAT = `^(-?\d+)(\.\d+)?`

	// 正则模式-邮箱
	PATTERN_EMAIL = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

	// 正则模式-大陆手机号
	PATTERN_MOBILECN = `^1[3-9]\d{9}$`

	// 正则模式-固定电话
	PATTERN_TEL_FIX = `^(010|02\d{1}|0[3-9]\d{2})-\d{7,9}(-\d+)?$`

	// 正则模式-400或800
	PATTERN_TEL_4800 = `^[48]00\d?(-?\d{3,4}){2}$`

	// 正则模式-座机号(固定电话或400或800)
	PATTERN_TELEPHONE = `(` + PATTERN_TEL_FIX + `)|(` + PATTERN_TEL_4800 + `)`

	// 正则模式-电话(手机或固话)
	PATTERN_PHONE = `(` + PATTERN_MOBILECN + `)|(` + PATTERN_TEL_FIX + `)`

	// 正则模式-日期时间
	PATTERN_DATETIME = `^[0-9]{4}(|\-[0-9]{2}(|\-[0-9]{2}(|\s+[0-9]{2}(|:[0-9]{2}(|:[0-9]{2})))))$`

	// 正则模式-身份证号码,18位或15位
	PATTERN_CREDIT_NO = `(^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`

	// 正则模式-小写英文
	PATTERN_ALPHA_LOWER = `^[a-z]+$`

	// 正则模式-大写英文
	PATTERN_ALPHA_UPPER = `^[A-Z]+$`

	// 正则模式-字母和数字
	PATTERN_ALPHA_NUMERIC = `^[a-zA-Z0-9]+$`

	// 正则模式-十六进制颜色
	PATTERN_HEXCOLOR = `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`

	// 正则模式-RGB颜色
	PATTERN_RGBCOLOR = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"

	// 正则模式-全空白字符
	PATTERN_WHITESPACE_ALL = "^[[:space:]]+$"

	// 正则模式-带空白字符
	PATTERN_WHITESPACE_HAS = ".*[[:space:]]"

	// 正则模式-连续空白符
	PATTERN_WHITESPACE_DUPLICATE = `[[:space:]]{2,}|[\s\p{Zs}]{2,}`

	// 正则模式-base64字符串
	PATTERN_BASE64 = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"

	// 正则模式-base64编码图片
	PATTERN_BASE64_IMAGE = `^data:\s*(image|img)\/(\w+);base64`

	// 正则模式-html标签
	PATTERN_HTML_TAGS = `<(.|\n)*?>`

	// 正则模式-DNS名称
	PATTERN_DNSNAME = `^([a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9_]{1}[a-zA-Z0-9_-]{0,62})*[\._]?$`

	// 正则模式-MD5
	PATTERN_MD5 = `^(?i)([0-9a-h]{32})$`

	// 正则模式-SHA1
	PATTERN_SHA1 = `^(?i)([0-9a-h]{40})$`

	// 正则模式-SHA256
	PATTERN_SHA256 = `^(?i)([0-9a-h]{64})$`

	// 正则模式-SHA512
	PATTERN_SHA512 = `^(?i)([0-9a-h]{128})$`
)

var (
	// Kuptime 当前服务启动时间
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

	// KEncr utilities
	KEncr LkkEncrypt

	// KDbug utilities
	KDbug LkkDebug

	// KPrivCidrs 私有网段的CIDR数组
	KPrivCidrs []*net.IPNet

	// KDelimiter 本库自定义分隔符
	KDelimiter = "$@#KSYS#@$"

	// 身份证区域
	CreditArea = map[string]string{"11": "北京", "12": "天津", "13": "河北", "14": "山西", "15": "内蒙古", "21": "辽宁", "22": "吉林", "23": "黑龙江", " 31": "上海", "32": "江苏", "33": "浙江", "34": "安徽", "35": "福建", "36": "江西", "37": "山东", "41": "河南", "42": "湖北", "43": "湖南", "44": "广东", "45": "广西", "46": "海南", "50": "重庆", "51": "四川", "52": "贵州", "53": "云南", "54": "西藏", "61": "陕西", "62": "甘肃", "63": "青海", "64": "宁夏", "65": "新疆", "71": "台湾", "81": "香港", "82": "澳门", "91": "国外"}

	// html抽取文本要排除的标签
	TextHtmlExcludeTags = []string{"head", "title", "img", "form", "textarea", "input", "select", "button", "iframe", "script", "style", "option"}

	// 已编译的正则
	RegFormatDir             = regexp.MustCompile(`[\/]{2,}`) //连续的"//"或"\\"或"\/"或"/\"
	RegChineseAll            = regexp.MustCompile(PATTERN_CHINESE_ALL)
	RegChineseName           = regexp.MustCompile(PATTERN_CHINESE_NAME)
	RegWord                  = regexp.MustCompile(PATTERN_WORD)
	RegMultiByte             = regexp.MustCompile(PATTERN_MULTIBYTE)
	RegFullWidth             = regexp.MustCompile(PATTERN_FULLWIDTH)
	RegHalfWidth             = regexp.MustCompile(PATTERN_HALFWIDTH)
	RegFloat                 = regexp.MustCompile(PATTERN_FLOAT)
	RegEmail                 = regexp.MustCompile(PATTERN_EMAIL)
	RegMobilecn              = regexp.MustCompile(PATTERN_MOBILECN)
	RegTelephone             = regexp.MustCompile(PATTERN_TELEPHONE)
	RegPhone                 = regexp.MustCompile(PATTERN_PHONE)
	RegDatetime              = regexp.MustCompile(PATTERN_DATETIME)
	RegCreditno              = regexp.MustCompile(PATTERN_CREDIT_NO)
	RegAlphaLower            = regexp.MustCompile(PATTERN_ALPHA_LOWER)
	RegAlphaUpper            = regexp.MustCompile(PATTERN_ALPHA_UPPER)
	RegAlphaNumeric          = regexp.MustCompile(PATTERN_ALPHA_NUMERIC)
	RegHexcolor              = regexp.MustCompile(PATTERN_HEXCOLOR)
	RegRgbcolor              = regexp.MustCompile(PATTERN_RGBCOLOR)
	RegWhitespace            = regexp.MustCompile(`\s`)
	RegWhitespaceAll         = regexp.MustCompile(PATTERN_WHITESPACE_ALL)
	RegWhitespaceHas         = regexp.MustCompile(PATTERN_WHITESPACE_HAS)
	RegWhitespaceDuplicate   = regexp.MustCompile(PATTERN_WHITESPACE_DUPLICATE)
	RegBase64                = regexp.MustCompile(PATTERN_BASE64)
	RegBase64Image           = regexp.MustCompile(PATTERN_BASE64_IMAGE)
	RegHtmlTag               = regexp.MustCompile(PATTERN_HTML_TAGS)
	RegDNSname               = regexp.MustCompile(PATTERN_DNSNAME)
	RegUrlBackslashDuplicate = regexp.MustCompile(`([^:])[\/]{2,}`) //url中连续的"//"或"\\"或"\/"或"/\"
	RegMd5                   = regexp.MustCompile(PATTERN_MD5)
	RegSha1                  = regexp.MustCompile(PATTERN_SHA1)
	RegSha256                = regexp.MustCompile(PATTERN_SHA256)
	RegSha512                = regexp.MustCompile(PATTERN_SHA512)

	//	RegAscii                 = regexp.MustCompile(PATTERN_ASCII)

)
