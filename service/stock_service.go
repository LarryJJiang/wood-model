package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strconv"
	"time"
	"woods/mapper"
	"woods/models"
	"woods/pkg/util"
	valid "woods/validation"
)

type stockService struct {
	BaseService
}

var StockService stockService

func init() {
	StockService = stockService{}
}

func (u stockService) GetMapper() *mapper.StockMapper {
	return new(mapper.StockMapper).GetDb()
}

// Scan 获取列表
func (u stockService) Scan(filter *mapper.ListStockFilter) (list []*models.Stock, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u stockService) First(filter mapper.ListStockFilter) (out models.Stock, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u stockService) Paginate(filter *mapper.ListStockFilter) (list []*models.Stock, total uint64, err error) {
	return u.GetMapper().Paginate(filter)
}

// GetByCustomerId
func (u stockService) GetByCustomerId(customerId int) (stock models.Stock, err error) {
	return u.First(mapper.ListStockFilter{CustomerId: customerId})
}

// GetByDate
func (u stockService) GetByDate(date uint64) (stockList []*models.Stock, err error) {
	field := []string{"id", "user_id", "crew_id", "customer_id", "total_amount"}
	stockList, _, err = u.GetMapper().GetByCondition(&mapper.ListStockFilter{Date: date}, field)
	return
}

// GetByUserId
func (u stockService) GetByUserId(userId int) (stock models.Stock, err error) {
	return u.First(mapper.ListStockFilter{UserId: userId})
}

// GetById
func (u stockService) GetById(id int) (stock models.Stock, err error) {
	return u.First(mapper.ListStockFilter{Id: id})
}

// GetByCrewId
//func (u stockService) GetByCrewId(crewId int) (stock models.Stock, err error) {
//	return u.First(mapper.ListStockFilter{CrewId: crewId})
//}

// CreateStock 创建系统用户
func (u stockService) CreateStock(userId int, postData *valid.StockValidate) error {
	// 检查货车表中Code是否与别的货车同名
	dateTime := util.Time2Date(time.Now())
	date, _ := util.Date2Time(dateTime)
	stock, err := u.GetMapper().FirstByCondition(&mapper.ListStockFilter{CrewId: postData.CrewId, Date: uint64(date.Unix())}, []string{"id"})
	if u.IsError(err) {
		return err
	}
	if stock.ID > 0 {
		length, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", postData.Length), 64)
		fmt.Println("长度：", postData.Length, length)
		result, _, err := StockRecordService.Scan(&mapper.ListStockRecordFilter{StockId: stock.ID, CrewId: postData.CrewId, Grade: postData.Grade, Code: postData.Code, Length: length})
		if u.IsError(err) {
			return err
		}
		if len(result) > 0 {
			return errors.New("Stock Record Exists")
		}
	}
	crew, err := CrewService.GetById(postData.CrewId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Crew Not Exists")
	}

	return u.Transaction(func(tx *gorm.DB) error {
		stockId := stock.ID
		if stock.ID == 0 {
			stockData := &models.Stock{
				UserId:      crew.UserId,
				CrewId:      crew.ID,
				CustomerId:  crew.CustomerId,
				Date:        int(date.Unix()),
				TotalAmount: 0.00,
			}
			err = tx.Create(stockData).Error
			if !u.IsErrorNil(err) {
				return err
			}
			stockId = stockData.ID
		}
		stockRecord := &models.StockRecord{
			StockId: stockId,
			CrewId:  crew.ID,
			Date:    int(date.Unix()),
			Grade:   postData.Grade,
			Code:    postData.Code,
			Length:  postData.Length,
			Amount:  0.00,
			Status:  mapper.Active_Status,
		}
		err = tx.Create(stockRecord).Error
		if !u.IsErrorNil(err) {
			return err
		}
		return err
	})
}

