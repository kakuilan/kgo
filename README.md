# kgo
k`s golang helper/library/utils  
golang 函数库/工具集

### 文档
[![GoDoc](https://godoc.org/github.com/kakuilan/kgo?status.svg)](https://godoc.org/github.com/kakuilan/kgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/kakuilan/kgo)](https://goreportcard.com/report/github.com/kakuilan/kgo)
[![Code Size](https://img.shields.io/github/languages/code-size/kakuilan/kgo.svg?style=flat-square)](https://github.com/kakuilan/kgo)
[![Starts](https://img.shields.io/github/stars/kakuilan/kgo.svg)](https://github.com/kakuilan/kgo)
[![Downloads](https://img.shields.io/github/downloads/kakuilan/kgo/total.svg)](https://github.com/kakuilan/kgo/releases)


### 测试
```shell
go test

#压测
time go test -bench=. -run=none
time go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="15s" -benchmem

#性能分析
time go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof profile.out
go tool pprof -http=":8081" [binary] [profile]
```

### 参考项目
- https://www.php2golang.com/
- https://github.com/b3log/gulu
- https://github.com/syyongx/php2go
- https://github.com/openset/php2go
- https://github.com/yioio/fun
- https://github.com/nutzam/zgo
- https://github.com/bitnami/gonit
- https://github.com/otiai10/copy
- https://github.com/polaris1119/goutils
- https://github.com/LyricTian/lib
- https://github.com/antongulenko/golib
- https://github.com/bocajim/helpers
- https://github.com/divinerapier/go-tools
- https://github.com/elimisteve/fun
- https://github.com/emirozer/go-helpers
- https://github.com/evilsocket/islazy
- https://github.com/fatih-yavuz/go-helpers
- https://github.com/fingerQin/gophp
- https://github.com/hhxsv5/go-helpers
- https://github.com/idoall/TokenExchangeCommon/tree/master/commonutils
- https://github.com/jiazhoulvke/goutil
- https://github.com/jimmykuu/webhelpers
- https://github.com/kooksee/cmn
- https://github.com/kooksee/common
- https://github.com/leifengyao/go2php
- https://github.com/leizongmin/go-utils
- https://github.com/lets-go-go/helper
- https://github.com/mylxsw/go-toolkit
- https://github.com/nletech/go-func
- https://github.com/orivil/helper
- https://github.com/relunctance/goutils
- https://github.com/seiflotfy/do
- https://github.com/shuangdeyu/helper_go
- https://github.com/sohaha/zlsgo
- https://github.com/stephanbaker/go-simpletime
- https://github.com/vence722/convert
- https://github.com/jinzhu/now

### 其他库
- https://github.com/lalamove/konfig
- https://github.com/jinzhu/configor
- https://github.com/denisbrodbeck/machineid
- github.com/karrick/godirwalk
- https://github.com/keysonZZZ/kgo

### 包
file
string
time
os
array
number
url
convert


