package logzap

import (
	"errors"
	loggercom "github.com/1-bi/log-api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"strings"
)

// ZapFactoryRegister
type ZapFactoryRegister struct {
}

func (myself *ZapFactoryRegister) CreateStructBean() loggercom.StructBean {
	var zapLb = new(zapLoggerBean)
	zapLb.fieldProps = make(map[string]zap.Field)
	return zapLb
}

// CreateLogger add logger name string
func (myself *ZapFactoryRegister) CreateLogger(loggerName string, multiopts ...loggercom.Option) (loggercom.Logger, error) {
	// --- check the current register implement is supported multi options or not.

	// --- check the multiops ---
	switch len(multiopts) {
	case 0:
		break
		//	case 1:
		//		return myself.useOneOption(multiopts[0])
	default:
		return myself.useMultiLoggerOption(multiopts)
	}

	// --- bindiing logger intance ----
	return nil, nil
}

/*

func (myself *ZapFactoryRegister) createZapLogger(opts loggercom.Option) (*zap.Logger, error) {

	// --- check preset key ---
	var customPresets = opts.GetProperties()[PRESETS]
	if customPresets == "" {
		// use default ==
		customPresets = PRESETS_DEV
	}

	var err error
	var logInst *zap.Logger

	switch customPresets {
	case PRESETS_EXAMPLE:
		logInst = zap.NewExample()
	case PRESETS_DEV:
		logInst, err = zap.NewDevelopment()
	case PRESETS_PROD:
		logInst, err = zap.NewProduction()
	case PRESETS_NOP:
		logInst = zap.NewNop()
	}

	if logInst != nil {
		return logInst, err
	}

	return nil, err

}
*/

// useMultiLoggerOption construct method
func (myself *ZapFactoryRegister) useMultiLoggerOption(multiopts []loggercom.Option) (loggercom.Logger, error) {

	// --- generate multiple logger ---
	var multiLogs = make([]loggercom.Logger, 0)

	var logInst loggercom.Logger
	var logInstErr error

	for _, opt := range multiopts {

		logInst, logInstErr = myself.createOneLoggerInstance(opt)

		if logInstErr != nil {
			return nil, errors.New(logInstErr.Error())
		}

		multiLogs = append(multiLogs, logInst)

	}

	return multiLogs[0], nil
}

func (myself *ZapFactoryRegister) createOneLoggerInstance(opt loggercom.Option) (loggercom.Logger, error) {

	// --- define the level  ---
	// set the runtime level
	var runtimeLevel byte
	if opt.GetLevel() == "" {
		// --- set the default level ---
		runtimeLevel = loggercom.INFO
	} else {

		// set the runtime level
		switch strings.ToUpper(opt.GetLevel()) {
		case "DEBUG":
			runtimeLevel = loggercom.DEBUG
		case "INFO":
			runtimeLevel = loggercom.INFO
		case "WARN":
			runtimeLevel = loggercom.WARN
		case "FATAL":
			runtimeLevel = loggercom.FATAL
		case "ERROR":
			runtimeLevel = loggercom.ERROR
		}
	}

	var level = levelEventFilter(runtimeLevel)

	// --- get and check appender

	var appenderMap = opt.GetAppenders()

	multiCores := make([]zapcore.Core, 0)

	for _, a := range appenderMap {

		var cores, err = myself.createZapCores(level, a.(zapAppender))

		if err != nil {
			return nil, err
		}
		// add stdout output
		multiCores = append(multiCores, cores[0])

		// add stderr output
		//multiCores = append(multiCores, cores[1])

	}

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	//kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.

	core := zapcore.NewTee(multiCores...)

	var zaplog = zap.New(core)

	defer func() {
		err := zaplog.Sync()
		if err != nil {
			log.Println(err)
		}
	}()

	loginst := new(logger)
	loginst.setZaplogger(zaplog)

	return loginst, nil

}

func (myself *ZapFactoryRegister) createZapCores(level zapcore.LevelEnabler, appender zapAppender) ([]zapcore.Core, error) {

	// --- check the layout --
	var layout, ok = appender.ZapLayout().(ZapLayout)

	if !ok {
		// --- use defulat layout ---
		return nil, errors.New(" Layout inputed is not suitable for \"ZapLayout\". Please check the layout class. ")
	}

	var cores = make([]zapcore.Core, 0)

	var outputCore = zapcore.NewCore(layout.BuildEncoder(), appender.Output(), level)

	var errorCore = zapcore.NewCore(layout.BuildEncoder(), appender.Error(), level)

	cores = append(cores, outputCore)
	cores = append(cores, errorCore)

	return cores, nil
}
