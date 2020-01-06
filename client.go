package spotify

import (
	"fmt"
	"io"
	"errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	AccessToken  string
	RefreshToken string
	ClientId     string
	ClientSecret string
}

func (c *Client) get(url string, body io.Reader) ([]byte, error) {
	return c.fetch("GET", url, body)
}

func (c *Client) post(url string, body io.Reader) ([]byte, error) {
	return c.fetch("POST", url, body)
}

func (c *Client) put(url string, body io.Reader) ([]byte, error) {
	return c.fetch("PUT", url, body)
}

func (c *Client) delete(url string, body io.Reader) ([]byte, error) {
	return c.fetch("DELETE", url, body)
}

func (c *Client) fetch(method string, url string, body io.Reader) ([]byte, error) {
	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body) 
	if err != nil {
		return []byte{}, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return []byte{}, err
	}

	if result["error"] != nil {
		return []byte{}, errors.New(fmt.Sprintf("%s", result["error_description"]))
	}

	return bytes, nil
}
