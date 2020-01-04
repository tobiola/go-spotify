package spotify

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Start/Resume a User's Playback
func (a Account) Play() error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// Start/Resume a User's Playback
func (a Account) PlayTrack(track Track) error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", strings.NewReader(fmt.Sprintf(`{"uris":["%s"]}`, track.Uri)))
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}

// Start/Resume a User's Playback
func (a Account) PlayTracks(tracks []Track) error {
	return errors.New("PlayTracks has not been implemented yet")
}

// Pause a User's Playback
func (a Account) Pause() error {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/pause", nil)
	if err != nil {
		return err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	_, err = client.Do(request)
	if err != nil {
		return err
	}

	return nil
}
