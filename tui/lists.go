package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	ListsScreenModel struct{
		Size Size
	}

	SwitchListsScreenMsg struct{}
)

var (
	listsScreenStyle = lipgloss.NewStyle().
		Padding(0, 4)
)

func SwitchListsScreen() tea.Msg {
	return SwitchListsScreenMsg{}
}

func NewListsScreen() ListsScreenModel {
	return ListsScreenModel{}
}

func (m ListsScreenModel) Init() tea.Cmd {
	return nil
}

func (m ListsScreenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m *ListsScreenModel) SetSize(width, height int) {
	m.Size = NewSize(width, height)
}

func (m ListsScreenModel) View() string {
	screen := listsScreenStyle.
		Width(m.Size.Width).
		Height(m.Size.Height)

	return screen.Render("lists screen")
}
