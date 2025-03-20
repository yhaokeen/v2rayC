package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// Log 全局日志对象
	Log *zap.Logger
)

// 初始化日志
func Init(logPath string, level string) {
	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(logPath), 0755); err != nil {
		panic("创建日志目录失败: " + err.Error())
	}

	// 设置日志级别
	var logLevel zapcore.Level
	switch level {
	case "debug":
		logLevel = zapcore.DebugLevel
	case "info":
		logLevel = zapcore.InfoLevel
	case "warn":
		logLevel = zapcore.WarnLevel
	case "error":
		logLevel = zapcore.ErrorLevel
	default:
		logLevel = zapcore.InfoLevel
	}

	// 配置编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 创建Core
	// 同时输出到控制台和文件
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)
	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

	// 打开日志文件
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic("打开日志文件失败: " + err.Error())
	}

	// 创建Core
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), logLevel),
		zapcore.NewCore(fileEncoder, zapcore.AddSync(logFile), logLevel),
	)

	// 创建Logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Log.Info("日志系统初始化成功", zap.String("level", level), zap.String("path", logPath))
}

// Debug 调试日志
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Info 信息日志
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Warn 警告日志
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Error 错误日志
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// Fatal 致命错误日志
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// WithFields 创建带有字段的日志记录器
func WithFields(fields ...zap.Field) *zap.Logger {
	return Log.With(fields...)
}
