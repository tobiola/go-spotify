# Golang Spotify API Wrapper

## Getting started

`
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
`
