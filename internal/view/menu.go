package view

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"spotify-terminal/internal/spotify"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	White  = "\033[37m"
)

func NewInitPage() *Page {
	return &Page{
		Name: "Main Menu",
		Items: []Item{
			{
				DisplayName: "Select Device",
				ID:          "device",
				OnEnter: func() (tea.Cmd, *Page) {
					devices := spotify.GetDevices()
					var pageItems []Item
					for _, device := range devices {
						pageItems = append(pageItems, Item{
							DisplayName: device.Name,
							ID:          device.ID,
							Active:      device.Active,
						})
					}

					// Add return choice
					pageItems = append(pageItems, Item{
						DisplayName: "Back to Menu",
						ID:          "main",
						OnEnter: func() (tea.Cmd, *Page) {
							initPage := NewInitPage()
							return tea.ClearScreen, initPage
						},
					})

					return tea.ClearScreen, &Page{
						Name:  "Devices",
						Items: pageItems,
					}
				},
			},
			{
				DisplayName: "Others (To be supported)",
				ID:          "nothing",
				OnEnter: func() (tea.Cmd, *Page) {
					// TODO this is just to debug this API
					spotify.GetPlayerStatus()
					return nil, nil
				},
			},
		},
	}
}

func (current *Page) HandleKeyMsg(keyMsg string) (tea.Cmd, *Page) {
	switch keyMsg {
	case "j", tea.KeyDown.String():
		if current.Cursor < len(current.Items)-1 {
			current.Cursor++
		}
	case "k", tea.KeyUp.String():
		if current.Cursor > 0 {
			current.Cursor--
		}
	case "enter":
		return current.Items[current.Cursor].OnEnter()
	case "q":
		return tea.Quit, nil
	}
	return nil, nil
}

func (current *Page) GetView() string {
	s := fmt.Sprintf("%s\t\t\t%s\n\n", Red, current.Name)
	s += fmt.Sprintf("%s%s\n\n", Blue, "Select an option:")
	for i, item := range current.Items {
		if i == current.Cursor {
			s += fmt.Sprintf("%s> %s\n", Green, getItemDisplay(item))
		} else {
			s += fmt.Sprintf("%s  %s\n", Yellow, getItemDisplay(item))
		}
	}
	return s
}

func getItemDisplay(input Item) string {
	if input.Active {
		return fmt.Sprintf("%s (Active)", input.DisplayName)
	}
	return fmt.Sprintf("%s", input.DisplayName)
}
