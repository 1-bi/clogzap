package appender

import (
	loggercom "github.com/1-bi/log-api"
	loggerzap "github.com/1-bi/log-zap"
	"go.uber.org/zap/zapcore"
)

type FileAppender struct {
	outWriter zapcore.WriteSyncer

	errWriter zapcore.WriteSyncer

	innerLayout loggerzap.ZapLayout
}

func NewFileAppender(layout loggerzap.ZapLayout) *FileAppender {

	var fileAppender = new(FileAppender)

	fileAppender.innerLayout = layout

	return fileAppender
}

func (myself *FileAppender) Initialize() {

}

func (myself *FileAppender) SetLayout(layout loggercom.Layout) {
	myself.innerLayout = layout.(loggerzap.ZapLayout)
}

func (myself *FileAppender) GetLayout() loggercom.Layout {
	return myself.innerLayout
}

func (myself *FileAppender) Output() zapcore.WriteSyncer {
	return myself.outWriter
}

func (myself *FileAppender) Error() zapcore.WriteSyncer {
	return myself.errWriter
}

func (myself *FileAppender) GetAppenderName() string {
	return "file"
}

// internal layout method
func (myself *FileAppender) ZapLayout() loggerzap.ZapLayout {
	return myself.innerLayout
}
