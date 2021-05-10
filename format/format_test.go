package format

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewFormat(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(new(LogFormatter))
	log.SetLevel(logrus.TraceLevel)
	log.SetReportCaller(true)
	log.Out = os.Stdout

	log.Error("error test")
	log.Debug("debug test")
	log.Debugln("debug test")
	log.Info("info test")
	log.Infof("info %s\n", "test")
	log.Warn("warn test")
}
