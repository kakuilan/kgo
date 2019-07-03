package gohelper

type (
	// LkkFile is the receiver of file utilities
	LkkFile byte
	// LkkStr is the receiver of string utilities
	LkkString byte
	// LkkArray is the receiver of array utilities
	LkkArray byte
	// LkkTime is the receiver of time utilities
	LkkTime byte

	// 枚举类型,文件是否覆盖
	LkkFileCover int8
	// 枚举类型,文件树查找类型
	LkkFileTree int8

)

const (
	// 版本号
	Version = "0.0.1"

	// 文件覆盖,允许
	FILE_COVER_ALLOW LkkFileCover = 1
	// 文件覆盖,忽略
	FILE_COVER_IGNORE LkkFileCover = 0
	// 文件覆盖,禁止
	FILE_COVER_DENY LkkFileCover = -1

	// 文件树,查找所有(包括目录和文件)
	FILE_TREE_ALL LkkFileTree = 3
	// 文件树,仅查找目录
	FILE_TREE_DIR LkkFileTree = 2
	// 文件树,仅查找文件
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
)