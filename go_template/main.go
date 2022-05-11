package main

import (
	"main/api"
	"main/config"
	_ "main/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// mysql、es、redis 連線
func init() {
	api.Mysql_client()
	api.Redis_client()
	api.Elastic_client()
}
// viper 讀取 server 連線設定
var Server_port = config.Viper.GetString("databases.server.port")

// @title go_template swagger
// @version 1.0
// @description go_template swagger
// @license.name Apache 2.0
// @contact.name go-swagger.doc
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host Server_port
// @BasePath /
// Server_port 無法讀取變數
// mysql、es、redis 路由配置
func main() {

	router := gin.Default()

	// elastic
	elastic := router.Group("/es")
	{
		elastic.GET("", api.ESquery)
	}

	// redis
	redis := router.Group("/redis")
	{
		redis.GET("/all", api.GetAllHash) // 取得所有資料
		redis.GET("", api.GetHash)
		redis.POST("", api.CreateHash)
		redis.PATCH("", api.UpdateHash)
		redis.DELETE("", api.DeleteHash)
		
	}
	// mysql
	mysql := router.Group("/sql")
	{
		mysql.GET("/:id", api.Show)
		mysql.POST("", api.Create)
		mysql.PATCH("/:id", api.Update)
		mysql.DELETE("/:id", api.Delete)
		mysql.GET("", api.Time)
	}


	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	router.Run(Server_port)

}
