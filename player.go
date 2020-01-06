package spotify

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type DevicesPayload struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	Id               string `json:"id"`
	IsActive         bool   `json:"is_active"`
	IsPrivateSession bool   `json:"is_private_session"`
	IsRestricted     bool   `json:"is_restricted"`
	Name             bool   `json:"name"`
	Type             string `json:"type"`
	VolumePercent    int    `json:"volume_percent"`
}

type CurrentlyPlayingContext struct {
	Device               Device  `json:"device"`
	RepeatState          string  `json:"repeat_state"`
	ShuffleState         bool    `json:"shuffle_state"`
	Context              Context `json:"context"`
	Timestamp            int     `json:"timestamp"`
	ProgressMs           int     `json:"progress_ms"`
	IsPlaying            bool    `json:"is_playing"`
	Item                 Track   `json:"item"`
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Actions              Actions `json:"actions"`
}

type Actions struct {
	Disallows Disallows `json:"disallows"`
}

type PlayHistory struct {
	Track    TrackSimple `json:"track"`
	PlayedAt string      `json:"played_at"`
	Context  Context     `json:"context"`
}

type PlayHistoryCursorBasedPaging struct {
	Href    string      `json:"href"`
	Items   PlayHistory `json:"items"`
	Limit   int         `json:"limit"`
	Next    string      `json:"next"`
	Cursors Cursor      `json:"cursors"`
	Total   int         `json:"total"`
}

// Get a User's Available Devices
func (c *Client) GetDevices() ([]Device, error) {
	devices := []Device{}

	bytes, err := c.get("https://api.spotify.com/v1/me/player/devices", nil)
	if err != nil {
		return devices, err
	}

	payload := DevicesPayload{}
	err = json.Unmarshal(bytes, &payload)
	if err != nil {
		return devices, err
	}

	return payload.Devices, nil
}

// Get information about the user’s current playback state, including track, track progress, and active device.
func (c *Client) GetCurrentPlayback() (CurrentlyPlayingContext, error) {
	playback := CurrentlyPlayingContext{}

	bytes, err := c.get("https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		return playback, err
	}

	err = json.Unmarshal(bytes, &playback)
	if err != nil {
		return playback, err
	}

	return playback, err
}

// Get tracks from the current user’s recently played tracks.
// Returns the most recent 20 tracks played by a user.
// Note that a track currently playing will not be visible in play history until it has completed.
// A track must be played for more than 30 seconds to be included in play history.
// Any tracks listened to while the user had “Private Session” enabled in their client will not be returned in the list of recently played tracks.
func (c *Client) GetRecentlyPlayedTracks() (PlayHistoryCursorBasedPaging, error) {
	cursor := PlayHistoryCursorBasedPaging{}
	bytes, err := c.get("https://api.spotify.com/v1/me/player/recently-played", nil)
	if err != nil {
		return cursor, err
	}

	err = json.Unmarshal(bytes, &cursor)
	if err != nil {
		return cursor, err
	}

	return cursor, err
}

// The endpoint uses a bidirectional cursor for paging.
// Follow the next field with the before parameter to move back in time, or use the after parameter to move forward in time.
// If you supply no before or after parameter, the endpoint will return the most recently played songs, and the next link will page back in time.
func (c *Client) GetRecentlyPlayedTracksBefore(before string) (PlayHistoryCursorBasedPaging, error) {
	cursor := PlayHistoryCursorBasedPaging{}
	bytes, err := c.get("https://api.spotify.com/v1/me/player/recently-played?before="+before, nil)
	if err != nil {
		return cursor, err
	}

	err = json.Unmarshal(bytes, &cursor)
	if err != nil {
		return cursor, err
	}

	return cursor, err
}

// The endpoint uses a bidirectional cursor for paging.
// Follow the next field with the before parameter to move back in time, or use the after parameter to move forward in time.
// If you supply no before or after parameter, the endpoint will return the most recently played songs, and the next link will page back in time.
func (c *Client) GetRecentlyPlayedTracksAfter(after string) (PlayHistoryCursorBasedPaging, error) {
	cursor := PlayHistoryCursorBasedPaging{}
	bytes, err := c.get("https://api.spotify.com/v1/me/player/recently-played?after="+after, nil)
	if err != nil {
		return cursor, err
	}

	err = json.Unmarshal(bytes, &cursor)
	if err != nil {
		return cursor, err
	}

	return cursor, err
}

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
