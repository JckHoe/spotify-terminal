package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentPage page
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
		return m, handleKeyMsg(&m.currentPage, msg.String())
	}
	return m, nil
}

func (m model) View() string {
	return getView(&m.currentPage)
}

func main() {
	startupPage := page{
		items: []string{"Item 1", "Item 2", "Item 3", "Item 4", "Item 5"},
	}
	p := tea.NewProgram(model{currentPage: startupPage})

	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
