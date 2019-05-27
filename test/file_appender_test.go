package test

import (
	"github.com/1-bi/log-zap/test/fixture"
	"github.com/smartystreets/gunit"
	"testing"
)

//  Test_BasicCase1_Debug define bug info
func Test_FileAppender(t *testing.T) {
	gunit.Run(new(fixture.FileAppenderFixTure), t)
}
