package valid

import (
	"github.com/gin-gonic/gin"
)

type CrewCarrierRelationValidate struct {
	CrewId    int `alias:"Crew Id" valid:"Required; " form:"crew_id" json:"crew_id"`          // 砍伐队ID
	CarrierId int `alias:"Carrier Id" valid:"Required; " form:"carrier_id" json:"carrier_id"` // 默认目的地id
}

// Valid 创建砍伐队校验
func (a *CrewCarrierRelationValidate) Valid(ctx *gin.Context) (err error) {
	err = a.Check(ctx)
	if err != nil {
		return err
	}
	return Validate(a)
}

func (a *CrewCarrierRelationValidate) Check(ctx *gin.Context) error {
	return nil
}
