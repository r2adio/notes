package altscreentoggle

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	keywordStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("204")).Background(lipgloss.Color("235"))
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type model struct {
	altscreen  bool
	quitting   bool
	suspending bool
}

func (m model) Init() tea.Cmd {}
