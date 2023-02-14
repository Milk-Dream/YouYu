package utils

import (
	"fmt"
	"os"
	"path"
	"Crypto/md5"
)

func IsExists(name string) (bool, string) {
	//./accounts/jack.json
	filepath := path.Join("accounts", fmt.Sprint(name, ".json"))
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		// fmt.Println("当前用户不存在")
		return false, filepath
	}
	return true, filepath

}

func Hash(str string) string{
	h := md5.New()
	h.Write([]byte(str))
	h.Write([]byte("Leo"))
	ret := fmt.Sprintf("%x", h.Sum(nil))
	return ret
}