package main

import (
	_ "main/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title gin+redis CRUD swagger
// @version 1.0
// @description gin+redis CRUD swagger
// @license.name Apache 2.0
// @contact.name go-swagger.doc
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8001
// @BasePath /
func main() {

	router := gin.Default()

	v1 := router.Group("/redis")
	{
		v1.GET("/all", GetAllHash) // 取得所有資料
		v1.GET("", GetHash)
		v1.POST("", CreateHash)
		v1.PATCH("", UpdateHash)
		v1.DELETE("", DeleteHash)
		
	}
	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	router.Run(":8001")

}
