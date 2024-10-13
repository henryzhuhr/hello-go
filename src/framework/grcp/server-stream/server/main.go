package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"net"

	pb "github.com/henryzhuhr/hello-go/src/framework/grcp/server-stream/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

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

var port = flag.Int("port", 50051, "The server port")

func main() {
	flag.Parse()
	// 监听端口，创建一个 gRPC 服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Server listening at port:", *port)

	// 创建一个 gRPC 服务
	// 设置单次接受和发送消息的最大值
	grpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*4),   // 1024*1024*4 bytes (4M)
		grpc.MaxSendMsgSize(math.MaxInt32), // 2147483647 bytes (2G)
	)

	// 将自己实现的 StreamServer 服务注册到 gRPC 服务中
	pb.RegisterValueStreamServer(grpcServer, &StreamServer{})

	// 往 gRPC 服务端注册反射服务
	reflection.Register(grpcServer)

	// 启动 gRPC 服务
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
