package mapper

import (
	"woods/models"
	bizcode "woods/pkg/bizerror"
)

type UserMapper struct {
	BaseMapper
}

var UserModel models.User

func (userMapper *UserMapper) GetDb() *UserMapper {
	userMapper.Db = userMapper.GetModel(userMapper.GetTable())
	return userMapper
}

func (userMapper *UserMapper) GetTableModel() *models.User {
	return new(models.User)
}

func (userMapper *UserMapper) GetTable() string {
	model := userMapper.GetTableModel()
	userMapper.Table = model.TableName()
	userMapper.ALis = ""
	userMapper.Field = model.GetFieldSlice()
	return userMapper.Table
}

type listUserFilter struct {
	Page         uint64
	Limit        uint64
	Mobile       string
	Id           int
	QuestionId   string
	QuestionName string
	ShotId       string
	GroupId      string
	StartTime    int
	StopTime     int
}

type ListUserFilter listUserFilter

func (userMapper *UserMapper) GetWhereOrder(filter *ListUserFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	if filter.Mobile != "" {
		whereOrder = append(whereOrder, *userMapper.WhereEq(filter.Mobile, "mobile"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *userMapper.WhereEq(filter.Id, "id"))
	}
	if filter.QuestionId != "" {
		whereOrder = append(whereOrder, *userMapper.WhereEq(filter.QuestionId, "question_id"))
	}
	if filter.QuestionName != "" {
		whereOrder = append(whereOrder, *userMapper.WhereLike(filter.QuestionName, "question_name"))
	}
	if filter.ShotId != "" {
		whereOrder = append(whereOrder, *userMapper.WhereLike(filter.ShotId, "shot_id"))
	}
	if filter.GroupId != "" {
		whereOrder = append(whereOrder, *userMapper.WhereLike(filter.GroupId, "group_id"))
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
func (userMapper *UserMapper) GetByCondition(filter *ListUserFilter, field []string) (list []*models.User, total uint64, err error) {
	whereOrder := userMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = userMapper.GetDb().GetPage(&list, &models.User{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = userMapper.GetDb().Scan(&list, &models.User{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (userMapper *UserMapper) FirstByCondition(filter *ListUserFilter, field []string) (out models.User, err error) {
	whereOrder := userMapper.GetWhereOrder(filter)
	err = userMapper.GetDb().Find(&out, &models.User{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (userMapper *UserMapper) Paginate(page uint64, pageSize uint64) (list []*models.User, total uint64, err error) {
	var field []string
	return userMapper.GetByCondition(&ListUserFilter{Page: page, Limit: pageSize}, field)
}
