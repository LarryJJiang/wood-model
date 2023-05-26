package models

import (
	"reflect"
	"strings"
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 客户/林场表
type Customer struct {
	Model
	UserId   int    `gorm:"column:user_id; size:11; not null; default:0;" json:"user_id"`     // 用户ID
	Name     string `gorm:"column:name; size:50; " json:"name"`                               // 名称
	Email    string `gorm:"column:email; size:150; " json:"email"`                            // 运输公司简称
	Status   int8   `gorm:"column:status; size:2; not null; default:0;" json:"status"`        // 状态
	DeleteAt int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"` //删除时间
	Account  string `gorm:"-" json:"account"`
	RoleId   int    `gorm:"-" json:"role_id"`
}

func (m *Customer) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "customer"
}

func (m *Customer) GetField() map[string]interface{} {
	v := reflect.TypeOf(m)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	var data = map[string]interface{}{"ID": "id", "CreateTime": "create_time", "UpdateTime": "update_time"}
	for i := 0; i < v.NumField(); i++ {
		var field = v.Field(i).Tag.Get("gorm")
		if field != "" {
			fieldArray := strings.Split(field, ";")
			for _, value := range fieldArray {
				if strings.Contains(value, "column:") {
					field = strings.Replace(value, "column:", "", 1)
					data[v.Field(i).Name] = field
				}
			}
		}
	}
	return data
}

func (m *Customer) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
