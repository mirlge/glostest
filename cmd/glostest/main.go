package main

import (
	"encoding/json"
	tui "glostest/internal/tui"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	var glossary []tui.Gloss
	for i, file := range os.Args {
		if i == 0 {
			continue
		}
		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatal("Error when opening file ${file}: ", err)
			os.Exit(1)
		}
		var glossarySlice []tui.Gloss
		err = json.Unmarshal(content, &glossarySlice)
		if err != nil {
			log.Fatal("Error during JSON parsing: ", err)
			os.Exit(1)
		}
		glossary = append(glossary, glossarySlice...)
	}

	m := tui.NewMenu(glossary)
	p := tea.NewProgram(m, tea.WithAltScreen())

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}
