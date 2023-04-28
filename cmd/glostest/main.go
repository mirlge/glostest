package main

import (
	"encoding/json"
	tui "glostest/internal/tui"
	"io/ioutil"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		os.Exit(1)
	}
	var glossary []tui.Gloss
	err = json.Unmarshal(content, &glossary)
	if err != nil {
		log.Fatal("Error during JSON parsing: ", err)
		os.Exit(1)
	}

	m := tui.NewMenu(glossary)
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err = p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
