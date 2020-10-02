package process

import (
	"cloudChat/client/utils"
	"fmt"
	"net"
	"os"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("-----------恭喜xxx登陆成功-----------")
	fmt.Println("-----------1. 显示在线列表-----------")
	fmt.Println("-----------2. 发送消息-----------")
	fmt.Println("-----------3. 信息列表-----------")
	fmt.Println("-----------4. 退出系统-----------")
	fmt.Println("请选择(1-4):")
	var key int
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		fmt.Println("显示在线列表")
	case 2:
		fmt.Println("发送消息")
	case 3:
		fmt.Println("信息列表")
	case 4:
		fmt.Println("你选择退出了系统")
		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确")
	}
}

//和服务器端保持通讯
func processServerMes (conn net.Conn) {
	//创建一个 transfer 实例，不停的读取服务器发送的消息
	transfer := &utils.Transfer{
		Conn : conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送消息")
		mes, err := transfer.ReadPkg()
		if err != nil {
			fmt.Println("transfer.ReadPkg() err = ", err)
			return
		}
		//如果读取到了数据
		fmt.Println("mes = ", mes)
	}

}
