package tui

import tea "github.com/charmbracelet/bubbletea"

type (
	NotificationScreenModel struct{
		Size Size
	}

	SwitchNotificationScreenMsg struct{}
)

func SwitchNotificationScreen() tea.Msg {
	return SwitchNotificationScreenMsg{}
}

func NewNotificationScreen() NotificationScreenModel {
	return NotificationScreenModel{}
}

func (m NotificationScreenModel) Init() tea.Cmd {
	return nil
}

func (m NotificationScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl-c":
			cmd = tea.Quit
		case "H":
			cmd = SwitchHomeScreen
		}
	}

	return m, cmd
}

func (m NotificationScreenModel) View() string {
	return "notification screen"
}
