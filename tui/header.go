package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	HeaderModel struct{
		Size Size
		Tabs []HeaderTabModel
	}

	HeaderTabModel struct{
		Size Size
		Active bool
		Name string
		Key rune
		Unread int
	}
)

func NewHeader() HeaderModel {
	return HeaderModel{
		Tabs: []HeaderTabModel {
			NewHeaderTab("Home", 'H', 10, true),
			NewHeaderTab("Notifications", 'N', 3, false),
			NewHeaderTab("Lists", 'L', 5, false),
			NewHeaderTab("Private", 'P', 0, false),
		},
	}
}

func NewHeaderTab(name string, key rune, unread int, current bool) HeaderTabModel {
	return HeaderTabModel{ Name: name, Key: key, Unread: unread, Active: current }
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
		Padding(0, 1).
		Border(tabBorder, true)

	activeHeaderTabStyle = headerTabStyle.
		Border(activeTabBorder, true)
)

func (m HeaderModel) Init() tea.Cmd {
	return nil
}

func (m HeaderModel) Update(msg tea.Msg) (HeaderModel, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case SwitchHomeScreenMsg:
		m.SetActive(0)
	case SwitchNotificationScreenMsg:
		m.SetActive(1)
	}

	return m, cmd
}

func (m *HeaderModel) SetActive(tabIndex int) {
	for i := range m.Tabs {
		m.Tabs[i].Active = i == tabIndex
	}
}

func (m *HeaderModel) SetWidth(width int) {
	tabWidth := width / len(m.Tabs)
	var maxTabHeight int
	for i := range m.Tabs {
		m.Tabs[i].SetSize(tabWidth)
		maxTabHeight = max(maxTabHeight, m.Tabs[i].Size.Height)
	}
	m.Size = Size{
		Width: width,
		Height: 3+maxTabHeight,
	}
}

func (m HeaderModel) View() string {
	header := lipgloss.NewStyle().
		Padding(1, 0).
		Align(lipgloss.Center).
		Width(m.Size.Width)

	renderedTabs := make([]string, len(m.Tabs))
	for i, tab := range m.Tabs {
		renderedTabs[i] = tab.View()
	}
	joinedTabs := lipgloss.JoinHorizontal(lipgloss.Bottom, renderedTabs...)

	return  header.Render(joinedTabs)
}

func (m HeaderTabModel) Init() tea.Cmd {
	return nil
}

func (m HeaderTabModel) Update(msg tea.Msg) (HeaderTabModel, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

func (m *HeaderTabModel) SetSize(width int) {
	m.Size = Size{
		Width: width,
		Height: 1,
	}
}

func (m HeaderTabModel) View() string {
	contentWidth := m.Size.Width - 4
	content := fmt.Sprintf("[%c] %s - %d", m.Key, m.Name, m.Unread)
	if len(content) > contentWidth {
		content = fmt.Sprintf("[%c] %d", m.Key, m.Unread)
		if len(content) > contentWidth {
			content = fmt.Sprintf("[%c]", m.Key)
		}
	}
	content = lipgloss.PlaceHorizontal(contentWidth, lipgloss.Center, content)

	if m.Active {
		return activeHeaderTabStyle.Render(content)
	}

	return headerTabStyle.Render(content)
}
