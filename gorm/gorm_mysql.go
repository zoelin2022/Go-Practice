package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	Sex bool
	Age int
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:8889)/ginDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}else{
		fmt.Printf("連接成功")
	}
	// 創建表格
	db.AutoMigrate(&User{}) 

	// **Create**
	db.Create(&User{
		Name:  "amy",
		Sex:   false,
		Age:   12,
	})

	// **Find**
	var user User
	db.First(&user, "name = ?", "amy")

	// **Find**
	var user []User
	db.Where("age < ? and sex = ?", 21, true).Find(&user)
	fmt.Println(user)

	// **Update** 單筆資料、單個值
	db.AutoMigrate(&User{}) 
	db.Where("id = ?", 1).First(&User{}).Update("name", "bob")

	// **Updates** 更新單筆＋多欄，使用 First() + Updates(strcut)
	db.AutoMigrate(&User{}) 
	db.Where("id = ?", 1).First(&User{}).Updates(User{
		Name:  "qq",
		Age:   3,
	})

	// **Updates** 更新單筆＋多欄，使用 First() + Updates(map)
	db.AutoMigrate(&User{}) 
	db.Where("id = ?", 1).First(&User{}).Updates(map[string]interface{}{
		"Name": "222",
		"Age": 188,
	})

	// **Updates** 批量修改，使用 Find() + Updates(map)
	db.AutoMigrate(&User{}) 
	db.Where("id in (?)", []int{1,2}).Find(&[]User{}).Updates(map[string]interface{}{
		"Name": "333",
		"Age": 333,
	})

	// **Delete** 軟刪除
	// 刪除單筆
	db.Delete(&User{},"id = ?", 1) 

	// 刪除多筆
	db.Where("id in (?)", []int{1,2}).Delete(&User{})

	// **Delete** 硬刪除 + Unscoped()
	// 刪除單筆
	db.Unscoped().Delete(&User{},"id = ?", 1) 



	defer db.Close() // 關閉資料庫連線
	
}