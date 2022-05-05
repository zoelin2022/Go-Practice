package main



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