package model

type Response struct {
	Msg  string
	Data interface{}
}

type DiskData struct {
	Host   string
	Divice string
	Rate   float64
	Total  float64
	Used   float64
	Nday   float64

}

type Data struct {
	Time   string
	First  float64
	Last   float64
	Max	   float64
	Min	   float64
}

//elastic
// type Employee struct {
// 	Time    string            `json:"時間"`
// 	IP      string            `json:"IP"`
// 	Message string            `json:"RawData"`
// }
