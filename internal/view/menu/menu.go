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
			OnEnter:     DeviceOnEnter,
		},
		{
			DisplayName: "Liked Songs",
			OnEnter:     LikedOnEnter,
		},
	}
}

func backItem() model.Item {
	return model.Item{
		DisplayName: "Back",
		OnEnter:     OnEnter,
	}
}
