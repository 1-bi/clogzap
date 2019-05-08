package test

import (
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"log"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_Option_singleOption1(t *testing.T) {
	//var multiOpts = make([]logapi.Option, 0)

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02T15:04:05.000Z UTC")
	jsonLayout.SetTimezoneId("UTC")

	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = logzap.NewLoggerOption()
	loggerOpt1.SetLevel("debug")
	loggerOpt1.AddAppender(consoleAppender)

	//multiOpts = append(multiOpts, loggerOpt1)

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), loggerOpt1)
	if err != nil {
		log.Println(err)
	}

	//logger := lfm.GetLogger()
	logger := logapi.GetLogger("testapp.testmodule1.testfun1")

	logger.Debug("Info message singleOption1", nil)

}

func Test_Option_multiOptions(t *testing.T) {
	//var multiOpts = make([]logapi.Option, 0)

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02T15:04:05.000Z UTC")
	jsonLayout.SetTimezoneId("UTC")

	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var rootOpt = logzap.NewLoggerOption()
	rootOpt.SetLevel("debug")
	rootOpt.AddAppender(consoleAppender)

	jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02T15:04:05.000Z")
	//jsonLayout.SetTimezoneId("UTC")

	var specOpt1 = logzap.NewLoggerOption()
	specOpt1.SetLoggerPattern("testapp.testmodule1")
	specOpt1.SetLevel("warn")
	specOpt1.AddAppender(appender.NewConsoleAppender(jsonLayout))

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), rootOpt, specOpt1)
	if err != nil {
		log.Println(err)
	}

	//logger := lfm.GetLogger()
	logger := logapi.GetLogger("testapp.testmodule1.testfun1")

	logger.Debug("Debug message singleOption1", nil)
	logger.Info("Info message singleOption1", nil)
	logger.Warn("Warn message singleOption1", nil)
	logger.Error("Error message singleOption1", nil)

}
