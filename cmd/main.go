package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	spotify "spotify-terminal/internal/spotify"
	lib "spotify-terminal/internal/view"
)

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
}

func main() {
	p := tea.NewProgram(lib.NewCore())

	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running application: %v\n", err)
		os.Exit(1)
	}
}
