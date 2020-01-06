package spotify

import (
	"errors"
	"fmt"
	"encoding/json"

	"github.com/google/go-querystring/query"
)

type SearchResponse struct {
	Tracks    TrackPaging    `json:"tracks"`
	Albums    AlbumPaging    `json:"albums"`
	Artists   ArtistPaging   `json:"artists"`
	Playlists PlaylistPaging `json:"playlists"`
}

type SearchOptions struct {
	// Required.
	// Search query keywords and optional field filters and operators. 
	// Example: "One Dance Drake"
	Query           string `url:"q"`

	// Required.
	// A comma-separated list of item types to search across.
	// Valid types are: album, artist, playlist, and track.
	// Example: "album,track" returns both albums and tracks.
	Type            string `url:"type"`

	// Optional.
	// An ISO 3166-1 alpha-2 country code or the string from_token. 
	// If a country code is specified, only artists, albums, and tracks with content that is playable in that market is returned.
	Market          string `url:"market,omitempty"`

	// Optional.
	// Maximum number of results to return.
	// Default: 20
	// Minimum: 1
	// Maximum: 50
	// The limit is applied within each type, not on the total response. 
	// For example, if the limit value is 3 and the type is artist, album, the response contains 3 artists and 3 albums.
	Limit           int    `url:"limit,omitempty"`

	// Optional.
	// The index of the first result to return. 
	// Default: 0 (the first result). 
	// Maximum offset (including limit): 10,000. 
	// Use with limit to get the next page of search results.
	Offset          int    `url:"offset,omitempty"`

	// Optional. 
	// Possible values: audio 
	// If include_external=audio is specified the response will include any relevant audio content that is hosted externally. 
	// By default external content is filtered out from responses.
	IncludeExternal string `url:"include_external,omitempty"`
}

// Search for an item
func (c *Client) Search(options SearchOptions) (SearchResponse, error) {
	if options.Query == "" {
		return SearchResponse{}, errors.New("Query must not be empty")
	}

	v, err := query.Values(options)
	if err != nil {
		return SearchResponse{}, err
	}

	url := fmt.Sprintf("https://api.spotify.com/v1/search?%s", v.Encode())
	bytes, err := c.get(url, nil)
	if err != nil {
		return SearchResponse{}, err
	}

	result := SearchResponse{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return SearchResponse{}, err
	}

	return result, nil
}
