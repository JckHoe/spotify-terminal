package spotify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

var SpotifyKey string
var RefreshToken string
var ClientId string
var AccessToken string

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func refreshAccessToken() {
	formData := url.Values{}
	formData.Set("refresh_token", RefreshToken)
	formData.Set("client_id", ClientId)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", SpotifyKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var authResp AuthResponse
	err = json.Unmarshal(body, &authResp)
	// Set The access token
	AccessToken = authResp.AccessToken
}
