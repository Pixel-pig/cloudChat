package main

import (
	"fmt"
	"net"
)


//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//调用总控,既创建一个 Processor 实例
	processor := &Processor{
		Conn:conn,
	}
	err := processor.Process3()
	if err != nil {
		fmt.Println("客户端与服务器通讯的协程 err =", err)
		return
	}

}


func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	//一旦监听成功，就等待客户端来链接服务器
	for {
		fmt.Println("等待客户端来链接服务器.......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦链接成功，则启动一个协程和客户端保持通讯...
		go process(conn)
	}
}
