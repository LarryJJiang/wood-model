package service

import (
	"github.com/astaxie/beego/validation"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"sync"
	"woods/mapper"
	"woods/models"
	bizcode "woods/pkg/bizerror"
	"woods/pkg/e"
	"woods/pkg/util"
)

type BaseServiceInterface interface {
	GetMapper() interface{}
}

type BaseService struct {
	Mutex *sync.Mutex
}

//获取基础mapper
func (base BaseService) GetMapper() *mapper.BaseMapper {
	return new(mapper.BaseMapper)
}

func (base BaseService) Lock() {
	if base.Mutex == nil {
		bizcode.BizCode{}.PanicError(e.ErrorSystemLock, "")
	}
	base.Mutex.Lock()
}

func (base BaseService) Unlock() {
	if base.Mutex == nil {
		bizcode.BizCode{}.PanicError(e.ErrorSystemLock, "")
	}
	base.Mutex.Unlock()
}

func (base BaseService) CheckParam(params interface{}) {
	if params != nil {
		valid := validation.Validation{}
		result, err := valid.Valid(params)
		bizcode.Check(err)
		if !result {
			for _, err := range valid.Errors {
				bizcode.BizCode{}.PanicError(e.InvalidParams, err.Field+" "+err.Message)
			}
		}
	}
}

// Service层的事物处理
// 使用此方法，如需要返回错误信息给客户端，需要在controller层定义 defer con.PanicHandler(http.StatusOK) 来处理异常
// 也可使用 models.DB().Transaction()来替代，但需要自行处理异常信息
// 事务内的连接(tx) 必须与开启事务的连接（tx）一致，如果事务中用了新建立的连接，会与事务无关
func (base BaseService) Transaction(invoker func(tx *gorm.DB) error) (err error) {
	err = models.Db().Transaction(invoker)
	bizcode.DbCheck(err)
	return
}

// 判定记录是否存在
func (base BaseService) IsNotFound(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}

// 判定是否是错误
func (base BaseService) IsError(err error) bool {
	return !base.IsErrorNil(err) && !base.IsNotFound(err)
}

// 判定是否是错误
func (base BaseService) IsErrorNil(err error) bool {
	return err == nil
}

// 判定是否是错误
func (base BaseService) Error(err string) error {
	return errors.New(err)
}

// 判定是否是错误
func (base BaseService) NotFound(field string) error {
	return base.Error(field + " Record Not Found")
}

// 判定是否是存在
func (base BaseService) RecordExists(field string) error {
	return base.Error(field + " Record Exists")
}

// 转出指定的字段
func (base BaseService) FilterField(data interface{}, out interface{}) error {
	return util.Interface2Interface(data, out)
}
