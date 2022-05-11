package config

import "github.com/spf13/viper"

var Viper *viper.Viper


func init(){
	Viper = viper.New() // 實體化viper物件
	Viper.SetConfigName("conn") // 要讀取的檔名
	Viper.SetConfigType("yaml") // 要讀取的附檔名
	Viper.AddConfigPath("./config")  // 要讀取的路徑(在config資料夾內)
	Viper.SetDefault("databases.mysql.port", 8080) // 設定參數預設值
	err := Viper.ReadInConfig() // 讀取設定檔
	if err != nil {
		panic("讀取設定檔出現錯誤，原因為：" + err.Error())
	}

	Viper.WatchConfig()
	
	}
