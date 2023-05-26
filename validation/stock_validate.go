package valid

// 货车验证
type StockValidate struct {
	CrewId int     `alias:"Crew Id" valid:"Required; " form:"crew_id" json:"crew_id"` // 砍伐队ID
	Grade  string  `alias:"Grade" valid:"Required; " form:"grade" json:"grade"`       // 木材等级
	Code   string  `alias:"Code" valid:"Required; " form:"code" json:"code"`          // 编号
	Length float64 `alias:"Length" valid:"Required; " form:"length" json:"length"`    // 木材长度
	Amount float64 `alias:"Amount" form:"amount" json:"amount"`                       // 数量
}

// Valid 创建库存验证
func (a *StockValidate) Valid() (err error) {
	return Validate(a)
}
