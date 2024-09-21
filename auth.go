package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

// auth vars
var spotifyKey string
var refreshToken string
var accessToken string

func refreshAccessToken() {
	formData := url.Values{}
	formData.Set("code", refreshToken)
	formData.Set("redirect_uri", "http://localhost:8888/callback")
	formData.Set("grant_type", "authorization_code")

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

	log.Printf("Status : %s", resp.Status)
	log.Printf("Content: %s", string(body))
}
