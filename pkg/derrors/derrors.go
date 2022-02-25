package derrors

import (
	"errors"
	"net/http"

	"github.com/matiniiuu/mcommon/pkg/logger"
	"github.com/matiniiuu/mcommon/pkg/translator/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	kind uint

	serverError struct {
		kind    kind
		message string
	}
)

const (
	_ kind = iota
	KindInvalid
	KindNotFound
	KindUnauthorized
	KindUnexpected
	KindNotAllowed
	KindForbidden
)

var (
	httpErrors = map[kind]int{
		KindInvalid:      http.StatusBadRequest,
		KindNotFound:     http.StatusNotFound,
		KindUnauthorized: http.StatusUnauthorized,
		KindUnexpected:   http.StatusInternalServerError,
		KindNotAllowed:   http.StatusMethodNotAllowed,
		KindForbidden:    http.StatusForbidden,
	}
	grpcErrors = map[codes.Code]int{
		codes.InvalidArgument:  http.StatusBadRequest,
		codes.NotFound:         http.StatusNotFound,
		codes.Unauthenticated:  http.StatusUnauthorized,
		codes.Internal:         http.StatusInternalServerError,
		codes.PermissionDenied: http.StatusBadRequest,
	}
)

//New is constructor of the errors package
func New(kind kind, msg string) error {
	return serverError{
		kind:    kind,
		message: msg,
	}
}
func NewGrpcError(kind kind, msg string) error {
	httpCode := httpErrors[kind]
	for key, value := range grpcErrors {
		if value == httpCode {
			return status.New(key, msg).Err()
		}
	}
	return status.New(codes.Internal, msg).Err()
}

func NewWithLogger(kind kind, msg string, logger logger.Logger, function string, err error) error {
	logger.Error(err.Error(), map[string]interface{}{
		"Function":        function,
		"ResponseMessage": msg,
	})
	return New(kind, msg)
}

//Error return message of error
func (e serverError) Error() string {
	return e.message
}

//HttpError convert kind of error to Http status error
//if error type is not serverError return 400 status code
func HttpError(err error) (string, int) {
	var serverErr serverError

	ok := errors.As(err, &serverErr)
	if !ok {
		return messages.GeneralError, http.StatusInternalServerError
	}

	code, ok := httpErrors[serverErr.kind]
	if !ok {
		return serverErr.message, http.StatusBadRequest
	}

	return serverErr.message, code

}

//GrpcError convert kind of error to Http status error
//if error type is not serverError return 400 status code
func ConvertGrpcErrorToHttpError(err error) (string, int) {
	gError, ok := status.FromError(err)
	if !ok {
		return messages.GeneralError, http.StatusInternalServerError
	}
	code, ok := grpcErrors[gError.Code()]
	if !ok {
		return gError.Message(), http.StatusBadRequest
	}
	return gError.Message(), code
}

func As(err error) bool {
	return errors.As(err, &serverError{})
}
