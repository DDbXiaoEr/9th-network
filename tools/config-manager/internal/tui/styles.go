package tui

import "github.com/charmbracelet/lipgloss"

var (
	colorAccent  = lipgloss.Color("45")
	colorAccent2 = lipgloss.Color("213")
	colorDim     = lipgloss.Color("244")
	colorText    = lipgloss.Color("252")
	colorWarn    = lipgloss.Color("214")
	colorErr     = lipgloss.Color("203")
	colorOK      = lipgloss.Color("84")

	titleStyle   = lipgloss.NewStyle().Bold(true).Foreground(colorAccent)
	headingStyle = lipgloss.NewStyle().Bold(true).Foreground(colorAccent2)
	dimStyle     = lipgloss.NewStyle().Foreground(colorDim)
	cursorStyle  = lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	labelStyle   = lipgloss.NewStyle().Foreground(colorText)
	keyStyle     = lipgloss.NewStyle().Foreground(colorAccent)
	okStyle      = lipgloss.NewStyle().Foreground(colorOK)
	errStyle     = lipgloss.NewStyle().Foreground(colorErr)
	helpStyle    = lipgloss.NewStyle().Foreground(colorDim)
	badgeStyle   = lipgloss.NewStyle().Foreground(colorWarn)
	boxStyle     = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colorDim).Padding(0, 1)
)
