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
	// FileFilter 文件过滤函数
	FileFilter func(string) bool
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
)
