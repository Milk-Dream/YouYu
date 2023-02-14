package main

import (
	"crypto/md5"
	"fmt"

)

func main() {
	//第1种哈希计算
	str := "小白"
	ret := fmt.Sprintf("%x",md5.Sum([]byte(str+"Leo")))
	fmt.Println(ret)
	fmt.Println(Hash("小白"))
	
}

//第2种哈希计算
//Hash计算
func Hash(str string) string{
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte("Leo"))
	ret := fmt.Sprintf("%x", h.Sum(nil))
	return ret
}