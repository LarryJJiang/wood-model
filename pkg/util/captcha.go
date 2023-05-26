package util

import (
	captcha "github.com/mojocn/base64Captcha"
)

var store = captcha.DefaultMemStore

func NewDriver() *captcha.DriverString {
	driver := new(captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = captcha.OptionShowSineLine | captcha.OptionShowSlimeLine | captcha.OptionShowHollowLine
	driver.Length = 6
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

// Generate 生成图形验证码
func Generate() (id, content string, err error) {
	var driver = NewDriver().ConvertFonts()
	//var driver = captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	c := captcha.NewCaptcha(driver, store)
	id, content, err = c.Generate()
	return
}

// Verify 验证
func Verify(id string, captcha string) bool {
	return store.Verify(id, captcha, true)
}
