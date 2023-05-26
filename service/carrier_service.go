package service

import (
	"github.com/jinzhu/gorm"
	"wood-client-api/mapper"
	"wood-client-api/models"
	"wood-client-api/pkg/util"
)

type carrierService struct {
	BaseService
}

var CarrierService carrierService

func init() {
	CarrierService = carrierService{}
}

func (u carrierService) GetMapper() *mapper.CarrierMapper {
	return new(mapper.CarrierMapper).GetDb()
}

// Scan 获取列表
func (u carrierService) Scan(filter *mapper.ListCarrierFilter) (list []*models.Carrier, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u carrierService) First(filter mapper.ListCarrierFilter) (out models.Carrier, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

// GetByName 通过名称获取一条用户信息
func (u carrierService) GetByName(name string) (user models.Carrier, err error) {
	return u.First(mapper.ListCarrierFilter{Name: name, Status: mapper.Active_Status})
}

// GetByUserId 通过用户ID获取一条用户信息
func (u carrierService) GetByUserId(userId int) (user models.Carrier, err error) {
	return u.First(mapper.ListCarrierFilter{UserId: userId, Status: mapper.Active_Status})
}

// GetById 通过ID获取一条用户信息
func (u carrierService) GetById(id int) (user models.Carrier, err error) {
	return u.First(mapper.ListCarrierFilter{Id: id, Status: mapper.Active_Status})
}

// GetByIds 通过ID获取一条用户信息
func (u carrierService) GetByIds(id string) (user []*models.Carrier, err error) {
	user, _, err = u.Scan(&mapper.ListCarrierFilter{IdIn: id, Status: mapper.Active_Status})
	return
}

//删除运输公司账号
func (u carrierService) DeleteCarrier(id int) error {
	crew, err := u.GetById(id)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Carrier")
	}
	return u.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&models.Carrier{}).Where(map[string]interface{}{"id": id}).Update(map[string]interface{}{
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
