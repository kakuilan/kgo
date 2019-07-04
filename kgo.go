package gohelper

type (
	// LkkFile is the receiver of file utilities
	LkkFile byte
	// LkkString is the receiver of string utilities
	LkkString byte
	// LkkArray is the receiver of array utilities
	LkkArray byte
	// LkkTime is the receiver of time utilities
	LkkTime byte
	// LkkConvert is the receiver of convert utilities
	LkkConvert byte

	// LkkFileCover 枚举类型,文件是否覆盖
	LkkFileCover int8
	// LkkFileTree 枚举类型,文件树查找类型
	LkkFileTree int8
)

const (
	// Version 版本号
	Version = "0.0.1"

	//UINT_MAX 无符号整型uint最大值
	UINT_MAX uint = ^uint(0)

	//UINT_MIN 无符号整型uint最小值
	UINT_MIN uint = 0

	//INT_MAX 有符号整型int最大值
	INT_MAX = int(^uint(0) >> 1)

	//INT_MIN 有符号整型int最小值
	INT_MIN = ^INT_MAX

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
)

var (
	// KFile utilities
	KFile LkkFile

	// KStr utilities
	KStr LkkString

	// KArr utilities
	KArr LkkArray

	// KTime utilities
	KTime LkkTime

	// KConv utilities
	KConv LkkConvert
)
