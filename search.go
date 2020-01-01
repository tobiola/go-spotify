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

type Tracks struct {
	Items []Track `json:"items"`
}

type Albums struct {
	Items []AlbumSimple `json:"items"`
}

type Artists struct {
	Items []ArtistSimple `json:"items"`
}

type Playlists struct {
	Items []PlaylistSimple `json:"items"`
}

type SearchResponse struct {
	Tracks  Tracks  `json:"tracks"`
	Albums  Albums  `json:"albums"`
	Artists Artists `json:"artists"`
}

// Search for an item
func Search(accessToken string, query string) (SearchResponse, error) {
	result := SearchResponse{}

	params := make(url.Values)
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "20")

	url := "https://api.spotify.com/v1/search?" + params.Encode()
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	resp, err := client.Do(request)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
