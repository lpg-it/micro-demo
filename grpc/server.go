package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "micro_demo/myproto"
	"net"
)

const (
	post = "127.0.0.1:18881"
)

// 对象要和 proto 文件内定义的服务一样
type server struct{}

// 实现 RPC SayHello 接口
// 函数关键字 () 函数名(context, 客户端发送过来的参数) (发送给客户端的参数, 错误)
func (this *server) SayHello(ctx context.Context, in *pt.HelloRequest) (*pt.HelloResponse, error) {
	return &pt.HelloResponse{Msg: "Hello " + in.Name}, nil
}

// 实现 RPC PrintMotto 接口
func (this *server) PrintMotto(ctx context.Context, in *pt.MottoRequest) (*pt.MottoResponse, error) {
	return &pt.MottoResponse{
		Msg: in.Motto,
	}, nil
}

func main() {
	// 监听网络
	ln, err := net.Listen("tcp", post)
	if err != nil {
		fmt.Println("net.Listen error: ", err.Error())
		return
	}

	// 创建一个 gRPC 的句柄
	srv := grpc.NewServer()
	// 将 server 结构体注册到 gRPC 中
	pt.RegisterHelloServerServer(srv, &server{})

	// 监听 gRPC 服务
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("srv.Serve error: ", err.Error())
		return
	}
}
