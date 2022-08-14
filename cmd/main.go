package main

import (
	// "fmt"
	"context"
	// "net/http"
	// "log"
	"github.com/wanwan0622/BaestaMap-Backend"
)


func main() {
	location := function.SearchLocation{
		Lat: 35.615304235976,
		Lng: 139.7175761816,
	}
	ctx := context.Background()
	client := function.LocalCreateClient(ctx)
	function.FetchNearPosts(ctx, client, location, 0.1)
}

// func WebServer() {
// 	http.HandleFunc("/", localGcloudMain)
// 	fmt.Println("Server started on: http://127.0.0.1:8080")
// 	http.ListenAndServe("127.0.0.1:8080", nil)
// }

// func localGcloudMain(w http.ResponseWriter, r *http.Request) {
// 	ctx := context.Background()
// 	client := function.LocalCreateClient(ctx)
// 	result := function.GcloudFirestore(ctx, client)
// 	defer client.Close()
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(result)
// }
