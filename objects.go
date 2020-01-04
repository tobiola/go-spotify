package spotify

// TODO: Add Description
type Album struct {
	AlbumType            string            `json:"album_type"`
	Artists              []ArtistSimple    `json:"artists"`
	AvailableMarkets     []string          `json:"available_markets"`
	Copyrights           []Copyright       `json:"copyrights"`
	ExternalIds          map[string]string `json:"external_ids"`
	ExternalUrls         map[string]string `json:"external_urls"`
	Genres               []string          `json:"genres"`
	Href                 string            `json:"href"`
	Id                   string            `json:"id"`
	Images               []Image           `json:"images"`
	Label                string            `json:"label"`
	Name                 string            `json:"name"`
	Popularity           int               `json:"popularity"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Restrictions         Restrictions      `json:"restrictions"`
	Tracks               []TrackSimple     `json:"tracks"`
	Type                 string            `json:"type"`
	Uri                  string            `json:"uri"`
}

// TODO: Add Description
// TODO: Add Description
// TODO: Add Description
// TODO: Add Description
type AlbumSimple struct {
	AlbumGroup           string            `json:"album_group"`
	AlbumType            string            `json:"album_type"`
	Artists              []ArtistSimple    `json:"artists"`
	AvailableMarkets     []string          `json:"available_markets"`
	ExternalUrls         map[string]string `json:"external_url"`
	Genres               []string          `json:"genres"`
	Href                 string            `json:"href"`
	Id                   string            `json:"id"`
	Images               []Image           `json:"images"`
	Name                 string            `json:"name"`
	ReleaseDate          string            `json:"release_date"`
	ReleaseDatePrecision string            `json:"release_date_precision"`
	Restrictions         Restrictions      `json:"restrictions"`
	Type                 string            `json:"type"`
	Uri                  string            `json:"uri"`
}

// TODO: Add Description
type Artist struct {
	ExternalUrls map[string]string `json:"external_url"`
	Followers    Followers         `json:"followers"`
	Genres       []string          `json:"genres"`
	Href         string            `json:"href"`
	Id           string            `json:"id"`
	Images       []Image           `json:"images"`
	Name         string            `json:"name"`
	Popularity   int               `json:"popularity"`
	Uri          string            `json:"uri"`
	Type         string            `json:"type"`
}

// TODO: Add Description
type ArtistSimple struct {
	ExternalUrls map[string]string `json:"external_url"`
	Href         string            `json:"href"`
	Id           string            `json:"id"`
	Name         string            `json:"name"`
	Uri          string            `json:"uri"`
	Type         string            `json:"type"`
}

// TODO: Add Description
type AudioFeatures struct {
	Acousticness     float32 `json:"acousticness"`
	AnalysisUrl      string  `json:"analysis_url"`
	Danceability     float32 `json:"danceability"`
	DurationMs       int     `json:"duration_ms"`
	Energy           float32 `json:"energy"`
	Id               string  `json:"id"`
	Instrumentalness float32 `json:"instrumentalness"`
	Key              int     `json:"key"`
	Liveness         float32 `json:"liveness"`
	Loudness         float32 `json:"loudness"`
	Mode             int     `json:"mode"`
	Speechiness      float32 `json:"speechiness"`
	Tempo            float32 `json:"tempo"`
	TimeSignature    int     `json:"time_signature"`
	TrackHref        string  `json:"track_href"`
	Type             string  `json:"type"`
	Uri              string  `json:"uri"`
	Valence          float32 `json:"valence"`
}

// TODO: Add Description
type Category struct {
	Href  string  `json:"href"`
	Icons []Image `json:"icons"`
	Id    string  `json:"id"`
	Name  string  `json:"name"`
}

// TODO: Add Description
type Context struct {
	Type         string            `json:"type"`
	Href         string            `json:"href"`
	ExternalUrls map[string]string `json:"external_urls"`
	Uri          string            `json:"uri"`
}

// TODO: Add Description
type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// TODO: Add Description
type Cursor struct {
	After string `json:"after"`
}

// TODO: Add Description
type Disallows string

// TODO: Add Description
type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

// TODO: Add Description
type PlayerError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}

// TODO: Add Description
type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

// TODO: Add Description
type Image struct {
	Height int    `json:"height"`
	Url    string `json:"url"`
	Width  int    `json:"width"`
}

// TODO: Add Description
type Paging struct {
	Href     string      `json:"href"`
	Items    interface{} `json:"items"`
	Limit    int         `json:"limit"`
	Next     string      `json:"next"`
	Offset   int         `json:"offset"`
	Previous string      `json:"previous"`
	Total    int         `json:"total"`
}

// TODO: Add Description
type TrackPaging struct {
	Href     string  `json:"href"`
	Items    []Track `json:"items"`
	Limit    int     `json:"limit"`
	Next     string  `json:"next"`
	Offset   int     `json:"offset"`
	Previous string  `json:"previous"`
	Total    int     `json:"total"`
}

// TODO: Add Description
type AlbumPaging struct {
	Href     string        `json:"href"`
	Items    []AlbumSimple `json:"items"`
	Limit    int           `json:"limit"`
	Next     string        `json:"next"`
	Offset   int           `json:"offset"`
	Previous string        `json:"previous"`
	Total    int           `json:"total"`
}

// TODO: Add Description
type ArtistPaging struct {
	Href     string   `json:"href"`
	Items    []Artist `json:"items"`
	Limit    int      `json:"limit"`
	Next     string   `json:"next"`
	Offset   int      `json:"offset"`
	Previous string   `json:"previous"`
	Total    int      `json:"total"`
}

// TODO: Add Description
type PlaylistPaging struct {
	Href     string           `json:"href"`
	Items    []PlaylistSimple `json:"items"`
	Limit    int              `json:"limit"`
	Next     string           `json:"next"`
	Offset   int              `json:"offset"`
	Previous string           `json:"previous"`
	Total    int              `json:"total"`
}

// TODO: Add Description
type CursorBasedPaging struct {
	Href    string      `json:"href"`
	Items   interface{} `json:"items"`
	Limit   int         `json:"limit"`
	Next    string      `json:"next"`
	Cursors Cursor      `json:"cursors"`
	Total   int         `json:"total"`
}

// TODO: Add Description
type PlayHistory struct {
	Track    TrackSimple `json:"track"`
	PlayedAt Timestamp   `json:"played_at"`
	Context  Context     `json:"context"`
}

// TODO: Add Description
type Playlist struct {
	Collaborative bool              `json:"collaborative"`
	Descriptions  string            `json:"descriptions"`
	ExternalUrls  map[string]string `json:"external_urls"`
	Followers     Followers         `json:"followers"`
	Href          string            `json:"href"`
	Id            string            `json:"id"`
	Images        []Image           `json:"images"`
	Name          string            `json:"name"`
	Owner         User              `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotId    string            `json:"snapshot_id"`
	Tracks        Paging            `json:"tracks"`
	Type          string            `json:"type"`
	Uri           string            `json:"uri"`
}

