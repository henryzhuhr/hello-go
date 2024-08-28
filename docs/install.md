---
outline: deep
---
# Go 开发环境

## 安装

### 安装预编译版本

### 从源码编译

参考 [Installing Go from source](https://go.dev/doc/install/source)


## 配置
安装 Go 后，配置 GOPATH， `export GOPATH="$HOME/project/goproject"` ，开发时，代码存放在 `$GOPATH/src` 下， 工程经过 `go build`、`go install` 或`go get` 等指令后，会将下载的第三方包源代码文件放在 `$GOPATH/src` 目录下， 产生的二进制可执行文件放在 `$GOPATH/bin` 目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下


### GOPROXY 代理

配置国内代理 [Goproxy.cn](https://goproxy.cn)
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### GOPRIVATE 私有仓库



## VSCode 配置

VSCode 需要安装如下插件，可以在命令面板(`Ctrl+Shift+P`)中搜索 `Go: Install/Update Tools` 安装

- [`gopls`](https://github.com/golang/tools/blob/master/gopls/README.md#installation)
- [`gotests`](https://github.com/cweill/gotests)
- [`gomodifytags`](https://github.com/fatih/gomodifytags)
- [`impl`](https://github.com/josharian/impl)
- [`goplay`](https://github.com/haya14busa/goplay)
- [`dlv`](https://github.com/go-delve/delve)
- [`staticcheck`](https://github.com/dominikh/go-tools)

如果存在网络问题，可以手动使用下面命令手动安装，或者参照各个工具的手动安装手册，下载源码编译安装
```bash
go install golang.org/x/tools/gopls@latest
go get -u  github.com/cweill/gotests/...
go install github.com/fatih/gomodifytags@latest
go install github.com/josharian/impl@latest
# go get -u  github.com/haya14busa/goplay            # as a Library
go get -u  github.com/haya14busa/goplay/cmd/goplay # as a command line tool
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
```