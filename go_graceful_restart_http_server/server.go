package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world"))
    })

    PORT := "5000"

    fmt.Println("Server running on port:", PORT)

    http.ListenAndServe(":" + PORT, nil)

}