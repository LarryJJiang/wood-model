package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 库存记录表
type StockRecord struct {
	Model
	StockId  int     `gorm:"column:stock_id; size:11; not null; default:0;" json:"stock_id"`         // 库存ID
	CrewId   int     `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`           // 砍伐队ID
	Date     int     `gorm:"column:date; size:11; not null; default:0;" json:"date"`                 // 日期
	Grade    string  `gorm:"column:grade; size:20; " json:"grade"`                                   // 木材等级
	Code     string  `gorm:"column:code; size:20; " json:"code"`                                     // 编号
	Species  string  `gorm:"column:species; size:20; not null; default:'';" json:"species"`          // 木材种类
	Length   float64 `gorm:"type:decimal(5,2);column:length; not null; default:0.00;" json:"length"` // 木材长度
	Amount   float64 `gorm:"type:decimal(5,2);column:amount; not null; default:0.00;" json:"amount"` // 库存（垛）
	Status   int     `gorm:"column:status; size:11; not null; default:0;" json:"status"`             //状态
	DeleteAt int     `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`       //删除时间
	Account  string  `gorm:"-" json:"account"`
}

func (m *StockRecord) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "stock_record"
}

func (m *StockRecord) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
