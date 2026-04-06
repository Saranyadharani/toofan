package tui

import "github.com/charmbracelet/lipgloss"

// col renders text in a fixed-width column, handling ANSI escape codes correctly.
func col(w int, s string) string {
	return lipgloss.NewStyle().Width(w).Render(s)
}
