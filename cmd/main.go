package main

import (
	"fmt"
	"context"
	// "net/http"
	"log"
	"github.com/wanwan0622/BaestaMap-Backend"
)

func main() {
	postIds := [...] string{"CghEoYyv-4e", "CghbnAEPfU8"}
	function.GetCoordinates(postIds[0])
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

func Crowling() {
	location := "新橋ランチ"
	postIDs, err := function.Crawling(location)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(postIDs)
	ctx := context.Background()
	client := function.LocalCreateClient(ctx)
	postDocs := function.PostDocs{
		PostId: postIDs[0],
		SearchWord: location,
		Location: function.Location{
			Lat: "135.00",
			Lng: "34.3900",
		},
	}
	ok := function.FireStoreInsert(ctx, client, postDocs)
	if !ok {
		log.Fatal("FireStoreInsert failed")
	}
	defer client.Close()
}