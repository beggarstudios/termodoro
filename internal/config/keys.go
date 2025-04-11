package config

type Keys struct {
	ForceQuit []string `toml:"force_quit"`
	Exit      []string `toml:"exit"`
	Help      []string `toml:"help"`
	Up        []string `toml:"up"`
	Down      []string `toml:"down"`
	Left      []string `toml:"left"`
	Right     []string `toml:"right"`
}

func DefaultKeys() *Keys {
	return &Keys{
		ForceQuit: []string{"ctrl+c"},
		Exit:      []string{"esc"},
		Help:      []string{"?"},
		Up:        []string{"w"},
		Down:      []string{"s"},
		Left:      []string{"a"},
		Right:     []string{"d"},
	}
}
