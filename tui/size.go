package tui

import tea "github.com/charmbracelet/bubbletea"

type Size struct{
	Width, Height int
}

func NewSize(width, height int) Size {
	return Size { Width: width, Height: height }
}

func NewSizeFromWindow(size tea.WindowSizeMsg) Size {
	return Size { Width: size.Width, Height: size.Height }
}
