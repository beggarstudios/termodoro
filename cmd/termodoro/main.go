package main

import (
	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
)

// Build command: go build -o termodoro.exe .

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

	err := ctx.Run(&cli.StartupParameters)
	ctx.FatalIfErrorf(err)
}

func handleDefaultGlobals(g *StartupParameters) error {
	if g.Config == "" {
		var err error
		g.Config, err = xdg.ConfigFile("./termodoro/config.toml")
		if err != nil {
			return err
		}
	}
	if g.DB == "" {
		var err error
		g.DB, err = xdg.DataFile("./termodoro/termodoro.db")
		if err != nil {
			return err
		}
	}
	return nil
}
