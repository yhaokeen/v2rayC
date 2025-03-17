package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/yhaokeen/v2rayC/ui/components/header"
	"github.com/yhaokeen/v2rayC/ui/components/list"
	"github.com/yhaokeen/v2rayC/ui/components/tabs"
	"github.com/yhaokeen/v2rayC/ui/context"
)

// 主应用模型
type Model struct {
	ctx         *context.AppContext
	header      header.Model
	tabs        tabs.Model
	list        list.Model
	currentView string
	width       int
	height      int
}

// TODO: 准备把subscription和server的model合并，用一个组件来表示
// 初始化主模型
func NewModel() Model {
	ctx := context.NewAppContext()

	return Model{
		ctx:         ctx,
		header:      header.NewModel(ctx),
		tabs:        tabs.NewModel(ctx),
		list:        list.NewModel(ctx, "subscription"),
		currentView: "subscription",
	}
}

func (m Model) Init() tea.Cmd {
	// 组合所有组件的初始化命令
	return tea.Batch(
		m.header.Init(),
		m.tabs.Init(),
		m.list.Init(),
		tea.EnterAltScreen, // 进入全屏模式
	)
}

// 更新方法
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// 处理按键事件
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	// case tea.MouseMsg:
	// 处理鼠标事件
	// 这里我们不直接处理，而是将事件传递给子组件
	// 子组件会根据自己的区域判断是否处理该事件
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tabs.TabChangedMsg:
		m.currentView = msg.Tab
		m.list = list.NewModel(m.ctx, strings.ToLower(msg.Tab))
	}

	// 更新header
	var headerModel header.Model
	headerModel, cmd = m.header.Update(msg)
	m.header = headerModel
	cmds = append(cmds, cmd)

	// 更新tabs
	var tabsModel tabs.Model
	tabsModel, cmd = m.tabs.Update(msg)
	m.tabs = tabsModel
	cmds = append(cmds, cmd)

	// 根据当前视图更新相应组件
	cmds = append(cmds, m.list.Update(msg)...)

	return m, tea.Batch(cmds...)
}

// 视图渲染
func (m Model) View() string {
	headerView := m.header.View()
	tabsView := m.tabs.View()

	// 根据当前视图选择要显示的内容
	var contentView string
	currentViewLower := strings.ToLower(m.currentView)

	contentView = m.list.View()

	// 状态栏
	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#336699")).
		Padding(0, 1)

	// 这里可以根据实际情况显示不同的状态信息
	statusText := "状态: 已连接 | 延迟: 120ms | 上传: 1.2MB/s | 下载: 5.6MB/s"
	statusBar := statusStyle.Render(statusText)

	// 帮助信息
	helpStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Padding(0, 1)

	helpText := helpStyle.Render("按 q 退出 • 使用 ← → 切换标签 • 使用 ↑ ↓ 导航")

	return lipgloss.JoinVertical(
		lipgloss.Center,
		headerView,
		tabsView,
		contentView,
		statusBar,
		helpText,
	)
}
