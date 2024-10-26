package spotify

type DeviceResponse struct {
	Devices []Device `json:"devices"`
}

type Device struct {
	ID     string `json:"id"`
	Active bool   `json:"is_active"`
	Name   string `json:"name"`
}
