package song

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SongResponse struct {
	Items []SongItem `json:"items"`
}

type SongItem struct {
	Track Track `json:"track"`
}

type Track struct {
	Name string `json:"name"`
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
