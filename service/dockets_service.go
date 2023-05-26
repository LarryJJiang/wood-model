package service

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"wood-client-api/mapper"
	"wood-client-api/models"
	"wood-client-api/pkg/e"
	"wood-client-api/pkg/google_cloud"
	"wood-client-api/pkg/gredis"
	"wood-client-api/pkg/logging"
	"wood-client-api/pkg/util"
	"wood-client-api/pkg/util/convert"
	valid "wood-client-api/validation"
)

type docketsService struct {
	BaseService
}

var DocketsService docketsService

func init() {
	DocketsService = docketsService{}
}

func (ds docketsService) GetMapper() *mapper.DocketsMapper {
	return new(mapper.DocketsMapper).GetDb()
}

// Scan 获取列表
func (ds docketsService) Scan(filter *mapper.ListDocketsFilter) (list []*models.Dockets, total uint64, err error) {
	var field []string
	list, total, err = ds.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (ds docketsService) First(filter mapper.ListDocketsFilter) (out models.Dockets, err error) {
	var field []string
	out, err = ds.GetMapper().FirstByCondition(&filter, field)
	return
}

func (ds docketsService) Paginate(filter *mapper.ListDocketsFilter) (list []*models.Dockets, total uint64, err error) {
	return ds.GetMapper().Paginate(filter)
}

// GetById
func (ds docketsService) GetById(id int) (dockets models.Dockets, err error) {
	return ds.First(mapper.ListDocketsFilter{Id: id})
}

// CreateDocket
func (ds docketsService) CreateDocket(userId, taskId int, postData *valid.DocketCreateValidate) (err error) {
	crew, err := CrewService.GetByUserId(userId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Crew")
	}
	task, err := TaskService.GetById(taskId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Task")
	}
	stockRecord, err := StockRecordService.GetById(postData.StockRecordId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Stock Record")
	}
	destination, err := DestinationService.GetById(postData.DestinationId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Destination")
	}
	vehicle, err := VehicleService.GetById(task.VehicleId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Vehicle")
	}
	date := util.GetDateTimestamp(util.Now()) * 1000
	count, err := ds.GetMapper().Count(&models.Dockets{TaskId: taskId, CrewId: crew.ID, Date: date, StockRecordId: postData.StockRecordId, DestinationId: postData.DestinationId, DeleteAt: 0})
	if ds.IsError(err) {
		return err
	}
	if count > 0 {
		return errors.New("Docket Exists")
	}
	docketNumber := ds.GetDocketNumber("")
	if ds.IsError(err) {
		return err
	}
	return ds.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&models.Dockets{
			TaskId:        taskId,
			CrewId:        crew.ID,
			Logger:        crew.Code,
			VehicleId:     task.VehicleId,
			TruckCode:     task.VehicleCode,
			CarrierId:     vehicle.CarrierId,
			Carrier:       vehicle.CarrierName,
			DocketNumber:  convert.ToString(docketNumber),
			OperationType: postData.OperationType,
			Date:          date,
			Customer:      postData.Customer,
			Forest:        crew.Forest,
			Compartment:   crew.Compartment,
			Setting:       crew.Setting,
			StockRecordId: postData.StockRecordId,
			Grade:         stockRecord.Grade,
			Species:       stockRecord.Species,
			Code:          stockRecord.Code,
			Length:        stockRecord.Length,
			DestinationId: postData.DestinationId,
			Destination:   destination.Code,
			TareWeight:    util.MulFloat64(postData.TareWeight, 1000),
			NetWeight:     util.MulFloat64(postData.NetWeight, 1000),
			GrossWeight:   util.MulFloat64(postData.GrossWeight, 1000),
			Rate:          crew.Rate,
		}).Error
		return err
	})
}

// UpdateDocket
func (ds docketsService) UpdateDocket(userId, docketId int, postData *valid.DocketCreateValidate) (err error) {
	crew, err := CrewService.GetByUserId(userId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Crew")
	}
	docket, err := ds.GetById(docketId)
	if ds.IsError(err) {
		return err
	}
	if ds.IsNotFound(err) {
		return ds.NotFound("Docket")
	}
	if docket.CrewId != crew.ID {
		return ds.NotFound("Docket")
	}
	updateData := &models.Dockets{
		OperationType: postData.OperationType,
		Customer:      postData.Customer,
		StockRecordId: postData.StockRecordId,
		DestinationId: postData.DestinationId,
		TareWeight:    util.MulFloat64(postData.TareWeight, 1000),
		NetWeight:     util.MulFloat64(postData.NetWeight, 1000),
		GrossWeight:   util.MulFloat64(postData.GrossWeight, 1000),
	}
	if postData.StockRecordId != docket.StockRecordId {
		stockRecord, err := StockRecordService.GetById(postData.StockRecordId)
		if ds.IsError(err) {
			return err
		}
		if ds.IsNotFound(err) {
			return ds.NotFound("Stock Record")
		}
		updateData.StockRecordId = postData.StockRecordId
		updateData.Grade = stockRecord.Grade
		updateData.Species = stockRecord.Species
		updateData.Code = stockRecord.Code
		updateData.Length = stockRecord.Length
	}

	var destination models.Destination
	if postData.DestinationId != docket.DestinationId {
		destination, err = DestinationService.GetById(postData.DestinationId)
		if ds.IsError(err) {
			return err
		}
		if ds.IsNotFound(err) {
			return ds.NotFound("Destination")
		}
		updateData.Destination = destination.Code
	}
	return ds.GetMapper().UpdateById(docketId, updateData)
}

// GetDocketNumber 生成docket number
func (ds docketsService) GetDocketNumber(prefix string) (docketNumber string) {
	var err error
	var number int64
	var step = 3
	if !gredis.Exists(e.DocketNumberGenerator) { // 该key不存在，是使用incr初始化
		step = e.DocketNumberMin
	}
	number, err = gredis.IncrBy(e.DocketNumberGenerator, step, 0)
	if err != nil {
		logging.Error("Generate:", err.Error())
	}
	return fmt.Sprintf("%v%v", prefix, convert.ToString(number))
}

// GetDocketNumber 生成docket number
func (ds docketsService) UpdateDocketData(docketId int, file *google_cloud.FileInfo) error {
	fmt.Println("单据ID", docketId)
	result, err := google_cloud.IdentifyFile(file)
	if err != nil {
		return err
	}
	updateData, err := json.Marshal(result)
	if err != nil {
		return err
	}
	go ds.GetMapper().UpdateById(docketId, &models.Dockets{Detail: string(updateData)})
	return nil
}
