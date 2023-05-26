package mapper

import (
	"woods/models"
	bizcode "woods/pkg/bizerror"
)

type StockRecordMapper struct {
	BaseMapper
}

var StockRecordModel models.StockRecord

func (stockRecordMapper *StockRecordMapper) GetDb() *StockRecordMapper {
	stockRecordMapper.Db = stockRecordMapper.GetModel(stockRecordMapper.GetTable())
	return stockRecordMapper
}

func (stockRecordMapper *StockRecordMapper) GetTableModel() *models.StockRecord {
	return new(models.StockRecord)
}

func (stockRecordMapper *StockRecordMapper) GetTable() string {
	model := stockRecordMapper.GetTableModel()
	stockRecordMapper.Table = model.TableName()
	stockRecordMapper.ALis = ""
	stockRecordMapper.Field = model.GetFieldSlice()
	return stockRecordMapper.Table
}

type listStockRecordFilter struct {
	Page     uint64
	Limit    uint64
	Id       int
	StockId  int
	CrewId   int
	UserId   int
	Code     string
	Date     int64
	Grade    string
	Length   float64
	DeleteAt int
}

type ListStockRecordFilter listStockRecordFilter

func (stockRecordMapper *StockRecordMapper) GetWhereOrder(filter *ListStockRecordFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.DeleteAt, stockRecordMapper.Table+".delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.Id, "id"))
	}
	if filter.StockId != 0 {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.StockId, "stock_id"))
	}
	if filter.CrewId != 0 {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.CrewId, "crew_id"))
	}
	if filter.Date != 0 {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.Date, "date"))
	}
	if filter.Grade != "" {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.Grade, "grade"))
	}
	if filter.Code != "" {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.Code, "code"))
	}
	if filter.Length != 0 {
		whereOrder = append(whereOrder, *stockRecordMapper.WhereEq(filter.Length, "length"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (stockRecordMapper *StockRecordMapper) GetByCondition(filter *ListStockRecordFilter, field []string) (list []*models.StockRecord, total uint64, err error) {
	whereOrder := stockRecordMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = stockRecordMapper.GetDb().GetPage(&list, &models.StockRecord{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = stockRecordMapper.GetDb().Scan(&list, &models.StockRecord{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (stockRecordMapper *StockRecordMapper) FirstByCondition(filter *ListStockRecordFilter, field []string) (out models.StockRecord, err error) {
	whereOrder := stockRecordMapper.GetWhereOrder(filter)
	err = stockRecordMapper.GetDb().Find(&out, &models.StockRecord{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (stockRecordMapper *StockRecordMapper) Paginate(filter *ListStockRecordFilter) (list []*models.StockRecord, total uint64, err error) {
	var field []string
	return stockRecordMapper.GetByCondition(filter, field)
}
