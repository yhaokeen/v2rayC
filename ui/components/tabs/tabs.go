package tabs

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yhaokeen/v2rayC/ui/context"
)

type TabChangedMsg struct {
	Tab string
}

type Model struct {
	ctx       *context.AppContext
	Tabs      []string
	ActiveTab int
	width     int
}

func NewModel(ctx *context.AppContext) Model {
	// 初始固定标签
	tabs := []string{"SUBSCRIPTION", "SERVER"}

	// 添加订阅标签
	for _, sub := range ctx.Subscriptions {
		tabs = append(tabs, sub.Name)
	}

	return Model{
		ctx:       ctx,
		Tabs:      tabs,
		ActiveTab: 0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "left":
			if m.ActiveTab > 0 {
				m.ActiveTab--
				return m, func() tea.Msg {
					return TabChangedMsg{Tab: m.Tabs[m.ActiveTab]}
				}
			}
		case "right":
			if m.ActiveTab < len(m.Tabs)-1 {
				m.ActiveTab++
				return m, func() tea.Msg {
					return TabChangedMsg{Tab: m.Tabs[m.ActiveTab]}
				}
			}
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
	}
	return m, nil
}

func (m Model) View() string {
	// 定义基础样式
	baseStyle := lipgloss.NewStyle().
		Padding(0, 2).
		MarginRight(0)

	// 定义激活状态的样式
	activeStyle := baseStyle.
		Foreground(lipgloss.Color("#FFE4B5")). // 黑色文字
		// Background(lipgloss.Color("#FFE4B5")). // 浅黄色背景
		Bold(true)

	// 定义未激活状态的样式
	inactiveStyle := baseStyle.
		Foreground(lipgloss.Color("#666666")). // 灰色文字
		// Background(lipgloss.Color("#FFFFFF")). // 白色背景
		Bold(false)

	separatorStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#4293f5")).
		Bold(true)

	// 创建容器样式（用于整个tab组）
	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).            // 圆角边框
		BorderForeground(lipgloss.Color("#CCCCCC")). // 边框颜色
		Padding(0).
		BorderTop(true).
		BorderBottom(true).
		BorderLeft(true).
		BorderRight(true)

	var renderedTabs []string
	renderedTabs = append(renderedTabs, separatorStyle.Render("|"))
	for i, tab := range m.Tabs {
		var style lipgloss.Style
		if i == m.ActiveTab {
			style = activeStyle
		} else {
			style = inactiveStyle
		}
		renderedTabs = append(renderedTabs, style.Render(tab), separatorStyle.Render("|"))
	}

	// 将所有标签水平连接
	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)

	// 将连接后的标签放入容器中
	return containerStyle.Render(row)
}

// 添加新的订阅时更新标签
func (m *Model) AddSubscriptionTab(name string) {
	m.Tabs = append(m.Tabs, name)
}

// 其他方法省略...
