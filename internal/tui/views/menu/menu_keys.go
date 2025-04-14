package menu

import (
	"github.com/charmbracelet/bubbles/key"
)

type menuKeyMap struct {
	Exit   key.Binding
	Help   key.Binding
	Up     key.Binding
	Down   key.Binding
	Select key.Binding
	Quit   key.Binding
}

func (k menuKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k menuKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Select},
		{k.Help, k.Quit},
	}
}

func defaultMenuKeyMap() *menuKeyMap {
	keys := &menuKeyMap{
		Exit:   key.NewBinding(key.WithKeys("esc"), key.WithHelp("esc", "exit")),
		Help:   key.NewBinding(key.WithKeys("h"), key.WithHelp("h", "help")),
		Up:     key.NewBinding(key.WithKeys("up"), key.WithHelp("↑", "move up")),
		Down:   key.NewBinding(key.WithKeys("down"), key.WithHelp("↓", "move down")),
		Select: key.NewBinding(key.WithKeys("enter"), key.WithHelp("↵", "select")),
		Quit:   key.NewBinding(key.WithKeys("q"), key.WithHelp("q", "quit")),
	}

	return keys
}
