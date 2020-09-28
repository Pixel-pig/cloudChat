package main

import (
	"cloudChat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

//将读取数据包
func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 1024*4)
	//conn.Read 在conn 没有被关闭时 才会被堵塞，关闭则不会堵塞
	_, err  = conn.Read(buf[:4])//没有读到东西会堵塞
	if err != nil {
		return
	}
	//fmt.Println("读到的buf=", buf[:4])

	//根据buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])

	//根据pkgLen 读取消息内容(将pkgLen长度的内容读到buf里)
	n, err := conn.Read(buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		err = errors.New("read pkg body err")
		return
	}

	//把PkgLen 反序列化成 -> message.Message类型（mes）
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err = ", err)
		return
	}

	return
}

//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的读取客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg()，返回Message，err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return
			} else {
				fmt.Println("readPkg() err = ", err)
				return
			}
		}
		fmt.Println("mes = ", mes)
	}

}


func main() {
	//提示信息
	fmt.Println("服务器在8889端口监听....")
	listen, err := net.Listen("tcp","0.0.0.0:8889")
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