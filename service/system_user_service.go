package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wood-client-api/mapper"
	"wood-client-api/models"
	"wood-client-api/pkg/crypt"
)

type systemUserService struct {
	BaseService
}

var SystemUserService systemUserService

func init() {
	SystemUserService = systemUserService{}
}

func (su *systemUserService) GetMapper() *mapper.SystemUserMapper {
	return new(mapper.SystemUserMapper).GetDb()
}

// Scan 获取列表
func (su *systemUserService) Scan(filter *mapper.ListSystemUserFilter) (list []*models.SystemUser, total uint64, err error) {
	var field []string
	list, total, err = su.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (su *systemUserService) First(filter mapper.ListSystemUserFilter) (out models.SystemUser, err error) {
	var field []string
	out, err = su.GetMapper().FirstByCondition(&filter, field)
	return
}

func (su *systemUserService) Paginate(filter *mapper.ListSystemUserFilter) (list []*models.SystemUser, total uint64, err error) {
	return su.GetMapper().Paginate(filter)
}

// GetByAccount 通过账号获取一条用户信息
func (su *systemUserService) GetByAccount(account string) (user models.SystemUser, err error) {
	return su.First(mapper.ListSystemUserFilter{Account: account})
}

// GetByUserName 通过用户名获取一条用户信息
func (su *systemUserService) GetByUserName(userName string) (user models.SystemUser, err error) {
	return su.First(mapper.ListSystemUserFilter{UserName: userName})
}

func (su *systemUserService) CheckPwd(systemUser *models.SystemUser) (*models.SystemUser, error) {
	findUser, err := su.GetByAccount(systemUser.Account)
	isNotFound := gorm.IsRecordNotFoundError(err)
	if err != nil && !isNotFound {
		return nil, err
	}
	if isNotFound {
		return nil, errors.New("Account is not exists.")
	}
	systemUser.Password = crypt.GetSystemPassword(systemUser.Password, findUser.Salt)
	if findUser.Status == 0 || findUser.DeleteAt > 0 || !crypt.VerifyCryptPwd(findUser.Password, systemUser.Password) {
		return nil, errors.New("Password is incorrect.")
	}
	findUser.Salt = ""
	findUser.Password = ""
	return &findUser, nil
}

// GetUserById 通过user_id获取一条用户信息
func (su *systemUserService) GetUserById(userId int) (user models.SystemUser, err error) {
	return su.First(mapper.ListSystemUserFilter{Id: userId, DeleteAt: 0})
}
