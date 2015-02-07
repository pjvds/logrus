package logrus

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stdout`. You can also set this to
	// something more adventorous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks levelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter
	// The logging level the logger should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged. `logrus.Debug` is useful in
	Level Level
	// Used to sync writing to the log.
	mu sync.Mutex
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(JSONFormatter),
//      Hooks: make(levelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stdout,
		Formatter: new(TextFormatter),
		Hooks:     make(levelHooks),
		Level:     InfoLevel,
	}
}

// Reports whether log level is at least debug level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
// when the log level isn't at least debug.
func (logger *Logger) IsDebug() bool {
	return logger.Level >= DebugLevel
}

// Reports whether log level is at least info level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
// when the log level isn't at least debug.
func (logger *Logger) IsInfo() bool {
	return logger.Level >= InfoLevel
}

// Reports whether log level is at least warning level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
// when the log level isn't at least debug.
func (logger *Logger) IsWarn() bool {
	return logger.Level >= WarnLevel
}

// Reports whether log level is at least error level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
func (logger *Logger) IsError() bool {
	return logger.Level >= ErrorLevel
}

// Reports whether log level is at least fatal level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
func (logger *Logger) IsFatal() bool {
	return logger.Level >= FatalLevel
}

// Reports whether log level is at least panic level.
//
// This method can be used to prevent evaluation of arguments if
// the log level isn't high enough:
//
//      if log.IsDebug() {
//           log.Debugf("π = %v", calculatePi())
//      }
//
// In this case the method `calculatePi` will not be evaluated
func (logger *Logger) IsPanic() bool {
	return logger.Level >= PanicLevel
}

// Adds a field to the log entry, note that you it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// Ff you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	return NewEntry(logger).WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	return NewEntry(logger).WithFields(fields)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	NewEntry(logger).Debugf(format, args...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	NewEntry(logger).Infof(format, args...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	NewEntry(logger).Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	NewEntry(logger).Warnf(format, args...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	NewEntry(logger).Warnf(format, args...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	NewEntry(logger).Errorf(format, args...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	NewEntry(logger).Fatalf(format, args...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	NewEntry(logger).Panicf(format, args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	NewEntry(logger).Debug(args...)
}

func (logger *Logger) Info(args ...interface{}) {
	NewEntry(logger).Info(args...)
}

func (logger *Logger) Print(args ...interface{}) {
	NewEntry(logger).Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	NewEntry(logger).Warn(args...)
}

func (logger *Logger) Warning(args ...interface{}) {
	NewEntry(logger).Warn(args...)
}

func (logger *Logger) Error(args ...interface{}) {
	NewEntry(logger).Error(args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	NewEntry(logger).Fatal(args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	NewEntry(logger).Panic(args...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	NewEntry(logger).Debugln(args...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	NewEntry(logger).Infoln(args...)
}

func (logger *Logger) Println(args ...interface{}) {
	NewEntry(logger).Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	NewEntry(logger).Warnln(args...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	NewEntry(logger).Warnln(args...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	NewEntry(logger).Errorln(args...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	NewEntry(logger).Fatalln(args...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	NewEntry(logger).Panicln(args...)
}
