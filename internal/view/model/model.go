package model

import (
	"spotify-terminal/internal/spotify"
)

const (
	// Text colors
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"

	// Borders
	CurveTopLeft   = "╭"
	CurveTopRight  = "╮"
	Verticle       = "│"
	Horizontal     = "─"
	CurveBtmLeft   = "╰"
	CurveBtmRight  = "╯"
	SharpLeftDown  = "┐"
	SharpRightDown = "┌"
	SharpLeftUp    = "┘"
	SharpRightUp   = "└"
)

type PageState struct {
	Name      string
	Items     []Item
	Cursor    int
	NoSubMenu int

	CurrentDeviceId string
	SClient         *spotify.Client

	// TODO refactor this handling
	FetchUrl string
}

type Item struct {
	DisplayName string
	Active      bool
	OnEnter     func(currentPage *PageState)
}
