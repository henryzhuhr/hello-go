---
outline: deep
---

# Go 工具链

Go 开发过程中依赖于一些工具链

- [`gopls`](https://github.com/golang/tools/blob/master/gopls/README.md)：Go 语言的 LSP 服务

```shell
go install golang.org/x/tools/gopls@latest
```


- [`gotests`](https://github.com/cweill/gotests)：生成测试代码

```shell
go install github.com/cweill/gotests/...@latest
```


- [`gomodifytags`](https://github.com/fatih/gomodifytags)：修改结构体标签

```shell
go install github.com/fatih/gomodifytags@latest
```


- [`impl`](https://github.com/josharian/impl)：生成接口实现

```shell
go install github.com/josharian/impl@latest
```


- [`goplay`](https://github.com/haya14busa/goplay)：在 Go Playground 运行代码

```shell
# as a Library
go get -u github.com/haya14busa/goplay
# as a command line tool
go get -u github.com/haya14busa/goplay/cmd/goplay
```


- [`dlv`](https://github.com/go-delve/delve)：Go 语言调试器

```shell
go install github.com/go-delve/delve/cmd/dlv@latest
```


- [`golangci-lint`](https://github.com/golangci/golangci-lint)：Go 语言代码静态检查工具



不建议使用 `go get` 安装 `golangci-lint`，因为它会安装最新的 `golangci-lint`，而不是最新的稳定版本。
```shell
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Ubuntu 用户可以使用以下命令安装 `golangci-lint`：
```shell
apt install golangci-lint
```


```shell
# 官方提示不要使用 -u
go get golang.org/x/tools/gopls@latest
go get github.com/cweill/gotests/...@latest
go get github.com/fatih/gomodifytags@latest
go get github.com/josharian/impl@latest
go get -u github.com/haya14busa/goplay/cmd/goplay
go get github.com/go-delve/delve/cmd/dlv@latest
go get github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```