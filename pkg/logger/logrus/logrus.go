package logrus

import (
	"github.com/matiniiuu/mcommon/pkg/logger"
	"github.com/sirupsen/logrus"
)

type log struct {
	logger *logrus.Logger
}
type Options struct {
	logger.Options
}

// New is constructor of the logrus package
func New(opt *Options) (logger.Logger, error) {

	if opt == nil {
		return nil, logger.ErrNilOption
	}
	writer, err := logger.GetLoggerWriter(&opt.Options)
	if err != nil {
		return nil, err
	}
	l := &log{logger: logrus.New()}
	l.logger.SetOutput(writer)
	l.logger.SetFormatter(&logrus.JSONFormatter{})

	return l, nil
}

func (l *log) Sync() {}

// Info is logger with level info
func (l *log) Info(msg string, kv map[string]interface{}) {
	l.logger.WithFields(kv).Info(msg)
}

// Warning is logger with level warning
func (l *log) Warning(msg string, kv map[string]interface{}) {
	l.logger.WithFields(kv).Warning(msg)
}

// Error is logger with level error
func (l *log) Error(msg string, kv map[string]interface{}) {
	l.logger.WithFields(kv).Error(msg)
}
