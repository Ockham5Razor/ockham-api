package database

import (
	"fmt"
	"gol-c/config"
	"gol-c/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Init() {
	// 读取配置
	conf := config.GetConfig()
	dbConf := conf.Db
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Schema, dbConf.Charset,
	)
	// 连接数据库
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}

	// 移植数据库
	model.Migrate(DBConn)
}
