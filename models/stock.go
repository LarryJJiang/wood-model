package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 库存表
type Stock struct {
	Model
	UserId      int            `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`                       // 砍伐队用户ID
	CrewId      int            `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`                       // 砍伐队ID
	CustomerId  int            `gorm:"column:customer_id; size:11; not null; default:0;" json:"customer_id"`               // 林场ID
	Date        int            `gorm:"column:date; size:11; not null; default:0;" json:"date"`                             // 日期，用于查询每天的库存
	TotalAmount float64        `gorm:"type:decimal(5,2);column:total_amount; not null; default:0.00;" json:"total_amount"` // 总库存（垛）
	DeleteAt    int            `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`                   //删除时间
	Account     string         `gorm:"-" json:"account"`
	StockList   []*StockRecord `gorm:"ForeignKey:StockId;" json:"stock_list"`
}

func (m *Stock) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "stock"
}

func (m *Stock) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
