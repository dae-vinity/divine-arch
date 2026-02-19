package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type dict map[string]any

type element struct {
	XMLName  xml.Name
	NameAttr xml.Attr   `xml:"name,attr"`
	Attrs    []xml.Attr `xml:",any,attr"`
	Elements []element  `xml:",any"`
	Content  string     `xml:",chardata"`
}

func convertXMLToInst(xmlElement element) dict {
	var instance dict = make(dict)
	instance["class"] = xmlElement.XMLName.Local
	if xmlElement.NameAttr.Value != "" {
		instance["name"] = xmlElement.NameAttr.Value
	} else {
		instance["name"] = instance["class"]
	}
	instance["content"] = xmlElement.Content
	for i := 0; i < len(xmlElement.Elements); i++ {
		var child dict = convertXMLToInst(xmlElement.Elements[i])
		instance[child["name"].(string)] = child
	}
	for i := 0; i < len(xmlElement.Attrs); i++ {
		instance[xmlElement.Attrs[i].Name.Local] = xmlElement.Attrs[i].Value
	}
	return instance
}

func initialModel() dict {
	xmlFile, err := os.Open("model.xml")
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()
	byteValue, _ := io.ReadAll(xmlFile)
	var xmlModel element
	xml.Unmarshal(byteValue, &xmlModel)

	var model dict = convertXMLToInst(xmlModel)
	return model
}

func (model dict) Init() tea.Cmd {
	return nil
}

func (model dict) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {

	case tea.KeyMsg:
		switch message.String() {

		case "ctrl+c", "q":
			return model, tea.Quit
		}
	}
	return model, nil
}

func (model dict) View() string {
	return model["name"].(string)
}

func main() {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := program.Run(); err != nil {
		fmt.Printf("you... IDIOT. %v", err)
		os.Exit(1)
	}
}
