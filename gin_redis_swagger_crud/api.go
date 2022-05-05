package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

// @Summary      Read All Data
// @Description  獲取所有資料
// @Tags         redis_all
// @Accept       json
// @Produce      json
// @Param id path int false "id"
// @Success      200  {object} Response
// @Failure      404  {object} Response
// @Router       /redis/all [get]
func GetAllHash(c *gin.Context) {
	keys, err := rdb.Keys(ctx, "*").Result()
	if err != nil {
		panic(err)
	}
	fmt.Printf("keys: %v\n", keys)
	var x []map[string]string
	for i := range keys {
		all := rdb.HGetAll(ctx, keys[i])
		c := all.Val()
		c["id"] = keys[i]
		x = append(x, c)
	}
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":    x,
	})
}

// @Summary      Create
// @Description  新增
// @Tags         redis
// @Accept       json
// @Produce      json
// @Param id query string true "id" Format(id)
// @Param name query string true "name" Format(name)
// @Param age query string true "18"  Format(age)
// @Success      200  {object} Response
// @Failure      404  {object} Response
// @Router       /redis [post]
func CreateHash(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	name := c.Request.URL.Query().Get("name")
	age := c.Request.URL.Query().Get("age")
	rdb.HSet(ctx, id, map[string]string{"name": name, "age": age})
	all := rdb.HGetAll(ctx, id)
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":    all.Val(),
	})
}

// @Summary      Read
// @Description  查詢
// @Tags         redis
// @Accept       json
// @Produce      json
// @Param id query string true "id" Format(id)
// @Success      200  {object} Response
// @Failure      404  {object} Response
// @Router       /redis [get]
func GetHash(c *gin.Context) {
	id := c.Query("id")
	all := rdb.HGetAll(ctx, id)
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":    all.Val(),
	})
}

// @Summary      Update
// @Description  修改
// @Tags         redis
// @Accept       json
// @Produce      json
// @Param id query string true "id" Format(id)
// @Param name query string true "name" Format(name)
// @Param age query string true "18"  Format(age)
// @Success      200  {object} Response
// @Failure      404  {object} Response
// @Router       /redis [patch]
func UpdateHash(c *gin.Context) {
	id := c.Request.URL.Query().Get("id")
	name := c.Request.URL.Query().Get("name")
	age := c.Request.URL.Query().Get("age")
	// 初始化hash數據的多個字段值
	data := make(map[string]interface{})
	data["name"] = name
	data["age"] = age

	// 一次性保存多個hash字段值
	err := rdb.HMSet(ctx,id,data).Err()
	if err != nil {
		panic(err)
	}
	all := rdb.HGetAll(ctx, id)
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":  all.Val(),
	})
}

// @Summary      Delete
// @Description  刪除
// @Tags         redis
// @Accept       json
// @Produce      json
// @Param id query string true "id" Format(id)
// @Success      200  {object} Response
// @Failure      404  {object} Response
// @Router       /redis [delete]
func DeleteHash(c *gin.Context) {
	id := c.Query("id")
	rdb.Del(ctx, id)
	
	c.JSON(http.StatusOK, gin.H{
		"messege": "deleted successfuly",
		"data":    "",
	})
}
