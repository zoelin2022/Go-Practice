package api

import (
	"context"
	"fmt"
	"main/config"
	model "main/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// influxdb
	//influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

//var client influxdb2.Client

//var bucket = config.Viper.GetString("databases.influx.bucket")
var org    = config.Viper.GetString("databases.influx.org")
//var bucket    = config.Viper.GetString("databases.influx.bucket")


// @Summary      取得 Disk 90% 滿載天數
// @Description  取得 Disk 90% 滿載天數
// @Tags         Influx
// @Accept       json
// @Produce      json
// @Param bucket query string true "bucket" Format(bucket)
// @Success      200  {object} model.Response
// @Failure      404  {object} model.Response
// @Router       /influx [get]
// 計算還有幾天到達90%滿載
func Calculate(c *gin.Context) {
	bucket := c.Request.URL.Query().Get("bucket")
	var client = ConnInflux()
	var n float64
	var disk model.DiskData
	var disk_arr []model.DiskData
	queryAPI := client.QueryAPI(org)

	// 查詢 sda3 disk 過去7天 平均每天使用量
	query := fmt.Sprintf( `
	import "experimental/aggregate"
	from(bucket: "%s")
	  |> range(start: -1d)
	  |> filter(fn: (r) => r["_measurement"] == "disk")
	  |> filter(fn: (r) => r["_field"] == "used")
	  |> map(fn: (r) => ({ r with _value: float(v:r._value) / 1024.0 / 1024.0 / 1024.0}))
	  |> aggregate.rate(every: 1d, unit: 1h, groupColumns: ["host", "device"])
      |> keep(columns: ["host", "device","_value"])
	  |> first()
	`,bucket)
	result, err := queryAPI.Query(context.Background(), query)
	if err == nil {
		for result.Next() {

			disk.Host = fmt.Sprintf("%v", result.Record().ValueByKey("host"))
			disk.Divice = fmt.Sprintf("%v", result.Record().ValueByKey("device"))
			r := fmt.Sprintf("%v",result.Record().Value())
			n , _ = strconv.ParseFloat(r,64) // string 轉 float
			disk.Rate = n
			disk_arr = append(disk_arr, disk)

		}
	} else {
		panic(err)
	}

	// 查詢  disk used
	// 查詢  disk total x90%
	query1 := fmt.Sprintf( `
	from(bucket: "%s")
	|> range(start: -5m)
	|> filter(fn: (r) => r["_measurement"] == "disk")
  	|> filter(fn: (r) => r["_field"] == "total" or r["_field"] == "used")
	|> last()
  	|> pivot(rowKey:["_time"], columnKey: ["_field"], valueColumn: "_value")
	|> map(fn: (r) => ({ r with total: float(v:r.total) / 1024.0 / 1024.0 / 1024.0 * 0.9}))
  	|> map(fn: (r) => ({ r with used: float(v:r.used) / 1024.0 / 1024.0 / 1024.0}))
	|> keep(columns: ["host", "device","total","used"])
	`,bucket)
	
	result1, err := queryAPI.Query(context.Background(), query1)
	if err == nil {
		for result1.Next() {

			total := fmt.Sprintf("%v",result1.Record().ValueByKey("total")) // string
			total1 , _ := strconv.ParseFloat(total,64) // 轉 float

			used := fmt.Sprintf("%v",result1.Record().ValueByKey("used")) // string
			used1 , _ := strconv.ParseFloat(used,64) // 轉 float

			for idx ,item := range disk_arr {

				if item.Divice == result1.Record().ValueByKey("device") {

					// 計算 剩下 n 天 
					days := ((total1 - used1) / item.Rate)

					disk_arr[idx].Total = total1
					disk_arr[idx].Used = used1

					if (days < 0) {
						disk_arr[idx].Nday = 0
					} else if (item.Rate == 0) {
						disk_arr[idx].Nday = 999

					} else {
						disk_arr[idx].Nday = days
					}
					
				}
				
				
			}
		
		}
	} else {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":    disk_arr,
	})

}




