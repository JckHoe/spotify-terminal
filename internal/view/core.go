package view

import (
	"fmt"
	"spotify-terminal/internal/spotify"
	"spotify-terminal/internal/view/menu"
	"spotify-terminal/internal/view/model"

	tea "github.com/charmbracelet/bubbletea"
)

type Core struct {
	currentPage *model.PageState
}

func NewCore() *Core {
	mainPage := &model.PageState{
		SClient: spotify.NewClient("https://api.spotify.com/v1"),
	}
	menu.OnEnter(mainPage)
	return &Core{
		currentPage: mainPage,
	}
}

func (m Core) Init() tea.Cmd {
	return tea.Batch(tea.ClearScreen, tea.SetWindowTitle("Spotify Terminal"))
}

func (m Core) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", tea.KeyDown.String():
			if m.currentPage.Cursor < len(m.currentPage.Items)-1 {
				m.currentPage.Cursor++
			}
		case "k", tea.KeyUp.String():
			if m.currentPage.Cursor > 0 {
				m.currentPage.Cursor--
			}
		case "enter":
			m.currentPage.Items[m.currentPage.Cursor].OnEnter(m.currentPage)
		case "q":
			return m, tea.Quit
		}

	}
	return m, nil
}

func (m Core) View() string {
	s := fmt.Sprintf("%s\t\t\t%s\n\n", model.Red, m.currentPage.Name)
	s += fmt.Sprintf("%s%s\n\n", model.Blue, "Select an option:")
	for i, item := range m.currentPage.Items {
		if i == len(m.currentPage.Items)-m.currentPage.NoSubMenu {
			s += "\n"
		}

		if i == m.currentPage.Cursor {
			s += fmt.Sprintf("%s> %s\n", model.Green, getItemDisplay(item))
		} else {
			s += fmt.Sprintf("%s  %s\n", model.Yellow, getItemDisplay(item))
		}
	}

	return s
}

func getItemDisplay(input model.Item) string {
	if input.Active {
		return fmt.Sprintf("%s (Active)", input.DisplayName)
	}
	return fmt.Sprintf("%s", input.DisplayName)
}
