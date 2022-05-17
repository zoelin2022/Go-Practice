package api

import (
	"main/config"

	// influxdb
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// Influxdb 連線
func ConnInflux() influxdb2.Client {
	var port = config.Viper.GetString("databases.influx.port")
	var token = config.Viper.GetString("databases.influx.token")
	client := influxdb2.NewClient(port, token)
	// always close client at the end
	defer client.Close()
	return client
}



