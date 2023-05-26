package valid

// 货车验证
type StockUpdateValidate struct {
	CrewId int     `alias:"Crew Id" valid:"Required; " form:"crew_id" json:"crew_id"`           // 砍伐队ID
	Amount float64 `alias:"Amount" valid:"Required; check:age > 0" form:"amount" json:"amount"` // 木材数量
}

// Valid 更新库存验证
func (a *StockUpdateValidate) Valid() (err error) {
	return Validate(a)
}
