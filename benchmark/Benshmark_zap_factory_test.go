package benchmark

import (
	loggercom "github.com/1-bi/log-api"
	loggerzap "github.com/1-bi/log-zap"
	appender "github.com/1-bi/log-zap/appender"
	zaplayout "github.com/1-bi/log-zap/layout"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Benchmark_Zap_Factory_case1_advanced_example(b *testing.B) {
	//b.StopTimer()
	var multiOpts []loggercom.Option
	multiOpts = make([]loggercom.Option, 0)
	// --- construct layout ---
	var jsonLayout = zaplayout.NewJsonLayout()
	// --- set appender
	var consoleAppender = appender.NewConsoleAppender(jsonLayout)

	var loggerOpt1 = loggerzap.NewLoggerOption()
	loggerOpt1.SetLevel("info")
	loggerOpt1.AddAppender(consoleAppender)
	multiOpts = append(multiOpts, loggerOpt1)

	// use new or struct binding
	// create instance from implement
	loggercom.RegisterLoggerFactory(new(loggerzap.ZapFactoryRegister), multiOpts...)

	// --- create logger factory manager

	logger := loggercom.GetLogger("benshmark.test")

	//logger.Debug("debug message for  example", nil)
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("info message for  example", nil)
		//logger.Warn("warn message for  example", nil)
	}
	//logger.Warn("warn message for  example", nil)
	//logger.Error("error  message for  example", nil)

}
