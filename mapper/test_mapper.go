package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type TestMapper struct {
	BaseMapper
}

func (testMapper *TestMapper) GetDb() *TestMapper {
	testMapper.Db = testMapper.GetModel(testMapper.GetTable())
	return testMapper
}

func (testMapper *TestMapper) GetTableModel() *models.TestModel {
	return new(models.TestModel)
}

func (testMapper *TestMapper) GetTable() string {
	model := testMapper.GetTableModel()
	testMapper.Table = model.TableName()
	testMapper.ALis = ""
	testMapper.Field = model.GetFieldSlice()
	return testMapper.Table
}

type listFilter struct {
	Page         uint64
	Limit        uint64
	SerialNumber string
	CaseId       int64
	QuestionId   string
	QuestionName string
	ShotId       string
	GroupId      string
	StartTime    int
	StopTime     int
}

type ListFilter listFilter

type AssessQuestionAnswerRes struct {
	AddTime            int    `json:"add_time"`
	CaseId             int64  `json:"case_id"`
	QuestionId         string `json:"question_id"`
	QuestionName       string `json:"question_name"`
	ShotId             string `json:"shot_id"`
	AnswerScore        int32  `json:"answer_score"`
	GroupId            string `json:"group_id"`
	BoxNum             int32  `json:"box_num"`
	BoxIndex           int32  `json:"box_index"`
	QuestionTotalScore int32  `json:"question_total_score"`
	AddExtra           string `json:"add_extra"`
}

func (testMapper *TestMapper) GetWhereOrder(filter *ListFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	if filter.SerialNumber != "" {
		whereOrder = append(whereOrder, *testMapper.WhereEq(filter.SerialNumber, "serial_number"))
	}
	if filter.CaseId != 0 {
		whereOrder = append(whereOrder, *testMapper.WhereEq(filter.CaseId, "case_id"))
	}
	if filter.QuestionId != "" {
		whereOrder = append(whereOrder, *testMapper.WhereEq(filter.QuestionId, "question_id"))
	}
	if filter.QuestionName != "" {
		whereOrder = append(whereOrder, *testMapper.WhereLike(filter.QuestionName, "question_name"))
	}
	if filter.ShotId != "" {
		whereOrder = append(whereOrder, *testMapper.WhereLike(filter.ShotId, "shot_id"))
	}
	if filter.GroupId != "" {
		whereOrder = append(whereOrder, *testMapper.WhereLike(filter.GroupId, "group_id"))
	}
	if filter.StartTime > 0 {
		whereOrder = append(whereOrder, *testMapper.WhereEGT(filter.StartTime, "add_time"))
	}
	if filter.StopTime > 0 {
		whereOrder = append(whereOrder, *testMapper.WhereELT(filter.StopTime, "add_time"))
	}
	return whereOrder
}

// 获取列表
func (testMapper *TestMapper) GetByCondition(filter *ListFilter, field []string) (list []*models.TestModel, total uint64, err error) {
	whereOrder := testMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = testMapper.GetDb().GetPage(&list, &models.TestModel{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = testMapper.GetDb().Scan(&list, &models.TestModel{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// 获取一条数据
func (testMapper *TestMapper) FirstByCondition(filter *ListFilter, field []string) (out models.TestModel, err error) {
	whereOrder := testMapper.GetWhereOrder(filter)
	err = testMapper.GetDb().Find(&out, &models.TestModel{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}
