package bizcode

var MsgFlags = map[int]string{
	SUCCESS:               "Success",
	ERROR:                 "System Error",
	InvalidParams:         "Param Error",
	NoneToken:             "Token Not Exists",
	NotPermission:         "没有权限",
	NotExist:              "请求资源不存在",
	ServerNotAuthorize:    "服务端未授权",
	RepeatedAction:        "你的操作太快了，请稍候再试",
	ErrorRequestHeader:    "请求头部信息错误",
	Unknown:               "未知错误",
	DialTimeout:           "访问超时",
	ErrorAuthorizeDisable: "Authorization Not Available",
	ErrorIdentity:         "Error Identity",

	ErrorUsername:              "用户名错误",
	ErrorPassword:              "用户密码错误",
	ErrorAuthCheckTokenFail:    "用户token验证失败",
	ErrorAuthCheckTokenTimeout: "用户token验证超时",
	ErrorUsernameExists:        "用户名已存在",
	ErrorPhoneExists:           "手机号已被绑定",
	ErrorUser:                  "用户不存在",
	ErrorUserDescTooLong:       "用户简介字数过长(最多800字符)",
	ErrorUserRepeatLogin:       "用户已登录，请勿重复登录",
	ErrorUserOtherLogin:        "账号已在另一部设备登录，请确认是否本人操作",
	ErrorEmptyToken:            "请求头部token为空",

	ErrorRedis:      "redis相关错误",
	ErrorDb:         "Operate Database Error",
	ErrorImgBase64:  "图片base64不正确",
	ErrorImgExt:     "仅支持png/jpg/jpeg/gif图片",
	ErrorImgTooMax:  "图片大小超过限制",
	ErrorSystemLock: "系统锁不存在",

	ErrorFileFormat:    "错误的文件格式",
	ErrorFileSave:      "文件保存失败",
	ErrorFileNotExists: "文件不存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
