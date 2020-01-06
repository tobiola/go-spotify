package spotify

import (
	"testing"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func TestRefreshToken(t *testing.T) {
	client := Client{
		ClientId: os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		RefreshToken: os.Getenv("REFRESHTOKEN"),
	}

	err := client.RefreshAccessToken()
	if err != nil {
		t.Errorf("Auth Error: %s\n", err)
	}
}