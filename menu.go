package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type page struct {
	name string

	isInit   bool
	items    []item
	cursor   int
	selected item
}

type item struct {
	displayName string
	ID          string
	active      bool
}

func (current *page) handleKeyMsg(keyMsg string) (tea.Cmd, *page) {
	switch keyMsg {
	case "j", tea.KeyDown.String():
		if current.cursor < len(current.items)-1 {
			current.cursor++
		}
	case "k", tea.KeyUp.String():
		if current.cursor > 0 {
			current.cursor--
		}
	case "enter":
		if current.items[current.cursor].ID == "device" {
			current.selected = current.items[current.cursor]
			devices := getDevices()
			var pageItems []item
			for _, device := range devices {
				pageItems = append(pageItems, item{
					displayName: device.Name,
					ID:          device.ID,
					active:      device.Active,
				})
			}
			return tea.ClearScreen, &page{
				name:  "Devices",
				items: pageItems,
			}
		}
	case "q", "ctrl+c":
		return tea.Quit, nil
	}
	return nil, nil
}

func (current *page) getView() string {
	s := fmt.Sprintf("*** %s ***\n", current.name)
	s += "Select an option:\n\n"
	for i, item := range current.items {
		if i == current.cursor {
			s += fmt.Sprintf("> %s\n", item.displayName)
		} else {
			s += fmt.Sprintf("  %s\n", item.displayName)
		}
	}
	s += fmt.Sprintf("You selected: %s\n", current.selected.displayName)
	return s

}
