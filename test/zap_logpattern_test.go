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
func Test_LoggberPattern(t *testing.T) {
	var multiOpts = make([]logapi.Option, 0)

	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	//jsonLayout.SetTimeFormat("2006-01-02 15:04:05")
	jsonLayout.SetTimeFormat("2006-01-02 15:04:05 +0800 UTC")
	jsonLayout.SetTimezoneId("UTC")

	//fmt.Println( time.Now().Location() )

	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = logzap.NewLoggerOption()
	loggerOpt1.SetLevel("debug")
	loggerOpt1.AddAppender(consoleAppender)

	multiOpts = append(multiOpts, loggerOpt1)

	//multiOpts = append(multiOpts, loggerOpt2)

	// use new or struct binding
	// create instance from implement
	mainLog, err := logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), multiOpts...)
	if err != nil {
		log.Println(err)
	}

	logapi.InitLoggerPattern([]string{
		"testapp.testmodule.fun1"}, mainLog)

	//logger := lfm.GetLogger()
	logger := logapi.GetLogger("testapp.testmodule.fun1")

	var loggerBean = logapi.NewStructBean()

	logger.Debug("logger bean test case:", loggerBean)

}
