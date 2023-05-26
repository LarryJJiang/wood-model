package service

import (
	"github.com/jinzhu/gorm"
	"wood-client-api/mapper"
	"wood-client-api/models"
	"wood-client-api/pkg/util"
)

type customerService struct {
	BaseService
}

var CustomerService customerService

func init() {
	CustomerService = customerService{}
}

func (u customerService) GetMapper() *mapper.CustomerMapper {
	return new(mapper.CustomerMapper).GetDb()
}

// Scan 获取列表
func (u customerService) Scan(filter *mapper.ListCustomerFilter) (list []*models.Customer, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u customerService) First(filter mapper.ListCustomerFilter) (out models.Customer, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

// GetByName
func (u customerService) GetByName(name string) (user models.Customer, err error) {
	return u.First(mapper.ListCustomerFilter{Name: name, Status: mapper.Active_Status})
}

// GetByUserId
func (u customerService) GetByUserId(userId int) (user models.Customer, err error) {
	return u.First(mapper.ListCustomerFilter{UserId: userId, Status: mapper.Active_Status})
}

// GetById
func (u customerService) GetById(id int) (user models.Customer, err error) {
	return u.First(mapper.ListCustomerFilter{Id: id, Status: mapper.Active_Status})
}

//删除林场账号
func (u customerService) DeleteCustomer(id int) error {
	crew, err := u.GetById(id)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Customer")
	}
	return u.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&models.Customer{}).Where(map[string]interface{}{"id": id}).Update(map[string]interface{}{
			"delete_at":   util.Now().Year(),
			"update_time": util.NowMilli(),
		}).Error
		if !u.IsErrorNil(err) {
			return err
		}
		err = tx.Model(&models.SystemUser{}).Where(map[string]interface{}{"id": crew.UserId}).Update(map[string]interface{}{
			"delete_at":   util.Now().Year(),
			"update_time": util.NowMilli(),
		}).Error
		if !u.IsErrorNil(err) {
			return err
		}
		return nil
	})
}
