package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	HomeScreenModel struct{
		Size Size
	}

	SwitchHomeScreenMsg struct{}
)

var (
	homeScreenStyle = lipgloss.NewStyle().
		Padding(0, 4)
)

func SwitchHomeScreen() tea.Msg {
	return SwitchHomeScreenMsg{}
}

func NewHomeScreen() HomeScreenModel {
	return HomeScreenModel{}
}

func (m HomeScreenModel) Init() tea.Cmd {
	return nil
}

func (m HomeScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	}

	return m, cmd
}

func (m HomeScreenModel) SetSize(width, height int) HomeScreenModel {
	m.Size = NewSize(width, height)
	return m
}

func (m HomeScreenModel) View() string {
	screen := homeScreenStyle.
		Width(m.Size.Width).
		Height(m.Size.Height)

	return screen.Render(
		"HomeScreen",
	)
}
