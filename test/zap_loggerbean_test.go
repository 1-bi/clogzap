package test

import (
	"github.com/1-bi/clog/loggercom"
	"github.com/1-bi/clog/loggerzap"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_LoggerBean_Debug(t *testing.T) {

	var lfm loggercom.LoggerFactory
	var lfo = loggerzap.NewLoggerOption()
	lfo.SetProperty(loggerzap.P_PRESETS, loggerzap.PRESETS_DEV)

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	var loggerBean = lfm.NewLoggerBean()

	loggerBean.LogString("testStringfield", "logstring filed")
	loggerBean.LogBool("testBoolField", true)

	lfm.GetLogger().Debug("logger bean test case:", loggerBean)

}

func Test_BasicCase1_Info(t *testing.T) {
	var lfm loggercom.LoggerFactory
	var lfo = loggerzap.NewLoggerOption()
	lfo.SetProperty(loggerzap.P_PRESETS, loggerzap.PRESETS_DEV)

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	var loggerBean = lfm.NewLoggerBean()

	loggerBean.LogString("testStringfield2", "logstring filed")
	loggerBean.LogBool("testBoolField2", false)
	loggerBean.LogFloat32("testfloat32Field2", 32.32)

	lfm.GetLogger().Info("logger bean test case:", loggerBean)
}
