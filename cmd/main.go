package main

import (
	"fmt"
	"net/http"
	"github.com/wanwan0622/BaestaMap-Backend"
)

func main() {
	http.HandleFunc("/", function.GcloudMain)
	fmt.Println("Server started on: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
