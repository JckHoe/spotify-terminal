package spotify

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

func (c *Client) GetDevices() []Device {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/me/player/devices", c.baseUrl), nil)
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

	var response DeviceResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalln(err)
	}

	return response.Devices

}
