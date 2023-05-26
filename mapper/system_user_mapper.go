package mapper

import (
	"woods/models"
	bizcode "woods/pkg/bizerror"
)

type SystemUserMapper struct {
	BaseMapper
}

var SystemUserModel models.SystemUser

func (systemUserMapper *SystemUserMapper) GetDb() *SystemUserMapper {
	systemUserMapper.Db = systemUserMapper.GetModel(systemUserMapper.GetTable())
	return systemUserMapper
}

func (systemUserMapper *SystemUserMapper) GetTableModel() *models.SystemUser {
	return new(models.SystemUser)
}

func (systemUserMapper *SystemUserMapper) GetTable() string {
	model := systemUserMapper.GetTableModel()
	systemUserMapper.Table = model.TableName()
	systemUserMapper.ALis = ""
	systemUserMapper.Field = model.GetFieldSlice()
	return systemUserMapper.Table
}

type listSystemUserFilter struct {
	Page      uint64
	Limit     uint64
	Id        int
	IdNotSelf int
	Identity  int
	Status    int
	DeleteAt  int
	Account   string
	UserName  string
}

const Identity_SYS = 1
const Identity_TRUCK_DRIVER = 2
const Identity_CREW = 3

var UserIdentity = map[string]int{
	"SYS":          1,
	"TRUCK_DRIVER": 2,
	"CREW":         3,
}

type ListSystemUserFilter listSystemUserFilter

func (systemUserMapper *SystemUserMapper) GetWhereOrder(filter *ListSystemUserFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.DeleteAt, "delete_at"))
	if filter.Identity != 0 {
		whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.Identity, "identity"))
	}
	if filter.IdNotSelf != 0 {
		whereOrder = append(whereOrder, *systemUserMapper.WhereNEq(filter.IdNotSelf, "id"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.Id, "id"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.Status, "status"))
	}
	if filter.Account != "" {
		whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.Account, "account"))
	}
	if filter.UserName != "" {
		whereOrder = append(whereOrder, *systemUserMapper.WhereEq(filter.UserName, "user_name"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (systemUserMapper *SystemUserMapper) GetByCondition(filter *ListSystemUserFilter, field []string) (list []*models.SystemUser, total uint64, err error) {
	whereOrder := systemUserMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = systemUserMapper.GetDb().GetPage(&list, &models.SystemUser{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = systemUserMapper.GetDb().Scan(&list, &models.SystemUser{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (systemUserMapper *SystemUserMapper) FirstByCondition(filter *ListSystemUserFilter, field []string) (out models.SystemUser, err error) {
	whereOrder := systemUserMapper.GetWhereOrder(filter)
	err = systemUserMapper.GetDb().Find(&out, &models.SystemUser{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (systemUserMapper *SystemUserMapper) Paginate(filter *ListSystemUserFilter) (list []*models.SystemUser, total uint64, err error) {
	var field []string
	return systemUserMapper.GetByCondition(filter, field)
}

// DeleteSystemUser
func (systemUserMapper *SystemUserMapper) DeleteSystemUser(id int) error {
	return systemUserMapper.SoftDeleteById(id)
}
