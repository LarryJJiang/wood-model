package service

import (
	"woods/mapper"
	"woods/models"
)

type userService struct {
	BaseService
}

var UserService userService

func init() {
	UserService = userService{}
}

func (u userService) GetMapper() *mapper.UserMapper {
	return new(mapper.UserMapper).GetDb()
}

// Scan 获取列表
func (u userService) Scan(filter mapper.ListUserFilter) (list []*models.User, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(&filter, field)
	return
}

// First 获取一条数据
func (u userService) First(filter mapper.ListUserFilter) (out models.User, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u userService) Paginate(page uint64, pageSize uint64) (list []*models.User, total uint64, err error) {
	return u.GetMapper().Paginate(page, pageSize)
}

// GetByMobile
func (u userService) GetByMobile(mobile string) (user models.User, err error) {
	return u.First(mapper.ListUserFilter{Mobile: mobile})
}
