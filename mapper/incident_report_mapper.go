package mapper

import (
	"woods/models"
	bizcode "woods/pkg/bizerror"
)

type IncidentReportMapper struct {
	BaseMapper
}

var IncidentReportModel models.IncidentReport

func (userMapper *IncidentReportMapper) GetDb() *IncidentReportMapper {
	userMapper.Db = userMapper.GetModel(userMapper.GetTable())
	return userMapper
}

func (userMapper *IncidentReportMapper) GetTableModel() *models.IncidentReport {
	return new(models.IncidentReport)
}

func (userMapper *IncidentReportMapper) GetTable() string {
	model := userMapper.GetTableModel()
	userMapper.Table = model.TableName()
	userMapper.ALis = ""
	userMapper.Field = model.GetFieldSlice()
	return userMapper.Table
}

type listIncidentReportFilter struct {
	Page      uint64
	Limit     uint64
	Id        int
	UserId    int
	VehicleId int
	Forest    string
	StartTime int
	StopTime  int
}

type ListIncidentReportFilter listIncidentReportFilter

func (userMapper *IncidentReportMapper) GetWhereOrder(filter *ListIncidentReportFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *userMapper.WhereEq(filter.Id, "id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *userMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.VehicleId != 0 {
		whereOrder = append(whereOrder, *userMapper.WhereLike(filter.VehicleId, "vehicle_id"))
	}
	if filter.StartTime > 0 {
		whereOrder = append(whereOrder, *userMapper.WhereEGT(filter.StartTime, "add_time"))
	}
	if filter.StopTime > 0 {
		whereOrder = append(whereOrder, *userMapper.WhereELT(filter.StopTime, "add_time"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (userMapper *IncidentReportMapper) GetByCondition(filter *ListIncidentReportFilter, field []string) (list []*models.IncidentReport, total uint64, err error) {
	whereOrder := userMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = userMapper.GetDb().GetPage(&list, &models.IncidentReport{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = userMapper.GetDb().Scan(&list, &models.IncidentReport{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (userMapper *IncidentReportMapper) FirstByCondition(filter *ListIncidentReportFilter, field []string) (out models.IncidentReport, err error) {
	whereOrder := userMapper.GetWhereOrder(filter)
	err = userMapper.GetDb().Find(&out, &models.IncidentReport{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (userMapper *IncidentReportMapper) Paginate(filter *ListIncidentReportFilter) (list []*models.IncidentReport, total uint64, err error) {
	var field []string
	return userMapper.GetByCondition(filter, field)
}
