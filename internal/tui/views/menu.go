package views

import (
	"fmt"

	"termodoro/internal/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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
___________                             .___                   
\__    ___/__________  _____   ____   __| _/___________  ____  
  |    |_/ __ \_  __ \/     \ /  _ \ / __ |/  _ \_  __ \/  _ \ 
  |    |\  ___/|  | \/  Y Y  (  <_> ) /_/ (  <_> )  | \(  <_> )
  |____| \___  >__|  |__|_|  /\____/\____ |\____/|__|   \____/ 
             \/            \/            \/                                                                                                                                                                                                          


`
)

var _ tea.Model = &MenuModel{}

type MenuModel struct {
	options []string
	cursor  int
	keys    *menuKeyMap
	help    help.Model
	spinner spinner.Model
}

func (m MenuModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {

		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Up):
			if m.cursor > 0 {
				m.cursor--
			}

		case key.Matches(msg, m.keys.Down):
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
	s := titleStr

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			//cursor = m.spinner.View()
		}

		// Render the row
		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	helpView := m.help.View(m.keys)

	// The footer
	s += "\n\n"

	s += helpView

	// Send the UI for rendering
	return s
}

func NewMenuModel(_ *tui.MenuInput) MenuModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return MenuModel{
		options: []string{
			"Timers",
			"Settings",
			"Rewards",
			"Quit",
		},
		keys:    defaultMenuKeyMap(),
		help:    help.New(),
		spinner: s,
	}
}
