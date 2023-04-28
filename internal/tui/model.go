package tui

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

/// A gloss
type Gloss struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}

func (g Gloss) FilterValue() string {
	return g.Definition
}

func (g Gloss) Title() string {
	return g.Definition
}

/// Results is the view of the results
type Results struct {
	correct int
	glosses int
	err     error
}

func NewResults(correct, glosses int) Results {
	return Results{
		correct: correct,
		glosses: glosses,
	}
}

func (m Results) Init() tea.Cmd {
	return nil
}

func (m Results) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Results) View() string {
	return fmt.Sprintf("Results: %d correct, %d total", m.correct, m.glosses)
}

/// Write is spelling test
type Write struct {
	glossary      []Gloss
	glossIdx      int
	glosses       int
	firstGloss    bool
	correctAmount int
	correct       bool
	input         textinput.Model
	err           error
}

func NewWrite(glossary []Gloss) Write {
	glosses := len(glossary)
	write := Write{
		glossary:   glossary,
		glosses:    glosses,
		glossIdx:   rand.Intn(glosses - 1),
		firstGloss: true,
	}
	write.input = textinput.New()
	write.input.Focus()
	return write
}

func (m Write) Init() tea.Cmd {
	return textinput.Blink
}

func (m Write) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.input.Value() == m.glossary[m.glossIdx].Definition {
				m.correctAmount++
				m.correct = true
			} else {
				m.correct = false
			}
			if m.firstGloss {
				m.firstGloss = false
			}
			m.glossary[m.glossIdx] = m.glossary[len(m.glossary)-1]
			m.glossary = m.glossary[:len(m.glossary)-1]

			if len(m.glossary) == 0 {
				r := NewResults(m.correctAmount, m.glosses)
				return r.Update(nil)
			}
			m.glossIdx = rand.Intn(len(m.glossary))
			m.input = textinput.New()
			m.input.Focus()
			return m, cmd
		}

	case ErrMsg:
		m.err = msg
		return m, nil
	}
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

func (m Write) View() string {
	correctMsg := ""
	if !m.firstGloss {
		if !m.correct {
			correctMsg = "Incorrect"
		} else if m.correct {
			correctMsg = "Correct! Good job!"
		}
	}
	return lipgloss.JoinVertical(lipgloss.Left, correctMsg, fmt.Sprintf("What does '%s' mean?", m.glossary[m.glossIdx].Term), m.input.View())
}

/// MenuItem is an item of the Menu
type MenuItem struct {
	title       string
	description string
}

func (i MenuItem) FilterValue() string {
	return i.title
}

func (i MenuItem) Title() string {
	return i.title
}

func (i MenuItem) Description() string {
	return i.description
}

/// Menu is the root state of the app.
type Menu struct {
	glossary []Gloss
	loaded   bool
	list     list.Model
	err      error
}

/// NewMenu configures the initial model at runtime.
func NewMenu(glossary []Gloss) Menu {
	return Menu{
		glossary: glossary,
	}
}

/// initList initializes the menu list
func (m *Menu) initList(width, height int) {
	menuItems := list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	menuItems.Title = "Glossary Test"
	menuItems.SetItems([]list.Item{
		MenuItem{
			title:       "Write",
			description: "Test if you remember the definitions of words how to spell them",
		},
		MenuItem{
			title:       "Multiple choice",
			description: "Test if you remember the definition of words",
		},
	})
	m.list = menuItems
}

/// Init returns any number of tea.Cmds at runtime.
func (m Menu) Init() tea.Cmd {
	return nil
}

/// Update handles all tea.Msgs in the Bubble Tea event loop.
func (m Menu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		if !m.loaded {
			m.initList(msg.Width, msg.Height)
			m.loaded = true
		}

	// Handle keypress messages.
	case tea.KeyMsg:
		switch msg.String() {

		case "enter":
			switch m.list.Index() {

			case 0:
				w := NewWrite(m.glossary)
				return w.Update(nil)
			}
		}

	case ErrMsg:
		m.err = msg
		return m, nil

	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

/// View renders a string representation of the Menu.
func (m Menu) View() string {
	if m.loaded {
		return m.list.View()
	}
	return "Loading..."
}
