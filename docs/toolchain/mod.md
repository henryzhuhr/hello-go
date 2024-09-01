---
outline: deep
---
# Go Modules 包管理

## Go Modules 产生原因

早期 Go 语言的包管理是通过 `GOPATH` 环境变量来实现的，所有的第三方包都会下载到 `GOPATH` 目录下，这样就会导致不同项目之间的包版本冲突，而且也不方便管理。


在没有 Go Modules 机制之前，Go 项目的包管理是通过 `GOPATH` 环境变量来实现的，开发的项目都会放在 `GOPATH` 目录下，所有的第三方包也会下载到 `GOPATH` 目录下，这样就会导致不同项目之间的包版本冲突，而且也不方便管理。

为了解决上述的问题，从 Go 1.11 [开始引入了 Go Modules 机制](https://go.dev/doc/go1.11#modules)，通过环境变量 `GO111MODULE` 启用。Go 1.13 版本后，[Go Modules](https://go.dev/ref/mod) 成为默认的包管理工具，并通过 `go.mod` 文件来管理项目的依赖。

> Go Modules 的使用可以参考官方文章 [_Using Go Modules_](https://go.dev/blog/using-go-modules)

## Go Modules 简介

Go Modules 是 Go 1.11 版本引入的包管理工具，可以通过环境变量 `GO111MODULE` 设置，有三个值 `off`、`on` 和 `auto`(default)，也可以通过 `go env -w` 命令设置

```shell
go env -w GO111MODULE=on
```

Go 1.13 版本后，Go Modules 成为默认的包管理工具，通过 `go mod` 命令来管理项目的依赖，因此不需要设置 `GO111MODULE` 环境变量。

## Go Mod 命令


```shell
go mod
# Go mod provides access to operations on modules.

# Note that support for modules is built into all the go commands,
# not just 'go mod'. For example, day-to-day adding, removing, upgrading,
# and downgrading of dependencies should be done using 'go get'.
# See 'go help modules' for an overview of module functionality.

# Usage:

#         go mod <command> [arguments]

# The commands are:

#         download    download modules to local cache
#         edit        edit go.mod from tools or scripts
#         graph       print module requirement graph
#         init        initialize new module in current directory
#         tidy        add missing and remove unused modules
#         vendor      make vendored copy of dependencies
#         verify      verify dependencies have expected content
#         why         explain why packages or modules are needed

# Use "go help mod <command>" for more information about a command.
```


## 使用 Go Mod 创建项目并管理

创建 modules 的命令
```shell
go mod init github.com/<username>/<projectname>
```

> `go.mod` 文件一旦创建后，它的内容将会被 go toolchain 全面掌控。go toolchain 会在各类命令执行时，比如 `go get`、`go build`、`go mod` 等修改和维护 `go.mod` 文件。



拉取私有仓库

Go 1.13 版本后，拉取私有仓库的方式发生了变化，需要配置环境变量，否则会报错

```shell
export GOPRIVATE=github.com/<username>
```
同时由于 `go get` 默认使用的是 https 协议，如果你的私有仓库是 ssh 协议，需要配置 `~/.gitconfig` 文件，添加如下配置

```shell
git config --global url."git@git.xxx.com".insteadOf "https://git.xxx.com/"
```