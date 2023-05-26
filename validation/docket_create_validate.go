package valid

// 创建Docket验证
type DocketCreateValidate struct {
	StockRecordId int     `alias:"Stock Record Id" valid:"Required; " form:"stock_record_id" json:"stock_record_id"` // 砍伐库存记录ID
	DestinationId int     `alias:"Destination Id" valid:"Required; " form:"destination_id" json:"destination_id"`    // 目的地ID
	OperationType string  `alias:"Operation Type" valid:"Required; " form:"operation_type" json:"operation_type"`    // 砍伐队工作方式
	Customer      string  `alias:"Customer" valid:"Required; " form:"customer" json:"customer"`                      // 客户
	GrossWeight   float64 `alias:"Gross Weight" valid:"Required;" form:"gross_weight" json:"gross_weight"`           // 总重(kg)
	TareWeight    float64 `alias:"Tare Weight" valid:"Required;" form:"tare_weight" json:"tare_weight"`              // 自重(kg)
	NetWeight     float64 `alias:"Net Weight" valid:"Required;" form:"net_weight" json:"net_weight"`                 // 净重(kg)
}

// Valid DocketCreateValidate
func (a *DocketCreateValidate) Valid() (err error) {
	return Validate(a)
}
