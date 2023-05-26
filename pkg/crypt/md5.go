package crypt

import (
	"crypto/md5"
	"fmt"
)

var key = "!@#$%^&*()_GOBLOG"

func GetMd5(str string) string {
	if str == "" {
		return str
	}

	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

func GetSystemPassword(password string, salt string) string {
	if password == "" {
		return password
	}
	data := GetMd5(password)
	data += salt
	data = GetMd5(data)
	return data
}
