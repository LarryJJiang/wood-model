package e

const (
	TokenKey                = "Authorization"                      // 页面token键名
	AuthorLanguage          = "author-language"                    // 语言类型：cn-中文，en-英文；(默认：cn)
	AuthorNumber            = "author-number"                      // 客户端授权id
	UserIdKey               = "X-USER-ID"                          // 页面用户ID键名
	UserUuidKey             = "X-UUID"                             // 页面UUID键名
	UserNameKey             = "X-USERNAME"                         // 页面用户名
	UserEmailKey            = "X-EMAIL"                            // 页面邮箱
	UserIdentityKey         = "X-USER-IDENTITY"                    // 页面用户身份
	UserInfo                = "X-USER-INFO"                        // 用户信息
	UserRoleId              = "X-USER-Role"                        // 用户角色
	UserLoginTokenState     = "local:user:%s:token:state"          // 用户登录的token状态
	UserLoginErrorTimes     = "local:user:%s:%s:login:error:times" // 用户错误登录次数
	CustomerTokenState      = "local:customer:%s:token:state"      // 客户端用户登录的token状态
	Day                     = 24 * 60 * 60                         // 一天
	Hour                    = 60 * 60                              // 一小时
	Minute                  = 60                                   // 一分钟
	CustomerTokenExpireTime = 60 * 60 * 5                          // 5小时
	AdminTokenExpireTime    = 60 * 60 * 2                          // 2小时
	StockRecordUpdateToken  = "crew:%s:stock:%s:record:%s:token"   // 砍伐队更新库存记录token
	DocketNumberGenerator   = "docket_number_generator"            // docket number 生成最初值
	DocketNumberMin         = 4530012                              // docket number 生成最初值
	DriverTaskGeo           = "geo:task_%s"                        // 卡车司机任务过程中的位置信息
)
