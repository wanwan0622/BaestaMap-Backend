package function

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"flag"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
	"time"
)

func LocalCreateClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	sa := option.WithCredentialsFile("serviceAccount.json")
	conf := &firebase.Config{ProjectID: "baestamap"}
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to Create Client: %v", err)
	}
	// Close client when done with
	return client
}

func remoteCreateClient(ctx context.Context) *firestore.Client {
	projectID := "baestamap"
	flag.StringVar(&projectID, "project", projectID, "The Google Cloud Platform project ID.")
	flag.Parse()

	// [START firestore_setup_client_create]
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

type SearchLocation struct {
	Lat float64
	Lng float64
}

type Location struct {
	Lat        float64
	Lng        float64
	LocationId int32
	Name       string
}

type PostDocs struct {
	HashTagDocsId string
	Location      Location
	Permalink     string
	Timestamp     time.Time
}

func FetchNearPosts(ctx context.Context, client *firestore.Client, location SearchLocation, diff float64) ([]*firestore.DocumentSnapshot, error) {
	iter := client.Collection("posts").Where("location.lat", ">=", location.Lat-diff).Where("location.lat", "<=", location.Lat+diff).Limit(100).Documents(ctx)
	nearPosts := []*firestore.DocumentSnapshot{}
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		lng := doc.Data()["location"].(map[string]interface{})["lng"].(float64)
		if location.Lng-diff <= lng && lng <= location.Lng+diff {
			nearPosts = append(nearPosts, doc)
		}
	}
	return nearPosts, nil
}
