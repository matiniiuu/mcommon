package logger

import (
	"io"
	"path/filepath"

	"github.com/alecthomas/units"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/xhit/go-str2duration/v2"
)

// GetLoggerWriter return io.Writer which can create different
// files with custom names at different time intervals
func GetLoggerWriter(opt *Options) (io.Writer, error) {
	maxAge, err := str2duration.ParseDuration(opt.MaxAge)
	if err != nil {
		return nil, err
	}

	rotationTime, err := str2duration.ParseDuration(opt.RotationTime)
	if err != nil {
		return nil, err
	}

	rotationSize, err := units.ParseBase2Bytes(opt.RotationSize)
	if err != nil {
		return nil, err
	}

	return rotators.New(
		filepath.Join(opt.Path, opt.Pattern),
		rotators.WithMaxAge(maxAge),
		rotators.WithRotationTime(rotationTime),
		rotators.WithRotationSize(int64(rotationSize)),
	)
}
