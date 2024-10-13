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
- **一元RPC (Unary RPC)**：单一请求-单一响应的模式。客户端发起请求，服务器返回一个响应，和调用函数一样
- **服务器流式 (Server-side streaming RPC)**：客户端发起请求，服务器返回一个流，客户端从流中读取数据，直到流中没有任何消息，适用于服务器返回的数据量较大，客户端无法一次性接收的场景。
- **客户端流式 (Client-side streaming RPC)**：客户端发起一个流，服务器返回一个响应，客户端继续发送数据，适用于客户端发送的数据量较大，服务器无法一次性接收的场景。
- **双向流式 (Bidirectional streaming RPC)**：客户端和服务器之间建立一个双向流，客户端和服务器可以同时发送和接收数据，适用于客户端和服务器需要同时发送和接收数据的场景。

### 一元 RPC 

一元RPC (Unary RPC)，指客户端发起请求，服务器返回一个响应，和调用函数一样。这是最简单的 RPC 模式，客户端发起请求，服务器返回一个响应，适用于请求和响应数据量较小的场景。

参考官方的 [gRPC Hello World](https://github.com/grpc/grpc-go/tree/master/examples/helloworld) 示例，需要完成以下步骤以完成一个简单的 gRPC 服务：
1. 定义服务接口：需要编写 `.proto` 文件定义服务接口和数据结构，然后使用 `protoc` 编译器生成 Go 代码。
2. 编写服务端代码
3. 编写客户端代码
4. 分别运行服务端和客户端

#### 定义服务接口

RPC 服务需要从编写基于 protobuf 语法规范的 `.proto` 文件开始，该文件规定了服务接口形式、交互消息格式，而不涉及任何有关服务具体功能代码的编写，然后使用 `protoc` 编译器生成对应语言的代码。

protobuf 语法具体参考 [Protobuf 语法指南](https://colobu.com/2015/01/07/Protobuf-language-guide/)

```bash
src/framework/grcp/unary
└── proto
    └── hello.proto # 定义服务接口
```

例如 `hello.proto` 文件定义了一个简单的服务接口：
```proto
// 指定 protobuf 的版本，proto3 是最新的语法版本
syntax = "proto3";
// 包名，可以不需要
package helloworld;
// 生成go代码相关的选项，是必须的，否则无法生成go代码
option go_package = "./;helloworld";

// 定义 Greet 服务，包含一个 SayHello 方法
service Greeter {
  // 定义 SayHello 方法，接收一个 HelloRequest 消息，返回一个 HelloReply 消息
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// 使用 message 定义消息格式以及字段，可以看成 Go 语言中的结构体
// 定义请求数据结构
message HelloRequest {
  string name = 1;  // string 类型的字段 name, 序号为1
}
// 定义响应数据结构
message HelloReply {
  string message = 1; // string 类型的字段 message, 序号为1
}
```

`syntax = "proto3";` 指定了使用的 protobuf 版本，proto3 是最新的语法版本，如果不指定版本则可能会报错
```proto
syntax = "proto3";
```

`option go_package` 指定生成的Go代码在你项目中的导入路径，例如下面的代码表示生成的Go代码在当前目录 `.` 下，包名为 `proto`：
```proto
option go_package = "./;proto";
```

用 `service` 定义服务接口，并在其中定义一个 `rpc` 方法，`rpc` 方法接收一个消息类型，返回一个消息类型：
```proto
service ServiceName {
  rpc MethodName (RequestType) returns (ResponseType) {}
}
```

`message` 定义消息格式以及字段，可以看成 Go 语言中的结构体：
```proto
message MessageName {
  FieldType FieldName = 1; // 字段类型 字段名 = 序号
  FieldType FieldName = 2;
}
```


随后就可以使用 `protoc` 编译器生成 Go 代码：
```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  src/framework/grcp/unary/proto/hello.proto
```

执行上述命令生成两个文件 `hello.pb.go` 和 `hello_grpc.pb.go`，分别对应 `.proto` 文件中定义的**消息类型**和**服务接口**：
```bash
└── proto
    ├── hello_grpc.pb.go  # 生成的 gRPC 服务代码
    ├── hello.pb.go       # 生成的 Go 代码
    └── hello.proto       # 定义服务接口
```

上面的命令表示：
- `--go_out` 指定 `xxpb.go` 文件的生成路径
- `--go-grpc_out` 指定 gRPC 服务定义生成的 Go 代码 `xx_grpc.pb.go` 文件的生成路径
- `--go_opt` 是一个选项，例如可以指定生成的 Go 代码的包路径与源文件的相对路径一致，即 `paths=source_relative`
- `--go-grpc_opt` 则是 `xx_grpc.pb.go` 文件的选项
- `src/framework/grcp/unary/proto/hello.proto`：输入 `proto` 文件路径

生成的 `hello.pb.go` 和 `hello_grpc.pb.go` 文件中包含了定义的消息类型和服务接口的 Go 代码，可以在服务端和客户端代码中引用。

其中 `hello_grpc.pb.go` 文件包括了客户端和服务器的接口 `GreeterClient` 和 `GreeterServer`，以及服务接口的实现 `UnimplementedGreeterServer`：
```go
// GreeterClient is the client API for Greeter service.
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
// 定义 Greet 服务，包含一个 SayHello 方法
type GreeterClient interface {
	// 定义 SayHello 方法，接收一个 HelloRequest 消息，返回一个 HelloReply 消息
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer for forward compatibility.
// 定义 Greet 服务，包含一个 SayHello 方法
type GreeterServer interface {
	// 定义 SayHello 方法，接收一个 HelloRequest 消息，返回一个 HelloReply 消息
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}
```

接下来需要编写服务端和客户端代码，实现服务接口。

#### 编写服务端代码


```bash
src/framework/grcp/unary
├── proto
│   ├── hello_grpc.pb.go  # 生成的 gRPC 服务代码
│   ├── hello.pb.go       # 生成的 Go 代码
│   └── hello.proto       # 定义服务接口
└── server/main.go  # 服务端启动代码
```

`server.go` 文件中实现了服务接口 `GreeterServer` 并启动一个 gRPC 服务：


首先是 实现服务接口部分
```go
package main
import (
  // 其他的一些包
	//引入 proto 包并重命名包名，也可以使用 hello
	pb "github.com/henryzhuhr/hello-go/src/framework/grcp/unary/proto"
)
// 定义服务，用于实现服务接口
type GreeterServer struct {
  pb.UnimplementedGreeterServer
}
// 实现 SayHello 方法
func (s *GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("收到一个 grpc 请求，请求参数：", request)
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}
```

`GreeterServer` 的结构体中包含一个 `pb.UnimplementedGreeterServer`，表示未实现的服务接口，如果没有实现服务接口，编译器会报错。

然后在 `main` 函数中启动一个 gRPC 服务：
```go
func main() {
	// 监听端口，创建一个 gRPC 服务
	lis, err := net.Listen("tcp", 50051)
	if err != nil { log.Fatalf("failed to listen: %v", err) }

	// 创建一个 gRPC 服务
	grpcServer := grpc.NewServer()
	// 也可以设置单次接受和发送消息的最大值
	// grpcServer := grpc.NewServer(
	// 	grpc.MaxRecvMsgSize(1024*1024*4),   // 1024*1024*4 bytes (4M)
	// 	grpc.MaxSendMsgSize(math.MaxInt32), // 2147483647 bytes (2G)
	// )

	// 将自己实现的 GreeterServer 服务注册到 gRPC 服务中
	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	// 启动 gRPC 服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

::: details 服务端代码 `src/framework/grcp/unary/server/main.go` 完整代码
<<< @/../src/framework/grcp/unary/server/main.go
:::

#### 编写客户端代码

```bash
src/framework/grcp/unary
├── proto
│   ├── hello_grpc.pb.go
│   ├── hello.pb.go
│   └── hello.proto
├── client/main.go
└── server/main.go
```

`client.go` 文件中实现了一个简单的 gRPC 客户端，向服务端发送请求并接收响应：
```go
var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)
func main() {
	flag.Parse()
	// 连接 RPC 服务器
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil { log.Fatalf("did not connect: %v", err) }
	// 延迟关闭连接
	defer conn.Close()

	// 初始化一个 Greeter 客户端
	c := pb.NewGreeterClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil { log.Fatalf("could not greet: %v", err) }
	// 打印服务的返回的消息
	log.Printf("Greeting: %s", r.GetMessage())
}
```


::: details 客户端代码 `src/framework/grcp/unary/client/main.go` 完整代码
<<< @/../src/framework/grcp/unary/client/main.go
:::

#### 运行服务

首先在一个终端中启动服务端：
```bash
go run src/framework/grcp/unary/server/main.go
```
再另一个新终端中启动客户端：
```bash
go run src/framework/grcp/unary/client/main.go
```

然后客户端会输出服务端返回的消息，就像是调用一个函数一样：
```bash
# 服务端的终端输出
Received a gRPC request, request parameters: name:"world"
# 客户端的终端输出
Greeting: Hello world
```


### 服务器流式 RPC

服务器流式 RPC (Server-side streaming RPC)，客户端在其中向服务器发送请求，并获取流以读取回一系列消息。客户端从返回的流中读取，直到没有更多消息为止。gRPC保证在单个RPC调用中对消息进行排序。


#### 定义服务接口

```proto
syntax = "proto3";
option go_package = "./;server_stream_rpc";
// 定义发送请求消息结构
message StreamRequest{
    string data = 1;
}
// 定义流式响应消息结构
message StreamResponse{
    string stream_value = 1;
}
// 服务端流式rpc，只要在响应数据前加stream（可定义多个服务,每个服务可定义多个接口）
service StreamServer{
    rpc ListValue(StreamRequest) returns (stream StreamResponse){};
}
```

生成 Go 代码：
```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  src/framework/grcp/server-stream/proto/server_stream.proto
```

#### 编写服务端代码

服务端代码需要实现的是 `StreamServer` 服务接口中的 `GetValueStream` 方法，该方法接收一个 `StreamRequest` 消息，返回一个 `StreamResponse` 流。
```go
// StreamServer 服务端定义我们的服务
type StreamServer struct {
	pb.UnimplementedValueStreamServer
}
// 实现 proto 中定义的 GetValueStream 方法
func (s *StreamServer) GetValueStream(
	req *pb.StreamRequest, // 接收客户端请求
	stream pb.ValueStream_GetValueStreamServer, // 服务端流式处理
) error {
  // 循环发送消息
	for i := 0; i < 5; i++ {
		res := fmt.Sprintf("Hello %s, %d", req.Data, i)
		err := stream.Send(&pb.StreamResponse{StreamValue: res})
		if err != nil {
			log.Printf("send error: %v", err)
			return err
		}
		log.Printf("send: %s", res)
	}
	return nil
}
```
上述只是实现服务的部分，启动服务的代码与前一个例子类似，可以查看完整代码：

::: details 客户端代码 `src/framework/grcp/server-stream/server/main.go` 完整代码
<<< @/../src/framework/grcp/server-stream/server/main.go
:::

#### 编写客户端代码


这里只给出调用部分的代码
```go
func main() {
	// ...
	// 初始化一个 NewValueStream 客户端
	rpcClient := pb.NewValueStreamClient(conn)

	// 调用 GetValueStream 接口，发送一条消息
	req := &pb.StreamRequest{Data: *name}
	rpcStream, err := rpcClient.GetValueStream(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		resp, err := rpcStream.Recv()
		// 如果流已经结束，退出循环
		if err == io.EOF {
			log.Println("Stream response is EOF")
			break
		}
		if err != nil {
			log.Fatalf("Cannot receive stream response: %v", err)
		}
		// 打印服务的返回的消息
		log.Printf("Stream value: %s", resp.StreamValue)
	}
}
```

完整代码：

::: details 客户端代码 `src/framework/grcp/server-stream/client/main.go` 完整代码
<<< @/../src/framework/grcp/server-stream/client/main.go
:::


#### 运行服务

首先在一个终端中启动服务端：
```bash
go run src/framework/grcp/server-stream/server/main.go
```
再另一个新终端中启动客户端：
```bash
go run src/framework/grcp/server-stream/client/main.go
```

然后客户端会输出服务端返回的消息，就像是调用一个函数一样：
```bash
# 服务端的终端输出
send: Hello world, 0
...
send: Hello world, 5
# 客户端的终端输出
Stream value: Hello world, 0
...
Stream value: Hello world, 5
```

  

### 客户端流式 RPC
客户端在其中编写消息序列，然后再次使用提供的流将其发送到服务器。客户端写完消息后，它将等待服务器读取消息并返回其响应。gRPC再次保证了在单个RPC调用中的消息顺序。

### 双向流式RPC
双方都使用读写流发送一系列消息。这两个流独立运行，因此客户端和服务器可以按照自己喜欢的顺序进行读写：例如，服务器可以在写响应之前等待接收所有客户端消息，或者可以先读取消息再写入消息，或读写的其他组合。每个流中的消息顺序都会保留。

## 参考

- [Protobuf 语法指南](https://colobu.com/2015/01/07/Protobuf-language-guide/)
- [GRPC (2): 四种通信模式‍ - Canghaimingue的文章 - 知乎](https://zhuanlan.zhihu.com/p/547806941)：以订单管理系统为例
