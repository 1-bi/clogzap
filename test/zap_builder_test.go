package test

import (
	"github.com/1-bi/log-api"
	logzap "github.com/1-bi/log-zap"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_case1_base(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetLevel("warn")

	// use new or struct binding
	// create instance from implement
	logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)

	// --- create logger factory manager

}

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_prop_presets_example(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetProperty(logzap.P_PRESETS, logzap.PRESETS_EXAMPLE)

	// use new or struct binding
	// create instance from implement
	logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)

	logger := logapi.GetLogger("test.case 2 ")

	logger.Debug("debug message for  example", nil)
	logger.Info("info message for  example", nil)
	logger.Warn("warn message for  example", nil)
	logger.Error("error  message for  example", nil)
}

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_prop_presets_production(t *testing.T) {

	var lfo = logzap.NewLoggerOption()
	lfo.SetProperty(logzap.P_PRESETS, logzap.PRESETS_PROD)

	// use new or struct binding
	// create instance from implement
	logapi.RegisterLoggerFactory(new(logzap.ZapFactoryRegister), lfo)

	logger := logapi.GetLogger("test.case1")

	logger.Debug("debug message for  prod", nil)
	logger.Info("info message for  prod", nil)
	logger.Warn("warn message for  prod", nil)
	logger.Error("error  message for  prod", nil)

}
