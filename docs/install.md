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


配置代理
```bash
export GOPROXY="https://goproxy.cn,direct"
```

