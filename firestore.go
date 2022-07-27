package function

import (
	"context"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
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

func fireStoreInsert(ctx context.Context, client *firestore.Client) {
	// Get a Firestore client.
	_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
		"first": "Ada",
		"last":  "Lovelace",
		"born":  1815,
	})
	if err != nil {
		log.Fatalf("Failed adding alovelace: %v", err)
	}
}

func fireStoreRead(ctx context.Context, client *firestore.Client) []*firestore.DocumentSnapshot {
	iter := client.Collection("users").Documents(ctx)
	var results []*firestore.DocumentSnapshot
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}
		fmt.Println(doc.Data()) // TODO: remove
		results = append(results, doc)
	}
	return results
}
