package view

import (
	"fmt"
	"log"
	"spotify-terminal/internal/spotify"
	"spotify-terminal/internal/view/menu"
	"spotify-terminal/internal/view/model"
	"strings"

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
	// TODO dynamic check border lengths
	maxHorizontalLength := 100
	var viewBuilder strings.Builder

	if err := renderTitle(&viewBuilder, m.currentPage.Name, uint(maxHorizontalLength)); err != nil {
		log.Fatalf("Failed to render Title %v", err)
		return ""
	}

	for i, item := range m.currentPage.Items {
		if i == len(m.currentPage.Items)-m.currentPage.NoSubMenu {
			if err := renderEmptyRow(&viewBuilder, uint(maxHorizontalLength)); err != nil {
				log.Fatalln(err)
				return ""
			}
		}

		if i == m.currentPage.Cursor {
			if err := renderItem(&viewBuilder, fmt.Sprintf("> %s", getItemDisplay(item)), uint(maxHorizontalLength), true); err != nil {
				log.Fatalln(err)
				return ""
			}
		} else {
			if err := renderItem(&viewBuilder, fmt.Sprintf("  %s", getItemDisplay(item)), uint(maxHorizontalLength), false); err != nil {
				log.Fatalln(err)
				return ""
			}
		}
	}

	if err := renderFooter(&viewBuilder, uint(maxHorizontalLength)); err != nil {
		log.Fatalln(err)
		return ""
	}

	return viewBuilder.String()
}

func getItemDisplay(input model.Item) string {
	if input.Active {
		return fmt.Sprintf("%s (Active)", input.DisplayName)
	}
	return fmt.Sprintf("%s", input.DisplayName)
}
