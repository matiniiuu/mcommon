package logger

import (
	"errors"
)

var ErrNilOption = errors.New("option can not be nil")

type (
	Options struct {
		Path, Pattern, MaxAge, RotationTime, RotationSize string
	}
	Logger interface {
		Info(string, map[string]interface{})
		Warning(string, map[string]interface{})
		Error(string, map[string]interface{})
		Sync()
	}
)
