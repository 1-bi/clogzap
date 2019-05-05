package logzap

import (
	loggercom "github.com/1-bi/log-api"
	"go.uber.org/zap/zapcore"
)

// zapAppender appender struct
type zapAppender interface {
	loggercom.Appender

	// Output output writer
	Output() zapcore.WriteSyncer

	// Error error writer
	Error() zapcore.WriteSyncer

	ZapLayout() ZapLayout
}