// UpdateStock 更新货车司机用户
func (u stockService) UpdateStock(stockRecordId int, postData *valid.StockUpdateValidate) error {
	// 检查库存记录是否存在
	stockRecord, err := StockRecordService.GetById(stockRecordId)
	if u.IsError(err) {
		return err
	}
	// 检查库存是否存在
	stock, err := StockService.GetById(stockRecord.StockId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Stock Not Exists")
	}
	date := util.GetDateTimestamp(time.Now())
	if date != int64(stock.Date) {
		return errors.New("Can not edit history stock")
	}
	// 检查砍伐队是否存在
	_, err = CrewService.GetById(stockRecord.CrewId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Crew Not Exists")
	}
	updateData := map[string]interface{}{
		"amount":      postData.Amount,
		"update_time": util.NowMilli(),
	}
	return u.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&models.StockRecord{}).Where(map[string]interface{}{"id": stockRecordId, "delete_at": 0}).Update(updateData).Error
		if !u.IsErrorNil(err) {
			return err
		}
		totalStock := util.AddFloat64(util.SubFloat64(stock.TotalAmount, stockRecord.Amount), postData.Amount)
		return u.UpdateStockAmount(tx, stock.ID, totalStock)
	})
}

//删除
func (u stockService) DeleteStock(id int) error {
	stockRecord, err := StockRecordService.GetById(id)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Stock")
	}
	stock, err := StockService.GetById(id)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Stock")
	}
	date := util.GetDateTimestamp(time.Now())
	if date != int64(stock.Date) {
		return errors.New("Can not edit history stock")
	}
	return u.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&models.StockRecord{}).Where(map[string]interface{}{"id": id}).Update(map[string]interface{}{"delete_at": util.Now().Year(), "update_time": util.Now().UnixMilli()}).Error
		if !u.IsErrorNil(err) {
			return err
		}
		// 如果数量大于0，则需要减总库存
		if stockRecord.Amount > 0 {
			totalAmount := util.SubFloat64(stock.TotalAmount, stockRecord.Amount)
			return u.UpdateStockAmount(tx, stock.ID, totalAmount)
		}
		return err
	})
}

// GetByCrewId 通过store_id和goods_id获取商城商品信息
func (u stockService) GetByCrewId(filter *mapper.ListStockFilter) (*models.Crew, error) {
	fmt.Println("uint64转int64：", int64(filter.Date/1000))
	date := util.GetDateTimestamp(util.Unix2Time(int64(filter.Date / 1000)))
	where := map[string]interface{}{"id": filter.CrewId, "delete_at": 0}

	var crew models.Crew
	crewModel := new(mapper.CrewMapper).GetTableModel()
	stockModel := new(mapper.StockMapper).GetTableModel()
	var field map[string][]string
	field = map[string][]string{
		crewModel.TableName():  {"*"},
		stockModel.TableName(): {"id as stock_id", "total_amount"},
	}
	crewMapper := CrewService.GetMapper()
	crewMapper.WhereCondition = where
	_, err := crewMapper.LeftJoin(stockModel, crewModel.TableName()+".id", "=", stockModel.TableName()+".crew_id").
		Preloads("RecordList", "date = ? and delete_at = 0", date).First(&crew, u.GetMapper().GetField(field))
	if u.IsError(err) {
		return nil, err
	}
	if u.IsNotFound(err) {
		return nil, errors.New("Crew Not Exists")
	}
	return &crew, nil
}

func (u stockService) Create(tx *gorm.DB, stock *models.Stock) (int, error) {
	insertStock := &models.Stock{
		UserId:      stock.UserId,
		CrewId:      stock.CrewId,
		CustomerId:  stock.CustomerId,
		TotalAmount: stock.TotalAmount,
		Date:        stock.Date,
	}
	fmt.Println("库存数据：", insertStock)
	err := tx.Create(insertStock).Error
	if u.IsError(err) {
		return 0, err
	}
	return insertStock.ID, nil
}

// 更新库存记录库存
func (u stockService) UpdateStockAmount(tx *gorm.DB, stockId int, totalStock float64) error {
	stockUpdate := mapper.Map{
		"total_amount": totalStock,
		"update_time":  util.NowMilli(),
	}
	return tx.Model(&models.Stock{}).Where(map[string]interface{}{"id": stockId, "delete_at": 0}).Update(stockUpdate).Error
}
