package main

import (
	"fmt"
	"context"
	"net/http"
	"github.com/wanwan0622/BaestaMap-Backend"
)


func main() {
	WebServer()
}

func WebServer() {
	http.HandleFunc("/", localGcloudMain)
	fmt.Println("Server started on: http://127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func localGcloudMain(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := function.LocalCreateClient(ctx)
	location := function.SearchLocation{
		Lat: 35.615304235976,
		Lng: 139.7175761816,
	}
	result, err := function.GcloudFirestore(ctx, client, location)
	defer client.Close()
	
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("{'success':false,error:'unexpected error!'}"))
	} else {
		w.Write(result)
	}
}

