package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"micro_demo/proto-learn/prototest"
)

func main() {
	test := &prototest.Test{
		Name:    "李培冠",
		Age:     18,
		Hobbies: []string{"dance", "read"},
	}
	fmt.Println(test)
	// 将 struct test 转换成 protobuf
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Print("转码失败: ", err.Error())
	}
	fmt.Print(data)
}
