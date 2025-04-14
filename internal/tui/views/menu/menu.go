package menu

import (
	"fmt"

	"termodoro/internal/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

// TODO Move ti seperate package
var (
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	mainStyle    = lipgloss.NewStyle().MarginLeft(2)

	selectedStyle = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("219"))

	// Gradient colors we'll use for the progress bar
	//ramp = makeRampStyles("#B14FFF", "#00FFA3", progressBarWidth)
)

var _ tea.Model = &MenuModel{}

type MenuModel struct {
	options []tui.MenuItem
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
		case key.Matches(msg, m.keys.Select):
			// Call Action on the menu item
			return m, m.options[m.cursor].Action()
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m MenuModel) View() string {
	// header
	s := ""

	for i, choice := range m.options {
		cursor := " "
		if m.cursor == i {
			cursor = keywordStyle.Render("➤")
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, selectedStyle.Render(choice.Label))
		} else {
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, mainStyle.Render(choice.Label))
		}
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
		options: []tui.MenuItem{
			{Label: "Timers", Action: func() tea.Cmd { return navigateToTimers() }},
			{Label: "Settings", Action: func() tea.Cmd { return navigateToSettings() }},
			{Label: "Rewards", Action: func() tea.Cmd { return navigateToRewards() }},
			{Label: "Quit", Action: func() tea.Cmd { return quitApp() }},
		},
		keys:    defaultMenuKeyMap(),
		help:    help.New(),
		spinner: s,
	}
}
