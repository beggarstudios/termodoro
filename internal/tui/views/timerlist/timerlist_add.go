package timerlist

import (
	"database/sql"
	"termodoro/internal/data"
	"termodoro/internal/tui"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
)

var _ tea.Model = &TimerListAddModel{}

type TimerListAddModel struct {
	repo    data.TimerRepository
	timerId *int64
	help    help.Model
	spinner spinner.Model
	keys    *timerlistKeyMap
}

func (m TimerListAddModel) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m TimerListAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		// case key.Matches(msg, m.keys.Up):
		// 	if m.cursor > 0 {
		// 		m.cursor--
		// 	}

		// case key.Matches(msg, m.keys.Down):
		// 	if m.cursor < len(m.timers)-1 {
		// 		m.cursor++
		// 	}
		// }

		default:
			var cmd tea.Cmd
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m TimerListAddModel) View() string {
	return "Testing TimerListAddModel"
}

func NewTimerListAddModel(input *tui.TimerListAddInput, db *sql.DB) (TimerListAddModel, error) {
	repo := data.NewTimerRepositorySQLite(db)

	addModel := TimerListAddModel{
		keys:    defaultMenuKeyMap(),
		help:    help.New(),
		spinner: spinner.New(),
		repo:    repo,
	}

	if input.TimerId == nil {
		addModel.timerId = input.TimerId
	}

	return addModel, nil
}
