package main

import (
	"encoding/json"
	"fmt"
	
)

type Account struct {
	Name string `json:"name"`
	Pwd string `json:"pwd"`
	Balance float64 `json:"balance"`
}

func main() {
	account := Account{"jack", "123", 13}
	//序列化方法
	data1, _ := json.Marshal(account)
	fmt.Println(string(data1))
	fmt.Printf("%T\n", data1)

	//反序列化方法
	var account2 Account
	json.Unmarshal(data1, &account2)
	fmt.Println(account2)
	fmt.Printf("%T\n", account2)
}