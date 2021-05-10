package format

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

type LogFormatter struct{}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	var file, function string
	var len int
	level := entry.Level.String()
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		len = entry.Caller.Line
		function = entry.Caller.Function
	}

	prefix := fmt.Sprintf("%s [%s] %s:%d %s", timestamp, strings.ToUpper(level), file, len, function)
	red := color.New(color.FgHiRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	switch strings.ToUpper(level) {
	case "DEBUG", "TRACE":
		return []byte(fmt.Sprintf("%s: %s\n", green(prefix), entry.Message)), nil
	case "WARN", "WARNING":
		return []byte(fmt.Sprintf("%s: %s\n", yellow(prefix), entry.Message)), nil
	case "ERROR", "PANIC", "FATAL":
		return []byte(fmt.Sprintf("%s: %s\n", red(prefix), entry.Message)), nil
	default:
		return []byte(fmt.Sprintf("%s: %s\n", cyan(prefix), entry.Message)), nil
	}

}
