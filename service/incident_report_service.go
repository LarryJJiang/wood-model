package service

import (
	"woods/mapper"
	"woods/models"
	valid "woods/validation"
)

type incidentService struct {
	BaseService
}

var IncidentReportService incidentService

func init() {
	IncidentReportService = incidentService{}
}

func (u incidentService) GetMapper() *mapper.IncidentReportMapper {
	return new(mapper.IncidentReportMapper).GetDb()
}

// Scan 获取列表
func (u incidentService) Scan(filter mapper.ListIncidentReportFilter) (list []*models.IncidentReport, total uint64, err error) {
	var field []string
	list, total, err = u.GetMapper().GetByCondition(&filter, field)
	return
}

// First 获取一条数据
func (u incidentService) First(filter mapper.ListIncidentReportFilter) (out models.IncidentReport, err error) {
	var field []string
	out, err = u.GetMapper().FirstByCondition(&filter, field)
	return
}

func (u incidentService) Paginate(filter *mapper.ListIncidentReportFilter) (list []*models.IncidentReport, total uint64, err error) {
	return u.GetMapper().Paginate(filter)
}

// GetByUserId
func (u incidentService) GetByUserId(userId int) (list []*models.IncidentReport, err error) {
	list, _, err = u.Scan(mapper.ListIncidentReportFilter{UserId: userId})
	return
}

// Create
func (u incidentService) Create(userId int, postData *valid.IncidentReportValidate) (err error) {
	vehicle, err := VehicleService.GetByUserId(userId)
	insertData := &models.IncidentReport{
		UserId:       userId,
		VehicleId:    vehicle.ID,
		IncidentType: postData.IncidentType,
		Location:     postData.Location,
		Forest:       postData.Forest,
		Remark:       postData.Remark,
		Images:       postData.Images,
	}
	return u.GetMapper().Insert(insertData)
}
