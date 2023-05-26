package mapper

// 使用方法
// 在继承的子mapper中继承该父mapper，并实现getDb()方法的重写，如：
//func (testMapper *TestMapper) getDb() *TestMapper {
//	testMapper.db = testMapper.GetModel(new(models.CaseLabelLog).TableName())
//	return testMapper
//}
//err := caseLogMapper.getDb().GetPage(&list, &models.CaseLabelLog{}, field, filter.Page, filter.Limit, &total, whereOrder...)
//err := caseLogMapper.getDb().Find(&list, &models.CaseLabelLog{}, field, 0, whereOrder...)
//err := caseLogMapper.getDb().Last(&list, &models.CaseLabelLog{CaseId: 1234}, field)
//err := caseLogMapper.getDb().Scan(&list, &models.CaseLabelLog{CaseId: 1234}, field, 0)
//err := caseLogMapper.getDb().Last(&list, &models.CaseLabelLog{CaseId: 1234}, field)

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"strings"
	"wood-client-api/models"
	bizcode "wood-client-api/pkg/bizerror"
	"wood-client-api/pkg/logging"
	"wood-client-api/pkg/util"
)

type PageWhereOrder struct {
	Order string
	Where string
	Value []interface{}
}

type Mapper interface {
	GetModel(tableName string) gorm.DB

	Insert(model interface{}) error

	Update(model interface{}) error

	Delete(model interface{}) error

	Scan(out interface{}, where interface{}, cols []string, limit int64, whereOrder ...PageWhereOrder) error

	Transaction(invoker func(tx *gorm.DB) error) (err error)

	UpdateWithOut(where interface{}, value interface{}, out interface{}, selectFields ...[]string) error

	First(out interface{}, where interface{}, cols []string) (bool, error)

	Last(out interface{}, where interface{}, cols []string) (notFound bool, err error)

	Find(out interface{}, where interface{}, selectFields []string, limit int64, whereOrder ...PageWhereOrder) error
}

type BetweenParam struct {
	Start interface{}
	End   interface{}
}

type Preload struct {
	Column     string
	Conditions []interface{}
}

type BaseMapper struct {
	Db             *gorm.DB
	Table          string
	JoinTable      string
	ALis           string      // 要使用这个别名,需要定义好查询字段,否则会出现因GORM不支持别名导致的一些sql语法问题,一般默认为空即可
	PreloadList    []*Preload  // 用来实现关联模型
	WhereCondition interface{} // 使用这个,可以使得查询比较灵活,可传参,也可以使用这种方式
	Field          []string
}

var db *gorm.DB

const Active_Status = 1
const Inactive_Status = 2

type Map map[string]interface{}

var DeleteMap = Map{"delete_at": util.Now().Year(), "update_time": util.Now().UnixMilli()}

func InitMapper() *gorm.DB {
	db := models.Db()
	_ = BaseMapper{Db: db}
	return db
}

func (baseMapper *BaseMapper) GetModel(tableName string) *gorm.DB {
	alis := ""
	if baseMapper.ALis != "" {
		alis = " AS " + baseMapper.ALis
	}
	tableName += alis
	return InitMapper().Table(tableName)
}

func (baseMapper *BaseMapper) GetDb() *BaseMapper {
	return baseMapper
}

// Join 连接
// 支持多个Join，可以在 On 后面的字段定义表别名，或是默认表别名
// tableName 表名
// first 必须为驱动表的主键,不能为被驱动表的外键,否则别名对不上,出现sql语句错误
// op 操作符
// second 必须为连接表的外键,否则会出现别名对应不上的问题
// joinType 连接方式
func (baseMapper *BaseMapper) Join(model models.BaseModel, first string, op string, second string, joinType string) *BaseMapper {
	firstTableAlis := baseMapper.ALis
	if firstTableAlis == "" {
		firstTableAlis = baseMapper.Table
	}
	// 如果 first字段中含有英文字符'.'，则说明是有表字段别名的，是有定义别名的
	fmt.Println("First = ", first)
	if !strings.Contains(first, ".") && baseMapper.JoinTable != "" { //如果前面有使用Join操作，默认将前面的tableName作为下一次join的第一个表
		firstTableAlis = baseMapper.JoinTable
		first = baseMapper.JoinTable + "." + first
	}
	fmt.Println("First = ", first)
	fmt.Println("Second = ", second)
	// 如果 second字段中含有英文字符'.'，则说明是有表字段别名的，是有定义别名的
	if !strings.Contains(second, ".") { //如果前面有使用Join操作，默认将前面的tableName作为下一次join的第一个表
		second = model.TableName() + "." + second
	}
	fmt.Println("Second = ", second)
	joinString := fmt.Sprintf("%s JOIN %s ON %s %s %s", joinType, model.TableName(), first, op, second)
	baseMapper.Db = baseMapper.Db.Joins(joinString)
	baseMapper.JoinTable = model.TableName()
	return baseMapper
}

