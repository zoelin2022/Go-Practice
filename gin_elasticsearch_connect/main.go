package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

type Data struct {
	Time    string            `json:"時間"`
	IP      string            `json:"IP"`
	Message string            `json:"RawData"`
	//Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
 }

var (
	client *elastic.Client
	url    = "http://192.168.1.104:9200"
	ctx    = context.Background()
 )

 func init() {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(url))
	//測試es連接狀態
	_, _, err := client.Ping(url).Do(ctx)
	if err != nil {
	   log.Println("連接es失敗", err)
	}
	log.Println("連接es成功")
	//查看es當前版本
	version, err := client.ElasticsearchVersion(url)
	if err != nil {
	   log.Println("查詢es版本錯誤", err)
	}
	log.Println("Elasticsearch version: ", version)
 }

 // 搜尋ID獲取數據
 func getData() {
	res, _ := client.Get().Index("accesslog").Id("ZF1RZoABWSdVd1arkef2").Do(ctx)
	if res.Found {
	   fmt.Printf("Got document %s in version %d from index %s, type %s\n", res.Id, res.Version, res.Index, res.Type)
	   var data Data
	   _ = json.Unmarshal(*res.Source, &data)
	   fmt.Println("數據内容: ",data)
	}
 }

 func main() {
	getData()

 }
 