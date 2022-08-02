package log

import (
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	logrus.SetLevel(logrus.InfoLevel)
	log = logrus.New()
}

func WithBot(name string) *logrus.Entry {
	return log.WithField("bot_name", name)
}

func Info(args ...interface{}) {
	log.Info(args)
}

func Warn(args ...interface{}) {
	log.Warn(args)
}

func Debug(args ...interface{}) {
	log.Debug(args)
}

func Fatal(args ...interface{}) {
	log.Fatal(args)
}

func Error(args ...interface{}) {
	log.Error(args)
}
