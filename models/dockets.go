package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 单据表
type Dockets struct {
	Model
	TaskId        int     `gorm:"column:task_id; size:11; not null; default:0;" json:"task_id"`                       // 任务ID
	CrewId        int     `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`                       // 砍伐队ID
	Logger        string  `gorm:"column:logger; size:50; not null; default:'';" json:"logger"`                        // 砍伐队code
	VehicleId     int     `gorm:"column:vehicle_id; size:11; not null; default:0;" json:"vehicle_id"`                 // 货车ID
	TruckCode     string  `gorm:"column:truck_code; size:20; not null; default:'';" json:"truck_code"`                // 卡车代号
	CarrierId     int     `gorm:"column:carrier_id; size:11; not null; default:0;" json:"carrier_id"`                 // 运输公司ID
	Carrier       string  `gorm:"column:carrier; size:50; not null; default:'';" json:"carrier"`                      // 运输公司简称
	DocketNumber  string  `gorm:"column:docket_number; size:50; not null; default:'';" json:"docket_number"`          // 单据编号
	OperationType string  `gorm:"column:operation_type; size:20; not null; default:'';" json:"operation_type"`        // 操作：RL | CF | THIN
	Date          int64   `gorm:"column:date; size:20; not null; default:0;" json:"date"`                             // 日期
	Images        string  `gorm:"column:images; size:500; not null; default:'';" json:"images"`                       // 单据图片
	Detail        string  `gorm:"type:json;column:detail; size:500;" json:"detail"`                                   // 单据详情
	Customer      string  `gorm:"column:customer; size:50; not null; default:'';" json:"customer"`                    // 客户
	Forest        string  `gorm:"column:forest; size:50; not null; default:'';" json:"forest"`                        // 林场名称
	StockRecordId int     `gorm:"column:stock_record_id; size:11; not null; default:0;" json:"stock_record_id"`       // 库存记录ID
	Grade         string  `gorm:"column:grade; size:20; not null; default:'';" json:"grade"`                          // 等级
	Species       string  `gorm:"column:species; size:20; not null; default:'';" json:"species"`                      // 等级
	Length        float64 `gorm:"type:decimal(8,2);column:length; not null; default:0.00;" json:"length"`             // 长度
	Code          string  `gorm:"column:code; size:20; not null; default:'';" json:"code"`                            // 编码
	Compartment   string  `gorm:"column:compartment; size:50; not null; default:'';" json:"compartment"`              // 坐标和频道
	Setting       string  `gorm:"column:setting; size:50; not null; default:'';" json:"setting"`                      // 设置
	DestinationId int     `gorm:"column:destination_id; size:11; not null; default:0;" json:"destination_id"`         // 目的地ID
	Destination   string  `gorm:"column:destination; size:50; not null; default:'';" json:"destination"`              // 目的地
	TareWeight    float64 `gorm:"type:decimal(8,2);column:tare_weight; not null; default:0.00;" json:"tare_weight"`   // 自重(kg)
	GrossWeight   float64 `gorm:"type:decimal(8,2);column:gross_weight; not null; default:0.00;" json:"gross_weight"` // 总重(kg)
	NetWeight     float64 `gorm:"type:decimal(8,2);column:net_weight; not null; default:0.00;" json:"net_weight"`     // 净重(kg)
	Rate          float64 `gorm:"type:decimal(8,2);column:rate; not null; default:0.00;" json:"rate"`                 // 费率（$）
	//Status        int8     `gorm:"column:status; size:2; not null; default:0;" json:"status"`                          // 状态
	DeleteAt int      `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"` //删除时间
	Crew     *Crew    `json:"crew_info"`
	Vehicle  *Vehicle `json:"vehicle_info"`
}

func (m *Dockets) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "dockets"
}

func (m *Dockets) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
