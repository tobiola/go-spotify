package spotify

import (
	"encoding/json"
)

func (c *Client) GetAlbum(albumId string) (Album, error) {
	album := Album{}

	/*
		params := make(url.Values)
		params.Add("q", query)
		params.Add("type", "track")
		params.Add("limit", "20")
	*/

	url := "https://api.spotify.com/v1/albums/" + albumId
	bytes, err := c.get(url, nil)
	if err != nil {
		return album, err
	}

	err = json.Unmarshal(bytes, &album)
	if err != nil {
		return album, err
	}

	return album, nil
}

func (c *Client) GetAlbumTracks(albumId string) (TrackPaging, error) {
	tracks := TrackPaging{}

	/*
		params := make(url.Values)
		params.Add("q", query)
		params.Add("type", "track")
		params.Add("limit", "20")
	*/

	url := "https://api.spotify.com/v1/albums/" + albumId + "/tracks"
	bytes, err := c.get(url, nil)
	if err != nil {
		return tracks, err
	}

	err = json.Unmarshal(bytes, &tracks)
	if err != nil {
		return tracks, err
	}

	return tracks, nil
}
