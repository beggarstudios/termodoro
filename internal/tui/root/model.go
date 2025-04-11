package root

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"termodoro/internal/config"
	"termodoro/internal/tui"
	"termodoro/internal/tui/views"

	"github.com/Broderick-Westrope/charmutils"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	view     tui.View
	switchIn tui.SwitchViewInput
	db       *sql.DB
	cfg      *config.Config
}

func NewInput(view tui.View, switchIn tui.SwitchViewInput, db *sql.DB, cfg *config.Config) *Input {
	return &Input{
		view:     view,
		switchIn: switchIn,
		db:       db,
		cfg:      cfg,
	}
}

var _ tea.Model = &RootModel{}

type RootModel struct {
	child tea.Model
	db    *sql.DB
	cfg   *config.Config

	width  int
	height int

	ExitError error
}

func NewRootModel(in *Input) (*RootModel, error) {
	model := &RootModel{
		db:  in.db,
		cfg: in.cfg,
	}

	err := model.setChild(in.view, in.switchIn)
	if err != nil {
		return nil, fmt.Errorf("setting child model: %w", err)
	}
	return model, nil
}

func (m *RootModel) Init() tea.Cmd {
	return m.initChild()
}

func (m *RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tui.SwitchViewMsg:
		err := m.setChild(msg.Target, msg.Input)
		if err != nil {
			return m, FatalErrorCmd(fmt.Errorf("setting child model: %w", err))
		}
		cmd := m.initChild()
		return m, cmd

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.child, cmd = m.child.Update(msg)

	return m, cmd
}

func (m *RootModel) View() string {
	return m.child.View()
}

func (m *RootModel) initChild() tea.Cmd {
	var cmds []tea.Cmd
	cmd := m.child.Init()
	cmds = append(cmds, cmd)
	m.child, cmd = m.child.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}

func (m *RootModel) setChild(mode tui.View, switchIn tui.SwitchViewInput) error {
	if rv := reflect.ValueOf(switchIn); !rv.IsValid() || rv.IsNil() {
		return errors.New("switchIn is not valid")
	}

	switch mode {
	case tui.MenuView:
		menuIn, ok := switchIn.(*tui.MenuInput)
		if !ok {
			return fmt.Errorf("switchIn is not a MenuInput: %w", charmutils.ErrInvalidTypeAssertion)
		}
		m.child = views.NewMenuModel(menuIn)
	default:
		return errors.New("invalid Mode")
	}
	return nil
}

// FatalErrorMsg encloses an error which should be set on the starter model before exiting the program.
type FatalErrorMsg error

// FatalErrorCmd returns a command for creating a new FatalErrorMsg with the given error.
func FatalErrorCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return FatalErrorMsg(err)
	}
}
