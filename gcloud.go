package function

import (
	"context"
	"fmt"
	"net/http"
	"encoding/json"

	"cloud.google.com/go/firestore"
)

func HelloCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'text':'Hello World!'}"))
}

func GcloudFirestore(ctx context.Context, client *firestore.Client) []byte {
	// fireStoreInsert(ctx, client)
	result := fireStoreRead(ctx, client)
	jsonRes, err := json.Marshal(result)
    if err != nil {
        fmt.Println("JSON marshal error: ", err)
    }
	return jsonRes
}

func GcloudMain(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	client := remoteCreateClient(ctx)
	result := GcloudFirestore(ctx, client)
	defer client.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
