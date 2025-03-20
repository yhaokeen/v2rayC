package list

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yhaokeen/v2rayC/pkg/logger"
	"github.com/yhaokeen/v2rayC/ui/context"
	"go.uber.org/zap"
)

type Model struct {
	ctx    *context.AppContext
	table  table.Model
	width  int
	height int
	mode   string // "subscription" 或 "server"
}

func NewModel(ctx *context.AppContext, mode string) Model {
	logger.Debug("创建列表模型", zap.String("mode", mode))

	var columns []table.Column

	// 根据模式设置不同的列
	if mode == "subscription" {
		columns = []table.Column{
			{Title: "[ ]", Width: 5},
			{Title: "ID", Width: 10},
			{Title: "域名", Width: 30},
			{Title: "别名", Width: 30},
			{Title: "上次更新时间", Width: 20},
			{Title: "节点数", Width: 10},
			{Title: "操作", Width: 20},
		}
	} else { // server 模式
		columns = []table.Column{
			{Title: "[ ]", Width: 5},
			{Title: "ID↑", Width: 10},
			{Title: "节点名", Width: 20},
			{Title: "节点地址", Width: 30},
			{Title: "协议", Width: 30},
			{Title: "时延", Width: 8},
			{Title: "操作", Width: 20},
		}
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)

	// 设置表格样式
	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderTop(true).
		BorderLeft(false).
		BorderRight(false).
		Bold(true)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#ffffff")).
		Background(lipgloss.Color("#666666")).
		Bold(true)

	s.Cell = s.Cell.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderLeft(false).
		BorderRight(false)

	t.SetStyles(s)

	return Model{
		ctx:   ctx,
		table: t,
		mode:  mode,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.table.SetWidth(msg.Width)
		m.table.SetHeight(msg.Height - 4)
		logger.Debug("调整列表大小", zap.Int("width", msg.Width), zap.Int("height", msg.Height-4))
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	m.UpdateRows()
	return m.table.View()
}

// 更新表格数据
func (m *Model) UpdateRows() {
	logger.Debug("更新表格数据", zap.String("mode", m.mode))
	rows := make([]table.Row, 0)

	if m.mode == "subscription" {
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
	} else { // server 模式
		// 添加服务器列表数据
		testData := []struct {
			id       int
			name     string
			address  string
			protocol string
		}{
			{1, "HKCN | 香港 | 负载均衡", "gtm.dociadd.com:30002", "SS(aes-256-gcm+simple-obfs)"},
			// ... 其他测试数据 ...
		}

		for _, data := range testData {
			rows = append(rows, table.Row{
				"□",
				strconv.Itoa(data.id),
				data.name,
				data.address,
				data.protocol,
				"--",
				"连接 查看 分享",
			})
		}
	}

	m.table.SetRows(rows)
}

func (m Model) loadRows() {

}

// 根据ListType获取表头
func (m Model) initTableColumns() {
	var columns []table.Column
	switch m.ctx.CurrentList {
	case context.SubListType:
		// 使用subColumns创建表头
		for _, column := range context.SubColumns {
			columns = append(columns, table.Column{Title: column.Title, Width: column.Width})
		}
	case context.SerListType:
		for _, column := range context.SerColumns {
			columns = append(columns, table.Column{Title: column.Title, Width: column.Width})
		}
	default:
		panic("invalid list type")
	}
	m.table.SetColumns(columns)
}
