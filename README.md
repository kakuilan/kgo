# kgo
k`s golang helper/library/utils  
golang 函数库/工具集,仅测试支持64位linux

### 文档
[![GoDoc](https://godoc.org/github.com/kakuilan/kgo?status.svg)](https://godoc.org/github.com/kakuilan/kgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/kakuilan/kgo)](https://goreportcard.com/report/github.com/kakuilan/kgo)
[![Build Status](https://travis-ci.org/kakuilan/kgo.svg?branch=master)](https://travis-ci.org/kakuilan/kgo)
[![Coverage Status](https://coveralls.io/repos/github/kakuilan/kgo/badge.svg?branch=master)](https://coveralls.io/github/kakuilan/kgo?branch=master)
[![Code Size](https://img.shields.io/github/languages/code-size/kakuilan/kgo.svg?style=flat-square)](https://github.com/kakuilan/kgo)
[![Starts](https://img.shields.io/github/stars/kakuilan/kgo.svg)](https://github.com/kakuilan/kgo)
[![Downloads](https://img.shields.io/github/downloads/kakuilan/kgo/total.svg)](https://github.com/kakuilan/kgo/releases)


### 测试
```shell
go test

#清理包
go mod tidy

#代码覆盖率
go test -cover #概览

go test -coverprofile=coverage.out #生成统计信息
go test -v -covermode=count -coverprofile=coverage.out
go tool cover -func=coverage.out #查看统计信息
go tool cover -html=coverage.out #将统计信息转换为html


#压测
time go test -bench=. -run=none
time go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="15s" -benchmem

#性能分析
time go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof profile.out
go tool pprof -http=192.168.56.10:8081 /usr/bin/dot profile.out
```

### 参考项目
- https://www.php2golang.com/
- https://github.com/b3log/gulu -x
- https://github.com/syyongx/php2go
- https://github.com/openset/php2go
- https://github.com/yioio/fun
- https://github.com/henrylee2cn/goutil
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
- https://github.com/thinkeridea/go-extend
- https://github.com/lalamove/nui
- https://github.com/go-eyas/toolkit
- https://github.com/hwholiday/learning_tools
- https://github.com/nothollyhigh/kiss
- https://github.com/thinkeridea/go-extend


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
validate
encrypt


### 临时
https://www.calhoun.io/creating-random-strings-in-go/
http://ju.outofmemory.cn/entry/221647
https://www.golangnote.com/topic/90.html
https://stackoverflow.com/questions/38554353/how-to-check-if-a-string-only-contains-alphabetic-characters-in-go
https://socketloop.com/tutorials/golang-regular-expression-alphanumeric-underscore

encrypt
https://www.ctolib.com/topics-4140.html
https://www.cnblogs.com/lavin/p/5373188.html
https://github.com/sinlov/fastEncryptDecode
http://www.voidcn.com/article/p-pmecifun-bod.html
https://blog.csdn.net/yue7603835/article/details/73395580
https://gist.github.com/yinheli/3370e0e901329b639be4
https://www.jishuwen.com/d/2PrM#tuit
https://stackoverflow.com/questions/24072026/golang-aes-ecb-encryption
http://ciika.com/2016/11/golang-aes-demo/
http://wsztrush.com/2017/08/28/golang-first/
http://www.lampnick.com/php/728
https://segmentfault.com/a/1190000004151272
https://www.tech1024.com/original/3015.html
https://github.com/thinkoner/openssl
https://heroims.github.io/2018/11/21/AES-128-CBC%20Base64%E5%8A%A0%E5%AF%86%E2%80%94%E2%80%94OC,Java,Golang%E8%81%94%E8%B0%83/
https://smalltowntechblog.wordpress.com/2014/12/28/%E5%A6%82%E4%BD%95%E8%AE%93-aes-%E5%9C%A8-golang-%E8%88%87-androidjava-%E5%BE%97%E5%88%B0%E4%B8%80%E8%87%B4%E7%9A%84%E5%8A%A0%E8%A7%A3%E5%AF%86%E7%B5%90%E6%9E%9C/comment-page-1/
https://www.itread01.com/content/1547577396.html
https://stackoverflow.com/questions/18817336/golang-encrypting-a-string-with-aes-and-base64

map merge
https://github.com/imdario/mergo
https://gist.github.com/wuriyanto48/d35d7e0d322cb08a567a5305a41732dd
https://maxbittker.com/merging-maps
https://www.golangnote.com/topic/209.html
https://stackoverflow.com/questions/12172215/merging-maps-in-go

pool
https://github.com/gobwas/pool

json库
https://github.com/json-iterator/go
https://github.com/tidwall/gjson
https://github.com/mailru/easyjson
https://github.com/buger/jsonparser
