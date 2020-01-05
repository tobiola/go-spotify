package spotify

import (
	"io"
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
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
	bytes := []byte{}

	client := &http.Client{Timeout: time.Second * 5}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return bytes, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AccessToken))

	resp, err := client.Do(request)
	if err != nil {
		return bytes, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
