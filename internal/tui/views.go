package tui

type View int

const (
	MenuView = View(iota)
)

type SwitchViewMsg struct {
	Target View
	Input  SwitchViewInput
}

// TODO wtf is this

type SwitchViewInput interface {
	isSwitchViewInput()
}

type MenuInput struct{}

func (in *MenuInput) isSwitchViewInput() {}

func NewMenuInput() *MenuInput {
	return &MenuInput{}
}
