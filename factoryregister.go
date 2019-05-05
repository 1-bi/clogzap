package log_zap

import (
	"errors"
	loggercom "github.com/1-bi/log-api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

// ZapFactoryRegister
type ZapFactoryRegister struct {
}

func (myself *ZapFactoryRegister) CreateLoggerBean() loggercom.LoggerBean {
	var zapLb = new(zapLoggerBean)
	zapLb.fieldProps = make(map[string]zap.Field, 0)
	return zapLb
}

func (myself *ZapFactoryRegister) CreateLogger(multiopts ...loggercom.Option) (loggercom.Logger, error) {
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

func (myself *ZapFactoryRegister) useOneOption(opts loggercom.Option) (loggercom.Logger, error) {

	var runtimeLevel byte
	if opts.GetLevel() == "" {
		// --- set the default level ---
		runtimeLevel = loggercom.DEVEL_INFO
	} else {

		// set the runtime level
		switch strings.ToUpper(opts.GetLevel()) {
		case "DEBUG":
			runtimeLevel = loggercom.DEVEL_DEBUG
			break
		case "INFO":
			runtimeLevel = loggercom.DEVEL_INFO
			break
		case "WARN":
			runtimeLevel = loggercom.DEVEL_WARN
			break
		case "FATAL":
			runtimeLevel = loggercom.DEVEL_FATAL
			break
		case "ERROR":
			runtimeLevel = loggercom.DEVEL_ERROR
			break
		default:
			return nil, errors.New("Custom Log Level is not predefined from input \"" + opts.GetLevel() + "\". Please choose one from 'debug' , 'info' , 'warn', 'fatal', 'error' .")
		}

	}

	loginst := new(logger)
	loginst.setRuntimeLevel(runtimeLevel)

	var zaplog *zap.Logger
	var err error

	zaplog, err = myself.createZapLogger(opts)
	// --- build new error
	if err != nil {
		return nil, err
	}

	loginst.setZaplogger(zaplog)

	return loginst, err
}

func (myself *ZapFactoryRegister) createZapLogger(opts loggercom.Option) (*zap.Logger, error) {

	// --- check preset key ---
	var customPresets = opts.GetProperties()[P_PRESETS]
	if customPresets == "" {
		// use default ==
		customPresets = PRESETS_DEV
	}

	var err error
	var logInst *zap.Logger

	switch customPresets {
	case PRESETS_EXAMPLE:
		logInst = zap.NewExample()
		break
	case PRESETS_DEV:
		logInst, err = zap.NewDevelopment()
		break
	case PRESETS_PROD:
		logInst, err = zap.NewProduction()
		break
	case PRESETS_NOP:
		logInst = zap.NewNop()
		break
	}

	if logInst != nil {
		return logInst, err
	}

	return nil, err

}

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
		runtimeLevel = loggercom.DEVEL_INFO
	} else {

		// set the runtime level
		switch strings.ToUpper(opt.GetLevel()) {
		case "DEBUG":
			runtimeLevel = loggercom.DEVEL_DEBUG
			break
		case "INFO":
			runtimeLevel = loggercom.DEVEL_INFO
			break
		case "WARN":
			runtimeLevel = loggercom.DEVEL_WARN
			break
		case "FATAL":
			runtimeLevel = loggercom.DEVEL_FATAL
			break
		case "ERROR":
			runtimeLevel = loggercom.DEVEL_ERROR
			break
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

	var zaplog *zap.Logger

	zaplog = zap.New(core)
	defer zaplog.Sync()

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
