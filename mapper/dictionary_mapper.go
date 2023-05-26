package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type DictionaryMapper struct {
	BaseMapper
}

var DictionaryModel models.Dictionary

func (dictionaryMapper *DictionaryMapper) GetDb() *DictionaryMapper {
	dictionaryMapper.Db = dictionaryMapper.GetModel(dictionaryMapper.GetTable())
	return dictionaryMapper
}

func (dictionaryMapper *DictionaryMapper) GetTableModel() *models.Dictionary {
	return new(models.Dictionary)
}

func (dictionaryMapper *DictionaryMapper) GetTable() string {
	model := dictionaryMapper.GetTableModel()
	dictionaryMapper.Table = model.TableName()
	dictionaryMapper.ALis = ""
	dictionaryMapper.Field = model.GetFieldSlice()
	return dictionaryMapper.Table
}

type listDictionaryFilter struct {
	Page     uint64
	Limit    uint64
	Id       int
	Key      string
	Name     string
	DeleteAt int
}

type ListDictionaryFilter listDictionaryFilter

func (dictionaryMapper *DictionaryMapper) GetWhereOrder(filter *ListDictionaryFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *dictionaryMapper.WhereEq(filter.DeleteAt, dictionaryMapper.GetTable()+".delete_at"))
	if filter.Name != "" {
		whereOrder = append(whereOrder, *dictionaryMapper.WhereEq(filter.Name, "name"))
	}
	if filter.Key != "" {
		whereOrder = append(whereOrder, *dictionaryMapper.WhereEq(filter.Key, "key"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *dictionaryMapper.WhereEq(filter.Id, "id"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (dictionaryMapper *DictionaryMapper) GetByCondition(filter *ListDictionaryFilter, field []string) (list []*models.Dictionary, total uint64, err error) {
	whereOrder := dictionaryMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = dictionaryMapper.GetDb().GetPage(&list, &models.Dictionary{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = dictionaryMapper.GetDb().Scan(&list, &models.Dictionary{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (dictionaryMapper *DictionaryMapper) FirstByCondition(filter *ListDictionaryFilter, field []string) (out models.Dictionary, err error) {
	whereOrder := dictionaryMapper.GetWhereOrder(filter)
	err = dictionaryMapper.GetDb().Find(&out, &models.Dictionary{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

// GetByName 获取多条数据
func (dictionaryMapper *DictionaryMapper) GetByName(filter *ListDictionaryFilter, field []string) (out models.Dictionary, err error) {
	whereOrder := dictionaryMapper.GetWhereOrder(filter)
	err = dictionaryMapper.GetDb().Find(&out, &models.Dictionary{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}
