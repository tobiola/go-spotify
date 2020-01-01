package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Search for an item
func Search(accessToken string, query string) (SearchResult, error) {
	params := make(url.Values)
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "7")

	url := "https://api.spotify.com/v1/search?" + params.Encode()
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	result := SearchResult{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
