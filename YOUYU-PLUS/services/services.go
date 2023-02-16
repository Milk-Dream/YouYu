package services

import (
	"fmt"
	"youyu/YOUYU-PLUS/models"
	"youyu/YOUYU-PLUS/utils"
)

var currentAccount *models.Account

//登录
func LoginServices(name, pwd string) (string, bool, *models.Account){
	//判断name是不是注册了
	ok, _ := utils.IsExists(name)
	if !ok {
		fmt.Println("用户名不存在")
		return "用户名不存在", false, nil
	}
	account := *models.GetAccount(name)
	models.SaveAccount(&account)
	
	/*account, ok := allAccounts[name]
	if !ok {
		fmt.Println("用户名不存在")
		return
	}*/
	//再次判断用户输入的密码和注册时保存的密码是否一致
	if utils.Hash(pwd) != account.Pwd {
		fmt.Println("密码不正确...")
		return "密码不正确...", false, nil
	}
	//account.isLogin = true
	currentAccount = &account
	
	return "登录成功", true, &account
}

//注册
func RegisterServices(name, pwd string) (string, bool){
	
	
	account := models.GetAccount(name)
	if account != nil {
		fmt.Println("账号已经存在!请更换一个用户名再注册")
		return "账号已经存在!请更换一个用户名再注册", false
	}
	
	newAccounts := models.NewAccount(name, utils.Hash(pwd))
	models.SaveAccount(&newAccounts)
	
	//allAccounts[user] = newAccounts
	//fmt.Println(user, "注册成功")
	return "注册成功", true
}

//展示余额
func ShowBalanceServices() (string, float64, bool){
	if currentAccount == nil {
		return "未登录", 0, false
	}

	return "余额查询成功", currentAccount.ShowBalance(), true
}

//查询明细
func ShowBalanceDetailsServices() (string, []models.Detail, bool){
	if currentAccount == nil {
		return "未登录", nil, false
	}
	details := currentAccount.ShowBalanceDetails()
	if len(details) == 0 {
		fmt.Println()
		return "你当前没有收支记录", nil, false
	}

	return "明细查询成功", details, true
	
}

//增加余额
func UpBalanceServices(amounts float64, message string) (string, bool){
	if currentAccount == nil {
		return "未登录", false
	}
	currentAccount.UpBalance(amounts, message)
	return "增加收入成功", true
}

//支出余额

func DownBalanceServices(amounts float64, message string) (string, bool) {
	if currentAccount == nil {
		return "未登录", false
	}
	if currentAccount.Balance > amounts {
		//fmt.Println("余额不足")
		return "账号余额不足", false
	}
	currentAccount.DownBalance(amounts, message)
	return "增加支出成功", true
}

