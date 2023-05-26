package service

import (
	"fmt"
	"github.com/pkg/errors"
	"woods/mapper"
	"woods/models"
	"woods/pkg/util/convert"
	valid "woods/validation"
)

type vehicleService struct {
	BaseService
}

var VehicleService vehicleService

func init() {
	VehicleService = vehicleService{}
}

func (u vehicleService) GetMapper() *mapper.VehicleMapper {
	return new(mapper.VehicleMapper).GetDb()
}

// Scan 获取列表
func (u vehicleService) Scan(filter *mapper.ListVehicleFilter) (list []*models.Vehicle, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u vehicleService) First(filter mapper.ListVehicleFilter) (out models.Vehicle, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u vehicleService) Paginate(filter *mapper.ListVehicleFilter) (list []*models.Vehicle, total uint64, err error) {
	return u.GetMapper().Paginate(filter)
}

// GetByCode
func (u vehicleService) GetByCode(code string) (vehicle models.Vehicle, err error) {
	return u.First(mapper.ListVehicleFilter{Code: code, Status: mapper.Active_Status})
}

// GetByRegistration
func (u vehicleService) GetByRegistration(registration string) (vehicle models.Vehicle, err error) {
	return u.First(mapper.ListVehicleFilter{Registration: registration, Status: mapper.Active_Status})
}

// GetByUserId
func (u vehicleService) GetByUserId(userId int) (vehicle models.Vehicle, err error) {
	return u.First(mapper.ListVehicleFilter{UserId: userId, Status: mapper.Active_Status})
}

// GetById
func (u vehicleService) GetById(id int) (vehicle models.Vehicle, err error) {
	return u.First(mapper.ListVehicleFilter{Id: id, Status: mapper.Active_Status})
}

// GetByCustomerId
func (u vehicleService) GetByCarrierId(carrierId int) (vehicle models.Vehicle, err error) {
	return u.First(mapper.ListVehicleFilter{CarrierId: carrierId, Status: mapper.Active_Status})
}

// AcceptTask 司机接受任务
func (u vehicleService) AcceptTask(taskId, userId int) error {
	vehicle, err := u.GetByUserId(userId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Vehicle Not Exists")
	}
	// 任务是否存在
	task, err := TaskService.GetById(taskId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Task Record Not Found")
	}
	// 检查任务状态是否是待接受状态
	if task.Status != mapper.Status_Pending {
		return errors.New("Task Status is Not Pending")
	}
	// 任务与货车是否匹配
	if vehicle.ID != task.VehicleId {
		return errors.New("You've Got Wrong Task")
	}

	return TaskService.GetMapper().UpdateById(taskId, &models.Task{Status: mapper.Status_Accept})
}

// DeclineTask 司机拒绝任务
func (u vehicleService) DeclineTask(taskId, userId int, postData *valid.TaskDeclineValidate) error {
	vehicle, err := u.GetByUserId(userId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Vehicle Not Exists")
	}
	// 任务是否存在
	task, err := TaskService.GetById(taskId)
	if u.IsError(err) {
		return err
	}
	if u.IsNotFound(err) {
		return errors.New("Task Record Not Found")
	}
	if task.Status == mapper.Status_Decline {
		return errors.New("The Task is Declined")
	}
	if task.Status == mapper.Status_Finished {
		return errors.New("The Task is Finished")
	}
	fmt.Println("任务状态：", task.Status)
	// 检查任务状态是否是待接受或已接受状态，这两种状态都可以拒单
	if task.Status != mapper.Status_Pending && task.Status != mapper.Status_Accept {
		return errors.New("Task Cannot Decline")
	}
	// 任务与货车是否匹配
	if vehicle.ID != task.VehicleId {
		return errors.New("You've Got Wrong Task")
	}

	return TaskService.GetMapper().UpdateById(taskId, &models.Task{Status: mapper.Status_Decline, DeclineReason: postData.DeclineReason})
}

// GetByIds
func (u vehicleService) GetByIds(ids []int) (list []*models.Vehicle, err error) {
	list, _, err = u.Scan(&mapper.ListVehicleFilter{IdIn: convert.Join2str(ids, ","), Status: mapper.Active_Status})
	return
}
