package service

import (
	"woods/mapper"
	"woods/models"
	"woods/pkg/util/convert"
)

type taskService struct {
	BaseService
}

var TaskService taskService

func init() {
	TaskService = taskService{}
}

func (u taskService) GetMapper() *mapper.TaskMapper {
	return new(mapper.TaskMapper).GetDb()
}

// Scan 获取列表
func (u taskService) Scan(filter *mapper.ListTaskFilter) (list []*models.Task, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(filter, field)
	return
}

// First 获取一条数据
func (u taskService) First(filter mapper.ListTaskFilter) (out models.Task, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u taskService) Paginate(filter *mapper.ListTaskFilter) (list []*models.Task, total uint64, err error) {
	list, _, err = u.GetMapper().Paginate(filter)
	if len(list) > 0 {
		vehicleIds := make([]int, 0)
		for _, task := range list {
			if len(task.DocketList) > 0 {
				for _, docket := range task.DocketList {
					vehicleIds = append(vehicleIds, docket.VehicleId)
				}
			}
		}
		vehicleIds = convert.UnrepeatedSlice(vehicleIds)
		vehicleList, _ := VehicleService.GetByIds(vehicleIds)
		if len(vehicleList) > 0 {
			vehicleListMap := make(map[int]*models.Vehicle, len(vehicleList))
			for _, vehicle := range vehicleList {
				vehicleListMap[vehicle.ID] = vehicle
			}
			for _, task := range list {
				if len(task.DocketList) > 0 {
					for _, docket := range task.DocketList {
						docket.Vehicle = vehicleListMap[docket.VehicleId]
					}
				}
			}
		}
	}
	return
}

// GetByUserId
func (u taskService) GetByUserId(userId int) (task models.Task, err error) {
	return u.First(mapper.ListTaskFilter{UserId: userId, Status: mapper.Active_Status})
}

// GetById
func (u taskService) GetById(id int) (task models.Task, err error) {
	return u.First(mapper.ListTaskFilter{Id: id})
}

// GetByCustomerId
func (u taskService) GetByCarrierId(carrierId int) (task models.Task, err error) {
	return u.First(mapper.ListTaskFilter{CarrierId: carrierId, Status: mapper.Active_Status})
}

// CreateSystemUser 创建系统用户
func (u taskService) CreateTask() error {
	return nil
}

// UpdateTask 更新货车司机用户
func (u taskService) UpdateTask() error {
	return nil
}

//删除货车账号
func (u taskService) DeleteTask(id int) error {
	return nil
}
