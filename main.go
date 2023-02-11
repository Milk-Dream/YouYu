package main

import (
	"fmt"
	"os"
)

//go mod init 名字

// go env -w GOPROXY=https://goproxy.cn,direct

type Account struct {
	Name string
	Pwd string
	Balance float64
}

//
func NewAccount(name, pwd string) Account {
	return Account {
		Name:name,
		Pwd:pwd,
		Balance: 0,
	}
}

var allAccounts = make(map[string]Account, 100)

func Register() {
	var (
		user string
		pwd string
		pwdAgain string
	)

	fmt.Print("请输入注册账号:")
	fmt.Scanln(&user)
	fmt.Print("请输入注册密码:")
	fmt.Scanln(&pwd)
	fmt.Print("请再次输入密码:")
	fmt.Scanln(&pwdAgain)

	if pwd != pwdAgain {
		fmt.Println("两次密码输入不一致~请重新输入")
		return
	}

	_, ok := allAccounts[user]
	if ok {
		fmt.Println("账号已经存在!请换一个账号再注册")
		return
	}

	newAccounts := NewAccount(user, pwd)
	allAccounts[user] = newAccounts
	



}

func main() {
	var choice int
	fmt.Println("欢迎使用小白YOUYU记账本")
	for {
		fmt.Println(`
			1.登录
			2.注册
			3.余额
			4.明细
			5.收入
			6.支出
			7.退出
		`)
		fmt.Print("请输入功能编号")
		fmt.Scanln(&choice)
		switch(choice) {
		case 1:
			fmt.Println("登录")
		case 2:
			fmt.Println("注册")
		case 3:
			fmt.Println("余额")
		case 4:
			fmt.Println("明细")
		case 5:
			fmt.Println("收入")
		case 6:
			fmt.Println("支出")
		case 7:
			fmt.Println("欢迎下次再来")
			os.Exit(-1)
		default:
			fmt.Println("你输入的功能指令有误，请重新输入")
		}
	}
}