// @Summary      取得 Disk 圖表資料
// @Description  取得 Disk 圖表資料
// @Tags         Influx
// @Accept       json
// @Produce      json
// 參數名 參數類型 參數的資料類型 是否必須 註解
// @Param bucket query string true "bucket" Format(bucket)
// @Param device query string true "device" Format(device)
// @Success      200  {object} model.Response
// @Failure      404  {object} model.Response
// @Router       /influx/2 [get]
func Query_box(c *gin.Context) {
	bucket := c.Request.URL.Query().Get("bucket")
	device := c.Request.URL.Query().Get("device")
	var data = model.Data{}
	var arr []model.Data
	var client = ConnInflux()
	query := fmt.Sprintf( `
	d1 = from(bucket: "%s")
	|> range(start: -1d)
	|> filter(fn: (r) => r["_measurement"] == "disk")
	|> filter(fn: (r) => r["_field"] == "used")
	|> filter(fn: (r) => r["device"] == "%s")
	|> aggregateWindow(every: 5m, fn: first, createEmpty: false)
	|> map(fn: (r) => ({ r with _value: float(v:r._value) / 1024.0 / 1024.0 / 1024.0}))
	|> keep(columns: ["_time", "_value"])
  
	d2 = from(bucket: "%s")
	|> range(start: -1d)
	|> filter(fn: (r) => r["_measurement"] == "disk")
	|> filter(fn: (r) => r["_field"] == "used")
	|> filter(fn: (r) => r["device"] == "%s")
	|> aggregateWindow(every: 5m, fn: last, createEmpty: false)
	|> map(fn: (r) => ({ r with _value: float(v:r._value) / 1024.0 / 1024.0 / 1024.0}))
	|> keep(columns: ["_time", "_value"])
  
  	d3 = from(bucket: "%s")
	|> range(start: -1d)
	|> filter(fn: (r) => r["_measurement"] == "disk")
	|> filter(fn: (r) => r["_field"] == "used")
	|> filter(fn: (r) => r["device"] == "%s")
	|> aggregateWindow(every: 5m, fn: min, createEmpty: false)
	|> map(fn: (r) => ({ r with _value: float(v:r._value) / 1024.0 / 1024.0 / 1024.0}))
	|> keep(columns: ["_time", "_value"])
  
  	d4 = from(bucket: "%s")
	|> range(start: -1d)
	|> filter(fn: (r) => r["_measurement"] == "disk")
	|> filter(fn: (r) => r["_field"] == "used")
	|> filter(fn: (r) => r["device"] == "%s")
	|> aggregateWindow(every: 5m, fn: max, createEmpty: false)
	|> map(fn: (r) => ({ r with _value: float(v:r._value) / 1024.0 / 1024.0 / 1024.0}))
	|> keep(columns: ["_time", "_value"])
  
  	r1 = join(tables: {first: d1, last: d2}, on: ["_time"], method: "inner")
  	r2 = join(tables: {min: d3, max: d4}, on: ["_time"], method: "inner")
  	join(tables: {r1: r1, r2: r2}, on: ["_time"], method: "inner")
	`,bucket,device,bucket,device,bucket,device,bucket,device)

	// 過去一天，平均每小時增加的用量
	queryAPI := client.QueryAPI(org)
	result, err := queryAPI.Query(context.Background(), query)

	if err == nil {

		for result.Next() {
			
			//device := fmt.Sprintf("%v",result.Record().ValueByKey("device"))

			first := fmt.Sprintf("%v", result.Record().ValueByKey("_value_first")) 
			first2 , _ := strconv.ParseFloat(first,64) // string 轉 float
			data.First = first2

			last := fmt.Sprintf("%v", result.Record().ValueByKey("_value_last")) 
			last2 , _ := strconv.ParseFloat(last,64) // string 轉 float
			data.Last = last2

			min := fmt.Sprintf("%v", result.Record().ValueByKey("_value_min")) 
			min2 , _ := strconv.ParseFloat(min,64) // string 轉 float
			data.Min = min2
			
			max := fmt.Sprintf("%v", result.Record().ValueByKey("_value_max")) 
			max2 , _ := strconv.ParseFloat(max,64) // string 轉 float
			data.Max = max2

			time := fmt.Sprintf("%v", result.Record().Time()) 
			data.Time = time
			arr = append(arr, data)
		}
		//fmt.Printf("data: %v\n", arr)
	} else {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"messege": "successfuly",
		"data":    arr,
	})
}
