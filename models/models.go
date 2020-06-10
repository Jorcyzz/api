package models

import (
	"api/pkg/logging"
	"api/pkg/setting"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetUp() {
	var (
		err          error
		databaseType = setting.DatabaseSetting.Type     //数据库类型
		user         = setting.DatabaseSetting.User     //数据库的用户
		pass         = setting.DatabaseSetting.Password //数据库的密码
		host         = setting.DatabaseSetting.Host     //数据库地址
		name         = setting.DatabaseSetting.Name     //数据库名称
	)

	db, err := gorm.Open(databaseType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, pass, host, name))
	if err != nil {
		logging.Fatal("数据库链接失败", err) //数据库链接失败是致命的错误，链接失败后可以关闭程序了，所以使用logging.Fatal方法
	}
	defer db.Close()
	//设置表名称的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true) //设置禁用表名的复数形式
	db.LogMode(true)       //打印日志，本地调试的时候可以打开看执行的sql语句

	db.DB().SetMaxIdleConns(10)  //设置空闲时的最大连接数
	db.DB().SetMaxOpenConns(100) //设置数据库的最大打开连接数
	db.LogMode(true)             //打印日志，本地调试的时候可以打开看执行的sql语句
	db.AutoMigrate(&User{})      //自动检查 Tag 结构是否变化，变化则进行迁移，需要的参数为数据库模型结构体
}
