package check

import (
	"github.com/astaxie/beego/validation"
	"woods/pkg/bizerror"
)

func CheckParams(params interface{}) {
	if params != nil {
		valid := validation.Validation{}
		result, err := valid.Valid(params)
		bizcode.Check(err)
		if !result {
			//for _, err := range valid.Errors {
			//	bizcode.BizCode400.PanicError()
			//	bizcode.BizCode400.PanicErrorMsg(err.Field+" "+err.Message)
			//}
		}
	}
}
