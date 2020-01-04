package spotify

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

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

func GetAccountFromCallback(requestUrl *url.URL, redirectUri string, clientId string, clientSecret string) (Account, error) {
	account := Account{}
	code := requestUrl.Query().Get("code")
	// state := requestUrl.Query().Get("state")

	/*
		storedStateCookie, err := r.Cookie("spotify_auth_state")
		storedState := storedStateCookie.Value

		if state != storedState {
			return account, errors.New("Auth State Mismatch Error")
		}

		http.SetCookie(w, &http.Cookie{Name: "spotify_auth_state", Value: "", Expires: time.Now()})
	*/

	client := &http.Client{Timeout: time.Second * 5}
	form := url.Values{"code": {code}, "redirect_uri": {redirectUri}, "grant_type": {"authorization_code"}}

	request, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(form.Encode()))
	if err != nil {
		return account, err
	}
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")
	request.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", clientId, clientSecret)))))

	resp, err := client.Do(request)
	if err != nil {
		return account, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return account, err
	}

	if resp.StatusCode != 200 {
		return account, errors.New("Invalid Request")
	}

	/*
		request, err = http.NewRequest("GET", "https://api.spotify.com/v1/me", strings.NewReader(url.Values{}.Encode()))
		if err != nil {
			log.Fatalln(err)
		}
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", result["access_token"]))
	*/

	account.AccessToken = fmt.Sprintf("%v", result["access_token"])
	account.RefreshToken = fmt.Sprintf("%v", result["refresh_token"])
	account.ClientId = clientId
	account.ClientSecret = clientSecret
	return account, nil
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

func (a Account) RefreshAccessToken() error {
	accessToken, err := GetAccessTokenFromRefreshToken(a.RefreshToken, a.ClientId, a.ClientSecret)
	if err != nil {
		return err
	}

	a.AccessToken = accessToken
	return nil
}
