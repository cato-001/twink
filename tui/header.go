package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	HeaderModel struct{
		Tabs []HeaderTabModel
		Width int
	}

	HeaderTabModel struct{
		Current bool
		Name string
		Unread int
	}
)

func NewHeader() HeaderModel {
	return HeaderModel{
		Tabs: []HeaderTabModel {
			NewHeaderTab("Home", 10, true),
			NewHeaderTab("Notifications", 3, false),
			NewHeaderTab("Lists", 5, false),
			NewHeaderTab("Private", 0, false),
		},
	}
}

func NewHeaderTab(name string, unread int, current bool) HeaderTabModel {
	return HeaderTabModel{ Name: name, Unread: unread, Current: current }
}

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}

	headerTabStyle = lipgloss.NewStyle().
		Border(tabBorder, true)

	activeHeaderTabStyle = headerTabStyle.
		Border(activeTabBorder, true)
)

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case Size:
		m.Width = msg.Width
	}

	return m, cmd
}

func (m HeaderModel) View() string {
	header := lipgloss.NewStyle().
		Padding(1, 0).
		Align(lipgloss.Center).
		Width(m.Width)

	renderedTabs := make([]string, len(m.Tabs))
	for i, tab := range m.Tabs {
		renderedTabs[i] = tab.View()
	}

	return header.Render(renderedTabs...)
}

func (m HeaderTabModel) Init() tea.Cmd {
	return nil
}

func (m HeaderTabModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m HeaderTabModel) View() string {
	content := lipgloss.JoinHorizontal(lipgloss.Center, m.Name, fmt.Sprintf("(%d)", m.Unread))

	if m.Current {
		return lipgloss.NewStyle().Underline(true).Render(content)
	}

	return content
}
