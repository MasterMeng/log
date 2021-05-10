package hook

import "github.com/sirupsen/logrus"

type ExampleHook struct{}

func (h *ExampleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (h *ExampleHook) Fire(entry *logrus.Entry) error {
	entry.Data["hook"] = "ExampleHook"
	return nil
}

func NewHook() logrus.Hook {
	return &ExampleHook{}
}
