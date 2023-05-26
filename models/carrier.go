package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 运输公司表
type Carrier struct {
	Model
	Name     string `gorm:"column:name; size:50; " json:"name"`                               // 名称
	Code     string `gorm:"column:code; size:20; " json:"code"`                               // 运输公司简称
	UserId   int    `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`     // 用户ID
	Status   int8   `gorm:"column:status; size:2; not null; default:0;" json:"status"`        // 状态
	DeleteAt int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"` //删除时间
	Account  string `gorm:"-" json:"account"`
	RoleId   int    `gorm:"-" json:"role_id"`
}

func (m *Carrier) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "carrier"
}

func (m *Carrier) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
