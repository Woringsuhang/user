package common

import (
	"crypto/md5"
	"fmt"
)

func Md5(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	return md5str1
}
