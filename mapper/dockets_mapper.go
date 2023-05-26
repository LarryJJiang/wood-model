package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type DocketsMapper struct {
	BaseMapper
}

var DocketsModel models.Dockets

func (docketsMapper *DocketsMapper) GetDb() *DocketsMapper {
	docketsMapper.Db = docketsMapper.GetModel(docketsMapper.GetTable())
	return docketsMapper
}

func (docketsMapper *DocketsMapper) GetTableModel() *models.Dockets {
	return new(models.Dockets)
}

func (docketsMapper *DocketsMapper) GetTable() string {
	model := docketsMapper.GetTableModel()
	docketsMapper.Table = model.TableName()
	docketsMapper.ALis = ""
	docketsMapper.Field = model.GetFieldSlice()
	return docketsMapper.Table
}

type listDocketsFilter struct {
	Page      uint64
	Limit     uint64
	Id        int
	UserId    int
	TaskId    int
	CrewId    int
	VehicleId int
	Date      int
	Status    int
	DeleteAt  int
	Order     string
}

type ListDocketsFilter listDocketsFilter

func (docketsMapper *DocketsMapper) GetWhereOrder(filter *ListDocketsFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.DeleteAt, docketsMapper.Table+".delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.Id, docketsMapper.Table+".id"))
	}
	if filter.TaskId != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.TaskId, "task_id"))
	}
	if filter.CrewId != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.CrewId, "crew_id"))
	}
	if filter.VehicleId != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.VehicleId, "vehicle_id"))
	}
	if filter.Date != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.Date, "date"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *docketsMapper.WhereEq(filter.Status, docketsMapper.Table+".status"))
	}
	if filter.Order != "" {
		whereOrder = append(whereOrder, *docketsMapper.WhereOrder(filter.Order))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (docketsMapper *DocketsMapper) GetByCondition(filter *ListDocketsFilter, field []string) (list []*models.Dockets, total uint64, err error) {
	whereOrder := docketsMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = docketsMapper.GetDb().GetPage(&list, &models.Dockets{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = docketsMapper.GetDb().Scan(&list, &models.Dockets{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (docketsMapper *DocketsMapper) FirstByCondition(filter *ListDocketsFilter, field []string) (out models.Dockets, err error) {
	whereOrder := docketsMapper.GetWhereOrder(filter)
	err = docketsMapper.GetDb().Find(&out, &models.Dockets{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

// Paginate 分页
func (docketsMapper *DocketsMapper) Paginate(filter *ListDocketsFilter) (list []*models.Dockets, total uint64, err error) {
	var field []string
	return docketsMapper.GetByCondition(filter, field)
}

func (docketsMapper *DocketsMapper) Preload(column string, conditions ...interface{}) *DocketsMapper {
	docketsMapper.Db = docketsMapper.Preloads(column, conditions...).Db
	return docketsMapper
}

// Select 查询字段 这里也是可用的
//func (docketsMapper *DocketsMapper) Select(query interface{}, args ...interface{}) *DocketsMapper {
//	docketsMapper.Db = docketsMapper.GetDb().Db.Select(query, args)
//	return docketsMapper
//}
