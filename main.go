package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yhaokeen/v2rayC/ui"
)

func main() {
	// 创建程序，启用鼠标支持和设置初始窗口大小
	p := tea.NewProgram(
		ui.NewModel(),
		tea.WithAltScreen(), // 使用备用屏幕
		tea.WithReportFocus(),
	)

	// 运行程序
	if _, err := p.Run(); err != nil {
		log.Fatal("Error running program:", err)
	}
}
