package theme

import "github.com/charmbracelet/lipgloss"

var Chameleon = Palette{
	Name:       "chameleon",
	Background: lipgloss.Color("0"),
	Foreground: lipgloss.Color("8"),
	Typed:      lipgloss.Color("7"),
	Error:      lipgloss.Color("1"),
	Cursor:     lipgloss.Color("15"),
	Accent:     lipgloss.Color("6"),
	Success:    lipgloss.Color("2"),
}
