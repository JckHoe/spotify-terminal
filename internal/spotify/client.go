package spotify

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	token      string
	baseUrl    string
}

func NewClient(baseUrl string) *Client {
	// TODO handle auth properly
	refreshAccessToken()
	return &Client{
		httpClient: &http.Client{},
		token:      AccessToken,
		baseUrl:    baseUrl,
	}
}

// TODO define interface

func (c *Client) GetPlayerStatus() {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", AccessToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("\n%s\n", string(body))
}
