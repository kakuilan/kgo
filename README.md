# kgo
k`s golang helper/library/utils  
golang 函数库/工具集

### 压测
```shell
time go test -bench=. -run=none
time go test -v -bench=. -cpu=4 -benchtime="10s" -timeout="15s" -benchmem

#性能分析
time go test -bench=. -benchmem -memprofile memprofile.out -cpuprofile profile.out
go tool pprof profile.out
go tool pprof -http=":8081" [binary] [profile]
```

### 参考项目
- https://github.com/b3log/gulu
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
- https://github.com/keysonZZZ/kgo

- https://stackoverflow.com/questions/32482673/how-to-get-directory-total-size
- https://www.codesd.com/item/golang-how-to-get-the-total-size-of-the-directory.html
- https://socketloop.com/tutorials/golang-find-file-size-disk-usage-with-filepath-walk
- https://github.com/karrick/godirwalk
- http://xahlee.info/golang/golang_walk_dir.html
