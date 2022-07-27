package main

import (
	// "fmt"
	// "context"
	// "net/http"
	"github.com/wanwan0622/BaestaMap-Backend"
)

func main() {
	function.Crawling()
// 	http.HandleFunc("/", localGcloudMain)
// 	fmt.Println("Server started on: http://127.0.0.1:8080")
// 	http.ListenAndServe("127.0.0.1:8080", nil)
}

// func localGcloudMain(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	client := function.LocalCreateClient(ctx)
// 	result := function.GcloudFirestore(ctx, client)
// 	defer client.Close()
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(result)
// }