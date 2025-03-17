package server

import (
	"strconv"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/yhaokeen/v2rayC/ui/context"
)

type Model struct {
	ctx    *context.AppContext
	table  table.Model
	width  int
	height int
}

func NewModel(ctx *context.AppContext) Model {
	columns := []table.Column{
		{Title: "□", Width: 3},
		{Title: "ID↑", Width: 4},
		{Title: "节点名", Width: 20},
		{Title: "节点地址", Width: 30},
		{Title: "协议", Width: 30},
		{Title: "时延", Width: 8},
		{Title: "操作", Width: 20},
	}

	// 创建基础表格
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

	// 添加行边框
	s.Cell = s.Cell.
		BorderStyle(lipgloss.NormalBorder()).
		BorderBottom(true).
		BorderLeft(false).
		BorderRight(false)

	t.SetStyles(s)

	return Model{
		ctx:   ctx,
		table: t,
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
	rows := make([]table.Row, 0)

	// 添加图片中的示例数据
	testData := []struct {
		id       int
		name     string
		address  string
		protocol string
	}{
		{1, "HKCN | 香港 | 负载均衡", "gtm.dociadd.com:30002", "SS(aes-256-gcm+simple-obfs)"},
		{2, "HKCN | 香港02", "gtm.dociadd.com:30003", "SS(aes-256-gcm+simple-obfs)"},
		{3, "HKCN | 香港03", "gtm.dociadd.com:30012", "SS(aes-256-gcm+simple-obfs)"},
		{4, "HKCN | 香港04", "gtm.dociadd.com:30013", "SS(aes-256-gcm+simple-obfs)"},
		{5, "HKCN | 香港 | 家宽", "gtm.dociadd.com:30015", "SS(aes-256-gcm+simple-obfs)"},
		{6, "CNCN | 台北", "gtm.dociadd.com:30032", "SS(aes-256-gcm+simple-obfs)"},
		{7, "CNCN | 台北 | 家宽", "gtm.dociadd.com:30034", "SS(aes-256-gcm+simple-obfs)"},
		{8, "SGSG | 新加坡01", "gtm.dociadd.com:30042", "SS(aes-256-gcm+simple-obfs)"},
		{9, "SGSG | 新加坡02", "gtm.dociadd.com:30043", "SS(aes-256-gcm+simple-obfs)"},
		{10, "JPJP | 东京01", "gtm.dociadd.com:30052", "SS(aes-256-gcm+simple-obfs)"},
		{11, "JPJP | 东京02", "gtm.dociadd.com:30053", "SS(aes-256-gcm+simple-obfs)"},
		{12, "KRKR | 首尔 | 家宽", "gtm.dociadd.com:30061", "SS(aes-256-gcm+simple-obfs)"},
		{13, "USUS | 芝加哥", "gtm.dociadd.com:30066", "SS(aes-256-gcm+simple-obfs)"},
		{14, "USUS | 洛杉矶", "gtm.dociadd.com:30068", "SS(aes-256-gcm+simple-obfs)"},
		{15, "GBGB | 伦敦", "gtm.dociadd.com:30067", "SS(aes-256-gcm+simple-obfs)"},
		{16, "DEDE | 法兰克福", "gtm.dociadd.com:30081", "SS(aes-256-gcm+simple-obfs)"},
		{17, "TRTR | 土耳其", "gtm.dociadd.com:30111", "SS(aes-256-gcm+simple-obfs)"},
	}

	for _, data := range testData {
		rows = append(rows, table.Row{
			"□",
			strconv.Itoa(data.id),
			data.name,
			data.address,
			data.protocol,
			"--", // 时延，初始为空
			"连接 查看 分享",
		})
	}

	m.table.SetRows(rows)
}
