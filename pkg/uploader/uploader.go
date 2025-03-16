package uploader

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/matiniiuu/mcommon/pkg/derrors"
)

type (
	Uploader interface {
		Upload(file multipart.File, folder, filename string) (path string, err error)
		UploadBase24(base64File, folder, filename string) (path string, err error)
	}
)

func GenerateFileName(originalFilename string) (string, error) {
	return fmt.Sprintf("%s%s", uuid.New().String(), filepath.Ext(originalFilename)), nil
}

func GetMimetype(file multipart.File) (string, error) {
	buff := make([]byte, 512)
	_, err := file.Read(buff)

	if err != nil {
		return "", derrors.InternalError()
	}
	return strings.Split(http.DetectContentType(buff), "/")[0], nil
}
