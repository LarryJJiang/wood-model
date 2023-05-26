package e

const (
	SUCCESS               = 200   // 成功
	ERROR                 = 500   // 系统错误
	InvalidParams         = 400   // 参数错误
	NoneToken             = 401   // token不存在
	InvalidToken          = 402   // 无效的token
	NotPermission         = 403   // 没有权限
	NotExist              = 404   // 请求资源不存在
	ServerNotAuthorize    = 10007 // 服务端未授权
	RepeatedAction        = 10008 // 你的操作太快了，请稍候再试
	ErrorRequestHeader    = 10009 // 请求头部信息错误
	Unknown               = 10011 // 未知错误
	DialTimeout           = 10012 // 访问超时
	ErrorAuthorizeDisable = 10016 // 授权不可用
	ErrorIdentity         = 10017 // 不正确的身份

	ErrorUsername              = 10101 // 用户名错误
	ErrorPassword              = 10102 // 用户密码错误
	ErrorAuthCheckTokenFail    = 10103 // 用户token验证失败
	ErrorAuthCheckTokenTimeout = 10104 // 用户token验证超时
	ErrorUsernameExists        = 10105 // 用户名已存在
	ErrorPhoneExists           = 10106 // 手机号已被绑定
	ErrorUser                  = 10107 // 用户不存在
	ErrorUserDescTooLong       = 10108 // 用户简介字数过长(最多800字符)
	ErrorUserRepeatLogin       = 10109 // 用户已登录，请勿重复登录
	ErrorUserOtherLogin        = 10110 // 账号已在另一部设备登录，请确认是否本人操作
	ErrorEmptyToken            = 10111 // 头部token为空

	ErrorRedis         = 10501 // redis相关错误
	ErrorDb            = 10502 // 数据库操作相关错误
	ErrorImgBase64     = 10503 // 图片base64不正确
	ErrorImgExt        = 10504 // 仅支持png/jpg/jpeg/gif图片
	ErrorImgTooMax     = 10505 // 图片大小超过限制
	ErrorFileFormat    = 10801 // 错误的文件格式
	ErrorFileSave      = 10802 // 文件保存失败
	ErrorFileNotExists = 10803 // 文件不存在
	ErrorSystemLock    = 10506 // 系统锁不存在BizError404900 = BizError{"404900", "系统锁不存在!"}
)
