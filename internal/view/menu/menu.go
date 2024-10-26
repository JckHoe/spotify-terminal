package menu

import (
	"spotify-terminal/internal/view/model"
)

func OnEnter(page *model.Page) {
	page.Cursor = 0
	page.Name = "Main Menu"
	page.Items = []model.Item{
		{
			DisplayName: "Select Device",
			ID:          "device",
			OnEnter:     DeviceOnEnter,
		},
	}
}

func backItem() model.Item {
	return model.Item{
		DisplayName: "Back",
		ID:          "back",
		OnEnter:     OnEnter,
	}
}
