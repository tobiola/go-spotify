package spotify

import (
	"fmt"
	"time"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func (a Account) GetAlbum(albumId string) (Album, error) {
	album := Album{}

	/*
	params := make(url.Values)
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "20")
	*/

	url := "https://api.spotify.com/v1/albums/"+ albumId
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return album, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	resp, err := client.Do(request)
	if err != nil {
		return album, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &album)
	if err != nil {
		return album, err
	}

	return album, nil
}

func (a Account) GetAlbumTracks(albumId string) (TrackPaging, error) {
	tracks := TrackPaging{}

	/*
	params := make(url.Values)
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "20")
	*/

	url := "https://api.spotify.com/v1/albums/" + albumId + "/tracks"
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return tracks, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	resp, err := client.Do(request)
	if err != nil {
		return tracks, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bytes, &tracks)
	if err != nil {
		return tracks, err
	}

	return tracks, nil
}