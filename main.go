package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentPage *page
}

func init() {
	spotifyKey = os.Getenv("SPOTIFY_KEY")
	if spotifyKey == "" {
		log.Fatal("SPOTIFY_KEY variable is not set")
	}

	refreshToken = os.Getenv("SPOTIFY_REFRESH_TOKEN")
	if refreshToken == "" {
		log.Fatal("SPOTIFY_REFRESH_TOKEN variable is not set")
	}

	clientId = os.Getenv("SPOTIFY_CLIENT_ID")
	if clientId == "" {
		log.Fatal("SPOTIFY_CLIENT_ID variable is not set")
	}
	go startup()
}

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd, nextPage := m.currentPage.handleKeyMsg(msg.String())
		if nextPage != nil {
			m.currentPage = nextPage
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	return m.currentPage.getView()
}

func main() {
	startupPage := page{
		name: "Main Menu",
		items: []item{
			{
				displayName: "Select Device",
				ID:          "device",
			},
			{
				displayName: "Others (To be supported)",
				ID:          "nothing",
			},
		},
	}
	p := tea.NewProgram(model{currentPage: &startupPage})

	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
