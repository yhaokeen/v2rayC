package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yhaokeen/v2rayC/pkg/logger"
	"github.com/yhaokeen/v2rayC/ui"
	"go.uber.org/zap"
)

func main() {
	// 初始化日志，设置日志文件路径和日志级别
	logger.Init("./logs/v2rayc.log", "debug")

	// 记录一条启动日志
	logger.Info("应用启动", zap.String("version", "1.0.0"))

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
