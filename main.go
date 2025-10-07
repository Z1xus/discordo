package main

import (
	"log/slog"

	"github.com/ayn2op/discordo/internal/root"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(root.NewModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		slog.Error("failed to run program", "err", err)
	}
}
