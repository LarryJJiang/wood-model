package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type CustomerMapper struct {
	BaseMapper
}

var CustomerModel models.Customer

func (customerMapper *CustomerMapper) GetDb() *CustomerMapper {
	customerMapper.Db = customerMapper.GetModel(customerMapper.GetTable())
	return customerMapper
}

func (customerMapper *CustomerMapper) GetTableModel() *models.Customer {
	return new(models.Customer)
}

func (customerMapper *CustomerMapper) GetTable() string {
	model := customerMapper.GetTableModel()
	customerMapper.Table = model.TableName()
	customerMapper.ALis = ""
	customerMapper.Field = model.GetFieldSlice()
	return customerMapper.Table
}

type listCustomerFilter struct {
	Page          uint64
	Limit         uint64
	Id            int
	IdNotSelf     int
	UserId        int
	UserIdNotSelf int
	Name          string
	NameLike      string
	Email         string
	Status        int
	DeleteAt      int
}

type ListCustomerFilter listCustomerFilter

func (customerMapper *CustomerMapper) GetWhereOrder(filter *ListCustomerFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.DeleteAt, customerMapper.GetTable()+".delete_at"))
	if filter.Name != "" {
		whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.Name, "name"))
	}
	if filter.NameLike != "" {
		whereOrder = append(whereOrder, *customerMapper.WhereLike(filter.NameLike, "name"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.Id, customerMapper.GetTable()+".id"))
	}
	if filter.IdNotSelf != 0 {
		whereOrder = append(whereOrder, *customerMapper.WhereNEq(filter.IdNotSelf, customerMapper.GetTable()+".id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.UserIdNotSelf != 0 {
		whereOrder = append(whereOrder, *customerMapper.WhereNEq(filter.UserIdNotSelf, "user_id"))
	}
	if filter.Email != "" {
		whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.Email, customerMapper.GetTable()+".email"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *customerMapper.WhereEq(filter.Status, customerMapper.GetTable()+".status"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (customerMapper *CustomerMapper) GetByCondition(filter *ListCustomerFilter, field []string) (list []*models.Customer, total uint64, err error) {
	whereOrder := customerMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = customerMapper.GetDb().GetPage(&list, &models.Customer{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = customerMapper.GetDb().Scan(&list, &models.Customer{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (customerMapper *CustomerMapper) FirstByCondition(filter *ListCustomerFilter, field []string) (out models.Customer, err error) {
	whereOrder := customerMapper.GetWhereOrder(filter)
	err = customerMapper.GetDb().Find(&out, &models.Customer{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}
