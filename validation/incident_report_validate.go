package valid

// 事故报告验证
type IncidentReportValidate struct {
	IncidentType string `alias:"Incident Type" valid:"Required; " form:"incident_type" json:"incident_type"` // 事件类型，多选，用英文逗号隔开
	Location     string `alias:"Location" valid:"Required; " form:"location" json:"location"`                // 位置
	Forest       string `alias:"Forest" valid:"Required; " form:"forest" json:"forest"`                      // 林场名称
	Remark       string `alias:"Remark" valid:"Required; " form:"remark" json:"remark"`                      // 事项说明
	Images       string `alias:"Images" form:"images" json:"images"`                                         // 事项拍照
}

// Valid 事故报告验证
func (a *IncidentReportValidate) Valid() (err error) {
	return Validate(a)
}
