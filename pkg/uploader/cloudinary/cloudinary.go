package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

var cld *cloudinary.Cloudinary

type (
	Cloudinary struct{}
)

func New() uploader.Uploader {
	cld = connectCloudinary()
	return &Cloudinary{}
}
