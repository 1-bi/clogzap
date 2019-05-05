package logzap

import (
	loggercom "github.com/1-bi/log-api"
	"go.uber.org/zap/zapcore"
)

// zapLayout custom layout
type ZapLayout interface {
	loggercom.Layout

	BuildEncoder() zapcore.Encoder
}
