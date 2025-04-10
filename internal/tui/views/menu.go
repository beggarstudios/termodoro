package views

import (
	"fmt"

	"termodoro/internal/data"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// const (
// 	timerListView uint = iota
// 	timerSettingsView
// 	timerView
// )

const (
	titleStr = `
 _________  _______   ________  _____ ______   ________  ________  ________  ________  ________     
|\___   ___\\  ___ \ |\   __  \|\   _ \  _   \|\   __  \|\   ___ \|\   __  \|\   __  \|\   __  \    
\|___ \  \_\ \   __/|\ \  \|\  \ \  \\\__\ \  \ \  \|\  \ \  \_|\ \ \  \|\  \ \  \|\  \ \  \|\  \   
     \ \  \ \ \  \_|/_\ \   _  _\ \  \\|__| \  \ \  \\\  \ \  \ \\ \ \  \\\  \ \   _  _\ \  \\\  \  
      \ \  \ \ \  \_|\ \ \  \\  \\ \  \    \ \  \ \  \\\  \ \  \_\\ \ \  \\\  \ \  \\  \\ \  \\\  \ 
       \ \__\ \ \_______\ \__\\ _\\ \__\    \ \__\ \_______\ \_______\ \_______\ \__\\ _\\ \_______\
        \|__|  \|_______|\|__|\|__|\|__|     \|__|\|_______|\|_______|\|_______|\|__|\|__|\|_______|                                                                                                                                                                                                    
`
)

var _ tea.Model = &MenuModel{}

type MenuModel struct {
	//state   uint
	options []string
	cursor  int
	spinner spinner.Model
	store   *data.Store
}

func (m MenuModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m MenuModel) View() string {
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

func newModel(store *data.Store) MenuModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return MenuModel{
		// Our pomodoro type menu options
		// 25/5 = 25 minutes of work, 5 minutes of break
		options: []string{"25/5", "45/15", "60/30"},
		//state:   timerListView,
		spinner: s,
		store:   store,
	}
}
