package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 质控考核题目答案表
type User struct {
	Model
	NickName        string  `gorm:"column:nick_name; size:50; " json:"nick_name"`                                                             // 昵称
	GradeId         int     `gorm:"column:grade_id; size:11; not null; default:0;" json:"grade_id"`                                           // 会员等级ID
	AvatarId        int     `gorm:"column:avatar_id; size:11; not null; default:0;" json:"avatar_id"`                                         // 头像文件ID
	Status          int8    `gorm:"column:status; size:2; not null; default:0;" json:"status"`                                                // 状态
	Gender          int8    `gorm:"column:gender; size:2; not null; default:0;" json:"gender"`                                                // 性别
	Mobile          string  `gorm:"column:mobile; size:11;" json:"mobile"`                                                                    // 用户手机号
	Country         string  `gorm:"column:country; size:50; not null; default:'';" json:"country"`                                            // 国家
	Province        string  `gorm:"column:province; size:50; not null; default:'';" json:"province"`                                          // 省份
	City            string  `gorm:"column:city; size:50; not null; default:'';" json:"city"`                                                  // 城市
	AddressId       int     `gorm:"column:address_id; size:11; not null; default:0;" json:"address_id"`                                       // 默认收货地址
	Balance         float32 `gorm:"column:balance; size:11; not null; default:0;" json:"balance" sql:"type:decimal(10,2);"`                   // 用户可用余额
	PayMoney        float32 `gorm:"column:pay_money; size:11; not null; default:0;" json:"pay_money" sql:"type:decimal(10,2);"`               // 用户总支付的金额
	ExpendMoney     float32 `gorm:"column:expend_money; size:11; not null; default:0;" json:"expend_money" sql:"type:decimal(10,2);"`         // 实际消费的金额(不含退款)
	TeamMoney       float32 `gorm:"column:team_money; size:11; not null; default:0;" json:"team_money" sql:"type:decimal(10,2);"`             // 团队业绩
	RecommendAmount float32 `gorm:"column:recommend_amount; size:11; not null; default:0;" json:"recommend_amount" sql:"type:decimal(10,2);"` // 推荐报单业绩
	Points          int     `gorm:"column:points; size:2; not null; default:0;" json:"points"`                                                // 用户可用积分
	LastLoginTime   int     `gorm:"column:last_login_time; size:1; not null; default:0;" json:"last_login_time"`                              // 最后登录时间
	Platform        string  `gorm:"column:platform; size:255;" json:"platform"`                                                               // 注册来源的平台 (APP、H5、小程序等)
	InviteCode      int     `gorm:"column:invite_code; size:11; not null; default:0;" json:"invite_code"`                                     // 邀请码
	IsDelete        int8    `gorm:"column:is_delete; size:2;" json:"is_delete"`                                                               // 是否删除
	StoreId         int     `gorm:"column:store_id; size:11;" json:"store_id"`                                                                // 商城ID
}

func (m *User) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "user"
}

func (m *User) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
