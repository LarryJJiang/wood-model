package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 砍伐队表
type Crew struct {
	Model
	Code            string                     `gorm:"column:code; size:50; " json:"code"`                                           // 砍伐队名称
	UserId          int                        `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`                 // 用户ID
	CustomerId      int                        `gorm:"column:customer_id; size:11; not null; default:0;" json:"customer_id"`         // 林场ID
	CustomerName    string                     `gorm:"column:customer_name; size:50; not null; default:'';" json:"customer_name"`    // 林场名称
	Email           string                     `gorm:"column:email; size:255; not null; default:'';" json:"email"`                   // 邮箱地址
	Notes           string                     `gorm:"column:notes; size:255; not null; default:'';" json:"notes"`                   // 笔记
	Setting         string                     `gorm:"column:setting; size:150; not null; default:'';" json:"setting"`               // 设置
	Compartment     string                     `gorm:"column:compartment; size:50; not null; default:'';" json:"compartment"`        // 坐标或位置
	Coordinate      string                     `gorm:"column:coordinate; size:50; not null; default:'';" json:"coordinate"`          // 坐标
	Channel         string                     `gorm:"column:channel; size:50; not null; default:'';" json:"channel"`                // 频道
	Forest          string                     `gorm:"column:forest; size:50; not null; default:'';" json:"forest"`                  // 林场
	Area            string                     `gorm:"column:area; size:50; not null; default:'';" json:"area"`                      // 区域
	Map             string                     `gorm:"column:map; size:50; not null; default:'';" json:"map"`                        // 地图链接
	OperationType   int                        `gorm:"column:operation_type; size:11; not null; default:0;" json:"operation_type"`   // 操作类型
	Rate            float64                    `gorm:"type:decimal(5,2);column:rate; not null; default:0.00;" json:"rate"`           // 费率
	RoundTripTime   int                        `gorm:"column:round_trip_time; size:11; not null; default:0;" json:"round_trip_time"` // 往返时间
	Status          int8                       `gorm:"column:status; type:tinyint(2); not null; default:0;" json:"status"`           // 状态
	DeleteAt        int                        `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`             //删除时间
	Account         string                     `gorm:"-" json:"account"`
	StockId         string                     `gorm:"-" json:"stock_id"`
	TotalAmount     float64                    `gorm:"-" json:"total_amount"`
	StockList       []*Stock                   `gorm:"ForeignKey:CrewId;" json:"stock_list"`
	RecordList      []*StockRecord             `gorm:"ForeignKey:CrewId;" json:"record_list"`
	CarrierList     []*CrewCarrierRelation     `gorm:"ForeignKey:CrewId;" json:"carrier_list"`
	DestinationList []*CrewDestinationRelation `gorm:"ForeignKey:CrewId;" json:"destination_list"`
	Customer        *Customer                  `gorm:"-" json:"customer"`
}

func (m *Crew) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "crew"
}

func (m *Crew) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
