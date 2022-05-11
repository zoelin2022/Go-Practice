package api

import (
	"context"
	"fmt"
	"main/config"
	"main/model"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// es 套件
	"log"
	"os"

	"github.com/olivere/elastic/v7"
)

// Redis 連線
func Redis_client()(*redis.Client,context.Context){
	// viper 讀取 redis 連線設定
	var port = config.Viper.GetString("databases.redis.port")
	rdb = redis.NewClient(&redis.Options{
		Addr: port,
	})
	ctx = context.Background()
	return rdb,ctx
}

// Mysql 連線
func Mysql_client()(*gorm.DB){
	//making an instance of the type DB from the gorm package

	// viper 讀取  Mysql 連線設定
	username := config.Viper.Get("databases.mysql.username")
	password := config.Viper.Get("databases.mysql.password")
	port := config.Viper.Get("databases.mysql.port")
	dbname := config.Viper.Get("databases.mysql.dbname")
	// dsn := "root:root@tcp(localhost:8889)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",username,password,port,dbname)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//handle the error comes from the connection with DB
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&model.Log{})
	return db
}

//var host = "http://140.137.219.56:9200/"
//初始化
func Elastic_client()(*elastic.Client) {
    // viper 讀取 redis 連線設定
    var host = config.Viper.GetString("databases.elastic.port")
    errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
    var err error
    client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
    if err != nil {
        panic(err)
    }
    client.Ping(host).Do(context.Background())
    if err != nil {
        panic(err)
    }
    client.ElasticsearchVersion(host)
    if err != nil {
        panic(err)
    }
    return client

}