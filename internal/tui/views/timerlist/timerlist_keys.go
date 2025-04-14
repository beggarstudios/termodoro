package timerlist

import (
	"github.com/charmbracelet/bubbles/key"
)

type timerlistKeyMap struct {
	Exit   key.Binding
	Help   key.Binding
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Quit   key.Binding
	New    key.Binding
	Edit   key.Binding
}

func (k timerlistKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.New}
}

func (k timerlistKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select},
		{k.New, k.Edit, k.Help, k.Quit},
	}
}

func defaultTimerListKeyMap() *timerlistKeyMap {
	keys := &timerlistKeyMap{
		Exit:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		Help:   key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "help")),
		Up:     key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
		Down:   key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
		Select: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "start timer")),
		Quit:   key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
		New:    key.NewBinding(key.WithKeys("n"), key.WithHelp("n", "new")),
		Edit:   key.NewBinding(key.WithKeys("e"), key.WithHelp("e", "edit")),
	}

	return keys
}

// ADD SCREEN KEYS

type timerlistAddKeyMap struct {
	Exit key.Binding
	Help key.Binding
	Up   key.Binding
	Down key.Binding
}

func (k timerlistAddKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Help}
}

func (k timerlistAddKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Help},
		{k.Exit},
	}
}

func defaultTimerListAddKeyMap() *timerlistAddKeyMap {
	keys := &timerlistAddKeyMap{
		Exit: key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		Help: key.NewBinding(key.WithKeys("h"), key.WithHelp("ctrl+h", "help")),
		Up:   key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
		Down: key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
	}

	return keys
}
