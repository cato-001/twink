package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	HomeScreenModel struct{
		Size Size
	}

	SwitchHomeScreenMsg struct{}
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

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl-c":
			cmd = tea.Quit
		case "N":
			cmd = SwitchNotificationScreen
		}
	case Size:
		m.Size = msg
	}

	return m, cmd
}

func (m HomeScreenModel) SetSize(width, height int) HomeScreenModel {
	m.Size = NewSize(width, height)
	return m
}

func (m HomeScreenModel) View() string {
	header := lipgloss.NewStyle().
		Align(lipgloss.Center).
		Width(m.Size.Width).
		SetString(fmt.Sprintf("Home\n%s", "@example@mastodon.social"))

	window := lipgloss.NewStyle().
		Width(m.Size.Width).
		Height(m.Size.Height)

	return window.Render(
		header.Render(),
	)
}
