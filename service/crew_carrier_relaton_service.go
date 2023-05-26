package service

import (
	"woods/mapper"
	"woods/models"
)

type crewCarrierRelationService struct {
	BaseService
}

var CrewCarrierRelationService crewCarrierRelationService

func init() {
	CrewCarrierRelationService = crewCarrierRelationService{}
}

func (u crewCarrierRelationService) GetMapper() *mapper.CrewCarrierRelationMapper {
	return new(mapper.CrewCarrierRelationMapper).GetDb()
}

// Scan 获取列表
func (u crewCarrierRelationService) Scan(filter mapper.ListCrewCarrierRelationFilter) (list []*models.CrewCarrierRelation, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(&filter, field)
	return
}

// First 获取一条数据
func (u crewCarrierRelationService) First(filter mapper.ListCrewCarrierRelationFilter) (out models.CrewCarrierRelation, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u crewCarrierRelationService) Paginate(page uint64, pageSize uint64) (list []*models.CrewCarrierRelation, total uint64, err error) {
	return u.GetMapper().Paginate(page, pageSize)
}

// GetByMobile
func (u crewCarrierRelationService) GetByCrewId(crewId int) (crewCarrierRelation models.CrewCarrierRelation, err error) {
	return u.First(mapper.ListCrewCarrierRelationFilter{CrewId: crewId})
}

// GetByCrewIdAndCarrierId
func (u crewCarrierRelationService) GetByCrewIdAndCarrierId(crewId, carrierId int) (crewCarrierRelation models.CrewCarrierRelation, err error) {
	return u.First(mapper.ListCrewCarrierRelationFilter{CrewId: crewId, CarrierId: carrierId})
}
