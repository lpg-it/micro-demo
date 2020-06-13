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
	fmt.Println("test: ", test)
	// 将 struct test 转换成 protobuf
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Print("转码失败: ", err.Error())
	}
	fmt.Println("data: ", data)

	// 得到一个新的 Test 结构体：newTest
	newTest := &prototest.Test{}
	// 将 data 转换为 test 结构体
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("转码失败：", err.Error())
		return
	}
	fmt.Println(newTest.String())

	// 获取字段
	fmt.Println("newTest.Name: ", newTest.GetName())
	fmt.Println("newTest.Age: ", newTest.GetAge())
	fmt.Println("newTest.Hobbies: ", newTest.GetHobbies())
}
