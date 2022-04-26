package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
type Log struct {
	// 自定義pk、欄位名稱、type
	Id int64
	Time string
	IP string
	Data string
	Status int64
}
// 自定義 TableName
func (u Log)TableName()string{
	return "test_log"
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", dsn)

	// 提交資料庫的修改
	// db.AutoMigrate(&User{})
	defer db.Close() // 關閉資料庫連線

	// 通過接口，接收前端post過來的json，json繼承user結構體後，再寫進db
	r:=gin.Default()

	// Post
	r.POST("/post",func(c *gin.Context) {
		var log Log
		// 取得 user 結構體 (繼承相關屬性)
		_ = c.BindJSON(&log)
		db.Create(&log)
	})

	// Get id
	r.GET("/get/:id",func(c *gin.Context) {
		id := c.Param("id")
		var log Log
		// 取得 user 結構體 (繼承相關屬性)
		_ = c.BindJSON(&log)
		db.First(&log, "id = ?", id)
		c.JSON(200, gin.H{
			"r":log,
		})
	})

	// Get time
	r.GET("/get/time/:Time",func(c *gin.Context) {
		time := c.Param("Time")
		var log []Log
		// 取得 user 結構體 (繼承相關屬性)
		_ = c.BindJSON(&log)
		db.Where("time < ?", time).Find(&log)
		c.JSON(200, gin.H{
			"r":log,
		})
	})

	// Get time 區間
	r.GET("/get/time_start_end",func(c *gin.Context) {
		start := c.Query("Start")
		end := c.Query("End")
		var log []Log
		// 取得 user 結構體 (繼承相關屬性)
		_ = c.BindJSON(&log)
		db.Where("time BETWEEN ? AND ?", start, end).Find(&log)
		c.JSON(200, gin.H{
			"r":log,
		})
	})

	// form 表單接收
	r.POST("/form", func(c *gin.Context) {
        //types := c.DefaultPostForm("type", "post")
        id := c.PostForm("username")
		start := c.PostForm("start")
		end := c.PostForm("end")
		var log []Log
		// 取得 user 結構體 (繼承相關屬性)
		_ = c.BindJSON(&log)

		// 查詢時間
		// db.Where("time < ?", userpassword).Find(&log)
		// c.JSON(http.StatusOK, gin.H{
		// 	"r":log,
		// })

		// 查詢時間區間
		_ = c.BindJSON(&log)
		db.Where("time BETWEEN ? AND ?", start, end).Find(&log)
		c.JSON(200, gin.H{
			"r":log,
		})


		// 查詢id
		db.First(&log, "id = ?", id)
		c.JSON(http.StatusOK, gin.H{
			"r":log,
		})




        // c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
        //c.String(http.StatusOK, fmt.Sprintf("r:%s,", userpassword))

		
    })
	

	r.Run(":8084")
	
}


