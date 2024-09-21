package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	items    []string
	cursor   int
	selected string
}

func (m model) Init() tea.Cmd {
	spotifyKey = os.Getenv("SPOTIFY_KEY")
	if spotifyKey == "" {
		log.Fatal("SPOTIFY_KEY variable is not set")
	}

	refreshToken = os.Getenv("SPOTIFY_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Fatal("SPOTIFY_REFRESH_TOKEN variable is not set")
	}

	go refreshAccessToken()

	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", tea.KeyDown.String():
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "k", tea.KeyUp.String():
			if m.cursor > 0 {
				m.cursor--
			}
		case "enter":
			m.selected = m.items[m.cursor]
			return m, tea.ClearScreen
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Select an item:\n\n"
	s += fmt.Sprintf("You selected: %s\n", m.selected)
	for i, item := range m.items {
		if i == m.cursor {
			s += fmt.Sprintf("> %s\n", item)
		} else {
			s += fmt.Sprintf("  %s\n", item)
		}
	}
	return s
}

func main() {
	items := []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"}
	p := tea.NewProgram(model{items: items})

	finalModel, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}

	finalModelState, ok := finalModel.(model)
	if ok && finalModelState.selected != "" {
		fmt.Printf("Final selected: %s\n", finalModelState.selected)
	}
}
