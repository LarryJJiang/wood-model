package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
	"wood/pkg/setting"
)

type Model struct {
	ID         int   `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
}

var db *gorm.DB
var NotFound = gorm.ErrRecordNotFound

func Db() *gorm.DB {
	return db
}

func SetupDb() {
	var err error
	//dsnFormat := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	dsnFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(dsnFormat,
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Port,
		setting.DatabaseSetting.Name)

	db, err = gorm.Open(setting.DatabaseSetting.Type, dsn)
	if err != nil {
		log.Fatalf("models.SetupDb err: %v", err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		if strings.HasPrefix(defaultTableName, setting.DatabaseSetting.TablePrefix) {
			return defaultTableName
		}
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	// 日志模式 设置为true之后控制台会输出对应的SQL语句
	if setting.DatabaseSetting.LogMode {
		db.LogMode(setting.DatabaseSetting.LogMode)
	}
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(30 * time.Minute)
	//db.AutoMigrate(&User{}, &SystemUser{}, &Crew{})
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().UnixMilli()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", time.Now().UnixMilli())
	}
}

type BaseModel interface {
	TableName() string
	//GetField() map[string]interface{}
	GetFieldSlice() []string
}