// InnerJoin 内连接
func (baseMapper *BaseMapper) InnerJoin(model models.BaseModel, first string, op string, second string) *BaseMapper {
	return baseMapper.Join(model, first, op, second, "INNER")
}

// LeftJoin 左连接
func (baseMapper *BaseMapper) LeftJoin(model models.BaseModel, first string, op string, second string) *BaseMapper {
	return baseMapper.Join(model, first, op, second, "LEFT")
}

// RightJoin 右连接
func (baseMapper *BaseMapper) RightJoin(model models.BaseModel, first string, op string, second string) *BaseMapper {
	return baseMapper.Join(model, first, op, second, "RIGHT")
}

// Transaction 事物处理
func (baseMapper *BaseMapper) Transaction(invoker func(tx *gorm.DB) error) (err error) {
	tx := models.Db()
	err = tx.Transaction(invoker)
	bizcode.DbCheck(err)
	return
}

// Insert 插入数据
func (baseMapper *BaseMapper) Insert(model interface{}) error {
	err := baseMapper.GetDb().Db.Create(model).Error
	if err != nil {
		logging.Error("插入数据错误：%#v", model)
	}
	//bizcode.DbCheck(err)
	fmt.Printf("返回数据：%#v", model)
	return err
}

// Update 更新数据
// 可带Where
// service.demoService.GetMapper().Update(&model, &where)
func (baseMapper *BaseMapper) Update(model, where interface{}) error {
	db := baseMapper.GetDb().Db
	if where != nil {
		db = db.Where(where)
	}
	db = baseMapper.GetPreload(db)
	db = baseMapper.GetWhere(db)
	err := db.Updates(model).Error
	bizcode.DbCheck(err)
	return err
}

// UpdateById 通过id更新数据
// service.demoService.GetMapper().UpdateById(123, &model)
func (baseMapper *BaseMapper) UpdateById(id int, model interface{}) error {
	baseMapper.WhereCondition = map[string]interface{}{"id": id}
	err := baseMapper.Update(model, nil)
	bizcode.DbCheck(err)
	return err
}

// SoftDeleteById 通过id更新数据
func (baseMapper *BaseMapper) SoftDeleteById(id int) error {
	return baseMapper.UpdateById(id, map[string]interface{}{"delete_at": util.Now().Year(), "update_time": util.Now().UnixMilli()})
}

// UpdateWithOut 更新并返回
// 用法
// 可带Where()
// where := models.Licence{Model: models.Model{ID: 14}}
// updateDta := models.Licence{Valid: 2}
// out := models.Licence{}
// err := service.LicenceService.GetMapper().UpdateWithOut(&value, &out, &selectFields)
func (baseMapper *BaseMapper) UpdateWithOut(value interface{}, out interface{}, selectFields ...[]string) error {
	db := baseMapper.GetDb().Db
	db = baseMapper.GetWhere(db)
	err := db.Updates(value).Error
	if err != nil {
		return err
	}
	if len(selectFields) > 0 {
		return db.Select(selectFields[0]).Find(out).Error
	}
	return db.Find(out).Error
}

// Delete 删除
func (baseMapper *BaseMapper) Delete(model interface{}) error {
	err := baseMapper.GetDb().Db.Delete(model).Error
	return err
}

