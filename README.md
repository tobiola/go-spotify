# Spotify API Wrapper - Golang

## Getting started

### Official Spotfy API
[https://developer.spotify.com/documentation/web-api/](https://developer.spotify.com/documentation/web-api/)

### Authentication

````Go
func HandleLogin(w http.ResponseWriter, r *http.Request) {
 	scope := "streaming user-read-private user-read-email user-modify-playback-state"
	clientId := "[Client ID Here]"
	clientSecret := "[Client Secret Here]"
 	redirectUri := "[URL that should handle the callback]"

 	redirectLink, _ := spotify.GetRedirectLink(redirectURI, clientId, scope)

 	http.Redirect(w, r, redirectLink, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	clientId := "[Client ID Here]"
	clientSecret := "[Client Secret Here]"
 	redirectUri := "[URL that should handle the callback]"

	account, _ := spotify.GetAccountFromCallback(r.URL, redirectURI, clientID, clientSecret)
}
````

### Player

````Go
// Play 
account.Play()
// or
account.PlayTrack(track)
// or
account.PlayTracks(tracks)

// Pause
account.Pause()
````

### Search

````Go
result, _ := accout.Search("bangers")
account.Play(result.Tracks.Items[0])
````
