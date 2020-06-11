package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pt "micro_demo/myproto"
)

func main() {
	// 客户端连接服务器
	conn, err := grpc.Dial("127.0.0.1:18881", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial error: ", err.Error())
		return
	}
	defer conn.Close()

	// 获取 gRPC 句柄
	c := pt.NewHelloServerClient(conn)

	// 远程调用 SayHello 接口
	r1, err := c.SayHello(context.Background(), &pt.HelloRequest{Name: "Conan"})
	if err != nil {
		fmt.Println("c.SayHello error: ", err.Error())
		return
	}
	fmt.Println("SayHello 返回的消息：", r1.Msg)

	// 远程调用 PrintMotto 接口
	r2, err := c.PrintMotto(context.Background(), &pt.MottoRequest{Motto: "宁可有准备而没机会 也不要有机会而无准备"})
	if err != nil {
		fmt.Println("c.PrintMotto error: ", err.Error())
		return
	}
	fmt.Println("PrintMotto 返回的消息：", r2.Msg)
}
