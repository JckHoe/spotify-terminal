package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type DeviceResponse struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	ID     string `json:"id"`
	Active bool   `json:"is_active"`
	Name   string `json:"name"`
}

func startup() {
	refreshAccessToken()
	getDevicesIds()
	// getPlayerStatus()
}

func getDevicesIds() []Device {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/devices", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var response DeviceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.Devices
}

func getPlayerStatus() {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	writeToLogFile("log.json", string(body))
}
