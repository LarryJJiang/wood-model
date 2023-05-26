package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 目的地/目的地表
type Destination struct {
	Model
	Name               string `gorm:"column:name; size:50; " json:"name"`                                                      // 名称
	Code               string `gorm:"column:code; size:20; " json:"code"`                                                      // 简称
	Type               int8   `gorm:"column:type; type:tinyint(2); not null; default:0;" json:"type"`                          // 类型
	WeighBridgeSupport int8   `gorm:"column:weight_bridge_support; size:2; not null; default:0;" json:"weight_bridge_support"` // 状态
	Status             int8   `gorm:"column:status; size:2; not null; default:0;" json:"status"`                               // 状态
	DeleteAt           int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`                        //删除时间
}

func (m *Destination) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "destination"
}

func (m *Destination) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
