package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type TaskMapper struct {
	BaseMapper
}

var TaskModel models.Task

func (taskMapper *TaskMapper) GetDb() *TaskMapper {
	taskMapper.Db = taskMapper.GetModel(taskMapper.GetTable())
	return taskMapper
}

func (taskMapper *TaskMapper) GetTableModel() *models.Task {
	return new(models.Task)
}

func (taskMapper *TaskMapper) GetTable() string {
	model := taskMapper.GetTableModel()
	taskMapper.Table = model.TableName()
	taskMapper.ALis = ""
	taskMapper.Field = model.GetFieldSlice()
	return taskMapper.Table
}

type listTaskFilter struct {
	Page      uint64
	Limit     uint64
	Id        int
	UserId    int
	CrewId    int
	VehicleId int
	CarrierId int
	Date      int64
	Status    int
	StatusIn  string
	DeleteAt  int
	Order     string
}

type ListTaskFilter listTaskFilter

const (
	Status_Pending           = 1
	Status_Accept            = 2
	Status_Scan_Docket       = 3
	Status_Scan_Weightbridge = 4
	Status_Finished          = 5
	Status_Decline           = -1
)

func (taskMapper *TaskMapper) GetWhereOrder(filter *ListTaskFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.DeleteAt, taskMapper.GetTable()+".delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.Id, taskMapper.GetTable()+".id"))
	}
	if filter.CrewId != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.CrewId, "crew_id"))
	}
	if filter.VehicleId != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.VehicleId, taskMapper.GetTable()+".vehicle_id"))
	}
	if filter.CarrierId != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.CarrierId, taskMapper.GetTable()+".carrier_id"))
	}
	if filter.Date != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.Date, taskMapper.GetTable()+".date"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *taskMapper.WhereEq(filter.Status, taskMapper.GetTable()+".status"))
	}
	if filter.StatusIn != "" {
		whereOrder = append(whereOrder, *taskMapper.WhereIn(filter.StatusIn, taskMapper.GetTable()+".status"))
	}
	if filter.Order != "" {
		whereOrder = append(whereOrder, *taskMapper.WhereOrder(filter.Order))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (taskMapper *TaskMapper) GetByCondition(filter *ListTaskFilter, field []string) (list []*models.Task, total uint64, err error) {
	whereOrder := taskMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = taskMapper.GetDb().GetPage(&list, &models.Task{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = taskMapper.GetDb().Scan(&list, &models.Task{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (taskMapper *TaskMapper) FirstByCondition(filter *ListTaskFilter, field []string) (out models.Task, err error) {
	whereOrder := taskMapper.GetWhereOrder(filter)
	err = taskMapper.GetDb().Find(&out, &models.Task{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

// Paginate 分页
func (taskMapper *TaskMapper) Paginate(filter *ListTaskFilter) (list []*models.Task, total uint64, err error) {
	whereOrder := taskMapper.GetWhereOrder(filter)
	var field []string
	field = []string{taskMapper.GetTable() + ".id", "crew_id", "vehicle_id", "vehicle_code", "start_time", "end_time", "rate", "notes", "destination", "decline_reason", "weightbridge_fail_remark", taskMapper.GetTable() + ".status", taskMapper.GetTable() + ".create_time", taskMapper.GetTable() + ".update_time"}
	err = taskMapper.Select(field).Preloads("DocketList", "delete_at = 0").Find(&list, &models.Task{}, field, 0, whereOrder...)
	return
}

func (taskMapper *TaskMapper) Preload(column string, conditions ...interface{}) *TaskMapper {
	taskMapper.Db = taskMapper.Preloads(column, conditions...).Db
	return taskMapper
}
