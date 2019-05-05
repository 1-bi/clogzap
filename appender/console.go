package appender

import (
	"fmt"
	loggercom "github.com/1-bi/log-api"
	loggerzap "github.com/1-bi/log-zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ConsoleAppender struct {

	// See Open for details.
	OutputPaths []string `json:"outputPaths" yaml:"outputPaths"`
	// ErrorOutputPaths is a list of URLs to write internal logger errors to.
	// The default is standard error.
	//
	// Note that this setting only affects internal errors; for sample code that
	// sends error-level logs to a different location from info- and debug-level
	// logs, see the package-level AdvancedConfiguration example.
	ErrorOutputPaths []string `json:"errorOutputPaths" yaml:"errorOutputPaths"`

	outWriter zapcore.WriteSyncer

	errWriter zapcore.WriteSyncer

	innerLayout loggerzap.ZapLayout
}

// NewConsoleAppender public constructer
func NewConsoleAppender(layout loggerzap.ZapLayout) *ConsoleAppender {

	// check the layout existe or not
	if layout == nil {
		// --- use default layout ----
	}

	var zlayout = layout

	var consoleAppender = new(ConsoleAppender)
	consoleAppender.OutputPaths = []string{"stderr"}
	consoleAppender.ErrorOutputPaths = []string{"stderr"}
	consoleAppender.SetLayout(zlayout)

	sink, errSink, err := consoleAppender.openSinks()
	consoleAppender.outWriter = sink
	consoleAppender.errWriter = errSink

	//consoleAppender.outWriter = zapcore.Lock(os.Stdout)
	//consoleAppender.errWriter = zapcore.Lock(os.Stderr)
	if err != nil {

		fmt.Println(err)

	}

	return consoleAppender
}

func (myself *ConsoleAppender) Initialize() {

}

func (myself *ConsoleAppender) Output() zapcore.WriteSyncer {
	return myself.outWriter
}

func (myself *ConsoleAppender) Error() zapcore.WriteSyncer {
	return myself.errWriter
}

func (myself *ConsoleAppender) GetAppenderName() string {
	return "console"
}

func (myself *ConsoleAppender) SetLayout(layout loggercom.Layout) {
	myself.innerLayout = layout.(loggerzap.ZapLayout)
}

func (myself *ConsoleAppender) GetLayout() loggercom.Layout {
	return myself.innerLayout
}

// internal layout method
func (myself *ConsoleAppender) ZapLayout() loggerzap.ZapLayout {
	return myself.innerLayout
}

func (myself *ConsoleAppender) openSinks() (zapcore.WriteSyncer, zapcore.WriteSyncer, error) {
	sink, closeOut, err := zap.Open(myself.OutputPaths...)
	if err != nil {
		return nil, nil, err
	}
	errSink, _, err := zap.Open(myself.ErrorOutputPaths...)
	if err != nil {
		closeOut()
		return nil, nil, err
	}
	return sink, errSink, nil
}
