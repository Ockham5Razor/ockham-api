package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	v1 "gol-c/api/v1"
	"gol-c/database"
	_ "gol-c/docs"
	"gol-c/model"
)

// @title Gol-C API
// @version 1.0
// @description All APIs of Gol-C Project
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization

// @BasePath /api
func main() {

	dbConn := database.Init()

	// 移植数据库
	model.Migrate(dbConn)
	// 初始化数据
	model.InitData(dbConn)

	// 创建 HTTP 路由
	r := gin.Default()
	r.Use(Cors())                              // Using Middleware to allowing CORS
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	v1.ApiV1(r)

	// 定义 404 响应
	v1.DefaultHttp404(r)

	// 启动 HTTP 服务
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic("Failed to listen HTTP port.")
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
