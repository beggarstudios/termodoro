package menu

import (
	"termodoro/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func navigateToTimers() tea.Cmd {
	return tui.SwitchViewCmd(tui.TimerListView, tui.NewTimerListInput())
}

func navigateToSettings() tea.Cmd {
	return nil
}

func navigateToRewards() tea.Cmd {
	return nil
}

func quitApp() tea.Cmd {
	return tea.Quit
}
