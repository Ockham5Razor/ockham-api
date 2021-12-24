package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"ockham-api/config"
)

var DBConn *gorm.DB

func Init() *gorm.DB {
	// 读取配置
	conf := config.GetConfig()
	dbConf := conf.Db
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConf.User, dbConf.Pass, dbConf.Host, dbConf.Port, dbConf.Schema, dbConf.Charset,
	)
	// 连接数据库
	var err error
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数名称作为表名
		},
	})
	if err != nil {
		panic("Failed to connect database.")
	}

	return DBConn
}
