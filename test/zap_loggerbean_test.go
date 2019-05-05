package test

import (
	"github.com/1-bi/log-api"
	logzap "github.com/1-bi/log-zap"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_LoggerBean_Debug(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetProperty(logzap.P_PRESETS, logzap.PRESETS_DEV)

	// use new or struct binding
	// create instance from implement
	logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)

	var loggerBean = logapi.NewStructBean()
	var logger = logapi.GetLogger("loggerbean.test")

	loggerBean.LogString("testStringfield", "logstring filed")
	loggerBean.LogBool("testBoolField", true)

	logger.Debug("logger bean test case:", loggerBean)

}

func Test_BasicCase1_Info(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetProperty(logzap.P_PRESETS, logzap.PRESETS_DEV)

	// use new or struct binding
	// create instance from implement
	logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)

	var loggerBean = logapi.NewStructBean()

	loggerBean.LogString("testStringfield2", "logstring filed")
	loggerBean.LogBool("testBoolField2", false)
	loggerBean.LogFloat32("testfloat32Field2", 32.32)

	logapi.GetLogger("logger bean ").Info("logger bean test case:", loggerBean)
}
