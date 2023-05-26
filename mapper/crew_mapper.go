package mapper

import (
	"fmt"
	"time"
	"woods/models"
	bizcode "woods/pkg/bizerror"
	"woods/pkg/util"
)

type CrewMapper struct {
	BaseMapper
}

var CrewModel models.Crew

func (crewMapper *CrewMapper) GetDb() *CrewMapper {
	crewMapper.Db = crewMapper.GetModel(crewMapper.GetTable())
	return crewMapper
}

func (crewMapper *CrewMapper) GetTableModel() *models.Crew {
	return new(models.Crew)
}

func (crewMapper *CrewMapper) GetTable() string {
	model := crewMapper.GetTableModel()
	crewMapper.Table = model.TableName()
	crewMapper.ALis = ""
	crewMapper.Field = model.GetFieldSlice()
	return crewMapper.Table
}

type listCrewFilter struct {
	Page       uint64
	Limit      uint64
	Id         int
	IdNotSelf  int
	UserId     int
	CustomerId int
	Code       string
	CodeLike   string
	Area       string
	Date       int
	Status     int
	DeleteAt   int
	Order      string
}

type ListCrewFilter listCrewFilter

func (crewMapper *CrewMapper) GetWhereOrder(filter *ListCrewFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.DeleteAt, crewMapper.Table+".delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.Id, crewMapper.Table+".id"))
	}
	if filter.IdNotSelf != 0 {
		whereOrder = append(whereOrder, *crewMapper.WhereNEq(filter.IdNotSelf, crewMapper.Table+".id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.CustomerId != 0 {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.CustomerId, "customer_id"))
	}
	if filter.Code != "" {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.Code, "code"))
	}
	if filter.Area != "" {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.Area, "area"))
	}
	if filter.CodeLike != "" {
		whereOrder = append(whereOrder, *crewMapper.WhereLike(filter.CodeLike, "code"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *crewMapper.WhereEq(filter.Status, crewMapper.Table+".status"))
	}
	if filter.Order != "" {
		whereOrder = append(whereOrder, *crewMapper.WhereOrder(filter.Order))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (crewMapper *CrewMapper) GetByCondition(filter *ListCrewFilter, field []string) (list []*models.Crew, total uint64, err error) {
	whereOrder := crewMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = crewMapper.GetDb().GetPage(&list, &models.Crew{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = crewMapper.GetDb().Scan(&list, &models.Crew{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (crewMapper *CrewMapper) FirstByCondition(filter *ListCrewFilter, field []string) (out models.Crew, err error) {
	whereOrder := crewMapper.GetWhereOrder(filter)
	err = crewMapper.GetDb().Find(&out, &models.Crew{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

// Paginate 分页
func (crewMapper *CrewMapper) Paginate(filter *ListCrewFilter) (list []*models.Crew, total uint64, err error) {
	var field []string
	field = []string{crewMapper.GetTable() + ".id", "code", "user_id", "customer_id", "customer_name", crewMapper.GetTable() + ".email", "notes", "setting", "compartment", "forest", "area", "coordinate", "channel", "destination_id", "default_destination", "operation_type", "rate", "round_trip_time", crewMapper.GetTable() + ".status", crewMapper.GetTable() + ".create_time", crewMapper.GetTable() + ".update_time", "t_system_user.email as account"}
	whereOrder := crewMapper.GetWhereOrder(filter)
	fmt.Println("查询条件：", whereOrder)
	SystemUserModel := new(SystemUserMapper).GetTableModel()
	// 添加关联模型
	crewMapper.PreloadList = append(crewMapper.PreloadList, &Preload{Column: "CarrierList", Conditions: []interface{}{"status=? and delete_at=?", Active_Status, 0}})
	crewMapper.PreloadList = append(crewMapper.PreloadList, &Preload{Column: "DestinationList", Conditions: []interface{}{"status=? and delete_at=?", Active_Status, 0}})
	err = crewMapper.Select(field).
		LeftJoin(SystemUserModel, "user_id", "=", "id").
		//Preloads("CarrierList", "status=? and delete_at=?", Active_Status, 0).
		//Preloads("DestinationList", "status=? and delete_at=?", Active_Status, 0).
		GetPage(&list, &models.Crew{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	return
}

func (crewMapper *CrewMapper) Preload(column string, conditions ...interface{}) *CrewMapper {
	crewMapper.Db = crewMapper.Preloads(column, conditions...).Db
	return crewMapper
}

// Select 查询字段 这里也是可用的
//func (crewMapper *CrewMapper) Select(query interface{}, args ...interface{}) *CrewMapper {
//	crewMapper.Db = crewMapper.GetDb().Db.Select(query, args)
//	return crewMapper
//}

// GetCrewList 获取下拉列表
func (crewMapper *CrewMapper) GetCrewList(filter *ListCrewFilter) (list []*models.Crew, err error) {
	var field []string
	stockModel := new(StockMapper).GetTableModel()
	field = []string{crewMapper.GetTable() + ".id", "code", crewMapper.GetTable() + ".user_id", crewMapper.GetTable() + ".customer_id", "customer_name", crewMapper.GetTable() + ".email", "notes", "setting", "compartment", "forest", "area", "channel", "coordinate", "operation_type", "rate", "round_trip_time", crewMapper.GetTable() + ".status", crewMapper.GetTable() + ".create_time", crewMapper.GetTable() + ".update_time", stockModel.TableName() + ".total_amount"}
	whereOrder := crewMapper.GetWhereOrder(filter)
	fmt.Println("查询条件：", whereOrder)
	date := util.GetDateTimestamp(time.Now())
	err = crewMapper.Select(field).Where(stockModel.TableName()+".date=? and "+stockModel.TableName()+".delete_at = 0", date).
		LeftJoin(stockModel, crewMapper.GetTable()+".id", "=", stockModel.TableName()+".crew_id").
		Scan(&list, &models.Crew{}, field, 0, whereOrder...)
	return
}

// GetAllByUserId
func (crewMapper *CrewMapper) GetAllByUserId(userId int) (crew models.Crew, err error) {
	var field []string
	// 添加关联模型
	//crewMapper.PreloadList = append(crewMapper.PreloadList, &Preload{Column: "CarrierList", Conditions: []interface{}{"status=? and delete_at=?", Active_Status, 0}})
	crewMapper.PreloadList = append(crewMapper.PreloadList, &Preload{Column: "DestinationList", Conditions: []interface{}{"status=? and delete_at=?", Active_Status, 0}})
	return crewMapper.FirstByCondition(&ListCrewFilter{UserId: userId, Status: Active_Status}, field)
}
