package derrors

import (
	"github.com/matiniiuu/mcommon/pkg/logger"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
)

func InternalError() error {
	return New(KindUnexpected, messages.GeneralError)
}

func InternalErrorWithLogger(function string, err error, logger logger.Logger) error {
	return NewWithLogger(
		KindUnexpected,
		messages.GeneralError,
		logger,
		function,
		err,
	)
}
