package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	
	"youyu/utils"
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

func (ac *Account) Filepath() string{
	filepath := path.Join("accounts", fmt.Sprint(ac.Name, ".json"))
	return filepath
}

func SaveAccount(account *Account) {
	
	data, _ := json.Marshal(account)
	os.WriteFile(account.Filepath(), data, 0666)
}

func GetAccount(name string) *Account{
	ok, filepath := utils.IsExists(name)
	if !ok {
		fmt.Println("用户名不存在")
		return nil
	}
	accountJsons, _ := ioutil.ReadFile(filepath)
	var account Account
	json.Unmarshal(accountJsons, &account)
	return &account

}

func Login() {
	var (
		name string
		pwd string
		account Account

	)
	fmt.Print("请输入用户名:")
	fmt.Scanln(&name)
	fmt.Print("请输入密码:")
	fmt.Scanln(&pwd)
	//判断name是不是注册了
	ok, _ := utils.IsExists(name)
	if !ok {
		fmt.Println("用户名不存在")
		return
	}
	account = *GetAccount(name)
	SaveAccount(&account)
	currentAccount = account
	/*account, ok := allAccounts[name]
	if !ok {
		fmt.Println("用户名不存在")
		return
	}*/
	//再次判断用户输入的密码和注册时保存的密码是否一致
	if utils.Hash(pwd) != account.Pwd {
		fmt.Println("密码不正确...")
		return
	}
	account.isLogin = true

	fmt.Println("登录成功...")
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
	//ok, _ := utils.IsExists(user)

	// filepath := "./accounts/" + user + ".json"
	// _, err := os.Stat(filepath)
	
	// if !os.IsNotExist(err) {
	// 	fmt.Println("账号已经存在!请换一个账号再注册")
	// 	return
	// }

	

	account := GetAccount(user)
	if account != nil {
		fmt.Println("当前账号不存在")
	}
	
	newAccounts := NewAccount(user, utils.Hash(pwd))
	SaveAccount(&newAccounts)
	
	allAccounts[user] = newAccounts
	fmt.Println(user, "注册成功")
}

//查看余额
func (ac *Account)ShowBalance() {
	if !ac.isLogin {
		fmt.Println("请先登录再操作")
		return
	}
	fmt.Printf("你当前余额:%v元", ac.Balance)
}

//增加收入
func (ac *Account)UpBalance() {
	if !ac.isLogin {
		fmt.Println("请先登录再操作")
		return
	}
	var (
		amounts float64
		message string
	)
	fmt.Print("请输入收入金额")
	fmt.Scanln(&amounts)
	fmt.Print("请输入收入缘由")
	fmt.Scanln(&message)
	ac.Balance += amounts
	details := Detail {
		Kind: "收入",
		Amounts: amounts,
		Message: message,
	}
	
	ac.Details = append(ac.Details, details)
	filepath := path.Join("accounts", fmt.Sprint(ac.Name, ".json"))
	data, _ := json.Marshal(ac)
	os.WriteFile(filepath, data, 0666)
	fmt.Println("收入记录成功")
}

//支出
func (ac *Account)DownBalance() {
	if !ac.isLogin {
		fmt.Println("请先登录再操作")
		return
	}
	var (
		amounts float64
		message string
	)
	fmt.Print("请输入支出金额")
	fmt.Scanln(&amounts)
	fmt.Print("请输入支出缘由")
	fmt.Scanln(&message)
	if ac.Balance > amounts {
		fmt.Println("余额不足")
		return
	}

	ac.Balance -= amounts

	details := Detail {
		Kind: "支出",
		Amounts: amounts,
		Message: message,
	}
	
	ac.Details = append(ac.Details, details)
	filepath := path.Join("accounts", fmt.Sprint(ac.Name, ".json"))
	data, _ := json.Marshal(ac)
	os.WriteFile(filepath, data, 0666)
	fmt.Println("支出记录成功")
}

//查看明细
func (ac *Account) ShowBalanceDetails() {
	if !ac.isLogin {
		fmt.Println("请先登录再操作")
		return
	}

	if len(ac.Details) == 0 {
		fmt.Println("你当前没有收支记录")
		return
	}
	fmt.Println("你的收支如下:")

	for index, detail := range(ac.Details) {
		fmt.Printf("(%d)%v\t\t%v\t\t%v", index + 1, detail.Kind, detail.Amounts, detail.Message)
	}
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
			Login()
		case 2:
			Register()
		case 3:
			currentAccount.ShowBalance()
		case 4:
			currentAccount.ShowBalanceDetails()
		case 5:
			currentAccount.UpBalance()
		case 6:
			currentAccount.DownBalance()
		case 7:
			fmt.Println("欢迎下次再来")
			os.Exit(-1)
		default:
			fmt.Println("你输入的功能指令有误，请重新输入")
		}
	}
}