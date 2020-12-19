package logger

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
	})
}

// SetLevel set the log level for the central logger
func SetLevel(logLevel logrus.Level) {
	logger.SetLevel(logLevel)
}

// Infoln logs an INFO level event
func Infoln(args ...interface{}) {
	logger.Infoln(args...)
}

// Debugln logs an DEBUG level event
func Debugln(args ...interface{}) {
	logger.Debugln(args...)
}

// Warningln logs an WARNING level event
func Warningln(args ...interface{}) {
	logger.Warningln(args...)
}

// Errorln logs an ERROR level event
func Errorln(args ...interface{}) {
	logger.Errorln(args...)
}

/*
 FORMATTED output functions
*/

// Infof logs an INFO level event
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Debugf logs an DEBUG level event
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Warningf logs an WARNING level event
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Errorf logs an ERROR level event
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}
