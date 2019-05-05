package test

import (
	"github.com/1-bi/clog/loggercom"
	"github.com/1-bi/clog/loggerzap"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_case1_base(t *testing.T) {

	var lfm loggercom.LoggerFactory
	var lfo = loggerzap.NewLoggerOption()
	lfo.SetLevel("warn")

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	// --- create logger factory manager
	if lfm == nil {
		t.Errorf(": logger factory  expected,[%v], actually: [%v]", " object ", " is null ")
	}

}

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_prop_presets_example(t *testing.T) {

	var lfm loggercom.LoggerFactory

	var lfo = loggerzap.NewLoggerOption()
	lfo.SetProperty(loggerzap.P_PRESETS, loggerzap.PRESETS_EXAMPLE)

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	// --- create logger factory manager
	if lfm == nil {
		t.Errorf(": logger factory  expected,[%v], actually: [%v]", " object ", " is null ")
	}

	logger := lfm.GetLogger()

	logger.Debug("debug message for  example", nil)
	logger.Info("info message for  example", nil)
	logger.Warn("warn message for  example", nil)
	logger.Error("error  message for  example", nil)
}

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Factory_prop_presets_production(t *testing.T) {

	var lfm loggercom.LoggerFactory

	var lfo = loggerzap.NewLoggerOption()
	lfo.SetProperty(loggerzap.P_PRESETS, loggerzap.PRESETS_PROD)

	// use new or struct binding
	// create instance from implement
	lfm = loggercom.NewLoggerFactory(new(loggerzap.ZapFactoryRegister), lfo)

	// --- create logger factory manager
	if lfm == nil {
		t.Errorf(": logger factory  expected,[%v], actually: [%v]", " object ", " is null ")
	}

	logger := lfm.GetLogger()

	logger.Debug("debug message for  prod", nil)
	logger.Info("info message for  prod", nil)
	logger.Warn("warn message for  prod", nil)
	logger.Error("error  message for  prod", nil)

}
