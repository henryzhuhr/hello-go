---
outline: deep
---

# gRPC

## 简介

grpc 的全称是 Google Remote Procedure Call，是一种高性能、开源和通用的 RPC 框架，基于 HTTP/2 协议，支持多种语言。

Go 语言的 gRPC 实现是 [grpc-go](https://github.com/grpc/grpc-go)

## 安装

### grpc 包的安装

只需将以下导入添加到代码中，然后使用 `go mod tidy` 或 `go [build|run|test]` 时将自动获取依赖包：
```go
import "google.golang.org/grpc"
```

也可以使用以下命令安装：
```bash
go get -u google.golang.org/grpc
```

如果网络环境无法访问Google服务器时，可以使用 `go mod` 的替换功能，为 `golang.org` 软件包创建别名（需要 Go 模块支持）：
```bash
go mod edit -replace=google.golang.org/grpc=github.com/grpc/grpc-go@latest
go mod tidy
go mod vendor
go build -mod=vendor
```

### protoc

RPC 服务的定义是通过 Protocol Buffers（简称 Protobuf）来定义的，Protobuf 是一种轻便高效的结构化数据序列化方式，类似于 XML 或 JSON，通过编写 `.proto` 文件定义数据结构和服务接口，然后使用 protoc 编译器生成对应语言的代码。

protoc 的安装可以安装官网 [_Protocol Buffer Compiler Installation_](https://grpc.io/docs/protoc-installation/) 的命令：
```bash
# Linux, using apt or apt-get, for example:
apt install -y protobuf-compiler
# MacOS, using Homebrew:
brew install protobuf
```
也可以从 [protocolbuffers/protobuf 的 Github Release](https://github.com/protocolbuffers/protobuf/releases) 下载，然后解压到 PATH 环境变量中。

除了 protoc 编译器外，还需要安装生成 Go 代码的 protoc 插件 `protoc-gen-go` 和生成 gRPC 服务代码的插件 `protoc-gen-go-grpc`：
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

或者从 github 下载 [`protoc-gen-go`](https://github.com/golang/protobuf) 和 [`protoc-gen-go-grpc`](https://github.com/grpc/grpc-go/tree/master/cmd/protoc-gen-go-grpc)
```bash
go install github.com/golang/protobuf/protoc-gen-go@latest
go install github.com/grpc/grpc-go/cmd/protoc-gen-go-grpc@latest
```

然后需要将 `protoc-gen-go` 和 `protoc-gen-go-grpc` 可执行文件添加到 PATH 环境变量中，以便 `protoc` 编译器能够找到：
```bash
export PATH="$PATH:$(go env GOPATH)/bin"
```

MacOS 下可以使用 Homebrew 安装：
```bash
brew install protoct-gen-go
brew install protoct-gen-go-grpc
```

### proto VSCode 插件

在 VSCode 中编写 `.proto` 文件时，可以安装 `vscode-proto3` 插件，该插件提供了语法高亮、代码提示、错误检查等功能。该插件依赖于 clang-format，需要先安装 clang-format：
```bash
apt  install clang-format # Debian/Ubuntu
brew install clang-format # MacOS
```

<!-- https://www.jianshu.com/p/15d153a77d88 -->


## gRPC 四种通信模式

gRPC 提供了四种主要的通信模式：**普通模式**、**服务器流式**、**客户端流式**和**双向流式**。每种模式都有不同的特点和适用场景:
- **普通模式 (Unary RPC)**：单一请求-单一响应的模式。客户发起请求，并等待服务器响应，这是最简单的 RPC 模式，适用于请求和响应数据量较小的场景。
- **服务器流式 (Server-side streaming RPC)**：客户端发起请求，服务器返回一个流，客户端从流中读取数据，直到流中没有任何消息，适用于服务器返回的数据量较大，客户端无法一次性接收的场景。
- **客户端流式 (Client-side streaming RPC)**：客户端发起一个流，服务器返回一个响应，客户端继续发送数据，适用于客户端发送的数据量较大，服务器无法一次性接收的场景。
- **双向流式 (Bidirectional streaming RPC)**：客户端和服务器之间建立一个双向流，客户端和服务器可以同时发送和接收数据，适用于客户端和服务器需要同时发送和接收数据的场景。

