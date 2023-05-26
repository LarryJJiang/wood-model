package bizcode

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
	"strconv"
	"woods/pkg/logging"
)

var (
	BizCode500 = BizCode{200, 500, "系统错误"}
	BizCode400 = BizCode{200, 400, "参数错误"}
)

type BizCode struct {
	HttpCode int
	Code     int
	Msg      interface{}
}

var BizCodeHander BizCode

func init() {
	BizCodeHander = BizCode{}
}

func (bizCode BizCode) PanicError(httpCode int, msg interface{}) {
	bizCode.HttpCode = httpCode
	bizCode.Code = httpCode
	message := GetMsg(httpCode)
	if msg != nil {
		msgType := reflect.TypeOf(msg)
		fmt.Printf("类型：%#v\n", msgType)
		if msgType.Kind().String() == "string" {
			bizCode.Msg = message + "：" + msg.(string)
		} else {
			bizCode.Msg = msg
		}
	} else {
		bizCode.Msg = message
	}
	panic(bizCode)
}

func (bizCode BizCode) Error() string {
	return strconv.FormatInt(int64(bizCode.Code), 10) + "-" + bizCode.Msg.(string)
}

func Check(err error) {
	if err != nil {
		fmt.Sprintf("错误：%#v", err)
		panic(err.Error())
	}
}

func CheckBizCode(err error, bizCode BizCode) {
	if err != nil {
		panic(err.Error() + bizCode.Error())
	}
}

func DbCheck(err error) {
	if err != nil && err != gorm.ErrRecordNotFound {
		logging.Error("数据库错误信息：%#v", err)
		panic(BizCode{http.StatusOK, ErrorDb, "数据库操作异常-" + err.Error()})
	}
}
