package benchmark

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Benchmark_Zap_orginal(b *testing.B) {
	//b.StopTimer()

	// First, define our level-handling logic.
	/*
		highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.InfoLevel
		})
	*/
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	// Assume that we have clients for two Kafka topics. The clients implement
	// zapcore.WriteSyncer and are safe for concurrent use. (If they only
	// implement io.Writer, we can use zapcore.AddSync to add a no-op Sync
	// method. If they're not safe for concurrent use, we can add a protecting
	// mutex with zapcore.Lock.)
	//topicDebugging := zapcore.AddSync(ioutil.Discard)
	//topicErrors := zapcore.AddSync(ioutil.Discard)

	// High-priority output should also go to standard error, and low-priority
	// output should also go to standard out.
	consoleDebugging := zapcore.Lock(os.Stdout)
	//consoleErrors := zapcore.Lock(os.Stderr)

	// Optimize the Kafka output for machine consumption and the console output
	// for human operators.
	//kafkaEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// Join the outputs, encoders, and level-handling functions into
	// zapcore.Cores, then tee the four cores together.
	core := zapcore.NewTee(
		//zapcore.NewCore(kafkaEncoder, topicErrors, highPriority),
		//zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
		//zapcore.NewCore(kafkaEncoder, topicDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	logger := zap.New(core)
	defer logger.Sync()
	//logger.Info("constructed a logger")
	//logger.Warn("constructed a logger waring ")
	//logger.Error("constructed a logger error ")

	//logger.Debug("debug message for  example", nil)
	//b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("constructed a logger info ")
		//logger.Warn("constructed a logger waring ")
		//logger.Warn("constructed a logger error ")
	}
	//logger.Warn("warn message for  example", nil)
	//logger.Error("error  message for  example", nil)

}
