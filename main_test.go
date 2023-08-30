package main

import (
	"ockham-api/model"
	"ockham-api/util"
	"testing"
)

func TestInitDB(t *testing.T) {
	dbConn := util.InitDatabase()
	// 移植数据库
	model.Migrate(dbConn)
	// 初始化数据
	model.InitData(dbConn)
}