// TODO: Add Description
type PlaylistSimple struct {
	Collaborative bool              `json:"collaborative"`
	Descriptions  string            `json:"descriptions"`
	ExternalUrls  map[string]string `json:"external_urls"`
	Href          string            `json:"href"`
	Id            string            `json:"id"`
	Images        []Image           `json:"images"`
	Name          string            `json:"name"`
	Owner         User              `json:"owner"`
	Public        bool              `json:"public"`
	SnapshotId    string            `json:"snapshot_id"`
	Tracks        Paging            `json:"tracks"`
	Type          string            `json:"type"`
	Uri           string            `json:"uri"`
}

// TODO: Add Description
type PlaylistTrack struct {
	AddedAt Timestamp `json:"added_at"`
	AddedBy User      `json:"added_by"`
	IsLocal bool      `json:"is_local"`
	Track   Track     `json:"track"`
}

// TODO: Add Description
type Recommendations struct {
	Seeds  []RecommendationSeed `json:"seeds"`
	Tracks []TrackSimple        `json:"tracks"`
}

// TODO: Add Description
type RecommendationSeed struct {
	AfterFilteringSize int    `json:"after_filtering_size"`
	AfterRelinkingSize int    `json:"after_relinking_size"`
	Href               string `json:"href"`
	Id                 string `json:"id"`
	InitialPoolSize    int    `json:"initial_pool_size"`
	Type               string `json:"type"`
}

// TODO: Add Description
type Restrictions interface{}

// TODO: Add Description
type SavedTrack struct {
	AddedAt Timestamp `json:"added_at"`
	Track   Track     `json:"track"`
}

// TODO: Add Description
type SavedAlbum struct {
	AddedAt Timestamp `json:"added_at"`
	Album   Album     `json:"album"`
}

// TODO: Add Description
type Timestamp string

// TODO: Add Description
type Track struct {
	Album            AlbumSimple       `json:"album"`
	Artists          []ArtistSimple    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIds      map[string]string `json:"external_ids"`
	Href             string            `json:"href"`
	Id               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       TrackLink         `json:"linked_from"`
	Restrictions     Restrictions      `json:"restrictions"`
	Name             string            `json:"name"`
	Popularity       int               `json:"popularity"`
	PreviewUrl       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	Uri              string            `json:"uri"`
	IsLocal          bool              `json:"is_local"`
}

// TODO: Add Description
type TrackSimple struct {
	Artists          []ArtistSimple    `json:"artists"`
	AvailableMarkets []string          `json:"available_markets"`
	DiscNumber       int               `json:"disc_number"`
	DurationMs       int               `json:"duration_ms"`
	Explicit         bool              `json:"explicit"`
	ExternalIds      map[string]string `json:"external_ids"`
	Href             string            `json:"href"`
	Id               string            `json:"id"`
	IsPlayable       bool              `json:"is_playable"`
	LinkedFrom       TrackLink         `json:"linked_from"`
	Restrictions     Restrictions      `json:"restrictions"`
	Name             string            `json:"name"`
	PreviewUrl       string            `json:"preview_url"`
	TrackNumber      int               `json:"track_number"`
	Type             string            `json:"type"`
	Uri              string            `json:"uri"`
	IsLocal          bool              `json:"is_local"`
}

// TODO: Add Description
type TrackLink struct {
	ExternalUrls map[string]string `json:"external_urls"`
	Href         string            `json:"href"`
	Id           string            `json:"id"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

// TODO: Add Description
type User struct {
	DisplayName  string            `json:"display_name"`
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    Followers         `json:"followers"`
	Href         string            `json:"href"`
	Id           string            `json:"id"`
	Images       []Image           `json:"images"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

// TODO: Add Description
type UserPrivate struct {
	Country      string            `json:"country"`
	DisplayName  string            `json:"display_name"`
	Email        string            `json:"email"`
	ExternalUrls map[string]string `json:"external_urls"`
	Followers    Followers         `json:"followers"`
	Href         string            `json:"href"`
	Id           string            `json:"id"`
	Images       []Image           `json:"images"`
	Product      string            `json:"product"`
	Type         string            `json:"type"`
	Uri          string            `json:"uri"`
}

// TODO: Add Description
type State struct {
	CurrentTrack Track `json:"currentTrack"`
	Position     int   `json:"position"`
	Duration     int   `json:"duration"`
	Paused       bool  `json:"paused"`
}
