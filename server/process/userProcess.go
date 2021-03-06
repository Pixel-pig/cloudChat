//（user的控制器）
package process

import (
	"cloudChat/common/message"
	"cloudChat/server/model"
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
	//(使用MyUserDao.Login方法验证redis中的数据与用户输入的数据是否一致)
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		//登录不合法
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500 //用户不存在
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300 //密码不正确
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505 //未知错误
			loginResMes.Error = "服务器内部错误......"
		}

	} else {
		//登录合法
		loginResMes.Code = 200
		fmt.Println(user, " 登录成功")
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

//编写一个函数serverProcessRegister函数，专门处理注册请求
func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并直接反序列化成registerMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}

	//2.1 先声明一个resMes用于返回客户端
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	//2.2 再声明一个 RegisterMes,并完成赋值
	var registerResMes message.RegisterResMes

	//3. 去 redis 数据中完成注册
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_TEXISTS {
			registerResMes.Code = 505 //用户已被占用
			registerResMes.Error = model.ERROR_USER_TEXISTS.Error()
		} else {
			registerResMes.Code = 506 //注册时发生未知错误
			registerResMes.Error = "注册时发生未知错误..."
		}
	} else {
		registerResMes.Code = 200 //注册成功
	}

	//4. 将 registerResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//5 将 registerResMes 序列化的data 赋值给resMes
	resMes.Data = string(data)

	//6. 对resMes 序列化 准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal(resMes) err = ", err)
		return
	}

	//7. 发送resMes序列化的数据 data (将其封装到writerPkg函数中)
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
