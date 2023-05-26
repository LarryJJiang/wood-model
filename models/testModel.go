package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 质控考核题目答案表
type TestModel struct {
	Model
	CaseId             int64  `gorm:"column:case_id; size:20; " json:"case_id"`                                   // 病例id：检查的开始时间(时间戳：秒)
	QuestionId         string `gorm:"column:question_id; size:20; " json:"question_id"`                           // 题目id
	QuestionName       string `gorm:"column:question_name; size:20; " json:"question_name"`                       // 题目的名称
	ShotId             string `gorm:"column:shot_id; size:200;" json:"shot_id"`                                   //  截图ID(也就是题目的答案)
	AnswerScore        int32  `gorm:"column:answer_score; size:5;" json:"answer_score"`                           //答题得分
	GroupId            string `gorm:"column:group_id; size:100; not null; default:'';" json:"group_id"`           //  分组id
	BoxNum             int32  `gorm:"column:box_num; size:2;" json:"box_num"`                                     // 方框数量
	BoxIndex           int32  `gorm:"column:box_index; size:2;" json:"box_index"`                                 // 方框index
	QuestionTotalScore int32  `gorm:"column:question_total_score; size:5;" json:"question_total_score"`           // 题目总分
	SerialNumber       string `gorm:"column:serial_number; size:20; not null; default:''; " json:"serial_number"` // 所属设备序列号
	AddExtra           string `gorm:"column:add_extra; size:512; not null; default:''" json:"add_extra"`          // 扩充
}

func (m *TestModel) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "assess_question_answer"
}

func (m *TestModel) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
