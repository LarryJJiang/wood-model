package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type StockMapper struct {
	BaseMapper
}

var StockModel models.Stock

func (stockMapper *StockMapper) GetDb() *StockMapper {
	stockMapper.Db = stockMapper.GetModel(stockMapper.GetTable())
	return stockMapper
}

func (stockMapper *StockMapper) GetTableModel() *models.Stock {
	return new(models.Stock)
}

func (stockMapper *StockMapper) GetTable() string {
	model := stockMapper.GetTableModel()
	stockMapper.Table = model.TableName()
	stockMapper.ALis = ""
	stockMapper.Field = model.GetFieldSlice()
	return stockMapper.Table
}

type listStockFilter struct {
	Page       uint64
	Limit      uint64
	Id         int
	UserId     int
	CrewId     int
	CustomerId int
	Date       uint64
	DeleteAt   int
}

type ListStockFilter listStockFilter

func (stockMapper *StockMapper) GetWhereOrder(filter *ListStockFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *stockMapper.WhereEq(filter.DeleteAt, stockMapper.Table+".delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *stockMapper.WhereEq(filter.Id, "id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *stockMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.CrewId != 0 {
		whereOrder = append(whereOrder, *stockMapper.WhereEq(filter.CrewId, "crew_id"))
	}
	if filter.Date != 0 {
		whereOrder = append(whereOrder, *stockMapper.WhereEq(int64(filter.Date), "date"))
	}
	if filter.CustomerId != 0 {
		whereOrder = append(whereOrder, *stockMapper.WhereEq(filter.CustomerId, "customer_id"))
	}

	return whereOrder
}

// GetByCondition 获取列表
func (stockMapper *StockMapper) GetByCondition(filter *ListStockFilter, field []string) (list []*models.Stock, total uint64, err error) {
	whereOrder := stockMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = stockMapper.GetDb().GetPage(&list, &models.Stock{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = stockMapper.GetDb().Scan(&list, &models.Stock{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (stockMapper *StockMapper) FirstByCondition(filter *ListStockFilter, field []string) (out models.Stock, err error) {
	whereOrder := stockMapper.GetWhereOrder(filter)
	err = stockMapper.GetDb().Find(&out, &models.Stock{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (stockMapper *StockMapper) Paginate(filter *ListStockFilter) (list []*models.Stock, total uint64, err error) {
	var field []string
	return stockMapper.GetByCondition(filter, field)
}

func (stockMapper *StockMapper) Preload(column string, conditions ...interface{}) *StockMapper {
	stockMapper.Db = stockMapper.Preloads(column, conditions...).Db
	return stockMapper
}
