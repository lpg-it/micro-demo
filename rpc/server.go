package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"net/rpc"
)

// 用来实现网页的 web 函数
func conanTest(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello Conan")
}

/*
- 方法是导出的
- 方法有两个参数，都是导出类型或内建类型
- 方法的第二个参数是指针
- 方法只有一个error接口类型的返回值
func (t *T) MethodName(argType T1, replyType *T2) error
*/

type Conan int

// 函数关键字（对象）函数名（客户端发送过来的内容，返回给客户端的内容）错误返回值
func (this *Conan) GetInfo(argType int, replyType *int) error {
	fmt.Println("打印对端发送过来的内容为：", argType)
	// 修改内容值
	*replyType = argType + 10010

	return nil
}

func main() {
	// 注册一个页面请求
	http.HandleFunc("/conan", conanTest)

	// 将类实例化为对象
	cn := new(Conan)
	// 服务端注册一个对象
	// Register 在默认服务中注册并公布 接收服务 cn对象 的方法
	rpc.Register(cn)

	rpc.HandleHTTP()
	// 建立网络监听
	ln, err := net.Listen("tcp", ":10086")
	if err != nil {
		fmt.Println("net.Listen error: ", err.Error())
		return
	}
	fmt.Println("正在监听 10086...")

	http.Serve(ln, nil)
}
