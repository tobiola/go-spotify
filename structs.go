// This package contains spotify structs and wrappers for network requests
package spotify

type Artist struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

type Image struct {
	Url    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type Album struct {
	Name   string  `json:"name"`
	Uri    string  `json:"uri"`
	Images []Image `json:"images"`
	Type   string  `json:"type"`
}

type Track struct {
	Album   Album    `json:"album"`
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
	Uri     string   `json:"uri"`
	Type    string   `json:"type"`
}

type User struct {
	Images      []Image `json:"images"`
	DisplayName string  `json:"display_name"`
	Uri         string  `json:"uri"`
}

type Tracks struct {
	Items []Track `json:"items"`
}

type Albums struct {
	Items []Album `json:"items"`
}

type Artists struct {
	Items []Artist `json:"items"`
}

/*
type Playlists struct {
	Items []Playlist `json:"items"`
}
*/

type SearchResult struct {
	Tracks  Tracks  `json:"tracks"`
	Albums  Albums  `json:"albums"`
	Artists Artists `json:"artists"`
	//	Playlists Playlists `json:"playlists"`
}

type State struct {
	CurrentTrack Track `json:"currentTrack"`
	Position     int   `json:"position"`
	Duration     int   `json:"duration"`
	Paused       bool  `json:"paused"`
}
