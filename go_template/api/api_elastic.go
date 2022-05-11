package api

import (
	_ "main/model"
	"net/http"

	"github.com/gin-gonic/gin"

	// es 套件
	"context"

	"github.com/olivere/elastic/v7"
)


var client *elastic.Client

//查找

// @Summary      查詢log
// @Description  查詢log
// @Tags         ElasticSearch
// @Accept       json
// @Produce      json
// @Param ip query string true "207.46.13.9" Format(ip)
// @Success      200  {object} model.Response
// @Failure      404  {object} model.Response
// @Router       /es [get]
func ESquery(c *gin.Context) {
    var res *elastic.SearchResult
    var err error
	ip := c.Query("ip")
    //字段相等 IP
    q := elastic.NewQueryStringQuery(ip)
    res, err = client.Search("accesslog").Type("_doc").Query(q).Do(context.Background())
    if err != nil {
        println(err.Error())
    }
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data": res,
	})
}