package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
	"wood-client-api/pkg/util"
)

type DestinationMapper struct {
	BaseMapper
}

var DestinationModel models.Destination

func (destinationMapper *DestinationMapper) GetDb() *DestinationMapper {
	destinationMapper.Db = destinationMapper.GetModel(destinationMapper.GetTable())
	return destinationMapper
}

func (destinationMapper *DestinationMapper) GetTableModel() *models.Destination {
	return new(models.Destination)
}

func (destinationMapper *DestinationMapper) GetTable() string {
	model := destinationMapper.GetTableModel()
	destinationMapper.Table = model.TableName()
	destinationMapper.ALis = ""
	destinationMapper.Field = model.GetFieldSlice()
	return destinationMapper.Table
}

type listDestinationFilter struct {
	Page               uint64
	Limit              uint64
	Id                 int
	IdNotSelf          int
	IdIn               string
	Name               string
	NameLike           string
	Code               string
	Type               int
	WeighBridgeSupport int
	Status             int
	DeleteAt           int
}

type ListDestinationFilter listDestinationFilter

func (destinationMapper *DestinationMapper) GetWhereOrder(filter *ListDestinationFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.DeleteAt, "delete_at"))
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.Id, "id"))
	}
	if filter.IdNotSelf != 0 {
		whereOrder = append(whereOrder, *destinationMapper.WhereNEq(filter.IdNotSelf, "id"))
	}
	if filter.IdIn != "" {
		whereOrder = append(whereOrder, *destinationMapper.WhereIn(filter.IdIn, "id"))
	}
	if filter.Name != "" {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.Name, "name"))
	}
	if filter.NameLike != "" {
		whereOrder = append(whereOrder, *destinationMapper.WhereLike(filter.NameLike, "name"))
	}
	if filter.Code != "" {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.Code, "code"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.Status, "status"))
	}
	if filter.Type != 0 {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.Type, "type"))
	}
	if filter.WeighBridgeSupport != 0 {
		whereOrder = append(whereOrder, *destinationMapper.WhereEq(filter.WeighBridgeSupport, "weight_bridge_support"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (destinationMapper *DestinationMapper) GetByCondition(filter *ListDestinationFilter, field []string) (list []*models.Destination, total uint64, err error) {
	whereOrder := destinationMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = destinationMapper.GetDb().GetPage(&list, &models.Destination{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = destinationMapper.GetDb().Scan(&list, &models.Destination{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (destinationMapper *DestinationMapper) FirstByCondition(filter *ListDestinationFilter, field []string) (out models.Destination, err error) {
	whereOrder := destinationMapper.GetWhereOrder(filter)
	err = destinationMapper.GetDb().Find(&out, &models.Destination{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}

func (destinationMapper *DestinationMapper) Paginate(filter *ListDestinationFilter) (list []*models.Destination, total uint64, err error) {
	var field []string
	return destinationMapper.GetByCondition(filter, field)
}

// Create
func (destinationMapper *DestinationMapper) Create(data *models.Destination) error {
	return destinationMapper.Insert(&models.Destination{
		Name:               data.Name,
		Code:               data.Code,
		Type:               data.Type,
		WeighBridgeSupport: data.WeighBridgeSupport,
		Status:             data.Status,
	})
}

// Save
func (destinationMapper *DestinationMapper) Save(id int, data *models.Destination) error {
	return destinationMapper.UpdateById(id, map[string]interface{}{
		"name":                  data.Name,
		"code":                  data.Code,
		"type":                  data.Type,
		"weight_bridge_support": data.WeighBridgeSupport,
		"status":                data.Status,
		"update_time":           util.Now().UnixMilli(),
	})
}

// DeleteDestination
func (destinationMapper *DestinationMapper) DeleteDestination(id int) error {
	return destinationMapper.SoftDeleteById(id)
}
