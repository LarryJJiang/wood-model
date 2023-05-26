package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 砍伐队目的地关联表
type CrewDestinationRelation struct {
	Model
	CrewId          int    `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`                    // 用户ID
	DestinationId   int    `gorm:"column:destination_id; size:11; not null; default:0;" json:"destination_id"`      // 林场ID
	DestinationName string `gorm:"column:destination_name; size:50; not null; default:'';" json:"destination_name"` // 林场ID
	Status          int8   `gorm:"column:status; type:tinyint(2); not null; default:0;" json:"status"`              // 状态
	DeleteAt        int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`                //删除时间
}

func (m *CrewDestinationRelation) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "crew_destination_relation"
}

func (m *CrewDestinationRelation) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
