package spotify

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type PlaylistItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PlaylistResponse struct {
	Items    []PlaylistItem `json:"items"`
	Previous string         `json:"previous"`
	Next     string         `json:"next"`
}

func (c *Client) GetPlaylist() PlaylistResponse {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/me/playlists?limit=20", c.baseUrl), nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response PlaylistResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func (c *Client) GetPlaylistWithUrl(url string) PlaylistResponse {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response PlaylistResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response

}
