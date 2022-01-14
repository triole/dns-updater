package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logging holds the logging module
type Logging struct {
	Logrus    *logrus.Logger
	LogToFile bool
}

// Init method, does what it says
func initLogging(logFile string) (lg Logging) {
	timeStampFormat := "2006-01-02 15:04:05.000 MST"
	lg.Logrus = logrus.New()

	lg.Logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeStampFormat,
		DisableQuote:    true,
		PadLevelText:    true,
	})

	openLogFile, err := os.OpenFile(
		logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
	)
	if err != nil {
		lg.LogFatal(
			"Can not open log file",
			logrus.Fields{
				"logfile": logFile,
				"error":   err.Error(),
			},
		)
	}

	if logFile != "/dev/stdout" {
		lg.LogToFile = true
	}

	mw := io.MultiWriter(os.Stdout, openLogFile)
	logrus.SetOutput(mw)
	return lg
}

// LogInfo logs an info message
func (lg Logging) LogInfo(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Info(msg)
	default:
		lg.Logrus.Info(msg)
	}
}

// LogWarn logs a warning
func (lg Logging) LogWarn(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Warning(msg)
	default:
		lg.Logrus.Warning(msg)
	}
}

// LogFatal logs fatal and exits
func (lg Logging) LogFatal(msg string, fields interface{}) {
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Fatal(msg)
	default:
		lg.Logrus.Fatal(msg)
	}
}

// LogError logs an error message
func (lg Logging) LogError(msg interface{}, fields interface{}) {
	var msgStr string
	switch val := msg.(type) {
	case error:
		msgStr = val.Error()
	default:
		msgStr = val.(string)
	}
	switch val := fields.(type) {
	case logrus.Fields:
		lg.Logrus.WithFields(val).Error(msgStr)
	default:
		lg.Logrus.Error(msgStr)
	}
}
