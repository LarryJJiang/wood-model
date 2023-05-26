package models

import (
	"woods/pkg/setting"
	"woods/pkg/util/convert"
)

// 实时消息表
type Message struct {
	Model
	CrewId     int    `gorm:"column:crew_id; size:11; not null; default:0;" json:"crew_id"`           // 砍伐队id
	SendUserId int    `gorm:"column:send_user_id; size:11; not null; default:0;" json:"send_user_id"` // 发送者user_id
	ToUserId   int    `gorm:"column:to_user_id; size:11; not null; default:0;" json:"to_user_id"`     // 接收者user_id
	Content    string `gorm:"column:content; size:11;" json:"content"`                                // 消息内容
	DeleteAt   int    `gorm:"column:delete_at; size:11; not null; default:0;" json:"delete_at"`       //删除时间
}

func (m *Message) TableName() string {
	return setting.DatabaseSetting.TablePrefix + "message"
}

func (m *Message) GetFieldSlice() []string {
	return convert.GetFieldSlice(m)
}
