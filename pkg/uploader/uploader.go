package uploader

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/matiniiuu/mcommon/pkg/derrors"
)

type (
	Uploader interface {
		Upload(file multipart.File, filename, mimetype string) (name string, path string, err error)
	}
)

func GetMimetype(file multipart.File) (string, error) {
	buff := make([]byte, 512)
	_, err := file.Read(buff)

	if err != nil {
		return "", derrors.InternalError()
	}
	return strings.Split(http.DetectContentType(buff), "/")[0], nil
}
