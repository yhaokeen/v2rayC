package header

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yhaokeen/v2rayC/ui/context"
)

type Model struct {
	ctx    *context.AppContext
	help   help.Model
	width  int
	height int
}

func NewModel(ctx *context.AppContext) Model {
	return Model{
		ctx:  ctx,
		help: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#1B1D1E")).
		Width(m.width).
		Align(lipgloss.Left).
		Padding(0, 1)

	return style.Render("V2RayA")
}

// 其他方法省略...
