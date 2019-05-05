package logzap

import (
	logapi "github.com/1-bi/log-api"
	"go.uber.org/zap"
)

/**
 * define logger implement object
 */

type logger struct {
	zaplogger    *zap.Logger
	runtimeLevel byte
}

func (log *logger) setZaplogger(logInst *zap.Logger) {
	log.zaplogger = logInst
}

func (log *logger) IsDebugEnabled() bool {
	return log.runtimeLevel == logapi.DEBUG
}

func (log *logger) IsInfoEnabled() bool {
	return log.runtimeLevel == logapi.INFO
}

func (log *logger) IsWarnEnabled() bool {
	return log.runtimeLevel == logapi.WARN
}

func (log *logger) IsErrorEnabled() bool {
	return log.runtimeLevel == logapi.ERROR
}

// Debug debug logger message object
func (log *logger) Debug(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {

		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.Debug(msg, zab.convertToFields()...)

	} else {
		log.zaplogger.Debug(msg)
	}
}

func (log *logger) Info(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.Info(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.Info(msg)
	}
}

func (log *logger) Warn(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.Warn(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.Warn(msg)
	}

}

func (log *logger) Error(msg string, msgObj logapi.StructBean) {
	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.Error(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.Error(msg)
	}

}

// ==================================
//     Private method
// ==================================
