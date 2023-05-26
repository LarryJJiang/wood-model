package crypt

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha() (string, string) {
	driver := base64Captcha.DefaultDriverDigit
	//生成base64图片
	c := base64Captcha.NewCaptcha(driver, store)

	//获取
	id, b64s, err := c.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", ""
	}
	return id, b64s
}

func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	//同时在内存清理掉这个图片
	return store.Verify(id, val, true)
}
