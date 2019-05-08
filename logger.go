package logzap

import (
	"github.com/1-bi/log-api"
	"go.uber.org/zap"
)

/**
 * define logger implement object
 */

type logger struct {
	zaplogger    *zap.Logger
	name         string
	runtimeLevel byte
	parentLogger logapi.Logger
	additivity   bool
}

func (log *logger) GetName() string {
	return log.name
}

func (log *logger) SetParentLogger(parentLogger logapi.Logger) {
	log.parentLogger = parentLogger
}

func (log *logger) GetParentLogger() logapi.Logger {
	return log.parentLogger
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

func (log *logger) IsFatalEnabled() bool {
	return log.runtimeLevel == logapi.FATAL
}

// Debug debug logger message object
func (log *logger) Debug(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Debug(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Debug(msg)
	}

	if log.additivity && log.parentLogger != nil {
		log.parentLogger.Debug(msg, msgObj)
	}

}

func (log *logger) Info(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Info(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Info(msg)
	}

	if log.additivity && log.parentLogger != nil {
		log.parentLogger.Info(msg, msgObj)
	}
}

func (log *logger) Warn(msg string, msgObj logapi.StructBean) {

	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Warn(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Warn(msg)
	}
	if log.additivity && log.parentLogger != nil {
		log.parentLogger.Warn(msg, msgObj)
	}
}

func (log *logger) Error(msg string, msgObj logapi.StructBean) {
	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Error(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Error(msg)
	}

	if log.additivity && log.parentLogger != nil {
		log.parentLogger.
			Error(msg, msgObj)
	}

}

func (log *logger) Fatal(msg string, msgObj logapi.StructBean) {
	// --- convert zap field ----
	if msgObj != nil {
		zab := msgObj.(*zapLoggerBean)
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Fatal(msg, zab.convertToFields()...)
	} else {
		log.zaplogger.
			With(zap.String("loggerName", log.name)).
			Fatal(msg)
	}
	if log.additivity && log.parentLogger != nil {
		log.parentLogger.Fatal(msg, msgObj)
	}

}

// ==================================
//     Private method
// ==================================
