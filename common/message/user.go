package message

//定义一个用户的结构体
type User struct {
	UserID int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}
