package main

import (
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//引入 proto 包并重命名包名，也可以使用 hello
	pb "github.com/henryzhuhr/hello-go/src/framework/grcp/server-stream/proto"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", "world", "Name to greet")
)

func main() {
	flag.Parse()
	// 连接 RPC 服务器
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

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