// First 查询一条通过主键ASC排序的第一条数据
// out：返回字段，调用:&models.modelName{}
// where：为查询条件 调用：&models.modelName{Key:value}
// cols 查询字段
// 可带Where()
func (baseMapper *BaseMapper) First(out interface{}, cols []string) (notFound bool, err error) {
	db := baseMapper.GetDb().Db
	if len(cols) > 0 {
		db = db.Select(cols)
	} else {
		db = db.Select(baseMapper.Field)
	}
	db = baseMapper.GetPreload(db)
	db = baseMapper.GetWhere(db)
	err = db.First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
		if !notFound {
			bizcode.DbCheck(err)
		}
	}
	return
}

// Last 获取最后一条通过主键DESC排序的数据
// out：返回字段，调用:&models.modelName{}
// where：为查询条件 调用：&models.modelName{Key:value}
// 可带Where()
func (baseMapper *BaseMapper) Last(out interface{}, cols []string) (notFound bool, err error) {
	db := baseMapper.GetDb().Db
	if len(cols) > 0 {
		db = db.Select(cols)
	} else {
		db = db.Select(baseMapper.Field)
	}
	db = baseMapper.GetPreload(db)
	db = baseMapper.GetWhere(db)
	err = db.Last(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
		if !notFound {
			bizcode.DbCheck(err)
		}
	}
	return
}

// Count 统计
func (baseMapper *BaseMapper) Count(where interface{}) (int64, error) {
	db := baseMapper.GetDb().Db
	if where != nil {
		db = db.Where(where)
	}
	db = baseMapper.GetWhere(db)
	var count int64
	err := db.Count(&count).Error
	return count, err
}

// GetWhereOrder 获得查询条件和排序
func (baseMapper *BaseMapper) GetWhereOrder(db *gorm.DB, whereOrder ...PageWhereOrder) *gorm.DB {
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	return db
}

// GetPreload 获得关联查询
func (baseMapper *BaseMapper) GetPreload(db *gorm.DB) *gorm.DB {
	if len(baseMapper.PreloadList) > 0 {
		for _, item := range baseMapper.PreloadList {
			db = db.Preload(item.Column, item.Conditions...)
		}
	}
	return db
}

// GetWhere 获得查询条件
func (baseMapper *BaseMapper) GetWhere(db *gorm.DB) *gorm.DB {
	if baseMapper.WhereCondition != nil {
		db = db.Where(baseMapper.WhereCondition)
	}
	return db
}

// Scan 获取多行数据
// out 返回数据 用法：var res []*models.modelName  Scan(&res)
// where 条件查询 调用：&models.modelName{Key:value}  或者是map
// cols 字段
// whereOrder 拼接where语句和order语句
func (baseMapper *BaseMapper) Scan(out interface{}, where interface{}, cols []string, limit int64, whereOrder ...PageWhereOrder) error {
	db := baseMapper.GetDb().Db
	if where != nil {
		db = db.Where(where)
	}
	if len(whereOrder) > 0 {
		db = baseMapper.GetWhereOrder(db, whereOrder...)
	}
	if limit > 0 {
		db = db.Limit(limit)
	}
	if len(cols) > 0 {
		db = db.Select(cols)
	} else {
		db = db.Select(baseMapper.Field)
	}
	db = baseMapper.GetWhere(db)
	//db = baseMapper.GetPreload(db)
	err := db.Scan(out).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		bizcode.DbCheck(err)
	}
	return err
}

// Find 获取多条数据
// out：返回字段，调用:&models.modelName{}
// where：为查询条件 调用：&models.modelName{Key:value} 或者是map
func (baseMapper *BaseMapper) Find(out interface{}, where interface{}, selectFields []string, limit int64, whereOrder ...PageWhereOrder) error {
	db := baseMapper.GetDb().Db.Where(where)
	db = baseMapper.GetWhereOrder(db, whereOrder...)
	if limit > 0 {
		db = db.Limit(limit)
	}
	if len(selectFields) > 0 {
		db = db.Select(selectFields)
	} else {
		db = db.Select(baseMapper.Field)
	}
	db = baseMapper.GetPreload(db)
	err := db.Find(out).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		bizcode.DbCheck(err)
	}
	return err
}

