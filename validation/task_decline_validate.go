package valid

// 货车拒绝接单
type TaskDeclineValidate struct {
	DeclineReason string `alias:"Decline Reason" valid:"Required; " form:"decline_reason" json:"decline_reason"` // 拒绝接单原因
}

// Valid 货车拒绝接单验证
func (a *TaskDeclineValidate) Valid() (err error) {
	return Validate(a)
}
