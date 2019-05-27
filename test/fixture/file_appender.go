package fixture

import (
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"github.com/smartystreets/gunit"
	"log"
)

type FileAppenderFixTure struct {
	*gunit.Fixture
}

// Setup
func (myself *FileAppenderFixTure) Setup() {

}

func (myself *FileAppenderFixTure) Teardown() {

}

// TestCase1
func (myself *FileAppenderFixTure) TestCase1_jsonLayout() {

	var multiOpts = make([]logapi.Option, 0)

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var rootLogOpt = logzap.NewLoggerOption()
	rootLogOpt.SetLevel("debug")
	rootLogOpt.AddAppender(consoleAppender)

	multiOpts = append(multiOpts, rootLogOpt)

	var fileAppender = appender.NewFileAppender(jsonLayout)
	fileAppender.SetFileName("/var/log/test.log")

	// --- add file appender ---
	rootLogOpt.AddAppender(fileAppender)

	//var fileLoggerOpt = logzap.NewLoggerOption()
	//fileLoggerOpt.SetLevel("info")
	//fileLoggerOpt.AddAppender(fileAppender)

	//multiOpts = append(multiOpts, fileLoggerOpt)

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), multiOpts...)
	if err != nil {
		log.Println(err)
	}
	//logger := lfm.GetLogger()
	logger := logapi.GetLogger("module")

	logger.Debug("debug message for  example", nil)
	logger.Info("info message for  example", nil)
	logger.Warn("warn message for  example", nil)
	logger.Error("error  message for  example", nil)
}
