package main

import (
	_ "swaggo/docs"

	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//making an instance of the type DB from the gorm package
var db *gorm.DB = nil
var err error

// @title gin+gorm crud 測試swagger
// @version 1.0
// @description gin+gorm crud 測試swagger
// @license.name Apache 2.0
// @contact.name go-swagger幫助文檔
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
func main() {
	//establishing connection with mysql database 'CRUD'
	dsn := "root:root@tcp(localhost:8889)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&Log{})

	router := gin.Default()

	v1 := router.Group("/sql")
	{
		v1.GET("/:id", Show)
		v1.POST("", Create)
		v1.PATCH("/:id", Update)
		v1.DELETE("/:id", Delete)
		v1.GET("", Time)
	}
	

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	//start the server and listen on the port 8000
	router.Run(":8000")
}