package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", GcloudMain)
	fmt.Println("Server started on: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
