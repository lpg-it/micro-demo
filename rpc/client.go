package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// rpc 与服务端建立网络连接
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")
	if err != nil {
		fmt.Println("rpc.DialHTTP error: ", err.Error())
		return
	}

	var cn int
	// 远程调用函数（被调用的方法，传入的参数，返回的参数）
	err = cli.Call("Conan.GetInfo", 10086, &cn)
	if err != nil {
		fmt.Println("cli.Call error: ", err.Error())
		return
	}
	fmt.Println("结果为：", cn)
}
