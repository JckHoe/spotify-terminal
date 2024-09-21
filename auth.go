package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// auth vars
var spotifyKey string
var refreshToken string
var clientId string
var accessToken string

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}

func refreshAccessToken() {
	formData := url.Values{}
	formData.Set("refresh_token", refreshToken)
	formData.Set("client_id", clientId)
	formData.Set("grant_type", "refresh_token")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(formData.Encode()))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", spotifyKey))

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
	accessToken = authResp.AccessToken
}