// WhereOp 拼接操作
func (baseMapper *BaseMapper) WhereOp(param interface{}, field string, op string) *PageWhereOrder {
	paramType := reflect.TypeOf(param)
	paramValue := reflect.ValueOf(param)
	paramValueType := paramType.Kind().String()
	if paramValueType == "string" {
		if paramValue.String() != "" {
			var arr []interface{}
			value := paramValue.String()
			op = strings.ToUpper(op)
			if op == "IN" || op == "NOT IN" {
				ids := strings.Split(value, ",")
				arr = append(arr, ids)
				return &PageWhereOrder{Where: fmt.Sprintf("%v %v (?)", field, op), Value: arr}
			} else {
				if op == "LIKE" {
					value = "%" + value + "%"
				}
				arr = append(arr, value)
				return &PageWhereOrder{Where: fmt.Sprintf("%v %v ?", field, op), Value: arr}
			}
		}
	} else if paramValueType == "int" || paramValueType == "int8" || paramValueType == "int32" || paramValueType == "int64" || paramValueType == "uint64" {
		var arr []interface{}
		arr = append(arr, paramValue.Int())
		return &PageWhereOrder{Where: fmt.Sprintf("%v %v ?", field, op), Value: arr}
	} else if paramValueType == "float32" || paramValueType == "float64" {
		var arr []interface{}
		arr = append(arr, paramValue.Float())
		return &PageWhereOrder{Where: fmt.Sprintf("%v %v ?", field, op), Value: arr}
	} else if paramValueType == "struct" {
		var arr []interface{}
		op = strings.ToUpper(op)
		if op == "BETWEEN" || op == "NOT BETWEEN" {
			betweenData := (param).(BetweenParam)
			arr = append(arr, betweenData.Start)
			arr = append(arr, betweenData.End)
			return &PageWhereOrder{Where: fmt.Sprintf("%v %v ? AND ?", field, op), Value: arr}
		}
	}
	return nil
}

// WhereEq 等于
func (baseMapper *BaseMapper) WhereEq(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "=")
}

// WhereNEq 不等于
func (baseMapper *BaseMapper) WhereNEq(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "<>")
}

// WhereGT 大于
func (baseMapper *BaseMapper) WhereGT(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, ">")
}

// WhereEGT 大于等于
func (baseMapper *BaseMapper) WhereEGT(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, ">=")
}

// WhereLT 小于
func (baseMapper *BaseMapper) WhereLT(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "<")
}

// WhereELT 小于等于
func (baseMapper *BaseMapper) WhereELT(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "<=")
}

// WhereLike LIKE
func (baseMapper *BaseMapper) WhereLike(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "LIKE")
}

// IN
func (baseMapper *BaseMapper) WhereIn(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "IN")
}

// NOT IN
func (baseMapper *BaseMapper) WhereNotIn(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "NOT IN")
}

// BETWEEN 传一个结构体：struct { Start: start, End: end}
// 示例：mapper.BetweenParam{Start: 7, End: 9}}
func (baseMapper *BaseMapper) WhereBetween(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "BETWEEN")
}

// BETWEEN 传一个结构体：struct { Start: start, End: end}
// 示例：mapper.BetweenParam{Start: 7, End: 9}}
func (baseMapper *BaseMapper) WhereNotBetween(param interface{}, field string) *PageWhereOrder {
	return baseMapper.WhereOp(param, field, "NOT BETWEEN")
}

// WhereOrder ORDER排序
func (baseMapper *BaseMapper) WhereOrder(orderKey string) *PageWhereOrder {
	if orderKey != "" {
		return &PageWhereOrder{Order: orderKey}
	}
	return nil
}

