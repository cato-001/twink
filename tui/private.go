package tui

import tea "github.com/charmbracelet/bubbletea"

type (
	PrivateScreenModel struct{
		Size Size
	}

	SwitchPrivateScreenMsg struct{}
)

func SwitchPrivateScreen() tea.Msg {
	return SwitchPrivateScreenMsg{}
}

func NewPrivateScreen() PrivateScreenModel {
	return PrivateScreenModel{}
}

func (m PrivateScreenModel) Init() tea.Cmd {
	return nil
}

func (m PrivateScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m *PrivateScreenModel) SetSize(width, height int) {
	m.Size = NewSize(width, height)
}

func (m PrivateScreenModel) View() string {
	return "private messages screen"
}
