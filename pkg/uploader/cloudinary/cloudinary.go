package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

var cld *cloudinary.Cloudinary

type (
	Cloudinary struct{}
)

func New(cloudinaryURL string) uploader.Uploader {
	cld = connectCloudinary(cloudinaryURL)
	return &Cloudinary{}
}
