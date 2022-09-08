package function

import (
	"context"
	"log"
	"math"
	"sort"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func LocalCreateClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	sa := option.WithCredentialsFile("serviceAccount.json")
	app, err := firebase.NewApp(ctx, nil, sa)
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
	projectID := "baestamap-api-id"
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

type SearchLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Location struct {
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	LocationId int64   `json:"locationId"`
	Name       string  `json:"name"`
}

type PostDocs struct {
	Location      Location  `json:"location"`
	Permalink     string    `json:"permalink"`
	Timestamp     time.Time `json:"timestamp"`
}

func FetchNearPosts(ctx context.Context, client *firestore.Client, location SearchLocation) ([]*firestore.DocumentSnapshot, error) {
	// 徒歩30分圏内の円に外接する正方形に含まれる投稿を取得
	// 1km あたりの緯度は、だいたい X=0.0090133729745762
	// 徒歩5kmで行ける距離は「不動産の表示に関する公正競争規約施行規則」より400m
	// 徒歩30分圏内の円に外接する正方形の一辺/2の緯度 = X * 2.4(km)
	diff := 0.00901337 * 2.4
	iter := client.Collection("posts").Where("location.lat", ">=", location.Lat-diff).Where("location.lat", "<=", location.Lat+diff).Limit(200).Documents(ctx)
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
	// Locationから近い順に並び替え
	sort.Slice(nearPosts, func(x, y int) bool {
		latX := nearPosts[x].Data()["location"].(map[string]interface{})["lat"].(float64)
		lngX := nearPosts[x].Data()["location"].(map[string]interface{})["lng"].(float64)
		distX := (latX-location.Lat)*(latX-location.Lat)+(lngX-location.Lng)*(lngX-location.Lng)
		
		latY := nearPosts[y].Data()["location"].(map[string]interface{})["lat"].(float64)
		lngY := nearPosts[y].Data()["location"].(map[string]interface{})["lng"].(float64)
		distY := (latY-location.Lat)*(latY-location.Lat)+(lngY-location.Lng)*(lngY-location.Lng)
		return distX < distY
	})
	return nearPosts[:int(math.Min(100, float64(len(nearPosts))))], nil
}

func DSnaps2Obj(dSnaps []*firestore.DocumentSnapshot) []PostDocs {
	obj := []PostDocs{}
	for _, dSnap := range dSnaps {
		obj = append(obj, PostDocs{
			Location: Location{
				Lat:        dSnap.Data()["location"].(map[string]interface{})["lat"].(float64),
				Lng:        dSnap.Data()["location"].(map[string]interface{})["lng"].(float64),
				LocationId: int64(dSnap.Data()["location"].(map[string]interface{})["locationId"].(int64)),
				Name:       dSnap.Data()["location"].(map[string]interface{})["name"].(string),
			},
			Permalink: dSnap.Data()["permalink"].(string),
			Timestamp: dSnap.Data()["timestamp"].(time.Time),
		})
	}
	return obj
}
