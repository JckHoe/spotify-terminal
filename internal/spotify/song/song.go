package song

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
	Items []SongItem `json:"items"`
}

type SongItem struct {
	Track Track `json:"track"`
}

type Track struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}

func GetLiked(token string, httpClient *http.Client) SongResponse {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/tracks", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := httpClient.Do(req)
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

func Play(token string, httpClient *http.Client, request SongPlayRequest) error {
	reqBody, err := json.Marshal(request)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	req, err := http.NewRequest("PUT", "https://api.spotify.com/v1/me/player/play", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalln(err)
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	defer resp.Body.Close()

	// if resp.Status != "200 OK" {
	// }

	return nil
}