// GetPage 获取分页数据
// var userList []models.User
// var field []string
// var total *uint64
// var whereOrder mapper.PageWhereOrder
// 用法：GetPage(&userList, &models.User{}, field, 1, 10, &total, whereOrder)
func (baseMapper *BaseMapper) GetPage(out interface{}, where interface{}, field []string, pageIndex, pageSize uint64, totalCount *uint64, whereOrder ...PageWhereOrder) error {
	db := baseMapper.GetDb().Db.Where(where)
	db = baseMapper.GetWhereOrder(db, whereOrder...)
	db = baseMapper.GetWhere(db)
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	db = baseMapper.GetPreload(db)
	db = db.Offset((pageIndex - 1) * pageSize).Limit(pageSize)
	if len(field) > 0 {
		db = db.Select(field)
	} else {
		db = db.Select(baseMapper.Field)
	}
	return db.Find(out).Error
}

// Save 保存 如果不存在主键，则做插入操作，否则做更新操作
func (baseMapper *BaseMapper) Save(value interface{}) error {
	return baseMapper.GetDb().Db.Save(value).Error
}

// Where 查询条件
func (baseMapper *BaseMapper) Where(query interface{}, args ...interface{}) *BaseMapper {
	baseMapper.Db = baseMapper.GetDb().Db.Where(query, args)
	return baseMapper
}

// Select 查询字段
func (baseMapper *BaseMapper) Select(query interface{}, args ...interface{}) *BaseMapper {
	baseMapper.Db = baseMapper.GetDb().Db.Select(query, args)
	return baseMapper
}

// Limit 查询条数限制
func (baseMapper *BaseMapper) Limit(limit int) *BaseMapper {
	baseMapper.Db = baseMapper.GetDb().Db.Limit(limit)
	return baseMapper
}

// Group 查询条数限制
func (baseMapper *BaseMapper) Group(query string) *BaseMapper {
	baseMapper.Db = baseMapper.GetDb().Db.Group(query)
	return baseMapper
}

// Having specify HAVING conditions for GROUP BY
func (baseMapper *BaseMapper) Having(query interface{}, values ...interface{}) *BaseMapper {
	baseMapper.Db = baseMapper.GetDb().Db.Having(query, values...)
	return baseMapper
}

// Preloads preload associations with given conditions
func (baseMapper *BaseMapper) Preloads(column string, conditions ...interface{}) *BaseMapper {
	fmt.Println("查询条件：", conditions)
	baseMapper.Db = baseMapper.GetDb().Db.Preload(column, conditions...)
	return baseMapper
}

// Pluck 查询某列数据
// out：返回字段，必须为切片,不能为struct,例如:[]string
// column 查询字段
// 可带Where()
// 用法:
// var user []string
// service.UserService.GetMapper().Where(map[string]interface{}{"id": 19}).Pluck(&user, "nick_name")
func (baseMapper *BaseMapper) Pluck(out interface{}, column string) (notFound bool, err error) {
	db := baseMapper.GetDb().Db
	err = db.Pluck(column, out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
		if !notFound {
			bizcode.DbCheck(err)
		}
	}
	return
}

// Raw 直接执行sql语句
// 用法:
// users, _ := service.UserService.GetMapper().Raw(&user, "Select nick_name FROM t_user where id = ? and mobile = ?", 1, "18675343868")
func (baseMapper *BaseMapper) Raw(out interface{}, sql string, where ...interface{}) (notFound bool, err error) {
	db := InitMapper()
	err = db.Raw(sql, where...).Scan(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
		if !notFound {
			bizcode.DbCheck(err)
		}
	}
	return
}

// Value 查询某列一条数据
// column 查询字段
// 可带Where()
// 用法:
// nick_name, _, _ := service.UserService.GetMapper().Where(map[string]interface{}{"id": 19}).Value("nick_name")
func (baseMapper *BaseMapper) Value(column string) (out interface{}, notFound bool, err error) {
	var value []string
	notFound, err = baseMapper.Limit(1).Pluck(&value, column)
	if !notFound {
		out = value[0]
	}
	return
}

// GetField 多个表的字段合并到一个切片中
func (baseMapper *BaseMapper) GetField(field map[string][]string) []string {
	fieldSlice := make([]string, 0, 0)
	for key, value := range field {
		length := len(value)
		for i := 0; i < length; i++ {
			value[i] = fmt.Sprintf("%v.%v", key, value[i])
		}
		fieldSlice = append(fieldSlice, value...)
	}
	return fieldSlice
}
