package utils

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func LogINFO() {
	log.InfoLevelStyle = lipgloss.NewStyle().
		SetString("INFO").
		Padding(0, 1, 0, 1).
		Background(lipgloss.AdaptiveColor{
			Light: "86",
			Dark:  "86",
		}).
		Foreground(lipgloss.Color("0"))
	// Add a custom style for key `err`
	log.KeyStyles["info"] = lipgloss.NewStyle().Foreground(lipgloss.Color("86"))
	log.ValueStyles["info"] = lipgloss.NewStyle().Bold(true)
	log.Info("Whoops!", "info", "kitchen on fire")
}
