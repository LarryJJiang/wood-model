package service

import (
	"woods/mapper"
	"woods/models"
)

type testService struct {
	BaseService
}

var TestService testService

func init() {
	TestService = testService{}
}

func (answerService testService) GetMapper() *mapper.TestMapper {
	return new(mapper.TestMapper).GetDb()
}

// 获取列表
func (answerService testService) Scan(filter mapper.ListFilter) (list []*models.TestModel, total uint64, err error) {
	var field []string
	list, total, err = answerService.GetMapper().GetByCondition(&filter, field)
	return
}

// 获取一条数据
func (answerService testService) First(filter mapper.ListFilter) (out models.TestModel, err error) {
	var field []string
	out, err = answerService.GetMapper().FirstByCondition(&filter, field)
	return
}
