package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 事故报告表
type IncidentReport struct {
	Model
	UserId       int    `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`               // 用户ID
	VehicleId    int    `gorm:"column:vehicle_id; size:11; not null; default:0;" json:"vehicle_id"`         // 货车ID
	IncidentType string `gorm:"column:incident_type; size:255; not null; default:'';" json:"incident_type"` // 事件类型，多选，用英文逗号隔开
	Location     string `gorm:"column:location; size:120; not null; default:'';" json:"location"`           // 位置
	Forest       string `gorm:"column:forest; size:50; not null; default:'';" json:"forest"`                // 林场
	Remark       string `gorm:"column:remark; size:255; not null; default:'';" json:"remark"`               // 备注
	Images       string `gorm:"column:images; size:255; not null; default:'';" json:"images"`               // 图片地址，英文逗号隔开，多张，不带域名
	DeleteAt     int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`           // 删除时间
}

func (m *IncidentReport) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "incident_report"
}

func (m *IncidentReport) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
