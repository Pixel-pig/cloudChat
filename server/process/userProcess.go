package process

import (
	"cloudChat/common/message"
	"cloudChat/server/utils"
	"encoding/json"
	"fmt"
	"net"
)
//将这些方法关联到 UserProcess（用户处理）结构体中
type UserProcess struct {
	Conn net.Conn //从总控获取
}

//编写一个函数serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
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
	//创建Transfer实例，使用Transfer的方法
	transfer := &utils.Transfer{
		Conn: this.Conn,
	}
	err = transfer.WritePkg(data)
	if err != nil {
		fmt.Println("writerPkg(conn, data) err = ", err)
		return
	}
	return
}