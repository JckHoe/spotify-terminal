package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	// TODO refactor remove this main just invoke internal keep minimal in cmd
	lib "spotify-terminal/internal/pkg"
	spotify "spotify-terminal/internal/spotify"
)

type model struct {
	currentPage *lib.Page
}

func init() {
	spotify.SpotifyKey = os.Getenv("SPOTIFY_KEY")
	if spotify.SpotifyKey == "" {
		log.Fatal("SPOTIFY_KEY variable is not set")
	}

	// TODO handle it to not require
	spotify.RefreshToken = os.Getenv("SPOTIFY_REFRESH_TOKEN")
	if spotify.RefreshToken == "" {
		log.Fatal("SPOTIFY_REFRESH_TOKEN variable is not set")
	}

	spotify.ClientId = os.Getenv("SPOTIFY_CLIENT_ID")
	if spotify.ClientId == "" {
		log.Fatal("SPOTIFY_CLIENT_ID variable is not set")
	}
	spotify.Startup()
}

func (m model) Init() tea.Cmd {
	return tea.ClearScreen
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m model) View() string {
	return m.currentPage.GetView()
}

func main() {
	initPage := lib.NewInitPage()
	p := tea.NewProgram(model{currentPage: &initPage})

	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
