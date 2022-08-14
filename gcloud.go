package function

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"cloud.google.com/go/firestore"
)

func HelloCommand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{'text':'Hello World!'}"))
}

type APIResponse struct {
	Success bool       `json:"success"`
	Posts   []PostDocs `json:"posts"`
}

func GcloudFirestore(ctx context.Context, client *firestore.Client, location SearchLocation) ([]byte, error) {
	result, err := FetchNearPosts(ctx, client, location, 0.1)
	if err != nil {
		log.Fatalf("Failed to get posts: %v", err)
		return nil, err
	}
	posts := DSnaps2Obj(result)
	apiResponse := APIResponse{
		Success: true,
		Posts:   posts,
	}
	json, err := json.Marshal(apiResponse)
	if err != nil {
		log.Fatalf("Failed to parse json: %v", err)
		return nil, err
	}
	return json, nil
}

func GcloudMain(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("{'success':false,error:'Invalid request method.'}"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to parse form: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse request body.'}"))
		return
	}
	location := SearchLocation{}
	err = json.Unmarshal(body, &location)
	if err!= nil {
		log.Fatalf("Failed to parse json: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse json.'}"))
		return
	}
	ctx := context.Background()
	client := remoteCreateClient(ctx)
	result, err := GcloudFirestore(ctx, client, location)
	defer client.Close()

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("{'success':false,error:'unexpected error!'}"))
	} else {
		w.Write(result)
	}
}
