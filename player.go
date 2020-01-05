package spotify

import (
	"errors"
	"fmt"
	"strings"
)

// Start/Resume a User's Playback
func (c *Client) Play() error {
	_, err := c.put("https://api.spotify.com/v1/me/player/play", nil)
	if err != nil {
		return err
	}

	return nil
}

// Start/Resume a User's Playback
func (c *Client) PlayTrack(track Track) error {
	_, err := c.put("https://api.spotify.com/v1/me/player/play", strings.NewReader(fmt.Sprintf(`{"uris":["%s"]}`, track.Uri)))
	if err != nil {
		return err
	}

	return nil
}

// Start/Resume a User's Playback
func (c *Client) PlayTracks(tracks []Track) error {
	return errors.New("PlayTracks has not been implemented yet")
}

// Pause a User's Playback
func (c *Client) Pause() error {
	_, err := c.put("https://api.spotify.com/v1/me/player/pause", nil)
	if err != nil {
		return err
	}

	return nil
}
