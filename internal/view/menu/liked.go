package menu

import (
	"spotify-terminal/internal/view/model"
)

func LikedOnEnter(page *model.Page) {
	page.Cursor = 0
	page.Name = "Liked Songs"
	page.NoSubMenu = 1

	songs := page.SClient.GetLikedSong()

	var songList []model.Item
	for _, song := range songs.Items {
		temp := model.Item{
			DisplayName: song.Track.Name,
			OnEnter: func(currentPage *model.Page) {
				currentPage.SClient.PlaySelectedSong(song.Track.Uri)
			},
		}
		songList = append(songList, temp)
	}

	songList = append(songList, backItem())

	page.Items = songList

}
