package zlogging

import "go.uber.org/zap"

var Global *Logging

func init() {
	logging, err := New(Config{})
	if err != nil {
		panic(err)
	}

	Global = logging
}

func LoggerLevel() string {
	return Global.defaultLevel.String()
}

// MustGetLogger creates a logger with the specified name. If an invalid name
// is provided, the operation will panic.
func MustGetLogger(loggerName string) *zap.SugaredLogger {
	return Global.Logger(loggerName)
}

func SetLogLevel(level string) {
	Global.defaultLevel = NameToLevel(level)
	return
}