package model

import (
	"time"

	"gorm.io/gorm"
)

type Response struct {
	Msg  string
	Data interface{}
}


//elastic
type Employee struct {
	Time    string            `json:"時間"`
	IP      string            `json:"IP"`
	Message string            `json:"RawData"`
}

type Log struct {
	// 自定義pk、欄位名稱、type
	gorm.Model
	Id     int64
	Time   time.Time
	IP     string
	Data   string
	Status int
}

// 不顯示  gorm.Model
type Log2 struct {
	// 自定義pk、欄位名稱、type
	//gorm.Model
	Id     int64
	Time   time.Time
	IP     string
	Data   string
	Status int
}

// post 跟 update 預設輸入用
type Log1 struct {
	Time   time.Time
	IP     string
	Data   string
	Status int64
}