package models

import (
	"Go-App/pkg/logging"
	"Go-App/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetUp() {
	var (
		err          error
		databaseType = setting.DatabaseSetting.Type
		user         = setting.DatabaseSetting.User
		password     = setting.DatabaseSetting.Password
		host         = setting.DatabaseSetting.Host
		name         = setting.DatabaseSetting.Name
	)

	db, err = gorm.Open(databaseType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, password, host, name))
	if err != nil {
		logging.Fatal("数据库连接失败: ", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	// 禁用表明的复数形式
	db.SingularTable(true)
	// 打印日志，本地调试的时候可以打开看执行的sql语句
	db.LogMode(true)
	db.AutoMigrate(&User{})

	// 设置空闲时的最大连接数
	db.DB().SetMaxIdleConns(10)
	// 设置数据库的最大打开连接数
	db.DB().SetMaxOpenConns(100)

}
