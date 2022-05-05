package main

import (
	_ "swaggo/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/gorm"

	//elastic
	"github.com/olivere/elastic/v7"
)

//making an instance of the type DB from the gorm package
var db *gorm.DB = nil
var err error

//elastic
var client *elastic.Client
var host = "http://140.137.219.56:9200/"

// @title gin+ElasticSearch crud 測試 swagger
// @version 1.0
// @description gin+gorm crud 測試swagger
// @license.name Apache 2.0
// @contact.name go-swagger幫助文檔
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host 127.0.0.1:8001
// @BasePath /
func main() {
	

	router := gin.Default()

	v2 := router.Group("/es")
	{
		v2.GET("", ESquery)

	}

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8001
	router.Run(":8001")
}