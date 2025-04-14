package timerlist

import (
	"database/sql"
	"fmt"
	"termodoro/internal/data"
	"termodoro/internal/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	headerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	mainStyle     = lipgloss.NewStyle().MarginLeft(2)
	selectedStyle = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("219"))
)

var _ tea.Model = &TimerListModel{}

type TimerListModel struct {
	cursor  int
	timers  []data.Timer
	keys    *timerlistKeyMap
	help    help.Model
	spinner spinner.Model
	repo    data.TimerRepository
}

func (m TimerListModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m TimerListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			if m.cursor < len(m.timers)-1 {
				m.cursor++
			}

		case key.Matches(msg, m.keys.New):
			// Pass nill to SwitchViewCommand to create a new timer instead of editing one
			return m, tui.SwitchViewCmd(tui.TimerListAddView, tui.NewTimerListAddInput(nil))
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m TimerListModel) View() string {
	// header
	s := headerStyle.Render("Pomodoro timers") + "\n\n"

	if len(m.timers) == 0 {
		s += mainStyle.Render("No timers found.\n")
	}

	for i, timer := range m.timers {
		cursor := " "
		if m.cursor == i {
			cursor = headerStyle.Render("➤")
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, selectedStyle.Render(timer.Name))
		} else {
			// Render the row
			s += fmt.Sprintf("%s %s\n", cursor, mainStyle.Render(timer.Name))
		}
	}

	helpView := m.help.View(m.keys)

	// The footer
	s += "\n\n"
	s += helpView

	// Send the UI for rendering
	return s
}

func NewTimerListModel(_ *tui.TimerListInput, db *sql.DB) (TimerListModel, error) {
	repo := data.NewTimerRepositorySQLite(db)
	timers, err := repo.GetAllTimers()

	if err != nil {
		return TimerListModel{}, err
	}

	return TimerListModel{
		keys:    defaultTimerListKeyMap(),
		help:    help.New(),
		spinner: spinner.New(),
		repo:    repo,
		timers:  timers,
	}, nil
}
