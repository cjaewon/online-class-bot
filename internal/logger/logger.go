package logger

import "go.uber.org/zap"

var (
	defaultLogger, _ = zap.NewProduction()
	sugar            = defaultLogger.Sugar()
)

// Log returns zap.Logger instance
func Log() *zap.SugaredLogger {
	return sugar
}
