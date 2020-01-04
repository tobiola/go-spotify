# Spotify API Wrapper - Golang

## Getting started

### Authentication

````Go
func HandleLogin(w http.ResponseWriter, r *http.Request) {
 	scope := "streaming user-read-private user-read-email user-modify-playback-state"
	clientId := "[Client ID Here]"
	clientId := "[Client Secret Here]"
 	redirectUri := "[URL that should handle the callback]"

 	redirectLink, _ := spotify.GetRedirectLink(redirectURI, clientId, scope)

 	http.Redirect(w, r, redirectLink, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	clientId := "[Client ID Here]"
	clientId := "[Client Secret Here]"
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
query := "bangers"
result, _ := accout.Search(accessToken, query)

account.Play(accessToken, result.Tracks.Items[0])
````
