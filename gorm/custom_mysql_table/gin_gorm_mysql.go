package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	// 自定義pk、欄位名稱、type
	Name string `gorm:"primary_key;column:user_name;type:varchar(100);"`
}
// 自定義 TableName
func (u User)TableName()string{
	return "test_users"
}
func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/ginDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open("mysql", dsn)

	// 提交資料庫的修改
	db.AutoMigrate(&User{})
	defer db.Close() // 關閉資料庫連線
	
}