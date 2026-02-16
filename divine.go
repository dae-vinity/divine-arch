package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"

	// "github.com/76creates/stickers/flexbox"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	screen []string
}

func initialModel() model {
	modelFile, err := os.ReadFile("model.xml")
	if err != nil {
		log.Fatal(err)
	}
	var model model
	xml.Unmarshal(modelFile, &model)
	return model
}

func (model model) Init() tea.Cmd {
	return nil
}

func (model model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {

		case tea.KeyMsg:
        switch message.String() {

        case "ctrl+c", "q":
            return model, tea.Quit
		}
	}
	return model, nil
}

func (model model) View() string {
	var screen strings.Builder

	screen.WriteString("")

	return screen.String()
}

func main() {
	
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("you... IDIOT. %v", err)
		os.Exit(1)
	}
}
