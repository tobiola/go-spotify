package spotify

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Start/Resume a User's Playback
func Play(accessToken string) error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// Start/Resume a User's Playback
func PlayTrack(accessToken string, track Track) error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", strings.NewReader(fmt.Sprintf(`{"uris":["%s"]}`, track.Uri)))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// Pause a User's Playback
func Pause(accessToken string) error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
