package main

import (
	"main/api"
	"main/config"
	_ "main/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Influx 連線
func init() {
	api.ConnInflux()
}

// viper 讀取 server 連線設定
var Server_port = config.Viper.GetString("databases.server.port")

// @title go_DiskAPI swagger
// @version 1.0
// @description disk 預警天數 api
// @license.name Apache 2.0
// @contact.name go-swagger.doc
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
func main() {

	router := gin.Default()

	// influx
	influx := router.Group("/influx")
	{
		influx.GET("", api.Calculate) // 取得 Disk 90% 滿載天數
		influx.GET("/2", api.Query_box)	// 取得 Disk 圖表資料
	} 

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	router.Run(Server_port)

}
