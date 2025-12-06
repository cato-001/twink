package tui

import tea "github.com/charmbracelet/bubbletea"

func Start(args Args) error {
	app := NewApp()
	program := tea.NewProgram(app, tea.WithAltScreen())
	_, err := program.Run()
	return err
}
