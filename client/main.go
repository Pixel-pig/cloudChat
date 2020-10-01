package main

import (
	"cloudChat/client/login"
	"fmt"
	"os"
)

//定义两个变量，表示用户id和用户的密码
var userId int
var pwd string

//接收用户的选择
var key int

//判断是否还继续显示菜单
var loop = true

func main() {

	for {
		fmt.Println("--------------欢迎登陆多人聊天系统--------------")
		fmt.Println("\t\t\t 1 登陆聊天室")
		fmt.Println("\t\t\t 2 注册用户")
		fmt.Println("\t\t\t 3 退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Println("----------------------------------------------")

		fmt.Scanf("%d\n",&key)//获取用户的选择

		switch {
		case key == 1:
			fmt.Println("登陆聊天室")
			loop = false
		case key == 2:
			fmt.Println("注册用户")
		case key == 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	 	//根据用户输入，显示用户信息
		if key == 1 {
			//说明用户要登陆
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%v\n",&pwd)
			//先把登陆的函数，写到另一个文件，login.go
			login.Login(userId, pwd)
			//if err != nil {
			//	fmt.Println("登陆失败")
			//} else {
			//	fmt.Println("登陆成功")
			//}
		} else if key == 2{
			fmt.Println("进行用户注册")
		}

	}


}