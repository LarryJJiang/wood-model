package service

import (
	"wood-client-api/mapper"
	"wood-client-api/models"
)

type destinationService struct {
	BaseService
}

var DestinationService destinationService

func init() {
	DestinationService = destinationService{}
}

func (ds destinationService) GetMapper() *mapper.DestinationMapper {
	return new(mapper.DestinationMapper).GetDb()
}

// Scan 获取列表
func (ds destinationService) Scan(filter *mapper.ListDestinationFilter) (list []*models.Destination, total uint64, err error) {
	var field []string
	list, total, err = ds.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (ds destinationService) First(filter mapper.ListDestinationFilter) (out models.Destination, err error) {
	var field []string
	out, err = ds.GetMapper().FirstByCondition(&filter, field)
	return
}

func (ds destinationService) Paginate(filter *mapper.ListDestinationFilter) (list []*models.Destination, total uint64, err error) {
	return ds.GetMapper().Paginate(filter)
}

// GetByName
func (ds destinationService) GetByName(name string) (destination models.Destination, err error) {
	return ds.First(mapper.ListDestinationFilter{Name: name, Status: mapper.Active_Status})
}

// GetByCode
func (ds destinationService) GetByCode(code string) (destination models.Destination, err error) {
	return ds.First(mapper.ListDestinationFilter{Code: code, Status: mapper.Active_Status})
}

// GetById
func (ds destinationService) GetById(id int) (destination models.Destination, err error) {
	return ds.First(mapper.ListDestinationFilter{Id: id, Status: mapper.Active_Status})
}

// GetByIds
func (ds destinationService) GetByIds(ids string) (destination []*models.Destination, total uint64, err error) {
	return ds.Scan(&mapper.ListDestinationFilter{IdIn: ids, Status: mapper.Active_Status})
}
