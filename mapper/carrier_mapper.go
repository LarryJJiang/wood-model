package mapper

import (
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
)

type CarrierMapper struct {
	BaseMapper
}

var CarrierModel models.Carrier

func (carrierMapper *CarrierMapper) GetDb() *CarrierMapper {
	carrierMapper.Db = carrierMapper.GetModel(carrierMapper.GetTable())
	return carrierMapper
}

func (carrierMapper *CarrierMapper) GetTableModel() *models.Carrier {
	return new(models.Carrier)
}

func (carrierMapper *CarrierMapper) GetTable() string {
	model := carrierMapper.GetTableModel()
	carrierMapper.Table = model.TableName()
	carrierMapper.ALis = ""
	carrierMapper.Field = model.GetFieldSlice()
	return carrierMapper.Table
}

type listCarrierFilter struct {
	Page          uint64
	Limit         uint64
	Id            int
	IdIn          string
	UserId        int
	UserIdNotSelf int
	Name          string
	NameLike      string
	Code          string
	Status        int
	DeleteAt      int
}

type ListCarrierFilter listCarrierFilter

func (carrierMapper *CarrierMapper) GetWhereOrder(filter *ListCarrierFilter) []PageWhereOrder {
	var whereOrder []PageWhereOrder
	whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.DeleteAt, carrierMapper.GetTable()+".delete_at"))
	if filter.Name != "" {
		whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.Name, "name"))
	}
	if filter.NameLike != "" {
		whereOrder = append(whereOrder, *carrierMapper.WhereLike(filter.NameLike, "name"))
	}
	if filter.Id != 0 {
		whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.Id, carrierMapper.GetTable()+".id"))
	}
	if filter.IdIn != "" {
		whereOrder = append(whereOrder, *carrierMapper.WhereIn(filter.IdIn, carrierMapper.GetTable()+".id"))
	}
	if filter.UserId != 0 {
		whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.UserId, "user_id"))
	}
	if filter.UserIdNotSelf != 0 {
		whereOrder = append(whereOrder, *carrierMapper.WhereNEq(filter.UserIdNotSelf, "user_id"))
	}
	if filter.Code != "" {
		whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.Code, "code"))
	}
	if filter.Status != 0 {
		whereOrder = append(whereOrder, *carrierMapper.WhereEq(filter.Status, carrierMapper.GetTable()+".status"))
	}
	return whereOrder
}

// GetByCondition 获取列表
func (carrierMapper *CarrierMapper) GetByCondition(filter *ListCarrierFilter, field []string) (list []*models.Carrier, total uint64, err error) {
	whereOrder := carrierMapper.GetWhereOrder(filter)
	if filter.Page > 0 {
		err = carrierMapper.GetDb().GetPage(&list, &models.Carrier{}, field, filter.Page, filter.Limit, &total, whereOrder...)
	} else {
		err = carrierMapper.GetDb().Scan(&list, &models.Carrier{}, field, 0, whereOrder...)
	}
	return list, total, err
}

// FirstByCondition 获取一条数据
func (carrierMapper *CarrierMapper) FirstByCondition(filter *ListCarrierFilter, field []string) (out models.Carrier, err error) {
	whereOrder := carrierMapper.GetWhereOrder(filter)
	err = carrierMapper.GetDb().Find(&out, &models.Carrier{}, field, 1, whereOrder...)
	bizcode.DbCheck(err)
	return
}
