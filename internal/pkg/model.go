package pkg

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page struct {
	Name string

	IsInit   bool
	Items    []Item
	Cursor   int
	Selected Item

	OnEnter func() (tea.Cmd, *Page)
}

type Item struct {
	DisplayName string
	ID          string
	Active      bool
}
