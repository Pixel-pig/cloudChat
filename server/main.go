package main

import (
	"cloudChat/common/message"
	"cloudChat/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
)


//处理和客户端的通讯
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	//循环的读取客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg()，返回Message，err
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return
			} else {
				fmt.Println("readPkg() err = ", err)
				return
			}
		}
		//fmt.Println("mes = ", mes)
		err = serverProcessLogin(conn, &mes)
		if err != nil {
			fmt.Println("serverProcessLogin(conn, &mes) err = ", err)
			return
		}

	}

}

//编写一个函数serverProcessLogin函数，专门处理登录请求
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(mes.Data), &loginMes) err = ", err)
		return
	}

	/**
	 * 2.验证mes.Data中的数据是否与数据库的数据一致，既用户的id，passWord是否相同
	 * 验证后用resMes返回客户端
	 */
	//2.1 先声明一个resMes用于返回客户端
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.2 再声明一个LoginResMes,并完成赋值
	var loginResMes message.LoginResMes

	//2.3 验证（loginMes）并返回
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//登录合法
		loginResMes.Code = 200
	} else {
		//登录不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册后再使用"
	}

	//2.4 将LoginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal(loginResMes) err = ", err)
		return
	}

	//2.5 将LoginResMes 序列化的data 赋值给resMes
	resMes.Data = string(data)

	//2.6 对resMes 序列化 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err = ", err)
		return
	}

	//2.7 发送resMes序列化的数据 data (将其封装到writerPkg函数中)
	err = utils.WriterPkg(conn, data)
	if err != nil {
		fmt.Println("writerPkg(conn, data) err = ", err)
		return
	}
	return
}

//编写一个ServerProcessMes函数
//功能:根据客户端发送的消息种类不同。决定调用哪个函数来处理
func ServerProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		err = serverProcessLogin(conn, mes)
	case message.RegisterMesType:
	//处理注册的逻辑
	default:
		fmt.Println("消息类型不存在")
	}
	return
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
