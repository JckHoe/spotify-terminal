package device

import (
	"spotify-terminal/internal/view/model"
)

func OnEnter(page *model.Page) {
	page.Name = "Main Menu"

	devices := page.SClient.GetDevices()

	var deviceList []model.Item
	for _, device := range devices {
		temp := model.Item{
			DisplayName: device.Name,
			ID:          device.ID,
			OnEnter:     func(currentPage *model.Page) {},
		}
		deviceList = append(deviceList, temp)
	}

	page.Items = deviceList

}
