package context

type focusPanelType int

const (
	mainPanelFocus focusPanelType = iota
	nodePanelFocus
	editPanelFocus
)

type ListType int

const (
	SubListType ListType = iota
	SerListType
)

// AppContext 存储应用全局状态
type AppContext struct {
	FocusPanel  focusPanelType
	CurrentList ListType

	WindowWidth  int
	WindowHeight int
	// 其他全局状态...
}

func NewAppContext() *AppContext {
	return &AppContext{}
}

type Column struct {
	Title string
	Width int
}

type Columns []Column

var SubColumns = Columns{
	{Title: "[ ]", Width: 5},
	{Title: "ID", Width: 10},
	{Title: "域名", Width: 30},
	{Title: "别名", Width: 30},
	{Title: "上次更新时间", Width: 20},
	{Title: "节点数", Width: 10},
	{Title: "操作", Width: 20},
}

var SerColumns = Columns{
	{Title: "[ ]", Width: 5},
	{Title: "ID↑", Width: 10},
	{Title: "节点名", Width: 20},
	{Title: "节点地址", Width: 30},
	{Title: "协议", Width: 30},
	{Title: "时延", Width: 8},
	{Title: "操作", Width: 20},
}
