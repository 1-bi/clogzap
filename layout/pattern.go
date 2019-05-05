package layout

import (
	"github.com/1-bi/clog/loggerzap"
	"go.uber.org/zap/zapcore"
)

// jsonLayout
type patternLayout struct {

	// charset  default use utf-8
	charset string "utf-8"
	lyProps map[string]string
	pattern []byte

	// layout encoder for defone
	encoderConf zapcore.EncoderConfig

	timeFormater *loggerzap.TimeFormater
}

func NewPatternLayout() *patternLayout {
	var layout = new(patternLayout)

	layout.pattern = []byte(`{
	  "level": "debug",
	  "encoding": "layout",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "initialFields": {"foo": "bar"},
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	layout.encoderConf = zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	layout.timeFormater = loggerzap.NewTimeFormater("2006-01-02T15:04:05.000Z0700")

	return layout
}

func (myself *patternLayout) SetTimeFormat(timePattern string) {

	myself.timeFormater.SetPattern(timePattern)
	myself.encoderConf.EncodeTime = myself.timeFormater.CustomTimeEncoder
}

func (myself *patternLayout) SetCharset(newCharset string) {
	myself.charset = newCharset
}

func (myself *patternLayout) SetLayoutProps(props map[string]string) error {
	myself.lyProps = props
	return nil
}

func (myself *patternLayout) SetPattern(newPattern []byte) {
	myself.pattern = newPattern
}

func (myself *patternLayout) BuildEncoder() zapcore.Encoder {
	var encoder = zapcore.NewConsoleEncoder(myself.encoderConf)
	return encoder
}
