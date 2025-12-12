package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	AppModel struct{
		Size Size
		Header HeaderModel
		Screen tea.Model
		Screens []tea.Model
	}
)

func NewApp() AppModel {
	screens := []tea.Model{
		NewHomeScreen(),
		NewNotificationScreen(),
		NewListsScreen(),
		NewPrivateScreen(),
	}
	return AppModel{
		Header: NewHeader(),
		Screen: screens[0],
		Screens: screens,
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, headerCmd, screenCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl-c":
			cmd = tea.Quit
		case "H":
			cmd = SwitchHomeScreen
		case "N":
			cmd = SwitchNotificationScreen
		case "L":
			cmd = SwitchListsScreen
		case "P":
			cmd = SwitchPrivateScreen
		}
	case tea.WindowSizeMsg:
		m.SetSize(msg.Width, msg.Height)
	case SwitchHomeScreenMsg:
		m.Screen = m.Screens[0]
	case SwitchNotificationScreenMsg:
		m.Screen = m.Screens[1]
	case SwitchListsScreenMsg:
		m.Screen = m.Screens[2]
	case SwitchPrivateScreenMsg:
		m.Screen = m.Screens[3]
	}

	m.Header, headerCmd = m.Header.Update(msg)
	m.Screen, screenCmd = m.Screen.Update(msg)

	return m, tea.Batch(cmd, headerCmd, screenCmd)
}

func (m *AppModel) SetSize(width, height int) {
	m.Size = NewSize(width, height)
	m.Header.SetWidth(width)
	screenHeight := height - m.Header.Size.Height
	switch screen := m.Screen.(type) {
	case Sized:
		m.Screen = screen.SetSize(width, screenHeight)
	case SizedWidth:
		m.Screen = screen.SetWidth(width)
	case SizedHeight:
		m.Screen = screen.SetHeight(screenHeight)
	}
}

func (m AppModel) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.Header.View(),
		m.Screen.View(),
	)
}
