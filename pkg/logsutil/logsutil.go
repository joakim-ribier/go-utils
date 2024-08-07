package logsutil

import (
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap"
)

type Logger struct {
	log logr.Logger

	Info  func(msg string, keysAndValues ...any)
	Error func(err error, msg string, keysAndValues ...any)
}

// WithValues returns a new Logger instance with additional key/value pairs
func (l Logger) WithValues(keysAndValues ...any) Logger {
	log := l.log.WithValues(keysAndValues...)
	return Logger{
		log:   log,
		Info:  log.Info,
		Error: log.Error,
	}
}

// Namespace adds key/value pair namespace:{value}
func (l Logger) Namespace(value string) Logger {
	return l.WithValues("namespace", value)
}

// NewLogger builds and returns its own log struct which contains
// an implementation of {zap.SugaredLogger} library
func NewLogger(fileLog, appName string) (*Logger, error) {
	config := zap.NewDevelopmentConfig()

	config.OutputPaths = []string{fileLog}
	zapLog, err := config.Build()
	if err != nil {
		return nil, err
	}

	log := zapr.NewLogger(zapLog).WithValues("app", appName)

	return &Logger{
		log:   log,
		Info:  log.Info,
		Error: log.Error,
	}, nil
}
