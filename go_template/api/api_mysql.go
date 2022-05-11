package api

import (
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var db *gorm.DB
var err error

// @Summary      查詢log
// @Description  查詢log
// @Tags         MySQL
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object} model.Response
// @Failure      404  {object} model.Response
// @Router       /sql/{id} [get]
func Show(c *gin.Context) {
	log := getById(c)
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data": log,
	})
}

// @Summary      查詢time
// @Description  查詢time
// @Tags         MySQL_SearchTime
// @Accept       json
// @Produce      json
// @Param start query string true "2017-12-01 04:09" Format(time)
// @Param end query string true "2017-12-01 04:10"  Format(time)
// @Success      200  {object} model.Response
// @Router       /sql [get]
func Time(c *gin.Context) {
	start := c.Request.URL.Query().Get("start")
	end := c.Request.URL.Query().Get("end")
	var log []model.Log
	// 取得 user 結構體 (繼承相關屬性)
	_ = c.ShouldBind(&log)
	db.Where("time BETWEEN ? AND ?", start, end).Find(&log)
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data": log,
	})
}


// @Summary      新增log
// @Description  新增log
// @Tags         MySQL
// @Accept       json
// @Produce      json
// 參數名 參數類型 參數的資料類型 是否必須 註解
// @Param ip query string true "0.0.0.0" Format(ip)
// @Param data query string true "GET / HTTP/1.1"  Format(data)
// @Param status query int true "200"  Format(status)
// @Success      200  {object} model.Response
// @Failure      404  {object} model.Response
// @Router       /sql [post]
func Create(c *gin.Context) {
	ip := c.Request.URL.Query().Get("ip")
	data := c.Request.URL.Query().Get("data")
	status := c.Request.URL.Query().Get("status")
	astatus, _ := strconv.Atoi(status)
	log := model.Log{
		Time:   time.Now(),
		IP:     ip,
		Data:   data,
		Status: astatus,
	}
	db.Create(&log)
	c.JSON(http.StatusOK, model.Response{
		Msg:  "successfuly",
		Data: log,
	})
}



// @Summary      刪除log
// @Description  刪除log
// @Tags         MySQL
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Success      200  {object} model.Response
// @Router       /sql/{id} [delete]
func Delete(c *gin.Context) {
	log := getById(c)
	if log.Id == 0 {
		return
	}
	// **Delete** 硬刪除
	//db.Unscoped().Delete(&log)

	// **Delete** 軟刪除
	db.Delete(&model.Log{},"id = ?", log.Id) 
	
	c.JSON(http.StatusOK, model.Response{
		Msg:  "deleted successfuly",
		Data: "",
	})

}



// @Summary      更新log
// @Description  更新log
// @Tags         MySQL
// @Accept       json
// @Produce      json
// @Param id path int true "id"
// @Param ip query string true "0.0.0.0" Format(ip)
// @Param data query string true "GET / HTTP/1.1"  Format(data)
// @Param status query int true "200"  Format(status)
// @Success      200  {object} model.Response
// @Router       /sql/{id} [patch]
func Update(c *gin.Context) {
	oldlog := getById(c)
	ip := c.Request.URL.Query().Get("ip")
	data := c.Request.URL.Query().Get("data")
	status := c.Request.URL.Query().Get("status")
	astatus, _ := strconv.Atoi(status)
	newlog := model.Log{
		Time:   time.Now(),
		IP:     ip,
		Data:   data,
		Status: astatus,
	}
	oldlog.IP = newlog.IP
	oldlog.Time = newlog.Time
	oldlog.Data = newlog.Data
	oldlog.Status = newlog.Status
	db.Save(&oldlog)

	c.JSON(http.StatusOK, gin.H{
		"messege": "Log has been updated",
		"data":  oldlog,
	})

}

// 根据id查询
func getById(c *gin.Context) model.Log {
	id := c.Param("id")
	var log model.Log
	db.First(&log, id)
	return log
}