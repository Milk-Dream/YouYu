package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"youyu/YOUYU-PLUS/utils"
)

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

//查看余额
func (ac *Account)ShowBalance() float64{
	return ac.Balance
	//fmt.Printf("你当前余额:%v元", ac.Balance)
}

//增加收入
func (ac *Account)UpBalance(amounts float64, message string) {

	details := Detail {
		Kind: "收入",
		Amounts: amounts,
		Message: message,
	}
	
	ac.Details = append(ac.Details, details)
	filepath := path.Join("accounts", fmt.Sprint(ac.Name, ".json"))
	data, _ := json.Marshal(ac)
	os.WriteFile(filepath, data, 0666)
	
}

//支出
func (ac *Account)DownBalance(amounts float64, message string) {
	
	

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
	
}

//查看明细
func (ac *Account) ShowBalanceDetails() []Detail{
	return ac.Details
}



