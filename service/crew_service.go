package service

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"woods/mapper"
	"woods/models"
	"woods/pkg/util"
	valid "woods/validation"
)

type crewService struct {
	BaseService
}

var CrewService crewService

func init() {
	CrewService = crewService{}
}

func (u crewService) GetMapper() *mapper.CrewMapper {
	return new(mapper.CrewMapper).GetDb()
}

// Scan 获取列表
func (u crewService) Scan(filter *mapper.ListCrewFilter) (list []*models.Crew, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u crewService) First(filter mapper.ListCrewFilter) (out models.Crew, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u crewService) Paginate(filter *mapper.ListCrewFilter) (list []*models.Crew, total uint64, err error) {
	return u.GetMapper().Paginate(filter)
}

// GetByCode
func (u crewService) GetByCode(code string) (crew models.Crew, err error) {
	return u.First(mapper.ListCrewFilter{Code: code, Status: mapper.Active_Status})
}

// GetByUserId
func (u crewService) GetByUserId(userId int) (crew models.Crew, err error) {
	return u.First(mapper.ListCrewFilter{UserId: userId, Status: mapper.Active_Status})
}

// GetAllByUserId
func (u crewService) GetAllByUserId(userId int) (*models.Crew, error) {
	crew, err := u.GetMapper().GetAllByUserId(userId)
	if u.IsError(err) {
		return nil, err
	}
	if u.IsNotFound(err) {
		return nil, errors.New("Crew is Not Found")
	}
	customer, err := CustomerService.GetById(crew.CustomerId)
	if u.IsError(err) {
		return nil, err
	}
	if u.IsNotFound(err) {
		return nil, errors.New("Customer is Not Found")
	}
	crew.Customer = &customer
	return &crew, nil
}

// GetById
func (u crewService) GetById(id int) (crew models.Crew, err error) {
	return u.First(mapper.ListCrewFilter{Id: id, Status: mapper.Active_Status})
}

// GetByCustomerId
func (u crewService) GetByCustomerId(customerId int) (crew models.Crew, err error) {
	return u.First(mapper.ListCrewFilter{CustomerId: customerId, Status: mapper.Active_Status})
}

//删除砍伐队账号
func (u crewService) DeleteCrew(id int) error {
	crew, err := u.GetById(id)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Crew")
	}
	return u.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&models.Crew{}).Where(map[string]interface{}{"id": id}).Update(map[string]interface{}{
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

//砍伐队绑定运输公司
func (u crewService) BindCarrier(postData *valid.CrewCarrierRelationValidate) error {
	// 检查林场信息是否存在
	_, err := u.GetById(postData.CrewId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return u.NotFound("Carrier")
	}
	// 检查运输公司信息是否存在
	_, err = CarrierService.GetById(postData.CarrierId)
	if u.IsError(err) {
		return err
	}
	_, err = CrewCarrierRelationService.GetByCrewIdAndCarrierId(postData.CrewId, postData.CarrierId)
	if !u.IsNotFound(err) {
		return u.RecordExists("Bind Relation")
	}
	return CrewCarrierRelationService.GetMapper().Insert(&models.CrewCarrierRelation{CrewId: postData.CrewId, CarrierId: postData.CarrierId, Status: mapper.Active_Status})
}

// GetByCrewId 通过store_id和goods_id获取商城商品信息
func (u crewService) GetByCrewId(filter *mapper.ListCrewFilter) (crew models.Crew, err error) {
	//where := map[string]interface{}{"id": crewId, "delete_at": 0}
	var field []string
	condition := StockService.GetMapper().GetWhereOrder(&mapper.ListStockFilter{})
	return u.GetMapper().Preload("StockList", condition).FirstByCondition(filter, field)
}

// GetCrewList 获取砍伐队下拉列表
func (u crewService) GetCrewList(filter *mapper.ListCrewFilter) (list []*models.Crew, err error) {
	return u.GetMapper().GetCrewList(filter)
}
