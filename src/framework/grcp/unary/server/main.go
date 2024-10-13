package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	//引入 proto 包并重命名包名，也可以使用 hello
	pb "github.com/henryzhuhr/hello-go/src/framework/grcp/unary/proto"
)

var port = flag.Int("port", 50051, "The server port")

// 定义服务，用于实现服务接口
type GreeterServer struct {
	// cannot use &GreeterServer{} (value of type *GreeterServer) as
	// hello.GreeterServer value in argument to hello.RegisterGreeterServer: *GreeterServer does not implement hello.GreeterServer
	// (missing method mustEmbedUnimplementedGreeterServer)
	pb.UnimplementedGreeterServer
}

// 实现 SayHello 方法
func (s *GreeterServer) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("Received a gRPC request, request parameters:", request)
	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}

func main() {
	flag.Parse()
	// 监听端口，创建一个 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening at port:", *port)

	// 创建一个 gRPC 服务
	grpcServer := grpc.NewServer()

	// 将自己实现的 GreeterServer 服务注册到 gRPC 服务中
	pb.RegisterGreeterServer(grpcServer, &GreeterServer{})

	// 往 gRPC 服务端注册反射服务
	reflection.Register(grpcServer)

	// 启动 gRPC 服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
