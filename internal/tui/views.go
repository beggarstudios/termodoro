package tui

type View int

const (
	MenuView = View(iota)
)

// TODO wtf is this

type SwitchModeInput interface {
	isSwitchModeInput()
}

type MenuInput struct{}

func (in *MenuInput) isSwitchModeInput() {}
