package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// インスタグラムのAPIを用いてタグ検索から投稿IDを取得する
func Scraping(location string, clientId string) ([]byte, error) {
	tag := url.QueryEscape(location)
	url := fmt.Sprintf("https://i.instagram.com/api/v1/tags/logged_out_web_info/?tag_name=%s", tag)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ja")
	req.Header.Set("Referer", "https://www.instagram.com/")
	req.Header.Set("X-Ig-App-Id", clientId)
	client := new(http.Client)
	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	byteArray, _ := ioutil.ReadAll(res.Body)
	return byteArray, nil
}

func bytes2Json(bytes []byte) interface{} {
	var jsonObj interface{}
	_ = json.Unmarshal(bytes, &jsonObj)
	return jsonObj
}

func getPostIDs(jsonObj interface{}) []string {
	edges := jsonObj.(map[string]interface{})["data"].(map[string]interface{})["hashtag"].(map[string]interface{})["edge_hashtag_to_media"].(map[string]interface{})["edges"]
	edgeNum := len(edges.([]interface{}))
	getShortCode := func(idx int) string {
		return edges.([]interface{})[idx].(map[string]interface{})["node"].(map[string]interface{})["shortcode"].(string)
	}
	shortCodes := make([]string, edgeNum)
	for i := 0; i < edgeNum; i++ {
		shortCodes[i] = getShortCode(i)
	}
	return shortCodes
}

type Client struct {
	ClientId string `json:"client_id"`
}

func Crawling() {
	var client Client
	raw, err := ioutil.ReadFile("./client.json")
	if err != nil {
		log.Fatal(err)
		return
	}
	json.Unmarshal(raw, &client)
	location := "新橋ランチ"
	bytes, err := Scraping(location, client.ClientId)
	if err != nil {
		log.Fatal(err)
	}
	jsonObj := bytes2Json(bytes)
	postIDs := getPostIDs(jsonObj)

	fmt.Println(postIDs)
}
