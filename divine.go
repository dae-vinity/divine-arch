package main

import (
	"fmt"
	"os"
	"strings"

	// "github.com/76creates/stickers/flexbox"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentFocus int
	currentTab   int
	tabs         []string
	focusLimits  []int
}

func initialModel() model {
	return model{
		currentFocus: 0,
		currentTab:   0,

		tabs: []string{
			"System",  // system specs and other info
			"Network", // internet connection for package downloading
			"Users",   // user accounts and groups
			"Packages",
		},
		focusLimits: []int{1, 1, 1, 1},
	}
}

func (model model) Init() tea.Cmd {
	return nil
}

func (model model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
	case tea.KeyMsg:
		switch message.String() {
		case "left", "esc":
			if !(model.currentFocus < 1) {
				model.currentFocus -= 1
			}
		case "right", "enter", " ":
			if !(model.currentFocus > model.focusLimits[model.currentFocus]-1) {
				model.currentFocus += 1
			}
		case "up":
			if model.currentFocus == 0 && !(model.currentTab < 1) {
				model.currentTab -= 1
			}
		case "down":
			if model.currentFocus == 0 && !(model.currentTab > len(model.tabs)-2) {
				model.currentTab += 1
			}
		case "ctrl+c", "q":
			return model, tea.Quit
		}
	}
	return model, nil
}

func (model model) View() string {
	var screen strings.Builder

	screen.WriteString("Divine Arch\n\n\n")

	for i := 0; i < len(model.tabs); i++ {
		if model.currentTab == i {
			if model.currentFocus == 0 {
				screen.WriteString("\033[93m> [ " + model.tabs[i] + " ]\033[0m")
			} else {
				screen.WriteString("> [ " + model.tabs[i] + " ]")
			}
		} else {
			screen.WriteString("  [ " + model.tabs[i] + " ]")
		}
		screen.WriteString("\n\n")
	}

	return screen.String()
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("you... IDIOT. %v", err)
		os.Exit(1)
	}
}
