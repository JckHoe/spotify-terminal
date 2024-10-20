package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Page struct {
	Name string

	IsInit bool
	Items  []Item
	Cursor int
}

type Item struct {
	DisplayName string
	ID          string
	Active      bool
	OnEnter     func() (tea.Cmd, *Page)
}
