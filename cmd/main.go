package main

import (
	"fmt"
	sever "go_bazooka/server"
)

func main() {
	bserver := sever.NewServer(8314)
	bserver.BindAcceptErrorHandler(acceptError)
	_ = bserver.Open()
	select {}
}

func acceptError(err error) {
	fmt.Println("监听出现异常：" + err.Error())
}

func openSuccess() {
	fmt.Println("打开端口成功")
}
