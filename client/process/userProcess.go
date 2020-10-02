package process

import (
	"cloudChat/client/utils"
	"cloudChat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

//给 UserProcess 结构体绑定一些字段和方法
type UserProcess struct {

}

//Login函数，完成登陆
func (this *UserProcess) Login(userId int, pwd string) (err error) {
	//1.a:链接服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net,Dial err=", err)
		return
	}
	//1.b:延时关闭链接
	defer conn.Close()

	//2.准备通过conn发送消息个服务器
	var mes message.Message //定义message结构体
	mes.Type = message.LoginMesType

	//3.a:创建LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = pwd
	//3.b:将LoginMes结构体序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("LoginMes json.Marsahl err=", err)
		return
	}
	//3.c:把data信息 赋给 mes.Data字段
	mes.Data = string(data)

	//4.将mes结构体序列化(这个data既是我们要发送的消息)
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("mes json.Marsahl err=", err)
		return
	}

	/**
	 *5.计算data数据的长度，并将数据发送个服务端（避免丢包）
	 */
	//5.a:先获取data的长度(int 类型) -> 转为切片
	//错误类型：var len := len(data)
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pkgLen)
	//5.b:发送长度 n 为字节长度
	_, err = conn.Write(buf[:4])
	if err != nil {
		fmt.Println("conn.Write(buf) err=", err)
		return
	}

	//6.发送Mes(包含LoginMes)消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err=", err)
		return
	}

	//处理服务端返回的消息
	//创建一个 transfer 链接
	transfer := &utils.Transfer{
		Conn:conn,
	}
	mes, err = transfer.ReadPkg()
	if err != nil {
		fmt.Println("utils.ReadPkg(conn) err = ", err)
		return
	}
	//将mes.Data 反序列化 成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端，接收并显示在客户端
		go processServerMes(conn)

		//1. 循环显示我们登陆列表
		for {
			ShowMenu()
		}
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
		return
	}

	return
}
