package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Logging holds the logging module
type Logging struct {
	Logrus *logrus.Logger
}

// Init method, does what it says
func initLogging(logFile string) (lg Logging) {
	timeStampFormat := "2006-01-02 15:04:05.000 MST"
	lg.Logrus = logrus.New()

	lg.Logrus.SetLevel(logrus.InfoLevel)
	if CLI.Debug {
		lg.Logrus.SetLevel(logrus.DebugLevel)
	}

	lg.Logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "date",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyMsg:   "msg",
		},
		TimestampFormat: timeStampFormat,
	})

	lg.Logrus.SetOutput(os.Stdout)
	openLogFile, err := os.OpenFile(
		logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644,
	)
	lg.LogIfFileError("open", logFile, err, true)
	mw := io.MultiWriter(os.Stdout, openLogFile)

	lg.Logrus.SetOutput(mw)
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

func (lg Logging) LogIfError(err error, msg interface{}, fields interface{}) {
	if err != nil {
		lg.LogError(msg, fields)
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

func (lg Logging) LogStatus(msg string, ipd tIPData, forceUpdate bool) {
	lg.LogInfo(
		msg, logrus.Fields{
			"ip_data":      ipd,
			"force_update": forceUpdate,
		},
	)
}

func (lg Logging) LogIfFileError(msg, filename string, err error, fatal bool) {
	if err != nil {
		lg.LogError(msg+" file error", logrus.Fields{
			"err":      err.Error(),
			"filename": filename,
		})
		if fatal {
			os.Exit(1)
		}
	}
}

// LogDebug logs a debug message
// func (lg Logging) LogDebug(msg string, conf tConf) {
// 	if conf.Debug {
// 		conf.Token = "{{.TOKEN}}"
// 		conf.Hostname = "{{.HOSTNAME}}"
// 		conf.RequestHeaders = map[string]string{"map": "dummy"}
// 		lg.Logrus.WithFields(logrus.Fields{
// 			"all_data": conf,
// 		}).Debug(msg)
// 	}
// }
