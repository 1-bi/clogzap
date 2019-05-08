package test

import (
	"go.uber.org/zap"
	"testing"
	"time"
)

//  Test_BasicCase1_Debug define bug info
func Test_Zap_Dev(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()
	var url = "http://www.moco.com"

	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

	logger = zap.NewExample()
	defer logger.Sync()

	logger.With(
		zap.String("loggerName", "test"),
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	).Info("tracked some metrics")

}
