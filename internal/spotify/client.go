package spotify

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"spotify-terminal/internal/spotify/api"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	// TODO handle auth properly
	refreshAccessToken()
	return &Client{
		httpClient: &http.Client{},
	}
}

func (c *Client) GetDevices() []api.Device {
	return api.GetDevices(AccessToken, c.httpClient)
}

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
