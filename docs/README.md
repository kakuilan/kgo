### 本地文档

```sh
godoc -http=:6060

#查看
http://192.168.1.1:6060/pkg/github.com/kakuilan
```

### 生成markdown

```sh
go get -d github.com/robertkrimen/godocdown/godocdown
godocdown . > docs/v0.3.6.md
```

### 安装依赖

```shell
go get github.com/StackExchange/wmi@v1.2.1
go get github.com/brianvoe/gofakeit/v6
go get github.com/json-iterator/go
go get github.com/stretchr/testify
go get golang.org/x/crypto
go get golang.org/x/net
go get golang.org/x/sys
go get golang.org/x/text
go get gopkg.in/yaml.v3
```