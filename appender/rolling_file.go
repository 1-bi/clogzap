package appender

import (
	loggercom "github.com/1-bi/log-api"
	loggerzap "github.com/1-bi/log-zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RollingFileAppender struct {
	outWriter zapcore.WriteSyncer

	errWriter zapcore.WriteSyncer

	innerLayout loggerzap.ZapLayout

	fileName string

	maxSize    int
	maxBackups int
	maxAge     int
}

func NewFileAppender(layout loggerzap.ZapLayout) *RollingFileAppender {

	var fileAppender = new(RollingFileAppender)

	fileAppender.innerLayout = layout

	fileAppender.maxAge = 30 // days
	fileAppender.maxBackups = 3
	fileAppender.maxAge = 500 // MB

	return fileAppender
}

func (myself *RollingFileAppender) Initialize() {

	myself.outWriter = zapcore.AddSync(&lumberjack.Logger{
		Filename:   myself.fileName,
		MaxSize:    myself.maxSize, // megabytes, MB
		MaxBackups: myself.maxBackups,
		MaxAge:     myself.maxAge, //days
	})

}

func (myself *RollingFileAppender) SetFileName(fileName string) {
	myself.fileName = fileName
}

// SetMaxSize unit MB ,default is 500MB
func (myself *RollingFileAppender) SetMaxSize(size int) {
	myself.maxSize = size
}

func (myself *RollingFileAppender) SetMaxBackups(backups int) {
	myself.maxBackups = backups
}

// SetMaxAge unit days , default 30
func (myself *RollingFileAppender) SetMaxAge(age int) {
	myself.maxAge = age
}

func (myself *RollingFileAppender) SetLayout(layout loggercom.Layout) {
	myself.innerLayout = layout.(loggerzap.ZapLayout)
}

func (myself *RollingFileAppender) GetLayout() loggercom.Layout {
	return myself.innerLayout
}

func (myself *RollingFileAppender) Output() zapcore.WriteSyncer {
	return myself.outWriter
}

func (myself *RollingFileAppender) Error() zapcore.WriteSyncer {
	return myself.errWriter
}

func (myself *RollingFileAppender) GetAppenderName() string {
	return "file"
}

// internal layout method
func (myself *RollingFileAppender) ZapLayout() loggerzap.ZapLayout {
	return myself.innerLayout
}
