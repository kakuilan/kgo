### 本地文档
```sh
godoc -http=:6060

#查看
http://192.168.1.1:6060/pkg/github.com/kakuilan
```

### 生成markdown
```sh
go get github.com/robertkrimen/godocdown/godocdown
godocdown . > docs/v0.1.2.md
```