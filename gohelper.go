package gohelper

type (
	// LkkFile is the receiver of file utilities
	LkkFile byte
	// 枚举类型,文件是否覆盖
	LkkFileCover int8
)

const (
	// 版本号
	Version = "0.0.1"
	// 允许文件覆盖
	FCOVER_ALLOW LkkFileCover = 1
	// 忽略文件覆盖
	FCOVER_IGNORE LkkFileCover = 0
	// 禁止文件覆盖
	FCOVER_DENY LkkFileCover = -1
)

var (
	// File utilities
	File LkkFile
)