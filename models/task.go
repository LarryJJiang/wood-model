package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 任务表
type Task struct {
	Model
	CrewId                 int        `gorm:"column:crew_id; size:11; not null; default:0; " json:"crew_id"`                                    // 车牌号
	UserId                 int        `gorm:"column:user_id; size:11; not null; default:0; " json:"user_id"`                                    // 车牌号
	VehicleId              int        `gorm:"column:vehicle_id; size:11; not null; default:0; " json:"vehicle_id"`                              // 车牌号
	VehicleCode            string     `gorm:"column:vehicle_code; size:20; not null; default:''; " json:"vehicle_code"`                         // 货车代码
	Date                   int        `gorm:"column:date; size:11; not null; default:0; " json:"date"`                                          // 日期
	StartTime              int        `gorm:"column:start_time; size:11; not null; default:0;" json:"start_time"`                               // 用户ID
	EndTime                int        `gorm:"column:end_time; size:11; not null; default:0;" json:"end_time"`                                   // 用户ID
	Rate                   float32    `gorm:"type:decimal(5,2);column:rate; not null; default:0.00;" json:"rate"`                               // 费率
	Notes                  string     `gorm:"column:notes; size:255; not null; default:'';" json:"notes"`                                       // 所属运输公司名称
	Destination            string     `gorm:"column:destination; size:255; not null; default:'';" json:"destination"`                           // 所属运输公司名称
	DeclineReason          string     `gorm:"column:decline_reason; size:255; not null; default:'';" json:"decline_reason"`                     // 安装的App版本
	WeightBridgeFailRemark string     `gorm:"column:weightbridge_fail_remark; size:255; not null; default:'';" json:"weightbridge_fail_remark"` // 安装的App版本
	Status                 int8       `gorm:"column:status; size:2; not null; default:0;" json:"status"`                                        // 状态
	DeleteAt               int        `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`                                 //删除时间
	Account                string     `gorm:"-" json:"account"`
	Crew                   *Crew      `gorm:"-" json:"crew_info"`
	DocketList             []*Dockets `gorm:"ForeignKey:TaskId;" json:"docket_list"`
}

func (m *Task) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "task"
}

func (m *Task) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
