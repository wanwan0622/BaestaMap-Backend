package function

import (
	"context"
	"net/http"
)

func HelloCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'text':'Hello World!'}"))
}

func GcloudMain(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()
	fireStoreInsert(ctx, client)
	fireStoreRead(ctx, client)
}
