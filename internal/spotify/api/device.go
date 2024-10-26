package api

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

func GetDevices(token string, httpClient *http.Client) []Device {
	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/me/player/devices", nil)
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

	var response DeviceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.Devices

}
