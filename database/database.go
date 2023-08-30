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
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DbUser, config.DbPass, config.DbHost, config.DbPort, config.DbSchema, config.DbCharset,
	)
	// 连接数据库
	var err error
	fmt.Println(dsn)
	DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数名称作为表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Failed to connect database.")
	}

	return DBConn
}
