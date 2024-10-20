package view

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Core struct {
	currentPage *Page
}

func NewCore() *Core {
	return &Core{
		currentPage: NewInitPage(),
	}
}

func (m Core) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, tea.SetWindowTitle("Spotify Terminal"))
}

func (m Core) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd, nextPage := m.currentPage.HandleKeyMsg(msg.String())
		if nextPage != nil {
			m.currentPage = nextPage
		}
		return m, cmd
	}
	return m, nil
}

func (m Core) View() string {
	return m.currentPage.GetView()
}
