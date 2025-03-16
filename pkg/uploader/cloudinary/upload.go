package cloudinary

import (
	"bytes"
	"context"
	"encoding/base64"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (c *Cloudinary) Upload(file multipart.File, folder, filename string) (path string, err error) {
	resp, err := c.cld.Upload.Upload(
		context.Background(),
		file,
		uploader.UploadParams{
			Folder:   folder,
			PublicID: strings.TrimSuffix(filename, filepath.Ext(filename)),
			Format:   filepath.Ext(filename)[1:]})
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}

func (c *Cloudinary) UploadBase24(base64File, folder, filename string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(base64File[strings.IndexByte(base64File, ',')+1:])
	if err != nil {
		return "", err
	}
	resp, err := c.cld.Upload.Upload(
		context.Background(),
		bytes.NewReader(decode),
		uploader.UploadParams{
			Folder:   folder,
			PublicID: strings.TrimSuffix(filename, filepath.Ext(filename)),
			Format:   filepath.Ext(filename)[1:],
		})
	if err != nil {
		return "", err
	}
	return resp.SecureURL, nil
}
