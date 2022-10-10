# kgo

k`s golang helper/library/utils  
golang 常用函数库/工具集,仅测试支持有限的64位系统.  
总共460多个通用的方法,涵盖字符串、数组、文件、时间、加密以及类型转换等操作.

### 文档

[![GoDoc](https://godoc.org/github.com/kakuilan/kgo?status.svg)](https://pkg.go.dev/github.com/kakuilan/kgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/kakuilan/kgo)](https://goreportcard.com/report/github.com/kakuilan/kgo)
[![Build Status](https://github.com/kakuilan/kgo/workflows/kgo-test/badge.svg)](https://github.com/kakuilan/kgo/actions)
[![codecov](https://codecov.io/gh/kakuilan/kgo/branch/master/graph/badge.svg)](https://codecov.io/gh/kakuilan/kgo)
[![Code Size](https://img.shields.io/github/languages/code-size/kakuilan/kgo.svg?style=flat-square)](https://github.com/kakuilan/kgo)
[![Starts](https://img.shields.io/github/stars/kakuilan/kgo.svg)](https://github.com/kakuilan/kgo)
[![Version](https://img.shields.io/github/v/tag/kakuilan/kgo)](https://img.shields.io/github/v/tag/kakuilan/kgo)

### 测试支持

- GO版本
    - 1.16.x
    - 1.17.x
    - 1.18.x
    - 1.19.x
- OS系统
    - ubuntu-latest
    - macos-latest
    - windows-latest

### 依赖第三方库

- github.com/json-iterator/go
- github.com/StackExchange/wmi

### 安装使用

安装

```shell script
go get -u github.com/kakuilan/kgo
```

引入

```go
import "github.com/kakuilan/kgo"
```

### 函数接收器

- *KFile* 为文件操作,如

```go
chk := KFile.IsExist(filename)
```

- *KStr* 为字符串操作,如

```go
res := KStr.Trim(" hello world ")
```

- *KNum* 为数值操作,如

```go
res := KNum.NumberFormat(123.4567890, 3, ".", "")
```

- *KArr* 为数组(切片/字典)操作,如

```go
mp := map[string]string{
"a": "aa",
"b": "bb",
}
chk := KArr.InArray("bb", mp)    
```

- *KTime* 为时间操作,如

```go
res, err := KTime.Str2Timestamp("2019-07-11 10:11:23")
```

- *KConv* 为类型转换操作,如

```go
res := KConv.ToStr(false)
```

- *KOS* 为系统和网络操作,如

```go
res, err := KOS.LocalIP()
```

- *KEncr* 为加密操作,如

```go
res, err := KEncr.PasswordHash([]byte("123456"))
```

- *KDbug* 为调试操作,如

```go
KDbug.DumpPrint(1.2)
```

具体函数请查看[godoc](https://pkg.go.dev/github.com/kakuilan/kgo),更多示例请参考*_test.go文件.

### 测试

```shell
#使用go mod
go mod tidy
go mod vendor

#单元测试
go test -race

#压测
time go test -bench=. -run=none
time go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="15s" -benchmem

#代码覆盖率
go test -cover #概览

go test -coverprofile=coverage.out #生成统计信息
go test -v -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out #查看统计信息
go tool cover -html=coverage.out #将统计信息转换为html

#性能分析
time go test -timeout 30m -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof profile.out
go tool pprof -http=192.168.1.2:8081 /usr/bin/dot profile.out
```

### 更新日志

详见[[Changelog]](/docs/changelog.md)

### 鸣谢

感谢[JetBrains](https://www.jetbrains.com/?from=kakuilan/kgo)的赞助.  
![JetBrains](testdata/jetbrains.svg)

