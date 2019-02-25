package logger

import (
	logging "github.com/op/go-logging"
	"os"
)

type ILogger interface {
	Critical(format string, args ...interface{})
	Error(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Debug(format string, args ...interface{})
}

func NewLogger(logLevel string) ILogger {
	moduleNameLog := "main"
	var log = logging.MustGetLogger(moduleNameLog)
	var format = logging.MustStringFormatter(`{"date": "%{time:2006-01-02 15:04:05.000}", "level": "%{level:.4s}", "id": "%{id:03x}", "shortfunc": "%{shortfunc}", "message": "%{message}"}`)
	backendLogger := logging.NewLogBackend(os.Stderr, "", 0)
	backendLeveled := logging.AddModuleLevel(logging.NewBackendFormatter(backendLogger, format))

	var level logging.Level
	switch logLevel {
	case "debug":
		level = logging.DEBUG
	case "info":
		level = logging.INFO
	case "notice":
		level = logging.NOTICE
	case "warning":
		level = logging.WARNING
	case "error":
		level = logging.ERROR
	case "critical":
		level = logging.CRITICAL
	default:
		level = logging.WARNING
	}

	backendLeveled.SetLevel(level, moduleNameLog)
	logging.SetBackend(backendLeveled)

	return log
}
