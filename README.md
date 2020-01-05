# Spotify API Wrapper - Golang

![Version: 1.0.0](https://img.shields.io/badge/version-1.0.0-brightgreen.svg)
[![GoDoc](https://godoc.org/github.com/tobiola/Spotify?status.svg)](https://godoc.org/github.com/tobiola/Spotify)

```
go get github.com/tobiola/Spotify
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
result, _ := accout.Search("bangers")
client.Play(result.Tracks.Items[0])
````
