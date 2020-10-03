package main

import (
	"cloudChat/common/message"
	process2 "cloudChat/server/process"
	"cloudChat/server/utils"
	"fmt"
	"io"
	"net"
)

//创建 Processor 的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes函数
//功能:根据客户端发送的消息种类不同。决定调用哪个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录的逻辑
		//创建一个UserProcess实例
		user := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册的逻辑
		//创建一个UserProcess实例
		user := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = user.ServerProcessRegister(mes)
	default:
		fmt.Println("消息类型不存在")
	}
	return
}

func (this *Processor) Process3() (err error) {
	//循环的读取客户端发送的信息
	for {
		//这里我们将读取数据包，直接封装成一个函数readPkg()，返回Message，err
		//创建 Transfer 实例完成读包的任务
		transfer := &utils.Transfer{
			Conn: this.Conn,
		}
		mes, err := transfer.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出")
				return err
			} else {
				fmt.Println("readPkg() err = ", err)
				return err
			}
		}
		//fmt.Println("mes = ", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			fmt.Println("serverProcessLogin(conn, &mes) err = ", err)
			return err
		}

	}
}
