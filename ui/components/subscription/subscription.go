package subscription

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/yhaokeen/v2rayC/ui/context"
)

type Model struct {
	ctx    *context.AppContext
	table  table.Model
	width  int
	height int
}

// 定义消息类型
type SubscriptionLoadedMsg struct {
	Subscriptions []context.Subscription
}

func NewModel(ctx *context.AppContext) Model {
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "域名", Width: 30},
		{Title: "别名", Width: 30},
		{Title: "上次更新时间", Width: 20},
		{Title: "节点数", Width: 10},
		{Title: "操作", Width: 20},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	return Model{
		ctx:   ctx,
		table: t,
	}
}

func (m Model) Init() tea.Cmd {
	// TODO: 从文件中加载订阅数据

	return nil
}

// 加载订阅数据的命令
func loadSubscriptionsCmd(m Model) tea.Cmd {
	return func() tea.Msg {
		// 这里可以是异步的 HTTP 请求或数据库查询
		subs := m.ctx.Subscriptions
		return SubscriptionLoadedMsg{Subscriptions: subs}
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SubscriptionLoadedMsg:
		// 处理加载完成的数据
		m.ctx.Subscriptions = msg.Subscriptions
		m.UpdateRows()
	case tea.KeyMsg:
		if msg.String() == "r" {
			// 按 r 键刷新数据
			return m, loadSubscriptionsCmd(m)
		}
	}
	return m, nil
}

func (m Model) View() string {
	m.UpdateRows() // 确保显示最新数据
	return m.table.View()
}

// 更新表格数据
func (m *Model) UpdateRows() {
	rows := make([]table.Row, 0)
	for _, sub := range m.ctx.Subscriptions {
		rows = append(rows, table.Row{
			strconv.Itoa(sub.ID),
			sub.Name,
			"", // 别名
			sub.LastUpdate,
			strconv.Itoa(sub.NodeCount),
			"更新 修改 分享",
		})
	}
	m.table.SetRows(rows)
}
