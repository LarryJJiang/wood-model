package service

import (
	"github.com/jinzhu/gorm"
	"wood-client-api/mapper"
	"wood-client-api/models"
)

type stockRecordService struct {
	BaseService
}

var StockRecordService stockRecordService

func init() {
	StockRecordService = stockRecordService{}
}

func (u stockRecordService) GetMapper() *mapper.StockRecordMapper {
	return new(mapper.StockRecordMapper).GetDb()
}

// Scan 获取列表
func (u stockRecordService) Scan(filter *mapper.ListStockRecordFilter) (list []*models.StockRecord, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u stockRecordService) First(filter mapper.ListStockRecordFilter) (out models.StockRecord, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u stockRecordService) Paginate(filter *mapper.ListStockRecordFilter) (list []*models.StockRecord, total uint64, err error) {

	return u.GetMapper().Paginate(filter)
}

// GetByDate
func (u stockRecordService) GetByDate(date int64) (stockRecord models.StockRecord, err error) {
	return u.First(mapper.ListStockRecordFilter{Date: date})
}

// GetByCrewId
func (u stockRecordService) GetByCrewId(crewId int) (stockRecord models.StockRecord, err error) {
	return u.First(mapper.ListStockRecordFilter{CrewId: crewId})
}

// GetById
func (u stockRecordService) GetById(id int) (stockRecord models.StockRecord, err error) {
	return u.First(mapper.ListStockRecordFilter{Id: id})
}

// GetByStockId
func (u stockRecordService) GetByStockId(stockId int) (stockRecordList []*models.StockRecord, err error) {
	var field = []string{"id", "stock_id", "crew_id", "date", "grade", "code", "length", "amount", "status", "create_time"}
	stockRecordList, _, err = u.GetMapper().GetByCondition(&mapper.ListStockRecordFilter{StockId: stockId}, field)
	return
}

func (u stockRecordService) Create(tx *gorm.DB, stockRecord *models.StockRecord) (int, error) {
	insertData := &models.StockRecord{
		StockId: stockRecord.StockId,
		CrewId:  stockRecord.CrewId,
		Date:    stockRecord.Date,
		Grade:   stockRecord.Grade,
		Code:    stockRecord.Code,
		Length:  stockRecord.Length,
		Amount:  stockRecord.Amount,
		Status:  mapper.Active_Status,
	}
	err := tx.Create(insertData).Error
	if err != nil {
		return 0, err
	}
	return insertData.ID, nil
}
