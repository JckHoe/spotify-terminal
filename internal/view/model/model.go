package model

import (
	"spotify-terminal/internal/spotify"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

type Page struct {
	Name string

	Items   []Item
	Cursor  int
	SClient *spotify.Client
}

type Item struct {
	DisplayName string
	Active      bool
	OnEnter     func(currentPage *Page)
}
