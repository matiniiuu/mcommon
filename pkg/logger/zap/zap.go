package zap

import (
	"github.com/matiniiuu/mcommon/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type log struct {
	zap *zap.SugaredLogger
}

type Options struct {
	logger.Options
	Level zapcore.Level
}

func New(opt *Options) (logger.Logger, error) {
	if opt == nil {
		return nil, logger.ErrNilOption
	}
	writer, err := logger.GetLoggerWriter(&opt.Options)
	if err != nil {
		return nil, err
	}
	ws := zapcore.AddSync(writer)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	enc := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(enc, ws, opt.Level)

	z := zap.New(core)
	sugarLogger := z.Sugar()
	return &log{sugarLogger}, nil
}

func mapToZapFields(m map[string]interface{}) []interface{} {
	keysAndValues := make([]interface{}, 0, len(m)*2)
	for k, v := range m {
		keysAndValues = append(keysAndValues, k, v)
	}
	return keysAndValues
}

func (l *log) Sync() {
	l.zap.Sync()
}

func (l *log) Error(msg string, kv map[string]interface{}) {
	l.zap.Errorw(msg, mapToZapFields(kv)...)
}

func (l *log) Warning(msg string, kv map[string]interface{}) {
	l.zap.Warnw(msg, mapToZapFields(kv)...)
}

func (l *log) Info(msg string, kv map[string]interface{}) {
	l.zap.Infow(msg, mapToZapFields(kv)...)
}
