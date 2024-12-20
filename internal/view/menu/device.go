package menu

import (
	"spotify-terminal/internal/view/model"
)

func DeviceOnEnter(page *model.PageState) {
	page.Cursor = 0
	page.Name = "Main Menu"

	devices := page.SClient.GetDevices()

	var deviceList []model.Item
	for _, device := range devices {
		temp := model.Item{
			DisplayName: device.Name,
			Active:      device.Active,
			OnEnter: func(currentPage *model.PageState) {
				currentPage.CurrentDeviceId = device.ID
			},
		}
		deviceList = append(deviceList, temp)
	}

	deviceList = append(deviceList, backItem())
	page.NoSubMenu = 1

	page.Items = deviceList

}
