package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 货车表
type Vehicle struct {
	Model
	Registration string `gorm:"column:registration; size:50; " json:"registration"` // 车牌号
	Code         string `gorm:"column:code; size:20; " json:"code"`                 // 货车代码
	//PinCode        string `gorm:"column:pin_code; size:20; " json:"pin_code"`                                    // 频道密码
	UserId         int    `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`                  // 用户ID
	CarrierId      int    `gorm:"column:carrier_id; size:11; not null; default:0;" json:"carrier_id"`            // 所属运输公司ID
	CarrierName    string `gorm:"column:carrier_name; size:50; not null; default:'';" json:"carrier_name"`       // 所属运输公司名称
	AppVersion     string `gorm:"column:app_version; size:20; not null; default:'';" json:"app_version"`         // 安装的App版本
	CarHopperCount int8   `gorm:"column:car_hopper_count; size:2; not null; default:1;" json:"car_hopper_count"` // 车斗数量
	Status         int8   `gorm:"column:status; size:2; not null; default:0;" json:"status"`                     // 状态
	DeleteAt       int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`              //删除时间
	Account        string `gorm:"-" json:"account"`
}

func (m *Vehicle) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "vehicle"
}

func (m *Vehicle) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
