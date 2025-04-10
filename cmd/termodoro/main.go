package main

import (
	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
)

type CLI struct {
	StartupParameters

	Menu StartCmd `cmd:"" help:"Start in the menu" default:"1"`
}

type StartupParameters struct {
	Config string `help:"Path to config file. Leaving this empty will use default XDG directory." default:""`
	DB     string `help:"Path to the database file. Leaving this empty will use the default XDG directory." default:""`
}

func main() {
	cli := CLI{}
	ctx := kong.Parse(&cli,
		kong.Name("termodoro"),
		kong.Description("A pomodoro timer TUI written in Go"),
		kong.UsageOnError(),
	)

	if err := handleDefaultGlobals(&cli.StartupParameters); err != nil {
		ctx.FatalIfErrorf(err)
	}

	// Call the Run() method of the selected parsed command.
	err := ctx.Run(&cli.StartupParameters)
	ctx.FatalIfErrorf(err)

	// store := &Store{}

	// if err := store.Init(); err != nil {
	// 	fmt.Printf("Error initializing store: %v\n", err)
	// }

	// p := tea.NewProgram(newModel(store))
	// if _, err := p.Run(); err != nil {
	// 	fmt.Printf("Error running program: %v\n", err)
	// 	os.Exit(1)
	// }
}

func handleDefaultGlobals(g *StartupParameters) error {
	if g.Config == "" {
		var err error
		g.Config, err = xdg.ConfigFile("./tetrigo/config.toml")
		if err != nil {
			return err
		}
	}
	if g.DB == "" {
		var err error
		g.DB, err = xdg.DataFile("./tetrigo/tetrigo.db")
		if err != nil {
			return err
		}
	}
	return nil
}
