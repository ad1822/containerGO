package utils

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func Logger(level log.Level, status string, message string, bgColor string, fgColor string) {
	// Initialize default styles
	styles := log.DefaultStyles()

	// Set the background and foreground colors for the level message (e.g., SUCCESS with colored background)
	styles.Levels[level] = lipgloss.NewStyle().
		SetString(status). // The "SUCCESS" or status message
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color(bgColor)).
		Foreground(lipgloss.Color(fgColor))

	// Create logger and set styles
	logger := log.New(os.Stderr)
	logger.SetStyles(styles)

	// Log the status and the message with the provided level
	logger.Log(level, status+" "+message)
}
