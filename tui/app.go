package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type (
	AppModel struct{
		Size Size
		Header HeaderModel
		Screen tea.Model
	}
)

func NewApp() AppModel {
	return AppModel{
		Header: NewHeader(),
		Screen: NewHomeScreen(),
	}
}

func (m AppModel) Init() tea.Cmd {
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd, headerCmd, screenCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Size = NewSizeFromWindow(msg)
	}

	m.Header, headerCmd = m.Header.Update(msg)
	m.Screen, screenCmd = m.Screen.Update(msg)

	return m, tea.Batch(cmd, headerCmd, screenCmd)
}

func (m AppModel) View() string {
	return strings.Join([]string{
		m.Header.View(),
		m.Screen.View(),
	}, "\n")
}
