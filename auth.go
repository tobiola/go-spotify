package spotify

import (
	"net/url"
	"math/rand"
	"net/http"
	"time"
	"fmt"
	"encoding/base64"
	"encoding/json"
	"strings"
	"errors"

)

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func generateRandomString(n int) string {
	randomString := make([]rune, n)
	allowedCharacters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	for i := range randomString {
		randomString[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}

	return string(randomString)
}

func GetRedirectLink(callbackUrl string, clientId string, scope string) (string, error) {
	// stateKey := "spotify_auth_state"
	state := generateRandomString(16)

	// http.SetCookie(w, &http.Cookie{Name: stateKey, Value: state, Expires: time.Now().AddDate(0, 0, 1)})

	redirectLink, err := url.Parse("https://accounts.spotify.com/authorize")
	if err != nil {
		return "", nil
	}

	params := url.Values{}
	params.Add("response_type", "code")
	params.Add("client_id", clientId)
	params.Add("scope", scope)
	params.Add("redirect_uri", callbackUrl)
	params.Add("state", state)
	redirectLink.RawQuery = params.Encode()

	return redirectLink.String(), nil
}

func GetTokensFromCallback(requestUrl *url.URL, redirectUri string, clientId string, clientSecret string) (Tokens, error) {
	tokens := Tokens{}
	code := requestUrl.Query().Get("code")
	// state := requestUrl.Query().Get("state")

	/*
	storedStateCookie, err := r.Cookie("spotify_auth_state")
	storedState := storedStateCookie.Value

	if state != storedState {
		return tokens, errors.New("Auth State Mismatch Error")
	}

	http.SetCookie(w, &http.Cookie{Name: "spotify_auth_state", Value: "", Expires: time.Now()})
	*/

	client := &http.Client{Timeout: time.Second * 5}
	form := url.Values{"code": {code}, "redirect_uri": {redirectUri}, "grant_type": {"authorization_code"}}

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
	if err != nil {
		return tokens, err
	}
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))))

	resp, err := client.Do(request)
	if err != nil {
		return tokens, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return tokens, err
	}

	if resp.StatusCode != 200 {
		return tokens, errors.New("Invalid Request")
	}

	/*
		request, err = http.NewRequest("GET", "https://api.spotify.com/v1/me", strings.NewReader(url.Values{}.Encode()))
		if err != nil {
			log.Fatalln(err)
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", result["access_token"]))
	*/

	tokens.AccessToken = fmt.Sprintf("%v", result["access_token"])
	tokens.RefreshToken = fmt.Sprintf("%v", result["refresh_token"])
	return tokens, nil
}

func GetAccessTokenFromRefreshToken(refreshToken string, clientId string, clientSecret string) (string, error) {
	accessToken := ""
	client := &http.Client{Timeout: time.Second * 5}
	form := url.Values{"refresh_token": {refreshToken}, "grant_type": {"authorization_code"}}

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
	if err != nil {
		return accessToken, err
	}
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))))

	resp, err := client.Do(request)
	if err != nil {
		return accessToken, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return accessToken, nil
	}

	accessToken = fmt.Sprintf("%s", result["access_token"])
	return accessToken, nil
}
