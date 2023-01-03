package common

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
)

// go slog 测试
func init() {
	TimerSetLogProperty()
	// add logs

	slog.Trace("this is a log message")
	slog.Debug("this is a log message")
	slog.Info("this is a log message")
	slog.Notice("this is a log message")
	slog.Warn("this is a log message")
	slog.Error("this is a log message")
	//slog.Fatal("this is a log message")     // 默认打印fatal 就退出程序
}
func TimerSetLogProperty() {

	defer slog.MustFlush()

	// DangerLevels 包含： slog.PanicLevel, slog.ErrorLevel, slog.WarnLevel
	h1 := handler.MustFileHandler("./logs/error.log", handler.WithLogLevels(slog.DangerLevels))

	// NormalLevels 包含： slog.InfoLevel, slog.NoticeLevel, slog.DebugLevel, slog.TraceLevel
	h2 := handler.MustFileHandler("./logs/info.log", handler.WithLogLevels(slog.NormalLevels))

	slog.PushHandler(h1)
	slog.PushHandler(h2)

}
