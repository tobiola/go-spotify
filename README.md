# Spotify API Wrapper - Golang


[![GoDoc](https://godoc.org/github.com/tobiola/spotify?status.svg)](https://godoc.org/github.com/tobiola/spotify)
[![Build Status](https://travis-ci.com/tobiola/spotify.svg?branch=master)](https://travis-ci.com/tobiola/spotify)

```
go get github.com/tobiola/spotify
```

## Getting started

### Official Spotfy API
[https://developer.spotify.com/documentation/web-api/](https://developer.spotify.com/documentation/web-api/)

### Authentication

````Go
clientId := "[Client ID Here]"
clientSecret := "[Client Secret Here]"
redirectUri := "[URL that should handle the callback]"

func HandleLogin(w http.ResponseWriter, r *http.Request) {
 	scope := "streaming user-read-private user-read-email user-modify-playback-state"

 	redirectLink, _ := spotify.GetRedirectLink(redirectUri, clientId, scope)

 	http.Redirect(w, r, redirectLink, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	client, _ := spotify.GetClientFromCallback(r.URL, redirectUri, clientId, clientSecret)
}
````

### Player

````Go
// Play 
client.Play()
// or
client.PlayTrack(track)
// or
client.PlayTracks(tracks)

// Pause
client.Pause()
````

### Search

````Go
options := spotify.SearchOptions{Query: "hotline bling", Type: "track"}
result, _ := client.Search(options)

client.Play(result.Tracks.Items[0])
````

### Progress
- [x] Authentication
- [ ] Albums
- [ ] Browse
- [ ] Follow
- [ ] Library
- [ ] Personalization
- [ ] Player
- [ ] Playlists
- [x] Search
- [ ] Tracks
- [ ] Users Profile
