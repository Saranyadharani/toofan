package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"toofan/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.New(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
