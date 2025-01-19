package menu

import (
	"spotify-terminal/internal/spotify"
	"spotify-terminal/internal/view/model"
)

func PlaylistTracksOnEnter(page *model.PageState) {
	page.Cursor = 0
	page.NoSubMenu = 1

	var songs spotify.SongResponse
	if page.FetchUrl != "" {
		songs = page.SClient.GetTracksWithUrl(page.FetchUrl)
	}
	var items []model.Item
	for _, song := range songs.Items {
		songItem := model.Item{
			DisplayName: song.Track.Name,
			OnEnter: func(currentPage *model.PageState) {
				currentPage.SClient.PlaySelectedSong(song.Track.Uri, currentPage.CurrentDeviceId)
			},
		}
		items = append(items, songItem)
	}

	if songs.Previous != "" {
		page.NoSubMenu++
		prevItem := model.Item{
			DisplayName: "Previous",
			OnEnter: func(currentPage *model.PageState) {
				currentPage.FetchUrl = songs.Previous
				PlaylistTracksOnEnter(currentPage)
			},
		}
		items = append(items, prevItem)
	}

	if songs.Next != "" {
		page.NoSubMenu++
		nextItem := model.Item{
			DisplayName: "Next",
			OnEnter: func(currentPage *model.PageState) {
				currentPage.FetchUrl = songs.Next
				PlaylistTracksOnEnter(currentPage)
			},
		}
		items = append(items, nextItem)
	}

	items = append(items, backItem())

	page.Items = items

}
