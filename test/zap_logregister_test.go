package test

import (
	"github.com/1-bi/log-api"
	"github.com/1-bi/log-zap"
	"github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"testing"
)

func Test_Default_Case1(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetLevel("info")

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02 15:04:05 +0800 CST")
	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	lfo.AddAppender(consoleAppender)

	// use new or struct binding
	// create instance from implement
	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)
	if err != nil {
		t.Fatal(err)
		return
	}

	// --- create logger factory manager
	logapi.GetLogger("main").Debug("Debug message in default case.", nil)
	logapi.GetLogger("main").Info("Info message in default case.", nil)
	logapi.GetLogger("main").Warn("Warn message in default case.", nil)
	logapi.GetLogger("main").Error("Error message in default case.", nil)
	logapi.GetLogger("main").Fatal("Fatal message in default case.", nil)

}

//  Test_BasicCase1_Debug define bug info
func Test_Default_case2_withoutOption(t *testing.T) {

	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	// --- create logger factory manager
	logapi.GetLogger("main").Debug("Debug message in default case.", nil)
	logapi.GetLogger("main").Info("Info message in default case.", nil)
	logapi.GetLogger("main").Warn("Warn message in default case.", nil)
	logapi.GetLogger("main").Error("Error message in default case.", nil)
	logapi.GetLogger("main").Fatal("Fatal message in default case.", nil)
}

//  Test_BasicCase1_Debug define bug info
func Test_Default_case3_withoutOption(t *testing.T) {

	_, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), nil)
	if err != nil {
		t.Fatal(err)
		return
	}

	// --- create logger factory manager
	logapi.GetLogger("appender").Debug("Debug message in default case.", nil)
	logapi.GetLogger("appender").Info("Info message in default case.", nil)
	logapi.GetLogger("appender").Warn("Warn message in default case.", nil)
	logapi.GetLogger("appender").Error("Error message in default case.", nil)
	logapi.GetLogger("appender").Fatal("Fatal message in default case.", nil)
}
