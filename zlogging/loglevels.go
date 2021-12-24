package zlogging

import (
	"sync"

	"go.uber.org/zap/zapcore"
)

//LoggerLevels tracks the logging level of named loggers.
type LoggerLevels struct {
	mutex        sync.RWMutex
	//levelCache   map[string]zapcore.Level
	//specs        map[string]zapcore.Level
	defaultLevel zapcore.Level
	//minLevel     zapcore.Level
}

//// ActivateSpec is used to modify logging levels.
////
//// The logging specification has the following form:
////   [<logger>[,<logger>...]=]<level>[:[<logger>[,<logger>...]=]<level>...]
//func (l *LoggerLevels) ActivateSpec(spec string) error {
//	l.mutex.Lock()
//	defer l.mutex.Unlock()
//
//	defaultLevel := zapcore.InfoLevel
//	specs := map[string]zapcore.Level{}
//	for _, field := range strings.Split(spec, ":") {
//		split := strings.Split(field, "=")
//		switch len(split) {
//		case 1: // level
//			//if field != "" && !IsValidLevel(field) {
//			//	return errors.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
//			//}
//			defaultLevel = NameToLevel(field)
//
//		case 2: // <logger>[,<logger>...]=<level>
//			if split[0] == "" {
//				return errors.Errorf("invalid logging specification '%s': no logger specified in segment '%s'", spec, field)
//			}
//			//if field != "" && !IsValidLevel(split[1]) {
//			//	return errors.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
//			//}
//
//			level := NameToLevel(split[1])
//			loggers := strings.Split(split[0], ",")
//			for _, logger := range loggers {
//				// check if the logger name in the spec is valid. The
//				// trailing period is trimmed as logger names in specs
//				// ending with a period signifies that this part of the
//				// spec refers to the exact logger name (i.e. is not a prefix)
//				//if !isValidLoggerName(strings.TrimSuffix(logger, ".")) {
//				//	return errors.Errorf("invalid logging specification '%s': bad logger name '%s'", spec, logger)
//				//}
//				specs[logger] = level
//			}
//
//		default:
//			return errors.Errorf("invalid logging specification '%s': bad segment '%s'", spec, field)
//		}
//	}
//
//	minLevel := defaultLevel
//	for _, lvl := range specs {
//		if lvl < minLevel {
//			minLevel = lvl
//		}
//	}
//
//	l.minLevel = minLevel
//	l.defaultLevel = defaultLevel
//	l.specs = specs
//	l.levelCache = map[string]zapcore.Level{}
//
//	return nil
//}
//
// Enabled function is an enabled check that evaluates the minimum active logging level.
// It serves as a fast check before the (relatively) expensive Check call in the core.
func (l *LoggerLevels) Enabled(lvl zapcore.Level) bool {
	l.mutex.RLock()
	enabled := l.defaultLevel.Enabled(lvl)
	l.mutex.RUnlock()
	return enabled
}
//
//// Level returns the effective logging level for a logger. If a level has not
//// been explicitly set for the logger, the default logging level will be
//// returned.
//func (l *LoggerLevels) Level(loggerName string) zapcore.Level {
//	if level, ok := l.cachedLevel(loggerName); ok {
//		return level
//	}
//
//	l.mutex.Lock()
//	level := l.calculateLevel(loggerName)
//	l.levelCache[loggerName] = level
//	l.mutex.Unlock()
//
//	return level
//}
//
//// cachedLevel attempts to retrieve the effective log level for a logger from the
//// cache. If the logger is not found, ok will be false.
//func (l *LoggerLevels) cachedLevel(loggerName string) (lvl zapcore.Level, ok bool) {
//	l.mutex.RLock()
//	level, ok := l.levelCache[loggerName]
//	l.mutex.RUnlock()
//	return level, ok
//}
//
//// calculateLevel walks the logger name back to find the appropriate
//// log level from the current spec.
//func (l *LoggerLevels) calculateLevel(loggerName string) zapcore.Level {
//	candidate := loggerName + "."
//	for {
//		if lvl, ok := l.specs[candidate]; ok {
//			return lvl
//		}
//
//		idx := strings.LastIndex(candidate, ".")
//		if idx <= 0 {
//			return l.defaultLevel
//		}
//		candidate = candidate[:idx]
//	}
//}