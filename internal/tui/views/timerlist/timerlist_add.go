package timerlist

import (
	"database/sql"
	"fmt"
	"strings"
	"termodoro/internal/data"
	"termodoro/internal/tui"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

var _ tea.Model = &TimerListAddModel{}

type TimerListAddModel struct {
	repo       data.TimerRepository
	focusIndex int
	inputs     []textinput.Model
	timerId    *int64
	help       help.Model
	spinner    spinner.Model
	cursorMode cursor.Mode
	keys       *timerlistAddKeyMap
}

func (m TimerListAddModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TimerListAddModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll

		case key.Matches(msg, m.keys.Exit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Up) || key.Matches(msg, m.keys.Down):
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *TimerListAddModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m TimerListAddModel) View() string {
	var b strings.Builder

	b.WriteString(focusedStyle.Render("Add a new timer") + "\n\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteRune('\n')

	b.WriteString(m.help.View(m.keys))

	return b.String()
}

func NewTimerListAddModel(input *tui.TimerListAddInput, db *sql.DB) (TimerListAddModel, error) {
	repo := data.NewTimerRepositorySQLite(db)

	m := TimerListAddModel{
		keys:    defaultTimerListAddKeyMap(),
		help:    help.New(),
		spinner: spinner.New(),
		inputs:  make([]textinput.Model, 4),
		repo:    repo,
	}

	if input.TimerId == nil {
		m.timerId = input.TimerId
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.Cursor.SetMode(cursor.CursorHide)
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Timer name"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
		case 1:
			t.Placeholder = "Description"
			t.CharLimit = 64
		case 2:
			t.Placeholder = "Minutes to work"
			t.CharLimit = 64
		case 3:
			t.Placeholder = "Minutes to rest"
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m, nil
}
