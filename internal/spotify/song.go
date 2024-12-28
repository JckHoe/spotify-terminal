package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SongPlayRequest struct {
	Uris []string `json:"uris"`
}

type SongResponse struct {
	Items    []SongItem `json:"items"`
	Previous string     `json:"previous"`
	Next     string     `json:"next"`
}

type SongItem struct {
	Track Track `json:"track"`
}

type Track struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

func (c *Client) GetLiked() SongResponse {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/me/tracks", c.baseUrl), nil)
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

	var response SongResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func (c *Client) GetLikedWithUrl(requestUrl string) SongResponse {
	req, err := http.NewRequest("GET", requestUrl, nil)
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

	var response SongResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}

func (c *Client) PlaySelectedSong(track, deviceId string) error {
	request := SongPlayRequest{
		Uris: []string{track},
	}
	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	resourcePath := "/me/player/play"
	if deviceId != "" {
		resourcePath = resourcePath + fmt.Sprintf("?device_id=%s", deviceId)
	}

	req, err := http.NewRequest("PUT", c.baseUrl+resourcePath, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer resp.Body.Close()

	return nil
}
