package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	timerListView uint = iota
	timerSettingsView
	timerView
)

type model struct {
	state   uint
	options []string
	cursor  int
	spinner spinner.Model
	store   *Store
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Quit the program
		case "ctrl+c", "q":
			return m, tea.Quit

		// Move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// Move the cursor down
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {
	// header
	s := "Termodoro\n\n"

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = m.spinner.View()
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func newModel(store *Store) model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		// Our pomodoro type menu options
		// 25/5 = 25 minutes of work, 5 minutes of break
		options: []string{"25/5", "45/15", "60/30"},
		state:   timerListView,
		spinner: s,
		store:   store,
	}
}

func main() {
	store := &Store{}

	if err := store.Init(); err != nil {
		fmt.Printf("Error initializing store: %v\n", err)
	}

	p := tea.NewProgram(newModel(store))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v\n", err)
		os.Exit(1)
	}
}
