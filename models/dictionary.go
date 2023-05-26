package models

import (
	"wood/pkg/setting"
	"wood/pkg/util/convert"
)

// 字典表
type Dictionary struct {
	Model
	Name     string `gorm:"column:name; size:50; " json:"name"`                               // 砍伐队名称
	Key      string `gorm:"column:key; size:50; " json:"key"`                                 // 砍伐队名称
	Value    string `gorm:"column:value; size:100; " json:"value"`                            // 砍伐队名称
	DeleteAt int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"` //删除时间
}

func (m *Dictionary) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "dictionary"
}

func (m *Dictionary) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
