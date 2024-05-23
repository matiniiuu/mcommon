package cloudinary

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
	"github.com/matiniiuu/mcommon/pkg/derrors"
)

func (c *Cloudinary) Upload(file multipart.File, filename, mimetype string) (name string, path string, err error) {
	newFilename := uuid.New().String()
	resp, err := cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{Folder: mimetype, PublicID: newFilename})
	if err != nil {
		return "", "", derrors.InternalError()
	}
	filePath := resp.SecureURL
	return newFilename + "." + resp.Format, filePath, nil
}
