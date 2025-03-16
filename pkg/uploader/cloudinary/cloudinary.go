package cloudinary

import (
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

type (
	Cloudinary struct {
		cld *cloudinary.Cloudinary
	}
)

func New(cloudinaryURL string) uploader.Uploader {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalln(err)
	}
	return &Cloudinary{cld: cld}
}
