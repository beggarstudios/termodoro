package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type View int

const (
	MenuView = View(iota)
	TimerListView
	TimerListAddView
)

var modeToStrMap = map[View]string{
	MenuView:         "Menu",
	TimerListView:    "Timers",
	TimerListAddView: "Add Timer",
}

type MenuItem struct {
	Label  string
	Action func() tea.Cmd
}

func (m View) String() string {
	return modeToStrMap[m]
}

func SwitchViewCmd(target View, in SwitchViewInput) tea.Cmd {
	return func() tea.Msg {
		return SwitchViewMsg{
			Target: target,
			Input:  in,
		}
	}
}

type SwitchViewMsg struct {
	Target View
	Input  SwitchViewInput
}

type SwitchViewInput interface {
	isSwitchViewInput()
}

// MENU VIEW
type MenuInput struct{}

func (in *MenuInput) isSwitchViewInput() {}

func NewMenuInput() *MenuInput {
	return &MenuInput{}
}

// TIMERLIST VIEW

type TimerListInput struct{}

func (in *TimerListInput) isSwitchViewInput() {}

func NewTimerListInput() *TimerListInput {
	return &TimerListInput{}
}

// TIMERLIST ADD VIEW

type TimerListAddInput struct {
	// Pointer to ID for nullable values
	TimerId *int64
}

func NewTimerListAddInput(timerId *int64) *TimerListAddInput {
	return &TimerListAddInput{
		TimerId: timerId,
	}
}

func (in *TimerListAddInput) isSwitchViewInput() {}
