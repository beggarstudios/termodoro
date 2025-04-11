package main

import (
	"fmt"
	"termodoro/internal/config"
	"termodoro/internal/data"
	"termodoro/internal/tui"
	"termodoro/internal/tui/root"

	tea "github.com/charmbracelet/bubbletea"
)

type StartCmd struct{}

func (c *StartCmd) Run(params *StartupParameters) error {
	return launchStarter(params, tui.MenuView, tui.NewMenuInput())
}

func launchStarter(globals *StartupParameters, starterView tui.View, switchIn tui.SwitchViewInput) error {
	db, err := data.InitializeStore(globals.DB)
	if err != nil {
		return fmt.Errorf("opening database: %w", err)
	}

	cfg, err := config.GetConfig(globals.Config)
	if err != nil {
		return fmt.Errorf("getting config: %w", err)
	}

	model, err := root.NewRootModel(
		root.NewInput(starterView, switchIn, db, cfg),
	)
	if err != nil {
		return fmt.Errorf("creating starter model: %w", err)
	}

	exitModel, err := tea.NewProgram(model, tea.WithAltScreen()).Run()
	if err != nil {
		return fmt.Errorf("failed to run program: %w", err)
	}

	typedExitModel, ok := exitModel.(*root.RootModel)
	if !ok {
		return fmt.Errorf("failed to assert exit model type: %w", err)
	}

	if err = typedExitModel.ExitError; err != nil {
		return fmt.Errorf("starter model exited with an error: %w", err)
	}

	return nil
}
