package main

import (
	"cloudChat/client/process"
	"fmt"
	"os"
)

//定义两个变量，表示用户id和用户的密码
var userId int
var pwd string
var userName string

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

		fmt.Scanf("%d\n", &key) //获取用户的选择

		switch {
		case key == 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &pwd)
			//完成登陆
			//创建一个 user 实例获取用户输入的密码和用户名
			user := &process.UserProcess{}
			user.Login(userId, pwd)

			loop = false
		case key == 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户的id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &pwd)
			fmt.Println("请输入用户的名字(昵称)")
			fmt.Scanf("%s\n", &userName)
			//完成注册
			//创建一个 user 实例创建新用户
			user := &process.UserProcess{}
			user.Register(userId, pwd, userName)
		case key == 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误，请重新输入")
		}

	}
}
