package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// es 套件
	"context"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/olivere/elastic/v7"
)

//查找

// @Summary      查詢log
// @Description  查詢log
// @Tags         ElasticSearch
// @Accept       json
// @Produce      json
// @Param ip query string true "207.46.13.9" Format(ip)
// @Success      200  {object} Response
// @Failure      400  {object} Response
// @Failure      404  {object} Response
// @Failure      500  {object} Response
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
//初始化
func init() {
    errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
    var err error
    client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
    if err != nil {
        panic(err)
    }
    info, code, err := client.Ping(host).Do(context.Background())
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

    esversion, err := client.ElasticsearchVersion(host)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Elasticsearch version %s\n", esversion)

}

//打印查询到的Employee
func printEmployee(res *elastic.SearchResult, err error) {
    if err != nil {
        print(err.Error())
        return
    }
    var typ Employee
    for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
        t := item.(Employee)
        fmt.Printf("%#v\n", t)
    }
}