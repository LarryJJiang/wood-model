package valid

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego/validation"
	"reflect"
	"strings"
)

func InitValidate() {
	SetDefaultMessage()
}

type ValidateError struct {
	Code string
	Msg  string
}

type Base struct {
	DeviceId string `json:"device_id"  alias:"设备ID"  valid:""`
}

func SetDefaultMessage() {
	if len(MessageTmpls) == 0 {
		return
	}
	//将默认的提示信息转为自定义
	for k, _ := range MessageTmpls {
		validation.MessageTmpls[k] = MessageTmpls[k]
	}

	//增加默认的自定义验证方法
	_ = validation.AddCustomFunc("Unique", Unique)
	_ = validation.AddCustomFunc("url", Url)
	_ = validation.AddCustomFunc("RequiredIf", RequireIf)
}

var Unique validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {

}

var Url validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
	if !govalidator.IsURL(fmt.Sprintf("%v", obj)) {
		v.AddError(key, "url链接非法")
		return
	}
}

var RequireIf validation.CustomFunc = func(v *validation.Validation, obj interface{}, key string) {
	fmt.Printf("验证：%v ab%vaa, %v\n", v, obj, key)
	return
}

var MessageTmpls = map[string]string{
	"Required":     "can not empty",
	"Min":          "min value is %d",
	"Max":          "max value is %d",
	"Range":        "range %d to %d",
	"MinSize":      "min length is %d",
	"MaxSize":      "max length is %d",
	"Length":       "length must be %d",
	"Alpha":        "must be valid char",
	"Numeric":      "must be valid number",
	"AlphaNumeric": "must be valid char or number",
	"Match":        "必须匹配格式 %s",
	"NoMatch":      "必须不匹配格式 %s",
	"AlphaDash":    "必须是有效的字母或数字或破折号(-_)字符",
	"Email":        "必须是有效的邮件地址",
	"IP":           "必须是有效的IP地址",
	"Base64":       "必须是有效的base64字符",
	"Mobile":       "必须是有效手机号码",
	"Tel":          "必须是有效电话号码",
	"Phone":        "必须是有效的电话号码或者手机号码",
	"ZipCode":      "必须是有效的邮政编码",
}

func Validate(data interface{}) (err error) {
	InitValidate()
	valid := validation.Validation{}
	b, _ := valid.Valid(data)
	if !b {
		//表示获取验证的结构体
		v := reflect.TypeOf(data)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct { // 非结构体返回错误提示
			return errors.New("校验数据格式错误")
		}
		for i := 0; i < v.NumField(); i++ {
			fi := v.Field(i)
			for _, err := range valid.Errors {
				//获取验证的字段名和提示信息的别名
				if fi.Name == err.Field {
					var alias = v.Field(i).Tag.Get("alias")
					var message = strings.Replace(err.Message, err.Field+" ", "", 1)
					//返回验证的错误信息
					return errors.New(alias + message)
				}
			}
		}
	}
	return nil

}
