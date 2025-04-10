package main

import (
	"termodoro/internal/tui"
)

type StartCmd struct{}

func (c *StartCmd) Run(globals *StartupParameters) error {
	return launchStarter(globals, tui.MenuView, tui.NewMenuInput())
}

func launchStarter(globals *StartupParameters, starterView tui.View, switchIn tui.SwitchModeInput) error {
	// db, err := data.NewDB(globals.DB)
	// if err != nil {
	// 	return fmt.Errorf("opening database: %w", err)
	// }

	// cfg, err := config.GetConfig(globals.Config)
	// if err != nil {
	// 	return fmt.Errorf("getting config: %w", err)
	// }

	// model, err := starter.NewModel(
	// 	starter.NewInput(starterMode, switchIn, db, cfg),
	// )
	// if err != nil {
	// 	return fmt.Errorf("creating starter model: %w", err)
	// }

	// exitModel, err := tea.NewProgram(model, tea.WithAltScreen()).Run()
	// if err != nil {
	// 	return fmt.Errorf("failed to run program: %w", err)
	// }

	// typedExitModel, ok := exitModel.(*starter.Model)
	// if !ok {
	// 	return fmt.Errorf("failed to assert exit model type: %w", err)
	// }

	// if err = typedExitModel.ExitError; err != nil {
	// 	return fmt.Errorf("starter model exited with an error: %w", err)
	// }

	return nil
}
