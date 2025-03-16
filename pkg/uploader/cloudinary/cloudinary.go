package cloudinary

import (
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/matiniiuu/mcommon/pkg/mconfig"
	"github.com/matiniiuu/mcommon/pkg/uploader"
)

type (
	Cloudinary struct {
		cld *cloudinary.Cloudinary
	}
)

func New(cfg *mconfig.Cloudinary) (uploader.Uploader, error) {
	cld, err := cloudinary.NewFromURL(cfg.Url)
	if err != nil {
		return nil, err
	}
	return &Cloudinary{cld: cld}, nil
}
