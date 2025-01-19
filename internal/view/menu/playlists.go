package menu

import (
	"fmt"
	"spotify-terminal/internal/spotify"
	"spotify-terminal/internal/view/model"
)

func MyPlayListOnEnter(page *model.PageState) {
	page.Cursor = 0
	page.Name = "Playlists"
	page.NoSubMenu = 1

	var playlists spotify.PlaylistResponse
	if page.FetchUrl != "" {
		playlists = page.SClient.GetPlaylistWithUrl(page.FetchUrl)
	} else {
		playlists = page.SClient.GetPlaylist()
	}

	var items []model.Item
	for _, playlist := range playlists.Items {
		playlistItem := model.Item{
			DisplayName: playlist.Name,
			OnEnter: func(currentPage *model.PageState) {
				currentPage.Name = playlist.Name
				currentPage.FetchUrl = fmt.Sprintf("https://api.spotify.com/v1/playlists/%s/tracks?limit=20", playlist.ID)
				PlaylistTracksOnEnter(currentPage)
			},
		}
		items = append(items, playlistItem)
	}

	if playlists.Previous != "" {
		page.NoSubMenu++
		prevItem := model.Item{
			DisplayName: "Previous",
			OnEnter: func(currentPage *model.PageState) {
				currentPage.FetchUrl = playlists.Previous
				MyPlayListOnEnter(currentPage)
			},
		}
		items = append(items, prevItem)
	}

	if playlists.Next != "" {
		page.NoSubMenu++
		nextItem := model.Item{
			DisplayName: "Next",
			OnEnter: func(currentPage *model.PageState) {
				currentPage.FetchUrl = playlists.Next
				MyPlayListOnEnter(currentPage)
			},
		}
		items = append(items, nextItem)
	}

	items = append(items, backItem())

	page.Items = items

}
