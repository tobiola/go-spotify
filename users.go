package spotify

import (
	"fmt"
	"encoding/json"
)

func (c *Client) GetCurrentUserProfile() (UserPrivate, error) {
	url := "https://api.spotify.com/v1/me"
	bytes, err := c.get(url, nil)
	if err != nil {
		return UserPrivate{}, err
	}

	user := UserPrivate{}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return UserPrivate{}, err
	}

	return user, nil
}

func (c *Client) GetUserProfile(userId string) (User, error) {
	url := fmt.Sprintf("https://api.spotify.com/v1/users/%s", userId)
	bytes, err := c.get(url, nil)
	if err != nil {
		return User{}, err
	}

	user := User{}
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}