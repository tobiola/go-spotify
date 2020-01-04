# Spotify API Wrapper - Golang

## Getting started

### Authentication

````Go
func HandleLogin(w http.ResponseWriter, r *http.Request) {
 	scope := "streaming user-read-private user-read-email user-modify-playback-state"
 	clientId = "[Your Client ID Here]"
 	redirectUri = "[Your url that should handle the callback]"

 	redirectLink, err := spotify.GetRedirectLink(s.RedirectURI, clientId, scope)
 	if err != nil {
   		log.Fatalln(err)
 	}

 	http.Redirect(w, r, redirectLink, http.StatusSeeOther)
}

func HandleCallback(w http.ResponseWriter, r *http.Request) {
	tokens, err := spotify.GetTokensFromCallback(r.URL, s.RedirectURI, s.ClientID, s.ClientSecret)
	if err != nil {
		log.Fatalln(err)
	}

	accessToken := tokens.AccessToken
	refreshToken := tokens.RefreshToken
}
````

### Player

````Go
// Play 
spotify.Play(accessToken)

// Pause
spotify.Pause(accessToken)
````

### Search

````Go
query := "bangers"
result, err := spotify.Search(accessToken, query)
if err != nil {
	log.Fatalln(err)
}
	
spotify.Play(accessToken, result.Tracks.Items[0])
````
