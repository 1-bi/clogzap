package layout

import (
	loggerzap "github.com/1-bi/log-zap"
	"go.uber.org/zap/zapcore"
)

// jsonLayout
type jsonLayout struct {

	// charset  default use utf-8
	charset string
	lyProps map[string]string
	pattern []byte

	// layout encoder for defone
	encoderConf zapcore.EncoderConfig

	timeFormater *loggerzap.TimeFormater
}

func NewJsonLayout() *jsonLayout {
	var layout = new(jsonLayout)

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
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	layout.timeFormater = loggerzap.NewTimeFormater("2006-01-02T15:04:05.000Z0700")

	return layout
}

func (myself *jsonLayout) SetTimeFormat(timePattern string) {

	myself.timeFormater.SetPattern(timePattern)
	myself.encoderConf.EncodeTime = myself.timeFormater.CustomTimeEncoder
}

func (myself *jsonLayout) SetTimezoneId(timezoneId string) {
	myself.timeFormater.SetTimeZone(timezoneId)
}

func (myself *jsonLayout) SetCharset(newCharset string) {
	myself.charset = newCharset
}

func (myself *jsonLayout) SetLayoutProps(props map[string]string) error {
	myself.lyProps = props
	return nil
}

func (myself *jsonLayout) SetPattern(newPattern []byte) {
	myself.pattern = newPattern
}

func (myself *jsonLayout) BuildEncoder() zapcore.Encoder {
	var encoder = zapcore.NewJSONEncoder(myself.encoderConf)
	return encoder
}
