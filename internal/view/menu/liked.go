package menu

import (
	"spotify-terminal/internal/view/model"
)

func LikedOnEnter(page *model.PageState) {
	page.Cursor = 0
	page.Name = "Liked Songs"
	page.NoSubMenu = 1

	songs := page.SClient.GetLiked()

	var songList []model.Item
	for _, song := range songs.Items {
		temp := model.Item{
			DisplayName: song.Track.Name,
			OnEnter: func(currentPage *model.PageState) {
				currentPage.SClient.PlaySelectedSong(song.Track.Uri, currentPage.CurrentDeviceId)
			},
		}
		songList = append(songList, temp)
	}

	songList = append(songList, backItem())

	page.Items = songList

}
