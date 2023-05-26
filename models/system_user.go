package models

import (
	"fmt"
	"strconv"
	"time"
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

//func init() {
//	// 需要在init中注册定义的model
//	orm.RegisterModel(new(SystemUser))
//}

type SystemUser struct {
	Model
	UserName string    `gorm:"column:user_name; size:50; not null; default:'';" json:"user_name"`    //用户名
	Salt     string    `gorm:"column:salt; size:20; not null; default:'';" json:"salt"`              //用户名
	Password string    `gorm:"column:password; size:80; not null; default:'';" json:"password"`      //用户密码
	Account  string    `gorm:"column:account; size:150; not null; default:'';" json:"account"`       //登录账号
	Status   int       `gorm:"column:status; size:1; not null; default:0;" json:"status"`            //用户状态
	Identity int       `gorm:"column:identity; size:11; not null; default:0;" json:"identity"`       //用户身份 1：系统用户 2：卡车司机 3：砍伐队
	Portrait string    `gorm:"column:portrait; size:255; not null; default:'';" json:"portrait"`     //昵称
	LoginAt  time.Time `gorm:"column:login_at; type(datetime);not null; default:0;" json:"login_at"` //登陆时间
	Remark   string    `gorm:"column:remark; size:255; not null; default:'';" json:"remark"`         //备注
	DeleteAt int       `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`     //删除时间

	Captcha string `gorm:"-" json:"captcha"` //验证码
	Token   string `gorm:"-" json:"token"`   //token
}

func (systemUser *SystemUser) String() string {
	return fmt.Sprintf("{Id:%s,UserName:%s,Salt:%s,Password:%s,Identity:%s,Account:%s,Status:%s,Portrait:%s,DeleteAt:%s,LoginAt:%s}",
		strconv.Itoa(systemUser.ID),
		systemUser.UserName,
		systemUser.Salt,
		systemUser.Password,
		systemUser.Identity,
		systemUser.Account,
		systemUser.Status,
		systemUser.Portrait,
		systemUser.DeleteAt,
		systemUser.LoginAt,
	)
}

//表名
func (systemUser *SystemUser) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "system_user"
}

func (systemUser *SystemUser) GetFieldSlice() []string {
	return convert.GetFieldSlice(systemUser)
}
