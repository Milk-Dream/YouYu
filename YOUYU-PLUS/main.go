package main

import (
	"fmt"
	"youyu/YOUYU-PLUS/services"

	"os"
)

//go mod init 名字

// go env -w GOPROXY=https://goproxy.cn,direct

//明细记录
type Detail struct {
	Kind string `json:"kind"`
	Amounts float64 `json:"amounts"`
	Message string `json:"message"`
}

type Account struct {
	Name string `json:"name"`
	Pwd string `json:"pwd"`
	Balance float64 `json:"balance"`
	Details []Detail `json:"details"`
	isLogin bool
}

var currentAccount Account

//
func NewAccount(name, pwd string) Account {
	return Account {
		Name:name,
		Pwd:pwd,
		Balance: 0,
		
	}
}

//存放注册的账号
var allAccounts = make(map[string]Account, 100)







func Login() {
	var (
		name string
		pwd string
		

	)
	fmt.Print("请输入用户名:")
	fmt.Scanln(&name)
	fmt.Print("请输入密码:")
	fmt.Scanln(&pwd)
	//判断name是不是注册了
	message, _, _ := services.LoginServices(name, pwd)
	fmt.Println(message)
}

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

	message, _ := services.RegisterServices(user, pwd)
	fmt.Println(message)
}

//查看余额
func ShowBalance() {
	message, _, _ := services.ShowBalanceServices()
	fmt.Println(message)
}

//增加收入
func UpBalance() {
	var (
		amounts float64
		message string
	)
	fmt.Print("请输入收入金额")
	fmt.Scanln(&amounts)
	fmt.Print("请输入收入缘由")
	fmt.Scanln(&message)
	

	msg, _ := services.UpBalanceServices(amounts, message)

	fmt.Println(msg)


}

//支出
func DownBalance() {
	var (
		amounts float64
		message string
	)
	fmt.Print("请输入支出金额")
	fmt.Scanln(&amounts)
	fmt.Print("请输入支出缘由")
	fmt.Scanln(&message)
	msg, _ := services.DownBalanceServices(amounts, message)

	fmt.Println(msg)
	
}

//查看明细
func ShowBalanceDetails() {
	msg, details, ok := services.ShowBalanceDetailsServices()
	fmt.Println(msg)
	if !ok {
		return
	}
	fmt.Println("你的明细如下:")
	for index, detail := range(details) {
		fmt.Printf("(%d)%v\t\t%v\t\t%v", index + 1, detail.Kind, detail.Amounts, detail.Message)
	}

	
}

func menu() {
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
			Login()
		case 2:
			Register()
		case 3:
			ShowBalance()
		case 4:
			ShowBalanceDetails()
		case 5:
			UpBalance()
		case 6:
			DownBalance()
		case 7:
			fmt.Println("欢迎下次再来")
			os.Exit(-1)
		default:
			fmt.Println("你输入的功能指令有误，请重新输入")
		}
	}
}

func main() {
	//菜单选项
	menu()
}