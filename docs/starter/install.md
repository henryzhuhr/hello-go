---
outline: deep
---
# Go 开发环境

## 安装

### 安装预编译版本

### 从源码编译

参考 [Installing Go from source](https://go.dev/doc/install/source)


## 配置
<!-- 安装 Go 后，配置 GOPATH， `export GOPATH="$HOME/project/goproject"` ，开发时，代码存放在 `$GOPATH/src` 下， 工程经过 `go build`、`go install` 或`go get` 等指令后，会将下载的第三方包源代码文件放在 `$GOPATH/src` 目录下， 产生的二进制可执行文件放在 `$GOPATH/bin` 目录下，生成的中间缓存文件会被保存在 `$GOPATH/pkg` 下 -->

PATH 中包含 go 即可，可以通过 `go env` 查看 Go 的环境变量

```shell
go env
```

### GOPROXY 代理

配置国内代理 [Goproxy.cn](https://goproxy.cn)
```bash
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

### GOPRIVATE 私有仓库

Go 1.13 版本后，拉取私有仓库的方式发生了变化，需要配置环境变量，否则会报错

```shell
export GOPRIVATE=github.com/<username>
```
同时由于 `go get` 默认使用的是 https 协议，如果你的私有仓库是 ssh 协议，需要配置 `~/.gitconfig` 文件，添加如下配置

```shell
git config --global url."git@git.xxx.com".insteadOf "https://git.xxx.com/"
```

