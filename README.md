# kgo
k`s golang helper/library/utils  
golang 函数库/工具集,仅测试支持64位linux

### 文档
[![GoDoc](https://godoc.org/github.com/kakuilan/kgo?status.svg)](https://godoc.org/github.com/kakuilan/kgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/kakuilan/kgo)](https://goreportcard.com/report/github.com/kakuilan/kgo)
[![Build Status](https://travis-ci.org/kakuilan/kgo.svg?branch=master)](https://travis-ci.org/kakuilan/kgo)
[![codecov](https://codecov.io/gh/kakuilan/kgo/branch/master/graph/badge.svg)](https://codecov.io/gh/kakuilan/kgo)
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
- https://github.com/openset/php2go -x
- https://github.com/yioio/fun  -x
- https://github.com/henrylee2cn/goutil -x
- https://github.com/nutzam/zgo -x
- https://github.com/bitnami/gonit  -x
- https://github.com/otiai10/copy   -x
- https://github.com/polaris1119/goutils    -x
- https://github.com/LyricTian/lib  -x
- https://github.com/antongulenko/golib -x
- https://github.com/bocajim/helpers    -x
- https://github.com/elimisteve/fun -x
- https://github.com/emirozer/go-helpers    -x
- https://github.com/evilsocket/islazy  -x
- https://github.com/fatih-yavuz/go-helpers -x
- https://github.com/fingerQin/gophp    -x
- https://github.com/hhxsv5/go-helpers  -x
- https://github.com/idoall/TokenExchangeCommon/tree/master/commonutils -x
- https://github.com/jiazhoulvke/goutil -x
- https://github.com/jimmykuu/webhelpers    -x
- https://github.com/kooksee/cmn utils
- https://github.com/kooksee/common
- https://github.com/leifengyao/go2php  -x
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
- https://github.com/huandu/xstrings
- https://github.com/bytedance/go-tagexpr
- https://github.com/pibigstar/go-demo


### 其他库
- https://github.com/lalamove/konfig
- https://github.com/jinzhu/configor
- https://github.com/denisbrodbeck/machineid
- github.com/karrick/godirwalk
- https://github.com/keysonZZZ/kgo
- https://github.com/gobwas/pool
- https://github.com/shirou/gopsutil
- https://github.com/sunmi-OS/gocore


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
net
debug


### 临时
https://github.com/jimmykuu/webhelpers/blob/master/gravatar.go
https://github.com/jimmykuu/webhelpers/blob/master/text.go



encrypt
https://github.com/henrylee2cn/goutil/blob/master/encrypt.go
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
https://github.com/bocajim/helpers/blob/master/crypt.go


json库
https://github.com/json-iterator/go
https://github.com/tidwall/gjson
https://github.com/mailru/easyjson
https://github.com/buger/jsonparser

pack/unpack
https://stackoverflow.com/questions/40182289/golang-equivalent-of-pythons-struct-pack-struct-unpack
https://gist.github.com/cloveryume/9a59e8d77f5836f11720#file-golang_struct_packed-go
https://github.com/lunixbochs/struc
https://github.com/roman-kachanovsky/go-binary-pack
https://learnku.com/articles/31460
https://studygolang.com/articles/2791
https://golangtc.com/t/55237dd8421aa9704b0000cb
https://stackoverflow.com/questions/8039552/byte-endian-convert-by-using-encoding-binary-in-go

检查端口绑定
https://golangnote.com/topic/241.html
https://www.jianshu.com/p/b4ce0794fa32

pid
https://github.com/bitnami/gonit/blob/master/utils/process.go
https://github.com/henrylee2cn/goutil/blob/master/pid_file.go

ping
https://github.com/bocajim/helpers/blob/master/ping.go

http/curl
https://github.com/mreiferson/go-httpclient
https://github.com/elimisteve/fun/blob/master/fetch.go
https://github.com/nareix/curl
https://github.com/andelf/go-curl
https://github.com/parnurzeal/gorequest
https://github.com/go-resty/resty
https://github.com/gojek/heimdall
https://github.com/dghubble/sling
https://github.com/h2non/gentleman
https://github.com/guonaihong/gout
https://github.com/levigross/grequests


queue
https://github.com/evilsocket/islazy/blob/master/async/queue.go

ismail
https://github.com/idoall/TokenExchangeCommon/blob/master/commonutils/checkmail/checkmail.go



debug
https://colobu.com/2018/11/03/get-function-name-in-go/
https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/
https://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces
https://github.com/rfyiamcool/stack_dump
https://www.jianshu.com/p/abbe6663b672
https://github.com/go-delve/delve

array sort
https://stackoverflow.com/questions/36122668/how-to-sort-struct-with-multiple-sort-parameters
https://yourbasic.org/golang/how-to-sort-in-go/
https://itimetraveler.github.io/2016/09/07/%E3%80%90Go%E8%AF%AD%E8%A8%80%E3%80%91%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B%E6%8E%92%E5%BA%8F%E5%92%8C%20slice%20%E6%8E%92%E5%BA%8F/
https://blog.csdn.net/chenbaoke/article/details/42340301
https://stackoverflow.com/questions/37695209/golang-sort-slice-ascending-or-descending
