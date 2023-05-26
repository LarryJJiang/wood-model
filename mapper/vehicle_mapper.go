package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type VehicleMapper struct {
	BaseMapper
}

var VehicleModel models.Vehicle

func (vehicleMapper *VehicleMapper) GetDb() *VehicleMapper {
	vehicleMapper.Db = vehicleMapper.GetModel(vehicleMapper.GetTable())
	return vehicleMapper
}

func (vehicleMapper *VehicleMapper) GetTableModel() *models.Vehicle {
	return new(models.Vehicle)
}

func (vehicleMapper *VehicleMapper) GetTable() string {
	model := vehicleMapper.GetTableModel()
	vehicleMapper.Table = model.TableName()
	vehicleMapper.ALis = ""
	vehicleMapper.Field = model.GetFieldSlice()
	return vehicleMapper.Table
}

type listVehicleFilter struct {
	Page             uint64
	Limit            uint64
	Id               int
	IdIn             string
	IdNotSelf        int
	UserId           int
	RegistrationLike string
	Registration     string
	Code             string
	CarrierId        int
	Status           int
	DeleteAt         int
}

type ListVehicleFilter listVehicleFilter

func (vehicleMapper *VehicleMapper) GetWhereOrder(filter *ListVehicleFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.DeleteAt, vehicleMapper.GetTable()+".delete_at"))
	if filter.Registration != "" {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.Registration, "registration"))
	}
	if filter.RegistrationLike != "" {
		whereOrder = append(whereOrder, *vehicleMapper.WhereLike(filter.RegistrationLike, "registration"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.Id, vehicleMapper.GetTable()+".id"))
	}
	if filter.IdIn != "" {
		whereOrder = append(whereOrder, *vehicleMapper.WhereIn(filter.IdIn, vehicleMapper.GetTable()+".id"))
	}
	if filter.IdNotSelf != 0 {
		whereOrder = append(whereOrder, *vehicleMapper.WhereNEq(filter.IdNotSelf, vehicleMapper.GetTable()+".id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.Code != "" {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.Code, vehicleMapper.GetTable()+".code"))
	}
	if filter.CarrierId != 0 {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.CarrierId, vehicleMapper.GetTable()+".carrier_id"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *vehicleMapper.WhereEq(filter.Status, vehicleMapper.GetTable()+".status"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (vehicleMapper *VehicleMapper) GetByCondition(filter *ListVehicleFilter, field []string) (list []*models.Vehicle, total uint64, err error) {
	whereOrder := vehicleMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = vehicleMapper.GetDb().GetPage(&list, &models.Vehicle{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = vehicleMapper.GetDb().Scan(&list, &models.Vehicle{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (vehicleMapper *VehicleMapper) FirstByCondition(filter *ListVehicleFilter, field []string) (out models.Vehicle, err error) {
	whereOrder := vehicleMapper.GetWhereOrder(filter)
	err = vehicleMapper.GetDb().Find(&out, &models.Vehicle{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

// Paginate 分页
func (vehicleMapper *VehicleMapper) Paginate(filter *ListVehicleFilter) (list []*models.Vehicle, total uint64, err error) {
	var field []string
	field = []string{vehicleMapper.GetTable() + ".id", "user_id", "registration", "code", "pin_code", "carrier_id", "carrier_name", "app_version", "car_hopper_count", vehicleMapper.GetTable() + ".status", vehicleMapper.GetTable() + ".create_time", vehicleMapper.GetTable() + ".update_time", "t_system_user.email as account"}
	whereOrder := vehicleMapper.GetWhereOrder(filter)
	SystemUserModel := new(SystemUserMapper).GetTableModel()
	err = vehicleMapper.Select(field).LeftJoin(SystemUserModel, "user_id", "=", "id").GetPage(&list, &models.Crew{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	return
}
