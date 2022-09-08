package function

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"googlemaps.github.io/maps"
)

func Query2Coordinate(query string) (SearchLocation, error) {
	// TODO: これでいいんか...??
	bytes, err := ioutil.ReadFile("./serverless_function_source_code/api_key.txt")
	if err != nil {
		log.Fatalf("Failed to read api_key.txt: %v", err)
		return SearchLocation{}, fmt.Errorf("internal server error")
	}
	API_KEY := string(bytes)
	c, err := maps.NewClient(maps.WithAPIKey(API_KEY))
	if err != nil {
		return SearchLocation{}, fmt.Errorf("failed to create client")
	}
	r := &maps.TextSearchRequest{
		Query: query,
	}

	res, err := c.TextSearch(context.Background(), r)
	if err != nil {
		return SearchLocation{}, fmt.Errorf("failed to Search")
	}

	if len(res.Results) == 0 {
		return SearchLocation{}, fmt.Errorf("we did a search and found no such place")
	}
	return SearchLocation{
		Lat: res.Results[0].Geometry.Location.Lat,
		Lng: res.Results[0].Geometry.Location.Lng,
	}, nil
}

type SearchQuery struct {
	Query string `json:"query"`
}

// TODO: refactor
func GetPostFromQuery(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// main request
	// validation
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		w.Write([]byte("{'success':false,error:'Invalid request method.'}"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to parse form: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse request body.'}"))
		return
	}
	query := SearchQuery{}
	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Fatalf("Failed to parse json: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse json.'}"))
		return
	}

	location, err := Query2Coordinate(query.Query)
	if err != nil {
		log.Fatalf("Failed to Convert Query to Coordinate: %v", err)
		w.Write([]byte(fmt.Sprintf("{'success':false,error:'%s'}", err.Error())))
		return
	}

	// main program
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

type LocationAPIResponse struct {
	Success  bool           `json:"success"`
	Location SearchLocation `json:"location"`
}

func GetLocationFromQuery(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for the preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// main request
	// validation
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		w.Write([]byte("{'success':false,error:'Invalid request method.'}"))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Failed to parse form: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse request body.'}"))
		return
	}
	query := SearchQuery{}
	err = json.Unmarshal(body, &query)
	if err != nil {
		log.Fatalf("Failed to parse json: %v", err)
		w.Write([]byte("{'success':false,error:'Failed to parse json.'}"))
		return
	}

	location, err := Query2Coordinate(query.Query)
	if err != nil {
		log.Fatalf("Failed to Convert Query to Coordinate: %v", err)
		w.Write([]byte(fmt.Sprintf("{'success':false,error:'%s'}", err.Error())))
		return
	}

	apiResponse := LocationAPIResponse{
		Success:  true,
		Location: location,
	}
	result, err := json.Marshal(apiResponse)
	if err != nil {
		log.Fatalf("Failed to parse json: %v", err)
		w.Write([]byte(fmt.Sprintf("{'success':false,error:'%s'}", err.Error())))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.Write([]byte("{'success':false,error:'unexpected error!'}"))
	} else {
		w.Write(result)
	}
}
