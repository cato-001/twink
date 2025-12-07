package tui

import tea "github.com/charmbracelet/bubbletea"

type (
	Size struct{
		Width, Height int
	}

	Sized interface{
		SetSize(width, height int) tea.Model
	}
	SizedWidth interface{
		SetWidth(width int) tea.Model
	}
	SizedHeight interface{
		SetHeight(height int) tea.Model
	}
)

func NewSize(width, height int) Size {
	return Size { Width: width, Height: height }
}

func NewSizeFromWindow(size tea.WindowSizeMsg) Size {
	return Size { Width: size.Width, Height: size.Height }
}
