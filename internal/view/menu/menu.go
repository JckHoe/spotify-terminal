package menu

import (
	"spotify-terminal/internal/view/model"
)

func OnEnter(page *model.PageState) {
	page.Cursor = 0
	page.Name = "Main Menu"
	page.NoSubMenu = 0
	page.Items = []model.Item{
		{
			DisplayName: "Select Device",
			OnEnter:     DeviceOnEnter,
		},
		{
			DisplayName: "Liked Songs",
			OnEnter:     LikedOnEnter,
		},
		{
			DisplayName: "Pause Player",
			OnEnter: func(currentPage *model.PageState) {
				currentPage.SClient.PausePlayer()
			},
		},
	}
}

func backItem() model.Item {
	return model.Item{
		DisplayName: "Back",
		OnEnter:     OnEnter,
	}
}
