package spotify

import (
	"testing"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func TestUserProfile(t *testing.T) {
	client := Client{
		ClientId:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		RefreshToken: os.Getenv("REFRESHTOKEN"),
	}

	err := client.RefreshAccessToken()
	if err != nil {
		t.Errorf("Auth Error: %v\n", err)
	}

	user, err := client.GetCurrentUserProfile()
	if err != nil {
		t.Errorf("User Error: %v\n", err)
	}

	if user.DisplayName != "Tobi Ola" {
		t.Errorf("User Error: %v\n", "User data is incorrect or missing")
	}
}
