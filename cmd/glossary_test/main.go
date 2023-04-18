package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	tui "glossary_test/internal/tui"
)

func main() {
	m := tui.NewModel()
	p := tea.NewProgram(m)

	_, err := p.Run()
	if err != nil {
		log.Fatal(err)
	}
}

