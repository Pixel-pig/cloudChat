package message

const (
	LoginMesType =  "LoginMes"
	LoginResMesType = "LoginResMes "

)

//client与server数据交流对象(消息结构体)
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

/**
 *定义多个消息
 */
//1.a 登陆请求信息
type LoginMes struct {
	UserId int `json:"userId"` //用户id
	UserPwd string `json:"userPwd"` //用户密码
	UserName string  `json:"userName"` //用户名
}
//1.b 登陆响应消息
type LoginResMes struct {
	Code  int `json:"code"` //返回转态码， 500 表示该账户为注册 ；200 表示登陆成功
	Error string `json:"error"` //返回错误信息
}