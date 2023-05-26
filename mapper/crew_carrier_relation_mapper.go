package mapper

import (
	"woods/models"
	bizcode "woods/pkg/bizerror"
)

type CrewCarrierRelationMapper struct {
	BaseMapper
}

var CrewCarrierRelationModel models.CrewCarrierRelation

func (crewCarrierRelationMapper *CrewCarrierRelationMapper) GetDb() *CrewCarrierRelationMapper {
	crewCarrierRelationMapper.Db = crewCarrierRelationMapper.GetModel(crewCarrierRelationMapper.GetTable())
	return crewCarrierRelationMapper
}

func (crewCarrierRelationMapper *CrewCarrierRelationMapper) GetTableModel() *models.CrewCarrierRelation {
	return new(models.CrewCarrierRelation)
}

func (crewCarrierRelationMapper *CrewCarrierRelationMapper) GetTable() string {
	model := crewCarrierRelationMapper.GetTableModel()
	crewCarrierRelationMapper.Table = model.TableName()
	crewCarrierRelationMapper.ALis = ""
	crewCarrierRelationMapper.Field = model.GetFieldSlice()
	return crewCarrierRelationMapper.Table
}

type listCrewCarrierRelationFilter struct {
	Page      uint64
	Limit     uint64
	Id        int
	CrewId    int
	CarrierId int
	Status    int
	DeleteAt  int
}

type ListCrewCarrierRelationFilter listCrewCarrierRelationFilter

func (crewCarrierRelationMapper *CrewCarrierRelationMapper) GetWhereOrder(filter *ListCrewCarrierRelationFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *crewCarrierRelationMapper.WhereEq(filter.DeleteAt, "delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *crewCarrierRelationMapper.WhereEq(filter.Id, "id"))
	}
	if filter.CrewId != 0 {
		whereOrder = append(whereOrder, *crewCarrierRelationMapper.WhereEq(filter.CrewId, "crew_id"))
	}
	if filter.CarrierId != 0 {
		whereOrder = append(whereOrder, *crewCarrierRelationMapper.WhereEq(filter.CarrierId, "carrier_id"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *crewCarrierRelationMapper.WhereEq(filter.Status, "status"))
	}
	//crewCarrierRelationMapper.Db = crewCarrierRelationMapper.Db.Where(whereOrder)
	return whereOrder
}

// GetByCondition 获取列表
func (crewCarrierRelationMapper *CrewCarrierRelationMapper) GetByCondition(filter *ListCrewCarrierRelationFilter, field []string) (list []*models.CrewCarrierRelation, total uint64, err error) {
	whereOrder := crewCarrierRelationMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = crewCarrierRelationMapper.GetDb().GetPage(&list, &models.CrewCarrierRelation{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = crewCarrierRelationMapper.GetDb().Scan(&list, &models.CrewCarrierRelation{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (crewCarrierRelationMapper *CrewCarrierRelationMapper) FirstByCondition(filter *ListCrewCarrierRelationFilter, field []string) (out models.CrewCarrierRelation, err error) {
	whereOrder := crewCarrierRelationMapper.GetWhereOrder(filter)
	err = crewCarrierRelationMapper.GetDb().Find(&out, &models.CrewCarrierRelation{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (crewCarrierRelationMapper *CrewCarrierRelationMapper) Paginate(page uint64, pageSize uint64) (list []*models.CrewCarrierRelation, total uint64, err error) {
	var field []string
	return crewCarrierRelationMapper.GetByCondition(&ListCrewCarrierRelationFilter{Page: page, Limit: pageSize}, field)
}
