package spotify

import (
	"testing"
	//	"log"
	//	"github.com/joho/godotenv"
	"os"
)

import _ "github.com/joho/godotenv/autoload"

func TestSearch(t *testing.T) {
	client := Client{
		ClientId:     os.Getenv("CLIENTID"),
		ClientSecret: os.Getenv("CLIENTSECRET"),
		RefreshToken: os.Getenv("REFRESHTOKEN"),
	}

	err := client.RefreshAccessToken()
	if err != nil {
		t.Errorf("Auth Error: %v\n", err)
	}

	options := SearchOptions{Query: "hotline bling", Type: "track"}
	_, err = client.Search(options)
	if err != nil {
		t.Errorf("Search Error: %v\n", err)
	}
}
