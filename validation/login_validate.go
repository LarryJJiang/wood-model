package valid

import (
	"github.com/gin-gonic/gin"
)

type LoginValidate struct {
	Account  string `alias:"account" valid:"Required; " form:"account" json:"account"`    // 账号
	Password string `alias:"Password" valid:"Required; " form:"password" json:"password"` // 密码
}

// Valid 创建登录校验
func (a *LoginValidate) Valid(ctx *gin.Context) (err error) {
	return Validate(a)
}
