package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type page struct {
	name string

	items    []string
	cursor   int
	selected string
}

func (current *page) handleKeyMsg(keyMsg string) tea.Cmd {
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
		if current.selected != current.items[current.cursor] {
			current.selected = current.items[current.cursor]
			return tea.ClearScreen
		}
	case "q", "ctrl+c":
		return tea.Quit
	}
	return nil
}

func (current *page) getView() string {
	s := "Select an item:\n\n"
	s += fmt.Sprintf("You selected: %s\n", current.selected)
	for i, item := range current.items {
		if i == current.cursor {
			s += fmt.Sprintf("> %s\n", item)
		} else {
			s += fmt.Sprintf("  %s\n", item)
		}
	}
	return s

}
