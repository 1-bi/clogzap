package log_zap

import (
	logapi "github.com/1-bi/log-api"
)

// runtimeOption
type runtimeOption struct {
	level  string "info"
	props  map[string]string
	layout logapi.Layout

	// define appender
	appenders map[string]zapAppender
}

func NewLoggerOption() *runtimeOption {
	var o = new(runtimeOption)
	o.props = make(map[string]string, 0)
	o.appenders = make(map[string]zapAppender, 0)
	return o
}

func (myself *runtimeOption) GetLevel() string {
	return myself.level
}

func (myself *runtimeOption) SetLevel(newLevel string) {
	myself.level = newLevel
}

func (myself *runtimeOption) SetProperty(key string, val string) {
	myself.props[key] = val
}

func (myself *runtimeOption) DelProperty(key string) {
	delete(myself.props, key)
}

func (myself *runtimeOption) GetProperties() map[string]string {
	return myself.props
}

func (myself *runtimeOption) GetAppenders() map[string]logapi.Appender {
	var commonAppenderMap = make(map[string]logapi.Appender, 0)

	for k, a := range myself.appenders {
		commonAppenderMap[k] = a
	}
	return commonAppenderMap
}

func (myself *runtimeOption) AddAppender(appender zapAppender) {
	myself.appenders[appender.GetAppenderName()] = appender
}

func (myself *runtimeOption) getAppenderMap() map[string]zapAppender {
	return myself.appenders
}
