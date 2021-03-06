package model

import (
	"fmt"
	otgorm "github.com/eddycjy/opentracing-gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/noChaos1012/tour/blog_service/global"
	"github.com/noChaos1012/tour/blog_service/pkg/setting"
	"time"
)

const (
	STATE_OPEN  = 1
	STATE_CLOSE = 0
)

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id,omitempty"` //omitempty当没有取值时，不返回
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
	CreatedOn  uint32 `json:"created_on,omitempty"`
	ModifiedOn uint32 `json:"modified_on,omitempty"`
	DeletedOn  uint32 `json:"deleted_on,omitempty"`
	IsDel      uint8  `json:"is_del,omitempty"`
}

//创建GORM引擎
func NewDBEngine(dbSetting *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(dbSetting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSetting.UserName,
		dbSetting.Password,
		dbSetting.Host,
		dbSetting.DBName,
		dbSetting.Charset,
		dbSetting.ParseTime))

	if err != nil {
		return nil, err
	}
	//debug模式就打开日志
	if global.ServerSetting.RunMode == "debug" {
		db.LogMode(true)
	}
	db.SingularTable(true)
	//注册回调
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)

	db.DB().SetMaxIdleConns(dbSetting.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSetting.MaxOpenConns)

	otgorm.AddGormCallbacks(db) //注册链路回调
	return db, nil
}

//配置回调函数
//新增行为的回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		//判断是否包含所需字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			//如果是空则配置值
			if createTimeField.IsBlank {
				_ = createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				_ = modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新行为的回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//获取当前设置标识的gorm:update_column字段属性
	//若没有，则在更新回调中设置默认字段ModifiedOn为当前时间戳
	if _, ok := scope.Get("gorm:update_column"); !ok {
		_ = scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除行为的回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		//获取删除标识 gorm:delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteOnField, hasDeletedOnField := scope.FieldByName("DeleteOn")
		isDelField, hasIsDelField := scope.FieldByName("IsDel")
		//进行软删除
		if !scope.Search.Unscoped && hasDeletedOnField && hasIsDelField {
			nowTime := time.Now().Unix()
			scope.Raw(fmt.Sprintf("UPDATE %v SET %v=%v,%v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteOnField.DBName),
				scope.AddToVars(nowTime),
				scope.Quote(isDelField.DBName),
				scope.AddToVars(1),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption), )).Exec()
		} else {
			scope.Raw(fmt.Sprintf("DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				//scope.CombinedConditionSql 条件sql拼装
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption))).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
