package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 砍伐队运输公司表
type CrewCarrierRelation struct {
	Model
	CrewId      int    `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`            // 用户ID
	CarrierId   int    `gorm:"column:carrier_id; size:11; not null; default:0;" json:"carrier_id"`      // 运输公司
	CarrierCode string `gorm:"column:carrier_code; size:20; not null; default:'';" json:"carrier_code"` // 运输公司编号
	Status      int8   `gorm:"column:status; type:tinyint(2); not null; default:0;" json:"status"`      // 状态
	DeleteAt    int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`        //删除时间
}

func (m *CrewCarrierRelation) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "crew_carrier_relation"
}

func (m *CrewCarrierRelation) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
