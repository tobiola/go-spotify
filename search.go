package spotify

import (
	"encoding/json"
	"net/url"
)

/*
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
*/

type SearchResponse struct {
	Tracks    TrackPaging    `json:"tracks"`
	Albums    AlbumPaging    `json:"albums"`
	Artists   ArtistPaging   `json:"artists"`
	Playlists PlaylistPaging `json:"playlists"`
}

// Search for an item
func (c *Client) Search(query string) (SearchResponse, error) {
	result := SearchResponse{}

	params := make(url.Values)
	params.Add("q", query)
	params.Add("type", "track")
	params.Add("limit", "20")

	url := "https://api.spotify.com/v1/search?" + params.Encode()

	bytes, err := c.get(url, nil)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
