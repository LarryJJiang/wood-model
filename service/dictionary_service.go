package service

import (
	"woods/mapper"
	"woods/models"
)

type dictionaryService struct {
	BaseService
}

var DictionaryService dictionaryService

func init() {
	DictionaryService = dictionaryService{}
}

func (u dictionaryService) GetMapper() *mapper.DictionaryMapper {
	return new(mapper.DictionaryMapper).GetDb()
}

type Dictionary struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Scan 获取列表
func (u dictionaryService) Scan(filter *mapper.ListDictionaryFilter) ([]*Dictionary, uint64, error) {
	var field []string
	list, total, err := u.GetMapper().GetByCondition(filter, field)
	var dictionaryList []*Dictionary
	u.FilterField(list, &dictionaryList)
	return dictionaryList, total, err
}

// First 获取一条数据
func (u dictionaryService) First(filter mapper.ListDictionaryFilter) (out models.Dictionary, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

// GetByName 通过名称获取一条用户信息
func (u dictionaryService) GetByName(name string) (user models.Dictionary, err error) {
	return u.First(mapper.ListDictionaryFilter{Name: name})
}

// GetById 通过ID获取一条用户信息
func (u dictionaryService) GetById(id int) (user models.Dictionary, err error) {
	return u.First(mapper.ListDictionaryFilter{Id: id})
}

// GetById 通过ID获取一条用户信息
func (u dictionaryService) GetIncidentReportingCase() (list []*Dictionary, err error) {
	filter := &mapper.ListDictionaryFilter{Name: "incident_reporting_case"}
	list, _, err = u.Scan(filter)
	return
}

// GetById 通过ID获取一条用户信息
func (u dictionaryService) GetDeclineJobReason() (list []*Dictionary, err error) {
	filter := &mapper.ListDictionaryFilter{Name: "decline_job_reason"}
	list, _, err = u.Scan(filter)
	return
}
